package v5

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/hss/v5/model"
)

type HssClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewHssClient(hcClient *httpclient.HcHttpClient) *HssClient {
	return &HssClient{HcClient: hcClient}
}

func HssClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AddBaselineWhiteList 新增基线白名单
//
// 新增基线白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) AddBaselineWhiteList(request *model.AddBaselineWhiteListRequest) (*model.AddBaselineWhiteListResponse, error) {
	requestDef := GenReqDefForAddBaselineWhiteList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddBaselineWhiteListResponse), nil
	}
}

// AddBaselineWhiteListInvoker 新增基线白名单
func (c *HssClient) AddBaselineWhiteListInvoker(request *model.AddBaselineWhiteListRequest) *AddBaselineWhiteListInvoker {
	requestDef := GenReqDefForAddBaselineWhiteList()
	return &AddBaselineWhiteListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddCceIntegrationProtection 新建cce集成防护配置
//
// 新建cce集成防护配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) AddCceIntegrationProtection(request *model.AddCceIntegrationProtectionRequest) (*model.AddCceIntegrationProtectionResponse, error) {
	requestDef := GenReqDefForAddCceIntegrationProtection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddCceIntegrationProtectionResponse), nil
	}
}

// AddCceIntegrationProtectionInvoker 新建cce集成防护配置
func (c *HssClient) AddCceIntegrationProtectionInvoker(request *model.AddCceIntegrationProtectionRequest) *AddCceIntegrationProtectionInvoker {
	requestDef := GenReqDefForAddCceIntegrationProtection()
	return &AddCceIntegrationProtectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddHostsGroup 创建服务器组
//
// 创建服务器组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) AddHostsGroup(request *model.AddHostsGroupRequest) (*model.AddHostsGroupResponse, error) {
	requestDef := GenReqDefForAddHostsGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddHostsGroupResponse), nil
	}
}

// AddHostsGroupInvoker 创建服务器组
func (c *HssClient) AddHostsGroupInvoker(request *model.AddHostsGroupRequest) *AddHostsGroupInvoker {
	requestDef := GenReqDefForAddHostsGroup()
	return &AddHostsGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddLoginWhiteList 添加登录白名单
//
// 添加登录白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) AddLoginWhiteList(request *model.AddLoginWhiteListRequest) (*model.AddLoginWhiteListResponse, error) {
	requestDef := GenReqDefForAddLoginWhiteList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddLoginWhiteListResponse), nil
	}
}

// AddLoginWhiteListInvoker 添加登录白名单
func (c *HssClient) AddLoginWhiteListInvoker(request *model.AddLoginWhiteListRequest) *AddLoginWhiteListInvoker {
	requestDef := GenReqDefForAddLoginWhiteList()
	return &AddLoginWhiteListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddPolicy 添加防护策略
//
// 添加防护策略：创建防护策略，包含策略名称、相关规则开启状态、防护动作以及检测规则配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) AddPolicy(request *model.AddPolicyRequest) (*model.AddPolicyResponse, error) {
	requestDef := GenReqDefForAddPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddPolicyResponse), nil
	}
}

// AddPolicyInvoker 添加防护策略
func (c *HssClient) AddPolicyInvoker(request *model.AddPolicyRequest) *AddPolicyInvoker {
	requestDef := GenReqDefForAddPolicy()
	return &AddPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddProtectionPolicy 添加防护策略
//
// 添加防护策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) AddProtectionPolicy(request *model.AddProtectionPolicyRequest) (*model.AddProtectionPolicyResponse, error) {
	requestDef := GenReqDefForAddProtectionPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddProtectionPolicyResponse), nil
	}
}

// AddProtectionPolicyInvoker 添加防护策略
func (c *HssClient) AddProtectionPolicyInvoker(request *model.AddProtectionPolicyRequest) *AddProtectionPolicyInvoker {
	requestDef := GenReqDefForAddProtectionPolicy()
	return &AddProtectionPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddSystemUserWhiteList 添加系统用户白名单
//
// 添加系统用户白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) AddSystemUserWhiteList(request *model.AddSystemUserWhiteListRequest) (*model.AddSystemUserWhiteListResponse, error) {
	requestDef := GenReqDefForAddSystemUserWhiteList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddSystemUserWhiteListResponse), nil
	}
}

// AddSystemUserWhiteListInvoker 添加系统用户白名单
func (c *HssClient) AddSystemUserWhiteListInvoker(request *model.AddSystemUserWhiteListRequest) *AddSystemUserWhiteListInvoker {
	requestDef := GenReqDefForAddSystemUserWhiteList()
	return &AddSystemUserWhiteListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociatePolicyGroup 部署策略组
//
// 部署策略组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) AssociatePolicyGroup(request *model.AssociatePolicyGroupRequest) (*model.AssociatePolicyGroupResponse, error) {
	requestDef := GenReqDefForAssociatePolicyGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociatePolicyGroupResponse), nil
	}
}

// AssociatePolicyGroupInvoker 部署策略组
func (c *HssClient) AssociatePolicyGroupInvoker(request *model.AssociatePolicyGroupRequest) *AssociatePolicyGroupInvoker {
	requestDef := GenReqDefForAssociatePolicyGroup()
	return &AssociatePolicyGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchAddAccounts 批量添加账号
//
// 批量添加账号
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) BatchAddAccounts(request *model.BatchAddAccountsRequest) (*model.BatchAddAccountsResponse, error) {
	requestDef := GenReqDefForBatchAddAccounts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddAccountsResponse), nil
	}
}

// BatchAddAccountsInvoker 批量添加账号
func (c *HssClient) BatchAddAccountsInvoker(request *model.BatchAddAccountsRequest) *BatchAddAccountsInvoker {
	requestDef := GenReqDefForBatchAddAccounts()
	return &BatchAddAccountsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateTags 批量创建标签
//
// 批量创建标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) BatchCreateTags(request *model.BatchCreateTagsRequest) (*model.BatchCreateTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateTagsResponse), nil
	}
}

// BatchCreateTagsInvoker 批量创建标签
func (c *HssClient) BatchCreateTagsInvoker(request *model.BatchCreateTagsRequest) *BatchCreateTagsInvoker {
	requestDef := GenReqDefForBatchCreateTags()
	return &BatchCreateTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteAgentDaemonset 批量卸载集群daemonset
//
// 批量卸载集群daemonset
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) BatchDeleteAgentDaemonset(request *model.BatchDeleteAgentDaemonsetRequest) (*model.BatchDeleteAgentDaemonsetResponse, error) {
	requestDef := GenReqDefForBatchDeleteAgentDaemonset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteAgentDaemonsetResponse), nil
	}
}

// BatchDeleteAgentDaemonsetInvoker 批量卸载集群daemonset
func (c *HssClient) BatchDeleteAgentDaemonsetInvoker(request *model.BatchDeleteAgentDaemonsetRequest) *BatchDeleteAgentDaemonsetInvoker {
	requestDef := GenReqDefForBatchDeleteAgentDaemonset()
	return &BatchDeleteAgentDaemonsetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchScanSwrImage 镜像仓库镜像批量扫描
//
// 镜像仓库镜像批量扫描
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) BatchScanSwrImage(request *model.BatchScanSwrImageRequest) (*model.BatchScanSwrImageResponse, error) {
	requestDef := GenReqDefForBatchScanSwrImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchScanSwrImageResponse), nil
	}
}

// BatchScanSwrImageInvoker 镜像仓库镜像批量扫描
func (c *HssClient) BatchScanSwrImageInvoker(request *model.BatchScanSwrImageRequest) *BatchScanSwrImageInvoker {
	requestDef := GenReqDefForBatchScanSwrImage()
	return &BatchScanSwrImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchStartProtection 批量开启勒索病毒防护2.0
//
// 批量开启勒索病毒防护,若开启备份防护，请保证该region有cbr云备份服务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) BatchStartProtection(request *model.BatchStartProtectionRequest) (*model.BatchStartProtectionResponse, error) {
	requestDef := GenReqDefForBatchStartProtection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStartProtectionResponse), nil
	}
}

// BatchStartProtectionInvoker 批量开启勒索病毒防护2.0
func (c *HssClient) BatchStartProtectionInvoker(request *model.BatchStartProtectionRequest) *BatchStartProtectionInvoker {
	requestDef := GenReqDefForBatchStartProtection()
	return &BatchStartProtectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpgradeAgentDaemonset 批量升级集群daemonset
//
// 批量升级集群daemonset
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) BatchUpgradeAgentDaemonset(request *model.BatchUpgradeAgentDaemonsetRequest) (*model.BatchUpgradeAgentDaemonsetResponse, error) {
	requestDef := GenReqDefForBatchUpgradeAgentDaemonset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpgradeAgentDaemonsetResponse), nil
	}
}

// BatchUpgradeAgentDaemonsetInvoker 批量升级集群daemonset
func (c *HssClient) BatchUpgradeAgentDaemonsetInvoker(request *model.BatchUpgradeAgentDaemonsetRequest) *BatchUpgradeAgentDaemonsetInvoker {
	requestDef := GenReqDefForBatchUpgradeAgentDaemonset()
	return &BatchUpgradeAgentDaemonsetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeAntivirusPolicy 编辑自定义查杀策略
//
// 编辑自定义查杀策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ChangeAntivirusPolicy(request *model.ChangeAntivirusPolicyRequest) (*model.ChangeAntivirusPolicyResponse, error) {
	requestDef := GenReqDefForChangeAntivirusPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeAntivirusPolicyResponse), nil
	}
}

// ChangeAntivirusPolicyInvoker 编辑自定义查杀策略
func (c *HssClient) ChangeAntivirusPolicyInvoker(request *model.ChangeAntivirusPolicyRequest) *ChangeAntivirusPolicyInvoker {
	requestDef := GenReqDefForChangeAntivirusPolicy()
	return &ChangeAntivirusPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeBaselineWhiteList 修改基线白名单
//
// 修改基线白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ChangeBaselineWhiteList(request *model.ChangeBaselineWhiteListRequest) (*model.ChangeBaselineWhiteListResponse, error) {
	requestDef := GenReqDefForChangeBaselineWhiteList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeBaselineWhiteListResponse), nil
	}
}

// ChangeBaselineWhiteListInvoker 修改基线白名单
func (c *HssClient) ChangeBaselineWhiteListInvoker(request *model.ChangeBaselineWhiteListRequest) *ChangeBaselineWhiteListInvoker {
	requestDef := GenReqDefForChangeBaselineWhiteList()
	return &ChangeBaselineWhiteListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeBlockedIp 解除已拦截IP
//
// 解除已拦截IP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ChangeBlockedIp(request *model.ChangeBlockedIpRequest) (*model.ChangeBlockedIpResponse, error) {
	requestDef := GenReqDefForChangeBlockedIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeBlockedIpResponse), nil
	}
}

// ChangeBlockedIpInvoker 解除已拦截IP
func (c *HssClient) ChangeBlockedIpInvoker(request *model.ChangeBlockedIpRequest) *ChangeBlockedIpInvoker {
	requestDef := GenReqDefForChangeBlockedIp()
	return &ChangeBlockedIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeCheckRuleAction 对未通过的配置检查项进行忽略/取消忽略/修复/验证操作
//
// 对未通过的配置检查项进行忽略/取消忽略/修复/验证操作
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ChangeCheckRuleAction(request *model.ChangeCheckRuleActionRequest) (*model.ChangeCheckRuleActionResponse, error) {
	requestDef := GenReqDefForChangeCheckRuleAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeCheckRuleActionResponse), nil
	}
}

// ChangeCheckRuleActionInvoker 对未通过的配置检查项进行忽略/取消忽略/修复/验证操作
func (c *HssClient) ChangeCheckRuleActionInvoker(request *model.ChangeCheckRuleActionRequest) *ChangeCheckRuleActionInvoker {
	requestDef := GenReqDefForChangeCheckRuleAction()
	return &ChangeCheckRuleActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeClusterEvents 修改告警状态
//
// 修改告警状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ChangeClusterEvents(request *model.ChangeClusterEventsRequest) (*model.ChangeClusterEventsResponse, error) {
	requestDef := GenReqDefForChangeClusterEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeClusterEventsResponse), nil
	}
}

// ChangeClusterEventsInvoker 修改告警状态
func (c *HssClient) ChangeClusterEventsInvoker(request *model.ChangeClusterEventsRequest) *ChangeClusterEventsInvoker {
	requestDef := GenReqDefForChangeClusterEvents()
	return &ChangeClusterEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeClusterProtectionPolicy 修改集群防护策略
//
// 修改集群防护策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ChangeClusterProtectionPolicy(request *model.ChangeClusterProtectionPolicyRequest) (*model.ChangeClusterProtectionPolicyResponse, error) {
	requestDef := GenReqDefForChangeClusterProtectionPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeClusterProtectionPolicyResponse), nil
	}
}

// ChangeClusterProtectionPolicyInvoker 修改集群防护策略
func (c *HssClient) ChangeClusterProtectionPolicyInvoker(request *model.ChangeClusterProtectionPolicyRequest) *ChangeClusterProtectionPolicyInvoker {
	requestDef := GenReqDefForChangeClusterProtectionPolicy()
	return &ChangeClusterProtectionPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeEvent 处理告警事件
//
// 处理告警事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ChangeEvent(request *model.ChangeEventRequest) (*model.ChangeEventResponse, error) {
	requestDef := GenReqDefForChangeEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeEventResponse), nil
	}
}

// ChangeEventInvoker 处理告警事件
func (c *HssClient) ChangeEventInvoker(request *model.ChangeEventRequest) *ChangeEventInvoker {
	requestDef := GenReqDefForChangeEvent()
	return &ChangeEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeHostsGroup 编辑服务器组
//
// 编辑服务器组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ChangeHostsGroup(request *model.ChangeHostsGroupRequest) (*model.ChangeHostsGroupResponse, error) {
	requestDef := GenReqDefForChangeHostsGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeHostsGroupResponse), nil
	}
}

// ChangeHostsGroupInvoker 编辑服务器组
func (c *HssClient) ChangeHostsGroupInvoker(request *model.ChangeHostsGroupRequest) *ChangeHostsGroupInvoker {
	requestDef := GenReqDefForChangeHostsGroup()
	return &ChangeHostsGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeIsolatedFile 恢复已隔离文件
//
// 恢复已隔离文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ChangeIsolatedFile(request *model.ChangeIsolatedFileRequest) (*model.ChangeIsolatedFileResponse, error) {
	requestDef := GenReqDefForChangeIsolatedFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeIsolatedFileResponse), nil
	}
}

// ChangeIsolatedFileInvoker 恢复已隔离文件
func (c *HssClient) ChangeIsolatedFileInvoker(request *model.ChangeIsolatedFileRequest) *ChangeIsolatedFileInvoker {
	requestDef := GenReqDefForChangeIsolatedFile()
	return &ChangeIsolatedFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangePasswordComplexityStatus 对口令复杂度检测未通过的主机进行忽略/取消忽略
//
// 对口令复杂度检测未通过的主机进行忽略/取消忽略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ChangePasswordComplexityStatus(request *model.ChangePasswordComplexityStatusRequest) (*model.ChangePasswordComplexityStatusResponse, error) {
	requestDef := GenReqDefForChangePasswordComplexityStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangePasswordComplexityStatusResponse), nil
	}
}

// ChangePasswordComplexityStatusInvoker 对口令复杂度检测未通过的主机进行忽略/取消忽略
func (c *HssClient) ChangePasswordComplexityStatusInvoker(request *model.ChangePasswordComplexityStatusRequest) *ChangePasswordComplexityStatusInvoker {
	requestDef := GenReqDefForChangePasswordComplexityStatus()
	return &ChangePasswordComplexityStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeVulScanPolicy 修改漏洞扫描策略
//
// 修改漏洞扫描策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ChangeVulScanPolicy(request *model.ChangeVulScanPolicyRequest) (*model.ChangeVulScanPolicyResponse, error) {
	requestDef := GenReqDefForChangeVulScanPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeVulScanPolicyResponse), nil
	}
}

// ChangeVulScanPolicyInvoker 修改漏洞扫描策略
func (c *HssClient) ChangeVulScanPolicyInvoker(request *model.ChangeVulScanPolicyRequest) *ChangeVulScanPolicyInvoker {
	requestDef := GenReqDefForChangeVulScanPolicy()
	return &ChangeVulScanPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeVulStatus 修改漏洞的状态
//
// 修改漏洞的状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ChangeVulStatus(request *model.ChangeVulStatusRequest) (*model.ChangeVulStatusResponse, error) {
	requestDef := GenReqDefForChangeVulStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeVulStatusResponse), nil
	}
}

// ChangeVulStatusInvoker 修改漏洞的状态
func (c *HssClient) ChangeVulStatusInvoker(request *model.ChangeVulStatusRequest) *ChangeVulStatusInvoker {
	requestDef := GenReqDefForChangeVulStatus()
	return &ChangeVulStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAgentDaemonset 创建集群daemonset
//
// 创建集群daemonset
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) CreateAgentDaemonset(request *model.CreateAgentDaemonsetRequest) (*model.CreateAgentDaemonsetResponse, error) {
	requestDef := GenReqDefForCreateAgentDaemonset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAgentDaemonsetResponse), nil
	}
}

// CreateAgentDaemonsetInvoker 创建集群daemonset
func (c *HssClient) CreateAgentDaemonsetInvoker(request *model.CreateAgentDaemonsetRequest) *CreateAgentDaemonsetInvoker {
	requestDef := GenReqDefForCreateAgentDaemonset()
	return &CreateAgentDaemonsetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAntiVirusPolicy 创建自定义查杀策略
//
// 创建自定义查杀策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) CreateAntiVirusPolicy(request *model.CreateAntiVirusPolicyRequest) (*model.CreateAntiVirusPolicyResponse, error) {
	requestDef := GenReqDefForCreateAntiVirusPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAntiVirusPolicyResponse), nil
	}
}

// CreateAntiVirusPolicyInvoker 创建自定义查杀策略
func (c *HssClient) CreateAntiVirusPolicyInvoker(request *model.CreateAntiVirusPolicyRequest) *CreateAntiVirusPolicyInvoker {
	requestDef := GenReqDefForCreateAntiVirusPolicy()
	return &CreateAntiVirusPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAntiVirusTask 创建病毒扫描任务
//
// 创建病毒扫描任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) CreateAntiVirusTask(request *model.CreateAntiVirusTaskRequest) (*model.CreateAntiVirusTaskResponse, error) {
	requestDef := GenReqDefForCreateAntiVirusTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAntiVirusTaskResponse), nil
	}
}

// CreateAntiVirusTaskInvoker 创建病毒扫描任务
func (c *HssClient) CreateAntiVirusTaskInvoker(request *model.CreateAntiVirusTaskRequest) *CreateAntiVirusTaskInvoker {
	requestDef := GenReqDefForCreateAntiVirusTask()
	return &CreateAntiVirusTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateClusterProtectionPolicy 新建集群防护策略
//
// 新建集群防护策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) CreateClusterProtectionPolicy(request *model.CreateClusterProtectionPolicyRequest) (*model.CreateClusterProtectionPolicyResponse, error) {
	requestDef := GenReqDefForCreateClusterProtectionPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterProtectionPolicyResponse), nil
	}
}

// CreateClusterProtectionPolicyInvoker 新建集群防护策略
func (c *HssClient) CreateClusterProtectionPolicyInvoker(request *model.CreateClusterProtectionPolicyRequest) *CreateClusterProtectionPolicyInvoker {
	requestDef := GenReqDefForCreateClusterProtectionPolicy()
	return &CreateClusterProtectionPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateClustersInfo 同步集群信息
//
// 同步集群信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) CreateClustersInfo(request *model.CreateClustersInfoRequest) (*model.CreateClustersInfoResponse, error) {
	requestDef := GenReqDefForCreateClustersInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClustersInfoResponse), nil
	}
}

// CreateClustersInfoInvoker 同步集群信息
func (c *HssClient) CreateClustersInfoInvoker(request *model.CreateClustersInfoRequest) *CreateClustersInfoInvoker {
	requestDef := GenReqDefForCreateClustersInfo()
	return &CreateClustersInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateContainerNetworkPolicy 容器集群网络添加配置策略
//
// 容器集群网络添加配置策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) CreateContainerNetworkPolicy(request *model.CreateContainerNetworkPolicyRequest) (*model.CreateContainerNetworkPolicyResponse, error) {
	requestDef := GenReqDefForCreateContainerNetworkPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateContainerNetworkPolicyResponse), nil
	}
}

// CreateContainerNetworkPolicyInvoker 容器集群网络添加配置策略
func (c *HssClient) CreateContainerNetworkPolicyInvoker(request *model.CreateContainerNetworkPolicyRequest) *CreateContainerNetworkPolicyInvoker {
	requestDef := GenReqDefForCreateContainerNetworkPolicy()
	return &CreateContainerNetworkPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDecoyPortPolicy 新增动态端口蜜罐策略
//
// 新增动态端口蜜罐策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) CreateDecoyPortPolicy(request *model.CreateDecoyPortPolicyRequest) (*model.CreateDecoyPortPolicyResponse, error) {
	requestDef := GenReqDefForCreateDecoyPortPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDecoyPortPolicyResponse), nil
	}
}

// CreateDecoyPortPolicyInvoker 新增动态端口蜜罐策略
func (c *HssClient) CreateDecoyPortPolicyInvoker(request *model.CreateDecoyPortPolicyRequest) *CreateDecoyPortPolicyInvoker {
	requestDef := GenReqDefForCreateDecoyPortPolicy()
	return &CreateDecoyPortPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGlobalAssetScanTask 创建全局资产扫描任务
//
// 创建全局资产扫描任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) CreateGlobalAssetScanTask(request *model.CreateGlobalAssetScanTaskRequest) (*model.CreateGlobalAssetScanTaskResponse, error) {
	requestDef := GenReqDefForCreateGlobalAssetScanTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGlobalAssetScanTaskResponse), nil
	}
}

// CreateGlobalAssetScanTaskInvoker 创建全局资产扫描任务
func (c *HssClient) CreateGlobalAssetScanTaskInvoker(request *model.CreateGlobalAssetScanTaskRequest) *CreateGlobalAssetScanTaskInvoker {
	requestDef := GenReqDefForCreateGlobalAssetScanTask()
	return &CreateGlobalAssetScanTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMultiCloudClusters 创建多云集群
//
// 创建多云集群
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) CreateMultiCloudClusters(request *model.CreateMultiCloudClustersRequest) (*model.CreateMultiCloudClustersResponse, error) {
	requestDef := GenReqDefForCreateMultiCloudClusters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMultiCloudClustersResponse), nil
	}
}

// CreateMultiCloudClustersInvoker 创建多云集群
func (c *HssClient) CreateMultiCloudClustersInvoker(request *model.CreateMultiCloudClustersRequest) *CreateMultiCloudClustersInvoker {
	requestDef := GenReqDefForCreateMultiCloudClusters()
	return &CreateMultiCloudClustersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateQuotasOrder HSS服务创建订单订购配额
//
// HSS服务创建订单订购配额，只支持包周期计费模式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) CreateQuotasOrder(request *model.CreateQuotasOrderRequest) (*model.CreateQuotasOrderResponse, error) {
	requestDef := GenReqDefForCreateQuotasOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateQuotasOrderResponse), nil
	}
}

// CreateQuotasOrderInvoker HSS服务创建订单订购配额
func (c *HssClient) CreateQuotasOrderInvoker(request *model.CreateQuotasOrderRequest) *CreateQuotasOrderInvoker {
	requestDef := GenReqDefForCreateQuotasOrder()
	return &CreateQuotasOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSecurityGroupPolicy 创建安全组策略
//
// 创建安全组策略(云原生网络模型)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) CreateSecurityGroupPolicy(request *model.CreateSecurityGroupPolicyRequest) (*model.CreateSecurityGroupPolicyResponse, error) {
	requestDef := GenReqDefForCreateSecurityGroupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecurityGroupPolicyResponse), nil
	}
}

// CreateSecurityGroupPolicyInvoker 创建安全组策略
func (c *HssClient) CreateSecurityGroupPolicyInvoker(request *model.CreateSecurityGroupPolicyRequest) *CreateSecurityGroupPolicyInvoker {
	requestDef := GenReqDefForCreateSecurityGroupPolicy()
	return &CreateSecurityGroupPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVulnerabilityScanTask 创建漏洞扫描任务
//
// 创建漏洞扫描任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) CreateVulnerabilityScanTask(request *model.CreateVulnerabilityScanTaskRequest) (*model.CreateVulnerabilityScanTaskResponse, error) {
	requestDef := GenReqDefForCreateVulnerabilityScanTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVulnerabilityScanTaskResponse), nil
	}
}

// CreateVulnerabilityScanTaskInvoker 创建漏洞扫描任务
func (c *HssClient) CreateVulnerabilityScanTaskInvoker(request *model.CreateVulnerabilityScanTaskRequest) *CreateVulnerabilityScanTaskInvoker {
	requestDef := GenReqDefForCreateVulnerabilityScanTask()
	return &CreateVulnerabilityScanTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAccount 删除账号
//
// 删除账号
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) DeleteAccount(request *model.DeleteAccountRequest) (*model.DeleteAccountResponse, error) {
	requestDef := GenReqDefForDeleteAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAccountResponse), nil
	}
}

// DeleteAccountInvoker 删除账号
func (c *HssClient) DeleteAccountInvoker(request *model.DeleteAccountRequest) *DeleteAccountInvoker {
	requestDef := GenReqDefForDeleteAccount()
	return &DeleteAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAgentDaemonset 删除集群daemonset
//
// 删除集群daemonset
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) DeleteAgentDaemonset(request *model.DeleteAgentDaemonsetRequest) (*model.DeleteAgentDaemonsetResponse, error) {
	requestDef := GenReqDefForDeleteAgentDaemonset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAgentDaemonsetResponse), nil
	}
}

// DeleteAgentDaemonsetInvoker 删除集群daemonset
func (c *HssClient) DeleteAgentDaemonsetInvoker(request *model.DeleteAgentDaemonsetRequest) *DeleteAgentDaemonsetInvoker {
	requestDef := GenReqDefForDeleteAgentDaemonset()
	return &DeleteAgentDaemonsetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAntivirusPolicy 删除自定义查杀策略
//
// 删除自定义查杀策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) DeleteAntivirusPolicy(request *model.DeleteAntivirusPolicyRequest) (*model.DeleteAntivirusPolicyResponse, error) {
	requestDef := GenReqDefForDeleteAntivirusPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAntivirusPolicyResponse), nil
	}
}

