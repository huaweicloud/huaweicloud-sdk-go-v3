package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/deh/v1/model"
)

type DeHClient struct {
	HcClient *http_client.HcHttpClient
}

func NewDeHClient(hcClient *http_client.HcHttpClient) *DeHClient {
	return &DeHClient{HcClient: hcClient}
}

func DeHClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// 批量添加专属主机标签
//
// 为指定专属主机批量添加标签。
//
// 标签管理服务（TMS）使用该接口批量添加专属主机的标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DeHClient) BatchCreateDedicatedHostTags(request *model.BatchCreateDedicatedHostTagsRequest) (*model.BatchCreateDedicatedHostTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateDedicatedHostTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateDedicatedHostTagsResponse), nil
	}
}

// 批量删除专属主机标签
//
// 批量删除指定专属主机标签。
//
// 标签管理服务（TMS）使用该接口批量删除专属主机的标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DeHClient) BatchDeleteDedicatedHostTags(request *model.BatchDeleteDedicatedHostTagsRequest) (*model.BatchDeleteDedicatedHostTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteDedicatedHostTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteDedicatedHostTagsResponse), nil
	}
}

// 分配专属主机
//
// 分配一台或多台专属主机，需要设置实例规格、所属AZ、数量等参数。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DeHClient) CreateDedicatedHost(request *model.CreateDedicatedHostRequest) (*model.CreateDedicatedHostResponse, error) {
	requestDef := GenReqDefForCreateDedicatedHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDedicatedHostResponse), nil
	}
}

// 释放专属主机
//
// 释放专属主机。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DeHClient) DeleteDedicatedHost(request *model.DeleteDedicatedHostRequest) (*model.DeleteDedicatedHostResponse, error) {
	requestDef := GenReqDefForDeleteDedicatedHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDedicatedHostResponse), nil
	}
}

// 查询可用的专属主机类型
//
// 查询某一AZ内可用的专属主机类型。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DeHClient) ListDedicatedHostTypes(request *model.ListDedicatedHostTypesRequest) (*model.ListDedicatedHostTypesResponse, error) {
	requestDef := GenReqDefForListDedicatedHostTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDedicatedHostTypesResponse), nil
	}
}

// 查询专属主机列表
//
// 通过该接口查询专属主机列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DeHClient) ListDedicatedHosts(request *model.ListDedicatedHostsRequest) (*model.ListDedicatedHostsResponse, error) {
	requestDef := GenReqDefForListDedicatedHosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDedicatedHostsResponse), nil
	}
}

// 按标签查询专属主机列表
//
// 使用标签过滤专属主机列表，并返回专属主机使用的所有标签。
//
// 标签管理服务（TMS）使用该接口过滤专属主机列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DeHClient) ListDedicatedHostsByTags(request *model.ListDedicatedHostsByTagsRequest) (*model.ListDedicatedHostsByTagsResponse, error) {
	requestDef := GenReqDefForListDedicatedHostsByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDedicatedHostsByTagsResponse), nil
	}
}

// 查询专属主机上的云服务器
//
// 查询专属主机上已部署的云服务器信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DeHClient) ListServersDedicatedHost(request *model.ListServersDedicatedHostRequest) (*model.ListServersDedicatedHostResponse, error) {
	requestDef := GenReqDefForListServersDedicatedHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServersDedicatedHostResponse), nil
	}
}

// 查询专属主机详情
//
// 查询某一台专属主机的详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DeHClient) ShowDedicatedHost(request *model.ShowDedicatedHostRequest) (*model.ShowDedicatedHostResponse, error) {
	requestDef := GenReqDefForShowDedicatedHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDedicatedHostResponse), nil
	}
}

// 查询指定专属主机标签
//
// 查询指定专属主机的标签信息。
//
// 标签管理服务（TMS）使用该接口查询指定专属主机的全部标签数据。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DeHClient) ShowDedicatedHostTags(request *model.ShowDedicatedHostTagsRequest) (*model.ShowDedicatedHostTagsResponse, error) {
	requestDef := GenReqDefForShowDedicatedHostTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDedicatedHostTagsResponse), nil
	}
}

// 查询租户的专属主机配额
//
// 该接口用于查询租户的专属主机配额。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DeHClient) ShowQuotaSets(request *model.ShowQuotaSetsRequest) (*model.ShowQuotaSetsResponse, error) {
	requestDef := GenReqDefForShowQuotaSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotaSetsResponse), nil
	}
}

// 更新专属主机属性
//
// 该接口用于变更专属主机的“auto_placement”和“name”属性。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DeHClient) UpdateDedicatedHost(request *model.UpdateDedicatedHostRequest) (*model.UpdateDedicatedHostResponse, error) {
	requestDef := GenReqDefForUpdateDedicatedHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDedicatedHostResponse), nil
	}
}
