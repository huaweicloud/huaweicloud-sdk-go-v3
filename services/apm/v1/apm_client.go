package v1

import (
	http_client "github.com/RandolphCYG/huaweicloud-sdk-go-v3/core"

	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/services/apm/v1/model"
)

type ApmClient struct {
	HcClient *http_client.HcHttpClient
}

func NewApmClient(hcClient *http_client.HcHttpClient) *ApmClient {
	return &ApmClient{HcClient: hcClient}
}

func ApmClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//
func (c *ApmClient) ListBusiness(request *model.ListBusinessRequest) (*model.ListBusinessResponse, error) {
	requestDef := GenReqDefForListBusiness()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBusinessResponse), nil
	}
}