// DeleteAntivirusPolicyInvoker 删除自定义查杀策略
func (c *HssClient) DeleteAntivirusPolicyInvoker(request *model.DeleteAntivirusPolicyRequest) *DeleteAntivirusPolicyInvoker {
	requestDef := GenReqDefForDeleteAntivirusPolicy()
	return &DeleteAntivirusPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBaselineWhiteList 删除基线白名单
//
// 删除基线白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) DeleteBaselineWhiteList(request *model.DeleteBaselineWhiteListRequest) (*model.DeleteBaselineWhiteListResponse, error) {
	requestDef := GenReqDefForDeleteBaselineWhiteList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBaselineWhiteListResponse), nil
	}
}

// DeleteBaselineWhiteListInvoker 删除基线白名单
func (c *HssClient) DeleteBaselineWhiteListInvoker(request *model.DeleteBaselineWhiteListRequest) *DeleteBaselineWhiteListInvoker {
	requestDef := GenReqDefForDeleteBaselineWhiteList()
	return &DeleteBaselineWhiteListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteClusterProtectionPolicy 删除集群防护策略
//
// 删除集群防护策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) DeleteClusterProtectionPolicy(request *model.DeleteClusterProtectionPolicyRequest) (*model.DeleteClusterProtectionPolicyResponse, error) {
	requestDef := GenReqDefForDeleteClusterProtectionPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClusterProtectionPolicyResponse), nil
	}
}

// DeleteClusterProtectionPolicyInvoker 删除集群防护策略
func (c *HssClient) DeleteClusterProtectionPolicyInvoker(request *model.DeleteClusterProtectionPolicyRequest) *DeleteClusterProtectionPolicyInvoker {
	requestDef := GenReqDefForDeleteClusterProtectionPolicy()
	return &DeleteClusterProtectionPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteContainerNetworkPolicy 容器集群网络删除配置策略
//
// 容器集群网络删除配置策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) DeleteContainerNetworkPolicy(request *model.DeleteContainerNetworkPolicyRequest) (*model.DeleteContainerNetworkPolicyResponse, error) {
	requestDef := GenReqDefForDeleteContainerNetworkPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteContainerNetworkPolicyResponse), nil
	}
}

// DeleteContainerNetworkPolicyInvoker 容器集群网络删除配置策略
func (c *HssClient) DeleteContainerNetworkPolicyInvoker(request *model.DeleteContainerNetworkPolicyRequest) *DeleteContainerNetworkPolicyInvoker {
	requestDef := GenReqDefForDeleteContainerNetworkPolicy()
	return &DeleteContainerNetworkPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDecoyPortHostPolicy 关闭主机动态端口蜜罐策略
//
// 关闭主机动态端口蜜罐策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) DeleteDecoyPortHostPolicy(request *model.DeleteDecoyPortHostPolicyRequest) (*model.DeleteDecoyPortHostPolicyResponse, error) {
	requestDef := GenReqDefForDeleteDecoyPortHostPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDecoyPortHostPolicyResponse), nil
	}
}

// DeleteDecoyPortHostPolicyInvoker 关闭主机动态端口蜜罐策略
func (c *HssClient) DeleteDecoyPortHostPolicyInvoker(request *model.DeleteDecoyPortHostPolicyRequest) *DeleteDecoyPortHostPolicyInvoker {
	requestDef := GenReqDefForDeleteDecoyPortHostPolicy()
	return &DeleteDecoyPortHostPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDecoyPortPolicy 删除动态端口蜜罐策略
//
// 删除动态端口蜜罐策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) DeleteDecoyPortPolicy(request *model.DeleteDecoyPortPolicyRequest) (*model.DeleteDecoyPortPolicyResponse, error) {
	requestDef := GenReqDefForDeleteDecoyPortPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDecoyPortPolicyResponse), nil
	}
}

// DeleteDecoyPortPolicyInvoker 删除动态端口蜜罐策略
func (c *HssClient) DeleteDecoyPortPolicyInvoker(request *model.DeleteDecoyPortPolicyRequest) *DeleteDecoyPortPolicyInvoker {
	requestDef := GenReqDefForDeleteDecoyPortPolicy()
	return &DeleteDecoyPortPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHostsGroup 删除服务器组
//
// 删除服务器组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) DeleteHostsGroup(request *model.DeleteHostsGroupRequest) (*model.DeleteHostsGroupResponse, error) {
	requestDef := GenReqDefForDeleteHostsGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHostsGroupResponse), nil
	}
}

// DeleteHostsGroupInvoker 删除服务器组
func (c *HssClient) DeleteHostsGroupInvoker(request *model.DeleteHostsGroupRequest) *DeleteHostsGroupInvoker {
	requestDef := GenReqDefForDeleteHostsGroup()
	return &DeleteHostsGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteIsolatedFile 删除已隔离文件
//
// 删除已隔离文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) DeleteIsolatedFile(request *model.DeleteIsolatedFileRequest) (*model.DeleteIsolatedFileResponse, error) {
	requestDef := GenReqDefForDeleteIsolatedFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteIsolatedFileResponse), nil
	}
}

// DeleteIsolatedFileInvoker 删除已隔离文件
func (c *HssClient) DeleteIsolatedFileInvoker(request *model.DeleteIsolatedFileRequest) *DeleteIsolatedFileInvoker {
	requestDef := GenReqDefForDeleteIsolatedFile()
	return &DeleteIsolatedFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePolicy 删除防护策略
//
// 删除防护策略：删除策略，已经在使用的防护策略不能删除
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) DeletePolicy(request *model.DeletePolicyRequest) (*model.DeletePolicyResponse, error) {
	requestDef := GenReqDefForDeletePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePolicyResponse), nil
	}
}

// DeletePolicyInvoker 删除防护策略
func (c *HssClient) DeletePolicyInvoker(request *model.DeletePolicyRequest) *DeletePolicyInvoker {
	requestDef := GenReqDefForDeletePolicy()
	return &DeletePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteProtectionPolicy 删除防护策略
//
// 删除防护策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) DeleteProtectionPolicy(request *model.DeleteProtectionPolicyRequest) (*model.DeleteProtectionPolicyResponse, error) {
	requestDef := GenReqDefForDeleteProtectionPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteProtectionPolicyResponse), nil
	}
}

// DeleteProtectionPolicyInvoker 删除防护策略
func (c *HssClient) DeleteProtectionPolicyInvoker(request *model.DeleteProtectionPolicyRequest) *DeleteProtectionPolicyInvoker {
	requestDef := GenReqDefForDeleteProtectionPolicy()
	return &DeleteProtectionPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteResourceInstanceTag 删除资源标签
//
// 删除单个资源下的标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) DeleteResourceInstanceTag(request *model.DeleteResourceInstanceTagRequest) (*model.DeleteResourceInstanceTagResponse, error) {
	requestDef := GenReqDefForDeleteResourceInstanceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResourceInstanceTagResponse), nil
	}
}

// DeleteResourceInstanceTagInvoker 删除资源标签
func (c *HssClient) DeleteResourceInstanceTagInvoker(request *model.DeleteResourceInstanceTagRequest) *DeleteResourceInstanceTagInvoker {
	requestDef := GenReqDefForDeleteResourceInstanceTag()
	return &DeleteResourceInstanceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSecurityGroupPolicy 删除安全组策略
//
// 删除安全组策略(云原生网络模型)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) DeleteSecurityGroupPolicy(request *model.DeleteSecurityGroupPolicyRequest) (*model.DeleteSecurityGroupPolicyResponse, error) {
	requestDef := GenReqDefForDeleteSecurityGroupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecurityGroupPolicyResponse), nil
	}
}

// DeleteSecurityGroupPolicyInvoker 删除安全组策略
func (c *HssClient) DeleteSecurityGroupPolicyInvoker(request *model.DeleteSecurityGroupPolicyRequest) *DeleteSecurityGroupPolicyInvoker {
	requestDef := GenReqDefForDeleteSecurityGroupPolicy()
	return &DeleteSecurityGroupPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportAntiVirusResult 导出病毒扫描结果列表
//
// 导出病毒扫描结果列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ExportAntiVirusResult(request *model.ExportAntiVirusResultRequest) (*model.ExportAntiVirusResultResponse, error) {
	requestDef := GenReqDefForExportAntiVirusResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportAntiVirusResultResponse), nil
	}
}

// ExportAntiVirusResultInvoker 导出病毒扫描结果列表
func (c *HssClient) ExportAntiVirusResultInvoker(request *model.ExportAntiVirusResultRequest) *ExportAntiVirusResultInvoker {
	requestDef := GenReqDefForExportAntiVirusResult()
	return &ExportAntiVirusResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportContainerList 创建容器导出任务
//
// 创建容器导出任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ExportContainerList(request *model.ExportContainerListRequest) (*model.ExportContainerListResponse, error) {
	requestDef := GenReqDefForExportContainerList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportContainerListResponse), nil
	}
}

// ExportContainerListInvoker 创建容器导出任务
func (c *HssClient) ExportContainerListInvoker(request *model.ExportContainerListRequest) *ExportContainerListInvoker {
	requestDef := GenReqDefForExportContainerList()
	return &ExportContainerListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportVuls 导出漏洞及漏洞影响的主机的相关信息
//
// 导出漏洞及漏洞影响的主机的相关信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ExportVuls(request *model.ExportVulsRequest) (*model.ExportVulsResponse, error) {
	requestDef := GenReqDefForExportVuls()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportVulsResponse), nil
	}
}

// ExportVulsInvoker 导出漏洞及漏洞影响的主机的相关信息
func (c *HssClient) ExportVulsInvoker(request *model.ExportVulsRequest) *ExportVulsInvoker {
	requestDef := GenReqDefForExportVuls()
	return &ExportVulsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// HandleAntiVirusResult 处置病毒扫描结果
//
// 处置病毒扫描结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) HandleAntiVirusResult(request *model.HandleAntiVirusResultRequest) (*model.HandleAntiVirusResultResponse, error) {
	requestDef := GenReqDefForHandleAntiVirusResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.HandleAntiVirusResultResponse), nil
	}
}

// HandleAntiVirusResultInvoker 处置病毒扫描结果
func (c *HssClient) HandleAntiVirusResultInvoker(request *model.HandleAntiVirusResultRequest) *HandleAntiVirusResultInvoker {
	requestDef := GenReqDefForHandleAntiVirusResult()
	return &HandleAntiVirusResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccounts 查询多账号列表
//
// 查询多账号列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListAccounts(request *model.ListAccountsRequest) (*model.ListAccountsResponse, error) {
	requestDef := GenReqDefForListAccounts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAccountsResponse), nil
	}
}

// ListAccountsInvoker 查询多账号列表
func (c *HssClient) ListAccountsInvoker(request *model.ListAccountsRequest) *ListAccountsInvoker {
	requestDef := GenReqDefForListAccounts()
	return &ListAccountsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAgentInstallScript 查询agent安装脚本
//
// 查询agent安装脚本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListAgentInstallScript(request *model.ListAgentInstallScriptRequest) (*model.ListAgentInstallScriptResponse, error) {
	requestDef := GenReqDefForListAgentInstallScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAgentInstallScriptResponse), nil
	}
}

// ListAgentInstallScriptInvoker 查询agent安装脚本
func (c *HssClient) ListAgentInstallScriptInvoker(request *model.ListAgentInstallScriptRequest) *ListAgentInstallScriptInvoker {
	requestDef := GenReqDefForListAgentInstallScript()
	return &ListAgentInstallScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlarmWhiteList 查询告警白名单列表
//
// 查询告警白名单列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListAlarmWhiteList(request *model.ListAlarmWhiteListRequest) (*model.ListAlarmWhiteListResponse, error) {
	requestDef := GenReqDefForListAlarmWhiteList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmWhiteListResponse), nil
	}
}

// ListAlarmWhiteListInvoker 查询告警白名单列表
func (c *HssClient) ListAlarmWhiteListInvoker(request *model.ListAlarmWhiteListRequest) *ListAlarmWhiteListInvoker {
	requestDef := GenReqDefForListAlarmWhiteList()
	return &ListAlarmWhiteListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAntiVirusHost 查询病毒查杀可选服务器列表
//
// 查询病毒查杀可选服务器列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListAntiVirusHost(request *model.ListAntiVirusHostRequest) (*model.ListAntiVirusHostResponse, error) {
	requestDef := GenReqDefForListAntiVirusHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAntiVirusHostResponse), nil
	}
}

// ListAntiVirusHostInvoker 查询病毒查杀可选服务器列表
func (c *HssClient) ListAntiVirusHostInvoker(request *model.ListAntiVirusHostRequest) *ListAntiVirusHostInvoker {
	requestDef := GenReqDefForListAntiVirusHost()
	return &ListAntiVirusHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAntiVirusPolicy 查询自定义查杀策略列表
//
// 查询自定义查杀策略列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListAntiVirusPolicy(request *model.ListAntiVirusPolicyRequest) (*model.ListAntiVirusPolicyResponse, error) {
	requestDef := GenReqDefForListAntiVirusPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAntiVirusPolicyResponse), nil
	}
}

// ListAntiVirusPolicyInvoker 查询自定义查杀策略列表
func (c *HssClient) ListAntiVirusPolicyInvoker(request *model.ListAntiVirusPolicyRequest) *ListAntiVirusPolicyInvoker {
	requestDef := GenReqDefForListAntiVirusPolicy()
	return &ListAntiVirusPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAntiVirusResult 查询病毒扫描结果列表
//
// 查询病毒扫描结果列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListAntiVirusResult(request *model.ListAntiVirusResultRequest) (*model.ListAntiVirusResultResponse, error) {
	requestDef := GenReqDefForListAntiVirusResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAntiVirusResultResponse), nil
	}
}

// ListAntiVirusResultInvoker 查询病毒扫描结果列表
func (c *HssClient) ListAntiVirusResultInvoker(request *model.ListAntiVirusResultRequest) *ListAntiVirusResultInvoker {
	requestDef := GenReqDefForListAntiVirusResult()
	return &ListAntiVirusResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAntiVirusTask 查看病毒扫描任务列表
//
// 查看病毒扫描任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListAntiVirusTask(request *model.ListAntiVirusTaskRequest) (*model.ListAntiVirusTaskResponse, error) {
	requestDef := GenReqDefForListAntiVirusTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAntiVirusTaskResponse), nil
	}
}

// ListAntiVirusTaskInvoker 查看病毒扫描任务列表
func (c *HssClient) ListAntiVirusTaskInvoker(request *model.ListAntiVirusTaskRequest) *ListAntiVirusTaskInvoker {
	requestDef := GenReqDefForListAntiVirusTask()
	return &ListAntiVirusTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAntivirusHandleHistory 查询病毒扫描历史处置记录列表
//
// 查询病毒扫描历史处置记录列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListAntivirusHandleHistory(request *model.ListAntivirusHandleHistoryRequest) (*model.ListAntivirusHandleHistoryResponse, error) {
	requestDef := GenReqDefForListAntivirusHandleHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAntivirusHandleHistoryResponse), nil
	}
}

// ListAntivirusHandleHistoryInvoker 查询病毒扫描历史处置记录列表
func (c *HssClient) ListAntivirusHandleHistoryInvoker(request *model.ListAntivirusHandleHistoryRequest) *ListAntivirusHandleHistoryInvoker {
	requestDef := GenReqDefForListAntivirusHandleHistory()
	return &ListAntivirusHandleHistoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppChangeHistories 获取软件信息的历史变动记录
//
// 获取软件信息的历史变动记录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListAppChangeHistories(request *model.ListAppChangeHistoriesRequest) (*model.ListAppChangeHistoriesResponse, error) {
	requestDef := GenReqDefForListAppChangeHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppChangeHistoriesResponse), nil
	}
}

// ListAppChangeHistoriesInvoker 获取软件信息的历史变动记录
func (c *HssClient) ListAppChangeHistoriesInvoker(request *model.ListAppChangeHistoriesRequest) *ListAppChangeHistoriesInvoker {
	requestDef := GenReqDefForListAppChangeHistories()
	return &ListAppChangeHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppStatistics 查询软件列表
//
// 查询软件列表，支持通过软件名称查询对应的服务器数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListAppStatistics(request *model.ListAppStatisticsRequest) (*model.ListAppStatisticsResponse, error) {
	requestDef := GenReqDefForListAppStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppStatisticsResponse), nil
	}
}

// ListAppStatisticsInvoker 查询软件列表
func (c *HssClient) ListAppStatisticsInvoker(request *model.ListAppStatisticsRequest) *ListAppStatisticsInvoker {
	requestDef := GenReqDefForListAppStatistics()
	return &ListAppStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApps 查询软件的服务器列表
//
// 查询软件的服务器列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListApps(request *model.ListAppsRequest) (*model.ListAppsResponse, error) {
	requestDef := GenReqDefForListApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppsResponse), nil
	}
}

// ListAppsInvoker 查询软件的服务器列表
func (c *HssClient) ListAppsInvoker(request *model.ListAppsRequest) *ListAppsInvoker {
	requestDef := GenReqDefForListApps()
	return &ListAppsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAutoLaunchChangeHistories 获取自启动项的历史变动记录
//
// 获取自启动项的历史变动记录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListAutoLaunchChangeHistories(request *model.ListAutoLaunchChangeHistoriesRequest) (*model.ListAutoLaunchChangeHistoriesResponse, error) {
	requestDef := GenReqDefForListAutoLaunchChangeHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAutoLaunchChangeHistoriesResponse), nil
	}
}

// ListAutoLaunchChangeHistoriesInvoker 获取自启动项的历史变动记录
func (c *HssClient) ListAutoLaunchChangeHistoriesInvoker(request *model.ListAutoLaunchChangeHistoriesRequest) *ListAutoLaunchChangeHistoriesInvoker {
	requestDef := GenReqDefForListAutoLaunchChangeHistories()
	return &ListAutoLaunchChangeHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAutoLaunchStatistics 查询自启动项信息
//
// 查询自启动信息，支持通过传入自启动名称查询启动类型和服务器数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListAutoLaunchStatistics(request *model.ListAutoLaunchStatisticsRequest) (*model.ListAutoLaunchStatisticsResponse, error) {
	requestDef := GenReqDefForListAutoLaunchStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAutoLaunchStatisticsResponse), nil
	}
}

// ListAutoLaunchStatisticsInvoker 查询自启动项信息
func (c *HssClient) ListAutoLaunchStatisticsInvoker(request *model.ListAutoLaunchStatisticsRequest) *ListAutoLaunchStatisticsInvoker {
	requestDef := GenReqDefForListAutoLaunchStatistics()
	return &ListAutoLaunchStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAutoLaunchs 查询自启动项的服务列表
//
// 查询自启动项的服务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListAutoLaunchs(request *model.ListAutoLaunchsRequest) (*model.ListAutoLaunchsResponse, error) {
	requestDef := GenReqDefForListAutoLaunchs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAutoLaunchsResponse), nil
	}
}

// ListAutoLaunchsInvoker 查询自启动项的服务列表
func (c *HssClient) ListAutoLaunchsInvoker(request *model.ListAutoLaunchsRequest) *ListAutoLaunchsInvoker {
	requestDef := GenReqDefForListAutoLaunchs()
	return &ListAutoLaunchsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBackupVaults 查询备份存储库列表
//
// 查询备份存储库列表，若进行绑定主机，则需要额外判断，同时满足以下条件：1.存储库状态为“可用”状态；2.备份策略状态为“已启用”；3.存储库有剩余可用备份容量；4.存储库绑定的服务器数量少于256。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListBackupVaults(request *model.ListBackupVaultsRequest) (*model.ListBackupVaultsResponse, error) {
	requestDef := GenReqDefForListBackupVaults()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackupVaultsResponse), nil
	}
}

// ListBackupVaultsInvoker 查询备份存储库列表
func (c *HssClient) ListBackupVaultsInvoker(request *model.ListBackupVaultsRequest) *ListBackupVaultsInvoker {
	requestDef := GenReqDefForListBackupVaults()
	return &ListBackupVaultsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBaselineWhiteLists 查询基线白名单列表
//
// 查询基线白名单列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListBaselineWhiteLists(request *model.ListBaselineWhiteListsRequest) (*model.ListBaselineWhiteListsResponse, error) {
	requestDef := GenReqDefForListBaselineWhiteLists()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBaselineWhiteListsResponse), nil
	}
}

// ListBaselineWhiteListsInvoker 查询基线白名单列表
func (c *HssClient) ListBaselineWhiteListsInvoker(request *model.ListBaselineWhiteListsRequest) *ListBaselineWhiteListsInvoker {
	requestDef := GenReqDefForListBaselineWhiteLists()
	return &ListBaselineWhiteListsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBlockedIp 查询已拦截IP列表
//
// 查询已拦截IP列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListBlockedIp(request *model.ListBlockedIpRequest) (*model.ListBlockedIpResponse, error) {
	requestDef := GenReqDefForListBlockedIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBlockedIpResponse), nil
	}
}

// ListBlockedIpInvoker 查询已拦截IP列表
func (c *HssClient) ListBlockedIpInvoker(request *model.ListBlockedIpRequest) *ListBlockedIpInvoker {
	requestDef := GenReqDefForListBlockedIp()
	return &ListBlockedIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCceClusterConfig 获取集群配置
//
// 获取集群配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListCceClusterConfig(request *model.ListCceClusterConfigRequest) (*model.ListCceClusterConfigResponse, error) {
	requestDef := GenReqDefForListCceClusterConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCceClusterConfigResponse), nil
	}
}

// ListCceClusterConfigInvoker 获取集群配置
func (c *HssClient) ListCceClusterConfigInvoker(request *model.ListCceClusterConfigRequest) *ListCceClusterConfigInvoker {
	requestDef := GenReqDefForListCceClusterConfig()
	return &ListCceClusterConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCceClusterDetectRisk 批量获取容器集群风险信息
//
// 批量获取容器集群风险信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListCceClusterDetectRisk(request *model.ListCceClusterDetectRiskRequest) (*model.ListCceClusterDetectRiskResponse, error) {
	requestDef := GenReqDefForListCceClusterDetectRisk()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCceClusterDetectRiskResponse), nil
	}
}

// ListCceClusterDetectRiskInvoker 批量获取容器集群风险信息
func (c *HssClient) ListCceClusterDetectRiskInvoker(request *model.ListCceClusterDetectRiskRequest) *ListCceClusterDetectRiskInvoker {
	requestDef := GenReqDefForListCceClusterDetectRisk()
	return &ListCceClusterDetectRiskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCheckFeatureRule 查询检测规则列表
//
// 查询检测规则列表：查询默认检测规则信息，包含14种检测规则，默认都不开启
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListCheckFeatureRule(request *model.ListCheckFeatureRuleRequest) (*model.ListCheckFeatureRuleResponse, error) {
	requestDef := GenReqDefForListCheckFeatureRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCheckFeatureRuleResponse), nil
	}
}

// ListCheckFeatureRuleInvoker 查询检测规则列表
func (c *HssClient) ListCheckFeatureRuleInvoker(request *model.ListCheckFeatureRuleRequest) *ListCheckFeatureRuleInvoker {
	requestDef := GenReqDefForListCheckFeatureRule()
	return &ListCheckFeatureRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusterAuditLogs 查询k8s集群审计日志列表
//
// 查询k8s集群审计日志列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListClusterAuditLogs(request *model.ListClusterAuditLogsRequest) (*model.ListClusterAuditLogsResponse, error) {
	requestDef := GenReqDefForListClusterAuditLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterAuditLogsResponse), nil
	}
}

// ListClusterAuditLogsInvoker 查询k8s集群审计日志列表
func (c *HssClient) ListClusterAuditLogsInvoker(request *model.ListClusterAuditLogsRequest) *ListClusterAuditLogsInvoker {
	requestDef := GenReqDefForListClusterAuditLogs()
	return &ListClusterAuditLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusterEventLogs 查询k8s集群事件列表
//
// 查询k8s集群事件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListClusterEventLogs(request *model.ListClusterEventLogsRequest) (*model.ListClusterEventLogsResponse, error) {
	requestDef := GenReqDefForListClusterEventLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterEventLogsResponse), nil
	}
}

// ListClusterEventLogsInvoker 查询k8s集群事件列表
func (c *HssClient) ListClusterEventLogsInvoker(request *model.ListClusterEventLogsRequest) *ListClusterEventLogsInvoker {
	requestDef := GenReqDefForListClusterEventLogs()
	return &ListClusterEventLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusterEvents 获取所有集群中告警事件
//
// 获取所有集群中告警事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListClusterEvents(request *model.ListClusterEventsRequest) (*model.ListClusterEventsResponse, error) {
	requestDef := GenReqDefForListClusterEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterEventsResponse), nil
	}
}

// ListClusterEventsInvoker 获取所有集群中告警事件
func (c *HssClient) ListClusterEventsInvoker(request *model.ListClusterEventsRequest) *ListClusterEventsInvoker {
	requestDef := GenReqDefForListClusterEvents()
	return &ListClusterEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusterProtectOverview 集群防护概览
//
// 集群防护概览
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListClusterProtectOverview(request *model.ListClusterProtectOverviewRequest) (*model.ListClusterProtectOverviewResponse, error) {
	requestDef := GenReqDefForListClusterProtectOverview()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterProtectOverviewResponse), nil
	}
}

// ListClusterProtectOverviewInvoker 集群防护概览
func (c *HssClient) ListClusterProtectOverviewInvoker(request *model.ListClusterProtectOverviewRequest) *ListClusterProtectOverviewInvoker {
	requestDef := GenReqDefForListClusterProtectOverview()
	return &ListClusterProtectOverviewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusterProtectPolicyTemplates 查询集群组件防护策略模板列表
//
// 查询集群防护策略模板列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListClusterProtectPolicyTemplates(request *model.ListClusterProtectPolicyTemplatesRequest) (*model.ListClusterProtectPolicyTemplatesResponse, error) {
	requestDef := GenReqDefForListClusterProtectPolicyTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterProtectPolicyTemplatesResponse), nil
	}
}

