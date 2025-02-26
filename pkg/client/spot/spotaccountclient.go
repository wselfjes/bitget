package spot

import (
	"github.com/wselfjes/bitget/constants"
	"github.com/wselfjes/bitget/internal"
	"github.com/wselfjes/bitget/internal/common"
	"github.com/wselfjes/bitget/pkg/model/spot/account"
)

type SpotAccountClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *SpotAccountClient) Init() *SpotAccountClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init(nil)
	return p
}

/**
 * Obtain account assets
 * @return ResponseResult
 */
func (p *SpotAccountClient) Assets() (string, error) {

	params := internal.NewParams()

	uri := constants.SpotAccount + "/assets"

	resp, err := p.BitgetRestClient.DoGet(uri, params)

	return resp, err

}

/**
 * Obtain transfer records
 * @param coinId
 * @param fromType
 * @param limit
 * @param after
 * @param before
 * @return ResponseResult
 */
func (p *SpotAccountClient) TransferRecords(coinId string, fromType string, limit string, after string, before string) (string, error) {

	params := internal.NewParams()
	if len(coinId) > 0 {
		params["coinId"] = coinId
	}
	if len(fromType) > 0 {
		params["fromType"] = fromType
	}
	if len(limit) > 0 {
		params["limit"] = limit
	}
	if len(after) > 0 {
		params["after"] = after
	}
	if len(before) > 0 {
		params["before"] = before
	}

	uri := constants.SpotAccount + "/transferRecords"

	resp, err := p.BitgetRestClient.DoGet(uri, params)

	return resp, err

}

/**
 * Get the bill flow
 * @param BillsReq
 * @return ResponseResult
 */
func (p *SpotAccountClient) Bills(params account.BillsReq) (string, error) {

	postBody, jsonErr := internal.ToJson(params)

	if jsonErr != nil {
		return "", jsonErr
	}

	uri := constants.SpotAccount + "/bills"

	resp, err := p.BitgetRestClient.DoPost(uri, postBody)

	return resp, err
}
