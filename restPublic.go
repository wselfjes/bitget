package bitget

import (
	"github.com/wselfjes/bitget/config"
	"github.com/wselfjes/bitget/internal/common"

	v2 "github.com/wselfjes/bitget/pkg/client/v2"
)

type Rest struct {
	c *common.BitgetRestClient
}

func NewRest(creds config.InterApiCreds) *Rest {
	return &Rest{
		c: new(common.BitgetRestClient).Init(creds),
	}
}

type InterRestPublic interface {
	GetContractConfig(productType string) (*v2.GetContractConfig, error)
}

func (r *Rest) GetContractConfig(productType string) (*v2.GetContractConfig, error) {
	return v2.NewGetContractConfig(r.c, productType)
}