// ListClusterProtectPolicyTemplatesInvoker 查询集群组件防护策略模板列表
func (c *HssClient) ListClusterProtectPolicyTemplatesInvoker(request *model.ListClusterProtectPolicyTemplatesRequest) *ListClusterProtectPolicyTemplatesInvoker {
	requestDef := GenReqDefForListClusterProtectPolicyTemplates()
	return &ListClusterProtectPolicyTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusterProtectionDefaultPolicy 获取集群防护默认策略列表
//
// 获取集群防护默认策略列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListClusterProtectionDefaultPolicy(request *model.ListClusterProtectionDefaultPolicyRequest) (*model.ListClusterProtectionDefaultPolicyResponse, error) {
	requestDef := GenReqDefForListClusterProtectionDefaultPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterProtectionDefaultPolicyResponse), nil
	}
}

// ListClusterProtectionDefaultPolicyInvoker 获取集群防护默认策略列表
func (c *HssClient) ListClusterProtectionDefaultPolicyInvoker(request *model.ListClusterProtectionDefaultPolicyRequest) *ListClusterProtectionDefaultPolicyInvoker {
	requestDef := GenReqDefForListClusterProtectionDefaultPolicy()
	return &ListClusterProtectionDefaultPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusterProtectionInfo 查询集群防护信息
//
// 查询集群防护信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListClusterProtectionInfo(request *model.ListClusterProtectionInfoRequest) (*model.ListClusterProtectionInfoResponse, error) {
	requestDef := GenReqDefForListClusterProtectionInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterProtectionInfoResponse), nil
	}
}

// ListClusterProtectionInfoInvoker 查询集群防护信息
func (c *HssClient) ListClusterProtectionInfoInvoker(request *model.ListClusterProtectionInfoRequest) *ListClusterProtectionInfoInvoker {
	requestDef := GenReqDefForListClusterProtectionInfo()
	return &ListClusterProtectionInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusterProtectionItem 获取集群所有防护项
//
// 获取集群所有防护项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListClusterProtectionItem(request *model.ListClusterProtectionItemRequest) (*model.ListClusterProtectionItemResponse, error) {
	requestDef := GenReqDefForListClusterProtectionItem()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterProtectionItemResponse), nil
	}
}

// ListClusterProtectionItemInvoker 获取集群所有防护项
func (c *HssClient) ListClusterProtectionItemInvoker(request *model.ListClusterProtectionItemRequest) *ListClusterProtectionItemInvoker {
	requestDef := GenReqDefForListClusterProtectionItem()
	return &ListClusterProtectionItemInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusterProtectionPolicy 获取集群防护策略列表
//
// 获取集群防护策略列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListClusterProtectionPolicy(request *model.ListClusterProtectionPolicyRequest) (*model.ListClusterProtectionPolicyResponse, error) {
	requestDef := GenReqDefForListClusterProtectionPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterProtectionPolicyResponse), nil
	}
}

// ListClusterProtectionPolicyInvoker 获取集群防护策略列表
func (c *HssClient) ListClusterProtectionPolicyInvoker(request *model.ListClusterProtectionPolicyRequest) *ListClusterProtectionPolicyInvoker {
	requestDef := GenReqDefForListClusterProtectionPolicy()
	return &ListClusterProtectionPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusterProtectionPolicyDetail 查看指定策略的详情
//
// 查看指定策略的详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListClusterProtectionPolicyDetail(request *model.ListClusterProtectionPolicyDetailRequest) (*model.ListClusterProtectionPolicyDetailResponse, error) {
	requestDef := GenReqDefForListClusterProtectionPolicyDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterProtectionPolicyDetailResponse), nil
	}
}

// ListClusterProtectionPolicyDetailInvoker 查看指定策略的详情
func (c *HssClient) ListClusterProtectionPolicyDetailInvoker(request *model.ListClusterProtectionPolicyDetailRequest) *ListClusterProtectionPolicyDetailInvoker {
	requestDef := GenReqDefForListClusterProtectionPolicyDetail()
	return &ListClusterProtectionPolicyDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCommonTips 获取部分提示信息
//
// 获取部分提示信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListCommonTips(request *model.ListCommonTipsRequest) (*model.ListCommonTipsResponse, error) {
	requestDef := GenReqDefForListCommonTips()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCommonTipsResponse), nil
	}
}

// ListCommonTipsInvoker 获取部分提示信息
func (c *HssClient) ListCommonTipsInvoker(request *model.ListCommonTipsRequest) *ListCommonTipsInvoker {
	requestDef := GenReqDefForListCommonTips()
	return &ListCommonTipsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListContainerCmdLogs 查询容器内运行的命令列表
//
// 查询容器内运行的命令列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListContainerCmdLogs(request *model.ListContainerCmdLogsRequest) (*model.ListContainerCmdLogsResponse, error) {
	requestDef := GenReqDefForListContainerCmdLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListContainerCmdLogsResponse), nil
	}
}

// ListContainerCmdLogsInvoker 查询容器内运行的命令列表
func (c *HssClient) ListContainerCmdLogsInvoker(request *model.ListContainerCmdLogsRequest) *ListContainerCmdLogsInvoker {
	requestDef := GenReqDefForListContainerCmdLogs()
	return &ListContainerCmdLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListContainerImageLogs 查询容器镜像操作日志
//
// 查询容器镜像操作日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListContainerImageLogs(request *model.ListContainerImageLogsRequest) (*model.ListContainerImageLogsResponse, error) {
	requestDef := GenReqDefForListContainerImageLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListContainerImageLogsResponse), nil
	}
}

// ListContainerImageLogsInvoker 查询容器镜像操作日志
func (c *HssClient) ListContainerImageLogsInvoker(request *model.ListContainerImageLogsRequest) *ListContainerImageLogsInvoker {
	requestDef := GenReqDefForListContainerImageLogs()
	return &ListContainerImageLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListContainerImages 查询容器镜像列表
//
// 查询容器镜像列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListContainerImages(request *model.ListContainerImagesRequest) (*model.ListContainerImagesResponse, error) {
	requestDef := GenReqDefForListContainerImages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListContainerImagesResponse), nil
	}
}

// ListContainerImagesInvoker 查询容器镜像列表
func (c *HssClient) ListContainerImagesInvoker(request *model.ListContainerImagesRequest) *ListContainerImagesInvoker {
	requestDef := GenReqDefForListContainerImages()
	return &ListContainerImagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListContainerLogs 查询容器日志列表
//
// 查询容器日志列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListContainerLogs(request *model.ListContainerLogsRequest) (*model.ListContainerLogsResponse, error) {
	requestDef := GenReqDefForListContainerLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListContainerLogsResponse), nil
	}
}

// ListContainerLogsInvoker 查询容器日志列表
func (c *HssClient) ListContainerLogsInvoker(request *model.ListContainerLogsRequest) *ListContainerLogsInvoker {
	requestDef := GenReqDefForListContainerLogs()
	return &ListContainerLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListContainerNetworkClusters 查询容器防护的集群信息
//
// 查询容器防护的集群信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListContainerNetworkClusters(request *model.ListContainerNetworkClustersRequest) (*model.ListContainerNetworkClustersResponse, error) {
	requestDef := GenReqDefForListContainerNetworkClusters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListContainerNetworkClustersResponse), nil
	}
}

// ListContainerNetworkClustersInvoker 查询容器防护的集群信息
func (c *HssClient) ListContainerNetworkClustersInvoker(request *model.ListContainerNetworkClustersRequest) *ListContainerNetworkClustersInvoker {
	requestDef := GenReqDefForListContainerNetworkClusters()
	return &ListContainerNetworkClustersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListContainerNetworkNodeList 查询容器集群VPC网络的节点列表
//
// 查询容器集群网络的策略列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListContainerNetworkNodeList(request *model.ListContainerNetworkNodeListRequest) (*model.ListContainerNetworkNodeListResponse, error) {
	requestDef := GenReqDefForListContainerNetworkNodeList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListContainerNetworkNodeListResponse), nil
	}
}

// ListContainerNetworkNodeListInvoker 查询容器集群VPC网络的节点列表
func (c *HssClient) ListContainerNetworkNodeListInvoker(request *model.ListContainerNetworkNodeListRequest) *ListContainerNetworkNodeListInvoker {
	requestDef := GenReqDefForListContainerNetworkNodeList()
	return &ListContainerNetworkNodeListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListContainerNetworkPolicy 查询容器集群网络的策略列表
//
// 查询容器集群网络的策略列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListContainerNetworkPolicy(request *model.ListContainerNetworkPolicyRequest) (*model.ListContainerNetworkPolicyResponse, error) {
	requestDef := GenReqDefForListContainerNetworkPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListContainerNetworkPolicyResponse), nil
	}
}

// ListContainerNetworkPolicyInvoker 查询容器集群网络的策略列表
func (c *HssClient) ListContainerNetworkPolicyInvoker(request *model.ListContainerNetworkPolicyRequest) *ListContainerNetworkPolicyInvoker {
	requestDef := GenReqDefForListContainerNetworkPolicy()
	return &ListContainerNetworkPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListContainerNodes 查询容器节点列表
//
// 查询容器节点列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListContainerNodes(request *model.ListContainerNodesRequest) (*model.ListContainerNodesResponse, error) {
	requestDef := GenReqDefForListContainerNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListContainerNodesResponse), nil
	}
}

// ListContainerNodesInvoker 查询容器节点列表
func (c *HssClient) ListContainerNodesInvoker(request *model.ListContainerNodesRequest) *ListContainerNodesInvoker {
	requestDef := GenReqDefForListContainerNodes()
	return &ListContainerNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListContainers 查询容器基本信息列表
//
// 查询容器基本信息列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListContainers(request *model.ListContainersRequest) (*model.ListContainersResponse, error) {
	requestDef := GenReqDefForListContainers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListContainersResponse), nil
	}
}

// ListContainersInvoker 查询容器基本信息列表
func (c *HssClient) ListContainersInvoker(request *model.ListContainersRequest) *ListContainersInvoker {
	requestDef := GenReqDefForListContainers()
	return &ListContainersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDecoyPortPolicy 查看动态端口蜜罐策略
//
// 查看动态端口蜜罐策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListDecoyPortPolicy(request *model.ListDecoyPortPolicyRequest) (*model.ListDecoyPortPolicyResponse, error) {
	requestDef := GenReqDefForListDecoyPortPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDecoyPortPolicyResponse), nil
	}
}

// ListDecoyPortPolicyInvoker 查看动态端口蜜罐策略
func (c *HssClient) ListDecoyPortPolicyInvoker(request *model.ListDecoyPortPolicyRequest) *ListDecoyPortPolicyInvoker {
	requestDef := GenReqDefForListDecoyPortPolicy()
	return &ListDecoyPortPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDownloadExportedFile 下载导出文件
//
// 下载导出文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListDownloadExportedFile(request *model.ListDownloadExportedFileRequest) (*model.ListDownloadExportedFileResponse, error) {
	requestDef := GenReqDefForListDownloadExportedFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDownloadExportedFileResponse), nil
	}
}

// ListDownloadExportedFileInvoker 下载导出文件
func (c *HssClient) ListDownloadExportedFileInvoker(request *model.ListDownloadExportedFileRequest) *ListDownloadExportedFileInvoker {
	requestDef := GenReqDefForListDownloadExportedFile()
	return &ListDownloadExportedFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEventHandleHistory 查询告警事件历史处置记录列表
//
// 查询告警事件历史处置记录列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListEventHandleHistory(request *model.ListEventHandleHistoryRequest) (*model.ListEventHandleHistoryResponse, error) {
	requestDef := GenReqDefForListEventHandleHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEventHandleHistoryResponse), nil
	}
}

// ListEventHandleHistoryInvoker 查询告警事件历史处置记录列表
func (c *HssClient) ListEventHandleHistoryInvoker(request *model.ListEventHandleHistoryRequest) *ListEventHandleHistoryInvoker {
	requestDef := GenReqDefForListEventHandleHistory()
	return &ListEventHandleHistoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGlobalAssetScanTask 查询资产全局扫描任务状态
//
// 查询资产全局扫描任务状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListGlobalAssetScanTask(request *model.ListGlobalAssetScanTaskRequest) (*model.ListGlobalAssetScanTaskResponse, error) {
	requestDef := GenReqDefForListGlobalAssetScanTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGlobalAssetScanTaskResponse), nil
	}
}

// ListGlobalAssetScanTaskInvoker 查询资产全局扫描任务状态
func (c *HssClient) ListGlobalAssetScanTaskInvoker(request *model.ListGlobalAssetScanTaskRequest) *ListGlobalAssetScanTaskInvoker {
	requestDef := GenReqDefForListGlobalAssetScanTask()
	return &ListGlobalAssetScanTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHostGroups 查询服务器组列表
//
// 查询服务器组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListHostGroups(request *model.ListHostGroupsRequest) (*model.ListHostGroupsResponse, error) {
	requestDef := GenReqDefForListHostGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostGroupsResponse), nil
	}
}

// ListHostGroupsInvoker 查询服务器组列表
func (c *HssClient) ListHostGroupsInvoker(request *model.ListHostGroupsRequest) *ListHostGroupsInvoker {
	requestDef := GenReqDefForListHostGroups()
	return &ListHostGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHostProtectHistoryInfo 查询主机静态网页防篡改防护动态
//
// 查询主机静态网页防篡改防护动态：展示服务器名称、服务器ip、防护策略、检测时间、防护文件、事件描述信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListHostProtectHistoryInfo(request *model.ListHostProtectHistoryInfoRequest) (*model.ListHostProtectHistoryInfoResponse, error) {
	requestDef := GenReqDefForListHostProtectHistoryInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostProtectHistoryInfoResponse), nil
	}
}

// ListHostProtectHistoryInfoInvoker 查询主机静态网页防篡改防护动态
func (c *HssClient) ListHostProtectHistoryInfoInvoker(request *model.ListHostProtectHistoryInfoRequest) *ListHostProtectHistoryInfoInvoker {
	requestDef := GenReqDefForListHostProtectHistoryInfo()
	return &ListHostProtectHistoryInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHostRaspProtectHistoryInfo 查询主机动态网页防篡改防护动态
//
// 查询主机动态网页防篡改防护动态：包含告警级别、服务器ip、服务器名称、威胁类型、告警时间、攻击源ip、攻击源url信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListHostRaspProtectHistoryInfo(request *model.ListHostRaspProtectHistoryInfoRequest) (*model.ListHostRaspProtectHistoryInfoResponse, error) {
	requestDef := GenReqDefForListHostRaspProtectHistoryInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostRaspProtectHistoryInfoResponse), nil
	}
}

// ListHostRaspProtectHistoryInfoInvoker 查询主机动态网页防篡改防护动态
func (c *HssClient) ListHostRaspProtectHistoryInfoInvoker(request *model.ListHostRaspProtectHistoryInfoRequest) *ListHostRaspProtectHistoryInfoInvoker {
	requestDef := GenReqDefForListHostRaspProtectHistoryInfo()
	return &ListHostRaspProtectHistoryInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHostStatus 查询云服务器列表
//
// 查询云服务器列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListHostStatus(request *model.ListHostStatusRequest) (*model.ListHostStatusResponse, error) {
	requestDef := GenReqDefForListHostStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostStatusResponse), nil
	}
}

// ListHostStatusInvoker 查询云服务器列表
func (c *HssClient) ListHostStatusInvoker(request *model.ListHostStatusRequest) *ListHostStatusInvoker {
	requestDef := GenReqDefForListHostStatus()
	return &ListHostStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHostVuls 查询单台服务器漏洞信息
//
// 查询单台服务器漏洞信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListHostVuls(request *model.ListHostVulsRequest) (*model.ListHostVulsResponse, error) {
	requestDef := GenReqDefForListHostVuls()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostVulsResponse), nil
	}
}

// ListHostVulsInvoker 查询单台服务器漏洞信息
func (c *HssClient) ListHostVulsInvoker(request *model.ListHostVulsRequest) *ListHostVulsInvoker {
	requestDef := GenReqDefForListHostVuls()
	return &ListHostVulsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListImageLocal 本地镜像列表查询
//
// 本地镜像列表查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListImageLocal(request *model.ListImageLocalRequest) (*model.ListImageLocalResponse, error) {
	requestDef := GenReqDefForListImageLocal()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImageLocalResponse), nil
	}
}

// ListImageLocalInvoker 本地镜像列表查询
func (c *HssClient) ListImageLocalInvoker(request *model.ListImageLocalRequest) *ListImageLocalInvoker {
	requestDef := GenReqDefForListImageLocal()
	return &ListImageLocalInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListImageRiskConfigRules 查询镜像指定安全配置项的检查项列表
//
// 查询镜像指定安全配置项的检查项列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListImageRiskConfigRules(request *model.ListImageRiskConfigRulesRequest) (*model.ListImageRiskConfigRulesResponse, error) {
	requestDef := GenReqDefForListImageRiskConfigRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImageRiskConfigRulesResponse), nil
	}
}

// ListImageRiskConfigRulesInvoker 查询镜像指定安全配置项的检查项列表
func (c *HssClient) ListImageRiskConfigRulesInvoker(request *model.ListImageRiskConfigRulesRequest) *ListImageRiskConfigRulesInvoker {
	requestDef := GenReqDefForListImageRiskConfigRules()
	return &ListImageRiskConfigRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListImageRiskConfigs 查询镜像安全配置检测结果列表
//
// 查询镜像安全配置检测结果列表,当前支持检测CentOS 7、Debian 10、EulerOS和Ubuntu16镜像的系统配置项、SSH应用配置项。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListImageRiskConfigs(request *model.ListImageRiskConfigsRequest) (*model.ListImageRiskConfigsResponse, error) {
	requestDef := GenReqDefForListImageRiskConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImageRiskConfigsResponse), nil
	}
}

// ListImageRiskConfigsInvoker 查询镜像安全配置检测结果列表
func (c *HssClient) ListImageRiskConfigsInvoker(request *model.ListImageRiskConfigsRequest) *ListImageRiskConfigsInvoker {
	requestDef := GenReqDefForListImageRiskConfigs()
	return &ListImageRiskConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListImageVulnerabilities 查询镜像的漏洞信息
//
// 查询镜像的漏洞信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListImageVulnerabilities(request *model.ListImageVulnerabilitiesRequest) (*model.ListImageVulnerabilitiesResponse, error) {
	requestDef := GenReqDefForListImageVulnerabilities()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImageVulnerabilitiesResponse), nil
	}
}

// ListImageVulnerabilitiesInvoker 查询镜像的漏洞信息
func (c *HssClient) ListImageVulnerabilitiesInvoker(request *model.ListImageVulnerabilitiesRequest) *ListImageVulnerabilitiesInvoker {
	requestDef := GenReqDefForListImageVulnerabilities()
	return &ListImageVulnerabilitiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIsolatedFile 查询已隔离文件列表
//
// 查询已隔离文件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListIsolatedFile(request *model.ListIsolatedFileRequest) (*model.ListIsolatedFileResponse, error) {
	requestDef := GenReqDefForListIsolatedFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIsolatedFileResponse), nil
	}
}

// ListIsolatedFileInvoker 查询已隔离文件列表
func (c *HssClient) ListIsolatedFileInvoker(request *model.ListIsolatedFileRequest) *ListIsolatedFileInvoker {
	requestDef := GenReqDefForListIsolatedFile()
	return &ListIsolatedFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListJarPackageHostInfo 查询指定中间件的服务器列表
//
// 查询指定中间件的服务器列表，通过传入中间件名称参数，返回对应的中间件服务器列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListJarPackageHostInfo(request *model.ListJarPackageHostInfoRequest) (*model.ListJarPackageHostInfoResponse, error) {
	requestDef := GenReqDefForListJarPackageHostInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJarPackageHostInfoResponse), nil
	}
}

// ListJarPackageHostInfoInvoker 查询指定中间件的服务器列表
func (c *HssClient) ListJarPackageHostInfoInvoker(request *model.ListJarPackageHostInfoRequest) *ListJarPackageHostInfoInvoker {
	requestDef := GenReqDefForListJarPackageHostInfo()
	return &ListJarPackageHostInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListJarPackageStatistics 查询中间件列表
//
// 查询中间件列表，支持通过中间件名称查询对应的服务器树
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListJarPackageStatistics(request *model.ListJarPackageStatisticsRequest) (*model.ListJarPackageStatisticsResponse, error) {
	requestDef := GenReqDefForListJarPackageStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJarPackageStatisticsResponse), nil
	}
}

// ListJarPackageStatisticsInvoker 查询中间件列表
func (c *HssClient) ListJarPackageStatisticsInvoker(request *model.ListJarPackageStatisticsRequest) *ListJarPackageStatisticsInvoker {
	requestDef := GenReqDefForListJarPackageStatistics()
	return &ListJarPackageStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListK8sCronJobs 查询cronjobs基本信息列表
//
// 查询cronjobs基本信息列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListK8sCronJobs(request *model.ListK8sCronJobsRequest) (*model.ListK8sCronJobsResponse, error) {
	requestDef := GenReqDefForListK8sCronJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListK8sCronJobsResponse), nil
	}
}

// ListK8sCronJobsInvoker 查询cronjobs基本信息列表
func (c *HssClient) ListK8sCronJobsInvoker(request *model.ListK8sCronJobsRequest) *ListK8sCronJobsInvoker {
	requestDef := GenReqDefForListK8sCronJobs()
	return &ListK8sCronJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListK8sDaemonSets 查询daemonsets基本信息列表
//
// 查询daemonsets基本信息列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListK8sDaemonSets(request *model.ListK8sDaemonSetsRequest) (*model.ListK8sDaemonSetsResponse, error) {
	requestDef := GenReqDefForListK8sDaemonSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListK8sDaemonSetsResponse), nil
	}
}

// ListK8sDaemonSetsInvoker 查询daemonsets基本信息列表
func (c *HssClient) ListK8sDaemonSetsInvoker(request *model.ListK8sDaemonSetsRequest) *ListK8sDaemonSetsInvoker {
	requestDef := GenReqDefForListK8sDaemonSets()
	return &ListK8sDaemonSetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListK8sDeployments 查询deployment基本信息列表
//
// 查询deployment基本信息列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListK8sDeployments(request *model.ListK8sDeploymentsRequest) (*model.ListK8sDeploymentsResponse, error) {
	requestDef := GenReqDefForListK8sDeployments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListK8sDeploymentsResponse), nil
	}
}

// ListK8sDeploymentsInvoker 查询deployment基本信息列表
func (c *HssClient) ListK8sDeploymentsInvoker(request *model.ListK8sDeploymentsRequest) *ListK8sDeploymentsInvoker {
	requestDef := GenReqDefForListK8sDeployments()
	return &ListK8sDeploymentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListK8sJobs 查询jobs基本信息列表
//
// 查询jobs基本信息列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListK8sJobs(request *model.ListK8sJobsRequest) (*model.ListK8sJobsResponse, error) {
	requestDef := GenReqDefForListK8sJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListK8sJobsResponse), nil
	}
}

// ListK8sJobsInvoker 查询jobs基本信息列表
func (c *HssClient) ListK8sJobsInvoker(request *model.ListK8sJobsRequest) *ListK8sJobsInvoker {
	requestDef := GenReqDefForListK8sJobs()
	return &ListK8sJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListK8sPods 查询pod基本信息列表
//
// 查询pod基本信息列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListK8sPods(request *model.ListK8sPodsRequest) (*model.ListK8sPodsResponse, error) {
	requestDef := GenReqDefForListK8sPods()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListK8sPodsResponse), nil
	}
}

// ListK8sPodsInvoker 查询pod基本信息列表
func (c *HssClient) ListK8sPodsInvoker(request *model.ListK8sPodsRequest) *ListK8sPodsInvoker {
	requestDef := GenReqDefForListK8sPods()
	return &ListK8sPodsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListK8sStatefulSets 查询statefulset基本信息列表
//
// 查询statefulset基本信息列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListK8sStatefulSets(request *model.ListK8sStatefulSetsRequest) (*model.ListK8sStatefulSetsResponse, error) {
	requestDef := GenReqDefForListK8sStatefulSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListK8sStatefulSetsResponse), nil
	}
}

// ListK8sStatefulSetsInvoker 查询statefulset基本信息列表
func (c *HssClient) ListK8sStatefulSetsInvoker(request *model.ListK8sStatefulSetsRequest) *ListK8sStatefulSetsInvoker {
	requestDef := GenReqDefForListK8sStatefulSets()
	return &ListK8sStatefulSetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListKernelModuleHostInfo 资产管理-资产指纹-内核模块的服务器列表
//
// 资产管理-资产指纹-内核模块的服务器列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListKernelModuleHostInfo(request *model.ListKernelModuleHostInfoRequest) (*model.ListKernelModuleHostInfoResponse, error) {
	requestDef := GenReqDefForListKernelModuleHostInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKernelModuleHostInfoResponse), nil
	}
}

// ListKernelModuleHostInfoInvoker 资产管理-资产指纹-内核模块的服务器列表
func (c *HssClient) ListKernelModuleHostInfoInvoker(request *model.ListKernelModuleHostInfoRequest) *ListKernelModuleHostInfoInvoker {
	requestDef := GenReqDefForListKernelModuleHostInfo()
	return &ListKernelModuleHostInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListKernelModuleStatistics 资产管理-资产指纹-内核模块左侧树
//
// 资产管理-资产指纹-内核模块左侧树
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListKernelModuleStatistics(request *model.ListKernelModuleStatisticsRequest) (*model.ListKernelModuleStatisticsResponse, error) {
	requestDef := GenReqDefForListKernelModuleStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKernelModuleStatisticsResponse), nil
	}
}

// ListKernelModuleStatisticsInvoker 资产管理-资产指纹-内核模块左侧树
func (c *HssClient) ListKernelModuleStatisticsInvoker(request *model.ListKernelModuleStatisticsRequest) *ListKernelModuleStatisticsInvoker {
	requestDef := GenReqDefForListKernelModuleStatistics()
	return &ListKernelModuleStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListKubernetesClusterDetails 查询容器Kubernetes集群列表
//
// 查询容器Kubernetes集群列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListKubernetesClusterDetails(request *model.ListKubernetesClusterDetailsRequest) (*model.ListKubernetesClusterDetailsResponse, error) {
	requestDef := GenReqDefForListKubernetesClusterDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKubernetesClusterDetailsResponse), nil
	}
}

