package ws

import (
	"github.com/wselfjes/bitget/config"
	"github.com/wselfjes/bitget/constants"
	"github.com/wselfjes/bitget/internal/common"
	"github.com/wselfjes/bitget/internal/model"
)

type BitgetWsClient struct {
	bitgetBaseWsClient *common.BitgetBaseWsClient
}

func (p *BitgetWsClient) Init(needLogin bool, listener common.OnReceive, errorListener common.OnReceive, creds config.InterApiCreds) *BitgetWsClient {
	p.bitgetBaseWsClient = new(common.BitgetBaseWsClient).Init(needLogin, creds)
	p.bitgetBaseWsClient.SetListener(listener, errorListener)
	p.bitgetBaseWsClient.ConnectWebSocket()
	p.bitgetBaseWsClient.StartReadLoop()
	p.bitgetBaseWsClient.ExecuterPing()
	p.bitgetBaseWsClient.StartTickerLoop()
	p.bitgetBaseWsClient.Login()

	return p
}

func (p *BitgetWsClient) Connect() *BitgetWsClient {
	p.bitgetBaseWsClient.Connect()
	return p
}

func (p *BitgetWsClient) UnSubscribe(list []model.SubscribeReq) {

	var args []interface{}
	for i := 0; i < len(list); i++ {
		delete(p.bitgetBaseWsClient.ListenerMap, list[i])
		p.bitgetBaseWsClient.AllSubscribe.Add(list[i])
		p.bitgetBaseWsClient.AllSubscribe.Remove(list[i])
		args = append(args, list[i])
	}

	wsBaseReq := model.WsBaseReq{
		Op:   constants.WsOpUnsubscribe,
		Args: args,
	}

	p.SendMessageByType(wsBaseReq)
}

func (p *BitgetWsClient) SubscribeDef(list []model.SubscribeReq) {

	var args []interface{}
	for i := 0; i < len(list); i++ {
		req := list[i]
		args = append(args, req)
	}
	wsBaseReq := model.WsBaseReq{
		Op:   constants.WsOpSubscribe,
		Args: args,
	}

	p.SendMessageByType(wsBaseReq)
}

func (p *BitgetWsClient) Subscribe(list []model.SubscribeReq, listener common.OnReceive) {

	var args []interface{}
	for i := 0; i < len(list); i++ {
		req := list[i]
		args = append(args, req)

		p.bitgetBaseWsClient.ListenerMap[req] = listener
		p.bitgetBaseWsClient.AllSubscribe.Add(req)
	}

	wsBaseReq := model.WsBaseReq{
		Op:   constants.WsOpSubscribe,
		Args: args,
	}
	p.SendMessageByType(wsBaseReq)
}

func (p *BitgetWsClient) SendMessage(msg string) {
	p.bitgetBaseWsClient.Send(msg)
}

func (p *BitgetWsClient) SendMessageByType(req model.WsBaseReq) {
	p.bitgetBaseWsClient.SendByType(req)
}

func (p *BitgetWsClient) Close() {
	p.bitgetBaseWsClient.DisconnectWebSocket()
}
