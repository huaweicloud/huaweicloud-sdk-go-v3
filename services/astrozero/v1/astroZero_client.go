package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/astrozero/v1/model"
)

type AstroZeroClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewAstroZeroClient(hcClient *httpclient.HcHttpClient) *AstroZeroClient {
	return &AstroZeroClient{HcClient: hcClient}
}

func AstroZeroClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// ShowOrderStatus 查询活动时间内租户开通情况
//
// 查询活动时间租户开通情况
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AstroZeroClient) ShowOrderStatus(request *model.ShowOrderStatusRequest) (*model.ShowOrderStatusResponse, error) {
	requestDef := GenReqDefForShowOrderStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOrderStatusResponse), nil
	}
}

// ShowOrderStatusInvoker 查询活动时间内租户开通情况
func (c *AstroZeroClient) ShowOrderStatusInvoker(request *model.ShowOrderStatusRequest) *ShowOrderStatusInvoker {
	requestDef := GenReqDefForShowOrderStatus()
	return &ShowOrderStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
