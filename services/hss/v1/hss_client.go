package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/hss/v1/model"
)

type HssClient struct {
	HcClient *http_client.HcHttpClient
}

func NewHssClient(hcClient *http_client.HcHttpClient) *HssClient {
	return &HssClient{HcClient: hcClient}
}

func HssClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//查入侵事件列表
func (c *HssClient) ListEvents(request *model.ListEventsRequest) (*model.ListEventsResponse, error) {
	requestDef := GenReqDefForListEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEventsResponse), nil
	}
}

//查询弹性云服务器状态列表
func (c *HssClient) ListHosts(request *model.ListHostsRequest) (*model.ListHostsResponse, error) {
	requestDef := GenReqDefForListHosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostsResponse), nil
	}
}
