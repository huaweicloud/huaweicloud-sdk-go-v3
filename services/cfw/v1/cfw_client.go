package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cfw/v1/model"
)

type CfwClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCfwClient(hcClient *httpclient.HcHttpClient) *CfwClient {
	return &CfwClient{HcClient: hcClient}
}

func CfwClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AddAddressItem 添加地址组成员
//
// 添加地址组成员
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) AddAddressItem(request *model.AddAddressItemRequest) (*model.AddAddressItemResponse, error) {
	requestDef := GenReqDefForAddAddressItem()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddAddressItemResponse), nil
	}
}

// AddAddressItemInvoker 添加地址组成员
func (c *CfwClient) AddAddressItemInvoker(request *model.AddAddressItemRequest) *AddAddressItemInvoker {
	requestDef := GenReqDefForAddAddressItem()
	return &AddAddressItemInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddAddressSet 添加地址组
//
// 添加地址组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) AddAddressSet(request *model.AddAddressSetRequest) (*model.AddAddressSetResponse, error) {
	requestDef := GenReqDefForAddAddressSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddAddressSetResponse), nil
	}
}

// AddAddressSetInvoker 添加地址组
func (c *CfwClient) AddAddressSetInvoker(request *model.AddAddressSetRequest) *AddAddressSetInvoker {
	requestDef := GenReqDefForAddAddressSet()
	return &AddAddressSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddBlackWhiteList 创建黑白名单规则
//
// 创建黑白名单规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) AddBlackWhiteList(request *model.AddBlackWhiteListRequest) (*model.AddBlackWhiteListResponse, error) {
	requestDef := GenReqDefForAddBlackWhiteList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddBlackWhiteListResponse), nil
	}
}

// AddBlackWhiteListInvoker 创建黑白名单规则
func (c *CfwClient) AddBlackWhiteListInvoker(request *model.AddBlackWhiteListRequest) *AddBlackWhiteListInvoker {
	requestDef := GenReqDefForAddBlackWhiteList()
	return &AddBlackWhiteListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddDomainSet 添加域名组
//
// 添加域名组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) AddDomainSet(request *model.AddDomainSetRequest) (*model.AddDomainSetResponse, error) {
	requestDef := GenReqDefForAddDomainSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddDomainSetResponse), nil
	}
}

// AddDomainSetInvoker 添加域名组
func (c *CfwClient) AddDomainSetInvoker(request *model.AddDomainSetRequest) *AddDomainSetInvoker {
	requestDef := GenReqDefForAddDomainSet()
	return &AddDomainSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddDomains 添加域名列表
//
// 添加域名列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) AddDomains(request *model.AddDomainsRequest) (*model.AddDomainsResponse, error) {
	requestDef := GenReqDefForAddDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddDomainsResponse), nil
	}
}

// AddDomainsInvoker 添加域名列表
func (c *CfwClient) AddDomainsInvoker(request *model.AddDomainsRequest) *AddDomainsInvoker {
	requestDef := GenReqDefForAddDomains()
	return &AddDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddLogConfig 创建日志配置
//
// 创建日志配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) AddLogConfig(request *model.AddLogConfigRequest) (*model.AddLogConfigResponse, error) {
	requestDef := GenReqDefForAddLogConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddLogConfigResponse), nil
	}
}

// AddLogConfigInvoker 创建日志配置
func (c *CfwClient) AddLogConfigInvoker(request *model.AddLogConfigRequest) *AddLogConfigInvoker {
	requestDef := GenReqDefForAddLogConfig()
	return &AddLogConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddServiceItems 新建服务成员
//
// 批量添加服务组成员
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) AddServiceItems(request *model.AddServiceItemsRequest) (*model.AddServiceItemsResponse, error) {
	requestDef := GenReqDefForAddServiceItems()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddServiceItemsResponse), nil
	}
}

// AddServiceItemsInvoker 新建服务成员
func (c *CfwClient) AddServiceItemsInvoker(request *model.AddServiceItemsRequest) *AddServiceItemsInvoker {
	requestDef := GenReqDefForAddServiceItems()
	return &AddServiceItemsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddServiceSet 新建服务组
//
// 创建服务组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) AddServiceSet(request *model.AddServiceSetRequest) (*model.AddServiceSetResponse, error) {
	requestDef := GenReqDefForAddServiceSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddServiceSetResponse), nil
	}
}

// AddServiceSetInvoker 新建服务组
func (c *CfwClient) AddServiceSetInvoker(request *model.AddServiceSetRequest) *AddServiceSetInvoker {
	requestDef := GenReqDefForAddServiceSet()
	return &AddServiceSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteAddressItems 批量删除地址组成员
//
// 批量删除地址组成员
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) BatchDeleteAddressItems(request *model.BatchDeleteAddressItemsRequest) (*model.BatchDeleteAddressItemsResponse, error) {
	requestDef := GenReqDefForBatchDeleteAddressItems()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteAddressItemsResponse), nil
	}
}

// BatchDeleteAddressItemsInvoker 批量删除地址组成员
func (c *CfwClient) BatchDeleteAddressItemsInvoker(request *model.BatchDeleteAddressItemsRequest) *BatchDeleteAddressItemsInvoker {
	requestDef := GenReqDefForBatchDeleteAddressItems()
	return &BatchDeleteAddressItemsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteDomainSet 批量删除域名组
//
// 批量删除域名组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) BatchDeleteDomainSet(request *model.BatchDeleteDomainSetRequest) (*model.BatchDeleteDomainSetResponse, error) {
	requestDef := GenReqDefForBatchDeleteDomainSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteDomainSetResponse), nil
	}
}

// BatchDeleteDomainSetInvoker 批量删除域名组
func (c *CfwClient) BatchDeleteDomainSetInvoker(request *model.BatchDeleteDomainSetRequest) *BatchDeleteDomainSetInvoker {
	requestDef := GenReqDefForBatchDeleteDomainSet()
	return &BatchDeleteDomainSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteServiceItems 批量删除服务组成员信息
//
// 批量删除服务组成员信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) BatchDeleteServiceItems(request *model.BatchDeleteServiceItemsRequest) (*model.BatchDeleteServiceItemsResponse, error) {
	requestDef := GenReqDefForBatchDeleteServiceItems()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteServiceItemsResponse), nil
	}
}

// BatchDeleteServiceItemsInvoker 批量删除服务组成员信息
func (c *CfwClient) BatchDeleteServiceItemsInvoker(request *model.BatchDeleteServiceItemsRequest) *BatchDeleteServiceItemsInvoker {
	requestDef := GenReqDefForBatchDeleteServiceItems()
	return &BatchDeleteServiceItemsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelCaptureTask 取消抓包任务
//
// 取消抓包任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) CancelCaptureTask(request *model.CancelCaptureTaskRequest) (*model.CancelCaptureTaskResponse, error) {
	requestDef := GenReqDefForCancelCaptureTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelCaptureTaskResponse), nil
	}
}

