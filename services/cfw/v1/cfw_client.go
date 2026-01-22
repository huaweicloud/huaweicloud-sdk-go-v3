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

// DeleteIpBlacklist 删除已经导入的IP黑名单
//
// 删除流量过滤功能下已经导入的IP黑名单，指定生效范围进行删除。 标准版的墙只会存在生效范围为EIP的IP黑名单，专业版的墙会存在生效范围为EIP和NAT的IP黑名单。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteIpBlacklist(request *model.DeleteIpBlacklistRequest) (*model.DeleteIpBlacklistResponse, error) {
	requestDef := GenReqDefForDeleteIpBlacklist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteIpBlacklistResponse), nil
	}
}

// DeleteIpBlacklistInvoker 删除已经导入的IP黑名单
func (c *CfwClient) DeleteIpBlacklistInvoker(request *model.DeleteIpBlacklistRequest) *DeleteIpBlacklistInvoker {
	requestDef := GenReqDefForDeleteIpBlacklist()
	return &DeleteIpBlacklistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableIpBlacklist 开启或者关闭流量过滤的IP黑名单功能
//
// 开启或者关闭流量过滤功能，当前流量过滤是通过导入IP黑名单实现的。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) EnableIpBlacklist(request *model.EnableIpBlacklistRequest) (*model.EnableIpBlacklistResponse, error) {
	requestDef := GenReqDefForEnableIpBlacklist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableIpBlacklistResponse), nil
	}
}

// EnableIpBlacklistInvoker 开启或者关闭流量过滤的IP黑名单功能
func (c *CfwClient) EnableIpBlacklistInvoker(request *model.EnableIpBlacklistRequest) *EnableIpBlacklistInvoker {
	requestDef := GenReqDefForEnableIpBlacklist()
	return &EnableIpBlacklistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportIpBlacklist 导出用于流量过滤的IP黑名单
//
// 指定IP黑名单的名字进行导出，当前只有两种文件名，在EIP生效的文件名为ip-blacklist-eip.txt，在 NAT生效的文件名为ip-blacklist-nat.txt。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ExportIpBlacklist(request *model.ExportIpBlacklistRequest) (*model.ExportIpBlacklistResponse, error) {
	requestDef := GenReqDefForExportIpBlacklist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportIpBlacklistResponse), nil
	}
}

// ExportIpBlacklistInvoker 导出用于流量过滤的IP黑名单
func (c *CfwClient) ExportIpBlacklistInvoker(request *model.ExportIpBlacklistRequest) *ExportIpBlacklistInvoker {
	requestDef := GenReqDefForExportIpBlacklist()
	return &ExportIpBlacklistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportIpBlacklist 导入IP黑名单用于流量过滤
//
// 此接口用来导入IP黑名单，IP列表保存在request的body中，IP列表支持的格式如下：
// 单个IP地址，例如：100.1.1.10
// 连续的IP地址段，例如：80.1.1.3-80.1.1.30
// 掩码格式的网段，例如：6.6.6.0/24
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ImportIpBlacklist(request *model.ImportIpBlacklistRequest) (*model.ImportIpBlacklistResponse, error) {
	requestDef := GenReqDefForImportIpBlacklist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportIpBlacklistResponse), nil
	}
}

// ImportIpBlacklistInvoker 导入IP黑名单用于流量过滤
func (c *CfwClient) ImportIpBlacklistInvoker(request *model.ImportIpBlacklistRequest) *ImportIpBlacklistInvoker {
	requestDef := GenReqDefForImportIpBlacklist()
	return &ImportIpBlacklistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListIpBlacklist 获取导入的IP黑名单列表信息
//
// 获取防火墙实例中已经导入的IP黑名单信息，标准版防火墙只会显示一条EIP的记录，专业版防火墙可能显示EIP、NAT或EIP和NAT的记录，根据导入的情况确定。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListIpBlacklist(request *model.ListIpBlacklistRequest) (*model.ListIpBlacklistResponse, error) {
	requestDef := GenReqDefForListIpBlacklist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIpBlacklistResponse), nil
	}
}

// ListIpBlacklistInvoker 获取导入的IP黑名单列表信息
func (c *CfwClient) ListIpBlacklistInvoker(request *model.ListIpBlacklistRequest) *ListIpBlacklistInvoker {
	requestDef := GenReqDefForListIpBlacklist()
	return &ListIpBlacklistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIpBlacklistSwitch 获取流量过滤功能的开关信息
//
// 流量过滤功能可以打开或者关闭，通过此接口可以获取当前的开关信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListIpBlacklistSwitch(request *model.ListIpBlacklistSwitchRequest) (*model.ListIpBlacklistSwitchResponse, error) {
	requestDef := GenReqDefForListIpBlacklistSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIpBlacklistSwitchResponse), nil
	}
}

// ListIpBlacklistSwitchInvoker 获取流量过滤功能的开关信息
func (c *CfwClient) ListIpBlacklistSwitchInvoker(request *model.ListIpBlacklistSwitchRequest) *ListIpBlacklistSwitchInvoker {
	requestDef := GenReqDefForListIpBlacklistSwitch()
	return &ListIpBlacklistSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// RetryIpBlacklist 用于流量过滤的IP黑名单导入失败后进行重新导入
//
// 用于流量过滤的IP黑名单导入失败后，调用此接口进行重试。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) RetryIpBlacklist(request *model.RetryIpBlacklistRequest) (*model.RetryIpBlacklistResponse, error) {
	requestDef := GenReqDefForRetryIpBlacklist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RetryIpBlacklistResponse), nil
	}
}

// RetryIpBlacklistInvoker 用于流量过滤的IP黑名单导入失败后进行重新导入
func (c *CfwClient) RetryIpBlacklistInvoker(request *model.RetryIpBlacklistRequest) *RetryIpBlacklistInvoker {
	requestDef := GenReqDefForRetryIpBlacklist()
	return &RetryIpBlacklistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// DownloadExportResult 下载导出结果
//
// 下载导出结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DownloadExportResult(request *model.DownloadExportResultRequest) (*model.DownloadExportResultResponse, error) {
	requestDef := GenReqDefForDownloadExportResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadExportResultResponse), nil
	}
}

// DownloadExportResultInvoker 下载导出结果
func (c *CfwClient) DownloadExportResultInvoker(request *model.DownloadExportResultRequest) *DownloadExportResultInvoker {
	requestDef := GenReqDefForDownloadExportResult()
	return &DownloadExportResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadImportResult 下载导入规则
//
// 下载导入规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DownloadImportResult(request *model.DownloadImportResultRequest) (*model.DownloadImportResultResponse, error) {
	requestDef := GenReqDefForDownloadImportResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadImportResultResponse), nil
	}
}

