package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cfw/v1/model"
)

type CfwClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCfwClient(hcClient *http_client.HcHttpClient) *CfwClient {
	return &CfwClient{HcClient: hcClient}
}

func CfwClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// AddAddressItemsUsingPost 添加地址组成员
//
// 添加地址组成员
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) AddAddressItemsUsingPost(request *model.AddAddressItemsUsingPostRequest) (*model.AddAddressItemsUsingPostResponse, error) {
	requestDef := GenReqDefForAddAddressItemsUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddAddressItemsUsingPostResponse), nil
	}
}

// AddAddressItemsUsingPostInvoker 添加地址组成员
func (c *CfwClient) AddAddressItemsUsingPostInvoker(request *model.AddAddressItemsUsingPostRequest) *AddAddressItemsUsingPostInvoker {
	requestDef := GenReqDefForAddAddressItemsUsingPost()
	return &AddAddressItemsUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddAddressSetInfoUsingPost 添加地址组
//
// 添加地址组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) AddAddressSetInfoUsingPost(request *model.AddAddressSetInfoUsingPostRequest) (*model.AddAddressSetInfoUsingPostResponse, error) {
	requestDef := GenReqDefForAddAddressSetInfoUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddAddressSetInfoUsingPostResponse), nil
	}
}

// AddAddressSetInfoUsingPostInvoker 添加地址组
func (c *CfwClient) AddAddressSetInfoUsingPostInvoker(request *model.AddAddressSetInfoUsingPostRequest) *AddAddressSetInfoUsingPostInvoker {
	requestDef := GenReqDefForAddAddressSetInfoUsingPost()
	return &AddAddressSetInfoUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddBlackWhiteListUsingPost 创建黑白名单规则
//
// 创建黑白名单规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) AddBlackWhiteListUsingPost(request *model.AddBlackWhiteListUsingPostRequest) (*model.AddBlackWhiteListUsingPostResponse, error) {
	requestDef := GenReqDefForAddBlackWhiteListUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddBlackWhiteListUsingPostResponse), nil
	}
}

// AddBlackWhiteListUsingPostInvoker 创建黑白名单规则
func (c *CfwClient) AddBlackWhiteListUsingPostInvoker(request *model.AddBlackWhiteListUsingPostRequest) *AddBlackWhiteListUsingPostInvoker {
	requestDef := GenReqDefForAddBlackWhiteListUsingPost()
	return &AddBlackWhiteListUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddServiceItemsUsingPost 新建服务成员
//
// 批量添加服务组成员
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) AddServiceItemsUsingPost(request *model.AddServiceItemsUsingPostRequest) (*model.AddServiceItemsUsingPostResponse, error) {
	requestDef := GenReqDefForAddServiceItemsUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddServiceItemsUsingPostResponse), nil
	}
}

// AddServiceItemsUsingPostInvoker 新建服务成员
func (c *CfwClient) AddServiceItemsUsingPostInvoker(request *model.AddServiceItemsUsingPostRequest) *AddServiceItemsUsingPostInvoker {
	requestDef := GenReqDefForAddServiceItemsUsingPost()
	return &AddServiceItemsUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddServiceSetUsingPost 新建服务组
//
// 创建服务组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) AddServiceSetUsingPost(request *model.AddServiceSetUsingPostRequest) (*model.AddServiceSetUsingPostResponse, error) {
	requestDef := GenReqDefForAddServiceSetUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddServiceSetUsingPostResponse), nil
	}
}

// AddServiceSetUsingPostInvoker 新建服务组
func (c *CfwClient) AddServiceSetUsingPostInvoker(request *model.AddServiceSetUsingPostRequest) *AddServiceSetUsingPostInvoker {
	requestDef := GenReqDefForAddServiceSetUsingPost()
	return &AddServiceSetUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeEwProtectStatus 修改东西向防火墙防护状态
//
// 东西向防护资源防护开启/关闭
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ChangeEwProtectStatus(request *model.ChangeEwProtectStatusRequest) (*model.ChangeEwProtectStatusResponse, error) {
	requestDef := GenReqDefForChangeEwProtectStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeEwProtectStatusResponse), nil
	}
}