// CancelCaptureTaskInvoker 取消抓包任务
func (c *CfwClient) CancelCaptureTaskInvoker(request *model.CancelCaptureTaskRequest) *CancelCaptureTaskInvoker {
	requestDef := GenReqDefForCancelCaptureTask()
	return &CancelCaptureTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeEastWestFirewallStatus 修改东西向防火墙防护状态
//
// 东西向防护开启/关闭
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ChangeEastWestFirewallStatus(request *model.ChangeEastWestFirewallStatusRequest) (*model.ChangeEastWestFirewallStatusResponse, error) {
	requestDef := GenReqDefForChangeEastWestFirewallStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeEastWestFirewallStatusResponse), nil
	}
}

// ChangeEastWestFirewallStatusInvoker 修改东西向防火墙防护状态
func (c *CfwClient) ChangeEastWestFirewallStatusInvoker(request *model.ChangeEastWestFirewallStatusRequest) *ChangeEastWestFirewallStatusInvoker {
	requestDef := GenReqDefForChangeEastWestFirewallStatus()
	return &ChangeEastWestFirewallStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCaptureTask 创建抓包任务
//
// 创建抓包任务，每个任务只能执行一次。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) CreateCaptureTask(request *model.CreateCaptureTaskRequest) (*model.CreateCaptureTaskResponse, error) {
	requestDef := GenReqDefForCreateCaptureTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCaptureTaskResponse), nil
	}
}

// CreateCaptureTaskInvoker 创建抓包任务
func (c *CfwClient) CreateCaptureTaskInvoker(request *model.CreateCaptureTaskRequest) *CreateCaptureTaskInvoker {
	requestDef := GenReqDefForCreateCaptureTask()
	return &CreateCaptureTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEastWestFirewall 创建东西向防火墙
//
// 创建东西向防火墙
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) CreateEastWestFirewall(request *model.CreateEastWestFirewallRequest) (*model.CreateEastWestFirewallResponse, error) {
	requestDef := GenReqDefForCreateEastWestFirewall()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEastWestFirewallResponse), nil
	}
}

// CreateEastWestFirewallInvoker 创建东西向防火墙
func (c *CfwClient) CreateEastWestFirewallInvoker(request *model.CreateEastWestFirewallRequest) *CreateEastWestFirewallInvoker {
	requestDef := GenReqDefForCreateEastWestFirewall()
	return &CreateEastWestFirewallInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFirewall 创建防火墙
//
// 创建防火墙
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) CreateFirewall(request *model.CreateFirewallRequest) (*model.CreateFirewallResponse, error) {
	requestDef := GenReqDefForCreateFirewall()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFirewallResponse), nil
	}
}

// CreateFirewallInvoker 创建防火墙
func (c *CfwClient) CreateFirewallInvoker(request *model.CreateFirewallRequest) *CreateFirewallInvoker {
	requestDef := GenReqDefForCreateFirewall()
	return &CreateFirewallInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTag 标签创建接口
//
// 创建标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) CreateTag(request *model.CreateTagRequest) (*model.CreateTagResponse, error) {
	requestDef := GenReqDefForCreateTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTagResponse), nil
	}
}

// CreateTagInvoker 标签创建接口
func (c *CfwClient) CreateTagInvoker(request *model.CreateTagRequest) *CreateTagInvoker {
	requestDef := GenReqDefForCreateTag()
	return &CreateTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAddressItem 删除地址组成员
//
// 删除地址组成员
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteAddressItem(request *model.DeleteAddressItemRequest) (*model.DeleteAddressItemResponse, error) {
	requestDef := GenReqDefForDeleteAddressItem()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAddressItemResponse), nil
	}
}

// DeleteAddressItemInvoker 删除地址组成员
func (c *CfwClient) DeleteAddressItemInvoker(request *model.DeleteAddressItemRequest) *DeleteAddressItemInvoker {
	requestDef := GenReqDefForDeleteAddressItem()
	return &DeleteAddressItemInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAddressSet 删除地址组
//
// 删除地址组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteAddressSet(request *model.DeleteAddressSetRequest) (*model.DeleteAddressSetResponse, error) {
	requestDef := GenReqDefForDeleteAddressSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAddressSetResponse), nil
	}
}

// DeleteAddressSetInvoker 删除地址组
func (c *CfwClient) DeleteAddressSetInvoker(request *model.DeleteAddressSetRequest) *DeleteAddressSetInvoker {
	requestDef := GenReqDefForDeleteAddressSet()
	return &DeleteAddressSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBlackWhiteList 删除黑白名单规则
//
// 删除黑白名单规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteBlackWhiteList(request *model.DeleteBlackWhiteListRequest) (*model.DeleteBlackWhiteListResponse, error) {
	requestDef := GenReqDefForDeleteBlackWhiteList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBlackWhiteListResponse), nil
	}
}

// DeleteBlackWhiteListInvoker 删除黑白名单规则
func (c *CfwClient) DeleteBlackWhiteListInvoker(request *model.DeleteBlackWhiteListRequest) *DeleteBlackWhiteListInvoker {
	requestDef := GenReqDefForDeleteBlackWhiteList()
	return &DeleteBlackWhiteListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCaptureTask 批量删除抓包任务
//
// 批量删除抓包任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteCaptureTask(request *model.DeleteCaptureTaskRequest) (*model.DeleteCaptureTaskResponse, error) {
	requestDef := GenReqDefForDeleteCaptureTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCaptureTaskResponse), nil
	}
}

// DeleteCaptureTaskInvoker 批量删除抓包任务
func (c *CfwClient) DeleteCaptureTaskInvoker(request *model.DeleteCaptureTaskRequest) *DeleteCaptureTaskInvoker {
	requestDef := GenReqDefForDeleteCaptureTask()
	return &DeleteCaptureTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDomainSet 删除域名组
//
// 删除域名组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteDomainSet(request *model.DeleteDomainSetRequest) (*model.DeleteDomainSetResponse, error) {
	requestDef := GenReqDefForDeleteDomainSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDomainSetResponse), nil
	}
}

// DeleteDomainSetInvoker 删除域名组
func (c *CfwClient) DeleteDomainSetInvoker(request *model.DeleteDomainSetRequest) *DeleteDomainSetInvoker {
	requestDef := GenReqDefForDeleteDomainSet()
	return &DeleteDomainSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDomains 删除域名列表
//
// 删除域名列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteDomains(request *model.DeleteDomainsRequest) (*model.DeleteDomainsResponse, error) {
	requestDef := GenReqDefForDeleteDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDomainsResponse), nil
	}
}

// DeleteDomainsInvoker 删除域名列表
func (c *CfwClient) DeleteDomainsInvoker(request *model.DeleteDomainsRequest) *DeleteDomainsInvoker {
	requestDef := GenReqDefForDeleteDomains()
	return &DeleteDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFirewall 删除防火墙
//
// 删除防火墙，仅按需生效
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteFirewall(request *model.DeleteFirewallRequest) (*model.DeleteFirewallResponse, error) {
	requestDef := GenReqDefForDeleteFirewall()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFirewallResponse), nil
	}
}

