package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ucs/v1/model"
)

type UcsClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewUcsClient(hcClient *httpclient.HcHttpClient) *UcsClient {
	return &UcsClient{HcClient: hcClient}
}

func UcsClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// CreateAddonInstance 安装插件实例
//
// 安装插件实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) CreateAddonInstance(request *model.CreateAddonInstanceRequest) (*model.CreateAddonInstanceResponse, error) {
	requestDef := GenReqDefForCreateAddonInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAddonInstanceResponse), nil
	}
}

// CreateAddonInstanceInvoker 安装插件实例
func (c *UcsClient) CreateAddonInstanceInvoker(request *model.CreateAddonInstanceRequest) *CreateAddonInstanceInvoker {
	requestDef := GenReqDefForCreateAddonInstance()
	return &CreateAddonInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateClusterGroupPolicyInstance 创建舰队策略实例
//
// 创建舰队策略实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) CreateClusterGroupPolicyInstance(request *model.CreateClusterGroupPolicyInstanceRequest) (*model.CreateClusterGroupPolicyInstanceResponse, error) {
	requestDef := GenReqDefForCreateClusterGroupPolicyInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterGroupPolicyInstanceResponse), nil
	}
}

// CreateClusterGroupPolicyInstanceInvoker 创建舰队策略实例
func (c *UcsClient) CreateClusterGroupPolicyInstanceInvoker(request *model.CreateClusterGroupPolicyInstanceRequest) *CreateClusterGroupPolicyInstanceInvoker {
	requestDef := GenReqDefForCreateClusterGroupPolicyInstance()
	return &CreateClusterGroupPolicyInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateConfigSet 创建配置集合
//
// 创建配置集合
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) CreateConfigSet(request *model.CreateConfigSetRequest) (*model.CreateConfigSetResponse, error) {
	requestDef := GenReqDefForCreateConfigSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConfigSetResponse), nil
	}
}

// CreateConfigSetInvoker 创建配置集合
func (c *UcsClient) CreateConfigSetInvoker(request *model.CreateConfigSetRequest) *CreateConfigSetInvoker {
	requestDef := GenReqDefForCreateConfigSet()
	return &CreateConfigSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFederationCert 创建联邦网络连接并下载联邦kubeconfig
//
// 舰队开通联邦后，调用此接口，创建vpcep终端节点，连接到联邦apiserver，并下载联邦apiserver的kubeconfig
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) CreateFederationCert(request *model.CreateFederationCertRequest) (*model.CreateFederationCertResponse, error) {
	requestDef := GenReqDefForCreateFederationCert()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFederationCertResponse), nil
	}
}

// CreateFederationCertInvoker 创建联邦网络连接并下载联邦kubeconfig
func (c *UcsClient) CreateFederationCertInvoker(request *model.CreateFederationCertRequest) *CreateFederationCertInvoker {
	requestDef := GenReqDefForCreateFederationCert()
	return &CreateFederationCertInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFederationConnection 创建联邦网络连接
//
// 舰队开通联邦后，创建vpcep终端节点连接到联邦apiserver
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) CreateFederationConnection(request *model.CreateFederationConnectionRequest) (*model.CreateFederationConnectionResponse, error) {
	requestDef := GenReqDefForCreateFederationConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFederationConnectionResponse), nil
	}
}

// CreateFederationConnectionInvoker 创建联邦网络连接
func (c *UcsClient) CreateFederationConnectionInvoker(request *model.CreateFederationConnectionRequest) *CreateFederationConnectionInvoker {
	requestDef := GenReqDefForCreateFederationConnection()
	return &CreateFederationConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateProxyUnitedWorkload 创建联邦工作负载
//
// 创建联邦工作负载
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) CreateProxyUnitedWorkload(request *model.CreateProxyUnitedWorkloadRequest) (*model.CreateProxyUnitedWorkloadResponse, error) {
	requestDef := GenReqDefForCreateProxyUnitedWorkload()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProxyUnitedWorkloadResponse), nil
	}
}

// CreateProxyUnitedWorkloadInvoker 创建联邦工作负载
func (c *UcsClient) CreateProxyUnitedWorkloadInvoker(request *model.CreateProxyUnitedWorkloadRequest) *CreateProxyUnitedWorkloadInvoker {
	requestDef := GenReqDefForCreateProxyUnitedWorkload()
	return &CreateProxyUnitedWorkloadInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRecordSet 创建域名解析记录集
//
// 创建域名解析记录集
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) CreateRecordSet(request *model.CreateRecordSetRequest) (*model.CreateRecordSetResponse, error) {
	requestDef := GenReqDefForCreateRecordSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRecordSetResponse), nil
	}
}

// CreateRecordSetInvoker 创建域名解析记录集
func (c *UcsClient) CreateRecordSetInvoker(request *model.CreateRecordSetRequest) *CreateRecordSetInvoker {
	requestDef := GenReqDefForCreateRecordSet()
	return &CreateRecordSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRepo 创建仓库源
//
// 创建仓库源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) CreateRepo(request *model.CreateRepoRequest) (*model.CreateRepoResponse, error) {
	requestDef := GenReqDefForCreateRepo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRepoResponse), nil
	}
}

// CreateRepoInvoker 创建仓库源
func (c *UcsClient) CreateRepoInvoker(request *model.CreateRepoRequest) *CreateRepoInvoker {
	requestDef := GenReqDefForCreateRepo()
	return &CreateRepoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRule 创建权限策略
//
// 创建权限策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) CreateRule(request *model.CreateRuleRequest) (*model.CreateRuleResponse, error) {
	requestDef := GenReqDefForCreateRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRuleResponse), nil
	}
}

// CreateRuleInvoker 创建权限策略
func (c *UcsClient) CreateRuleInvoker(request *model.CreateRuleRequest) *CreateRuleInvoker {
	requestDef := GenReqDefForCreateRule()
	return &CreateRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAddonInstance 卸载插件实例
//
// 卸载插件实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) DeleteAddonInstance(request *model.DeleteAddonInstanceRequest) (*model.DeleteAddonInstanceResponse, error) {
	requestDef := GenReqDefForDeleteAddonInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAddonInstanceResponse), nil
	}
}

