package market

import "github.com/wselfjes/bitget/pkg/model"

type DepthtResp struct {
	model.BaseResp `json:",inline"`
	Data           Depth `json:"data"`
}

type Depth struct {
	Asks      []AskBid `json:"asks"`
	Bids      []AskBid `json:"bids"`
	Timestamp string   `json:"timestamp"`
}

type AskBid []string

// {
// 	"code": "00000",
// 	"msg": "success",
// 	"requestTime": 0,
// 	"data": {
// 	  "asks": [
// 		[
// 		  "23821.5",
// 		  "0.166"
// 		],
// 		[
// 		  "23822",
// 		  "5.924"
// 		],
// 		[
// 		  "23823.5",
// 		  "13.286"
// 		]
// 	  ],
// 	  "bids": [
// 		[
// 		  "23821",
// 		  "60.404"
// 		],
// 		[
// 		  "23820.5",
// 		  "28.646"
// 		],
// 	  ],
// 	  "timestamp": "1675357343893"
// 	}
// }
