package v2

import (
	"fmt"

	"github.com/wselfjes/bitget/internal"
	"github.com/wselfjes/bitget/internal/common"
	"github.com/wselfjes/bitget/pkg/model/v2/response"
	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/types/sJson"
)

type GetAllAccountBalance struct {
	c common.InterRestClient
}

func NewGetAllAccountBalance(c common.InterRestClient) (*GetAllAccountBalance, error) {
	return &GetAllAccountBalance{
		c: c,
	}, nil
}

func (p *GetAllAccountBalance) Do() (*response.AccountsResp, error) {

	uri := "/api/v2/account/all-account-balance"

	resp, err := p.c.DoGet(uri, internal.NewParams())
	if err != nil {
		return nil, err
	}

	sLog.Warn("resp: %s", resp)

	respObj := new(response.AccountsResp)

	err = sJson.ToObj(resp, respObj)
	if err != nil {
		return nil, fmt.Errorf("sJson.ToObj(): %s", err)
	}

	return respObj, nil
}
