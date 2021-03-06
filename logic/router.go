package main

import (
	log "code.google.com/p/log4go"
	"github.com/Terry-Mao/goim/libs/hash/ketama"
	inet "github.com/Terry-Mao/goim/libs/net"
	rproto "github.com/Terry-Mao/goim/proto/router"
	rpc "github.com/Terry-Mao/protorpc"
	"strconv"
)

var (
	routerServiceMap = map[string]**rpc.Client{}
	routerRing       *ketama.HashRing
)

const (
	routerService           = "RouterRPC"
	routerServiceConnect    = "RouterRPC.Connect"
	routerServiceDisconnect = "RouterRPC.Disconnect"
	routerServiceMGet       = "RouterRPC.MGet"
	routerServiceGetAll     = "RouterRPC.GetAll"
)

func InitRouter() (err error) {
	var (
		network, addr string
	)
	routerRing = ketama.NewRing(ketama.Base)
	for serverId, addrs := range Conf.RouterRPCAddrs {
		// WARN r must every recycle changed for reconnect
		var (
			r          *rpc.Client
			routerQuit = make(chan struct{}, 1)
		)
		if network, addr, err = inet.ParseNetwork(addrs); err != nil {
			log.Error("inet.ParseNetwork() error(%v)", err)
			return
		}
		r, err = rpc.Dial(network, addr)
		if err != nil {
			log.Error("rpc.Dial(\"%s\", \"%s\") error(%s)", network, addr, err)
		}
		go rpc.Reconnect(&r, routerQuit, network, addr)
		log.Debug("router rpc addr:%s connect", addr)
		routerServiceMap[serverId] = &r
		routerRing.AddNode(serverId, 1)
	}
	routerRing.Bake()
	return
}

func getRouters() map[string]**rpc.Client {
	return routerServiceMap
}

func getRouterByServer(server string) (*rpc.Client, error) {
	if client, ok := routerServiceMap[server]; !ok || *client == nil {
		return nil, ErrRouter
	} else {
		return *client, nil
	}
}

func getRouterByUID(userID int64) (*rpc.Client, error) {
	return getRouterByServer(routerRing.Hash(strconv.FormatInt(userID, 10)))
}

func getRouterNode(userID int64) string {
	return routerRing.Hash(strconv.FormatInt(userID, 10))
}

func connect(userID int64, server int32) (seq int32, err error) {
	var client *rpc.Client
	if client, err = getRouterByUID(userID); err != nil {
		return
	}
	arg := &rproto.ConnArg{UserId: userID, Server: server}
	reply := &rproto.ConnReply{}
	if err = client.Call(routerServiceConnect, arg, reply); err != nil {
		log.Error("c.Call(\"%s\",\"%v\") error(%s)", routerServiceConnect, arg, err)
	} else {
		seq = reply.Seq
	}
	return
}

func disconnect(userID int64, seq int32) (has bool, err error) {
	var client *rpc.Client
	if client, err = getRouterByUID(userID); err != nil {
		return
	}
	arg := &rproto.DisconnArg{UserId: userID, Seq: seq}
	reply := &rproto.DisconnReply{}
	if err = client.Call(routerServiceDisconnect, arg, reply); err != nil {
		log.Error("c.Call(\"%s\",\"%v\") error(%s)", routerServiceDisconnect, *arg, err)
	} else {
		has = reply.Has
	}
	return
}

func getSubkeys(serverId string, userIds []int64) (reply *rproto.MGetReply, err error) {
	var client *rpc.Client
	if client, err = getRouterByServer(serverId); err != nil {
		return
	}
	arg := &rproto.MGetArg{UserIds: userIds}
	reply = &rproto.MGetReply{}
	if err = client.Call(routerServiceMGet, arg, reply); err != nil {
		log.Error("client.Call(\"%s\",\"%v\") error(%s)", routerServiceMGet, arg, err)
	}
	return
}

func divideToRouter(userIds []int64) (divide map[int32][]string, err error) {
	var (
		i, j         int
		node, subkey string
		subkeys      []string
		reply        *rproto.MGetReply
		server       int32
		session      *rproto.GetReply
		uid          int64
		ids          []int64
		ok           bool
		m            = make(map[string][]int64)
	)
	divide = make(map[int32][]string) //map[comet.serverId][]subkey
	for i = 0; i < len(userIds); i++ {
		node = getRouterNode(userIds[i])
		if ids, ok = m[node]; !ok {
			ids = []int64{userIds[i]}
		} else {
			ids = append(ids, userIds[i])
		}
		m[node] = ids
	}
	// TODO muti-routine get
	for node, ids = range m {
		if reply, err = getSubkeys(node, ids); err != nil {
			log.Error("getSubkeys(\"%s\") error(%s)", node, err)
			return
		}
		for j = 0; j < len(reply.UserIds); j++ {
			session = reply.Sessions[j]
			uid = reply.UserIds[j]
			for i = 0; i < len(session.Seqs); i++ {
				subkey = encode(uid, session.Seqs[i])
				server = session.Servers[i]
				if subkeys, ok = divide[server]; !ok {
					subkeys = []string{subkey}
				} else {
					subkeys = append(subkeys, subkey)
				}
				divide[server] = subkeys
			}
		}
	}
	return
}