// ListKubernetesClusterDetailsInvoker 查询容器Kubernetes集群列表
func (c *HssClient) ListKubernetesClusterDetailsInvoker(request *model.ListKubernetesClusterDetailsRequest) *ListKubernetesClusterDetailsInvoker {
	requestDef := GenReqDefForListKubernetesClusterDetails()
	return &ListKubernetesClusterDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListKubernetesEndpointDetails 查询容器Kubernetes端点列表
//
// 查询容器Kubernetes端点列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListKubernetesEndpointDetails(request *model.ListKubernetesEndpointDetailsRequest) (*model.ListKubernetesEndpointDetailsResponse, error) {
	requestDef := GenReqDefForListKubernetesEndpointDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKubernetesEndpointDetailsResponse), nil
	}
}

// ListKubernetesEndpointDetailsInvoker 查询容器Kubernetes端点列表
func (c *HssClient) ListKubernetesEndpointDetailsInvoker(request *model.ListKubernetesEndpointDetailsRequest) *ListKubernetesEndpointDetailsInvoker {
	requestDef := GenReqDefForListKubernetesEndpointDetails()
	return &ListKubernetesEndpointDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListKubernetesServiceDetails 查询容器Kubernetes服务列表
//
// 查询容器Kubernetes服务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListKubernetesServiceDetails(request *model.ListKubernetesServiceDetailsRequest) (*model.ListKubernetesServiceDetailsResponse, error) {
	requestDef := GenReqDefForListKubernetesServiceDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKubernetesServiceDetailsResponse), nil
	}
}

// ListKubernetesServiceDetailsInvoker 查询容器Kubernetes服务列表
func (c *HssClient) ListKubernetesServiceDetailsInvoker(request *model.ListKubernetesServiceDetailsRequest) *ListKubernetesServiceDetailsInvoker {
	requestDef := GenReqDefForListKubernetesServiceDetails()
	return &ListKubernetesServiceDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLoginCommonIp 查询常用登录IP信息
//
// 查询常用登录IP信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListLoginCommonIp(request *model.ListLoginCommonIpRequest) (*model.ListLoginCommonIpResponse, error) {
	requestDef := GenReqDefForListLoginCommonIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLoginCommonIpResponse), nil
	}
}

// ListLoginCommonIpInvoker 查询常用登录IP信息
func (c *HssClient) ListLoginCommonIpInvoker(request *model.ListLoginCommonIpRequest) *ListLoginCommonIpInvoker {
	requestDef := GenReqDefForListLoginCommonIp()
	return &ListLoginCommonIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLoginCommonLocation 查询常用登录地信息
//
// 查询常用登录地信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListLoginCommonLocation(request *model.ListLoginCommonLocationRequest) (*model.ListLoginCommonLocationResponse, error) {
	requestDef := GenReqDefForListLoginCommonLocation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLoginCommonLocationResponse), nil
	}
}

// ListLoginCommonLocationInvoker 查询常用登录地信息
func (c *HssClient) ListLoginCommonLocationInvoker(request *model.ListLoginCommonLocationRequest) *ListLoginCommonLocationInvoker {
	requestDef := GenReqDefForListLoginCommonLocation()
	return &ListLoginCommonLocationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLoginWhiteIp 查询SSH登录IP白名单列表
//
// 查询SSH登录IP白名单列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListLoginWhiteIp(request *model.ListLoginWhiteIpRequest) (*model.ListLoginWhiteIpResponse, error) {
	requestDef := GenReqDefForListLoginWhiteIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLoginWhiteIpResponse), nil
	}
}

// ListLoginWhiteIpInvoker 查询SSH登录IP白名单列表
func (c *HssClient) ListLoginWhiteIpInvoker(request *model.ListLoginWhiteIpRequest) *ListLoginWhiteIpInvoker {
	requestDef := GenReqDefForListLoginWhiteIp()
	return &ListLoginWhiteIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLoginWhiteList 查询登录白名单列表
//
// 查询登录白名单列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListLoginWhiteList(request *model.ListLoginWhiteListRequest) (*model.ListLoginWhiteListResponse, error) {
	requestDef := GenReqDefForListLoginWhiteList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLoginWhiteListResponse), nil
	}
}

// ListLoginWhiteListInvoker 查询登录白名单列表
func (c *HssClient) ListLoginWhiteListInvoker(request *model.ListLoginWhiteListRequest) *ListLoginWhiteListInvoker {
	requestDef := GenReqDefForListLoginWhiteList()
	return &ListLoginWhiteListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMultiCloudClusters 查询多云集群
//
// 查询多云集群
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListMultiCloudClusters(request *model.ListMultiCloudClustersRequest) (*model.ListMultiCloudClustersResponse, error) {
	requestDef := GenReqDefForListMultiCloudClusters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMultiCloudClustersResponse), nil
	}
}

// ListMultiCloudClustersInvoker 查询多云集群
func (c *HssClient) ListMultiCloudClustersInvoker(request *model.ListMultiCloudClustersRequest) *ListMultiCloudClustersInvoker {
	requestDef := GenReqDefForListMultiCloudClusters()
	return &ListMultiCloudClustersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNamespaces 获取集群下的namespace
//
// 获取集群下的namespace
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListNamespaces(request *model.ListNamespacesRequest) (*model.ListNamespacesResponse, error) {
	requestDef := GenReqDefForListNamespaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNamespacesResponse), nil
	}
}

// ListNamespacesInvoker 获取集群下的namespace
func (c *HssClient) ListNamespacesInvoker(request *model.ListNamespacesRequest) *ListNamespacesInvoker {
	requestDef := GenReqDefForListNamespaces()
	return &ListNamespacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOperationLogsByVaultName 查询备份恢复任务列表
//
// 查询备份恢复任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListOperationLogsByVaultName(request *model.ListOperationLogsByVaultNameRequest) (*model.ListOperationLogsByVaultNameResponse, error) {
	requestDef := GenReqDefForListOperationLogsByVaultName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOperationLogsByVaultNameResponse), nil
	}
}

// ListOperationLogsByVaultNameInvoker 查询备份恢复任务列表
func (c *HssClient) ListOperationLogsByVaultNameInvoker(request *model.ListOperationLogsByVaultNameRequest) *ListOperationLogsByVaultNameInvoker {
	requestDef := GenReqDefForListOperationLogsByVaultName()
	return &ListOperationLogsByVaultNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOrganizationTree 查询账号组织
//
// 查询账号组织
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListOrganizationTree(request *model.ListOrganizationTreeRequest) (*model.ListOrganizationTreeResponse, error) {
	requestDef := GenReqDefForListOrganizationTree()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOrganizationTreeResponse), nil
	}
}

// ListOrganizationTreeInvoker 查询账号组织
func (c *HssClient) ListOrganizationTreeInvoker(request *model.ListOrganizationTreeRequest) *ListOrganizationTreeInvoker {
	requestDef := GenReqDefForListOrganizationTree()
	return &ListOrganizationTreeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPasswordComplexity 查询口令复杂度策略检测报告
//
// 查询口令复杂度策略检测报告
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListPasswordComplexity(request *model.ListPasswordComplexityRequest) (*model.ListPasswordComplexityResponse, error) {
	requestDef := GenReqDefForListPasswordComplexity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPasswordComplexityResponse), nil
	}
}

// ListPasswordComplexityInvoker 查询口令复杂度策略检测报告
func (c *HssClient) ListPasswordComplexityInvoker(request *model.ListPasswordComplexityRequest) *ListPasswordComplexityInvoker {
	requestDef := GenReqDefForListPasswordComplexity()
	return &ListPasswordComplexityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicies 查询主机策略列表
//
// 查询主机策略列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListPolicies(request *model.ListPoliciesRequest) (*model.ListPoliciesResponse, error) {
	requestDef := GenReqDefForListPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPoliciesResponse), nil
	}
}

// ListPoliciesInvoker 查询主机策略列表
func (c *HssClient) ListPoliciesInvoker(request *model.ListPoliciesRequest) *ListPoliciesInvoker {
	requestDef := GenReqDefForListPolicies()
	return &ListPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicyGroup 查询策略组列表
//
// 查询策略组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListPolicyGroup(request *model.ListPolicyGroupRequest) (*model.ListPolicyGroupResponse, error) {
	requestDef := GenReqDefForListPolicyGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyGroupResponse), nil
	}
}

// ListPolicyGroupInvoker 查询策略组列表
func (c *HssClient) ListPolicyGroupInvoker(request *model.ListPolicyGroupRequest) *ListPolicyGroupInvoker {
	requestDef := GenReqDefForListPolicyGroup()
	return &ListPolicyGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPortHost 资产指纹-端口-服务器列表
//
// 具备该端口的主机/容器信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListPortHost(request *model.ListPortHostRequest) (*model.ListPortHostResponse, error) {
	requestDef := GenReqDefForListPortHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPortHostResponse), nil
	}
}

// ListPortHostInvoker 资产指纹-端口-服务器列表
func (c *HssClient) ListPortHostInvoker(request *model.ListPortHostRequest) *ListPortHostInvoker {
	requestDef := GenReqDefForListPortHost()
	return &ListPortHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPortStatistics 查询开放端口统计信息
//
// 查询开放端口列表，支持通过传入端口或协议类型查询服务器数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListPortStatistics(request *model.ListPortStatisticsRequest) (*model.ListPortStatisticsResponse, error) {
	requestDef := GenReqDefForListPortStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPortStatisticsResponse), nil
	}
}

// ListPortStatisticsInvoker 查询开放端口统计信息
func (c *HssClient) ListPortStatisticsInvoker(request *model.ListPortStatisticsRequest) *ListPortStatisticsInvoker {
	requestDef := GenReqDefForListPortStatistics()
	return &ListPortStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPorts 查询单服务器的开放端口列表
//
// 查询单服务器的开放端口列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListPorts(request *model.ListPortsRequest) (*model.ListPortsResponse, error) {
	requestDef := GenReqDefForListPorts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPortsResponse), nil
	}
}

// ListPortsInvoker 查询单服务器的开放端口列表
func (c *HssClient) ListPortsInvoker(request *model.ListPortsRequest) *ListPortsInvoker {
	requestDef := GenReqDefForListPorts()
	return &ListPortsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProcessStatistics 查询进程列表
//
// 查询进程列表，通过传入进程路径参数查询对应的服务器数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListProcessStatistics(request *model.ListProcessStatisticsRequest) (*model.ListProcessStatisticsResponse, error) {
	requestDef := GenReqDefForListProcessStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProcessStatisticsResponse), nil
	}
}

// ListProcessStatisticsInvoker 查询进程列表
func (c *HssClient) ListProcessStatisticsInvoker(request *model.ListProcessStatisticsRequest) *ListProcessStatisticsInvoker {
	requestDef := GenReqDefForListProcessStatistics()
	return &ListProcessStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProcessesHost 资产指纹-进程-服务器列表
//
// 具备该进程的主机/容器信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListProcessesHost(request *model.ListProcessesHostRequest) (*model.ListProcessesHostResponse, error) {
	requestDef := GenReqDefForListProcessesHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProcessesHostResponse), nil
	}
}

// ListProcessesHostInvoker 资产指纹-进程-服务器列表
func (c *HssClient) ListProcessesHostInvoker(request *model.ListProcessesHostRequest) *ListProcessesHostInvoker {
	requestDef := GenReqDefForListProcessesHost()
	return &ListProcessesHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectTags 查询租户当前项目下所有用过的标签
//
// 查询租户当前项目下所有用过的标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListProjectTags(request *model.ListProjectTagsRequest) (*model.ListProjectTagsResponse, error) {
	requestDef := GenReqDefForListProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectTagsResponse), nil
	}
}

// ListProjectTagsInvoker 查询租户当前项目下所有用过的标签
func (c *HssClient) ListProjectTagsInvoker(request *model.ListProjectTagsRequest) *ListProjectTagsInvoker {
	requestDef := GenReqDefForListProjectTags()
	return &ListProjectTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProtectionPolicy 查询勒索病毒的防护策略列表
//
// 查询勒索病毒的防护策略列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListProtectionPolicy(request *model.ListProtectionPolicyRequest) (*model.ListProtectionPolicyResponse, error) {
	requestDef := GenReqDefForListProtectionPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProtectionPolicyResponse), nil
	}
}

// ListProtectionPolicyInvoker 查询勒索病毒的防护策略列表
func (c *HssClient) ListProtectionPolicyInvoker(request *model.ListProtectionPolicyRequest) *ListProtectionPolicyInvoker {
	requestDef := GenReqDefForListProtectionPolicy()
	return &ListProtectionPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProtectionServer 查询勒索防护服务器列表
//
// 查询勒索防护服务器列表，与云备份服务配合使用。因此使用勒索相关接口之前确保该局点有云备份服务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListProtectionServer(request *model.ListProtectionServerRequest) (*model.ListProtectionServerResponse, error) {
	requestDef := GenReqDefForListProtectionServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProtectionServerResponse), nil
	}
}

// ListProtectionServerInvoker 查询勒索防护服务器列表
func (c *HssClient) ListProtectionServerInvoker(request *model.ListProtectionServerRequest) *ListProtectionServerInvoker {
	requestDef := GenReqDefForListProtectionServer()
	return &ListProtectionServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProtectionServers 查询防护服务器列表
//
// 查询防护服务器列表：查询防护服务器相关数据，包含服务器名称、ip地址、操作系统、服务器组名称、防护策略、防护状态、微服务防护状态、RASP防护状态、RASP攻击数量信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListProtectionServers(request *model.ListProtectionServersRequest) (*model.ListProtectionServersResponse, error) {
	requestDef := GenReqDefForListProtectionServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProtectionServersResponse), nil
	}
}

// ListProtectionServersInvoker 查询防护服务器列表
func (c *HssClient) ListProtectionServersInvoker(request *model.ListProtectionServersRequest) *ListProtectionServersInvoker {
	requestDef := GenReqDefForListProtectionServers()
	return &ListProtectionServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQueryExportTask 查询导出任务信息
//
// 查询导出任务信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListQueryExportTask(request *model.ListQueryExportTaskRequest) (*model.ListQueryExportTaskResponse, error) {
	requestDef := GenReqDefForListQueryExportTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQueryExportTaskResponse), nil
	}
}

// ListQueryExportTaskInvoker 查询导出任务信息
func (c *HssClient) ListQueryExportTaskInvoker(request *model.ListQueryExportTaskRequest) *ListQueryExportTaskInvoker {
	requestDef := GenReqDefForListQueryExportTask()
	return &ListQueryExportTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQuotasDetail 查询配额详情
//
// 查询配额详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListQuotasDetail(request *model.ListQuotasDetailRequest) (*model.ListQuotasDetailResponse, error) {
	requestDef := GenReqDefForListQuotasDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotasDetailResponse), nil
	}
}

// ListQuotasDetailInvoker 查询配额详情
func (c *HssClient) ListQuotasDetailInvoker(request *model.ListQuotasDetailRequest) *ListQuotasDetailInvoker {
	requestDef := GenReqDefForListQuotasDetail()
	return &ListQuotasDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRansomwareProtectionNodes 查询勒索防护服务器列表2.0
//
// 查询勒索防护服务器列表，与云备份服务配合使用。因此使用勒索相关接口之前确保该局点有云备份服务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListRansomwareProtectionNodes(request *model.ListRansomwareProtectionNodesRequest) (*model.ListRansomwareProtectionNodesResponse, error) {
	requestDef := GenReqDefForListRansomwareProtectionNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRansomwareProtectionNodesResponse), nil
	}
}

// ListRansomwareProtectionNodesInvoker 查询勒索防护服务器列表2.0
func (c *HssClient) ListRansomwareProtectionNodesInvoker(request *model.ListRansomwareProtectionNodesRequest) *ListRansomwareProtectionNodesInvoker {
	requestDef := GenReqDefForListRansomwareProtectionNodes()
	return &ListRansomwareProtectionNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRaspEvents 查询应用防护事件列表
//
// 查询应用防护事件列表：展示防护事件相关信息，包含告警级别、服务器名称、告警名称、告警时间、攻击源ip、攻击源url数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListRaspEvents(request *model.ListRaspEventsRequest) (*model.ListRaspEventsResponse, error) {
	requestDef := GenReqDefForListRaspEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRaspEventsResponse), nil
	}
}

// ListRaspEventsInvoker 查询应用防护事件列表
func (c *HssClient) ListRaspEventsInvoker(request *model.ListRaspEventsRequest) *ListRaspEventsInvoker {
	requestDef := GenReqDefForListRaspEvents()
	return &ListRaspEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRaspPolicies 查询防护策略列表
//
// 查询防护策略列表：查询创建的防护策略信息，包含防护策略名称、检测规则、关联服务器数量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListRaspPolicies(request *model.ListRaspPoliciesRequest) (*model.ListRaspPoliciesResponse, error) {
	requestDef := GenReqDefForListRaspPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRaspPoliciesResponse), nil
	}
}

// ListRaspPoliciesInvoker 查询防护策略列表
func (c *HssClient) ListRaspPoliciesInvoker(request *model.ListRaspPoliciesRequest) *ListRaspPoliciesInvoker {
	requestDef := GenReqDefForListRaspPolicies()
	return &ListRaspPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRiskConfigCheckRules 查询指定安全配置项的检查项列表
//
// 查询指定安全配置项的检查项列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListRiskConfigCheckRules(request *model.ListRiskConfigCheckRulesRequest) (*model.ListRiskConfigCheckRulesResponse, error) {
	requestDef := GenReqDefForListRiskConfigCheckRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRiskConfigCheckRulesResponse), nil
	}
}

// ListRiskConfigCheckRulesInvoker 查询指定安全配置项的检查项列表
func (c *HssClient) ListRiskConfigCheckRulesInvoker(request *model.ListRiskConfigCheckRulesRequest) *ListRiskConfigCheckRulesInvoker {
	requestDef := GenReqDefForListRiskConfigCheckRules()
	return &ListRiskConfigCheckRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRiskConfigHosts 查询指定安全配置项的受影响服务器列表
//
// 查询指定安全配置项的受影响服务器列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListRiskConfigHosts(request *model.ListRiskConfigHostsRequest) (*model.ListRiskConfigHostsResponse, error) {
	requestDef := GenReqDefForListRiskConfigHosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRiskConfigHostsResponse), nil
	}
}

// ListRiskConfigHostsInvoker 查询指定安全配置项的受影响服务器列表
func (c *HssClient) ListRiskConfigHostsInvoker(request *model.ListRiskConfigHostsRequest) *ListRiskConfigHostsInvoker {
	requestDef := GenReqDefForListRiskConfigHosts()
	return &ListRiskConfigHostsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRiskConfigs 查询租户的服务器安全配置检测结果列表
//
// 查询租户的服务器安全配置检测结果列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListRiskConfigs(request *model.ListRiskConfigsRequest) (*model.ListRiskConfigsResponse, error) {
	requestDef := GenReqDefForListRiskConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRiskConfigsResponse), nil
	}
}

// ListRiskConfigsInvoker 查询租户的服务器安全配置检测结果列表
func (c *HssClient) ListRiskConfigsInvoker(request *model.ListRiskConfigsRequest) *ListRiskConfigsInvoker {
	requestDef := GenReqDefForListRiskConfigs()
	return &ListRiskConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecurityEvents 查入侵事件列表
//
// 查入侵事件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListSecurityEvents(request *model.ListSecurityEventsRequest) (*model.ListSecurityEventsResponse, error) {
	requestDef := GenReqDefForListSecurityEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityEventsResponse), nil
	}
}

// ListSecurityEventsInvoker 查入侵事件列表
func (c *HssClient) ListSecurityEventsInvoker(request *model.ListSecurityEventsRequest) *ListSecurityEventsInvoker {
	requestDef := GenReqDefForListSecurityEvents()
	return &ListSecurityEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecurityGroupPolicies 查询云原生网络模式2.0的集群已配置的安全组策略
//
// 查询云原生网络模式2.0的集群已配置的安全组策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListSecurityGroupPolicies(request *model.ListSecurityGroupPoliciesRequest) (*model.ListSecurityGroupPoliciesResponse, error) {
	requestDef := GenReqDefForListSecurityGroupPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityGroupPoliciesResponse), nil
	}
}

// ListSecurityGroupPoliciesInvoker 查询云原生网络模式2.0的集群已配置的安全组策略
func (c *HssClient) ListSecurityGroupPoliciesInvoker(request *model.ListSecurityGroupPoliciesRequest) *ListSecurityGroupPoliciesInvoker {
	requestDef := GenReqDefForListSecurityGroupPolicies()
	return &ListSecurityGroupPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecurityGroups 查询企业项目下所有的安全组列表
//
// 查询企业项目下所有的安全组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListSecurityGroups(request *model.ListSecurityGroupsRequest) (*model.ListSecurityGroupsResponse, error) {
	requestDef := GenReqDefForListSecurityGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityGroupsResponse), nil
	}
}

// ListSecurityGroupsInvoker 查询企业项目下所有的安全组列表
func (c *HssClient) ListSecurityGroupsInvoker(request *model.ListSecurityGroupsRequest) *ListSecurityGroupsInvoker {
	requestDef := GenReqDefForListSecurityGroups()
	return &ListSecurityGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSwrImageRepository 查询swr镜像仓库镜像列表
//
// 查询swr镜像仓库镜像列表,如果需要从swr同步最新镜像，需要先调用“从swr同步镜像”接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListSwrImageRepository(request *model.ListSwrImageRepositoryRequest) (*model.ListSwrImageRepositoryResponse, error) {
	requestDef := GenReqDefForListSwrImageRepository()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSwrImageRepositoryResponse), nil
	}
}

// ListSwrImageRepositoryInvoker 查询swr镜像仓库镜像列表
func (c *HssClient) ListSwrImageRepositoryInvoker(request *model.ListSwrImageRepositoryRequest) *ListSwrImageRepositoryInvoker {
	requestDef := GenReqDefForListSwrImageRepository()
	return &ListSwrImageRepositoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSystemUserWhiteList 查询系统用白名单列表
//
// 查询系统用户白名单列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListSystemUserWhiteList(request *model.ListSystemUserWhiteListRequest) (*model.ListSystemUserWhiteListResponse, error) {
	requestDef := GenReqDefForListSystemUserWhiteList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSystemUserWhiteListResponse), nil
	}
}

// ListSystemUserWhiteListInvoker 查询系统用白名单列表
func (c *HssClient) ListSystemUserWhiteListInvoker(request *model.ListSystemUserWhiteListRequest) *ListSystemUserWhiteListInvoker {
	requestDef := GenReqDefForListSystemUserWhiteList()
	return &ListSystemUserWhiteListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTwoFactorLoginHost 查询双因子主机列表
//
// 查询双因子主机列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListTwoFactorLoginHost(request *model.ListTwoFactorLoginHostRequest) (*model.ListTwoFactorLoginHostResponse, error) {
	requestDef := GenReqDefForListTwoFactorLoginHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTwoFactorLoginHostResponse), nil
	}
}

// ListTwoFactorLoginHostInvoker 查询双因子主机列表
func (c *HssClient) ListTwoFactorLoginHostInvoker(request *model.ListTwoFactorLoginHostRequest) *ListTwoFactorLoginHostInvoker {
	requestDef := GenReqDefForListTwoFactorLoginHost()
	return &ListTwoFactorLoginHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserChangeHistories 获取账户变动历史信息
//
// 获取账户变动历史记录信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListUserChangeHistories(request *model.ListUserChangeHistoriesRequest) (*model.ListUserChangeHistoriesResponse, error) {
	requestDef := GenReqDefForListUserChangeHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserChangeHistoriesResponse), nil
	}
}

// ListUserChangeHistoriesInvoker 获取账户变动历史信息
func (c *HssClient) ListUserChangeHistoriesInvoker(request *model.ListUserChangeHistoriesRequest) *ListUserChangeHistoriesInvoker {
	requestDef := GenReqDefForListUserChangeHistories()
	return &ListUserChangeHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserStatistics 查询账号信息列表
//
// 查询账号信息列表，支持通过传入账号名称参数查询对应的服务器数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListUserStatistics(request *model.ListUserStatisticsRequest) (*model.ListUserStatisticsResponse, error) {
	requestDef := GenReqDefForListUserStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserStatisticsResponse), nil
	}
}

// ListUserStatisticsInvoker 查询账号信息列表
func (c *HssClient) ListUserStatisticsInvoker(request *model.ListUserStatisticsRequest) *ListUserStatisticsInvoker {
	requestDef := GenReqDefForListUserStatistics()
	return &ListUserStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUsers 查询账号的服务器列表
//
// 查询账号的服务器列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListUsers(request *model.ListUsersRequest) (*model.ListUsersResponse, error) {
	requestDef := GenReqDefForListUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsersResponse), nil
	}
}

// ListUsersInvoker 查询账号的服务器列表
func (c *HssClient) ListUsersInvoker(request *model.ListUsersRequest) *ListUsersInvoker {
	requestDef := GenReqDefForListUsers()
	return &ListUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVulHandleHistory 查询漏洞历史处置记录列表
//
// 查询漏洞历史处置记录列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListVulHandleHistory(request *model.ListVulHandleHistoryRequest) (*model.ListVulHandleHistoryResponse, error) {
	requestDef := GenReqDefForListVulHandleHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVulHandleHistoryResponse), nil
	}
}

// ListVulHandleHistoryInvoker 查询漏洞历史处置记录列表
func (c *HssClient) ListVulHandleHistoryInvoker(request *model.ListVulHandleHistoryRequest) *ListVulHandleHistoryInvoker {
	requestDef := GenReqDefForListVulHandleHistory()
	return &ListVulHandleHistoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVulHosts 查询单个漏洞影响的云服务器信息
//
// 查询单个漏洞影响的云服务器信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListVulHosts(request *model.ListVulHostsRequest) (*model.ListVulHostsResponse, error) {
	requestDef := GenReqDefForListVulHosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVulHostsResponse), nil
	}
}