// DeleteAddonInstanceInvoker 卸载插件实例
func (c *UcsClient) DeleteAddonInstanceInvoker(request *model.DeleteAddonInstanceRequest) *DeleteAddonInstanceInvoker {
	requestDef := GenReqDefForDeleteAddonInstance()
	return &DeleteAddonInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteClusterGroup 删除容器舰队
//
// 容器舰队删除接口，只有在容器舰队为空时才可以删除该容器舰队，若需删除请先将集群移出容器舰队；传入的cluster ID必须符合k8s UUID的格式规则；同时需要用户有对应集群的操作权限，否则会鉴权失败
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) DeleteClusterGroup(request *model.DeleteClusterGroupRequest) (*model.DeleteClusterGroupResponse, error) {
	requestDef := GenReqDefForDeleteClusterGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClusterGroupResponse), nil
	}
}

// DeleteClusterGroupInvoker 删除容器舰队
func (c *UcsClient) DeleteClusterGroupInvoker(request *model.DeleteClusterGroupRequest) *DeleteClusterGroupInvoker {
	requestDef := GenReqDefForDeleteClusterGroup()
	return &DeleteClusterGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteConfigSet 删除配置集合
//
// 仅删除配置集合，不删除仓库源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) DeleteConfigSet(request *model.DeleteConfigSetRequest) (*model.DeleteConfigSetResponse, error) {
	requestDef := GenReqDefForDeleteConfigSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConfigSetResponse), nil
	}
}

// DeleteConfigSetInvoker 删除配置集合
func (c *UcsClient) DeleteConfigSetInvoker(request *model.DeleteConfigSetRequest) *DeleteConfigSetInvoker {
	requestDef := GenReqDefForDeleteConfigSet()
	return &DeleteConfigSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePolicyInstance 删除指定策略实例
//
// 删除指定策略实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) DeletePolicyInstance(request *model.DeletePolicyInstanceRequest) (*model.DeletePolicyInstanceResponse, error) {
	requestDef := GenReqDefForDeletePolicyInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePolicyInstanceResponse), nil
	}
}

// DeletePolicyInstanceInvoker 删除指定策略实例
func (c *UcsClient) DeletePolicyInstanceInvoker(request *model.DeletePolicyInstanceRequest) *DeletePolicyInstanceInvoker {
	requestDef := GenReqDefForDeletePolicyInstance()
	return &DeletePolicyInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteProxyUnitedWorkload 删除联邦工作负载
//
// 删除联邦工作负载
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) DeleteProxyUnitedWorkload(request *model.DeleteProxyUnitedWorkloadRequest) (*model.DeleteProxyUnitedWorkloadResponse, error) {
	requestDef := GenReqDefForDeleteProxyUnitedWorkload()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteProxyUnitedWorkloadResponse), nil
	}
}

// DeleteProxyUnitedWorkloadInvoker 删除联邦工作负载
func (c *UcsClient) DeleteProxyUnitedWorkloadInvoker(request *model.DeleteProxyUnitedWorkloadRequest) *DeleteProxyUnitedWorkloadInvoker {
	requestDef := GenReqDefForDeleteProxyUnitedWorkload()
	return &DeleteProxyUnitedWorkloadInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRepo 删除仓库源
//
// 删除仓库源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) DeleteRepo(request *model.DeleteRepoRequest) (*model.DeleteRepoResponse, error) {
	requestDef := GenReqDefForDeleteRepo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRepoResponse), nil
	}
}

// DeleteRepoInvoker 删除仓库源
func (c *UcsClient) DeleteRepoInvoker(request *model.DeleteRepoRequest) *DeleteRepoInvoker {
	requestDef := GenReqDefForDeleteRepo()
	return &DeleteRepoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRule 删除权限策略
//
// 删除权限策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) DeleteRule(request *model.DeleteRuleRequest) (*model.DeleteRuleResponse, error) {
	requestDef := GenReqDefForDeleteRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRuleResponse), nil
	}
}

// DeleteRuleInvoker 删除权限策略
func (c *UcsClient) DeleteRuleInvoker(request *model.DeleteRuleRequest) *DeleteRuleInvoker {
	requestDef := GenReqDefForDeleteRule()
	return &DeleteRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableClustergroupPolicy 舰队关闭策略中心
//
// 舰队关闭策略中心
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) DisableClustergroupPolicy(request *model.DisableClustergroupPolicyRequest) (*model.DisableClustergroupPolicyResponse, error) {
	requestDef := GenReqDefForDisableClustergroupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableClustergroupPolicyResponse), nil
	}
}

// DisableClustergroupPolicyInvoker 舰队关闭策略中心
func (c *UcsClient) DisableClustergroupPolicyInvoker(request *model.DisableClustergroupPolicyRequest) *DisableClustergroupPolicyInvoker {
	requestDef := GenReqDefForDisableClustergroupPolicy()
	return &DisableClustergroupPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableFederation 关闭容器集群联邦
//
// 关闭容器舰队联邦
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) DisableFederation(request *model.DisableFederationRequest) (*model.DisableFederationResponse, error) {
	requestDef := GenReqDefForDisableFederation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableFederationResponse), nil
	}
}

// DisableFederationInvoker 关闭容器集群联邦
func (c *UcsClient) DisableFederationInvoker(request *model.DisableFederationRequest) *DisableFederationInvoker {
	requestDef := GenReqDefForDisableFederation()
	return &DisableFederationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadFederationKubeconfig 下载联邦kubeconfig
//
// 舰队开通联邦并且创建网络连接之后，可以使用此接口下载联邦的kubeconfig
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) DownloadFederationKubeconfig(request *model.DownloadFederationKubeconfigRequest) (*model.DownloadFederationKubeconfigResponse, error) {
	requestDef := GenReqDefForDownloadFederationKubeconfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadFederationKubeconfigResponse), nil
	}
}