// DeleteFirewallInvoker 删除防火墙
func (c *CfwClient) DeleteFirewallInvoker(request *model.DeleteFirewallRequest) *DeleteFirewallInvoker {
	requestDef := GenReqDefForDeleteFirewall()
	return &DeleteFirewallInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteServiceItem 删除服务成员
//
// 删除服务组成员
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteServiceItem(request *model.DeleteServiceItemRequest) (*model.DeleteServiceItemResponse, error) {
	requestDef := GenReqDefForDeleteServiceItem()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServiceItemResponse), nil
	}
}

// DeleteServiceItemInvoker 删除服务成员
func (c *CfwClient) DeleteServiceItemInvoker(request *model.DeleteServiceItemRequest) *DeleteServiceItemInvoker {
	requestDef := GenReqDefForDeleteServiceItem()
	return &DeleteServiceItemInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteServiceSet 删除服务组
//
// 删除服务组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteServiceSet(request *model.DeleteServiceSetRequest) (*model.DeleteServiceSetResponse, error) {
	requestDef := GenReqDefForDeleteServiceSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServiceSetResponse), nil
	}
}

// DeleteServiceSetInvoker 删除服务组
func (c *CfwClient) DeleteServiceSetInvoker(request *model.DeleteServiceSetRequest) *DeleteServiceSetInvoker {
	requestDef := GenReqDefForDeleteServiceSet()
	return &DeleteServiceSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTag 删除标签
//
// 删除标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteTag(request *model.DeleteTagRequest) (*model.DeleteTagResponse, error) {
	requestDef := GenReqDefForDeleteTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTagResponse), nil
	}
}

// DeleteTagInvoker 删除标签
func (c *CfwClient) DeleteTagInvoker(request *model.DeleteTagRequest) *DeleteTagInvoker {
	requestDef := GenReqDefForDeleteTag()
	return &DeleteTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccessControlLogs 查询访问控制日志
//
// 查询访问控制日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListAccessControlLogs(request *model.ListAccessControlLogsRequest) (*model.ListAccessControlLogsResponse, error) {
	requestDef := GenReqDefForListAccessControlLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAccessControlLogsResponse), nil
	}
}

// ListAccessControlLogsInvoker 查询访问控制日志
func (c *CfwClient) ListAccessControlLogsInvoker(request *model.ListAccessControlLogsRequest) *ListAccessControlLogsInvoker {
	requestDef := GenReqDefForListAccessControlLogs()
	return &ListAccessControlLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAddressItems 查询地址组成员
//
// 查询地址组成员
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListAddressItems(request *model.ListAddressItemsRequest) (*model.ListAddressItemsResponse, error) {
	requestDef := GenReqDefForListAddressItems()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAddressItemsResponse), nil
	}
}

// ListAddressItemsInvoker 查询地址组成员
func (c *CfwClient) ListAddressItemsInvoker(request *model.ListAddressItemsRequest) *ListAddressItemsInvoker {
	requestDef := GenReqDefForListAddressItems()
	return &ListAddressItemsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAddressSetDetail 查询地址组详细信息
//
// 查询地址组详细
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListAddressSetDetail(request *model.ListAddressSetDetailRequest) (*model.ListAddressSetDetailResponse, error) {
	requestDef := GenReqDefForListAddressSetDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAddressSetDetailResponse), nil
	}
}

// ListAddressSetDetailInvoker 查询地址组详细信息
func (c *CfwClient) ListAddressSetDetailInvoker(request *model.ListAddressSetDetailRequest) *ListAddressSetDetailInvoker {
	requestDef := GenReqDefForListAddressSetDetail()
	return &ListAddressSetDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAddressSets 查询地址组列表
//
// 查询地址组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListAddressSets(request *model.ListAddressSetsRequest) (*model.ListAddressSetsResponse, error) {
	requestDef := GenReqDefForListAddressSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAddressSetsResponse), nil
	}
}

// ListAddressSetsInvoker 查询地址组列表
func (c *CfwClient) ListAddressSetsInvoker(request *model.ListAddressSetsRequest) *ListAddressSetsInvoker {
	requestDef := GenReqDefForListAddressSets()
	return &ListAddressSetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAttackLogs 查询攻击日志
//
// 查询攻击日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListAttackLogs(request *model.ListAttackLogsRequest) (*model.ListAttackLogsResponse, error) {
	requestDef := GenReqDefForListAttackLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAttackLogsResponse), nil
	}
}

// ListAttackLogsInvoker 查询攻击日志
func (c *CfwClient) ListAttackLogsInvoker(request *model.ListAttackLogsRequest) *ListAttackLogsInvoker {
	requestDef := GenReqDefForListAttackLogs()
	return &ListAttackLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBlackWhiteLists 查询黑白名单列表
//
// 查询黑白名单列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListBlackWhiteLists(request *model.ListBlackWhiteListsRequest) (*model.ListBlackWhiteListsResponse, error) {
	requestDef := GenReqDefForListBlackWhiteLists()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBlackWhiteListsResponse), nil
	}
}

// ListBlackWhiteListsInvoker 查询黑白名单列表
func (c *CfwClient) ListBlackWhiteListsInvoker(request *model.ListBlackWhiteListsRequest) *ListBlackWhiteListsInvoker {
	requestDef := GenReqDefForListBlackWhiteLists()
	return &ListBlackWhiteListsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCaptureResult 获取抓包任务结果
//
// 获取抓包任务结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListCaptureResult(request *model.ListCaptureResultRequest) (*model.ListCaptureResultResponse, error) {
	requestDef := GenReqDefForListCaptureResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCaptureResultResponse), nil
	}
}

// ListCaptureResultInvoker 获取抓包任务结果
func (c *CfwClient) ListCaptureResultInvoker(request *model.ListCaptureResultRequest) *ListCaptureResultInvoker {
	requestDef := GenReqDefForListCaptureResult()
	return &ListCaptureResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCaptureTask 查询抓包任务
//
// 查询抓包任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListCaptureTask(request *model.ListCaptureTaskRequest) (*model.ListCaptureTaskResponse, error) {
	requestDef := GenReqDefForListCaptureTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCaptureTaskResponse), nil
	}
}

// ListCaptureTaskInvoker 查询抓包任务
func (c *CfwClient) ListCaptureTaskInvoker(request *model.ListCaptureTaskRequest) *ListCaptureTaskInvoker {
	requestDef := GenReqDefForListCaptureTask()
	return &ListCaptureTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDnsServers 查询dns服务器列表
//
// 查询dns服务器列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListDnsServers(request *model.ListDnsServersRequest) (*model.ListDnsServersResponse, error) {
	requestDef := GenReqDefForListDnsServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDnsServersResponse), nil
	}
}