// ListVulHostsInvoker 查询单个漏洞影响的云服务器信息
func (c *HssClient) ListVulHostsInvoker(request *model.ListVulHostsRequest) *ListVulHostsInvoker {
	requestDef := GenReqDefForListVulHosts()
	return &ListVulHostsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVulScanTask 查询漏洞扫描任务列表
//
// 查询漏洞扫描任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListVulScanTask(request *model.ListVulScanTaskRequest) (*model.ListVulScanTaskResponse, error) {
	requestDef := GenReqDefForListVulScanTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVulScanTaskResponse), nil
	}
}

// ListVulScanTaskInvoker 查询漏洞扫描任务列表
func (c *HssClient) ListVulScanTaskInvoker(request *model.ListVulScanTaskRequest) *ListVulScanTaskInvoker {
	requestDef := GenReqDefForListVulScanTask()
	return &ListVulScanTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVulScanTaskHost 查询漏洞扫描任务对应的主机列表
//
// 查询漏洞扫描任务对应的主机列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListVulScanTaskHost(request *model.ListVulScanTaskHostRequest) (*model.ListVulScanTaskHostResponse, error) {
	requestDef := GenReqDefForListVulScanTaskHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVulScanTaskHostResponse), nil
	}
}

// ListVulScanTaskHostInvoker 查询漏洞扫描任务对应的主机列表
func (c *HssClient) ListVulScanTaskHostInvoker(request *model.ListVulScanTaskHostRequest) *ListVulScanTaskHostInvoker {
	requestDef := GenReqDefForListVulScanTaskHost()
	return &ListVulScanTaskHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVulnerabilities 查询漏洞列表
//
// 查询漏洞列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListVulnerabilities(request *model.ListVulnerabilitiesRequest) (*model.ListVulnerabilitiesResponse, error) {
	requestDef := GenReqDefForListVulnerabilities()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVulnerabilitiesResponse), nil
	}
}

// ListVulnerabilitiesInvoker 查询漏洞列表
func (c *HssClient) ListVulnerabilitiesInvoker(request *model.ListVulnerabilitiesRequest) *ListVulnerabilitiesInvoker {
	requestDef := GenReqDefForListVulnerabilities()
	return &ListVulnerabilitiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVulnerabilityCve 漏洞对应cve信息
//
// 漏洞对应cve信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListVulnerabilityCve(request *model.ListVulnerabilityCveRequest) (*model.ListVulnerabilityCveResponse, error) {
	requestDef := GenReqDefForListVulnerabilityCve()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVulnerabilityCveResponse), nil
	}
}

// ListVulnerabilityCveInvoker 漏洞对应cve信息
func (c *HssClient) ListVulnerabilityCveInvoker(request *model.ListVulnerabilityCveRequest) *ListVulnerabilityCveInvoker {
	requestDef := GenReqDefForListVulnerabilityCve()
	return &ListVulnerabilityCveInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWeakPasswordUsers 查询弱口令检测结果列表
//
// 查询弱口令检测结果列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListWeakPasswordUsers(request *model.ListWeakPasswordUsersRequest) (*model.ListWeakPasswordUsersResponse, error) {
	requestDef := GenReqDefForListWeakPasswordUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWeakPasswordUsersResponse), nil
	}
}

// ListWeakPasswordUsersInvoker 查询弱口令检测结果列表
func (c *HssClient) ListWeakPasswordUsersInvoker(request *model.ListWeakPasswordUsersRequest) *ListWeakPasswordUsersInvoker {
	requestDef := GenReqDefForListWeakPasswordUsers()
	return &ListWeakPasswordUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWebAppAndServiceStatistics 资产管理-资产指纹-左侧WebAppAndService名称树信息
//
// 资产管理-资产指纹-左侧WebAppAndService名称树信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListWebAppAndServiceStatistics(request *model.ListWebAppAndServiceStatisticsRequest) (*model.ListWebAppAndServiceStatisticsResponse, error) {
	requestDef := GenReqDefForListWebAppAndServiceStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWebAppAndServiceStatisticsResponse), nil
	}
}

// ListWebAppAndServiceStatisticsInvoker 资产管理-资产指纹-左侧WebAppAndService名称树信息
func (c *HssClient) ListWebAppAndServiceStatisticsInvoker(request *model.ListWebAppAndServiceStatisticsRequest) *ListWebAppAndServiceStatisticsInvoker {
	requestDef := GenReqDefForListWebAppAndServiceStatistics()
	return &ListWebAppAndServiceStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWebAppAndServices 资产管理-资产指纹-右侧WebAppAndService资产信息
//
// 资产管理-资产指纹-右侧WebAppAndService资产信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListWebAppAndServices(request *model.ListWebAppAndServicesRequest) (*model.ListWebAppAndServicesResponse, error) {
	requestDef := GenReqDefForListWebAppAndServices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWebAppAndServicesResponse), nil
	}
}

// ListWebAppAndServicesInvoker 资产管理-资产指纹-右侧WebAppAndService资产信息
func (c *HssClient) ListWebAppAndServicesInvoker(request *model.ListWebAppAndServicesRequest) *ListWebAppAndServicesInvoker {
	requestDef := GenReqDefForListWebAppAndServices()
	return &ListWebAppAndServicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWebFrameworkHostInfo 资产管理-资产指纹-Web框架的服务器列表
//
// 资产管理-资产指纹-Web框架的服务器列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListWebFrameworkHostInfo(request *model.ListWebFrameworkHostInfoRequest) (*model.ListWebFrameworkHostInfoResponse, error) {
	requestDef := GenReqDefForListWebFrameworkHostInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWebFrameworkHostInfoResponse), nil
	}
}

// ListWebFrameworkHostInfoInvoker 资产管理-资产指纹-Web框架的服务器列表
func (c *HssClient) ListWebFrameworkHostInfoInvoker(request *model.ListWebFrameworkHostInfoRequest) *ListWebFrameworkHostInfoInvoker {
	requestDef := GenReqDefForListWebFrameworkHostInfo()
	return &ListWebFrameworkHostInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWebFrameworkStatistics 资产管理-资产指纹-Web框架左侧树
//
// 资产管理-资产指纹-Web框架左侧树
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListWebFrameworkStatistics(request *model.ListWebFrameworkStatisticsRequest) (*model.ListWebFrameworkStatisticsResponse, error) {
	requestDef := GenReqDefForListWebFrameworkStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWebFrameworkStatisticsResponse), nil
	}
}

// ListWebFrameworkStatisticsInvoker 资产管理-资产指纹-Web框架左侧树
func (c *HssClient) ListWebFrameworkStatisticsInvoker(request *model.ListWebFrameworkStatisticsRequest) *ListWebFrameworkStatisticsInvoker {
	requestDef := GenReqDefForListWebFrameworkStatistics()
	return &ListWebFrameworkStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWebSiteHostInfo 资产管理-资产指纹-Web站点的服务器列表
//
// 资产管理-资产指纹-Web站点的服务器列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListWebSiteHostInfo(request *model.ListWebSiteHostInfoRequest) (*model.ListWebSiteHostInfoResponse, error) {
	requestDef := GenReqDefForListWebSiteHostInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWebSiteHostInfoResponse), nil
	}
}

// ListWebSiteHostInfoInvoker 资产管理-资产指纹-Web站点的服务器列表
func (c *HssClient) ListWebSiteHostInfoInvoker(request *model.ListWebSiteHostInfoRequest) *ListWebSiteHostInfoInvoker {
	requestDef := GenReqDefForListWebSiteHostInfo()
	return &ListWebSiteHostInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWebSiteStatistics 资产管理-资产指纹-Web站点左侧树
//
// 资产管理-资产指纹-Web站点左侧树
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListWebSiteStatistics(request *model.ListWebSiteStatisticsRequest) (*model.ListWebSiteStatisticsResponse, error) {
	requestDef := GenReqDefForListWebSiteStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWebSiteStatisticsResponse), nil
	}
}

// ListWebSiteStatisticsInvoker 资产管理-资产指纹-Web站点左侧树
func (c *HssClient) ListWebSiteStatisticsInvoker(request *model.ListWebSiteStatisticsRequest) *ListWebSiteStatisticsInvoker {
	requestDef := GenReqDefForListWebSiteStatistics()
	return &ListWebSiteStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkLoads 查询集群下某一命名空间下的工作负载
//
// 查询集群下某一命名空间下的工作负载
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListWorkLoads(request *model.ListWorkLoadsRequest) (*model.ListWorkLoadsResponse, error) {
	requestDef := GenReqDefForListWorkLoads()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkLoadsResponse), nil
	}
}

// ListWorkLoadsInvoker 查询集群下某一命名空间下的工作负载
func (c *HssClient) ListWorkLoadsInvoker(request *model.ListWorkLoadsRequest) *ListWorkLoadsInvoker {
	requestDef := GenReqDefForListWorkLoads()
	return &ListWorkLoadsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWtpProtectHost 查询防护列表
//
// 查询防护列表：查询网页防篡改主机防护状态列表信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListWtpProtectHost(request *model.ListWtpProtectHostRequest) (*model.ListWtpProtectHostResponse, error) {
	requestDef := GenReqDefForListWtpProtectHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWtpProtectHostResponse), nil
	}
}

// ListWtpProtectHostInvoker 查询防护列表
func (c *HssClient) ListWtpProtectHostInvoker(request *model.ListWtpProtectHostRequest) *ListWtpProtectHostInvoker {
	requestDef := GenReqDefForListWtpProtectHost()
	return &ListWtpProtectHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyDecoyPortPolicy 编辑动态端口蜜罐策略
//
// 编辑动态端口蜜罐策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ModifyDecoyPortPolicy(request *model.ModifyDecoyPortPolicyRequest) (*model.ModifyDecoyPortPolicyResponse, error) {
	requestDef := GenReqDefForModifyDecoyPortPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyDecoyPortPolicyResponse), nil
	}
}

// ModifyDecoyPortPolicyInvoker 编辑动态端口蜜罐策略
func (c *HssClient) ModifyDecoyPortPolicyInvoker(request *model.ModifyDecoyPortPolicyRequest) *ModifyDecoyPortPolicyInvoker {
	requestDef := GenReqDefForModifyDecoyPortPolicy()
	return &ModifyDecoyPortPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ParseMultiCloudClusterConfig 解析多云集群的配置文件
//
// 解析多云集群的配置文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ParseMultiCloudClusterConfig(request *model.ParseMultiCloudClusterConfigRequest) (*model.ParseMultiCloudClusterConfigResponse, error) {
	requestDef := GenReqDefForParseMultiCloudClusterConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ParseMultiCloudClusterConfigResponse), nil
	}
}

// ParseMultiCloudClusterConfigInvoker 解析多云集群的配置文件
func (c *HssClient) ParseMultiCloudClusterConfigInvoker(request *model.ParseMultiCloudClusterConfigRequest) *ParseMultiCloudClusterConfigInvoker {
	requestDef := GenReqDefForParseMultiCloudClusterConfig()
	return &ParseMultiCloudClusterConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveAlarmWhiteList 删除告警白名单
//
// 删除告警白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) RemoveAlarmWhiteList(request *model.RemoveAlarmWhiteListRequest) (*model.RemoveAlarmWhiteListResponse, error) {
	requestDef := GenReqDefForRemoveAlarmWhiteList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveAlarmWhiteListResponse), nil
	}
}

// RemoveAlarmWhiteListInvoker 删除告警白名单
func (c *HssClient) RemoveAlarmWhiteListInvoker(request *model.RemoveAlarmWhiteListRequest) *RemoveAlarmWhiteListInvoker {
	requestDef := GenReqDefForRemoveAlarmWhiteList()
	return &RemoveAlarmWhiteListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveLoginWhiteList 删除登录白名单
//
// 删除登录白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) RemoveLoginWhiteList(request *model.RemoveLoginWhiteListRequest) (*model.RemoveLoginWhiteListResponse, error) {
	requestDef := GenReqDefForRemoveLoginWhiteList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveLoginWhiteListResponse), nil
	}
}

// RemoveLoginWhiteListInvoker 删除登录白名单
func (c *HssClient) RemoveLoginWhiteListInvoker(request *model.RemoveLoginWhiteListRequest) *RemoveLoginWhiteListInvoker {
	requestDef := GenReqDefForRemoveLoginWhiteList()
	return &RemoveLoginWhiteListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveMultiCloudClusters 删除多云集群
//
// 删除多云集群
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) RemoveMultiCloudClusters(request *model.RemoveMultiCloudClustersRequest) (*model.RemoveMultiCloudClustersResponse, error) {
	requestDef := GenReqDefForRemoveMultiCloudClusters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveMultiCloudClustersResponse), nil
	}
}

// RemoveMultiCloudClustersInvoker 删除多云集群
func (c *HssClient) RemoveMultiCloudClustersInvoker(request *model.RemoveMultiCloudClustersRequest) *RemoveMultiCloudClustersInvoker {
	requestDef := GenReqDefForRemoveMultiCloudClusters()
	return &RemoveMultiCloudClustersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveSystemUserWhiteList 删除系统用户白名单
//
// 删除系统用户白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) RemoveSystemUserWhiteList(request *model.RemoveSystemUserWhiteListRequest) (*model.RemoveSystemUserWhiteListResponse, error) {
	requestDef := GenReqDefForRemoveSystemUserWhiteList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveSystemUserWhiteListResponse), nil
	}
}

// RemoveSystemUserWhiteListInvoker 删除系统用户白名单
func (c *HssClient) RemoveSystemUserWhiteListInvoker(request *model.RemoveSystemUserWhiteListRequest) *RemoveSystemUserWhiteListInvoker {
	requestDef := GenReqDefForRemoveSystemUserWhiteList()
	return &RemoveSystemUserWhiteListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunHostAssetManualCollect 采集单主机资产指纹
//
// 采集单主机资产指纹
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) RunHostAssetManualCollect(request *model.RunHostAssetManualCollectRequest) (*model.RunHostAssetManualCollectResponse, error) {
	requestDef := GenReqDefForRunHostAssetManualCollect()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunHostAssetManualCollectResponse), nil
	}
}

// RunHostAssetManualCollectInvoker 采集单主机资产指纹
func (c *HssClient) RunHostAssetManualCollectInvoker(request *model.RunHostAssetManualCollectRequest) *RunHostAssetManualCollectInvoker {
	requestDef := GenReqDefForRunHostAssetManualCollect()
	return &RunHostAssetManualCollectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunImageSynchronize 从SWR服务同步镜像列表
//
// 从SWR服务同步镜像列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) RunImageSynchronize(request *model.RunImageSynchronizeRequest) (*model.RunImageSynchronizeResponse, error) {
	requestDef := GenReqDefForRunImageSynchronize()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunImageSynchronizeResponse), nil
	}
}

// RunImageSynchronizeInvoker 从SWR服务同步镜像列表
func (c *HssClient) RunImageSynchronizeInvoker(request *model.RunImageSynchronizeRequest) *RunImageSynchronizeInvoker {
	requestDef := GenReqDefForRunImageSynchronize()
	return &RunImageSynchronizeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetRaspSwitch 开启/关闭动态网页防篡改防护
//
// 开启/关闭动态网页防篡改防护，下发/清空动态网页防篡改策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) SetRaspSwitch(request *model.SetRaspSwitchRequest) (*model.SetRaspSwitchResponse, error) {
	requestDef := GenReqDefForSetRaspSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetRaspSwitchResponse), nil
	}
}

// SetRaspSwitchInvoker 开启/关闭动态网页防篡改防护
func (c *HssClient) SetRaspSwitchInvoker(request *model.SetRaspSwitchRequest) *SetRaspSwitchInvoker {
	requestDef := GenReqDefForSetRaspSwitch()
	return &SetRaspSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetTwoFactorLoginConfig 设置双因子登录配置
//
// 设置双因子登录配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) SetTwoFactorLoginConfig(request *model.SetTwoFactorLoginConfigRequest) (*model.SetTwoFactorLoginConfigResponse, error) {
	requestDef := GenReqDefForSetTwoFactorLoginConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetTwoFactorLoginConfigResponse), nil
	}
}

// SetTwoFactorLoginConfigInvoker 设置双因子登录配置
func (c *HssClient) SetTwoFactorLoginConfigInvoker(request *model.SetTwoFactorLoginConfigRequest) *SetTwoFactorLoginConfigInvoker {
	requestDef := GenReqDefForSetTwoFactorLoginConfig()
	return &SetTwoFactorLoginConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetWtpProtectionStatusInfo 开启关闭网页防篡改防护
//
// 开启/关闭网页防篡改功能防护，下发/清空网页防篡改策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) SetWtpProtectionStatusInfo(request *model.SetWtpProtectionStatusInfoRequest) (*model.SetWtpProtectionStatusInfoResponse, error) {
	requestDef := GenReqDefForSetWtpProtectionStatusInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetWtpProtectionStatusInfoResponse), nil
	}
}

// SetWtpProtectionStatusInfoInvoker 开启关闭网页防篡改防护
func (c *HssClient) SetWtpProtectionStatusInfoInvoker(request *model.SetWtpProtectionStatusInfoRequest) *SetWtpProtectionStatusInfoInvoker {
	requestDef := GenReqDefForSetWtpProtectionStatusInfo()
	return &SetWtpProtectionStatusInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAgentDaemonsetDeployTemplate 获取部署模板
//
// 获取部署模板，在安装Daemonset的时候提供选择
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowAgentDaemonsetDeployTemplate(request *model.ShowAgentDaemonsetDeployTemplateRequest) (*model.ShowAgentDaemonsetDeployTemplateResponse, error) {
	requestDef := GenReqDefForShowAgentDaemonsetDeployTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAgentDaemonsetDeployTemplateResponse), nil
	}
}

// ShowAgentDaemonsetDeployTemplateInvoker 获取部署模板
func (c *HssClient) ShowAgentDaemonsetDeployTemplateInvoker(request *model.ShowAgentDaemonsetDeployTemplateRequest) *ShowAgentDaemonsetDeployTemplateInvoker {
	requestDef := GenReqDefForShowAgentDaemonsetDeployTemplate()
	return &ShowAgentDaemonsetDeployTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAgentDaemonsetDetailInfo 获取集群daemonset信息
//
// 获取集群daemonset信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowAgentDaemonsetDetailInfo(request *model.ShowAgentDaemonsetDetailInfoRequest) (*model.ShowAgentDaemonsetDetailInfoResponse, error) {
	requestDef := GenReqDefForShowAgentDaemonsetDetailInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAgentDaemonsetDetailInfoResponse), nil
	}
}

// ShowAgentDaemonsetDetailInfoInvoker 获取集群daemonset信息
func (c *HssClient) ShowAgentDaemonsetDetailInfoInvoker(request *model.ShowAgentDaemonsetDetailInfoRequest) *ShowAgentDaemonsetDetailInfoInvoker {
	requestDef := GenReqDefForShowAgentDaemonsetDetailInfo()
	return &ShowAgentDaemonsetDetailInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAgentStatisticsStatus 资产管理-概览-资产状态-主机Agent状态
//
// 资产管理-概览-资产状态-主机Agent状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowAgentStatisticsStatus(request *model.ShowAgentStatisticsStatusRequest) (*model.ShowAgentStatisticsStatusResponse, error) {
	requestDef := GenReqDefForShowAgentStatisticsStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAgentStatisticsStatusResponse), nil
	}
}

// ShowAgentStatisticsStatusInvoker 资产管理-概览-资产状态-主机Agent状态
func (c *HssClient) ShowAgentStatisticsStatusInvoker(request *model.ShowAgentStatisticsStatusRequest) *ShowAgentStatisticsStatusInvoker {
	requestDef := GenReqDefForShowAgentStatisticsStatus()
	return &ShowAgentStatisticsStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAntivirusPayPerScanStatus 查询“病毒查杀按次计费”开关状态
//
// 查询“病毒查杀按次计费”开关状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowAntivirusPayPerScanStatus(request *model.ShowAntivirusPayPerScanStatusRequest) (*model.ShowAntivirusPayPerScanStatusResponse, error) {
	requestDef := GenReqDefForShowAntivirusPayPerScanStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAntivirusPayPerScanStatusResponse), nil
	}
}

// ShowAntivirusPayPerScanStatusInvoker 查询“病毒查杀按次计费”开关状态
func (c *HssClient) ShowAntivirusPayPerScanStatusInvoker(request *model.ShowAntivirusPayPerScanStatusRequest) *ShowAntivirusPayPerScanStatusInvoker {
	requestDef := GenReqDefForShowAntivirusPayPerScanStatus()
	return &ShowAntivirusPayPerScanStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAntivirusStatistic 查询病毒查杀统计信息
//
// 查询病毒查杀统计信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowAntivirusStatistic(request *model.ShowAntivirusStatisticRequest) (*model.ShowAntivirusStatisticResponse, error) {
	requestDef := GenReqDefForShowAntivirusStatistic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAntivirusStatisticResponse), nil
	}
}

// ShowAntivirusStatisticInvoker 查询病毒查杀统计信息
func (c *HssClient) ShowAntivirusStatisticInvoker(request *model.ShowAntivirusStatisticRequest) *ShowAntivirusStatisticInvoker {
	requestDef := GenReqDefForShowAntivirusStatistic()
	return &ShowAntivirusStatisticInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAppRaspSwitchStatus 查询应用防护开启状态
//
// 查询应用防护开启状态：查询单台服务器的应用防护功能开启状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowAppRaspSwitchStatus(request *model.ShowAppRaspSwitchStatusRequest) (*model.ShowAppRaspSwitchStatusResponse, error) {
	requestDef := GenReqDefForShowAppRaspSwitchStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppRaspSwitchStatusResponse), nil
	}
}

// ShowAppRaspSwitchStatusInvoker 查询应用防护开启状态
func (c *HssClient) ShowAppRaspSwitchStatusInvoker(request *model.ShowAppRaspSwitchStatusRequest) *ShowAppRaspSwitchStatusInvoker {
	requestDef := GenReqDefForShowAppRaspSwitchStatus()
	return &ShowAppRaspSwitchStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAssetStatistic 统计资产信息，账号、端口、进程等
//
// 资产统计信息，账号、端口、进程等
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowAssetStatistic(request *model.ShowAssetStatisticRequest) (*model.ShowAssetStatisticResponse, error) {
	requestDef := GenReqDefForShowAssetStatistic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAssetStatisticResponse), nil
	}
}

// ShowAssetStatisticInvoker 统计资产信息，账号、端口、进程等
func (c *HssClient) ShowAssetStatisticInvoker(request *model.ShowAssetStatisticRequest) *ShowAssetStatisticInvoker {
	requestDef := GenReqDefForShowAssetStatistic()
	return &ShowAssetStatisticInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBackupPolicyInfo 查询HSS存储库绑定的备份策略信息
//
// 查询HSS存储库绑定的备份策略信息,确保已经购买了勒索防护存储库，可以从cbr云备份服务进行验证，确保已经存在HSS_projectid命名的存储库已经购买
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowBackupPolicyInfo(request *model.ShowBackupPolicyInfoRequest) (*model.ShowBackupPolicyInfoResponse, error) {
	requestDef := GenReqDefForShowBackupPolicyInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBackupPolicyInfoResponse), nil
	}
}

// ShowBackupPolicyInfoInvoker 查询HSS存储库绑定的备份策略信息
func (c *HssClient) ShowBackupPolicyInfoInvoker(request *model.ShowBackupPolicyInfoRequest) *ShowBackupPolicyInfoInvoker {
	requestDef := GenReqDefForShowBackupPolicyInfo()
	return &ShowBackupPolicyInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBaselineScanStatus 查询基线扫描手动检测结果
//
// 查询基线扫描手动检测结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowBaselineScanStatus(request *model.ShowBaselineScanStatusRequest) (*model.ShowBaselineScanStatusResponse, error) {
	requestDef := GenReqDefForShowBaselineScanStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBaselineScanStatusResponse), nil
	}
}

// ShowBaselineScanStatusInvoker 查询基线扫描手动检测结果
func (c *HssClient) ShowBaselineScanStatusInvoker(request *model.ShowBaselineScanStatusRequest) *ShowBaselineScanStatusInvoker {
	requestDef := GenReqDefForShowBaselineScanStatus()
	return &ShowBaselineScanStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBaselineWhiteList 查询基线白名单
//
// 查询基线白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowBaselineWhiteList(request *model.ShowBaselineWhiteListRequest) (*model.ShowBaselineWhiteListResponse, error) {
	requestDef := GenReqDefForShowBaselineWhiteList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBaselineWhiteListResponse), nil
	}
}

// ShowBaselineWhiteListInvoker 查询基线白名单
func (c *HssClient) ShowBaselineWhiteListInvoker(request *model.ShowBaselineWhiteListRequest) *ShowBaselineWhiteListInvoker {
	requestDef := GenReqDefForShowBaselineWhiteList()
	return &ShowBaselineWhiteListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCheckRuleDetail 查询配置检查项检测报告
//
// 查询配置检查项检测报告
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowCheckRuleDetail(request *model.ShowCheckRuleDetailRequest) (*model.ShowCheckRuleDetailResponse, error) {
	requestDef := GenReqDefForShowCheckRuleDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCheckRuleDetailResponse), nil
	}
}

// ShowCheckRuleDetailInvoker 查询配置检查项检测报告
func (c *HssClient) ShowCheckRuleDetailInvoker(request *model.ShowCheckRuleDetailRequest) *ShowCheckRuleDetailInvoker {
	requestDef := GenReqDefForShowCheckRuleDetail()
	return &ShowCheckRuleDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowClusterAssetStatistics 查询集群资产统计数量
//
// 查询集群资产统计数量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowClusterAssetStatistics(request *model.ShowClusterAssetStatisticsRequest) (*model.ShowClusterAssetStatisticsResponse, error) {
	requestDef := GenReqDefForShowClusterAssetStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusterAssetStatisticsResponse), nil
	}
}

// ShowClusterAssetStatisticsInvoker 查询集群资产统计数量
func (c *HssClient) ShowClusterAssetStatisticsInvoker(request *model.ShowClusterAssetStatisticsRequest) *ShowClusterAssetStatisticsInvoker {
	requestDef := GenReqDefForShowClusterAssetStatistics()
	return &ShowClusterAssetStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowClusterProtectPolicyTemplate 查询集群组件防护策略模板
//
// 查询集群防护策略模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowClusterProtectPolicyTemplate(request *model.ShowClusterProtectPolicyTemplateRequest) (*model.ShowClusterProtectPolicyTemplateResponse, error) {
	requestDef := GenReqDefForShowClusterProtectPolicyTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusterProtectPolicyTemplateResponse), nil
	}
}