// ChangeEwProtectStatusInvoker 修改东西向防火墙防护状态
func (c *CfwClient) ChangeEwProtectStatusInvoker(request *model.ChangeEwProtectStatusRequest) *ChangeEwProtectStatusInvoker {
	requestDef := GenReqDefForChangeEwProtectStatus()
	return &ChangeEwProtectStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeIpsProtectModeUsingPost 切换防护模式
//
// 切换防护模式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ChangeIpsProtectModeUsingPost(request *model.ChangeIpsProtectModeUsingPostRequest) (*model.ChangeIpsProtectModeUsingPostResponse, error) {
	requestDef := GenReqDefForChangeIpsProtectModeUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeIpsProtectModeUsingPostResponse), nil
	}
}

// ChangeIpsProtectModeUsingPostInvoker 切换防护模式
func (c *CfwClient) ChangeIpsProtectModeUsingPostInvoker(request *model.ChangeIpsProtectModeUsingPostRequest) *ChangeIpsProtectModeUsingPostInvoker {
	requestDef := GenReqDefForChangeIpsProtectModeUsingPost()
	return &ChangeIpsProtectModeUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAclRuleCount 删除规则击中次数
//
// 清除规则击中次数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteAclRuleCount(request *model.DeleteAclRuleCountRequest) (*model.DeleteAclRuleCountResponse, error) {
	requestDef := GenReqDefForDeleteAclRuleCount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAclRuleCountResponse), nil
	}
}

// DeleteAclRuleCountInvoker 删除规则击中次数
func (c *CfwClient) DeleteAclRuleCountInvoker(request *model.DeleteAclRuleCountRequest) *DeleteAclRuleCountInvoker {
	requestDef := GenReqDefForDeleteAclRuleCount()
	return &DeleteAclRuleCountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAddressItemUsingDelete 删除地址组成员
//
// 删除地址组成员
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteAddressItemUsingDelete(request *model.DeleteAddressItemUsingDeleteRequest) (*model.DeleteAddressItemUsingDeleteResponse, error) {
	requestDef := GenReqDefForDeleteAddressItemUsingDelete()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAddressItemUsingDeleteResponse), nil
	}
}

// DeleteAddressItemUsingDeleteInvoker 删除地址组成员
func (c *CfwClient) DeleteAddressItemUsingDeleteInvoker(request *model.DeleteAddressItemUsingDeleteRequest) *DeleteAddressItemUsingDeleteInvoker {
	requestDef := GenReqDefForDeleteAddressItemUsingDelete()
	return &DeleteAddressItemUsingDeleteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAddressSetInfoUsingDelete 删除地址组
//
// 删除地址组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteAddressSetInfoUsingDelete(request *model.DeleteAddressSetInfoUsingDeleteRequest) (*model.DeleteAddressSetInfoUsingDeleteResponse, error) {
	requestDef := GenReqDefForDeleteAddressSetInfoUsingDelete()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAddressSetInfoUsingDeleteResponse), nil
	}
}

// DeleteAddressSetInfoUsingDeleteInvoker 删除地址组
func (c *CfwClient) DeleteAddressSetInfoUsingDeleteInvoker(request *model.DeleteAddressSetInfoUsingDeleteRequest) *DeleteAddressSetInfoUsingDeleteInvoker {
	requestDef := GenReqDefForDeleteAddressSetInfoUsingDelete()
	return &DeleteAddressSetInfoUsingDeleteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBlackWhiteListUsingDelete 删除黑白名单规则
//
// 删除黑白名单规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteBlackWhiteListUsingDelete(request *model.DeleteBlackWhiteListUsingDeleteRequest) (*model.DeleteBlackWhiteListUsingDeleteResponse, error) {
	requestDef := GenReqDefForDeleteBlackWhiteListUsingDelete()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBlackWhiteListUsingDeleteResponse), nil
	}
}

