package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ies/v1/model"
)

type IesClient struct {
	HcClient *http_client.HcHttpClient
}

func NewIesClient(hcClient *http_client.HcHttpClient) *IesClient {
	return &IesClient{HcClient: hcClient}
}

func IesClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// CreateEdgeSite 创建边缘小站
//
// 创建边缘小站。
// - 一个边缘小站关联一个华为云指定的区域。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IesClient) CreateEdgeSite(request *model.CreateEdgeSiteRequest) (*model.CreateEdgeSiteResponse, error) {
	requestDef := GenReqDefForCreateEdgeSite()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEdgeSiteResponse), nil
	}
}

// CreateEdgeSiteInvoker 创建边缘小站
func (c *IesClient) CreateEdgeSiteInvoker(request *model.CreateEdgeSiteRequest) *CreateEdgeSiteInvoker {
	requestDef := GenReqDefForCreateEdgeSite()
	return &CreateEdgeSiteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEdgeSite 删除边缘小站
//
// 删除指定的边缘小站。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IesClient) DeleteEdgeSite(request *model.DeleteEdgeSiteRequest) (*model.DeleteEdgeSiteResponse, error) {
	requestDef := GenReqDefForDeleteEdgeSite()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEdgeSiteResponse), nil
	}
}

// DeleteEdgeSiteInvoker 删除边缘小站
func (c *IesClient) DeleteEdgeSiteInvoker(request *model.DeleteEdgeSiteRequest) *DeleteEdgeSiteInvoker {
	requestDef := GenReqDefForDeleteEdgeSite()
	return &DeleteEdgeSiteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEdgeSites 查询边缘小站列表
//
// 查询边缘小站列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IesClient) ListEdgeSites(request *model.ListEdgeSitesRequest) (*model.ListEdgeSitesResponse, error) {
	requestDef := GenReqDefForListEdgeSites()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEdgeSitesResponse), nil
	}
}

// ListEdgeSitesInvoker 查询边缘小站列表
func (c *IesClient) ListEdgeSitesInvoker(request *model.ListEdgeSitesRequest) *ListEdgeSitesInvoker {
	requestDef := GenReqDefForListEdgeSites()
	return &ListEdgeSitesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEdgeSite 查询边缘小站详情
//
// 查询边缘小站详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IesClient) ShowEdgeSite(request *model.ShowEdgeSiteRequest) (*model.ShowEdgeSiteResponse, error) {
	requestDef := GenReqDefForShowEdgeSite()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEdgeSiteResponse), nil
	}
}

// ShowEdgeSiteInvoker 查询边缘小站详情
func (c *IesClient) ShowEdgeSiteInvoker(request *model.ShowEdgeSiteRequest) *ShowEdgeSiteInvoker {
	requestDef := GenReqDefForShowEdgeSite()
	return &ShowEdgeSiteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEdgeSite 更新边缘小站
//
// 更新边缘小站。
// - 允许更新边缘小站描述或场地信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IesClient) UpdateEdgeSite(request *model.UpdateEdgeSiteRequest) (*model.UpdateEdgeSiteResponse, error) {
	requestDef := GenReqDefForUpdateEdgeSite()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEdgeSiteResponse), nil
	}
}

// UpdateEdgeSiteInvoker 更新边缘小站
func (c *IesClient) UpdateEdgeSiteInvoker(request *model.UpdateEdgeSiteRequest) *UpdateEdgeSiteInvoker {
	requestDef := GenReqDefForUpdateEdgeSite()
	return &UpdateEdgeSiteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEdgeSiteMetrics 查看站点容量信息
//
// 查看站点容量信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IesClient) ListEdgeSiteMetrics(request *model.ListEdgeSiteMetricsRequest) (*model.ListEdgeSiteMetricsResponse, error) {
	requestDef := GenReqDefForListEdgeSiteMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEdgeSiteMetricsResponse), nil
	}
}

// ListEdgeSiteMetricsInvoker 查看站点容量信息
func (c *IesClient) ListEdgeSiteMetricsInvoker(request *model.ListEdgeSiteMetricsRequest) *ListEdgeSiteMetricsInvoker {
	requestDef := GenReqDefForListEdgeSiteMetrics()
	return &ListEdgeSiteMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQuotas 查询配额
//
// 查询租户资源配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IesClient) ListQuotas(request *model.ListQuotasRequest) (*model.ListQuotasResponse, error) {
	requestDef := GenReqDefForListQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotasResponse), nil
	}
}

// ListQuotasInvoker 查询配额
func (c *IesClient) ListQuotasInvoker(request *model.ListQuotasRequest) *ListQuotasInvoker {
	requestDef := GenReqDefForListQuotas()
	return &ListQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSupportedRegions 查询支持的区域列表
//
// 查询支持智能边缘小站接入的华为云区域（region）列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IesClient) ListSupportedRegions(request *model.ListSupportedRegionsRequest) (*model.ListSupportedRegionsResponse, error) {
	requestDef := GenReqDefForListSupportedRegions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSupportedRegionsResponse), nil
	}
}

// ListSupportedRegionsInvoker 查询支持的区域列表
func (c *IesClient) ListSupportedRegionsInvoker(request *model.ListSupportedRegionsRequest) *ListSupportedRegionsInvoker {
	requestDef := GenReqDefForListSupportedRegions()
	return &ListSupportedRegionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