// ShowClusterProtectPolicyTemplateInvoker 查询集群组件防护策略模板
func (c *HssClient) ShowClusterProtectPolicyTemplateInvoker(request *model.ShowClusterProtectPolicyTemplateRequest) *ShowClusterProtectPolicyTemplateInvoker {
	requestDef := GenReqDefForShowClusterProtectPolicyTemplate()
	return &ShowClusterProtectPolicyTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCommonPort 呈现某一端口详细信息
//
// 呈现某一端口详细信息，如本地端口：80                      类型：TCP 危险程度：正常 端口描述：常用于SSH(SecureShell)-远程登录协议，用于安全登录文件传输（SCP，SFTP）及端口重新定向。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowCommonPort(request *model.ShowCommonPortRequest) (*model.ShowCommonPortResponse, error) {
	requestDef := GenReqDefForShowCommonPort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCommonPortResponse), nil
	}
}

// ShowCommonPortInvoker 呈现某一端口详细信息
func (c *HssClient) ShowCommonPortInvoker(request *model.ShowCommonPortRequest) *ShowCommonPortInvoker {
	requestDef := GenReqDefForShowCommonPort()
	return &ShowCommonPortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowContainerNetworkInfo 查询容器集群网络的网络信息
//
// 查询容器集群网络的网络信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowContainerNetworkInfo(request *model.ShowContainerNetworkInfoRequest) (*model.ShowContainerNetworkInfoResponse, error) {
	requestDef := GenReqDefForShowContainerNetworkInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowContainerNetworkInfoResponse), nil
	}
}

// ShowContainerNetworkInfoInvoker 查询容器集群网络的网络信息
func (c *HssClient) ShowContainerNetworkInfoInvoker(request *model.ShowContainerNetworkInfoRequest) *ShowContainerNetworkInfoInvoker {
	requestDef := GenReqDefForShowContainerNetworkInfo()
	return &ShowContainerNetworkInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowContainerNodeStatistics 查询容器节点防护总览数据
//
// 查询容器节点防护总览数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowContainerNodeStatistics(request *model.ShowContainerNodeStatisticsRequest) (*model.ShowContainerNodeStatisticsResponse, error) {
	requestDef := GenReqDefForShowContainerNodeStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowContainerNodeStatisticsResponse), nil
	}
}

// ShowContainerNodeStatisticsInvoker 查询容器节点防护总览数据
func (c *HssClient) ShowContainerNodeStatisticsInvoker(request *model.ShowContainerNodeStatisticsRequest) *ShowContainerNodeStatisticsInvoker {
	requestDef := GenReqDefForShowContainerNodeStatistics()
	return &ShowContainerNodeStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowContainerProtectionStatus 资产管理-概览-资产状态-容器节点防护状态
//
// 资产管理-概览-资产状态-容器节点防护状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowContainerProtectionStatus(request *model.ShowContainerProtectionStatusRequest) (*model.ShowContainerProtectionStatusResponse, error) {
	requestDef := GenReqDefForShowContainerProtectionStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowContainerProtectionStatusResponse), nil
	}
}

// ShowContainerProtectionStatusInvoker 资产管理-概览-资产状态-容器节点防护状态
func (c *HssClient) ShowContainerProtectionStatusInvoker(request *model.ShowContainerProtectionStatusRequest) *ShowContainerProtectionStatusInvoker {
	requestDef := GenReqDefForShowContainerProtectionStatus()
	return &ShowContainerProtectionStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDecoyPortPolicyDetails 查看动态端口蜜罐策略详情
//
// 查看动态端口蜜罐策略详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowDecoyPortPolicyDetails(request *model.ShowDecoyPortPolicyDetailsRequest) (*model.ShowDecoyPortPolicyDetailsResponse, error) {
	requestDef := GenReqDefForShowDecoyPortPolicyDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDecoyPortPolicyDetailsResponse), nil
	}
}

// ShowDecoyPortPolicyDetailsInvoker 查看动态端口蜜罐策略详情
func (c *HssClient) ShowDecoyPortPolicyDetailsInvoker(request *model.ShowDecoyPortPolicyDetailsRequest) *ShowDecoyPortPolicyDetailsInvoker {
	requestDef := GenReqDefForShowDecoyPortPolicyDetails()
	return &ShowDecoyPortPolicyDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHostAssetManualCollectStatus 查询单主机资产指纹采集状态
//
// 查询单主机资产指纹采集状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowHostAssetManualCollectStatus(request *model.ShowHostAssetManualCollectStatusRequest) (*model.ShowHostAssetManualCollectStatusResponse, error) {
	requestDef := GenReqDefForShowHostAssetManualCollectStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHostAssetManualCollectStatusResponse), nil
	}
}

// ShowHostAssetManualCollectStatusInvoker 查询单主机资产指纹采集状态
func (c *HssClient) ShowHostAssetManualCollectStatusInvoker(request *model.ShowHostAssetManualCollectStatusRequest) *ShowHostAssetManualCollectStatusInvoker {
	requestDef := GenReqDefForShowHostAssetManualCollectStatus()
	return &ShowHostAssetManualCollectStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHostProtectionStatus 资产管理-概览-资产状态-Agent状态
//
// 资产管理-概览-资产状态-Agent状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowHostProtectionStatus(request *model.ShowHostProtectionStatusRequest) (*model.ShowHostProtectionStatusResponse, error) {
	requestDef := GenReqDefForShowHostProtectionStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHostProtectionStatusResponse), nil
	}
}

// ShowHostProtectionStatusInvoker 资产管理-概览-资产状态-Agent状态
func (c *HssClient) ShowHostProtectionStatusInvoker(request *model.ShowHostProtectionStatusRequest) *ShowHostProtectionStatusInvoker {
	requestDef := GenReqDefForShowHostProtectionStatus()
	return &ShowHostProtectionStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowImageAssetStatistics 容器资产-镜像统计
//
// 容器资产-镜像统计
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowImageAssetStatistics(request *model.ShowImageAssetStatisticsRequest) (*model.ShowImageAssetStatisticsResponse, error) {
	requestDef := GenReqDefForShowImageAssetStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowImageAssetStatisticsResponse), nil
	}
}

// ShowImageAssetStatisticsInvoker 容器资产-镜像统计
func (c *HssClient) ShowImageAssetStatisticsInvoker(request *model.ShowImageAssetStatisticsRequest) *ShowImageAssetStatisticsInvoker {
	requestDef := GenReqDefForShowImageAssetStatistics()
	return &ShowImageAssetStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowImageCheckRuleDetail 查询镜像配置检查项检测报告
//
// 查询镜像配置检查项检测报告
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowImageCheckRuleDetail(request *model.ShowImageCheckRuleDetailRequest) (*model.ShowImageCheckRuleDetailResponse, error) {
	requestDef := GenReqDefForShowImageCheckRuleDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowImageCheckRuleDetailResponse), nil
	}
}

// ShowImageCheckRuleDetailInvoker 查询镜像配置检查项检测报告
func (c *HssClient) ShowImageCheckRuleDetailInvoker(request *model.ShowImageCheckRuleDetailRequest) *ShowImageCheckRuleDetailInvoker {
	requestDef := GenReqDefForShowImageCheckRuleDetail()
	return &ShowImageCheckRuleDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowK8sContainerDetail 查询容器详细信息
//
// 查询容器详细信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowK8sContainerDetail(request *model.ShowK8sContainerDetailRequest) (*model.ShowK8sContainerDetailResponse, error) {
	requestDef := GenReqDefForShowK8sContainerDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowK8sContainerDetailResponse), nil
	}
}

// ShowK8sContainerDetailInvoker 查询容器详细信息
func (c *HssClient) ShowK8sContainerDetailInvoker(request *model.ShowK8sContainerDetailRequest) *ShowK8sContainerDetailInvoker {
	requestDef := GenReqDefForShowK8sContainerDetail()
	return &ShowK8sContainerDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowK8sPodDetail 查询pod详细信息
//
// 查询pod详细信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowK8sPodDetail(request *model.ShowK8sPodDetailRequest) (*model.ShowK8sPodDetailResponse, error) {
	requestDef := GenReqDefForShowK8sPodDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowK8sPodDetailResponse), nil
	}
}

// ShowK8sPodDetailInvoker 查询pod详细信息
func (c *HssClient) ShowK8sPodDetailInvoker(request *model.ShowK8sPodDetailRequest) *ShowK8sPodDetailInvoker {
	requestDef := GenReqDefForShowK8sPodDetail()
	return &ShowK8sPodDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowKubernetesEndpointInfo 查询容器Kubernetes端点详情
//
// 查询容器Kubernetes端点详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowKubernetesEndpointInfo(request *model.ShowKubernetesEndpointInfoRequest) (*model.ShowKubernetesEndpointInfoResponse, error) {
	requestDef := GenReqDefForShowKubernetesEndpointInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKubernetesEndpointInfoResponse), nil
	}
}

// ShowKubernetesEndpointInfoInvoker 查询容器Kubernetes端点详情
func (c *HssClient) ShowKubernetesEndpointInfoInvoker(request *model.ShowKubernetesEndpointInfoRequest) *ShowKubernetesEndpointInfoInvoker {
	requestDef := GenReqDefForShowKubernetesEndpointInfo()
	return &ShowKubernetesEndpointInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowKubernetesServiceInfo 查询容器Kubernetes服务详情
//
// 查询容器Kubernetes服务详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowKubernetesServiceInfo(request *model.ShowKubernetesServiceInfoRequest) (*model.ShowKubernetesServiceInfoResponse, error) {
	requestDef := GenReqDefForShowKubernetesServiceInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKubernetesServiceInfoResponse), nil
	}
}

// ShowKubernetesServiceInfoInvoker 查询容器Kubernetes服务详情
func (c *HssClient) ShowKubernetesServiceInfoInvoker(request *model.ShowKubernetesServiceInfoRequest) *ShowKubernetesServiceInfoInvoker {
	requestDef := GenReqDefForShowKubernetesServiceInfo()
	return &ShowKubernetesServiceInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLatestExportTaskByType 查询导出任务信息-按查询条件
//
// 查询导出任务信息-按查询条件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowLatestExportTaskByType(request *model.ShowLatestExportTaskByTypeRequest) (*model.ShowLatestExportTaskByTypeResponse, error) {
	requestDef := GenReqDefForShowLatestExportTaskByType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLatestExportTaskByTypeResponse), nil
	}
}

// ShowLatestExportTaskByTypeInvoker 查询导出任务信息-按查询条件
func (c *HssClient) ShowLatestExportTaskByTypeInvoker(request *model.ShowLatestExportTaskByTypeRequest) *ShowLatestExportTaskByTypeInvoker {
	requestDef := GenReqDefForShowLatestExportTaskByType()
	return &ShowLatestExportTaskByTypeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMultiCloudClusterImageCommand 获取多云集群的上传镜像指令
//
// 获取多云集群的上传镜像指令
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowMultiCloudClusterImageCommand(request *model.ShowMultiCloudClusterImageCommandRequest) (*model.ShowMultiCloudClusterImageCommandResponse, error) {
	requestDef := GenReqDefForShowMultiCloudClusterImageCommand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMultiCloudClusterImageCommandResponse), nil
	}
}

// ShowMultiCloudClusterImageCommandInvoker 获取多云集群的上传镜像指令
func (c *HssClient) ShowMultiCloudClusterImageCommandInvoker(request *model.ShowMultiCloudClusterImageCommandRequest) *ShowMultiCloudClusterImageCommandInvoker {
	requestDef := GenReqDefForShowMultiCloudClusterImageCommand()
	return &ShowMultiCloudClusterImageCommandInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMultiCloudClusterProxyScript 获取多云集群的代理安装脚本
//
// 获取多云集群的代理安装脚本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowMultiCloudClusterProxyScript(request *model.ShowMultiCloudClusterProxyScriptRequest) (*model.ShowMultiCloudClusterProxyScriptResponse, error) {
	requestDef := GenReqDefForShowMultiCloudClusterProxyScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMultiCloudClusterProxyScriptResponse), nil
	}
}

// ShowMultiCloudClusterProxyScriptInvoker 获取多云集群的代理安装脚本
func (c *HssClient) ShowMultiCloudClusterProxyScriptInvoker(request *model.ShowMultiCloudClusterProxyScriptRequest) *ShowMultiCloudClusterProxyScriptInvoker {
	requestDef := GenReqDefForShowMultiCloudClusterProxyScript()
	return &ShowMultiCloudClusterProxyScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNetworkStatistics 集群网络策略总览
//
// 集群网络策略总览
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowNetworkStatistics(request *model.ShowNetworkStatisticsRequest) (*model.ShowNetworkStatisticsResponse, error) {
	requestDef := GenReqDefForShowNetworkStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNetworkStatisticsResponse), nil
	}
}

// ShowNetworkStatisticsInvoker 集群网络策略总览
func (c *HssClient) ShowNetworkStatisticsInvoker(request *model.ShowNetworkStatisticsRequest) *ShowNetworkStatisticsInvoker {
	requestDef := GenReqDefForShowNetworkStatistics()
	return &ShowNetworkStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOsStatisticsInfo 资产管理-概览-资产状态-操作系统统计信息
//
// 资产管理-概览-资产状态-操作系统统计信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowOsStatisticsInfo(request *model.ShowOsStatisticsInfoRequest) (*model.ShowOsStatisticsInfoResponse, error) {
	requestDef := GenReqDefForShowOsStatisticsInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOsStatisticsInfoResponse), nil
	}
}

// ShowOsStatisticsInfoInvoker 资产管理-概览-资产状态-操作系统统计信息
func (c *HssClient) ShowOsStatisticsInfoInvoker(request *model.ShowOsStatisticsInfoRequest) *ShowOsStatisticsInfoInvoker {
	requestDef := GenReqDefForShowOsStatisticsInfo()
	return &ShowOsStatisticsInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPageNotices 获取页面通知信息
//
// 获取页面通知信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowPageNotices(request *model.ShowPageNoticesRequest) (*model.ShowPageNoticesResponse, error) {
	requestDef := GenReqDefForShowPageNotices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPageNoticesResponse), nil
	}
}

// ShowPageNoticesInvoker 获取页面通知信息
func (c *HssClient) ShowPageNoticesInvoker(request *model.ShowPageNoticesRequest) *ShowPageNoticesInvoker {
	requestDef := GenReqDefForShowPageNotices()
	return &ShowPageNoticesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProductdataOfferingInfos 查询产商品信息
//
// 查询产商品信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowProductdataOfferingInfos(request *model.ShowProductdataOfferingInfosRequest) (*model.ShowProductdataOfferingInfosResponse, error) {
	requestDef := GenReqDefForShowProductdataOfferingInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProductdataOfferingInfosResponse), nil
	}
}

// ShowProductdataOfferingInfosInvoker 查询产商品信息
func (c *HssClient) ShowProductdataOfferingInfosInvoker(request *model.ShowProductdataOfferingInfosRequest) *ShowProductdataOfferingInfosInvoker {
	requestDef := GenReqDefForShowProductdataOfferingInfos()
	return &ShowProductdataOfferingInfosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQuotaStatisticsInfo 资产管理-概览-资产状态-防护配额统计信息
//
// 资产管理-概览-资产状态-防护配额统计信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowQuotaStatisticsInfo(request *model.ShowQuotaStatisticsInfoRequest) (*model.ShowQuotaStatisticsInfoResponse, error) {
	requestDef := GenReqDefForShowQuotaStatisticsInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotaStatisticsInfoResponse), nil
	}
}

// ShowQuotaStatisticsInfoInvoker 资产管理-概览-资产状态-防护配额统计信息
func (c *HssClient) ShowQuotaStatisticsInfoInvoker(request *model.ShowQuotaStatisticsInfoRequest) *ShowQuotaStatisticsInfoInvoker {
	requestDef := GenReqDefForShowQuotaStatisticsInfo()
	return &ShowQuotaStatisticsInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRaspPolicyDetail 查询防护策略详情
//
// 查询防护策略详情：查询防护策略配置的相关检测规则信息，包含14种检测规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowRaspPolicyDetail(request *model.ShowRaspPolicyDetailRequest) (*model.ShowRaspPolicyDetailResponse, error) {
	requestDef := GenReqDefForShowRaspPolicyDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRaspPolicyDetailResponse), nil
	}
}

// ShowRaspPolicyDetailInvoker 查询防护策略详情
func (c *HssClient) ShowRaspPolicyDetailInvoker(request *model.ShowRaspPolicyDetailRequest) *ShowRaspPolicyDetailInvoker {
	requestDef := GenReqDefForShowRaspPolicyDetail()
	return &ShowRaspPolicyDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRaspProtectStatistics 防护数据统计
//
// 防护数据统计：统计已添加防护服务器的数量以及近七天微服务RASP攻击数量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowRaspProtectStatistics(request *model.ShowRaspProtectStatisticsRequest) (*model.ShowRaspProtectStatisticsResponse, error) {
	requestDef := GenReqDefForShowRaspProtectStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRaspProtectStatisticsResponse), nil
	}
}

// ShowRaspProtectStatisticsInvoker 防护数据统计
func (c *HssClient) ShowRaspProtectStatisticsInvoker(request *model.ShowRaspProtectStatisticsRequest) *ShowRaspProtectStatisticsInvoker {
	requestDef := GenReqDefForShowRaspProtectStatistics()
	return &ShowRaspProtectStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRaspServerDetail 查询防护服务器java应用详情
//
// 查询防护服务器java应用详情：查询防护服务器的java应用状态列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowRaspServerDetail(request *model.ShowRaspServerDetailRequest) (*model.ShowRaspServerDetailResponse, error) {
	requestDef := GenReqDefForShowRaspServerDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRaspServerDetailResponse), nil
	}
}

// ShowRaspServerDetailInvoker 查询防护服务器java应用详情
func (c *HssClient) ShowRaspServerDetailInvoker(request *model.ShowRaspServerDetailRequest) *ShowRaspServerDetailInvoker {
	requestDef := GenReqDefForShowRaspServerDetail()
	return &ShowRaspServerDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceQuotas 查询配额信息
//
// 查询配额信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowResourceQuotas(request *model.ShowResourceQuotasRequest) (*model.ShowResourceQuotasResponse, error) {
	requestDef := GenReqDefForShowResourceQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceQuotasResponse), nil
	}
}

// ShowResourceQuotasInvoker 查询配额信息
func (c *HssClient) ShowResourceQuotasInvoker(request *model.ShowResourceQuotasRequest) *ShowResourceQuotasInvoker {
	requestDef := GenReqDefForShowResourceQuotas()
	return &ShowResourceQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRiskConfigDetail 查询指定安全配置项的检查结果
//
// 查询指定安全配置项的检查结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowRiskConfigDetail(request *model.ShowRiskConfigDetailRequest) (*model.ShowRiskConfigDetailResponse, error) {
	requestDef := GenReqDefForShowRiskConfigDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRiskConfigDetailResponse), nil
	}
}

// ShowRiskConfigDetailInvoker 查询指定安全配置项的检查结果
func (c *HssClient) ShowRiskConfigDetailInvoker(request *model.ShowRiskConfigDetailRequest) *ShowRiskConfigDetailInvoker {
	requestDef := GenReqDefForShowRiskConfigDetail()
	return &ShowRiskConfigDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSingleBackupPolicyInfo 查询单个备份策略信息
//
// 查询单个备份策略信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowSingleBackupPolicyInfo(request *model.ShowSingleBackupPolicyInfoRequest) (*model.ShowSingleBackupPolicyInfoResponse, error) {
	requestDef := GenReqDefForShowSingleBackupPolicyInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSingleBackupPolicyInfoResponse), nil
	}
}

// ShowSingleBackupPolicyInfoInvoker 查询单个备份策略信息
func (c *HssClient) ShowSingleBackupPolicyInfoInvoker(request *model.ShowSingleBackupPolicyInfoRequest) *ShowSingleBackupPolicyInfoInvoker {
	requestDef := GenReqDefForShowSingleBackupPolicyInfo()
	return &ShowSingleBackupPolicyInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVulScanPolicy 查询漏洞扫描策略
//
// 查询漏洞扫描策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowVulScanPolicy(request *model.ShowVulScanPolicyRequest) (*model.ShowVulScanPolicyResponse, error) {
	requestDef := GenReqDefForShowVulScanPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVulScanPolicyResponse), nil
	}
}

// ShowVulScanPolicyInvoker 查询漏洞扫描策略
func (c *HssClient) ShowVulScanPolicyInvoker(request *model.ShowVulScanPolicyRequest) *ShowVulScanPolicyInvoker {
	requestDef := GenReqDefForShowVulScanPolicy()
	return &ShowVulScanPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVulStatics 查询漏洞管理统计数据
//
// 查询漏洞管理统计数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowVulStatics(request *model.ShowVulStaticsRequest) (*model.ShowVulStaticsResponse, error) {
	requestDef := GenReqDefForShowVulStatics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVulStaticsResponse), nil
	}
}

// ShowVulStaticsInvoker 查询漏洞管理统计数据
func (c *HssClient) ShowVulStaticsInvoker(request *model.ShowVulStaticsRequest) *ShowVulStaticsInvoker {
	requestDef := GenReqDefForShowVulStatics()
	return &ShowVulStaticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartProtection 开启勒索病毒防护
//
// 开启勒索病毒防护,请保证该region有cbr云备份服务，勒索服务与云备份服务有关联关系
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) StartProtection(request *model.StartProtectionRequest) (*model.StartProtectionResponse, error) {
	requestDef := GenReqDefForStartProtection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartProtectionResponse), nil
	}
}

// StartProtectionInvoker 开启勒索病毒防护
func (c *HssClient) StartProtectionInvoker(request *model.StartProtectionRequest) *StartProtectionInvoker {
	requestDef := GenReqDefForStartProtection()
	return &StartProtectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopProtection 关闭勒索病毒防护
//
// 关闭勒索病毒防护
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) StopProtection(request *model.StopProtectionRequest) (*model.StopProtectionResponse, error) {
	requestDef := GenReqDefForStopProtection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopProtectionResponse), nil
	}
}

// StopProtectionInvoker 关闭勒索病毒防护
func (c *HssClient) StopProtectionInvoker(request *model.StopProtectionRequest) *StopProtectionInvoker {
	requestDef := GenReqDefForStopProtection()
	return &StopProtectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchAntivirusTask 取消扫描任务
//
// 取消扫描任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) SwitchAntivirusTask(request *model.SwitchAntivirusTaskRequest) (*model.SwitchAntivirusTaskResponse, error) {
	requestDef := GenReqDefForSwitchAntivirusTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchAntivirusTaskResponse), nil
	}
}

// SwitchAntivirusTaskInvoker 取消扫描任务
func (c *HssClient) SwitchAntivirusTaskInvoker(request *model.SwitchAntivirusTaskRequest) *SwitchAntivirusTaskInvoker {
	requestDef := GenReqDefForSwitchAntivirusTask()
	return &SwitchAntivirusTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchClusterProtectionMode 操作集群防护模式
//
// 操作集群防护模式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) SwitchClusterProtectionMode(request *model.SwitchClusterProtectionModeRequest) (*model.SwitchClusterProtectionModeResponse, error) {
	requestDef := GenReqDefForSwitchClusterProtectionMode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchClusterProtectionModeResponse), nil
	}
}

// SwitchClusterProtectionModeInvoker 操作集群防护模式
func (c *HssClient) SwitchClusterProtectionModeInvoker(request *model.SwitchClusterProtectionModeRequest) *SwitchClusterProtectionModeInvoker {
	requestDef := GenReqDefForSwitchClusterProtectionMode()
	return &SwitchClusterProtectionModeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchContainerProtectStatus 切换防护状态
//
// 切换防护状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) SwitchContainerProtectStatus(request *model.SwitchContainerProtectStatusRequest) (*model.SwitchContainerProtectStatusResponse, error) {
	requestDef := GenReqDefForSwitchContainerProtectStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchContainerProtectStatusResponse), nil
	}
}

// SwitchContainerProtectStatusInvoker 切换防护状态
func (c *HssClient) SwitchContainerProtectStatusInvoker(request *model.SwitchContainerProtectStatusRequest) *SwitchContainerProtectStatusInvoker {
	requestDef := GenReqDefForSwitchContainerProtectStatus()
	return &SwitchContainerProtectStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchDecoyPortHostPolicy 切换主机动态端口蜜罐策略
//
// 切换主机动态端口蜜罐策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) SwitchDecoyPortHostPolicy(request *model.SwitchDecoyPortHostPolicyRequest) (*model.SwitchDecoyPortHostPolicyResponse, error) {
	requestDef := GenReqDefForSwitchDecoyPortHostPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchDecoyPortHostPolicyResponse), nil
	}
}

// SwitchDecoyPortHostPolicyInvoker 切换主机动态端口蜜罐策略
func (c *HssClient) SwitchDecoyPortHostPolicyInvoker(request *model.SwitchDecoyPortHostPolicyRequest) *SwitchDecoyPortHostPolicyInvoker {
	requestDef := GenReqDefForSwitchDecoyPortHostPolicy()
	return &SwitchDecoyPortHostPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchHostsProtectStatus 切换防护状态
//
// 切换防护状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) SwitchHostsProtectStatus(request *model.SwitchHostsProtectStatusRequest) (*model.SwitchHostsProtectStatusResponse, error) {
	requestDef := GenReqDefForSwitchHostsProtectStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchHostsProtectStatusResponse), nil
	}
}

// SwitchHostsProtectStatusInvoker 切换防护状态
func (c *HssClient) SwitchHostsProtectStatusInvoker(request *model.SwitchHostsProtectStatusRequest) *SwitchHostsProtectStatusInvoker {
	requestDef := GenReqDefForSwitchHostsProtectStatus()
	return &SwitchHostsProtectStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchRasp 开启/关闭应用防护，更新防护端口
//
// 开启/关闭应用防护，选择开启的防护策略，选择需要防护的服务器，下发防护策略，可传入端口号更新防护端口，关闭防护则清空策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) SwitchRasp(request *model.SwitchRaspRequest) (*model.SwitchRaspResponse, error) {
	requestDef := GenReqDefForSwitchRasp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchRaspResponse), nil
	}
}

