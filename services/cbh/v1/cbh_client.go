package v1

import (
	http_client "github.com/RandolphCYG/huaweicloud-sdk-go-v3/core"

	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/services/cbh/v1/model"
)

type CbhClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCbhClient(hcClient *http_client.HcHttpClient) *CbhClient {
	return &CbhClient{HcClient: hcClient}
}

func CbhClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//获取CBH实例列表
func (c *CbhClient) ListCbhInstance(request *model.ListCbhInstanceRequest) (*model.ListCbhInstanceResponse, error) {
	requestDef := GenReqDefForListCbhInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCbhInstanceResponse), nil
	}
}