// DownloadImportResultInvoker 下载导入规则
func (c *CfwClient) DownloadImportResultInvoker(request *model.DownloadImportResultRequest) *DownloadImportResultInvoker {
	requestDef := GenReqDefForDownloadImportResult()
	return &DownloadImportResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadImportTemplate 下载导入模板
//
// 下载导入模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DownloadImportTemplate(request *model.DownloadImportTemplateRequest) (*model.DownloadImportTemplateResponse, error) {
	requestDef := GenReqDefForDownloadImportTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadImportTemplateResponse), nil
	}
}

// DownloadImportTemplateInvoker 下载导入模板
func (c *CfwClient) DownloadImportTemplateInvoker(request *model.DownloadImportTemplateRequest) *DownloadImportTemplateInvoker {
	requestDef := GenReqDefForDownloadImportTemplate()
	return &DownloadImportTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportRuleAcl 导出访问控制规则
//
// 导出访问控制规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ExportRuleAcl(request *model.ExportRuleAclRequest) (*model.ExportRuleAclResponse, error) {
	requestDef := GenReqDefForExportRuleAcl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportRuleAclResponse), nil
	}
}

// ExportRuleAclInvoker 导出访问控制规则
func (c *CfwClient) ExportRuleAclInvoker(request *model.ExportRuleAclRequest) *ExportRuleAclInvoker {
	requestDef := GenReqDefForExportRuleAcl()
	return &ExportRuleAclInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportRuleAcl 导入访问控制规则
//
// 导入访问控制规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ImportRuleAcl(request *model.ImportRuleAclRequest) (*model.ImportRuleAclResponse, error) {
	requestDef := GenReqDefForImportRuleAcl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportRuleAclResponse), nil
	}
}

// ImportRuleAclInvoker 导入访问控制规则
func (c *CfwClient) ImportRuleAclInvoker(request *model.ImportRuleAclRequest) *ImportRuleAclInvoker {
	requestDef := GenReqDefForImportRuleAcl()
	return &ImportRuleAclInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListAclRuleHitStatus 获取规则的命中次数和最后一次命中时间
//
// 获取规则的命中次数和最后一次命中时间
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListAclRuleHitStatus(request *model.ListAclRuleHitStatusRequest) (*model.ListAclRuleHitStatusResponse, error) {
	requestDef := GenReqDefForListAclRuleHitStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAclRuleHitStatusResponse), nil
	}
}

// ListAclRuleHitStatusInvoker 获取规则的命中次数和最后一次命中时间
func (c *CfwClient) ListAclRuleHitStatusInvoker(request *model.ListAclRuleHitStatusRequest) *ListAclRuleHitStatusInvoker {
	requestDef := GenReqDefForListAclRuleHitStatus()
	return &ListAclRuleHitStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowExportStatus 查询导出结果
//
// 查询导出结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ShowExportStatus(request *model.ShowExportStatusRequest) (*model.ShowExportStatusResponse, error) {
	requestDef := GenReqDefForShowExportStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowExportStatusResponse), nil
	}
}

// ShowExportStatusInvoker 查询导出结果
func (c *CfwClient) ShowExportStatusInvoker(request *model.ShowExportStatusRequest) *ShowExportStatusInvoker {
	requestDef := GenReqDefForShowExportStatus()
	return &ShowExportStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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
// ACL防护规则优先级设置
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

// BatchDeleteAddressSets 批量删除地址组
//
// 批量删除地址组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) BatchDeleteAddressSets(request *model.BatchDeleteAddressSetsRequest) (*model.BatchDeleteAddressSetsResponse, error) {
	requestDef := GenReqDefForBatchDeleteAddressSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteAddressSetsResponse), nil
	}
}

// BatchDeleteAddressSetsInvoker 批量删除地址组
func (c *CfwClient) BatchDeleteAddressSetsInvoker(request *model.BatchDeleteAddressSetsRequest) *BatchDeleteAddressSetsInvoker {
	requestDef := GenReqDefForBatchDeleteAddressSets()
	return &BatchDeleteAddressSetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// UpdateObjectConfigDesc 更新对象配置描述
//
// 更新对象配置描述
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateObjectConfigDesc(request *model.UpdateObjectConfigDescRequest) (*model.UpdateObjectConfigDescResponse, error) {
	requestDef := GenReqDefForUpdateObjectConfigDesc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateObjectConfigDescResponse), nil
	}
}

// UpdateObjectConfigDescInvoker 更新对象配置描述
func (c *CfwClient) UpdateObjectConfigDescInvoker(request *model.UpdateObjectConfigDescRequest) *UpdateObjectConfigDescInvoker {
	requestDef := GenReqDefForUpdateObjectConfigDesc()
	return &UpdateObjectConfigDescInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// BatchCreateBlackWhiteList 批量添加黑白名单列表
//
// 批量添加黑白名单列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) BatchCreateBlackWhiteList(request *model.BatchCreateBlackWhiteListRequest) (*model.BatchCreateBlackWhiteListResponse, error) {
	requestDef := GenReqDefForBatchCreateBlackWhiteList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateBlackWhiteListResponse), nil
	}
}

// BatchCreateBlackWhiteListInvoker 批量添加黑白名单列表
func (c *CfwClient) BatchCreateBlackWhiteListInvoker(request *model.BatchCreateBlackWhiteListRequest) *BatchCreateBlackWhiteListInvoker {
	requestDef := GenReqDefForBatchCreateBlackWhiteList()
	return &BatchCreateBlackWhiteListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteBlackWhiteLists 批量删除黑白名单列表
//
// 批量删除黑白名单列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) BatchDeleteBlackWhiteLists(request *model.BatchDeleteBlackWhiteListsRequest) (*model.BatchDeleteBlackWhiteListsResponse, error) {
	requestDef := GenReqDefForBatchDeleteBlackWhiteLists()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteBlackWhiteListsResponse), nil
	}
}

// BatchDeleteBlackWhiteListsInvoker 批量删除黑白名单列表
func (c *CfwClient) BatchDeleteBlackWhiteListsInvoker(request *model.BatchDeleteBlackWhiteListsRequest) *BatchDeleteBlackWhiteListsInvoker {
	requestDef := GenReqDefForBatchDeleteBlackWhiteLists()
	return &BatchDeleteBlackWhiteListsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// AddCustomDnsServer 添加指定DNS服务器
//
// 添加指定DNS服务器
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) AddCustomDnsServer(request *model.AddCustomDnsServerRequest) (*model.AddCustomDnsServerResponse, error) {
	requestDef := GenReqDefForAddCustomDnsServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddCustomDnsServerResponse), nil
	}
}

