package v1

import (
	http_client "github.com/dysodeng/huaweicloud-sdk-go-v3/core"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/secmaster/v1/model"
)

type SecMasterClient struct {
	HcClient *http_client.HcHttpClient
}

func NewSecMasterClient(hcClient *http_client.HcHttpClient) *SecMasterClient {
	return &SecMasterClient{HcClient: hcClient}
}

func SecMasterClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CheckProductHealthy 检查心跳健康（仅支持华北-北京四使用）
//
// SA提供心跳接口，集成产品定时（每五分钟）发送心跳报文到态势感知，用来确认集成产品与态势感知之间的通路是否健康。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CheckProductHealthy(request *model.CheckProductHealthyRequest) (*model.CheckProductHealthyResponse, error) {
	requestDef := GenReqDefForCheckProductHealthy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckProductHealthyResponse), nil
	}
}

// CheckProductHealthyInvoker 检查心跳健康（仅支持华北-北京四使用）
func (c *SecMasterClient) CheckProductHealthyInvoker(request *model.CheckProductHealthyRequest) *CheckProductHealthyInvoker {
	requestDef := GenReqDefForCheckProductHealthy()
	return &CheckProductHealthyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportEvents 上报安全产品数据（仅支持华北-北京四使用）
//
// 批量数据上报，每批次最多不超过50条。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ImportEvents(request *model.ImportEventsRequest) (*model.ImportEventsResponse, error) {
	requestDef := GenReqDefForImportEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportEventsResponse), nil
	}
}

// ImportEventsInvoker 上报安全产品数据（仅支持华北-北京四使用）
func (c *SecMasterClient) ImportEventsInvoker(request *model.ImportEventsRequest) *ImportEventsInvoker {
	requestDef := GenReqDefForImportEvents()
	return &ImportEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
