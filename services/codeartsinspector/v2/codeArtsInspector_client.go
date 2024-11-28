package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codeartsinspector/v2/model"
)

type CodeArtsInspectorClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCodeArtsInspectorClient(hcClient *httpclient.HcHttpClient) *CodeArtsInspectorClient {
	return &CodeArtsInspectorClient{HcClient: hcClient}
}

func CodeArtsInspectorClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreatePurchaseOrder 订购下单接口
//
// 订购下单接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsInspectorClient) CreatePurchaseOrder(request *model.CreatePurchaseOrderRequest) (*model.CreatePurchaseOrderResponse, error) {
	requestDef := GenReqDefForCreatePurchaseOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePurchaseOrderResponse), nil
	}
}

// CreatePurchaseOrderInvoker 订购下单接口
func (c *CodeArtsInspectorClient) CreatePurchaseOrderInvoker(request *model.CreatePurchaseOrderRequest) *CreatePurchaseOrderInvoker {
	requestDef := GenReqDefForCreatePurchaseOrder()
	return &CreatePurchaseOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePurchaseOrder 变更下单接口
//
// 变更下单接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsInspectorClient) UpdatePurchaseOrder(request *model.UpdatePurchaseOrderRequest) (*model.UpdatePurchaseOrderResponse, error) {
	requestDef := GenReqDefForUpdatePurchaseOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePurchaseOrderResponse), nil
	}
}

// UpdatePurchaseOrderInvoker 变更下单接口
func (c *CodeArtsInspectorClient) UpdatePurchaseOrderInvoker(request *model.UpdatePurchaseOrderRequest) *UpdatePurchaseOrderInvoker {
	requestDef := GenReqDefForUpdatePurchaseOrder()
	return &UpdatePurchaseOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