// AddCustomDnsServerInvoker 添加指定DNS服务器
func (c *CfwClient) AddCustomDnsServerInvoker(request *model.AddCustomDnsServerRequest) *AddCustomDnsServerInvoker {
	requestDef := GenReqDefForAddCustomDnsServer()
	return &AddCustomDnsServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListDnsServers 查询DNS服务器列表
//
// 查询DNS服务器列表
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

// ListDnsServersInvoker 查询DNS服务器列表
func (c *CfwClient) ListDnsServersInvoker(request *model.ListDnsServersRequest) *ListDnsServersInvoker {
	requestDef := GenReqDefForListDnsServers()
	return &ListDnsServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDomainParseDetail 查询域名解析IP地址
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

// ListDomainParseDetailInvoker 查询域名解析IP地址
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

// ListDomainResolveIp 获取域名解析结果
//
// 获取域名解析结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListDomainResolveIp(request *model.ListDomainResolveIpRequest) (*model.ListDomainResolveIpResponse, error) {
	requestDef := GenReqDefForListDomainResolveIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainResolveIpResponse), nil
	}
}

// ListDomainResolveIpInvoker 获取域名解析结果
func (c *CfwClient) ListDomainResolveIpInvoker(request *model.ListDomainResolveIpRequest) *ListDomainResolveIpInvoker {
	requestDef := GenReqDefForListDomainResolveIp()
	return &ListDomainResolveIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// UpdateDnsServers 更新DNS服务器列表
//
// 更新DNS服务器列表
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

// UpdateDnsServersInvoker 更新DNS服务器列表
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

// AddEipAlarmWhitelist 添加EIP告警白名单
//
// 添加EIP告警白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) AddEipAlarmWhitelist(request *model.AddEipAlarmWhitelistRequest) (*model.AddEipAlarmWhitelistResponse, error) {
	requestDef := GenReqDefForAddEipAlarmWhitelist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddEipAlarmWhitelistResponse), nil
	}
}

// AddEipAlarmWhitelistInvoker 添加EIP告警白名单
func (c *CfwClient) AddEipAlarmWhitelistInvoker(request *model.AddEipAlarmWhitelistRequest) *AddEipAlarmWhitelistInvoker {
	requestDef := GenReqDefForAddEipAlarmWhitelist()
	return &AddEipAlarmWhitelistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListAlarmWhitelist 查看EIP告警白名单
//
// 查看EIP告警白名单
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

// ListAlarmWhitelistInvoker 查看EIP告警白名单
func (c *CfwClient) ListAlarmWhitelistInvoker(request *model.ListAlarmWhitelistRequest) *ListAlarmWhitelistInvoker {
	requestDef := GenReqDefForListAlarmWhitelist()
	return &ListAlarmWhitelistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEipCount 查询EIP数量
//
// 查询EIP数量
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

// ListEipCountInvoker 查询EIP数量
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

// ShowAutoProtectStatus 获取EIP自动防护状态信息
//
// 获取EIP自动防护状态信息
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

// ShowAutoProtectStatusInvoker 获取EIP自动防护状态信息
func (c *CfwClient) ShowAutoProtectStatusInvoker(request *model.ShowAutoProtectStatusRequest) *ShowAutoProtectStatusInvoker {
	requestDef := GenReqDefForShowAutoProtectStatus()
	return &ShowAutoProtectStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchAutoProtectStatus 修改EIP自动防护开关
//
// 修改EIP自动防护开关
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

// SwitchAutoProtectStatusInvoker 修改EIP自动防护开关
func (c *CfwClient) SwitchAutoProtectStatusInvoker(request *model.SwitchAutoProtectStatusRequest) *SwitchAutoProtectStatusInvoker {
	requestDef := GenReqDefForSwitchAutoProtectStatus()
	return &SwitchAutoProtectStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchFirewallEipProtection 一键逃生/一键恢复开关
//
// 一键逃生/一键恢复开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) SwitchFirewallEipProtection(request *model.SwitchFirewallEipProtectionRequest) (*model.SwitchFirewallEipProtectionResponse, error) {
	requestDef := GenReqDefForSwitchFirewallEipProtection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchFirewallEipProtectionResponse), nil
	}
}

// SwitchFirewallEipProtectionInvoker 一键逃生/一键恢复开关
func (c *CfwClient) SwitchFirewallEipProtectionInvoker(request *model.SwitchFirewallEipProtectionRequest) *SwitchFirewallEipProtectionInvoker {
	requestDef := GenReqDefForSwitchFirewallEipProtection()
	return &SwitchFirewallEipProtectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowConfigQuota 查询防火墙配额信息
//
// 查询防火墙配额信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ShowConfigQuota(request *model.ShowConfigQuotaRequest) (*model.ShowConfigQuotaResponse, error) {
	requestDef := GenReqDefForShowConfigQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConfigQuotaResponse), nil
	}
}

// ShowConfigQuotaInvoker 查询防火墙配额信息
func (c *CfwClient) ShowConfigQuotaInvoker(request *model.ShowConfigQuotaRequest) *ShowConfigQuotaInvoker {
	requestDef := GenReqDefForShowConfigQuota()
	return &ShowConfigQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSnFirewallProtectionStatus 查询南北向防火墙防护状态
//
// 查询南北向防火墙防护状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ShowSnFirewallProtectionStatus(request *model.ShowSnFirewallProtectionStatusRequest) (*model.ShowSnFirewallProtectionStatusResponse, error) {
	requestDef := GenReqDefForShowSnFirewallProtectionStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSnFirewallProtectionStatusResponse), nil
	}
}

// ShowSnFirewallProtectionStatusInvoker 查询南北向防火墙防护状态
func (c *CfwClient) ShowSnFirewallProtectionStatusInvoker(request *model.ShowSnFirewallProtectionStatusRequest) *ShowSnFirewallProtectionStatusInvoker {
	requestDef := GenReqDefForShowSnFirewallProtectionStatus()
	return &ShowSnFirewallProtectionStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFirewallName 更改防火墙名称
//
// 更改防火墙名称
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateFirewallName(request *model.UpdateFirewallNameRequest) (*model.UpdateFirewallNameResponse, error) {
	requestDef := GenReqDefForUpdateFirewallName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFirewallNameResponse), nil
	}
}

// UpdateFirewallNameInvoker 更改防火墙名称
func (c *CfwClient) UpdateFirewallNameInvoker(request *model.UpdateFirewallNameRequest) *UpdateFirewallNameInvoker {
	requestDef := GenReqDefForUpdateFirewallName()
	return &UpdateFirewallNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteCustomerIps 批量删除自定义IPS规则
//
// 批量删除自定义IPS规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) BatchDeleteCustomerIps(request *model.BatchDeleteCustomerIpsRequest) (*model.BatchDeleteCustomerIpsResponse, error) {
	requestDef := GenReqDefForBatchDeleteCustomerIps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteCustomerIpsResponse), nil
	}
}