// ListDnsServersInvoker 查询dns服务器列表
func (c *CfwClient) ListDnsServersInvoker(request *model.ListDnsServersRequest) *ListDnsServersInvoker {
	requestDef := GenReqDefForListDnsServers()
	return &ListDnsServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDomainParseDetail 查询域名解析ip地址
//
// 测试域名有效性
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListDomainParseDetail(request *model.ListDomainParseDetailRequest) (*model.ListDomainParseDetailResponse, error) {
	requestDef := GenReqDefForListDomainParseDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainParseDetailResponse), nil
	}
}

// ListDomainParseDetailInvoker 查询域名解析ip地址
func (c *CfwClient) ListDomainParseDetailInvoker(request *model.ListDomainParseDetailRequest) *ListDomainParseDetailInvoker {
	requestDef := GenReqDefForListDomainParseDetail()
	return &ListDomainParseDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDomainParseIp 获取域名地址解析结果
//
// 获取域名地址解析结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListDomainParseIp(request *model.ListDomainParseIpRequest) (*model.ListDomainParseIpResponse, error) {
	requestDef := GenReqDefForListDomainParseIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainParseIpResponse), nil
	}
}

// ListDomainParseIpInvoker 获取域名地址解析结果
func (c *CfwClient) ListDomainParseIpInvoker(request *model.ListDomainParseIpRequest) *ListDomainParseIpInvoker {
	requestDef := GenReqDefForListDomainParseIp()
	return &ListDomainParseIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDomainSets 查询域名组列表
//
// 查询域名组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListDomainSets(request *model.ListDomainSetsRequest) (*model.ListDomainSetsResponse, error) {
	requestDef := GenReqDefForListDomainSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainSetsResponse), nil
	}
}

// ListDomainSetsInvoker 查询域名组列表
func (c *CfwClient) ListDomainSetsInvoker(request *model.ListDomainSetsRequest) *ListDomainSetsInvoker {
	requestDef := GenReqDefForListDomainSets()
	return &ListDomainSetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDomains 获取域名组下域名列表
//
// 获取域名组下域名列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListDomains(request *model.ListDomainsRequest) (*model.ListDomainsResponse, error) {
	requestDef := GenReqDefForListDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainsResponse), nil
	}
}

// ListDomainsInvoker 获取域名组下域名列表
func (c *CfwClient) ListDomainsInvoker(request *model.ListDomainsRequest) *ListDomainsInvoker {
	requestDef := GenReqDefForListDomains()
	return &ListDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEastWestFirewall 获取东西向防火墙信息
//
// 获取东西向防火墙信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListEastWestFirewall(request *model.ListEastWestFirewallRequest) (*model.ListEastWestFirewallResponse, error) {
	requestDef := GenReqDefForListEastWestFirewall()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEastWestFirewallResponse), nil
	}
}

// ListEastWestFirewallInvoker 获取东西向防火墙信息
func (c *CfwClient) ListEastWestFirewallInvoker(request *model.ListEastWestFirewallRequest) *ListEastWestFirewallInvoker {
	requestDef := GenReqDefForListEastWestFirewall()
	return &ListEastWestFirewallInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFirewallDetail 查询防火墙详细信息
//
// 查询防火墙实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListFirewallDetail(request *model.ListFirewallDetailRequest) (*model.ListFirewallDetailResponse, error) {
	requestDef := GenReqDefForListFirewallDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFirewallDetailResponse), nil
	}
}

// ListFirewallDetailInvoker 查询防火墙详细信息
func (c *CfwClient) ListFirewallDetailInvoker(request *model.ListFirewallDetailRequest) *ListFirewallDetailInvoker {
	requestDef := GenReqDefForListFirewallDetail()
	return &ListFirewallDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFirewallList 查询防火墙列表
//
// 查询防火墙列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListFirewallList(request *model.ListFirewallListRequest) (*model.ListFirewallListResponse, error) {
	requestDef := GenReqDefForListFirewallList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFirewallListResponse), nil
	}
}

// ListFirewallListInvoker 查询防火墙列表
func (c *CfwClient) ListFirewallListInvoker(request *model.ListFirewallListRequest) *ListFirewallListInvoker {
	requestDef := GenReqDefForListFirewallList()
	return &ListFirewallListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlowLogs 查询流日志
//
// 查询流日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListFlowLogs(request *model.ListFlowLogsRequest) (*model.ListFlowLogsResponse, error) {
	requestDef := GenReqDefForListFlowLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlowLogsResponse), nil
	}
}

// ListFlowLogsInvoker 查询流日志
func (c *CfwClient) ListFlowLogsInvoker(request *model.ListFlowLogsRequest) *ListFlowLogsInvoker {
	requestDef := GenReqDefForListFlowLogs()
	return &ListFlowLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListJob 获取CFW任务执行状态
//
// 获取CFW任务执行状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListJob(request *model.ListJobRequest) (*model.ListJobResponse, error) {
	requestDef := GenReqDefForListJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJobResponse), nil
	}
}

// ListJobInvoker 获取CFW任务执行状态
func (c *CfwClient) ListJobInvoker(request *model.ListJobRequest) *ListJobInvoker {
	requestDef := GenReqDefForListJob()
	return &ListJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLogConfig 获取日志配置
//
// 获取日志配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListLogConfig(request *model.ListLogConfigRequest) (*model.ListLogConfigResponse, error) {
	requestDef := GenReqDefForListLogConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLogConfigResponse), nil
	}
}

// ListLogConfigInvoker 获取日志配置
func (c *CfwClient) ListLogConfigInvoker(request *model.ListLogConfigRequest) *ListLogConfigInvoker {
	requestDef := GenReqDefForListLogConfig()
	return &ListLogConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectTags 查询标签信息
//
// 查询标签信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListProjectTags(request *model.ListProjectTagsRequest) (*model.ListProjectTagsResponse, error) {
	requestDef := GenReqDefForListProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectTagsResponse), nil
	}
}

// ListProjectTagsInvoker 查询标签信息
func (c *CfwClient) ListProjectTagsInvoker(request *model.ListProjectTagsRequest) *ListProjectTagsInvoker {
	requestDef := GenReqDefForListProjectTags()
	return &ListProjectTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProtectedVpcs 查询防护VPC数
//
// 查询防护vpc信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListProtectedVpcs(request *model.ListProtectedVpcsRequest) (*model.ListProtectedVpcsResponse, error) {
	requestDef := GenReqDefForListProtectedVpcs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProtectedVpcsResponse), nil
	}
}

// ListProtectedVpcsInvoker 查询防护VPC数
func (c *CfwClient) ListProtectedVpcsInvoker(request *model.ListProtectedVpcsRequest) *ListProtectedVpcsInvoker {
	requestDef := GenReqDefForListProtectedVpcs()
	return &ListProtectedVpcsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceTags 查询资源标签信息
//
// 查询资源标签信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListResourceTags(request *model.ListResourceTagsRequest) (*model.ListResourceTagsResponse, error) {
	requestDef := GenReqDefForListResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceTagsResponse), nil
	}
}