// SwitchRaspInvoker 开启/关闭应用防护，更新防护端口
func (c *HssClient) SwitchRaspInvoker(request *model.SwitchRaspRequest) *SwitchRaspInvoker {
	requestDef := GenReqDefForSwitchRasp()
	return &SwitchRaspInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SyncClusterList 同步容器集群最新数据
//
// 同步容器集群最新数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) SyncClusterList(request *model.SyncClusterListRequest) (*model.SyncClusterListResponse, error) {
	requestDef := GenReqDefForSyncClusterList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SyncClusterListResponse), nil
	}
}

// SyncClusterListInvoker 同步容器集群最新数据
func (c *HssClient) SyncClusterListInvoker(request *model.SyncClusterListRequest) *SyncClusterListInvoker {
	requestDef := GenReqDefForSyncClusterList()
	return &SyncClusterListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SyncClusterProtectionEvent 同步集群防护事件
//
// 同步集群防护事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) SyncClusterProtectionEvent(request *model.SyncClusterProtectionEventRequest) (*model.SyncClusterProtectionEventResponse, error) {
	requestDef := GenReqDefForSyncClusterProtectionEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SyncClusterProtectionEventResponse), nil
	}
}

// SyncClusterProtectionEventInvoker 同步集群防护事件
func (c *HssClient) SyncClusterProtectionEventInvoker(request *model.SyncClusterProtectionEventRequest) *SyncClusterProtectionEventInvoker {
	requestDef := GenReqDefForSyncClusterProtectionEvent()
	return &SyncClusterProtectionEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SyncContainerNetworkNode 同步集群下网络节点最新数据
//
// 同步集群下容器网络策略最新数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) SyncContainerNetworkNode(request *model.SyncContainerNetworkNodeRequest) (*model.SyncContainerNetworkNodeResponse, error) {
	requestDef := GenReqDefForSyncContainerNetworkNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SyncContainerNetworkNodeResponse), nil
	}
}

// SyncContainerNetworkNodeInvoker 同步集群下网络节点最新数据
func (c *HssClient) SyncContainerNetworkNodeInvoker(request *model.SyncContainerNetworkNodeRequest) *SyncContainerNetworkNodeInvoker {
	requestDef := GenReqDefForSyncContainerNetworkNode()
	return &SyncContainerNetworkNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SyncContainerNetworkPolicyList 同步集群下容器网络策略最新数据
//
// 同步集群下容器网络策略最新数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) SyncContainerNetworkPolicyList(request *model.SyncContainerNetworkPolicyListRequest) (*model.SyncContainerNetworkPolicyListResponse, error) {
	requestDef := GenReqDefForSyncContainerNetworkPolicyList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SyncContainerNetworkPolicyListResponse), nil
	}
}

// SyncContainerNetworkPolicyListInvoker 同步集群下容器网络策略最新数据
func (c *HssClient) SyncContainerNetworkPolicyListInvoker(request *model.SyncContainerNetworkPolicyListRequest) *SyncContainerNetworkPolicyListInvoker {
	requestDef := GenReqDefForSyncContainerNetworkPolicyList()
	return &SyncContainerNetworkPolicyListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SyncMultiCloudClusterStatus 同步多云集群的接入状态
//
// 同步多云集群的接入状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) SyncMultiCloudClusterStatus(request *model.SyncMultiCloudClusterStatusRequest) (*model.SyncMultiCloudClusterStatusResponse, error) {
	requestDef := GenReqDefForSyncMultiCloudClusterStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SyncMultiCloudClusterStatusResponse), nil
	}
}

// SyncMultiCloudClusterStatusInvoker 同步多云集群的接入状态
func (c *HssClient) SyncMultiCloudClusterStatusInvoker(request *model.SyncMultiCloudClusterStatusRequest) *SyncMultiCloudClusterStatusInvoker {
	requestDef := GenReqDefForSyncMultiCloudClusterStatus()
	return &SyncMultiCloudClusterStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SyncSecurityGroupPolicies 同步集群下安全组策略最新数据
//
// 同步集群下安全组策略最新数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) SyncSecurityGroupPolicies(request *model.SyncSecurityGroupPoliciesRequest) (*model.SyncSecurityGroupPoliciesResponse, error) {
	requestDef := GenReqDefForSyncSecurityGroupPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SyncSecurityGroupPoliciesResponse), nil
	}
}

// SyncSecurityGroupPoliciesInvoker 同步集群下安全组策略最新数据
func (c *HssClient) SyncSecurityGroupPoliciesInvoker(request *model.SyncSecurityGroupPoliciesRequest) *SyncSecurityGroupPoliciesInvoker {
	requestDef := GenReqDefForSyncSecurityGroupPolicies()
	return &SyncSecurityGroupPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAgentDaemonset 更新集群daemonset
//
// 更新集群daemonset
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) UpdateAgentDaemonset(request *model.UpdateAgentDaemonsetRequest) (*model.UpdateAgentDaemonsetResponse, error) {
	requestDef := GenReqDefForUpdateAgentDaemonset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAgentDaemonsetResponse), nil
	}
}

// UpdateAgentDaemonsetInvoker 更新集群daemonset
func (c *HssClient) UpdateAgentDaemonsetInvoker(request *model.UpdateAgentDaemonsetRequest) *UpdateAgentDaemonsetInvoker {
	requestDef := GenReqDefForUpdateAgentDaemonset()
	return &UpdateAgentDaemonsetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateBackupPolicyInfo 修改存储库绑定的备份策略
//
// 修改存储库绑定的备份策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) UpdateBackupPolicyInfo(request *model.UpdateBackupPolicyInfoRequest) (*model.UpdateBackupPolicyInfoResponse, error) {
	requestDef := GenReqDefForUpdateBackupPolicyInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBackupPolicyInfoResponse), nil
	}
}

// UpdateBackupPolicyInfoInvoker 修改存储库绑定的备份策略
func (c *HssClient) UpdateBackupPolicyInfoInvoker(request *model.UpdateBackupPolicyInfoRequest) *UpdateBackupPolicyInfoInvoker {
	requestDef := GenReqDefForUpdateBackupPolicyInfo()
	return &UpdateBackupPolicyInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateContainerNetworkPolicy 容器集群网络更新配置策略
//
// 容器集群网络更新配置策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) UpdateContainerNetworkPolicy(request *model.UpdateContainerNetworkPolicyRequest) (*model.UpdateContainerNetworkPolicyResponse, error) {
	requestDef := GenReqDefForUpdateContainerNetworkPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateContainerNetworkPolicyResponse), nil
	}
}

// UpdateContainerNetworkPolicyInvoker 容器集群网络更新配置策略
func (c *HssClient) UpdateContainerNetworkPolicyInvoker(request *model.UpdateContainerNetworkPolicyRequest) *UpdateContainerNetworkPolicyInvoker {
	requestDef := GenReqDefForUpdateContainerNetworkPolicy()
	return &UpdateContainerNetworkPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMultiCloudClusters 更新多云集群
//
// 更新多云集群
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) UpdateMultiCloudClusters(request *model.UpdateMultiCloudClustersRequest) (*model.UpdateMultiCloudClustersResponse, error) {
	requestDef := GenReqDefForUpdateMultiCloudClusters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMultiCloudClustersResponse), nil
	}
}

// UpdateMultiCloudClustersInvoker 更新多云集群
func (c *HssClient) UpdateMultiCloudClustersInvoker(request *model.UpdateMultiCloudClustersRequest) *UpdateMultiCloudClustersInvoker {
	requestDef := GenReqDefForUpdateMultiCloudClusters()
	return &UpdateMultiCloudClustersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePolicy 修改防护策略
//
// 修改防护策略：修改防护策略内容，包含策略名称、相关规则开启状态、防护动作以及检测规则配置，同时给使用该策略的服务器下发新的策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) UpdatePolicy(request *model.UpdatePolicyRequest) (*model.UpdatePolicyResponse, error) {
	requestDef := GenReqDefForUpdatePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePolicyResponse), nil
	}
}

// UpdatePolicyInvoker 修改防护策略
func (c *HssClient) UpdatePolicyInvoker(request *model.UpdatePolicyRequest) *UpdatePolicyInvoker {
	requestDef := GenReqDefForUpdatePolicy()
	return &UpdatePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProtectionPolicy 修改勒索防护策略
//
// 修改勒索防护策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) UpdateProtectionPolicy(request *model.UpdateProtectionPolicyRequest) (*model.UpdateProtectionPolicyResponse, error) {
	requestDef := GenReqDefForUpdateProtectionPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProtectionPolicyResponse), nil
	}
}

// UpdateProtectionPolicyInvoker 修改勒索防护策略
func (c *HssClient) UpdateProtectionPolicyInvoker(request *model.UpdateProtectionPolicyRequest) *UpdateProtectionPolicyInvoker {
	requestDef := GenReqDefForUpdateProtectionPolicy()
	return &UpdateProtectionPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSecurityGroupPolicy 更新安全组策略
//
// 更新安全组策略(云原生网络模型)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) UpdateSecurityGroupPolicy(request *model.UpdateSecurityGroupPolicyRequest) (*model.UpdateSecurityGroupPolicyResponse, error) {
	requestDef := GenReqDefForUpdateSecurityGroupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSecurityGroupPolicyResponse), nil
	}
}

// UpdateSecurityGroupPolicyInvoker 更新安全组策略
func (c *HssClient) UpdateSecurityGroupPolicyInvoker(request *model.UpdateSecurityGroupPolicyRequest) *UpdateSecurityGroupPolicyInvoker {
	requestDef := GenReqDefForUpdateSecurityGroupPolicy()
	return &UpdateSecurityGroupPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSystemUserWhiteList 修改系统用户白名单
//
// 修改系统用户白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) UpdateSystemUserWhiteList(request *model.UpdateSystemUserWhiteListRequest) (*model.UpdateSystemUserWhiteListResponse, error) {
	requestDef := GenReqDefForUpdateSystemUserWhiteList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSystemUserWhiteListResponse), nil
	}
}

// UpdateSystemUserWhiteListInvoker 修改系统用户白名单
func (c *HssClient) UpdateSystemUserWhiteListInvoker(request *model.UpdateSystemUserWhiteListRequest) *UpdateSystemUserWhiteListInvoker {
	requestDef := GenReqDefForUpdateSystemUserWhiteList()
	return &UpdateSystemUserWhiteListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeAntivirusPayPerScanStatus 修改“病毒查杀按次计费”开关状态
//
// 修改“病毒查杀按次计费”开关状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ChangeAntivirusPayPerScanStatus(request *model.ChangeAntivirusPayPerScanStatusRequest) (*model.ChangeAntivirusPayPerScanStatusResponse, error) {
	requestDef := GenReqDefForChangeAntivirusPayPerScanStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeAntivirusPayPerScanStatusResponse), nil
	}
}

// ChangeAntivirusPayPerScanStatusInvoker 修改“病毒查杀按次计费”开关状态
func (c *HssClient) ChangeAntivirusPayPerScanStatusInvoker(request *model.ChangeAntivirusPayPerScanStatusRequest) *ChangeAntivirusPayPerScanStatusInvoker {
	requestDef := GenReqDefForChangeAntivirusPayPerScanStatus()
	return &ChangeAntivirusPayPerScanStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAntiVirusPaidTask 创建付费病毒扫描任务
//
// 创建付费病毒扫描任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) CreateAntiVirusPaidTask(request *model.CreateAntiVirusPaidTaskRequest) (*model.CreateAntiVirusPaidTaskResponse, error) {
	requestDef := GenReqDefForCreateAntiVirusPaidTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAntiVirusPaidTaskResponse), nil
	}
}

// CreateAntiVirusPaidTaskInvoker 创建付费病毒扫描任务
func (c *HssClient) CreateAntiVirusPaidTaskInvoker(request *model.CreateAntiVirusPaidTaskRequest) *CreateAntiVirusPaidTaskInvoker {
	requestDef := GenReqDefForCreateAntiVirusPaidTask()
	return &CreateAntiVirusPaidTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAntiVirusPaidHosts 查询付费病毒查杀服务器列表
//
// 查询付费病毒查杀服务器列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListAntiVirusPaidHosts(request *model.ListAntiVirusPaidHostsRequest) (*model.ListAntiVirusPaidHostsResponse, error) {
	requestDef := GenReqDefForListAntiVirusPaidHosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAntiVirusPaidHostsResponse), nil
	}
}

// ListAntiVirusPaidHostsInvoker 查询付费病毒查杀服务器列表
func (c *HssClient) ListAntiVirusPaidHostsInvoker(request *model.ListAntiVirusPaidHostsRequest) *ListAntiVirusPaidHostsInvoker {
	requestDef := GenReqDefForListAntiVirusPaidHosts()
	return &ListAntiVirusPaidHostsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAntivirusFreeQuota 查询病毒查杀免费扫描次数
//
// 查询病毒查杀免费扫描次数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowAntivirusFreeQuota(request *model.ShowAntivirusFreeQuotaRequest) (*model.ShowAntivirusFreeQuotaResponse, error) {
	requestDef := GenReqDefForShowAntivirusFreeQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAntivirusFreeQuotaResponse), nil
	}
}

// ShowAntivirusFreeQuotaInvoker 查询病毒查杀免费扫描次数
func (c *HssClient) ShowAntivirusFreeQuotaInvoker(request *model.ShowAntivirusFreeQuotaRequest) *ShowAntivirusFreeQuotaInvoker {
	requestDef := GenReqDefForShowAntivirusFreeQuota()
	return &ShowAntivirusFreeQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddAppWhitelistPolicyHost 白名单策略添加主机
//
// 白名单策略添加主机
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) AddAppWhitelistPolicyHost(request *model.AddAppWhitelistPolicyHostRequest) (*model.AddAppWhitelistPolicyHostResponse, error) {
	requestDef := GenReqDefForAddAppWhitelistPolicyHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddAppWhitelistPolicyHostResponse), nil
	}
}

// AddAppWhitelistPolicyHostInvoker 白名单策略添加主机
func (c *HssClient) AddAppWhitelistPolicyHostInvoker(request *model.AddAppWhitelistPolicyHostRequest) *AddAppWhitelistPolicyHostInvoker {
	requestDef := GenReqDefForAddAppWhitelistPolicyHost()
	return &AddAppWhitelistPolicyHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddAppWhitelistPolicyProcess 新增进程白名单策略进程
//
// 新增进程白名单策略进程
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) AddAppWhitelistPolicyProcess(request *model.AddAppWhitelistPolicyProcessRequest) (*model.AddAppWhitelistPolicyProcessResponse, error) {
	requestDef := GenReqDefForAddAppWhitelistPolicyProcess()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddAppWhitelistPolicyProcessResponse), nil
	}
}

// AddAppWhitelistPolicyProcessInvoker 新增进程白名单策略进程
func (c *HssClient) AddAppWhitelistPolicyProcessInvoker(request *model.AddAppWhitelistPolicyProcessRequest) *AddAppWhitelistPolicyProcessInvoker {
	requestDef := GenReqDefForAddAppWhitelistPolicyProcess()
	return &AddAppWhitelistPolicyProcessInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeAppWhitelistPolicy 修改白名单策略
//
// 修改白名单策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ChangeAppWhitelistPolicy(request *model.ChangeAppWhitelistPolicyRequest) (*model.ChangeAppWhitelistPolicyResponse, error) {
	requestDef := GenReqDefForChangeAppWhitelistPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeAppWhitelistPolicyResponse), nil
	}
}

// ChangeAppWhitelistPolicyInvoker 修改白名单策略
func (c *HssClient) ChangeAppWhitelistPolicyInvoker(request *model.ChangeAppWhitelistPolicyRequest) *ChangeAppWhitelistPolicyInvoker {
	requestDef := GenReqDefForChangeAppWhitelistPolicy()
	return &ChangeAppWhitelistPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeAppWhitelistPolicyProcessStatus 标记进程白名单策略识别进程
//
// 标记进程白名单策略识别进程
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ChangeAppWhitelistPolicyProcessStatus(request *model.ChangeAppWhitelistPolicyProcessStatusRequest) (*model.ChangeAppWhitelistPolicyProcessStatusResponse, error) {
	requestDef := GenReqDefForChangeAppWhitelistPolicyProcessStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeAppWhitelistPolicyProcessStatusResponse), nil
	}
}

// ChangeAppWhitelistPolicyProcessStatusInvoker 标记进程白名单策略识别进程
func (c *HssClient) ChangeAppWhitelistPolicyProcessStatusInvoker(request *model.ChangeAppWhitelistPolicyProcessStatusRequest) *ChangeAppWhitelistPolicyProcessStatusInvoker {
	requestDef := GenReqDefForChangeAppWhitelistPolicyProcessStatus()
	return &ChangeAppWhitelistPolicyProcessStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAppWhitelistPolicy 创建白名单策略
//
// 创建白名单策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) CreateAppWhitelistPolicy(request *model.CreateAppWhitelistPolicyRequest) (*model.CreateAppWhitelistPolicyResponse, error) {
	requestDef := GenReqDefForCreateAppWhitelistPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppWhitelistPolicyResponse), nil
	}
}

// CreateAppWhitelistPolicyInvoker 创建白名单策略
func (c *HssClient) CreateAppWhitelistPolicyInvoker(request *model.CreateAppWhitelistPolicyRequest) *CreateAppWhitelistPolicyInvoker {
	requestDef := GenReqDefForCreateAppWhitelistPolicy()
	return &CreateAppWhitelistPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAppWhitelistPolicy 删除白名单策略
//
// 删除白名单策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) DeleteAppWhitelistPolicy(request *model.DeleteAppWhitelistPolicyRequest) (*model.DeleteAppWhitelistPolicyResponse, error) {
	requestDef := GenReqDefForDeleteAppWhitelistPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppWhitelistPolicyResponse), nil
	}
}

// DeleteAppWhitelistPolicyInvoker 删除白名单策略
func (c *HssClient) DeleteAppWhitelistPolicyInvoker(request *model.DeleteAppWhitelistPolicyRequest) *DeleteAppWhitelistPolicyInvoker {
	requestDef := GenReqDefForDeleteAppWhitelistPolicy()
	return &DeleteAppWhitelistPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAppWhitelistPolicyHost 白名单策略删除主机
//
// 白名单策略删除主机
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) DeleteAppWhitelistPolicyHost(request *model.DeleteAppWhitelistPolicyHostRequest) (*model.DeleteAppWhitelistPolicyHostResponse, error) {
	requestDef := GenReqDefForDeleteAppWhitelistPolicyHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppWhitelistPolicyHostResponse), nil
	}
}

// DeleteAppWhitelistPolicyHostInvoker 白名单策略删除主机
func (c *HssClient) DeleteAppWhitelistPolicyHostInvoker(request *model.DeleteAppWhitelistPolicyHostRequest) *DeleteAppWhitelistPolicyHostInvoker {
	requestDef := GenReqDefForDeleteAppWhitelistPolicyHost()
	return &DeleteAppWhitelistPolicyHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppWhitelistEvent 查询进程白名单可疑进程
//
// 查询进程白名单可疑进程
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListAppWhitelistEvent(request *model.ListAppWhitelistEventRequest) (*model.ListAppWhitelistEventResponse, error) {
	requestDef := GenReqDefForListAppWhitelistEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppWhitelistEventResponse), nil
	}
}

// ListAppWhitelistEventInvoker 查询进程白名单可疑进程
func (c *HssClient) ListAppWhitelistEventInvoker(request *model.ListAppWhitelistEventRequest) *ListAppWhitelistEventInvoker {
	requestDef := GenReqDefForListAppWhitelistEvent()
	return &ListAppWhitelistEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppWhitelistHostStatus 查询进程白名单可选服务器列表
//
// 查询进程白名单可选服务器列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListAppWhitelistHostStatus(request *model.ListAppWhitelistHostStatusRequest) (*model.ListAppWhitelistHostStatusResponse, error) {
	requestDef := GenReqDefForListAppWhitelistHostStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppWhitelistHostStatusResponse), nil
	}
}

// ListAppWhitelistHostStatusInvoker 查询进程白名单可选服务器列表
func (c *HssClient) ListAppWhitelistHostStatusInvoker(request *model.ListAppWhitelistHostStatusRequest) *ListAppWhitelistHostStatusInvoker {
	requestDef := GenReqDefForListAppWhitelistHostStatus()
	return &ListAppWhitelistHostStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppWhitelistPolicy 查询进程白名单策略列表
//
// 查询进程白名单策略列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListAppWhitelistPolicy(request *model.ListAppWhitelistPolicyRequest) (*model.ListAppWhitelistPolicyResponse, error) {
	requestDef := GenReqDefForListAppWhitelistPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppWhitelistPolicyResponse), nil
	}
}

// ListAppWhitelistPolicyInvoker 查询进程白名单策略列表
func (c *HssClient) ListAppWhitelistPolicyInvoker(request *model.ListAppWhitelistPolicyRequest) *ListAppWhitelistPolicyInvoker {
	requestDef := GenReqDefForListAppWhitelistPolicy()
	return &ListAppWhitelistPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppWhitelistPolicyHost 查询进程白名单策略关联主机列表
//
// 查询进程白名单策略关联主机列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListAppWhitelistPolicyHost(request *model.ListAppWhitelistPolicyHostRequest) (*model.ListAppWhitelistPolicyHostResponse, error) {
	requestDef := GenReqDefForListAppWhitelistPolicyHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppWhitelistPolicyHostResponse), nil
	}
}

// ListAppWhitelistPolicyHostInvoker 查询进程白名单策略关联主机列表
func (c *HssClient) ListAppWhitelistPolicyHostInvoker(request *model.ListAppWhitelistPolicyHostRequest) *ListAppWhitelistPolicyHostInvoker {
	requestDef := GenReqDefForListAppWhitelistPolicyHost()
	return &ListAppWhitelistPolicyHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppWhitelistPolicyProcess 查询进程白名单策略识别进程
//
// 查询进程白名单策略识别进程
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListAppWhitelistPolicyProcess(request *model.ListAppWhitelistPolicyProcessRequest) (*model.ListAppWhitelistPolicyProcessResponse, error) {
	requestDef := GenReqDefForListAppWhitelistPolicyProcess()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppWhitelistPolicyProcessResponse), nil
	}
}

// ListAppWhitelistPolicyProcessInvoker 查询进程白名单策略识别进程
func (c *HssClient) ListAppWhitelistPolicyProcessInvoker(request *model.ListAppWhitelistPolicyProcessRequest) *ListAppWhitelistPolicyProcessInvoker {
	requestDef := GenReqDefForListAppWhitelistPolicyProcess()
	return &ListAppWhitelistPolicyProcessInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppWhitelistPolicyProcessExtend 查询进程白名单策略进程扩展列表
//
// 查询进程白名单策略进程扩展列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListAppWhitelistPolicyProcessExtend(request *model.ListAppWhitelistPolicyProcessExtendRequest) (*model.ListAppWhitelistPolicyProcessExtendResponse, error) {
	requestDef := GenReqDefForListAppWhitelistPolicyProcessExtend()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppWhitelistPolicyProcessExtendResponse), nil
	}
}

// ListAppWhitelistPolicyProcessExtendInvoker 查询进程白名单策略进程扩展列表
func (c *HssClient) ListAppWhitelistPolicyProcessExtendInvoker(request *model.ListAppWhitelistPolicyProcessExtendRequest) *ListAppWhitelistPolicyProcessExtendInvoker {
	requestDef := GenReqDefForListAppWhitelistPolicyProcessExtend()
	return &ListAppWhitelistPolicyProcessExtendInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAppWhitelistAgentStatics 统计agent版本不匹配主机数量
//
// 统计agent版本不匹配主机数量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowAppWhitelistAgentStatics(request *model.ShowAppWhitelistAgentStaticsRequest) (*model.ShowAppWhitelistAgentStaticsResponse, error) {
	requestDef := GenReqDefForShowAppWhitelistAgentStatics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppWhitelistAgentStaticsResponse), nil
	}
}

// ShowAppWhitelistAgentStaticsInvoker 统计agent版本不匹配主机数量
func (c *HssClient) ShowAppWhitelistAgentStaticsInvoker(request *model.ShowAppWhitelistAgentStaticsRequest) *ShowAppWhitelistAgentStaticsInvoker {
	requestDef := GenReqDefForShowAppWhitelistAgentStatics()
	return &ShowAppWhitelistAgentStaticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAppWhitelistPolicy 查询进程白名单策略详情
//
// 查询进程白名单策略详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowAppWhitelistPolicy(request *model.ShowAppWhitelistPolicyRequest) (*model.ShowAppWhitelistPolicyResponse, error) {
	requestDef := GenReqDefForShowAppWhitelistPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppWhitelistPolicyResponse), nil
	}
}

// ShowAppWhitelistPolicyInvoker 查询进程白名单策略详情
func (c *HssClient) ShowAppWhitelistPolicyInvoker(request *model.ShowAppWhitelistPolicyRequest) *ShowAppWhitelistPolicyInvoker {
	requestDef := GenReqDefForShowAppWhitelistPolicy()
	return &ShowAppWhitelistPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchAppWhitelistPolicyHost 应用白名单策略
//
// 应用白名单策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) SwitchAppWhitelistPolicyHost(request *model.SwitchAppWhitelistPolicyHostRequest) (*model.SwitchAppWhitelistPolicyHostResponse, error) {
	requestDef := GenReqDefForSwitchAppWhitelistPolicyHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchAppWhitelistPolicyHostResponse), nil
	}
}

// SwitchAppWhitelistPolicyHostInvoker 应用白名单策略
func (c *HssClient) SwitchAppWhitelistPolicyHostInvoker(request *model.SwitchAppWhitelistPolicyHostRequest) *SwitchAppWhitelistPolicyHostInvoker {
	requestDef := GenReqDefForSwitchAppWhitelistPolicyHost()
	return &SwitchAppWhitelistPolicyHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchAppWhitelistPolicyLearnStatus 操作白名单策略学习状态
//
// 操作白名单策略学习状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) SwitchAppWhitelistPolicyLearnStatus(request *model.SwitchAppWhitelistPolicyLearnStatusRequest) (*model.SwitchAppWhitelistPolicyLearnStatusResponse, error) {
	requestDef := GenReqDefForSwitchAppWhitelistPolicyLearnStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchAppWhitelistPolicyLearnStatusResponse), nil
	}
}

