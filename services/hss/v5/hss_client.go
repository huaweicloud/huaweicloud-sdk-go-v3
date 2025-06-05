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
// 查询备份存储库列表
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

// ListRansomwareProtectionNodes 查询勒索防护服务器列表
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

// ListRansomwareProtectionNodesInvoker 查询勒索防护服务器列表
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