// BatchDeleteCustomerIpsInvoker 批量删除自定义IPS规则
func (c *CfwClient) BatchDeleteCustomerIpsInvoker(request *model.BatchDeleteCustomerIpsRequest) *BatchDeleteCustomerIpsInvoker {
	requestDef := GenReqDefForBatchDeleteCustomerIps()
	return &BatchDeleteCustomerIpsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateCustomerIpsAction 批量更新自定义IPS规则的动作
//
// 批量更新自定义IPS规则的动作
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) BatchUpdateCustomerIpsAction(request *model.BatchUpdateCustomerIpsActionRequest) (*model.BatchUpdateCustomerIpsActionResponse, error) {
	requestDef := GenReqDefForBatchUpdateCustomerIpsAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateCustomerIpsActionResponse), nil
	}
}

// BatchUpdateCustomerIpsActionInvoker 批量更新自定义IPS规则的动作
func (c *CfwClient) BatchUpdateCustomerIpsActionInvoker(request *model.BatchUpdateCustomerIpsActionRequest) *BatchUpdateCustomerIpsActionInvoker {
	requestDef := GenReqDefForBatchUpdateCustomerIpsAction()
	return &BatchUpdateCustomerIpsActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ChangeIpsRuleMode 改变IPS规则模式
//
// 改变IPS规则模式
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

// ChangeIpsRuleModeInvoker 改变IPS规则模式
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

// CreateCustomerIps 创建自定义IPS规则
//
// 创建自定义IPS规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) CreateCustomerIps(request *model.CreateCustomerIpsRequest) (*model.CreateCustomerIpsResponse, error) {
	requestDef := GenReqDefForCreateCustomerIps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCustomerIpsResponse), nil
	}
}

// CreateCustomerIpsInvoker 创建自定义IPS规则
func (c *CfwClient) CreateCustomerIpsInvoker(request *model.CreateCustomerIpsRequest) *CreateCustomerIpsInvoker {
	requestDef := GenReqDefForCreateCustomerIps()
	return &CreateCustomerIpsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAdvancedIpsRules 查询频率IPS规则信息
//
// 查询频率IPS规则信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListAdvancedIpsRules(request *model.ListAdvancedIpsRulesRequest) (*model.ListAdvancedIpsRulesResponse, error) {
	requestDef := GenReqDefForListAdvancedIpsRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAdvancedIpsRulesResponse), nil
	}
}

// ListAdvancedIpsRulesInvoker 查询频率IPS规则信息
func (c *CfwClient) ListAdvancedIpsRulesInvoker(request *model.ListAdvancedIpsRulesRequest) *ListAdvancedIpsRulesInvoker {
	requestDef := GenReqDefForListAdvancedIpsRules()
	return &ListAdvancedIpsRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCustomerIps 查看自定义IPS规则列表
//
// 查看自定义IPS规则列表
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

// ListCustomerIpsInvoker 查看自定义IPS规则列表
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

// ListIpsRules 获取IPS规则列表
//
// 获取IPS规则列表
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

// ListIpsRulesInvoker 获取IPS规则列表
func (c *CfwClient) ListIpsRulesInvoker(request *model.ListIpsRulesRequest) *ListIpsRulesInvoker {
	requestDef := GenReqDefForListIpsRules()
	return &ListIpsRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowCustomerIpsInfo 查询自定义IPS规则详情
//
// 功能说明：自定义IPS规则详情，特性:根据路径输入的IPS ID查看对应的自定义IPS详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ShowCustomerIpsInfo(request *model.ShowCustomerIpsInfoRequest) (*model.ShowCustomerIpsInfoResponse, error) {
	requestDef := GenReqDefForShowCustomerIpsInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCustomerIpsInfoResponse), nil
	}
}

// ShowCustomerIpsInfoInvoker 查询自定义IPS规则详情
func (c *CfwClient) ShowCustomerIpsInfoInvoker(request *model.ShowCustomerIpsInfoRequest) *ShowCustomerIpsInfoInvoker {
	requestDef := GenReqDefForShowCustomerIpsInfo()
	return &ShowCustomerIpsInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIpsUpdateTime 获取IPS规则细节
//
// 获取IPS规则细节
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

// ShowIpsUpdateTimeInvoker 获取IPS规则细节
func (c *CfwClient) ShowIpsUpdateTimeInvoker(request *model.ShowIpsUpdateTimeRequest) *ShowIpsUpdateTimeInvoker {
	requestDef := GenReqDefForShowIpsUpdateTime()
	return &ShowIpsUpdateTimeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAdvancedIpsRule 创建频率IPS规则
//
// 创建频率IPS规则
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

// UpdateAdvancedIpsRuleInvoker 创建频率IPS规则
func (c *CfwClient) UpdateAdvancedIpsRuleInvoker(request *model.UpdateAdvancedIpsRuleRequest) *UpdateAdvancedIpsRuleInvoker {
	requestDef := GenReqDefForUpdateAdvancedIpsRule()
	return &UpdateAdvancedIpsRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCustomerIps 更新自定义IPS规则
//
// 更新自定义IPS规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateCustomerIps(request *model.UpdateCustomerIpsRequest) (*model.UpdateCustomerIpsResponse, error) {
	requestDef := GenReqDefForUpdateCustomerIps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCustomerIpsResponse), nil
	}
}

// UpdateCustomerIpsInvoker 更新自定义IPS规则
func (c *CfwClient) UpdateCustomerIpsInvoker(request *model.UpdateCustomerIpsRequest) *UpdateCustomerIpsInvoker {
	requestDef := GenReqDefForUpdateCustomerIps()
	return &UpdateCustomerIpsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateIpsWhitelist 批量添加IPS白名单
//
// 批量添加IPS白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) BatchCreateIpsWhitelist(request *model.BatchCreateIpsWhitelistRequest) (*model.BatchCreateIpsWhitelistResponse, error) {
	requestDef := GenReqDefForBatchCreateIpsWhitelist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateIpsWhitelistResponse), nil
	}
}

// BatchCreateIpsWhitelistInvoker 批量添加IPS白名单
func (c *CfwClient) BatchCreateIpsWhitelistInvoker(request *model.BatchCreateIpsWhitelistRequest) *BatchCreateIpsWhitelistInvoker {
	requestDef := GenReqDefForBatchCreateIpsWhitelist()
	return &BatchCreateIpsWhitelistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteIpsWhitelist 批量删除IPS白名单
//
// 批量删除IPS白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) BatchDeleteIpsWhitelist(request *model.BatchDeleteIpsWhitelistRequest) (*model.BatchDeleteIpsWhitelistResponse, error) {
	requestDef := GenReqDefForBatchDeleteIpsWhitelist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteIpsWhitelistResponse), nil
	}
}

