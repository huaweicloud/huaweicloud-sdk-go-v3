package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/sa/v1/model"
)

type SaClient struct {
	HcClient *http_client.HcHttpClient
}

func NewSaClient(hcClient *http_client.HcHttpClient) *SaClient {
	return &SaClient{HcClient: hcClient}
}

func SaClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CheckProductHealthy 检查心跳健康
//
// SA提供心跳接口，集成产品定时（例如：每五分钟）发送心跳报文到SA，用来确认集成产品与SA之间的通路是否健康。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SaClient) CheckProductHealthy(request *model.CheckProductHealthyRequest) (*model.CheckProductHealthyResponse, error) {
	requestDef := GenReqDefForCheckProductHealthy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckProductHealthyResponse), nil
	}
}

// CheckProductHealthyInvoker 检查心跳健康
func (c *SaClient) CheckProductHealthyInvoker(request *model.CheckProductHealthyRequest) *CheckProductHealthyInvoker {
	requestDef := GenReqDefForCheckProductHealthy()
	return &CheckProductHealthyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportEvents 上报安全产品数据(V2)
//
// 批量数据上报，每批次最多不超过50条。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SaClient) ImportEvents(request *model.ImportEventsRequest) (*model.ImportEventsResponse, error) {
	requestDef := GenReqDefForImportEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportEventsResponse), nil
	}
}

// ImportEventsInvoker 上报安全产品数据(V2)
func (c *SaClient) ImportEventsInvoker(request *model.ImportEventsRequest) *ImportEventsInvoker {
	requestDef := GenReqDefForImportEvents()
	return &ImportEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