// ListResourceTagsInvoker 查询资源标签信息
func (c *CfwClient) ListResourceTagsInvoker(request *model.ListResourceTagsRequest) *ListResourceTagsInvoker {
	requestDef := GenReqDefForListResourceTags()
	return &ListResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServiceItems 查询服务成员列表
//
// 查询服务组成员列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListServiceItems(request *model.ListServiceItemsRequest) (*model.ListServiceItemsResponse, error) {
	requestDef := GenReqDefForListServiceItems()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceItemsResponse), nil
	}
}

// ListServiceItemsInvoker 查询服务成员列表
func (c *CfwClient) ListServiceItemsInvoker(request *model.ListServiceItemsRequest) *ListServiceItemsInvoker {
	requestDef := GenReqDefForListServiceItems()
	return &ListServiceItemsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServiceSetDetail 查询服务组详情
//
// 查询服务组细节
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListServiceSetDetail(request *model.ListServiceSetDetailRequest) (*model.ListServiceSetDetailResponse, error) {
	requestDef := GenReqDefForListServiceSetDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceSetDetailResponse), nil
	}
}

// ListServiceSetDetailInvoker 查询服务组详情
func (c *CfwClient) ListServiceSetDetailInvoker(request *model.ListServiceSetDetailRequest) *ListServiceSetDetailInvoker {
	requestDef := GenReqDefForListServiceSetDetail()
	return &ListServiceSetDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServiceSets 获取服务组列表
//
// 获取服务组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListServiceSets(request *model.ListServiceSetsRequest) (*model.ListServiceSetsResponse, error) {
	requestDef := GenReqDefForListServiceSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceSetsResponse), nil
	}
}

// ListServiceSetsInvoker 获取服务组列表
func (c *CfwClient) ListServiceSetsInvoker(request *model.ListServiceSetsRequest) *ListServiceSetsInvoker {
	requestDef := GenReqDefForListServiceSets()
	return &ListServiceSetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SaveTags 保存资源标签接口
//
// 保存资源标签接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) SaveTags(request *model.SaveTagsRequest) (*model.SaveTagsResponse, error) {
	requestDef := GenReqDefForSaveTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SaveTagsResponse), nil
	}
}

// SaveTagsInvoker 保存资源标签接口
func (c *CfwClient) SaveTagsInvoker(request *model.SaveTagsRequest) *SaveTagsInvoker {
	requestDef := GenReqDefForSaveTags()
	return &SaveTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAlarmConfig 获取告警配置信息
//
// 获取告警配置信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ShowAlarmConfig(request *model.ShowAlarmConfigRequest) (*model.ShowAlarmConfigResponse, error) {
	requestDef := GenReqDefForShowAlarmConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAlarmConfigResponse), nil
	}
}

// ShowAlarmConfigInvoker 获取告警配置信息
func (c *CfwClient) ShowAlarmConfigInvoker(request *model.ShowAlarmConfigRequest) *ShowAlarmConfigInvoker {
	requestDef := GenReqDefForShowAlarmConfig()
	return &ShowAlarmConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAntiVirusRule 获取防火墙反病毒规则信息
//
// 获取防火墙反病毒规则信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ShowAntiVirusRule(request *model.ShowAntiVirusRuleRequest) (*model.ShowAntiVirusRuleResponse, error) {
	requestDef := GenReqDefForShowAntiVirusRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAntiVirusRuleResponse), nil
	}
}

// ShowAntiVirusRuleInvoker 获取防火墙反病毒规则信息
func (c *CfwClient) ShowAntiVirusRuleInvoker(request *model.ShowAntiVirusRuleRequest) *ShowAntiVirusRuleInvoker {
	requestDef := GenReqDefForShowAntiVirusRule()
	return &ShowAntiVirusRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAntiVirusSwitch 查看反病毒开关
//
// 查看反病毒开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ShowAntiVirusSwitch(request *model.ShowAntiVirusSwitchRequest) (*model.ShowAntiVirusSwitchResponse, error) {
	requestDef := GenReqDefForShowAntiVirusSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAntiVirusSwitchResponse), nil
	}
}

// ShowAntiVirusSwitchInvoker 查看反病毒开关
func (c *CfwClient) ShowAntiVirusSwitchInvoker(request *model.ShowAntiVirusSwitchRequest) *ShowAntiVirusSwitchInvoker {
	requestDef := GenReqDefForShowAntiVirusSwitch()
	return &ShowAntiVirusSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDomainSetDetail 查看域名组详情
//
// 查看域名组详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ShowDomainSetDetail(request *model.ShowDomainSetDetailRequest) (*model.ShowDomainSetDetailResponse, error) {
	requestDef := GenReqDefForShowDomainSetDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainSetDetailResponse), nil
	}
}

// ShowDomainSetDetailInvoker 查看域名组详情
func (c *CfwClient) ShowDomainSetDetailInvoker(request *model.ShowDomainSetDetailRequest) *ShowDomainSetDetailInvoker {
	requestDef := GenReqDefForShowDomainSetDetail()
	return &ShowDomainSetDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAddressSet 更新地址组信息
//
// 更新地址组信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateAddressSet(request *model.UpdateAddressSetRequest) (*model.UpdateAddressSetResponse, error) {
	requestDef := GenReqDefForUpdateAddressSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAddressSetResponse), nil
	}
}

// UpdateAddressSetInvoker 更新地址组信息
func (c *CfwClient) UpdateAddressSetInvoker(request *model.UpdateAddressSetRequest) *UpdateAddressSetInvoker {
	requestDef := GenReqDefForUpdateAddressSet()
	return &UpdateAddressSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAlarmConfig 修改告警配置接口
//
// 修改告警配置接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateAlarmConfig(request *model.UpdateAlarmConfigRequest) (*model.UpdateAlarmConfigResponse, error) {
	requestDef := GenReqDefForUpdateAlarmConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAlarmConfigResponse), nil
	}
}

// UpdateAlarmConfigInvoker 修改告警配置接口
func (c *CfwClient) UpdateAlarmConfigInvoker(request *model.UpdateAlarmConfigRequest) *UpdateAlarmConfigInvoker {
	requestDef := GenReqDefForUpdateAlarmConfig()
	return &UpdateAlarmConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAntiVirusRule 修改反病毒规则
//
// 修改反病毒规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateAntiVirusRule(request *model.UpdateAntiVirusRuleRequest) (*model.UpdateAntiVirusRuleResponse, error) {
	requestDef := GenReqDefForUpdateAntiVirusRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAntiVirusRuleResponse), nil
	}
}

// UpdateAntiVirusRuleInvoker 修改反病毒规则
func (c *CfwClient) UpdateAntiVirusRuleInvoker(request *model.UpdateAntiVirusRuleRequest) *UpdateAntiVirusRuleInvoker {
	requestDef := GenReqDefForUpdateAntiVirusRule()
	return &UpdateAntiVirusRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAntiVirusSwitch 修改反病毒开关
//
// 修改反病毒开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateAntiVirusSwitch(request *model.UpdateAntiVirusSwitchRequest) (*model.UpdateAntiVirusSwitchResponse, error) {
	requestDef := GenReqDefForUpdateAntiVirusSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAntiVirusSwitchResponse), nil
	}
}