// BatchDeleteIpsWhitelistInvoker 批量删除IPS白名单
func (c *CfwClient) BatchDeleteIpsWhitelistInvoker(request *model.BatchDeleteIpsWhitelistRequest) *BatchDeleteIpsWhitelistInvoker {
	requestDef := GenReqDefForBatchDeleteIpsWhitelist()
	return &BatchDeleteIpsWhitelistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIpsWhitelists 查询IPS白名单列表详情
//
// 查询IPS白名单列表详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListIpsWhitelists(request *model.ListIpsWhitelistsRequest) (*model.ListIpsWhitelistsResponse, error) {
	requestDef := GenReqDefForListIpsWhitelists()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIpsWhitelistsResponse), nil
	}
}

// ListIpsWhitelistsInvoker 查询IPS白名单列表详情
func (c *CfwClient) ListIpsWhitelistsInvoker(request *model.ListIpsWhitelistsRequest) *ListIpsWhitelistsInvoker {
	requestDef := GenReqDefForListIpsWhitelists()
	return &ListIpsWhitelistsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateIpsWhitelist 更新IPS白名单
//
// 更新IPS白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateIpsWhitelist(request *model.UpdateIpsWhitelistRequest) (*model.UpdateIpsWhitelistResponse, error) {
	requestDef := GenReqDefForUpdateIpsWhitelist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIpsWhitelistResponse), nil
	}
}

// UpdateIpsWhitelistInvoker 更新IPS白名单
func (c *CfwClient) UpdateIpsWhitelistInvoker(request *model.UpdateIpsWhitelistRequest) *UpdateIpsWhitelistInvoker {
	requestDef := GenReqDefForUpdateIpsWhitelist()
	return &UpdateIpsWhitelistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAttackStatistic 查询攻击统计
//
// 根据防火墙攻击日志，查询攻击统计信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListAttackStatistic(request *model.ListAttackStatisticRequest) (*model.ListAttackStatisticResponse, error) {
	requestDef := GenReqDefForListAttackStatistic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAttackStatisticResponse), nil
	}
}

// ListAttackStatisticInvoker 查询攻击统计
func (c *CfwClient) ListAttackStatisticInvoker(request *model.ListAttackStatisticRequest) *ListAttackStatisticInvoker {
	requestDef := GenReqDefForListAttackStatistic()
	return &ListAttackStatisticInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlowStatistic 查询流量日志统计
//
// 查询流量日志统计
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListFlowStatistic(request *model.ListFlowStatisticRequest) (*model.ListFlowStatisticResponse, error) {
	requestDef := GenReqDefForListFlowStatistic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlowStatisticResponse), nil
	}
}

// ListFlowStatisticInvoker 查询流量日志统计
func (c *CfwClient) ListFlowStatisticInvoker(request *model.ListFlowStatisticRequest) *ListFlowStatisticInvoker {
	requestDef := GenReqDefForListFlowStatistic()
	return &ListFlowStatisticInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAccessDetail 查询访问控制统计详情
//
// 查询访问控制统计详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ShowAccessDetail(request *model.ShowAccessDetailRequest) (*model.ShowAccessDetailResponse, error) {
	requestDef := GenReqDefForShowAccessDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAccessDetailResponse), nil
	}
}

// ShowAccessDetailInvoker 查询访问控制统计详情
func (c *CfwClient) ShowAccessDetailInvoker(request *model.ShowAccessDetailRequest) *ShowAccessDetailInvoker {
	requestDef := GenReqDefForShowAccessDetail()
	return &ShowAccessDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAccessTop 查询访问日志统计信息
//
// 获取访问日志的TOP统计信息，如TOP命中规则等
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ShowAccessTop(request *model.ShowAccessTopRequest) (*model.ShowAccessTopResponse, error) {
	requestDef := GenReqDefForShowAccessTop()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAccessTopResponse), nil
	}
}

// ShowAccessTopInvoker 查询访问日志统计信息
func (c *CfwClient) ShowAccessTopInvoker(request *model.ShowAccessTopRequest) *ShowAccessTopInvoker {
	requestDef := GenReqDefForShowAccessTop()
	return &ShowAccessTopInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAttackDetail 查询攻击日志统计详情
//
// 查询攻击日志统计详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ShowAttackDetail(request *model.ShowAttackDetailRequest) (*model.ShowAttackDetailResponse, error) {
	requestDef := GenReqDefForShowAttackDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAttackDetailResponse), nil
	}
}

// ShowAttackDetailInvoker 查询攻击日志统计详情
func (c *CfwClient) ShowAttackDetailInvoker(request *model.ShowAttackDetailRequest) *ShowAttackDetailInvoker {
	requestDef := GenReqDefForShowAttackDetail()
	return &ShowAttackDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAttackTop 查询攻击日志TOP统计
//
// 查询攻击日志TOP统计
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ShowAttackTop(request *model.ShowAttackTopRequest) (*model.ShowAttackTopResponse, error) {
	requestDef := GenReqDefForShowAttackTop()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAttackTopResponse), nil
	}
}

// ShowAttackTopInvoker 查询攻击日志TOP统计
func (c *CfwClient) ShowAttackTopInvoker(request *model.ShowAttackTopRequest) *ShowAttackTopInvoker {
	requestDef := GenReqDefForShowAttackTop()
	return &ShowAttackTopInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAttackTotal 查询攻击概览
//
// 查询攻击概览
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ShowAttackTotal(request *model.ShowAttackTotalRequest) (*model.ShowAttackTotalResponse, error) {
	requestDef := GenReqDefForShowAttackTotal()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAttackTotalResponse), nil
	}
}

// ShowAttackTotalInvoker 查询攻击概览
func (c *CfwClient) ShowAttackTotalInvoker(request *model.ShowAttackTotalRequest) *ShowAttackTotalInvoker {
	requestDef := GenReqDefForShowAttackTotal()
	return &ShowAttackTotalInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAttackTrend 查询攻击趋势
//
// 查询攻击趋势
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ShowAttackTrend(request *model.ShowAttackTrendRequest) (*model.ShowAttackTrendResponse, error) {
	requestDef := GenReqDefForShowAttackTrend()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAttackTrendResponse), nil
	}
}

// ShowAttackTrendInvoker 查询攻击趋势
func (c *CfwClient) ShowAttackTrendInvoker(request *model.ShowAttackTrendRequest) *ShowAttackTrendInvoker {
	requestDef := GenReqDefForShowAttackTrend()
	return &ShowAttackTrendInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFlowDetail 查询流量日志统计详情
//
// 查询流量日志统计详情，如统计某个源IP访问详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ShowFlowDetail(request *model.ShowFlowDetailRequest) (*model.ShowFlowDetailResponse, error) {
	requestDef := GenReqDefForShowFlowDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFlowDetailResponse), nil
	}
}