// DownloadFederationKubeconfigInvoker 下载联邦kubeconfig
func (c *UcsClient) DownloadFederationKubeconfigInvoker(request *model.DownloadFederationKubeconfigRequest) *DownloadFederationKubeconfigInvoker {
	requestDef := GenReqDefForDownloadFederationKubeconfig()
	return &DownloadFederationKubeconfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableClustergroupPolicy 舰队启用策略中心
//
// 舰队启用策略中心
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) EnableClustergroupPolicy(request *model.EnableClustergroupPolicyRequest) (*model.EnableClustergroupPolicyResponse, error) {
	requestDef := GenReqDefForEnableClustergroupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableClustergroupPolicyResponse), nil
	}
}

// EnableClustergroupPolicyInvoker 舰队启用策略中心
func (c *UcsClient) EnableClustergroupPolicyInvoker(request *model.EnableClustergroupPolicyRequest) *EnableClustergroupPolicyInvoker {
	requestDef := GenReqDefForEnableClustergroupPolicy()
	return &EnableClustergroupPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableFederation 启用容器舰队联邦
//
// 启用容器舰队联邦
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) EnableFederation(request *model.EnableFederationRequest) (*model.EnableFederationResponse, error) {
	requestDef := GenReqDefForEnableFederation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableFederationResponse), nil
	}
}

// EnableFederationInvoker 启用容器舰队联邦
func (c *UcsClient) EnableFederationInvoker(request *model.EnableFederationRequest) *EnableFederationInvoker {
	requestDef := GenReqDefForEnableFederation()
	return &EnableFederationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// JoinGroup 集群加入容器舰队
//
// 集群加入容器舰队
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) JoinGroup(request *model.JoinGroupRequest) (*model.JoinGroupResponse, error) {
	requestDef := GenReqDefForJoinGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.JoinGroupResponse), nil
	}
}

// JoinGroupInvoker 集群加入容器舰队
func (c *UcsClient) JoinGroupInvoker(request *model.JoinGroupRequest) *JoinGroupInvoker {
	requestDef := GenReqDefForJoinGroup()
	return &JoinGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// LeaveGroup 集群移出容器舰队
//
// 集群退出容器舰队
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) LeaveGroup(request *model.LeaveGroupRequest) (*model.LeaveGroupResponse, error) {
	requestDef := GenReqDefForLeaveGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.LeaveGroupResponse), nil
	}
}

// LeaveGroupInvoker 集群移出容器舰队
func (c *UcsClient) LeaveGroupInvoker(request *model.LeaveGroupRequest) *LeaveGroupInvoker {
	requestDef := GenReqDefForLeaveGroup()
	return &LeaveGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAddonInstances 获取插件实例列表
//
// 获取插件实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) ListAddonInstances(request *model.ListAddonInstancesRequest) (*model.ListAddonInstancesResponse, error) {
	requestDef := GenReqDefForListAddonInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAddonInstancesResponse), nil
	}
}

// ListAddonInstancesInvoker 获取插件实例列表
func (c *UcsClient) ListAddonInstancesInvoker(request *model.ListAddonInstancesRequest) *ListAddonInstancesInvoker {
	requestDef := GenReqDefForListAddonInstances()
	return &ListAddonInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAddonTemplates 获取插件模板列表
//
// 获取插件模板列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) ListAddonTemplates(request *model.ListAddonTemplatesRequest) (*model.ListAddonTemplatesResponse, error) {
	requestDef := GenReqDefForListAddonTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAddonTemplatesResponse), nil
	}
}

// ListAddonTemplatesInvoker 获取插件模板列表
func (c *UcsClient) ListAddonTemplatesInvoker(request *model.ListAddonTemplatesRequest) *ListAddonTemplatesInvoker {
	requestDef := GenReqDefForListAddonTemplates()
	return &ListAddonTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusterGroup 获取容器舰队列表
//
// 获取所有容器舰队信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) ListClusterGroup(request *model.ListClusterGroupRequest) (*model.ListClusterGroupResponse, error) {
	requestDef := GenReqDefForListClusterGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterGroupResponse), nil
	}
}

// ListClusterGroupInvoker 获取容器舰队列表
func (c *UcsClient) ListClusterGroupInvoker(request *model.ListClusterGroupRequest) *ListClusterGroupInvoker {
	requestDef := GenReqDefForListClusterGroup()
	return &ListClusterGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConfigSets 获取所有配置集合信息
//
// 获取所有配置集合信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) ListConfigSets(request *model.ListConfigSetsRequest) (*model.ListConfigSetsResponse, error) {
	requestDef := GenReqDefForListConfigSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigSetsResponse), nil
	}
}

// ListConfigSetsInvoker 获取所有配置集合信息
func (c *UcsClient) ListConfigSetsInvoker(request *model.ListConfigSetsRequest) *ListConfigSetsInvoker {
	requestDef := GenReqDefForListConfigSets()
	return &ListConfigSetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicyDefinitions 获取策略定义列表
//
// 获取策略定义列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) ListPolicyDefinitions(request *model.ListPolicyDefinitionsRequest) (*model.ListPolicyDefinitionsResponse, error) {
	requestDef := GenReqDefForListPolicyDefinitions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyDefinitionsResponse), nil
	}
}

// ListPolicyDefinitionsInvoker 获取策略定义列表
func (c *UcsClient) ListPolicyDefinitionsInvoker(request *model.ListPolicyDefinitionsRequest) *ListPolicyDefinitionsInvoker {
	requestDef := GenReqDefForListPolicyDefinitions()
	return &ListPolicyDefinitionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicyInstances 获取策略实例列表
//
// 获取策略实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) ListPolicyInstances(request *model.ListPolicyInstancesRequest) (*model.ListPolicyInstancesResponse, error) {
	requestDef := GenReqDefForListPolicyInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyInstancesResponse), nil
	}
}

