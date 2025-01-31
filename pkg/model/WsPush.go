package model

import "github.com/wselfjes/bitget/internal/model"

type BasePush struct {
	Action string             `json:"action"`
	Arg    model.SubscribeReq `json:"arg"`
	// Data   []interface{} `json:"data"`
}