// UpdateAntiVirusSwitchInvoker 修改反病毒开关
func (c *CfwClient) UpdateAntiVirusSwitchInvoker(request *model.UpdateAntiVirusSwitchRequest) *UpdateAntiVirusSwitchInvoker {
	requestDef := GenReqDefForUpdateAntiVirusSwitch()
	return &UpdateAntiVirusSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateBlackWhiteList 更新黑白名单列表
//
// 更新黑白名单列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateBlackWhiteList(request *model.UpdateBlackWhiteListRequest) (*model.UpdateBlackWhiteListResponse, error) {
	requestDef := GenReqDefForUpdateBlackWhiteList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBlackWhiteListResponse), nil
	}
}

// UpdateBlackWhiteListInvoker 更新黑白名单列表
func (c *CfwClient) UpdateBlackWhiteListInvoker(request *model.UpdateBlackWhiteListRequest) *UpdateBlackWhiteListInvoker {
	requestDef := GenReqDefForUpdateBlackWhiteList()
	return &UpdateBlackWhiteListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDnsServers 更新dns服务器列表
//
// 更新dns服务器列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateDnsServers(request *model.UpdateDnsServersRequest) (*model.UpdateDnsServersResponse, error) {
	requestDef := GenReqDefForUpdateDnsServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDnsServersResponse), nil
	}
}

// UpdateDnsServersInvoker 更新dns服务器列表
func (c *CfwClient) UpdateDnsServersInvoker(request *model.UpdateDnsServersRequest) *UpdateDnsServersInvoker {
	requestDef := GenReqDefForUpdateDnsServers()
	return &UpdateDnsServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDomainSet 更新域名组
//
// 更新域名组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateDomainSet(request *model.UpdateDomainSetRequest) (*model.UpdateDomainSetResponse, error) {
	requestDef := GenReqDefForUpdateDomainSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDomainSetResponse), nil
	}
}

// UpdateDomainSetInvoker 更新域名组
func (c *CfwClient) UpdateDomainSetInvoker(request *model.UpdateDomainSetRequest) *UpdateDomainSetInvoker {
	requestDef := GenReqDefForUpdateDomainSet()
	return &UpdateDomainSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLogConfig 更新日志配置
//
// 更新日志配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateLogConfig(request *model.UpdateLogConfigRequest) (*model.UpdateLogConfigResponse, error) {
	requestDef := GenReqDefForUpdateLogConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLogConfigResponse), nil
	}
}

// UpdateLogConfigInvoker 更新日志配置
func (c *CfwClient) UpdateLogConfigInvoker(request *model.UpdateLogConfigRequest) *UpdateLogConfigInvoker {
	requestDef := GenReqDefForUpdateLogConfig()
	return &UpdateLogConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateServiceSet 修改服务组
//
// 更新服务组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateServiceSet(request *model.UpdateServiceSetRequest) (*model.UpdateServiceSetResponse, error) {
	requestDef := GenReqDefForUpdateServiceSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateServiceSetResponse), nil
	}
}

// UpdateServiceSetInvoker 修改服务组
func (c *CfwClient) UpdateServiceSetInvoker(request *model.UpdateServiceSetRequest) *UpdateServiceSetInvoker {
	requestDef := GenReqDefForUpdateServiceSet()
	return &UpdateServiceSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddAclRule 创建ACL规则
//
// 创建ACL规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) AddAclRule(request *model.AddAclRuleRequest) (*model.AddAclRuleResponse, error) {
	requestDef := GenReqDefForAddAclRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddAclRuleResponse), nil
	}
}

// AddAclRuleInvoker 创建ACL规则
func (c *CfwClient) AddAclRuleInvoker(request *model.AddAclRuleRequest) *AddAclRuleInvoker {
	requestDef := GenReqDefForAddAclRule()
	return &AddAclRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteAclRules 批量删除Acl规则
//
// 批量删除Acl规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) BatchDeleteAclRules(request *model.BatchDeleteAclRulesRequest) (*model.BatchDeleteAclRulesResponse, error) {
	requestDef := GenReqDefForBatchDeleteAclRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteAclRulesResponse), nil
	}
}

// BatchDeleteAclRulesInvoker 批量删除Acl规则
func (c *CfwClient) BatchDeleteAclRulesInvoker(request *model.BatchDeleteAclRulesRequest) *BatchDeleteAclRulesInvoker {
	requestDef := GenReqDefForBatchDeleteAclRules()
	return &BatchDeleteAclRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateAclRuleActions 批量更新规则动作
//
// 批量更新规则动作
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) BatchUpdateAclRuleActions(request *model.BatchUpdateAclRuleActionsRequest) (*model.BatchUpdateAclRuleActionsResponse, error) {
	requestDef := GenReqDefForBatchUpdateAclRuleActions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateAclRuleActionsResponse), nil
	}
}

// BatchUpdateAclRuleActionsInvoker 批量更新规则动作
func (c *CfwClient) BatchUpdateAclRuleActionsInvoker(request *model.BatchUpdateAclRuleActionsRequest) *BatchUpdateAclRuleActionsInvoker {
	requestDef := GenReqDefForBatchUpdateAclRuleActions()
	return &BatchUpdateAclRuleActionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAclRule 删除ACL规则
//
// 删除ACL规则组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteAclRule(request *model.DeleteAclRuleRequest) (*model.DeleteAclRuleResponse, error) {
	requestDef := GenReqDefForDeleteAclRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAclRuleResponse), nil
	}
}

// DeleteAclRuleInvoker 删除ACL规则
func (c *CfwClient) DeleteAclRuleInvoker(request *model.DeleteAclRuleRequest) *DeleteAclRuleInvoker {
	requestDef := GenReqDefForDeleteAclRule()
	return &DeleteAclRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAclRuleHitCount 删除规则击中次数
//
// 清除规则击中次数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteAclRuleHitCount(request *model.DeleteAclRuleHitCountRequest) (*model.DeleteAclRuleHitCountResponse, error) {
	requestDef := GenReqDefForDeleteAclRuleHitCount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAclRuleHitCountResponse), nil
	}
}

// DeleteAclRuleHitCountInvoker 删除规则击中次数
func (c *CfwClient) DeleteAclRuleHitCountInvoker(request *model.DeleteAclRuleHitCountRequest) *DeleteAclRuleHitCountInvoker {
	requestDef := GenReqDefForDeleteAclRuleHitCount()
	return &DeleteAclRuleHitCountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAclRuleHitCount 获取规则击中次数
//
// 获取规则击中次数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListAclRuleHitCount(request *model.ListAclRuleHitCountRequest) (*model.ListAclRuleHitCountResponse, error) {
	requestDef := GenReqDefForListAclRuleHitCount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAclRuleHitCountResponse), nil
	}
}