// DeleteBlackWhiteListUsingDeleteInvoker 删除黑白名单规则
func (c *CfwClient) DeleteBlackWhiteListUsingDeleteInvoker(request *model.DeleteBlackWhiteListUsingDeleteRequest) *DeleteBlackWhiteListUsingDeleteInvoker {
	requestDef := GenReqDefForDeleteBlackWhiteListUsingDelete()
	return &DeleteBlackWhiteListUsingDeleteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteServiceItemUsingDelete 删除服务成员
//
// 删除服务组成员
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteServiceItemUsingDelete(request *model.DeleteServiceItemUsingDeleteRequest) (*model.DeleteServiceItemUsingDeleteResponse, error) {
	requestDef := GenReqDefForDeleteServiceItemUsingDelete()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServiceItemUsingDeleteResponse), nil
	}
}

// DeleteServiceItemUsingDeleteInvoker 删除服务成员
func (c *CfwClient) DeleteServiceItemUsingDeleteInvoker(request *model.DeleteServiceItemUsingDeleteRequest) *DeleteServiceItemUsingDeleteInvoker {
	requestDef := GenReqDefForDeleteServiceItemUsingDelete()
	return &DeleteServiceItemUsingDeleteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteServiceSetUsingDelete 删除服务组
//
// 删除服务组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteServiceSetUsingDelete(request *model.DeleteServiceSetUsingDeleteRequest) (*model.DeleteServiceSetUsingDeleteResponse, error) {
	requestDef := GenReqDefForDeleteServiceSetUsingDelete()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServiceSetUsingDeleteResponse), nil
	}
}

// DeleteServiceSetUsingDeleteInvoker 删除服务组
func (c *CfwClient) DeleteServiceSetUsingDeleteInvoker(request *model.DeleteServiceSetUsingDeleteRequest) *DeleteServiceSetUsingDeleteInvoker {
	requestDef := GenReqDefForDeleteServiceSetUsingDelete()
	return &DeleteServiceSetUsingDeleteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListAddressItemsUsingGet 查询地址组成员
//
// 查询地址组成员
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListAddressItemsUsingGet(request *model.ListAddressItemsUsingGetRequest) (*model.ListAddressItemsUsingGetResponse, error) {
	requestDef := GenReqDefForListAddressItemsUsingGet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAddressItemsUsingGetResponse), nil
	}
}

// ListAddressItemsUsingGetInvoker 查询地址组成员
func (c *CfwClient) ListAddressItemsUsingGetInvoker(request *model.ListAddressItemsUsingGetRequest) *ListAddressItemsUsingGetInvoker {
	requestDef := GenReqDefForListAddressItemsUsingGet()
	return &ListAddressItemsUsingGetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAddressSetDetailUsingGet 查询地址组详细信息
//
// 查询地址组详细
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListAddressSetDetailUsingGet(request *model.ListAddressSetDetailUsingGetRequest) (*model.ListAddressSetDetailUsingGetResponse, error) {
	requestDef := GenReqDefForListAddressSetDetailUsingGet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAddressSetDetailUsingGetResponse), nil
	}
}

// ListAddressSetDetailUsingGetInvoker 查询地址组详细信息
func (c *CfwClient) ListAddressSetDetailUsingGetInvoker(request *model.ListAddressSetDetailUsingGetRequest) *ListAddressSetDetailUsingGetInvoker {
	requestDef := GenReqDefForListAddressSetDetailUsingGet()
	return &ListAddressSetDetailUsingGetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAddressSetListUsingGet 查询地址组列表
//
// 查询地址组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListAddressSetListUsingGet(request *model.ListAddressSetListUsingGetRequest) (*model.ListAddressSetListUsingGetResponse, error) {
	requestDef := GenReqDefForListAddressSetListUsingGet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAddressSetListUsingGetResponse), nil
	}
}