// ShowFlowDetailInvoker 查询流量日志统计详情
func (c *CfwClient) ShowFlowDetailInvoker(request *model.ShowFlowDetailRequest) *ShowFlowDetailInvoker {
	requestDef := GenReqDefForShowFlowDetail()
	return &ShowFlowDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFlowTop 查询流量TOP统计
//
// 查询流量TOP统计
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ShowFlowTop(request *model.ShowFlowTopRequest) (*model.ShowFlowTopResponse, error) {
	requestDef := GenReqDefForShowFlowTop()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFlowTopResponse), nil
	}
}

// ShowFlowTopInvoker 查询流量TOP统计
func (c *CfwClient) ShowFlowTopInvoker(request *model.ShowFlowTopRequest) *ShowFlowTopInvoker {
	requestDef := GenReqDefForShowFlowTop()
	return &ShowFlowTopInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFlowTrend 查询会话趋势
//
// 查询会话趋势
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ShowFlowTrend(request *model.ShowFlowTrendRequest) (*model.ShowFlowTrendResponse, error) {
	requestDef := GenReqDefForShowFlowTrend()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFlowTrendResponse), nil
	}
}

// ShowFlowTrendInvoker 查询会话趋势
func (c *CfwClient) ShowFlowTrendInvoker(request *model.ShowFlowTrendRequest) *ShowFlowTrendInvoker {
	requestDef := GenReqDefForShowFlowTrend()
	return &ShowFlowTrendInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLogsCount 查询日志数量
//
// 统计日志数量，如统计风险IP的数量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ShowLogsCount(request *model.ShowLogsCountRequest) (*model.ShowLogsCountResponse, error) {
	requestDef := GenReqDefForShowLogsCount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLogsCountResponse), nil
	}
}

// ShowLogsCountInvoker 查询日志数量
func (c *CfwClient) ShowLogsCountInvoker(request *model.ShowLogsCountRequest) *ShowLogsCountInvoker {
	requestDef := GenReqDefForShowLogsCount()
	return &ShowLogsCountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTrafficTrend 查询流量趋势
//
// 查询流量趋势，包括南北向、EIP、东西向的流量趋势
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ShowTrafficTrend(request *model.ShowTrafficTrendRequest) (*model.ShowTrafficTrendResponse, error) {
	requestDef := GenReqDefForShowTrafficTrend()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTrafficTrendResponse), nil
	}
}

// ShowTrafficTrendInvoker 查询流量趋势
func (c *CfwClient) ShowTrafficTrendInvoker(request *model.ShowTrafficTrendRequest) *ShowTrafficTrendInvoker {
	requestDef := GenReqDefForShowTrafficTrend()
	return &ShowTrafficTrendInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ExportLogs 导出防火墙日志
//
// 导出防火墙日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ExportLogs(request *model.ExportLogsRequest) (*model.ExportLogsResponse, error) {
	requestDef := GenReqDefForExportLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportLogsResponse), nil
	}
}

// ExportLogsInvoker 导出防火墙日志
func (c *CfwClient) ExportLogsInvoker(request *model.ExportLogsRequest) *ExportLogsInvoker {
	requestDef := GenReqDefForExportLogs()
	return &ExportLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListLogs 查询防火墙日志
//
// 查询防火墙日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListLogs(request *model.ListLogsRequest) (*model.ListLogsResponse, error) {
	requestDef := GenReqDefForListLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLogsResponse), nil
	}
}

// ListLogsInvoker 查询防火墙日志
func (c *CfwClient) ListLogsInvoker(request *model.ListLogsRequest) *ListLogsInvoker {
	requestDef := GenReqDefForListLogs()
	return &ListLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// BatchAddAccounts 批量添加账号
//
// 批量添加账号
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) BatchAddAccounts(request *model.BatchAddAccountsRequest) (*model.BatchAddAccountsResponse, error) {
	requestDef := GenReqDefForBatchAddAccounts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddAccountsResponse), nil
	}
}

// BatchAddAccountsInvoker 批量添加账号
func (c *CfwClient) BatchAddAccountsInvoker(request *model.BatchAddAccountsRequest) *BatchAddAccountsInvoker {
	requestDef := GenReqDefForBatchAddAccounts()
	return &BatchAddAccountsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchRemoveAccounts 批量移除账号
//
// 批量移除账号
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) BatchRemoveAccounts(request *model.BatchRemoveAccountsRequest) (*model.BatchRemoveAccountsResponse, error) {
	requestDef := GenReqDefForBatchRemoveAccounts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRemoveAccountsResponse), nil
	}
}

// BatchRemoveAccountsInvoker 批量移除账号
func (c *CfwClient) BatchRemoveAccountsInvoker(request *model.BatchRemoveAccountsRequest) *BatchRemoveAccountsInvoker {
	requestDef := GenReqDefForBatchRemoveAccounts()
	return &BatchRemoveAccountsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableMultiAccount 开启多账号管理
//
// 开启多账号管理
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) EnableMultiAccount(request *model.EnableMultiAccountRequest) (*model.EnableMultiAccountResponse, error) {
	requestDef := GenReqDefForEnableMultiAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableMultiAccountResponse), nil
	}
}

// EnableMultiAccountInvoker 开启多账号管理
func (c *CfwClient) EnableMultiAccountInvoker(request *model.EnableMultiAccountRequest) *EnableMultiAccountInvoker {
	requestDef := GenReqDefForEnableMultiAccount()
	return &EnableMultiAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccounts 查询账号列表
//
// 查询账号列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListAccounts(request *model.ListAccountsRequest) (*model.ListAccountsResponse, error) {
	requestDef := GenReqDefForListAccounts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAccountsResponse), nil
	}
}

// ListAccountsInvoker 查询账号列表
func (c *CfwClient) ListAccountsInvoker(request *model.ListAccountsRequest) *ListAccountsInvoker {
	requestDef := GenReqDefForListAccounts()
	return &ListAccountsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOrganizationAccounts 查询组织账号列表
//
// 查询组织账号列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListOrganizationAccounts(request *model.ListOrganizationAccountsRequest) (*model.ListOrganizationAccountsResponse, error) {
	requestDef := GenReqDefForListOrganizationAccounts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOrganizationAccountsResponse), nil
	}
}

// ListOrganizationAccountsInvoker 查询组织账号列表
func (c *CfwClient) ListOrganizationAccountsInvoker(request *model.ListOrganizationAccountsRequest) *ListOrganizationAccountsInvoker {
	requestDef := GenReqDefForListOrganizationAccounts()
	return &ListOrganizationAccountsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOrganizationTree 查询组织结构
//
// 查询组织结构
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListOrganizationTree(request *model.ListOrganizationTreeRequest) (*model.ListOrganizationTreeResponse, error) {
	requestDef := GenReqDefForListOrganizationTree()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOrganizationTreeResponse), nil
	}
}