// ListAclRuleHitCountInvoker 获取规则击中次数
func (c *CfwClient) ListAclRuleHitCountInvoker(request *model.ListAclRuleHitCountRequest) *ListAclRuleHitCountInvoker {
	requestDef := GenReqDefForListAclRuleHitCount()
	return &ListAclRuleHitCountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAclRules 查询防护规则
//
// 查询防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListAclRules(request *model.ListAclRulesRequest) (*model.ListAclRulesResponse, error) {
	requestDef := GenReqDefForListAclRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAclRulesResponse), nil
	}
}

// ListAclRulesInvoker 查询防护规则
func (c *CfwClient) ListAclRulesInvoker(request *model.ListAclRulesRequest) *ListAclRulesInvoker {
	requestDef := GenReqDefForListAclRules()
	return &ListAclRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRegions 查看region列表
//
// 查看region列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListRegions(request *model.ListRegionsRequest) (*model.ListRegionsResponse, error) {
	requestDef := GenReqDefForListRegions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRegionsResponse), nil
	}
}

// ListRegionsInvoker 查看region列表
func (c *CfwClient) ListRegionsInvoker(request *model.ListRegionsRequest) *ListRegionsInvoker {
	requestDef := GenReqDefForListRegions()
	return &ListRegionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRuleAclTags 查询规则标签
//
// 查询规则标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListRuleAclTags(request *model.ListRuleAclTagsRequest) (*model.ListRuleAclTagsResponse, error) {
	requestDef := GenReqDefForListRuleAclTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRuleAclTagsResponse), nil
	}
}

// ListRuleAclTagsInvoker 查询规则标签
func (c *CfwClient) ListRuleAclTagsInvoker(request *model.ListRuleAclTagsRequest) *ListRuleAclTagsInvoker {
	requestDef := GenReqDefForListRuleAclTags()
	return &ListRuleAclTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowImportStatus 查看导入状态接口
//
// 查看导入状态接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ShowImportStatus(request *model.ShowImportStatusRequest) (*model.ShowImportStatusResponse, error) {
	requestDef := GenReqDefForShowImportStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowImportStatusResponse), nil
	}
}

// ShowImportStatusInvoker 查看导入状态接口
func (c *CfwClient) ShowImportStatusInvoker(request *model.ShowImportStatusRequest) *ShowImportStatusInvoker {
	requestDef := GenReqDefForShowImportStatus()
	return &ShowImportStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAclRule 更新ACL规则
//
// 更新ACL规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateAclRule(request *model.UpdateAclRuleRequest) (*model.UpdateAclRuleResponse, error) {
	requestDef := GenReqDefForUpdateAclRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAclRuleResponse), nil
	}
}

// UpdateAclRuleInvoker 更新ACL规则
func (c *CfwClient) UpdateAclRuleInvoker(request *model.UpdateAclRuleRequest) *UpdateAclRuleInvoker {
	requestDef := GenReqDefForUpdateAclRule()
	return &UpdateAclRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAclRuleOrder ACL防护规则优先级设置
//
// # ACL防护规则优先级设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateAclRuleOrder(request *model.UpdateAclRuleOrderRequest) (*model.UpdateAclRuleOrderResponse, error) {
	requestDef := GenReqDefForUpdateAclRuleOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAclRuleOrderResponse), nil
	}
}

// UpdateAclRuleOrderInvoker ACL防护规则优先级设置
func (c *CfwClient) UpdateAclRuleOrderInvoker(request *model.UpdateAclRuleOrderRequest) *UpdateAclRuleOrderInvoker {
	requestDef := GenReqDefForUpdateAclRuleOrder()
	return &UpdateAclRuleOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeEipStatus 弹性IP开启关闭
//
// 开启关闭EIP，客户购买EIP后首次开启EIP防护前需使用ListEips同步EIP资产，sync字段设置为1。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ChangeEipStatus(request *model.ChangeEipStatusRequest) (*model.ChangeEipStatusResponse, error) {
	requestDef := GenReqDefForChangeEipStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeEipStatusResponse), nil
	}
}

// ChangeEipStatusInvoker 弹性IP开启关闭
func (c *CfwClient) ChangeEipStatusInvoker(request *model.ChangeEipStatusRequest) *ChangeEipStatusInvoker {
	requestDef := GenReqDefForChangeEipStatus()
	return &ChangeEipStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlarmWhitelist 查看eip告警白名单
//
// 查看eip告警白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListAlarmWhitelist(request *model.ListAlarmWhitelistRequest) (*model.ListAlarmWhitelistResponse, error) {
	requestDef := GenReqDefForListAlarmWhitelist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmWhitelistResponse), nil
	}
}

// ListAlarmWhitelistInvoker 查看eip告警白名单
func (c *CfwClient) ListAlarmWhitelistInvoker(request *model.ListAlarmWhitelistRequest) *ListAlarmWhitelistInvoker {
	requestDef := GenReqDefForListAlarmWhitelist()
	return &ListAlarmWhitelistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEipCount 查询Eip个数
//
// 查询Eip个数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListEipCount(request *model.ListEipCountRequest) (*model.ListEipCountResponse, error) {
	requestDef := GenReqDefForListEipCount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEipCountResponse), nil
	}
}

// ListEipCountInvoker 查询Eip个数
func (c *CfwClient) ListEipCountInvoker(request *model.ListEipCountRequest) *ListEipCountInvoker {
	requestDef := GenReqDefForListEipCount()
	return &ListEipCountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEips 弹性IP列表查询
//
// 弹性IP列表查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListEips(request *model.ListEipsRequest) (*model.ListEipsResponse, error) {
	requestDef := GenReqDefForListEips()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEipsResponse), nil
	}
}

// ListEipsInvoker 弹性IP列表查询
func (c *CfwClient) ListEipsInvoker(request *model.ListEipsRequest) *ListEipsInvoker {
	requestDef := GenReqDefForListEips()
	return &ListEipsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAutoProtectStatus 获取eip自动防护状态信息
//
// 获取eip自动防护状态信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ShowAutoProtectStatus(request *model.ShowAutoProtectStatusRequest) (*model.ShowAutoProtectStatusResponse, error) {
	requestDef := GenReqDefForShowAutoProtectStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAutoProtectStatusResponse), nil
	}
}

// ShowAutoProtectStatusInvoker 获取eip自动防护状态信息
func (c *CfwClient) ShowAutoProtectStatusInvoker(request *model.ShowAutoProtectStatusRequest) *ShowAutoProtectStatusInvoker {
	requestDef := GenReqDefForShowAutoProtectStatus()
	return &ShowAutoProtectStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchAutoProtectStatus 修改eip自动防护开关
//
// 修改eip自动防护开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) SwitchAutoProtectStatus(request *model.SwitchAutoProtectStatusRequest) (*model.SwitchAutoProtectStatusResponse, error) {
	requestDef := GenReqDefForSwitchAutoProtectStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchAutoProtectStatusResponse), nil
	}
}