// ListAddressSetListUsingGetInvoker 查询地址组列表
func (c *CfwClient) ListAddressSetListUsingGetInvoker(request *model.ListAddressSetListUsingGetRequest) *ListAddressSetListUsingGetInvoker {
	requestDef := GenReqDefForListAddressSetListUsingGet()
	return &ListAddressSetListUsingGetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListBlackWhiteListsUsingGet 查询黑白名单列表
//
// 查询黑白名单列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListBlackWhiteListsUsingGet(request *model.ListBlackWhiteListsUsingGetRequest) (*model.ListBlackWhiteListsUsingGetResponse, error) {
	requestDef := GenReqDefForListBlackWhiteListsUsingGet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBlackWhiteListsUsingGetResponse), nil
	}
}

// ListBlackWhiteListsUsingGetInvoker 查询黑白名单列表
func (c *CfwClient) ListBlackWhiteListsUsingGetInvoker(request *model.ListBlackWhiteListsUsingGetRequest) *ListBlackWhiteListsUsingGetInvoker {
	requestDef := GenReqDefForListBlackWhiteListsUsingGet()
	return &ListBlackWhiteListsUsingGetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListEastWestFirewall 获取东西向防火墙信息
//
// 查询东西向防火墙信息
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

// ListFirewallUsingGet 查询防火墙实例
//
// 查询防火墙实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListFirewallUsingGet(request *model.ListFirewallUsingGetRequest) (*model.ListFirewallUsingGetResponse, error) {
	requestDef := GenReqDefForListFirewallUsingGet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFirewallUsingGetResponse), nil
	}
}

// ListFirewallUsingGetInvoker 查询防火墙实例
func (c *CfwClient) ListFirewallUsingGetInvoker(request *model.ListFirewallUsingGetRequest) *ListFirewallUsingGetInvoker {
	requestDef := GenReqDefForListFirewallUsingGet()
	return &ListFirewallUsingGetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListIpsProtectModeUsingPost 查询防护模式
//
// 查询防护模式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListIpsProtectModeUsingPost(request *model.ListIpsProtectModeUsingPostRequest) (*model.ListIpsProtectModeUsingPostResponse, error) {
	requestDef := GenReqDefForListIpsProtectModeUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIpsProtectModeUsingPostResponse), nil
	}
}

// ListIpsProtectModeUsingPostInvoker 查询防护模式
func (c *CfwClient) ListIpsProtectModeUsingPostInvoker(request *model.ListIpsProtectModeUsingPostRequest) *ListIpsProtectModeUsingPostInvoker {
	requestDef := GenReqDefForListIpsProtectModeUsingPost()
	return &ListIpsProtectModeUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListParseDomainDetails 查询域名解析ip地址
//
// 测试域名有效性
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListParseDomainDetails(request *model.ListParseDomainDetailsRequest) (*model.ListParseDomainDetailsResponse, error) {
	requestDef := GenReqDefForListParseDomainDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListParseDomainDetailsResponse), nil
	}
}

// ListParseDomainDetailsInvoker 查询域名解析ip地址
func (c *CfwClient) ListParseDomainDetailsInvoker(request *model.ListParseDomainDetailsRequest) *ListParseDomainDetailsInvoker {
	requestDef := GenReqDefForListParseDomainDetails()
	return &ListParseDomainDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRuleHitCount 获取规则击中次数
//
// 获取规则击中次数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListRuleHitCount(request *model.ListRuleHitCountRequest) (*model.ListRuleHitCountResponse, error) {
	requestDef := GenReqDefForListRuleHitCount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRuleHitCountResponse), nil
	}
}