// ListOrganizationTreeInvoker 查询组织结构
func (c *CfwClient) ListOrganizationTreeInvoker(request *model.ListOrganizationTreeRequest) *ListOrganizationTreeInvoker {
	requestDef := GenReqDefForListOrganizationTree()
	return &ListOrganizationTreeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteSchedules 批量删除时间表
//
// 批量删除时间表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) BatchDeleteSchedules(request *model.BatchDeleteSchedulesRequest) (*model.BatchDeleteSchedulesResponse, error) {
	requestDef := GenReqDefForBatchDeleteSchedules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteSchedulesResponse), nil
	}
}

// BatchDeleteSchedulesInvoker 批量删除时间表
func (c *CfwClient) BatchDeleteSchedulesInvoker(request *model.BatchDeleteSchedulesRequest) *BatchDeleteSchedulesInvoker {
	requestDef := GenReqDefForBatchDeleteSchedules()
	return &BatchDeleteSchedulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSchedule 创建时间表
//
// 创建时间表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) CreateSchedule(request *model.CreateScheduleRequest) (*model.CreateScheduleResponse, error) {
	requestDef := GenReqDefForCreateSchedule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateScheduleResponse), nil
	}
}

// CreateScheduleInvoker 创建时间表
func (c *CfwClient) CreateScheduleInvoker(request *model.CreateScheduleRequest) *CreateScheduleInvoker {
	requestDef := GenReqDefForCreateSchedule()
	return &CreateScheduleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSchedule 删除时间表
//
// 删除时间表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteSchedule(request *model.DeleteScheduleRequest) (*model.DeleteScheduleResponse, error) {
	requestDef := GenReqDefForDeleteSchedule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteScheduleResponse), nil
	}
}

// DeleteScheduleInvoker 删除时间表
func (c *CfwClient) DeleteScheduleInvoker(request *model.DeleteScheduleRequest) *DeleteScheduleInvoker {
	requestDef := GenReqDefForDeleteSchedule()
	return &DeleteScheduleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSchedules 查询时间表列表
//
// 查询时间表列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListSchedules(request *model.ListSchedulesRequest) (*model.ListSchedulesResponse, error) {
	requestDef := GenReqDefForListSchedules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSchedulesResponse), nil
	}
}

// ListSchedulesInvoker 查询时间表列表
func (c *CfwClient) ListSchedulesInvoker(request *model.ListSchedulesRequest) *ListSchedulesInvoker {
	requestDef := GenReqDefForListSchedules()
	return &ListSchedulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSchedule 更新时间表
//
// 更新时间表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateSchedule(request *model.UpdateScheduleRequest) (*model.UpdateScheduleResponse, error) {
	requestDef := GenReqDefForUpdateSchedule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateScheduleResponse), nil
	}
}

// UpdateScheduleInvoker 更新时间表
func (c *CfwClient) UpdateScheduleInvoker(request *model.UpdateScheduleRequest) *UpdateScheduleInvoker {
	requestDef := GenReqDefForUpdateSchedule()
	return &UpdateScheduleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateReportProfile 创建安全报告模板
//
// 创建安全报告模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) CreateReportProfile(request *model.CreateReportProfileRequest) (*model.CreateReportProfileResponse, error) {
	requestDef := GenReqDefForCreateReportProfile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateReportProfileResponse), nil
	}
}

// CreateReportProfileInvoker 创建安全报告模板
func (c *CfwClient) CreateReportProfileInvoker(request *model.CreateReportProfileRequest) *CreateReportProfileInvoker {
	requestDef := GenReqDefForCreateReportProfile()
	return &CreateReportProfileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteReportProfile 删除安全报告模板
//
// 删除安全报告模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteReportProfile(request *model.DeleteReportProfileRequest) (*model.DeleteReportProfileResponse, error) {
	requestDef := GenReqDefForDeleteReportProfile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteReportProfileResponse), nil
	}
}

// DeleteReportProfileInvoker 删除安全报告模板
func (c *CfwClient) DeleteReportProfileInvoker(request *model.DeleteReportProfileRequest) *DeleteReportProfileInvoker {
	requestDef := GenReqDefForDeleteReportProfile()
	return &DeleteReportProfileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListReportHistory 查询安全报告发送历史
//
// 查询安全报告发送历史
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListReportHistory(request *model.ListReportHistoryRequest) (*model.ListReportHistoryResponse, error) {
	requestDef := GenReqDefForListReportHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListReportHistoryResponse), nil
	}
}

// ListReportHistoryInvoker 查询安全报告发送历史
func (c *CfwClient) ListReportHistoryInvoker(request *model.ListReportHistoryRequest) *ListReportHistoryInvoker {
	requestDef := GenReqDefForListReportHistory()
	return &ListReportHistoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListReportProfiles 查询安全报告模板列表
//
// 查询安全报告模板列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListReportProfiles(request *model.ListReportProfilesRequest) (*model.ListReportProfilesResponse, error) {
	requestDef := GenReqDefForListReportProfiles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListReportProfilesResponse), nil
	}
}

// ListReportProfilesInvoker 查询安全报告模板列表
func (c *CfwClient) ListReportProfilesInvoker(request *model.ListReportProfilesRequest) *ListReportProfilesInvoker {
	requestDef := GenReqDefForListReportProfiles()
	return &ListReportProfilesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFirewallReport 查询安全报告
//
// 查询安全报告
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ShowFirewallReport(request *model.ShowFirewallReportRequest) (*model.ShowFirewallReportResponse, error) {
	requestDef := GenReqDefForShowFirewallReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFirewallReportResponse), nil
	}
}

// ShowFirewallReportInvoker 查询安全报告
func (c *CfwClient) ShowFirewallReportInvoker(request *model.ShowFirewallReportRequest) *ShowFirewallReportInvoker {
	requestDef := GenReqDefForShowFirewallReport()
	return &ShowFirewallReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowReportProfile 获取安全报告模板
//
// 获取安全报告模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ShowReportProfile(request *model.ShowReportProfileRequest) (*model.ShowReportProfileResponse, error) {
	requestDef := GenReqDefForShowReportProfile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReportProfileResponse), nil
	}
}

// ShowReportProfileInvoker 获取安全报告模板
func (c *CfwClient) ShowReportProfileInvoker(request *model.ShowReportProfileRequest) *ShowReportProfileInvoker {
	requestDef := GenReqDefForShowReportProfile()
	return &ShowReportProfileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateReportProfile 更新安全报告模板
//
// 更新安全报告模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateReportProfile(request *model.UpdateReportProfileRequest) (*model.UpdateReportProfileResponse, error) {
	requestDef := GenReqDefForUpdateReportProfile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateReportProfileResponse), nil
	}
}