// ListPolicyInstancesInvoker 获取策略实例列表
func (c *UcsClient) ListPolicyInstancesInvoker(request *model.ListPolicyInstancesRequest) *ListPolicyInstancesInvoker {
	requestDef := GenReqDefForListPolicyInstances()
	return &ListPolicyInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicyJobs 获取策略实例任务列表
//
// 获取策略实例任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) ListPolicyJobs(request *model.ListPolicyJobsRequest) (*model.ListPolicyJobsResponse, error) {
	requestDef := GenReqDefForListPolicyJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyJobsResponse), nil
	}
}

// ListPolicyJobsInvoker 获取策略实例任务列表
func (c *UcsClient) ListPolicyJobsInvoker(request *model.ListPolicyJobsRequest) *ListPolicyJobsInvoker {
	requestDef := GenReqDefForListPolicyJobs()
	return &ListPolicyJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRecordSets 查询域名解析记录集列表
//
// 查询域名解析记录集列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) ListRecordSets(request *model.ListRecordSetsRequest) (*model.ListRecordSetsResponse, error) {
	requestDef := GenReqDefForListRecordSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecordSetsResponse), nil
	}
}

// ListRecordSetsInvoker 查询域名解析记录集列表
func (c *UcsClient) ListRecordSetsInvoker(request *model.ListRecordSetsRequest) *ListRecordSetsInvoker {
	requestDef := GenReqDefForListRecordSets()
	return &ListRecordSetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepos 获取仓库源列表
//
// 获取仓库源列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) ListRepos(request *model.ListReposRequest) (*model.ListReposResponse, error) {
	requestDef := GenReqDefForListRepos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListReposResponse), nil
	}
}

// ListReposInvoker 获取仓库源列表
func (c *UcsClient) ListReposInvoker(request *model.ListReposRequest) *ListReposInvoker {
	requestDef := GenReqDefForListRepos()
	return &ListReposInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRule 获取权限策略列表
//
// 获取权限策略列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) ListRule(request *model.ListRuleRequest) (*model.ListRuleResponse, error) {
	requestDef := GenReqDefForListRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRuleResponse), nil
	}
}

// ListRuleInvoker 获取权限策略列表
func (c *UcsClient) ListRuleInvoker(request *model.ListRuleRequest) *ListRuleInvoker {
	requestDef := GenReqDefForListRule()
	return &ListRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServerConfigs 查询服务配置信息
//
// 获取各个类型集群的全局配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) ListServerConfigs(request *model.ListServerConfigsRequest) (*model.ListServerConfigsResponse, error) {
	requestDef := GenReqDefForListServerConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServerConfigsResponse), nil
	}
}

// ListServerConfigsInvoker 查询服务配置信息
func (c *UcsClient) ListServerConfigsInvoker(request *model.ListServerConfigsRequest) *ListServerConfigsInvoker {
	requestDef := GenReqDefForListServerConfigs()
	return &ListServerConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RegisterClusterGroup 注册容器舰队
//
// 容器舰队创建API，生成容器舰队时可以选择关联的集群
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) RegisterClusterGroup(request *model.RegisterClusterGroupRequest) (*model.RegisterClusterGroupResponse, error) {
	requestDef := GenReqDefForRegisterClusterGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterClusterGroupResponse), nil
	}
}

// RegisterClusterGroupInvoker 注册容器舰队
func (c *UcsClient) RegisterClusterGroupInvoker(request *model.RegisterClusterGroupRequest) *RegisterClusterGroupInvoker {
	requestDef := GenReqDefForRegisterClusterGroup()
	return &RegisterClusterGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAddonInstance 获取插件实例
//
// 获取插件实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) ShowAddonInstance(request *model.ShowAddonInstanceRequest) (*model.ShowAddonInstanceResponse, error) {
	requestDef := GenReqDefForShowAddonInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAddonInstanceResponse), nil
	}
}

// ShowAddonInstanceInvoker 获取插件实例
func (c *UcsClient) ShowAddonInstanceInvoker(request *model.ShowAddonInstanceRequest) *ShowAddonInstanceInvoker {
	requestDef := GenReqDefForShowAddonInstance()
	return &ShowAddonInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowClusterGroup 获取单个容器舰队
//
// 获取单个容器舰队
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) ShowClusterGroup(request *model.ShowClusterGroupRequest) (*model.ShowClusterGroupResponse, error) {
	requestDef := GenReqDefForShowClusterGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusterGroupResponse), nil
	}
}

// ShowClusterGroupInvoker 获取单个容器舰队
func (c *UcsClient) ShowClusterGroupInvoker(request *model.ShowClusterGroupRequest) *ShowClusterGroupInvoker {
	requestDef := GenReqDefForShowClusterGroup()
	return &ShowClusterGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConfigSet 获取配置集合详细信息
//
// 获取配置集合详细信息，包含仓库源信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) ShowConfigSet(request *model.ShowConfigSetRequest) (*model.ShowConfigSetResponse, error) {
	requestDef := GenReqDefForShowConfigSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConfigSetResponse), nil
	}
}

// ShowConfigSetInvoker 获取配置集合详细信息
func (c *UcsClient) ShowConfigSetInvoker(request *model.ShowConfigSetRequest) *ShowConfigSetInvoker {
	requestDef := GenReqDefForShowConfigSet()
	return &ShowConfigSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFederationProgress 查询联邦开启进度
//
// 查询联邦开启进度
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) ShowFederationProgress(request *model.ShowFederationProgressRequest) (*model.ShowFederationProgressResponse, error) {
	requestDef := GenReqDefForShowFederationProgress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFederationProgressResponse), nil
	}
}

// ShowFederationProgressInvoker 查询联邦开启进度
func (c *UcsClient) ShowFederationProgressInvoker(request *model.ShowFederationProgressRequest) *ShowFederationProgressInvoker {
	requestDef := GenReqDefForShowFederationProgress()
	return &ShowFederationProgressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGitopsStatistics 统计某个用户所有配置集合的运行状态
//
// 统计某个用户所有配置集合的运行状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) ShowGitopsStatistics(request *model.ShowGitopsStatisticsRequest) (*model.ShowGitopsStatisticsResponse, error) {
	requestDef := GenReqDefForShowGitopsStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGitopsStatisticsResponse), nil
	}
}

