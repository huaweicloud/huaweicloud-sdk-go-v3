package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cloudpond/v1/model"
)

type CloudPondClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCloudPondClient(hcClient *httpclient.HcHttpClient) *CloudPondClient {
	return &CloudPondClient{HcClient: hcClient}
}

func CloudPondClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// CreateEdgeSite 创建边缘小站
//
// 创建边缘小站。
// - 一个边缘小站关联一个华为云指定的区域。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudPondClient) CreateEdgeSite(request *model.CreateEdgeSiteRequest) (*model.CreateEdgeSiteResponse, error) {
	requestDef := GenReqDefForCreateEdgeSite()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEdgeSiteResponse), nil
	}
}

// CreateEdgeSiteInvoker 创建边缘小站
func (c *CloudPondClient) CreateEdgeSiteInvoker(request *model.CreateEdgeSiteRequest) *CreateEdgeSiteInvoker {
	requestDef := GenReqDefForCreateEdgeSite()
	return &CreateEdgeSiteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEdgeSite 删除边缘小站
//
// 删除指定的边缘小站。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudPondClient) DeleteEdgeSite(request *model.DeleteEdgeSiteRequest) (*model.DeleteEdgeSiteResponse, error) {
	requestDef := GenReqDefForDeleteEdgeSite()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEdgeSiteResponse), nil
	}
}

// DeleteEdgeSiteInvoker 删除边缘小站
func (c *CloudPondClient) DeleteEdgeSiteInvoker(request *model.DeleteEdgeSiteRequest) *DeleteEdgeSiteInvoker {
	requestDef := GenReqDefForDeleteEdgeSite()
	return &DeleteEdgeSiteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEdgeSites 查询边缘小站列表
//
// 查询边缘小站列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudPondClient) ListEdgeSites(request *model.ListEdgeSitesRequest) (*model.ListEdgeSitesResponse, error) {
	requestDef := GenReqDefForListEdgeSites()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEdgeSitesResponse), nil
	}
}

// ListEdgeSitesInvoker 查询边缘小站列表
func (c *CloudPondClient) ListEdgeSitesInvoker(request *model.ListEdgeSitesRequest) *ListEdgeSitesInvoker {
	requestDef := GenReqDefForListEdgeSites()
	return &ListEdgeSitesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEdgeSite 查询边缘小站详情
//
// 查询边缘小站详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudPondClient) ShowEdgeSite(request *model.ShowEdgeSiteRequest) (*model.ShowEdgeSiteResponse, error) {
	requestDef := GenReqDefForShowEdgeSite()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEdgeSiteResponse), nil
	}
}

// ShowEdgeSiteInvoker 查询边缘小站详情
func (c *CloudPondClient) ShowEdgeSiteInvoker(request *model.ShowEdgeSiteRequest) *ShowEdgeSiteInvoker {
	requestDef := GenReqDefForShowEdgeSite()
	return &ShowEdgeSiteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEdgeSite 更新边缘小站
//
// 更新边缘小站。
// - 允许更新边缘小站描述或场地信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudPondClient) UpdateEdgeSite(request *model.UpdateEdgeSiteRequest) (*model.UpdateEdgeSiteResponse, error) {
	requestDef := GenReqDefForUpdateEdgeSite()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEdgeSiteResponse), nil
	}
}

// UpdateEdgeSiteInvoker 更新边缘小站
func (c *CloudPondClient) UpdateEdgeSiteInvoker(request *model.UpdateEdgeSiteRequest) *UpdateEdgeSiteInvoker {
	requestDef := GenReqDefForUpdateEdgeSite()
	return &UpdateEdgeSiteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEdgeSiteMetrics 查看站点容量信息
//
// 查看站点容量信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudPondClient) ListEdgeSiteMetrics(request *model.ListEdgeSiteMetricsRequest) (*model.ListEdgeSiteMetricsResponse, error) {
	requestDef := GenReqDefForListEdgeSiteMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEdgeSiteMetricsResponse), nil
	}
}

