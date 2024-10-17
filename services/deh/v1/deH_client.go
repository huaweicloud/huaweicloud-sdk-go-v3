package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/deh/v1/model"
)

type DeHClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewDeHClient(hcClient *httpclient.HcHttpClient) *DeHClient {
	return &DeHClient{HcClient: hcClient}
}

func DeHClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// BatchCreateDedicatedHostTags 批量添加专属主机标签
//
// 为指定专属主机批量添加标签。
//
// 标签管理服务（TMS）使用该接口批量添加专属主机的标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DeHClient) BatchCreateDedicatedHostTags(request *model.BatchCreateDedicatedHostTagsRequest) (*model.BatchCreateDedicatedHostTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateDedicatedHostTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateDedicatedHostTagsResponse), nil
	}
}

// BatchCreateDedicatedHostTagsInvoker 批量添加专属主机标签
func (c *DeHClient) BatchCreateDedicatedHostTagsInvoker(request *model.BatchCreateDedicatedHostTagsRequest) *BatchCreateDedicatedHostTagsInvoker {
	requestDef := GenReqDefForBatchCreateDedicatedHostTags()
	return &BatchCreateDedicatedHostTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteDedicatedHostTags 批量删除专属主机标签
//
// 批量删除指定专属主机标签。
//
// 标签管理服务（TMS）使用该接口批量删除专属主机的标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DeHClient) BatchDeleteDedicatedHostTags(request *model.BatchDeleteDedicatedHostTagsRequest) (*model.BatchDeleteDedicatedHostTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteDedicatedHostTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteDedicatedHostTagsResponse), nil
	}
}

// BatchDeleteDedicatedHostTagsInvoker 批量删除专属主机标签
func (c *DeHClient) BatchDeleteDedicatedHostTagsInvoker(request *model.BatchDeleteDedicatedHostTagsRequest) *BatchDeleteDedicatedHostTagsInvoker {
	requestDef := GenReqDefForBatchDeleteDedicatedHostTags()
	return &BatchDeleteDedicatedHostTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDedicatedHostTypes 查询可用的专属主机类型
//
// 查询某一AZ内可用的专属主机类型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DeHClient) ListDedicatedHostTypes(request *model.ListDedicatedHostTypesRequest) (*model.ListDedicatedHostTypesResponse, error) {
	requestDef := GenReqDefForListDedicatedHostTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDedicatedHostTypesResponse), nil
	}
}

// ListDedicatedHostTypesInvoker 查询可用的专属主机类型
func (c *DeHClient) ListDedicatedHostTypesInvoker(request *model.ListDedicatedHostTypesRequest) *ListDedicatedHostTypesInvoker {
	requestDef := GenReqDefForListDedicatedHostTypes()
	return &ListDedicatedHostTypesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDedicatedHosts 查询专属主机列表
//
// 通过该接口查询专属主机列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DeHClient) ListDedicatedHosts(request *model.ListDedicatedHostsRequest) (*model.ListDedicatedHostsResponse, error) {
	requestDef := GenReqDefForListDedicatedHosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDedicatedHostsResponse), nil
	}
}

// ListDedicatedHostsInvoker 查询专属主机列表
func (c *DeHClient) ListDedicatedHostsInvoker(request *model.ListDedicatedHostsRequest) *ListDedicatedHostsInvoker {
	requestDef := GenReqDefForListDedicatedHosts()
	return &ListDedicatedHostsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDedicatedHostsByTags 按标签查询专属主机列表
//
// 使用标签过滤专属主机列表，并返回专属主机使用的所有标签。
//
// 标签管理服务（TMS）使用该接口过滤专属主机列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DeHClient) ListDedicatedHostsByTags(request *model.ListDedicatedHostsByTagsRequest) (*model.ListDedicatedHostsByTagsResponse, error) {
	requestDef := GenReqDefForListDedicatedHostsByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDedicatedHostsByTagsResponse), nil
	}
}