// ShowGitopsStatisticsInvoker 统计某个用户所有配置集合的运行状态
func (c *UcsClient) ShowGitopsStatisticsInvoker(request *model.ShowGitopsStatisticsRequest) *ShowGitopsStatisticsInvoker {
	requestDef := GenReqDefForShowGitopsStatistics()
	return &ShowGitopsStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPolicyDefinition 获取策略定义
//
// 获取策略定义
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) ShowPolicyDefinition(request *model.ShowPolicyDefinitionRequest) (*model.ShowPolicyDefinitionResponse, error) {
	requestDef := GenReqDefForShowPolicyDefinition()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPolicyDefinitionResponse), nil
	}
}

// ShowPolicyDefinitionInvoker 获取策略定义
func (c *UcsClient) ShowPolicyDefinitionInvoker(request *model.ShowPolicyDefinitionRequest) *ShowPolicyDefinitionInvoker {
	requestDef := GenReqDefForShowPolicyDefinition()
	return &ShowPolicyDefinitionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPolicyInstance 获取指定策略实例
//
// 获取指定策略实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) ShowPolicyInstance(request *model.ShowPolicyInstanceRequest) (*model.ShowPolicyInstanceResponse, error) {
	requestDef := GenReqDefForShowPolicyInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPolicyInstanceResponse), nil
	}
}

// ShowPolicyInstanceInvoker 获取指定策略实例
func (c *UcsClient) ShowPolicyInstanceInvoker(request *model.ShowPolicyInstanceRequest) *ShowPolicyInstanceInvoker {
	requestDef := GenReqDefForShowPolicyInstance()
	return &ShowPolicyInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPolicyJob 获取指定策略实例相关任务
//
// 获取指定策略实例相关任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) ShowPolicyJob(request *model.ShowPolicyJobRequest) (*model.ShowPolicyJobResponse, error) {
	requestDef := GenReqDefForShowPolicyJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPolicyJobResponse), nil
	}
}

// ShowPolicyJobInvoker 获取指定策略实例相关任务
func (c *UcsClient) ShowPolicyJobInvoker(request *model.ShowPolicyJobRequest) *ShowPolicyJobInvoker {
	requestDef := GenReqDefForShowPolicyJob()
	return &ShowPolicyJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProxyUnitedWorkload 查询联邦工作负载
//
// 查询联邦工作负载
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) ShowProxyUnitedWorkload(request *model.ShowProxyUnitedWorkloadRequest) (*model.ShowProxyUnitedWorkloadResponse, error) {
	requestDef := GenReqDefForShowProxyUnitedWorkload()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProxyUnitedWorkloadResponse), nil
	}
}

// ShowProxyUnitedWorkloadInvoker 查询联邦工作负载
func (c *UcsClient) ShowProxyUnitedWorkloadInvoker(request *model.ShowProxyUnitedWorkloadRequest) *ShowProxyUnitedWorkloadInvoker {
	requestDef := GenReqDefForShowProxyUnitedWorkload()
	return &ShowProxyUnitedWorkloadInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQuota 获取配额信息
//
// 获取UCS配额信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) ShowQuota(request *model.ShowQuotaRequest) (*model.ShowQuotaResponse, error) {
	requestDef := GenReqDefForShowQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotaResponse), nil
	}
}

// ShowQuotaInvoker 获取配额信息
func (c *UcsClient) ShowQuotaInvoker(request *model.ShowQuotaRequest) *ShowQuotaInvoker {
	requestDef := GenReqDefForShowQuota()
	return &ShowQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAddonInstance 更新插件实例
//
// 更新插件实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) UpdateAddonInstance(request *model.UpdateAddonInstanceRequest) (*model.UpdateAddonInstanceResponse, error) {
	requestDef := GenReqDefForUpdateAddonInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAddonInstanceResponse), nil
	}
}

// UpdateAddonInstanceInvoker 更新插件实例
func (c *UcsClient) UpdateAddonInstanceInvoker(request *model.UpdateAddonInstanceRequest) *UpdateAddonInstanceInvoker {
	requestDef := GenReqDefForUpdateAddonInstance()
	return &UpdateAddonInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateClusterGroup 更新容器舰队描述信息
//
// 更新集群所属的容器舰队description信息；需要用户对相应容器舰队有更新权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) UpdateClusterGroup(request *model.UpdateClusterGroupRequest) (*model.UpdateClusterGroupResponse, error) {
	requestDef := GenReqDefForUpdateClusterGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateClusterGroupResponse), nil
	}
}

// UpdateClusterGroupInvoker 更新容器舰队描述信息
func (c *UcsClient) UpdateClusterGroupInvoker(request *model.UpdateClusterGroupRequest) *UpdateClusterGroupInvoker {
	requestDef := GenReqDefForUpdateClusterGroup()
	return &UpdateClusterGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateClusterGroupAssociatedClusters 向容器舰队中添加集群
//
// 向容器舰队中添加集群，同批次可以添加一个或多个集群，该接口无法清空或减少管理的集群。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) UpdateClusterGroupAssociatedClusters(request *model.UpdateClusterGroupAssociatedClustersRequest) (*model.UpdateClusterGroupAssociatedClustersResponse, error) {
	requestDef := GenReqDefForUpdateClusterGroupAssociatedClusters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateClusterGroupAssociatedClustersResponse), nil
	}
}