// ListRuleHitCountInvoker 获取规则击中次数
func (c *CfwClient) ListRuleHitCountInvoker(request *model.ListRuleHitCountRequest) *ListRuleHitCountInvoker {
	requestDef := GenReqDefForListRuleHitCount()
	return &ListRuleHitCountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServiceItemsDetails 查询服务成员列表
//
// 查询服务组成员列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListServiceItemsDetails(request *model.ListServiceItemsDetailsRequest) (*model.ListServiceItemsDetailsResponse, error) {
	requestDef := GenReqDefForListServiceItemsDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceItemsDetailsResponse), nil
	}
}

// ListServiceItemsDetailsInvoker 查询服务成员列表
func (c *CfwClient) ListServiceItemsDetailsInvoker(request *model.ListServiceItemsDetailsRequest) *ListServiceItemsDetailsInvoker {
	requestDef := GenReqDefForListServiceItemsDetails()
	return &ListServiceItemsDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServiceSet 获取服务组列表
//
// 获取服务组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListServiceSet(request *model.ListServiceSetRequest) (*model.ListServiceSetResponse, error) {
	requestDef := GenReqDefForListServiceSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceSetResponse), nil
	}
}

// ListServiceSetInvoker 获取服务组列表
func (c *CfwClient) ListServiceSetInvoker(request *model.ListServiceSetRequest) *ListServiceSetInvoker {
	requestDef := GenReqDefForListServiceSet()
	return &ListServiceSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServiceSetDetails 查询服务组详情
//
// 查询服务组细节
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListServiceSetDetails(request *model.ListServiceSetDetailsRequest) (*model.ListServiceSetDetailsResponse, error) {
	requestDef := GenReqDefForListServiceSetDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceSetDetailsResponse), nil
	}
}

// ListServiceSetDetailsInvoker 查询服务组详情
func (c *CfwClient) ListServiceSetDetailsInvoker(request *model.ListServiceSetDetailsRequest) *ListServiceSetDetailsInvoker {
	requestDef := GenReqDefForListServiceSetDetails()
	return &ListServiceSetDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAddressSetInfoUsingPut 更新地址组信息
//
// 更新地址组信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateAddressSetInfoUsingPut(request *model.UpdateAddressSetInfoUsingPutRequest) (*model.UpdateAddressSetInfoUsingPutResponse, error) {
	requestDef := GenReqDefForUpdateAddressSetInfoUsingPut()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAddressSetInfoUsingPutResponse), nil
	}
}

// UpdateAddressSetInfoUsingPutInvoker 更新地址组信息
func (c *CfwClient) UpdateAddressSetInfoUsingPutInvoker(request *model.UpdateAddressSetInfoUsingPutRequest) *UpdateAddressSetInfoUsingPutInvoker {
	requestDef := GenReqDefForUpdateAddressSetInfoUsingPut()
	return &UpdateAddressSetInfoUsingPutInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateBlackWhiteListUsingPut 更新黑白名单列表
//
// 更新黑白名单列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateBlackWhiteListUsingPut(request *model.UpdateBlackWhiteListUsingPutRequest) (*model.UpdateBlackWhiteListUsingPutResponse, error) {
	requestDef := GenReqDefForUpdateBlackWhiteListUsingPut()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBlackWhiteListUsingPutResponse), nil
	}
}

// UpdateBlackWhiteListUsingPutInvoker 更新黑白名单列表
func (c *CfwClient) UpdateBlackWhiteListUsingPutInvoker(request *model.UpdateBlackWhiteListUsingPutRequest) *UpdateBlackWhiteListUsingPutInvoker {
	requestDef := GenReqDefForUpdateBlackWhiteListUsingPut()
	return &UpdateBlackWhiteListUsingPutInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// UpdateServiceSetUsingPut 修改服务组
//
// 更新服务组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateServiceSetUsingPut(request *model.UpdateServiceSetUsingPutRequest) (*model.UpdateServiceSetUsingPutResponse, error) {
	requestDef := GenReqDefForUpdateServiceSetUsingPut()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateServiceSetUsingPutResponse), nil
	}
}