// ListDedicatedHostsByTagsInvoker 按标签查询专属主机列表
func (c *DeHClient) ListDedicatedHostsByTagsInvoker(request *model.ListDedicatedHostsByTagsRequest) *ListDedicatedHostsByTagsInvoker {
	requestDef := GenReqDefForListDedicatedHostsByTags()
	return &ListDedicatedHostsByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServersDedicatedHost 查询专属主机上的云服务器
//
// 查询专属主机上已部署的云服务器信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DeHClient) ListServersDedicatedHost(request *model.ListServersDedicatedHostRequest) (*model.ListServersDedicatedHostResponse, error) {
	requestDef := GenReqDefForListServersDedicatedHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServersDedicatedHostResponse), nil
	}
}

// ListServersDedicatedHostInvoker 查询专属主机上的云服务器
func (c *DeHClient) ListServersDedicatedHostInvoker(request *model.ListServersDedicatedHostRequest) *ListServersDedicatedHostInvoker {
	requestDef := GenReqDefForListServersDedicatedHost()
	return &ListServersDedicatedHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDedicatedHost 查询专属主机详情
//
// 查询某一台专属主机的详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DeHClient) ShowDedicatedHost(request *model.ShowDedicatedHostRequest) (*model.ShowDedicatedHostResponse, error) {
	requestDef := GenReqDefForShowDedicatedHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDedicatedHostResponse), nil
	}
}

// ShowDedicatedHostInvoker 查询专属主机详情
func (c *DeHClient) ShowDedicatedHostInvoker(request *model.ShowDedicatedHostRequest) *ShowDedicatedHostInvoker {
	requestDef := GenReqDefForShowDedicatedHost()
	return &ShowDedicatedHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDedicatedHostTags 查询指定专属主机标签
//
// 查询指定专属主机的标签信息。
//
// 标签管理服务（TMS）使用该接口查询指定专属主机的全部标签数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DeHClient) ShowDedicatedHostTags(request *model.ShowDedicatedHostTagsRequest) (*model.ShowDedicatedHostTagsResponse, error) {
	requestDef := GenReqDefForShowDedicatedHostTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDedicatedHostTagsResponse), nil
	}
}

// ShowDedicatedHostTagsInvoker 查询指定专属主机标签
func (c *DeHClient) ShowDedicatedHostTagsInvoker(request *model.ShowDedicatedHostTagsRequest) *ShowDedicatedHostTagsInvoker {
	requestDef := GenReqDefForShowDedicatedHostTags()
	return &ShowDedicatedHostTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQuotaSets 查询租户的专属主机配额
//
// 该接口用于查询租户的专属主机配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DeHClient) ShowQuotaSets(request *model.ShowQuotaSetsRequest) (*model.ShowQuotaSetsResponse, error) {
	requestDef := GenReqDefForShowQuotaSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotaSetsResponse), nil
	}
}

// ShowQuotaSetsInvoker 查询租户的专属主机配额
func (c *DeHClient) ShowQuotaSetsInvoker(request *model.ShowQuotaSetsRequest) *ShowQuotaSetsInvoker {
	requestDef := GenReqDefForShowQuotaSets()
	return &ShowQuotaSetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDedicatedHost 更新专属主机属性
//
// 该接口用于变更专属主机的“auto_placement”和“name”属性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DeHClient) UpdateDedicatedHost(request *model.UpdateDedicatedHostRequest) (*model.UpdateDedicatedHostResponse, error) {
	requestDef := GenReqDefForUpdateDedicatedHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDedicatedHostResponse), nil
	}
}

// UpdateDedicatedHostInvoker 更新专属主机属性
func (c *DeHClient) UpdateDedicatedHostInvoker(request *model.UpdateDedicatedHostRequest) *UpdateDedicatedHostInvoker {
	requestDef := GenReqDefForUpdateDedicatedHost()
	return &UpdateDedicatedHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