// ListEdgeSiteMetricsInvoker 查看站点容量信息
func (c *CloudPondClient) ListEdgeSiteMetricsInvoker(request *model.ListEdgeSiteMetricsRequest) *ListEdgeSiteMetricsInvoker {
	requestDef := GenReqDefForListEdgeSiteMetrics()
	return &ListEdgeSiteMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQuotas 查询配额
//
// 查询租户资源配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudPondClient) ListQuotas(request *model.ListQuotasRequest) (*model.ListQuotasResponse, error) {
	requestDef := GenReqDefForListQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotasResponse), nil
	}
}

// ListQuotasInvoker 查询配额
func (c *CloudPondClient) ListQuotasInvoker(request *model.ListQuotasRequest) *ListQuotasInvoker {
	requestDef := GenReqDefForListQuotas()
	return &ListQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRacks 查询机柜列表
//
// 查询机柜列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudPondClient) ListRacks(request *model.ListRacksRequest) (*model.ListRacksResponse, error) {
	requestDef := GenReqDefForListRacks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRacksResponse), nil
	}
}

// ListRacksInvoker 查询机柜列表
func (c *CloudPondClient) ListRacksInvoker(request *model.ListRacksRequest) *ListRacksInvoker {
	requestDef := GenReqDefForListRacks()
	return &ListRacksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRack 查询机柜详情
//
// 查询机柜详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudPondClient) ShowRack(request *model.ShowRackRequest) (*model.ShowRackResponse, error) {
	requestDef := GenReqDefForShowRack()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRackResponse), nil
	}
}

// ShowRackInvoker 查询机柜详情
func (c *CloudPondClient) ShowRackInvoker(request *model.ShowRackRequest) *ShowRackInvoker {
	requestDef := GenReqDefForShowRack()
	return &ShowRackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSupportedRegions 查询支持的区域列表
//
// 查询支持CloudPond接入的华为云区域（region）列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudPondClient) ListSupportedRegions(request *model.ListSupportedRegionsRequest) (*model.ListSupportedRegionsResponse, error) {
	requestDef := GenReqDefForListSupportedRegions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSupportedRegionsResponse), nil
	}
}

// ListSupportedRegionsInvoker 查询支持的区域列表
func (c *CloudPondClient) ListSupportedRegionsInvoker(request *model.ListSupportedRegionsRequest) *ListSupportedRegionsInvoker {
	requestDef := GenReqDefForListSupportedRegions()
	return &ListSupportedRegionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStoragePools 查询存储池列表
//
// 查询存储池列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudPondClient) ListStoragePools(request *model.ListStoragePoolsRequest) (*model.ListStoragePoolsResponse, error) {
	requestDef := GenReqDefForListStoragePools()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStoragePoolsResponse), nil
	}
}

// ListStoragePoolsInvoker 查询存储池列表
func (c *CloudPondClient) ListStoragePoolsInvoker(request *model.ListStoragePoolsRequest) *ListStoragePoolsInvoker {
	requestDef := GenReqDefForListStoragePools()
	return &ListStoragePoolsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowStoragePool 查询存储池详情
//
// 查询存储池详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudPondClient) ShowStoragePool(request *model.ShowStoragePoolRequest) (*model.ShowStoragePoolResponse, error) {
	requestDef := GenReqDefForShowStoragePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStoragePoolResponse), nil
	}
}

// ShowStoragePoolInvoker 查询存储池详情
func (c *CloudPondClient) ShowStoragePoolInvoker(request *model.ShowStoragePoolRequest) *ShowStoragePoolInvoker {
	requestDef := GenReqDefForShowStoragePool()
	return &ShowStoragePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSupportedZones 查询支持的地区列表
//
// 查询支持CloudPond接入的华为云地区列表。
// - 该接口支持企业项目细粒度权限的校验，具体细粒度请参见 ies:zone:list
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudPondClient) ListSupportedZones(request *model.ListSupportedZonesRequest) (*model.ListSupportedZonesResponse, error) {
	requestDef := GenReqDefForListSupportedZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSupportedZonesResponse), nil
	}
}

// ListSupportedZonesInvoker 查询支持的地区列表
func (c *CloudPondClient) ListSupportedZonesInvoker(request *model.ListSupportedZonesRequest) *ListSupportedZonesInvoker {
	requestDef := GenReqDefForListSupportedZones()
	return &ListSupportedZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