// UpdateServiceSetUsingPutInvoker 修改服务组
func (c *CfwClient) UpdateServiceSetUsingPutInvoker(request *model.UpdateServiceSetUsingPutRequest) *UpdateServiceSetUsingPutInvoker {
	requestDef := GenReqDefForUpdateServiceSetUsingPut()
	return &UpdateServiceSetUsingPutInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddRuleAclUsingPost 创建ACL规则
//
// 创建ACL规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) AddRuleAclUsingPost(request *model.AddRuleAclUsingPostRequest) (*model.AddRuleAclUsingPostResponse, error) {
	requestDef := GenReqDefForAddRuleAclUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddRuleAclUsingPostResponse), nil
	}
}

// AddRuleAclUsingPostInvoker 创建ACL规则
func (c *CfwClient) AddRuleAclUsingPostInvoker(request *model.AddRuleAclUsingPostRequest) *AddRuleAclUsingPostInvoker {
	requestDef := GenReqDefForAddRuleAclUsingPost()
	return &AddRuleAclUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRuleAclUsingDelete 删除ACL规则组
//
// 删除ACL规则组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteRuleAclUsingDelete(request *model.DeleteRuleAclUsingDeleteRequest) (*model.DeleteRuleAclUsingDeleteResponse, error) {
	requestDef := GenReqDefForDeleteRuleAclUsingDelete()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRuleAclUsingDeleteResponse), nil
	}
}

// DeleteRuleAclUsingDeleteInvoker 删除ACL规则组
func (c *CfwClient) DeleteRuleAclUsingDeleteInvoker(request *model.DeleteRuleAclUsingDeleteRequest) *DeleteRuleAclUsingDeleteInvoker {
	requestDef := GenReqDefForDeleteRuleAclUsingDelete()
	return &DeleteRuleAclUsingDeleteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRuleAclUsingPut ACL防护规则优先级设置
//
// ACL防护规则优先级设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListRuleAclUsingPut(request *model.ListRuleAclUsingPutRequest) (*model.ListRuleAclUsingPutResponse, error) {
	requestDef := GenReqDefForListRuleAclUsingPut()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRuleAclUsingPutResponse), nil
	}
}

// ListRuleAclUsingPutInvoker ACL防护规则优先级设置
func (c *CfwClient) ListRuleAclUsingPutInvoker(request *model.ListRuleAclUsingPutRequest) *ListRuleAclUsingPutInvoker {
	requestDef := GenReqDefForListRuleAclUsingPut()
	return &ListRuleAclUsingPutInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRuleAclsUsingGet 查询防护规则
//
// 查询防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListRuleAclsUsingGet(request *model.ListRuleAclsUsingGetRequest) (*model.ListRuleAclsUsingGetResponse, error) {
	requestDef := GenReqDefForListRuleAclsUsingGet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRuleAclsUsingGetResponse), nil
	}
}

// ListRuleAclsUsingGetInvoker 查询防护规则
func (c *CfwClient) ListRuleAclsUsingGetInvoker(request *model.ListRuleAclsUsingGetRequest) *ListRuleAclsUsingGetInvoker {
	requestDef := GenReqDefForListRuleAclsUsingGet()
	return &ListRuleAclsUsingGetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRuleAclUsingPut 更新ACL规则
//
// 更新ACL规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateRuleAclUsingPut(request *model.UpdateRuleAclUsingPutRequest) (*model.UpdateRuleAclUsingPutResponse, error) {
	requestDef := GenReqDefForUpdateRuleAclUsingPut()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRuleAclUsingPutResponse), nil
	}
}