// UpdateClusterGroupAssociatedClustersInvoker 向容器舰队中添加集群
func (c *UcsClient) UpdateClusterGroupAssociatedClustersInvoker(request *model.UpdateClusterGroupAssociatedClustersRequest) *UpdateClusterGroupAssociatedClustersInvoker {
	requestDef := GenReqDefForUpdateClusterGroupAssociatedClusters()
	return &UpdateClusterGroupAssociatedClustersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateClusterGroupAssociatedRules 更新容器舰队关联权限策略
//
// 更新容器舰队关联权限策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) UpdateClusterGroupAssociatedRules(request *model.UpdateClusterGroupAssociatedRulesRequest) (*model.UpdateClusterGroupAssociatedRulesResponse, error) {
	requestDef := GenReqDefForUpdateClusterGroupAssociatedRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateClusterGroupAssociatedRulesResponse), nil
	}
}

// UpdateClusterGroupAssociatedRulesInvoker 更新容器舰队关联权限策略
func (c *UcsClient) UpdateClusterGroupAssociatedRulesInvoker(request *model.UpdateClusterGroupAssociatedRulesRequest) *UpdateClusterGroupAssociatedRulesInvoker {
	requestDef := GenReqDefForUpdateClusterGroupAssociatedRules()
	return &UpdateClusterGroupAssociatedRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateClusterGroupAssociatedZones 更新容器舰队的联邦对应的zone
//
// 更新容器舰队的集群联邦关联的zone
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) UpdateClusterGroupAssociatedZones(request *model.UpdateClusterGroupAssociatedZonesRequest) (*model.UpdateClusterGroupAssociatedZonesResponse, error) {
	requestDef := GenReqDefForUpdateClusterGroupAssociatedZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateClusterGroupAssociatedZonesResponse), nil
	}
}

// UpdateClusterGroupAssociatedZonesInvoker 更新容器舰队的联邦对应的zone
func (c *UcsClient) UpdateClusterGroupAssociatedZonesInvoker(request *model.UpdateClusterGroupAssociatedZonesRequest) *UpdateClusterGroupAssociatedZonesInvoker {
	requestDef := GenReqDefForUpdateClusterGroupAssociatedZones()
	return &UpdateClusterGroupAssociatedZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateConfigSet 更新配置集合信息
//
// 仅更新配置集合，不更新仓库源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) UpdateConfigSet(request *model.UpdateConfigSetRequest) (*model.UpdateConfigSetResponse, error) {
	requestDef := GenReqDefForUpdateConfigSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateConfigSetResponse), nil
	}
}

// UpdateConfigSetInvoker 更新配置集合信息
func (c *UcsClient) UpdateConfigSetInvoker(request *model.UpdateConfigSetRequest) *UpdateConfigSetInvoker {
	requestDef := GenReqDefForUpdateConfigSet()
	return &UpdateConfigSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePolicyDefination 更新策略定义
//
// 更新策略定义
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) UpdatePolicyDefination(request *model.UpdatePolicyDefinationRequest) (*model.UpdatePolicyDefinationResponse, error) {
	requestDef := GenReqDefForUpdatePolicyDefination()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePolicyDefinationResponse), nil
	}
}

// UpdatePolicyDefinationInvoker 更新策略定义
func (c *UcsClient) UpdatePolicyDefinationInvoker(request *model.UpdatePolicyDefinationRequest) *UpdatePolicyDefinationInvoker {
	requestDef := GenReqDefForUpdatePolicyDefination()
	return &UpdatePolicyDefinationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePolicyInstance 更新指定策略实例
//
// 更新指定策略实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) UpdatePolicyInstance(request *model.UpdatePolicyInstanceRequest) (*model.UpdatePolicyInstanceResponse, error) {
	requestDef := GenReqDefForUpdatePolicyInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePolicyInstanceResponse), nil
	}
}

// UpdatePolicyInstanceInvoker 更新指定策略实例
func (c *UcsClient) UpdatePolicyInstanceInvoker(request *model.UpdatePolicyInstanceRequest) *UpdatePolicyInstanceInvoker {
	requestDef := GenReqDefForUpdatePolicyInstance()
	return &UpdatePolicyInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProxyUnitedWorkload 更新联邦工作负载
//
// 更新联邦工作负载
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) UpdateProxyUnitedWorkload(request *model.UpdateProxyUnitedWorkloadRequest) (*model.UpdateProxyUnitedWorkloadResponse, error) {
	requestDef := GenReqDefForUpdateProxyUnitedWorkload()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProxyUnitedWorkloadResponse), nil
	}
}

// UpdateProxyUnitedWorkloadInvoker 更新联邦工作负载
func (c *UcsClient) UpdateProxyUnitedWorkloadInvoker(request *model.UpdateProxyUnitedWorkloadRequest) *UpdateProxyUnitedWorkloadInvoker {
	requestDef := GenReqDefForUpdateProxyUnitedWorkload()
	return &UpdateProxyUnitedWorkloadInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRepo 更新仓库源信息
//
// 更新仓库源信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) UpdateRepo(request *model.UpdateRepoRequest) (*model.UpdateRepoResponse, error) {
	requestDef := GenReqDefForUpdateRepo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRepoResponse), nil
	}
}

// UpdateRepoInvoker 更新仓库源信息
func (c *UcsClient) UpdateRepoInvoker(request *model.UpdateRepoRequest) *UpdateRepoInvoker {
	requestDef := GenReqDefForUpdateRepo()
	return &UpdateRepoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRule 更新权限策略
//
// 更新权限策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) UpdateRule(request *model.UpdateRuleRequest) (*model.UpdateRuleResponse, error) {
	requestDef := GenReqDefForUpdateRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRuleResponse), nil
	}
}

// UpdateRuleInvoker 更新权限策略
func (c *UcsClient) UpdateRuleInvoker(request *model.UpdateRuleRequest) *UpdateRuleInvoker {
	requestDef := GenReqDefForUpdateRule()
	return &UpdateRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpgradeFederation 升级容器舰队联邦
//
// 容器舰队联邦版本升级
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) UpgradeFederation(request *model.UpgradeFederationRequest) (*model.UpgradeFederationResponse, error) {
	requestDef := GenReqDefForUpgradeFederation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpgradeFederationResponse), nil
	}
}

