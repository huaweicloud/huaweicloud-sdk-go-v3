package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dcos/v1/model"
)

type DcosClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewDcosClient(hcClient *httpclient.HcHttpClient) *DcosClient {
	return &DcosClient{HcClient: hcClient}
}

func DcosClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// ListOrder 客户查询服务单列表
//
// 客户查询服务单列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcosClient) ListOrder(request *model.ListOrderRequest) (*model.ListOrderResponse, error) {
	requestDef := GenReqDefForListOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOrderResponse), nil
	}
}

// ListOrderInvoker 客户查询服务单列表
func (c *DcosClient) ListOrderInvoker(request *model.ListOrderRequest) *ListOrderInvoker {
	requestDef := GenReqDefForListOrder()
	return &ListOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SaveOrder 客户创建服务单
//
// 客户创建服务单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcosClient) SaveOrder(request *model.SaveOrderRequest) (*model.SaveOrderResponse, error) {
	requestDef := GenReqDefForSaveOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SaveOrderResponse), nil
	}
}

// SaveOrderInvoker 客户创建服务单
func (c *DcosClient) SaveOrderInvoker(request *model.SaveOrderRequest) *SaveOrderInvoker {
	requestDef := GenReqDefForSaveOrder()
	return &SaveOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOrder 客户查询服务单详情
//
// 客户查询服务单详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcosClient) ShowOrder(request *model.ShowOrderRequest) (*model.ShowOrderResponse, error) {
	requestDef := GenReqDefForShowOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOrderResponse), nil
	}
}

// ShowOrderInvoker 客户查询服务单详情
func (c *DcosClient) ShowOrderInvoker(request *model.ShowOrderRequest) *ShowOrderInvoker {
	requestDef := GenReqDefForShowOrder()
	return &ShowOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOrderCatalogue 获取服务单目录列表
//
// 获取服务单目录列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcosClient) ShowOrderCatalogue(request *model.ShowOrderCatalogueRequest) (*model.ShowOrderCatalogueResponse, error) {
	requestDef := GenReqDefForShowOrderCatalogue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOrderCatalogueResponse), nil
	}
}

// ShowOrderCatalogueInvoker 获取服务单目录列表
func (c *DcosClient) ShowOrderCatalogueInvoker(request *model.ShowOrderCatalogueRequest) *ShowOrderCatalogueInvoker {
	requestDef := GenReqDefForShowOrderCatalogue()
	return &ShowOrderCatalogueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOrderInformation 获取服务服务单项目信息
//
// 获取服务服务单项目信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcosClient) ShowOrderInformation(request *model.ShowOrderInformationRequest) (*model.ShowOrderInformationResponse, error) {
	requestDef := GenReqDefForShowOrderInformation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOrderInformationResponse), nil
	}
}

// ShowOrderInformationInvoker 获取服务服务单项目信息
func (c *DcosClient) ShowOrderInformationInvoker(request *model.ShowOrderInformationRequest) *ShowOrderInformationInvoker {
	requestDef := GenReqDefForShowOrderInformation()
	return &ShowOrderInformationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPageAssetListResult 资产列表
//
// 资产列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcosClient) ShowPageAssetListResult(request *model.ShowPageAssetListResultRequest) (*model.ShowPageAssetListResultResponse, error) {
	requestDef := GenReqDefForShowPageAssetListResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPageAssetListResultResponse), nil
	}
}

// ShowPageAssetListResultInvoker 资产列表
func (c *DcosClient) ShowPageAssetListResultInvoker(request *model.ShowPageAssetListResultRequest) *ShowPageAssetListResultInvoker {
	requestDef := GenReqDefForShowPageAssetListResult()
	return &ShowPageAssetListResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadFile 上传附件
//
// 上传附件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcosClient) UploadFile(request *model.UploadFileRequest) (*model.UploadFileResponse, error) {
	requestDef := GenReqDefForUploadFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadFileResponse), nil
	}
}

// UploadFileInvoker 上传附件
func (c *DcosClient) UploadFileInvoker(request *model.UploadFileRequest) *UploadFileInvoker {
	requestDef := GenReqDefForUploadFile()
	return &UploadFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// VerifyOrder 验收服务单
//
// 验收服务单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcosClient) VerifyOrder(request *model.VerifyOrderRequest) (*model.VerifyOrderResponse, error) {
	requestDef := GenReqDefForVerifyOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.VerifyOrderResponse), nil
	}
}

// VerifyOrderInvoker 验收服务单
func (c *DcosClient) VerifyOrderInvoker(request *model.VerifyOrderRequest) *VerifyOrderInvoker {
	requestDef := GenReqDefForVerifyOrder()
	return &VerifyOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