// SwitchAppWhitelistPolicyLearnStatusInvoker 操作白名单策略学习状态
func (c *HssClient) SwitchAppWhitelistPolicyLearnStatusInvoker(request *model.SwitchAppWhitelistPolicyLearnStatusRequest) *SwitchAppWhitelistPolicyLearnStatusInvoker {
	requestDef := GenReqDefForSwitchAppWhitelistPolicyLearnStatus()
	return &SwitchAppWhitelistPolicyLearnStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHandleAffectBaseline 查询基线检查执行操作时影响的范围
//
// 查询基线检查执行操作时影响的范围
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListHandleAffectBaseline(request *model.ListHandleAffectBaselineRequest) (*model.ListHandleAffectBaselineResponse, error) {
	requestDef := GenReqDefForListHandleAffectBaseline()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHandleAffectBaselineResponse), nil
	}
}

// ListHandleAffectBaselineInvoker 查询基线检查执行操作时影响的范围
func (c *HssClient) ListHandleAffectBaselineInvoker(request *model.ListHandleAffectBaselineRequest) *ListHandleAffectBaselineInvoker {
	requestDef := GenReqDefForListHandleAffectBaseline()
	return &ListHandleAffectBaselineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CopyBaselinePolicyGroup 复制配置检测策略信息
//
// 复制配置检测策略信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) CopyBaselinePolicyGroup(request *model.CopyBaselinePolicyGroupRequest) (*model.CopyBaselinePolicyGroupResponse, error) {
	requestDef := GenReqDefForCopyBaselinePolicyGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyBaselinePolicyGroupResponse), nil
	}
}

// CopyBaselinePolicyGroupInvoker 复制配置检测策略信息
func (c *HssClient) CopyBaselinePolicyGroupInvoker(request *model.CopyBaselinePolicyGroupRequest) *CopyBaselinePolicyGroupInvoker {
	requestDef := GenReqDefForCopyBaselinePolicyGroup()
	return &CopyBaselinePolicyGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBaselineDirectory 查询基线检测策略的基线目录信息
//
// 查询基线检测策略的基线目录信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowBaselineDirectory(request *model.ShowBaselineDirectoryRequest) (*model.ShowBaselineDirectoryResponse, error) {
	requestDef := GenReqDefForShowBaselineDirectory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBaselineDirectoryResponse), nil
	}
}

// ShowBaselineDirectoryInvoker 查询基线检测策略的基线目录信息
func (c *HssClient) ShowBaselineDirectoryInvoker(request *model.ShowBaselineDirectoryRequest) *ShowBaselineDirectoryInvoker {
	requestDef := GenReqDefForShowBaselineDirectory()
	return &ShowBaselineDirectoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusterRiskAffectResources 查询集群风险影响的集群资源列表
//
// 查询集群风险影响的集群资源列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListClusterRiskAffectResources(request *model.ListClusterRiskAffectResourcesRequest) (*model.ListClusterRiskAffectResourcesResponse, error) {
	requestDef := GenReqDefForListClusterRiskAffectResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterRiskAffectResourcesResponse), nil
	}
}

// ListClusterRiskAffectResourcesInvoker 查询集群风险影响的集群资源列表
func (c *HssClient) ListClusterRiskAffectResourcesInvoker(request *model.ListClusterRiskAffectResourcesRequest) *ListClusterRiskAffectResourcesInvoker {
	requestDef := GenReqDefForListClusterRiskAffectResources()
	return &ListClusterRiskAffectResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusterRisks 查询集群风险列表
//
// 查询集群风险列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListClusterRisks(request *model.ListClusterRisksRequest) (*model.ListClusterRisksResponse, error) {
	requestDef := GenReqDefForListClusterRisks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterRisksResponse), nil
	}
}

// ListClusterRisksInvoker 查询集群风险列表
func (c *HssClient) ListClusterRisksInvoker(request *model.ListClusterRisksRequest) *ListClusterRisksInvoker {
	requestDef := GenReqDefForListClusterRisks()
	return &ListClusterRisksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowClusterScanStatistics 查询集群扫描统计数据
//
// 查询集群扫描统计数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowClusterScanStatistics(request *model.ShowClusterScanStatisticsRequest) (*model.ShowClusterScanStatisticsResponse, error) {
	requestDef := GenReqDefForShowClusterScanStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusterScanStatisticsResponse), nil
	}
}

// ShowClusterScanStatisticsInvoker 查询集群扫描统计数据
func (c *HssClient) ShowClusterScanStatisticsInvoker(request *model.ShowClusterScanStatisticsRequest) *ShowClusterScanStatisticsInvoker {
	requestDef := GenReqDefForShowClusterScanStatistics()
	return &ShowClusterScanStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUploadFiles 批量上传文件
//
// 批量上传文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) BatchUploadFiles(request *model.BatchUploadFilesRequest) (*model.BatchUploadFilesResponse, error) {
	requestDef := GenReqDefForBatchUploadFiles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUploadFilesResponse), nil
	}
}

// BatchUploadFilesInvoker 批量上传文件
func (c *HssClient) BatchUploadFilesInvoker(request *model.BatchUploadFilesRequest) *BatchUploadFilesInvoker {
	requestDef := GenReqDefForBatchUploadFiles()
	return &BatchUploadFilesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectConfigs 查询项目配置
//
// 查询项目配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListProjectConfigs(request *model.ListProjectConfigsRequest) (*model.ListProjectConfigsResponse, error) {
	requestDef := GenReqDefForListProjectConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectConfigsResponse), nil
	}
}

// ListProjectConfigsInvoker 查询项目配置
func (c *HssClient) ListProjectConfigsInvoker(request *model.ListProjectConfigsRequest) *ListProjectConfigsInvoker {
	requestDef := GenReqDefForListProjectConfigs()
	return &ListProjectConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyProjectConfigs 修改项目配置
//
// 修改项目配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ModifyProjectConfigs(request *model.ModifyProjectConfigsRequest) (*model.ModifyProjectConfigsResponse, error) {
	requestDef := GenReqDefForModifyProjectConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyProjectConfigsResponse), nil
	}
}

// ModifyProjectConfigsInvoker 修改项目配置
func (c *HssClient) ModifyProjectConfigsInvoker(request *model.ModifyProjectConfigsRequest) *ModifyProjectConfigsInvoker {
	requestDef := GenReqDefForModifyProjectConfigs()
	return &ModifyProjectConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SaveBrowsingHistory 保存用户访问记录
//
// 保存用户访问记录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) SaveBrowsingHistory(request *model.SaveBrowsingHistoryRequest) (*model.SaveBrowsingHistoryResponse, error) {
	requestDef := GenReqDefForSaveBrowsingHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SaveBrowsingHistoryResponse), nil
	}
}

// SaveBrowsingHistoryInvoker 保存用户访问记录
func (c *HssClient) SaveBrowsingHistoryInvoker(request *model.SaveBrowsingHistoryRequest) *SaveBrowsingHistoryInvoker {
	requestDef := GenReqDefForSaveBrowsingHistory()
	return &SaveBrowsingHistoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCicdConfigurations 删除CI/CD配置
//
// 删除CI/CD配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) DeleteCicdConfigurations(request *model.DeleteCicdConfigurationsRequest) (*model.DeleteCicdConfigurationsResponse, error) {
	requestDef := GenReqDefForDeleteCicdConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCicdConfigurationsResponse), nil
	}
}

// DeleteCicdConfigurationsInvoker 删除CI/CD配置
func (c *HssClient) DeleteCicdConfigurationsInvoker(request *model.DeleteCicdConfigurationsRequest) *DeleteCicdConfigurationsInvoker {
	requestDef := GenReqDefForDeleteCicdConfigurations()
	return &DeleteCicdConfigurationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportImageSecurityReportTask 创建镜像安全报告信息导出任务（支持全量/批量导出）
//
// 创建镜像安全报告信息导出任务（支持全量/批量导出）,支持导出恶意文件、软件信息、文件信息、敏感信息、软件合规、镜像构建指令风险。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ExportImageSecurityReportTask(request *model.ExportImageSecurityReportTaskRequest) (*model.ExportImageSecurityReportTaskResponse, error) {
	requestDef := GenReqDefForExportImageSecurityReportTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportImageSecurityReportTaskResponse), nil
	}
}

// ExportImageSecurityReportTaskInvoker 创建镜像安全报告信息导出任务（支持全量/批量导出）
func (c *HssClient) ExportImageSecurityReportTaskInvoker(request *model.ExportImageSecurityReportTaskRequest) *ExportImageSecurityReportTaskInvoker {
	requestDef := GenReqDefForExportImageSecurityReportTask()
	return &ExportImageSecurityReportTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAssociateRegistries 获取镜像同步任务关联的镜像仓的信息
//
// 获取镜像同步任务关联的镜像仓的信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListAssociateRegistries(request *model.ListAssociateRegistriesRequest) (*model.ListAssociateRegistriesResponse, error) {
	requestDef := GenReqDefForListAssociateRegistries()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAssociateRegistriesResponse), nil
	}
}

// ListAssociateRegistriesInvoker 获取镜像同步任务关联的镜像仓的信息
func (c *HssClient) ListAssociateRegistriesInvoker(request *model.ListAssociateRegistriesRequest) *ListAssociateRegistriesInvoker {
	requestDef := GenReqDefForListAssociateRegistries()
	return &ListAssociateRegistriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCicdConfigurations 查询cicd配置列表
//
// 查询cicd配置列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListCicdConfigurations(request *model.ListCicdConfigurationsRequest) (*model.ListCicdConfigurationsResponse, error) {
	requestDef := GenReqDefForListCicdConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCicdConfigurationsResponse), nil
	}
}

// ListCicdConfigurationsInvoker 查询cicd配置列表
func (c *HssClient) ListCicdConfigurationsInvoker(request *model.ListCicdConfigurationsRequest) *ListCicdConfigurationsInvoker {
	requestDef := GenReqDefForListCicdConfigurations()
	return &ListCicdConfigurationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyCicdConfiguration 修改CI/CD配置
//
// 修改CI/CD配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ModifyCicdConfiguration(request *model.ModifyCicdConfigurationRequest) (*model.ModifyCicdConfigurationResponse, error) {
	requestDef := GenReqDefForModifyCicdConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyCicdConfigurationResponse), nil
	}
}

// ModifyCicdConfigurationInvoker 修改CI/CD配置
func (c *HssClient) ModifyCicdConfigurationInvoker(request *model.ModifyCicdConfigurationRequest) *ModifyCicdConfigurationInvoker {
	requestDef := GenReqDefForModifyCicdConfiguration()
	return &ModifyCicdConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCicdConfiguration 查询CI/CD配置信息
//
// 查询CI/CD配置信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowCicdConfiguration(request *model.ShowCicdConfigurationRequest) (*model.ShowCicdConfigurationResponse, error) {
	requestDef := GenReqDefForShowCicdConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCicdConfigurationResponse), nil
	}
}

// ShowCicdConfigurationInvoker 查询CI/CD配置信息
func (c *HssClient) ShowCicdConfigurationInvoker(request *model.ShowCicdConfigurationRequest) *ShowCicdConfigurationInvoker {
	requestDef := GenReqDefForShowCicdConfiguration()
	return &ShowCicdConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFileEvents 变更文件列表
//
// 变更文件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListFileEvents(request *model.ListFileEventsRequest) (*model.ListFileEventsResponse, error) {
	requestDef := GenReqDefForListFileEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFileEventsResponse), nil
	}
}

// ListFileEventsInvoker 变更文件列表
func (c *HssClient) ListFileEventsInvoker(request *model.ListFileEventsRequest) *ListFileEventsInvoker {
	requestDef := GenReqDefForListFileEvents()
	return &ListFileEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFileHostEventDetails 某个服务器变更文件信息
//
// 某个服务器变更文件信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListFileHostEventDetails(request *model.ListFileHostEventDetailsRequest) (*model.ListFileHostEventDetailsResponse, error) {
	requestDef := GenReqDefForListFileHostEventDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFileHostEventDetailsResponse), nil
	}
}

// ListFileHostEventDetailsInvoker 某个服务器变更文件信息
func (c *HssClient) ListFileHostEventDetailsInvoker(request *model.ListFileHostEventDetailsRequest) *ListFileHostEventDetailsInvoker {
	requestDef := GenReqDefForListFileHostEventDetails()
	return &ListFileHostEventDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFileHosts 云服务器变更列表
//
// 云服务器变更列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListFileHosts(request *model.ListFileHostsRequest) (*model.ListFileHostsResponse, error) {
	requestDef := GenReqDefForListFileHosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFileHostsResponse), nil
	}
}

// ListFileHostsInvoker 云服务器变更列表
func (c *HssClient) ListFileHostsInvoker(request *model.ListFileHostsRequest) *ListFileHostsInvoker {
	requestDef := GenReqDefForListFileHosts()
	return &ListFileHostsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFileStatistic 获取服务器文件统计信息
//
// 获取服务器文件统计信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowFileStatistic(request *model.ShowFileStatisticRequest) (*model.ShowFileStatisticResponse, error) {
	requestDef := GenReqDefForShowFileStatistic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFileStatisticResponse), nil
	}
}

// ShowFileStatisticInvoker 获取服务器文件统计信息
func (c *HssClient) ShowFileStatisticInvoker(request *model.ShowFileStatisticRequest) *ShowFileStatisticInvoker {
	requestDef := GenReqDefForShowFileStatistic()
	return &ShowFileStatisticInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIacFileRiskPaths 获取iac文件风险路径列表
//
// 获取iac文件风险路径列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListIacFileRiskPaths(request *model.ListIacFileRiskPathsRequest) (*model.ListIacFileRiskPathsResponse, error) {
	requestDef := GenReqDefForListIacFileRiskPaths()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIacFileRiskPathsResponse), nil
	}
}

// ListIacFileRiskPathsInvoker 获取iac文件风险路径列表
func (c *HssClient) ListIacFileRiskPathsInvoker(request *model.ListIacFileRiskPathsRequest) *ListIacFileRiskPathsInvoker {
	requestDef := GenReqDefForListIacFileRiskPaths()
	return &ListIacFileRiskPathsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIacFileRisks 获取iac文件风险列表
//
// 获取iac文件风险列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListIacFileRisks(request *model.ListIacFileRisksRequest) (*model.ListIacFileRisksResponse, error) {
	requestDef := GenReqDefForListIacFileRisks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIacFileRisksResponse), nil
	}
}

// ListIacFileRisksInvoker 获取iac文件风险列表
func (c *HssClient) ListIacFileRisksInvoker(request *model.ListIacFileRisksRequest) *ListIacFileRisksInvoker {
	requestDef := GenReqDefForListIacFileRisks()
	return &ListIacFileRisksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIacFiles 获取iac文件列表
//
// 获取iac文件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListIacFiles(request *model.ListIacFilesRequest) (*model.ListIacFilesResponse, error) {
	requestDef := GenReqDefForListIacFiles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIacFilesResponse), nil
	}
}

// ListIacFilesInvoker 获取iac文件列表
func (c *HssClient) ListIacFilesInvoker(request *model.ListIacFilesRequest) *ListIacFilesInvoker {
	requestDef := GenReqDefForListIacFiles()
	return &ListIacFilesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTask 创建任务
//
// 创建任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) CreateTask(request *model.CreateTaskRequest) (*model.CreateTaskResponse, error) {
	requestDef := GenReqDefForCreateTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTaskResponse), nil
	}
}

// CreateTaskInvoker 创建任务
func (c *HssClient) CreateTaskInvoker(request *model.CreateTaskRequest) *CreateTaskInvoker {
	requestDef := GenReqDefForCreateTask()
	return &CreateTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTaskResources 查询当前任务关联的资源列表
//
// 查询当前任务关联的资源列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListTaskResources(request *model.ListTaskResourcesRequest) (*model.ListTaskResourcesResponse, error) {
	requestDef := GenReqDefForListTaskResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTaskResourcesResponse), nil
	}
}

// ListTaskResourcesInvoker 查询当前任务关联的资源列表
func (c *HssClient) ListTaskResourcesInvoker(request *model.ListTaskResourcesRequest) *ListTaskResourcesInvoker {
	requestDef := GenReqDefForListTaskResources()
	return &ListTaskResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTasks 查询任务列表
//
// 查询任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListTasks(request *model.ListTasksRequest) (*model.ListTasksResponse, error) {
	requestDef := GenReqDefForListTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTasksResponse), nil
	}
}

// ListTasksInvoker 查询任务列表
func (c *HssClient) ListTasksInvoker(request *model.ListTasksRequest) *ListTasksInvoker {
	requestDef := GenReqDefForListTasks()
	return &ListTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTaskStatistics 查询任务统计数据
//
// 查询任务统计数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowTaskStatistics(request *model.ShowTaskStatisticsRequest) (*model.ShowTaskStatisticsResponse, error) {
	requestDef := GenReqDefForShowTaskStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskStatisticsResponse), nil
	}
}

// ShowTaskStatisticsInvoker 查询任务统计数据
func (c *HssClient) ShowTaskStatisticsInvoker(request *model.ShowTaskStatisticsRequest) *ShowTaskStatisticsInvoker {
	requestDef := GenReqDefForShowTaskStatistics()
	return &ShowTaskStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportHandledVulnerabilities 创建历史处理的漏洞导出任务
//
// 创建历史处理的漏洞导出任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ExportHandledVulnerabilities(request *model.ExportHandledVulnerabilitiesRequest) (*model.ExportHandledVulnerabilitiesResponse, error) {
	requestDef := GenReqDefForExportHandledVulnerabilities()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportHandledVulnerabilitiesResponse), nil
	}
}

// ExportHandledVulnerabilitiesInvoker 创建历史处理的漏洞导出任务
func (c *HssClient) ExportHandledVulnerabilitiesInvoker(request *model.ExportHandledVulnerabilitiesRequest) *ExportHandledVulnerabilitiesInvoker {
	requestDef := GenReqDefForExportHandledVulnerabilities()
	return &ExportHandledVulnerabilitiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportVulHandleHistory 创建漏洞处置历史记录的导出任务
//
// 创建漏洞处置历史记录的导出任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ExportVulHandleHistory(request *model.ExportVulHandleHistoryRequest) (*model.ExportVulHandleHistoryResponse, error) {
	requestDef := GenReqDefForExportVulHandleHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportVulHandleHistoryResponse), nil
	}
}

// ExportVulHandleHistoryInvoker 创建漏洞处置历史记录的导出任务
func (c *HssClient) ExportVulHandleHistoryInvoker(request *model.ExportVulHandleHistoryRequest) *ExportVulHandleHistoryInvoker {
	requestDef := GenReqDefForExportVulHandleHistory()
	return &ExportVulHandleHistoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVulContainerApps 查询单个漏洞影响的容器app信息
//
// 查询单个漏洞影响的容器app信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListVulContainerApps(request *model.ListVulContainerAppsRequest) (*model.ListVulContainerAppsResponse, error) {
	requestDef := GenReqDefForListVulContainerApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVulContainerAppsResponse), nil
	}
}

// ListVulContainerAppsInvoker 查询单个漏洞影响的容器app信息
func (c *HssClient) ListVulContainerAppsInvoker(request *model.ListVulContainerAppsRequest) *ListVulContainerAppsInvoker {
	requestDef := GenReqDefForListVulContainerApps()
	return &ListVulContainerAppsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVulContainers 查询单个漏洞影响的容器信息
//
// 查询单个漏洞影响的容器信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListVulContainers(request *model.ListVulContainersRequest) (*model.ListVulContainersResponse, error) {
	requestDef := GenReqDefForListVulContainers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVulContainersResponse), nil
	}
}

// ListVulContainersInvoker 查询单个漏洞影响的容器信息
func (c *HssClient) ListVulContainersInvoker(request *model.ListVulContainersRequest) *ListVulContainersInvoker {
	requestDef := GenReqDefForListVulContainers()
	return &ListVulContainersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecordUserViewVulTask 记录用户查看漏洞任务管理页面的最后时间
//
// 记录用户查看漏洞任务管理页面的最后时间
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) RecordUserViewVulTask(request *model.RecordUserViewVulTaskRequest) (*model.RecordUserViewVulTaskResponse, error) {
	requestDef := GenReqDefForRecordUserViewVulTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecordUserViewVulTaskResponse), nil
	}
}

// RecordUserViewVulTaskInvoker 记录用户查看漏洞任务管理页面的最后时间
func (c *HssClient) RecordUserViewVulTaskInvoker(request *model.RecordUserViewVulTaskRequest) *RecordUserViewVulTaskInvoker {
	requestDef := GenReqDefForRecordUserViewVulTask()
	return &RecordUserViewVulTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVulTaskStatistics 获取漏洞任务的未读数量
//
// 获取漏洞任务的未读数量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowVulTaskStatistics(request *model.ShowVulTaskStatisticsRequest) (*model.ShowVulTaskStatisticsResponse, error) {
	requestDef := GenReqDefForShowVulTaskStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVulTaskStatisticsResponse), nil
	}
}

// ShowVulTaskStatisticsInvoker 获取漏洞任务的未读数量
func (c *HssClient) ShowVulTaskStatisticsInvoker(request *model.ShowVulTaskStatisticsRequest) *ShowVulTaskStatisticsInvoker {
	requestDef := GenReqDefForShowVulTaskStatistics()
	return &ShowVulTaskStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchStartWebTamperProtection 批量开启网页防篡改防护
//
// 批量开启网页防篡改防护
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) BatchStartWebTamperProtection(request *model.BatchStartWebTamperProtectionRequest) (*model.BatchStartWebTamperProtectionResponse, error) {
	requestDef := GenReqDefForBatchStartWebTamperProtection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStartWebTamperProtectionResponse), nil
	}
}

// BatchStartWebTamperProtectionInvoker 批量开启网页防篡改防护
func (c *HssClient) BatchStartWebTamperProtectionInvoker(request *model.BatchStartWebTamperProtectionRequest) *BatchStartWebTamperProtectionInvoker {
	requestDef := GenReqDefForBatchStartWebTamperProtection()
	return &BatchStartWebTamperProtectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportWebTamperHost 导出网页防篡改主机列表
//
// 导出网页防篡改主机列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ExportWebTamperHost(request *model.ExportWebTamperHostRequest) (*model.ExportWebTamperHostResponse, error) {
	requestDef := GenReqDefForExportWebTamperHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportWebTamperHostResponse), nil
	}
}

// ExportWebTamperHostInvoker 导出网页防篡改主机列表
func (c *HssClient) ExportWebTamperHostInvoker(request *model.ExportWebTamperHostRequest) *ExportWebTamperHostInvoker {
	requestDef := GenReqDefForExportWebTamperHost()
	return &ExportWebTamperHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWebTamperHost 查询网页防篡改可选服务器列表
//
// 查询网页防篡改可选服务器列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListWebTamperHost(request *model.ListWebTamperHostRequest) (*model.ListWebTamperHostResponse, error) {
	requestDef := GenReqDefForListWebTamperHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWebTamperHostResponse), nil
	}
}

// ListWebTamperHostInvoker 查询网页防篡改可选服务器列表
func (c *HssClient) ListWebTamperHostInvoker(request *model.ListWebTamperHostRequest) *ListWebTamperHostInvoker {
	requestDef := GenReqDefForListWebTamperHost()
	return &ListWebTamperHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWebTamperHostPolicy 查看网页防篡改策略信息
//
// 查看网页防篡改策略信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowWebTamperHostPolicy(request *model.ShowWebTamperHostPolicyRequest) (*model.ShowWebTamperHostPolicyResponse, error) {
	requestDef := GenReqDefForShowWebTamperHostPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWebTamperHostPolicyResponse), nil
	}
}

// ShowWebTamperHostPolicyInvoker 查看网页防篡改策略信息
func (c *HssClient) ShowWebTamperHostPolicyInvoker(request *model.ShowWebTamperHostPolicyRequest) *ShowWebTamperHostPolicyInvoker {
	requestDef := GenReqDefForShowWebTamperHostPolicy()
	return &ShowWebTamperHostPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWebTamperRaspPath 查询动态网页防篡改的Tomcat bin目录
//
// 查询动态网页防篡改的Tomcat bin目录：查询动态网页防篡改功能配置的Tomcat bin目录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowWebTamperRaspPath(request *model.ShowWebTamperRaspPathRequest) (*model.ShowWebTamperRaspPathResponse, error) {
	requestDef := GenReqDefForShowWebTamperRaspPath()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWebTamperRaspPathResponse), nil
	}
}

// ShowWebTamperRaspPathInvoker 查询动态网页防篡改的Tomcat bin目录
func (c *HssClient) ShowWebTamperRaspPathInvoker(request *model.ShowWebTamperRaspPathRequest) *ShowWebTamperRaspPathInvoker {
	requestDef := GenReqDefForShowWebTamperRaspPath()
	return &ShowWebTamperRaspPathInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWebTamperHostPolicy 编辑网页防篡改策略信息
//
// 编辑网页防篡改策略信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) UpdateWebTamperHostPolicy(request *model.UpdateWebTamperHostPolicyRequest) (*model.UpdateWebTamperHostPolicyResponse, error) {
	requestDef := GenReqDefForUpdateWebTamperHostPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWebTamperHostPolicyResponse), nil
	}
}

// UpdateWebTamperHostPolicyInvoker 编辑网页防篡改策略信息
func (c *HssClient) UpdateWebTamperHostPolicyInvoker(request *model.UpdateWebTamperHostPolicyRequest) *UpdateWebTamperHostPolicyInvoker {
	requestDef := GenReqDefForUpdateWebTamperHostPolicy()
	return &UpdateWebTamperHostPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWebTamperRaspPath 修改动态网页防篡改的Tomcat bin目录
//
// 修改动态网页防篡改的Tomcat bin目录：修改动态网页防篡改的Tomcat bin目录，重新下发动态网页防篡改策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) UpdateWebTamperRaspPath(request *model.UpdateWebTamperRaspPathRequest) (*model.UpdateWebTamperRaspPathResponse, error) {
	requestDef := GenReqDefForUpdateWebTamperRaspPath()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWebTamperRaspPathResponse), nil
	}
}

// UpdateWebTamperRaspPathInvoker 修改动态网页防篡改的Tomcat bin目录
func (c *HssClient) UpdateWebTamperRaspPathInvoker(request *model.UpdateWebTamperRaspPathRequest) *UpdateWebTamperRaspPathInvoker {
	requestDef := GenReqDefForUpdateWebTamperRaspPath()
	return &UpdateWebTamperRaspPathInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