// UpgradeFederationInvoker 升级容器舰队联邦
func (c *UcsClient) UpgradeFederationInvoker(request *model.UpgradeFederationRequest) *UpgradeFederationInvoker {
	requestDef := GenReqDefForUpgradeFederation()
	return &UpgradeFederationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateClusterConf 获取集群安装时所需的配置信息
//
// 获取集群安装时所需的配置信息，当前仅本地集群使用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) CreateClusterConf(request *model.CreateClusterConfRequest) (*model.CreateClusterConfResponse, error) {
	requestDef := GenReqDefForCreateClusterConf()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterConfResponse), nil
	}
}

// CreateClusterConfInvoker 获取集群安装时所需的配置信息
func (c *UcsClient) CreateClusterConfInvoker(request *model.CreateClusterConfRequest) *CreateClusterConfInvoker {
	requestDef := GenReqDefForCreateClusterConf()
	return &CreateClusterConfInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateClusterKubeconfig 获取集群kubeconfig
//
// 获取集群kubeconfig
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) CreateClusterKubeconfig(request *model.CreateClusterKubeconfigRequest) (*model.CreateClusterKubeconfigResponse, error) {
	requestDef := GenReqDefForCreateClusterKubeconfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterKubeconfigResponse), nil
	}
}

// CreateClusterKubeconfigInvoker 获取集群kubeconfig
func (c *UcsClient) CreateClusterKubeconfigInvoker(request *model.CreateClusterKubeconfigRequest) *CreateClusterKubeconfigInvoker {
	requestDef := GenReqDefForCreateClusterKubeconfig()
	return &CreateClusterKubeconfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateClusterPolicyInstance 创建集群建策略实例
//
// 创建集群建策略实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) CreateClusterPolicyInstance(request *model.CreateClusterPolicyInstanceRequest) (*model.CreateClusterPolicyInstanceResponse, error) {
	requestDef := GenReqDefForCreateClusterPolicyInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterPolicyInstanceResponse), nil
	}
}

// CreateClusterPolicyInstanceInvoker 创建集群建策略实例
func (c *UcsClient) CreateClusterPolicyInstanceInvoker(request *model.CreateClusterPolicyInstanceRequest) *CreateClusterPolicyInstanceInvoker {
	requestDef := GenReqDefForCreateClusterPolicyInstance()
	return &CreateClusterPolicyInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCluster 删除集群
//
// 用于集群解除注册；传入的cluster ID必须符合k8s UUID的格式规则；同时需要用户有对应集群的操作权限，否则会鉴权失败。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) DeleteCluster(request *model.DeleteClusterRequest) (*model.DeleteClusterResponse, error) {
	requestDef := GenReqDefForDeleteCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClusterResponse), nil
	}
}

// DeleteClusterInvoker 删除集群
func (c *UcsClient) DeleteClusterInvoker(request *model.DeleteClusterRequest) *DeleteClusterInvoker {
	requestDef := GenReqDefForDeleteCluster()
	return &DeleteClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableClusterPolicy 集群关闭策略中心
//
// 集群关闭策略中心
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) DisableClusterPolicy(request *model.DisableClusterPolicyRequest) (*model.DisableClusterPolicyResponse, error) {
	requestDef := GenReqDefForDisableClusterPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableClusterPolicyResponse), nil
	}
}

// DisableClusterPolicyInvoker 集群关闭策略中心
func (c *UcsClient) DisableClusterPolicyInvoker(request *model.DisableClusterPolicyRequest) *DisableClusterPolicyInvoker {
	requestDef := GenReqDefForDisableClusterPolicy()
	return &DisableClusterPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableGitOps 卸载集群gitops插件
//
// 卸载集群gitops插件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) DisableGitOps(request *model.DisableGitOpsRequest) (*model.DisableGitOpsResponse, error) {
	requestDef := GenReqDefForDisableGitOps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableGitOpsResponse), nil
	}
}

// DisableGitOpsInvoker 卸载集群gitops插件
func (c *UcsClient) DisableGitOpsInvoker(request *model.DisableGitOpsRequest) *DisableGitOpsInvoker {
	requestDef := GenReqDefForDisableGitOps()
	return &DisableGitOpsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableClusterPolicy 集群启用策略中心
//
// 集群启用策略中心
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) EnableClusterPolicy(request *model.EnableClusterPolicyRequest) (*model.EnableClusterPolicyResponse, error) {
	requestDef := GenReqDefForEnableClusterPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableClusterPolicyResponse), nil
	}
}

// EnableClusterPolicyInvoker 集群启用策略中心
func (c *UcsClient) EnableClusterPolicyInvoker(request *model.EnableClusterPolicyRequest) *EnableClusterPolicyInvoker {
	requestDef := GenReqDefForEnableClusterPolicy()
	return &EnableClusterPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableGitOps 启用Gitops插件
//
// 启用Gitops插件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) EnableGitOps(request *model.EnableGitOpsRequest) (*model.EnableGitOpsResponse, error) {
	requestDef := GenReqDefForEnableGitOps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableGitOpsResponse), nil
	}
}

// EnableGitOpsInvoker 启用Gitops插件
func (c *UcsClient) EnableGitOpsInvoker(request *model.EnableGitOpsRequest) *EnableGitOpsInvoker {
	requestDef := GenReqDefForEnableGitOps()
	return &EnableGitOpsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListManagedClusters 获取租户的CCE集群列表
//
// 获取当前租户的CCE集群列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) ListManagedClusters(request *model.ListManagedClustersRequest) (*model.ListManagedClustersResponse, error) {
	requestDef := GenReqDefForListManagedClusters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListManagedClustersResponse), nil
	}
}