// SwitchAutoProtectStatusInvoker 修改eip自动防护开关
func (c *CfwClient) SwitchAutoProtectStatusInvoker(request *model.SwitchAutoProtectStatusRequest) *SwitchAutoProtectStatusInvoker {
	requestDef := GenReqDefForSwitchAutoProtectStatus()
	return &SwitchAutoProtectStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeIpsProtectMode 切换防护模式
//
// 切换防护模式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ChangeIpsProtectMode(request *model.ChangeIpsProtectModeRequest) (*model.ChangeIpsProtectModeResponse, error) {
	requestDef := GenReqDefForChangeIpsProtectMode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeIpsProtectModeResponse), nil
	}
}

// ChangeIpsProtectModeInvoker 切换防护模式
func (c *CfwClient) ChangeIpsProtectModeInvoker(request *model.ChangeIpsProtectModeRequest) *ChangeIpsProtectModeInvoker {
	requestDef := GenReqDefForChangeIpsProtectMode()
	return &ChangeIpsProtectModeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeIpsRuleMode 改变ips规则模式
//
// 改变ips规则模式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ChangeIpsRuleMode(request *model.ChangeIpsRuleModeRequest) (*model.ChangeIpsRuleModeResponse, error) {
	requestDef := GenReqDefForChangeIpsRuleMode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeIpsRuleModeResponse), nil
	}
}

// ChangeIpsRuleModeInvoker 改变ips规则模式
func (c *CfwClient) ChangeIpsRuleModeInvoker(request *model.ChangeIpsRuleModeRequest) *ChangeIpsRuleModeInvoker {
	requestDef := GenReqDefForChangeIpsRuleMode()
	return &ChangeIpsRuleModeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeIpsSwitchStatus IPS特性开关操作
//
// 切换开关状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ChangeIpsSwitchStatus(request *model.ChangeIpsSwitchStatusRequest) (*model.ChangeIpsSwitchStatusResponse, error) {
	requestDef := GenReqDefForChangeIpsSwitchStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeIpsSwitchStatusResponse), nil
	}
}

// ChangeIpsSwitchStatusInvoker IPS特性开关操作
func (c *CfwClient) ChangeIpsSwitchStatusInvoker(request *model.ChangeIpsSwitchStatusRequest) *ChangeIpsSwitchStatusInvoker {
	requestDef := GenReqDefForChangeIpsSwitchStatus()
	return &ChangeIpsSwitchStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCustomerIps 获取自定义ips规则
//
// 获取自定义ips规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListCustomerIps(request *model.ListCustomerIpsRequest) (*model.ListCustomerIpsResponse, error) {
	requestDef := GenReqDefForListCustomerIps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomerIpsResponse), nil
	}
}

// ListCustomerIpsInvoker 获取自定义ips规则
func (c *CfwClient) ListCustomerIpsInvoker(request *model.ListCustomerIpsRequest) *ListCustomerIpsInvoker {
	requestDef := GenReqDefForListCustomerIps()
	return &ListCustomerIpsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIpsProtectMode 查询防护模式
//
// 查询防护模式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListIpsProtectMode(request *model.ListIpsProtectModeRequest) (*model.ListIpsProtectModeResponse, error) {
	requestDef := GenReqDefForListIpsProtectMode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIpsProtectModeResponse), nil
	}
}

// ListIpsProtectModeInvoker 查询防护模式
func (c *CfwClient) ListIpsProtectModeInvoker(request *model.ListIpsProtectModeRequest) *ListIpsProtectModeInvoker {
	requestDef := GenReqDefForListIpsProtectMode()
	return &ListIpsProtectModeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIpsRules 查询频率ips规则信息
//
// 查询频率ips规则信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListIpsRules(request *model.ListIpsRulesRequest) (*model.ListIpsRulesResponse, error) {
	requestDef := GenReqDefForListIpsRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIpsRulesResponse), nil
	}
}

// ListIpsRulesInvoker 查询频率ips规则信息
func (c *CfwClient) ListIpsRulesInvoker(request *model.ListIpsRulesRequest) *ListIpsRulesInvoker {
	requestDef := GenReqDefForListIpsRules()
	return &ListIpsRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIpsRules1 获取ips规则列表
//
// 获取ips规则列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListIpsRules1(request *model.ListIpsRules1Request) (*model.ListIpsRules1Response, error) {
	requestDef := GenReqDefForListIpsRules1()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIpsRules1Response), nil
	}
}

// ListIpsRules1Invoker 获取ips规则列表
func (c *CfwClient) ListIpsRules1Invoker(request *model.ListIpsRules1Request) *ListIpsRules1Invoker {
	requestDef := GenReqDefForListIpsRules1()
	return &ListIpsRules1Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIpsSwitchStatus 查询IPS特性开关状态
//
// 查询IPS特性开关状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListIpsSwitchStatus(request *model.ListIpsSwitchStatusRequest) (*model.ListIpsSwitchStatusResponse, error) {
	requestDef := GenReqDefForListIpsSwitchStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIpsSwitchStatusResponse), nil
	}
}

// ListIpsSwitchStatusInvoker 查询IPS特性开关状态
func (c *CfwClient) ListIpsSwitchStatusInvoker(request *model.ListIpsSwitchStatusRequest) *ListIpsSwitchStatusInvoker {
	requestDef := GenReqDefForListIpsSwitchStatus()
	return &ListIpsSwitchStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIpsUpdateTime 获取ips规则细节
//
// 获取ips规则细节
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ShowIpsUpdateTime(request *model.ShowIpsUpdateTimeRequest) (*model.ShowIpsUpdateTimeResponse, error) {
	requestDef := GenReqDefForShowIpsUpdateTime()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIpsUpdateTimeResponse), nil
	}
}

// ShowIpsUpdateTimeInvoker 获取ips规则细节
func (c *CfwClient) ShowIpsUpdateTimeInvoker(request *model.ShowIpsUpdateTimeRequest) *ShowIpsUpdateTimeInvoker {
	requestDef := GenReqDefForShowIpsUpdateTime()
	return &ShowIpsUpdateTimeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAdvancedIpsRule 创建频率ips规则
//
// 创建频率ips规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateAdvancedIpsRule(request *model.UpdateAdvancedIpsRuleRequest) (*model.UpdateAdvancedIpsRuleResponse, error) {
	requestDef := GenReqDefForUpdateAdvancedIpsRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAdvancedIpsRuleResponse), nil
	}
}

// UpdateAdvancedIpsRuleInvoker 创建频率ips规则
func (c *CfwClient) UpdateAdvancedIpsRuleInvoker(request *model.UpdateAdvancedIpsRuleRequest) *UpdateAdvancedIpsRuleInvoker {
	requestDef := GenReqDefForUpdateAdvancedIpsRule()
	return &UpdateAdvancedIpsRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