// UpdateRuleAclUsingPutInvoker 更新ACL规则
func (c *CfwClient) UpdateRuleAclUsingPutInvoker(request *model.UpdateRuleAclUsingPutRequest) *UpdateRuleAclUsingPutInvoker {
	requestDef := GenReqDefForUpdateRuleAclUsingPut()
	return &UpdateRuleAclUsingPutInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeProtectEip 弹性IP开启关闭
//
// 开启关闭EIP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ChangeProtectEip(request *model.ChangeProtectEipRequest) (*model.ChangeProtectEipResponse, error) {
	requestDef := GenReqDefForChangeProtectEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeProtectEipResponse), nil
	}
}

// ChangeProtectEipInvoker 弹性IP开启关闭
func (c *CfwClient) ChangeProtectEipInvoker(request *model.ChangeProtectEipRequest) *ChangeProtectEipInvoker {
	requestDef := GenReqDefForChangeProtectEip()
	return &ChangeProtectEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountEips 查询Eip个数
//
// 查询Eip个数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) CountEips(request *model.CountEipsRequest) (*model.CountEipsResponse, error) {
	requestDef := GenReqDefForCountEips()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountEipsResponse), nil
	}
}

// CountEipsInvoker 查询Eip个数
func (c *CfwClient) CountEipsInvoker(request *model.CountEipsRequest) *CountEipsInvoker {
	requestDef := GenReqDefForCountEips()
	return &CountEipsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEipResources 弹性IP列表查询
//
// 弹性IP列表查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListEipResources(request *model.ListEipResourcesRequest) (*model.ListEipResourcesResponse, error) {
	requestDef := GenReqDefForListEipResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEipResourcesResponse), nil
	}
}

// ListEipResourcesInvoker 弹性IP列表查询
func (c *CfwClient) ListEipResourcesInvoker(request *model.ListEipResourcesRequest) *ListEipResourcesInvoker {
	requestDef := GenReqDefForListEipResources()
	return &ListEipResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeIpsSwitchUsingPost IPS特性开关操作
//
// 切换开关状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ChangeIpsSwitchUsingPost(request *model.ChangeIpsSwitchUsingPostRequest) (*model.ChangeIpsSwitchUsingPostResponse, error) {
	requestDef := GenReqDefForChangeIpsSwitchUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeIpsSwitchUsingPostResponse), nil
	}
}

// ChangeIpsSwitchUsingPostInvoker IPS特性开关操作
func (c *CfwClient) ChangeIpsSwitchUsingPostInvoker(request *model.ChangeIpsSwitchUsingPostRequest) *ChangeIpsSwitchUsingPostInvoker {
	requestDef := GenReqDefForChangeIpsSwitchUsingPost()
	return &ChangeIpsSwitchUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIpsSwitchStatusUsingGet 查询IPS特性开关状态
//
// 查询IPS特性开关状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListIpsSwitchStatusUsingGet(request *model.ListIpsSwitchStatusUsingGetRequest) (*model.ListIpsSwitchStatusUsingGetResponse, error) {
	requestDef := GenReqDefForListIpsSwitchStatusUsingGet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIpsSwitchStatusUsingGetResponse), nil
	}
}

// ListIpsSwitchStatusUsingGetInvoker 查询IPS特性开关状态
func (c *CfwClient) ListIpsSwitchStatusUsingGetInvoker(request *model.ListIpsSwitchStatusUsingGetRequest) *ListIpsSwitchStatusUsingGetInvoker {
	requestDef := GenReqDefForListIpsSwitchStatusUsingGet()
	return &ListIpsSwitchStatusUsingGetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVpcProtects 查询防护VPC数
//
// 查询防护vpc信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListVpcProtects(request *model.ListVpcProtectsRequest) (*model.ListVpcProtectsResponse, error) {
	requestDef := GenReqDefForListVpcProtects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpcProtectsResponse), nil
	}
}

// ListVpcProtectsInvoker 查询防护VPC数
func (c *CfwClient) ListVpcProtectsInvoker(request *model.ListVpcProtectsRequest) *ListVpcProtectsInvoker {
	requestDef := GenReqDefForListVpcProtects()
	return &ListVpcProtectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