// ListManagedClustersInvoker 获取租户的CCE集群列表
func (c *UcsClient) ListManagedClustersInvoker(request *model.ListManagedClustersRequest) *ListManagedClustersInvoker {
	requestDef := GenReqDefForListManagedClusters()
	return &ListManagedClustersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRegisteredClusterVersions 查询支持接入UCS的集群版本列表
//
// 查询支持接入UCS的集群版本列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) ListRegisteredClusterVersions(request *model.ListRegisteredClusterVersionsRequest) (*model.ListRegisteredClusterVersionsResponse, error) {
	requestDef := GenReqDefForListRegisteredClusterVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRegisteredClusterVersionsResponse), nil
	}
}

// ListRegisteredClusterVersionsInvoker 查询支持接入UCS的集群版本列表
func (c *UcsClient) ListRegisteredClusterVersionsInvoker(request *model.ListRegisteredClusterVersionsRequest) *ListRegisteredClusterVersionsInvoker {
	requestDef := GenReqDefForListRegisteredClusterVersions()
	return &ListRegisteredClusterVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RegisterCluster 注册集群
//
// 集群注册接口。支持三方集群的注册和CCE导入集群的注册。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) RegisterCluster(request *model.RegisterClusterRequest) (*model.RegisterClusterResponse, error) {
	requestDef := GenReqDefForRegisterCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterClusterResponse), nil
	}
}

// RegisterClusterInvoker 注册集群
func (c *UcsClient) RegisterClusterInvoker(request *model.RegisterClusterRequest) *RegisterClusterInvoker {
	requestDef := GenReqDefForRegisterCluster()
	return &RegisterClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RetryClusterActivation 激活集群
//
// 激活集群接口；传入的cluster ID必须符合k8s UUID的格式规则；同时需要用户有对应集群的更新权限，否则会鉴权失败
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) RetryClusterActivation(request *model.RetryClusterActivationRequest) (*model.RetryClusterActivationResponse, error) {
	requestDef := GenReqDefForRetryClusterActivation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RetryClusterActivationResponse), nil
	}
}

// RetryClusterActivationInvoker 激活集群
func (c *UcsClient) RetryClusterActivationInvoker(request *model.RetryClusterActivationRequest) *RetryClusterActivationInvoker {
	requestDef := GenReqDefForRetryClusterActivation()
	return &RetryClusterActivationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCluster 获取单个集群
//
// 获取集群信息。传入的cluster ID必须符合k8s UUID的格式规则；同时需要用户有对应集群的获取权限，否则会鉴权失败
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) ShowCluster(request *model.ShowClusterRequest) (*model.ShowClusterResponse, error) {
	requestDef := GenReqDefForShowCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusterResponse), nil
	}
}

// ShowClusterInvoker 获取单个集群
func (c *UcsClient) ShowClusterInvoker(request *model.ShowClusterRequest) *ShowClusterInvoker {
	requestDef := GenReqDefForShowCluster()
	return &ShowClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowClusterAccessInfo 获取集群接入信息
//
// 该API接口用于获取集群接入信息；传入的cluster ID必须符合k8s UUID的格式规则；同时需要用户有对应集群证书的获取权限，否则会鉴权失败；agent证书只可以下载一次。仅用于获取三方集群的集群接入信息，CCE集群不从该接口获取，如果传入CCE集群ID，返回码为400
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) ShowClusterAccessInfo(request *model.ShowClusterAccessInfoRequest) (*model.ShowClusterAccessInfoResponse, error) {
	requestDef := GenReqDefForShowClusterAccessInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusterAccessInfoResponse), nil
	}
}

// ShowClusterAccessInfoInvoker 获取集群接入信息
func (c *UcsClient) ShowClusterAccessInfoInvoker(request *model.ShowClusterAccessInfoRequest) *ShowClusterAccessInfoInvoker {
	requestDef := GenReqDefForShowClusterAccessInfo()
	return &ShowClusterAccessInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowClusterList 获取集群列表
//
// 获取集群信息列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) ShowClusterList(request *model.ShowClusterListRequest) (*model.ShowClusterListResponse, error) {
	requestDef := GenReqDefForShowClusterList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusterListResponse), nil
	}
}

// ShowClusterListInvoker 获取集群列表
func (c *UcsClient) ShowClusterListInvoker(request *model.ShowClusterListRequest) *ShowClusterListInvoker {
	requestDef := GenReqDefForShowClusterList()
	return &ShowClusterListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGitOpsStatus 获取gitops启用进展
//
// 获取gitops启用进展
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) ShowGitOpsStatus(request *model.ShowGitOpsStatusRequest) (*model.ShowGitOpsStatusResponse, error) {
	requestDef := GenReqDefForShowGitOpsStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGitOpsStatusResponse), nil
	}
}

// ShowGitOpsStatusInvoker 获取gitops启用进展
func (c *UcsClient) ShowGitOpsStatusInvoker(request *model.ShowGitOpsStatusRequest) *ShowGitOpsStatusInvoker {
	requestDef := GenReqDefForShowGitOpsStatus()
	return &ShowGitOpsStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCluster 更新集群
//
// 更新集群。当前仅允许更新附着集群和本地集群的国家/城市，允许更新多云集群的工作节点个数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) UpdateCluster(request *model.UpdateClusterRequest) (*model.UpdateClusterResponse, error) {
	requestDef := GenReqDefForUpdateCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateClusterResponse), nil
	}
}

// UpdateClusterInvoker 更新集群
func (c *UcsClient) UpdateClusterInvoker(request *model.UpdateClusterRequest) *UpdateClusterInvoker {
	requestDef := GenReqDefForUpdateCluster()
	return &UpdateClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateClusterRules 集群关联权限策略
//
// 集群关联权限策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UcsClient) UpdateClusterRules(request *model.UpdateClusterRulesRequest) (*model.UpdateClusterRulesResponse, error) {
	requestDef := GenReqDefForUpdateClusterRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateClusterRulesResponse), nil
	}
}

// UpdateClusterRulesInvoker 集群关联权限策略
func (c *UcsClient) UpdateClusterRulesInvoker(request *model.UpdateClusterRulesRequest) *UpdateClusterRulesInvoker {
	requestDef := GenReqDefForUpdateClusterRules()
	return &UpdateClusterRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
