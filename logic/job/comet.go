package main

import (
	log "code.google.com/p/log4go"
	"github.com/Terry-Mao/goim/define"
	inet "github.com/Terry-Mao/goim/libs/net"
	cproto "github.com/Terry-Mao/goim/proto/comet"
	"github.com/Terry-Mao/protorpc"
	"time"
)

var (
	cometServiceMap = make(map[int32]**protorpc.Client)
)

const (
	CometService          = "PushRPC"
	CometServicePing      = "PushRPC.Ping"
	CometServicePushMsg   = "PushRPC.PushMsg"
	CometServicePushMsgs  = "PushRPC.PushMsgs"
	CometServiceMPushMsg  = "PushRPC.MPushMsg"
	CometServiceMPushMsgs = "PushRPC.MPushMsgs"
	CometServiceBroadcast = "PushRPC.Broadcast"
)

func InitComet(addrs map[int32]string) (err error) {
	for serverID, addrs := range addrs {
		var (
			rpcClient     *protorpc.Client
			quit          chan struct{}
			network, addr string
		)
		if network, addr, err = inet.ParseNetwork(addrs); err != nil {
			log.Error("inet.ParseNetwork() error(%v)", err)
			return
		}
		if rpcClient, err = protorpc.Dial(network, addr); err != nil {
			log.Error("protorpc.Dial(\"%s\") error(%s)", addr, err)
			return
		}
		go protorpc.Reconnect(&rpcClient, quit, network, addr)
		log.Info("rpc addr:%s connected", addr)
		cometServiceMap[serverID] = &rpcClient
	}
	return
}

// get comet server client by server id
func getCometByServerId(serverID int32) (*protorpc.Client, error) {
	if client, ok := cometServiceMap[serverID]; !ok || *client == nil {
		return nil, ErrComet
	} else {
		return *client, nil
	}
}

func mpushComet(c *protorpc.Client, serverId int32, subkeys []string, body []byte) {
	var (
		now  = time.Now()
		args = &cproto.MPushMsgArg{Keys: subkeys, Operation: define.OP_SEND_SMS_REPLY, Msg: body}
		rep  = &cproto.MPushMsgReply{}
		err  error
	)
	if err = c.Call(CometServiceMPushMsg, args, rep); err != nil {
		log.Error("c.Call(\"%s\", %v, reply) error(%v)", CometServiceMPushMsg, *args, err)
	} else {
		log.Info("push msg to serverId:%d index:%d(%f)", serverId, rep.Index, time.Now().Sub(now).Seconds())
	}
}

func broadcastComet(c *protorpc.Client, serverId int32, msg []byte) {
	var (
		now  = time.Now()
		args = &cproto.BoardcastArg{Ver: 0, Operation: define.OP_SEND_SMS_REPLY, Msg: msg}
		err  error
	)
	if err = c.Call(CometServiceBroadcast, args, nil); err != nil {
		log.Error("c.Call(\"%s\", %v, reply) error(%v)", CometServiceBroadcast, *args, err)
	} else {
		log.Info("broadcast msg to serverId:%d msg:%s(%f)", serverId, msg, time.Now().Sub(now).Seconds())
	}
}