// UpdateReportProfileInvoker 更新安全报告模板
func (c *CfwClient) UpdateReportProfileInvoker(request *model.UpdateReportProfileRequest) *UpdateReportProfileInvoker {
	requestDef := GenReqDefForUpdateReportProfile()
	return &UpdateReportProfileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// BatchDeleteServiceSets 批量删除服务组
//
// 批量删除服务组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) BatchDeleteServiceSets(request *model.BatchDeleteServiceSetsRequest) (*model.BatchDeleteServiceSetsResponse, error) {
	requestDef := GenReqDefForBatchDeleteServiceSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteServiceSetsResponse), nil
	}
}

// BatchDeleteServiceSetsInvoker 批量删除服务组
func (c *CfwClient) BatchDeleteServiceSetsInvoker(request *model.BatchDeleteServiceSetsRequest) *BatchDeleteServiceSetsInvoker {
	requestDef := GenReqDefForBatchDeleteServiceSets()
	return &BatchDeleteServiceSetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteServiceItem 删除服务组成员
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

// DeleteServiceItemInvoker 删除服务组成员
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

// BatchCreatePrivateNetworkSegments 创建私网网段
//
// 添加私网网段的接口，添加私网网段后该网段的流量将引流到VPC防火墙防护
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) BatchCreatePrivateNetworkSegments(request *model.BatchCreatePrivateNetworkSegmentsRequest) (*model.BatchCreatePrivateNetworkSegmentsResponse, error) {
	requestDef := GenReqDefForBatchCreatePrivateNetworkSegments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreatePrivateNetworkSegmentsResponse), nil
	}
}

// BatchCreatePrivateNetworkSegmentsInvoker 创建私网网段
func (c *CfwClient) BatchCreatePrivateNetworkSegmentsInvoker(request *model.BatchCreatePrivateNetworkSegmentsRequest) *BatchCreatePrivateNetworkSegmentsInvoker {
	requestDef := GenReqDefForBatchCreatePrivateNetworkSegments()
	return &BatchCreatePrivateNetworkSegmentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeletePrivateNetworkSegments 删除私网网段信息
//
// 删除私网网段的接口，根据用户输入的私网网段conf_id，删除对应ID的私网网段
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) BatchDeletePrivateNetworkSegments(request *model.BatchDeletePrivateNetworkSegmentsRequest) (*model.BatchDeletePrivateNetworkSegmentsResponse, error) {
	requestDef := GenReqDefForBatchDeletePrivateNetworkSegments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeletePrivateNetworkSegmentsResponse), nil
	}
}

// BatchDeletePrivateNetworkSegmentsInvoker 删除私网网段信息
func (c *CfwClient) BatchDeletePrivateNetworkSegmentsInvoker(request *model.BatchDeletePrivateNetworkSegmentsRequest) *BatchDeletePrivateNetworkSegmentsInvoker {
	requestDef := GenReqDefForBatchDeletePrivateNetworkSegments()
	return &BatchDeletePrivateNetworkSegmentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeEastWestFirewallStatus 更新VPC间防火墙防护状态
//
// 更新VPC间防火墙防护状态
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

// ChangeEastWestFirewallStatusInvoker 更新VPC间防火墙防护状态
func (c *CfwClient) ChangeEastWestFirewallStatusInvoker(request *model.ChangeEastWestFirewallStatusRequest) *ChangeEastWestFirewallStatusInvoker {
	requestDef := GenReqDefForChangeEastWestFirewallStatus()
	return &ChangeEastWestFirewallStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPrivateNetworkSegments 获取私网网段的信息
//
// 东西向私网网段查询接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListPrivateNetworkSegments(request *model.ListPrivateNetworkSegmentsRequest) (*model.ListPrivateNetworkSegmentsResponse, error) {
	requestDef := GenReqDefForListPrivateNetworkSegments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPrivateNetworkSegmentsResponse), nil
	}
}

// ListPrivateNetworkSegmentsInvoker 获取私网网段的信息
func (c *CfwClient) ListPrivateNetworkSegmentsInvoker(request *model.ListPrivateNetworkSegmentsRequest) *ListPrivateNetworkSegmentsInvoker {
	requestDef := GenReqDefForListPrivateNetworkSegments()
	return &ListPrivateNetworkSegmentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEwAssociatedEr 查询VPC间防火墙使用的企业路由器信息
//
// 查询VPC间防火墙使用的企业路由器信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ShowEwAssociatedEr(request *model.ShowEwAssociatedErRequest) (*model.ShowEwAssociatedErResponse, error) {
	requestDef := GenReqDefForShowEwAssociatedEr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEwAssociatedErResponse), nil
	}
}

// ShowEwAssociatedErInvoker 查询VPC间防火墙使用的企业路由器信息
func (c *CfwClient) ShowEwAssociatedErInvoker(request *model.ShowEwAssociatedErRequest) *ShowEwAssociatedErInvoker {
	requestDef := GenReqDefForShowEwAssociatedEr()
	return &ShowEwAssociatedErInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEwAssociatedVpc 查询VPC边界防火墙使用的引流VPC信息
//
// 查询VPC边界防火墙使用的引流VPC信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ShowEwAssociatedVpc(request *model.ShowEwAssociatedVpcRequest) (*model.ShowEwAssociatedVpcResponse, error) {
	requestDef := GenReqDefForShowEwAssociatedVpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEwAssociatedVpcResponse), nil
	}
}

// ShowEwAssociatedVpcInvoker 查询VPC边界防火墙使用的引流VPC信息
func (c *CfwClient) ShowEwAssociatedVpcInvoker(request *model.ShowEwAssociatedVpcRequest) *ShowEwAssociatedVpcInvoker {
	requestDef := GenReqDefForShowEwAssociatedVpc()
	return &ShowEwAssociatedVpcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePrivateNetworkSegment 更新私网网段
//
// 更新私网网段的REST接口，主要特性：
//  * 根据私网网段ID，更新对应的私网网段信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdatePrivateNetworkSegment(request *model.UpdatePrivateNetworkSegmentRequest) (*model.UpdatePrivateNetworkSegmentResponse, error) {
	requestDef := GenReqDefForUpdatePrivateNetworkSegment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePrivateNetworkSegmentResponse), nil
	}
}

// UpdatePrivateNetworkSegmentInvoker 更新私网网段
func (c *CfwClient) UpdatePrivateNetworkSegmentInvoker(request *model.UpdatePrivateNetworkSegmentRequest) *UpdatePrivateNetworkSegmentInvoker {
	requestDef := GenReqDefForUpdatePrivateNetworkSegment()
	return &UpdatePrivateNetworkSegmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
