package v3

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codeartsinspector/v3/model"
)

type CodeArtsInspectorClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCodeArtsInspectorClient(hcClient *httpclient.HcHttpClient) *CodeArtsInspectorClient {
	return &CodeArtsInspectorClient{HcClient: hcClient}
}

func CodeArtsInspectorClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AddGroup 批量创建主机组
//
// 批量创建主机组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsInspectorClient) AddGroup(request *model.AddGroupRequest) (*model.AddGroupResponse, error) {
	requestDef := GenReqDefForAddGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddGroupResponse), nil
	}
}

// AddGroupInvoker 批量创建主机组
func (c *CodeArtsInspectorClient) AddGroupInvoker(request *model.AddGroupRequest) *AddGroupInvoker {
	requestDef := GenReqDefForAddGroup()
	return &AddGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGroup 删除主机组
//
// 删除主机组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsInspectorClient) DeleteGroup(request *model.DeleteGroupRequest) (*model.DeleteGroupResponse, error) {
	requestDef := GenReqDefForDeleteGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGroupResponse), nil
	}
}

// DeleteGroupInvoker 删除主机组
func (c *CodeArtsInspectorClient) DeleteGroupInvoker(request *model.DeleteGroupRequest) *DeleteGroupInvoker {
	requestDef := GenReqDefForDeleteGroup()
	return &DeleteGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroups 获取主机组列表
//
// 获取主机组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsInspectorClient) ListGroups(request *model.ListGroupsRequest) (*model.ListGroupsResponse, error) {
	requestDef := GenReqDefForListGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupsResponse), nil
	}
}

// ListGroupsInvoker 获取主机组列表
func (c *CodeArtsInspectorClient) ListGroupsInvoker(request *model.ListGroupsRequest) *ListGroupsInvoker {
	requestDef := GenReqDefForListGroups()
	return &ListGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHostResults 获取主机漏洞扫描结果
//
// 获取主机漏洞扫描结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsInspectorClient) ListHostResults(request *model.ListHostResultsRequest) (*model.ListHostResultsResponse, error) {
	requestDef := GenReqDefForListHostResults()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostResultsResponse), nil
	}
}

// ListHostResultsInvoker 获取主机漏洞扫描结果
func (c *CodeArtsInspectorClient) ListHostResultsInvoker(request *model.ListHostResultsRequest) *ListHostResultsInvoker {
	requestDef := GenReqDefForListHostResults()
	return &ListHostResultsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchStartHostTasks 批量启动或取消主机扫描任务
//
// 批量启动或取消主机漏洞扫描任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsInspectorClient) BatchStartHostTasks(request *model.BatchStartHostTasksRequest) (*model.BatchStartHostTasksResponse, error) {
	requestDef := GenReqDefForBatchStartHostTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStartHostTasksResponse), nil
	}
}

// BatchStartHostTasksInvoker 批量启动或取消主机扫描任务
func (c *CodeArtsInspectorClient) BatchStartHostTasksInvoker(request *model.BatchStartHostTasksRequest) *BatchStartHostTasksInvoker {
	requestDef := GenReqDefForBatchStartHostTasks()
	return &BatchStartHostTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateHosts 批量创建主机资产
//
// 批量创建租户的主机资产
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsInspectorClient) BatchCreateHosts(request *model.BatchCreateHostsRequest) (*model.BatchCreateHostsResponse, error) {
	requestDef := GenReqDefForBatchCreateHosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateHostsResponse), nil
	}
}

// BatchCreateHostsInvoker 批量创建主机资产
func (c *CodeArtsInspectorClient) BatchCreateHostsInvoker(request *model.BatchCreateHostsRequest) *BatchCreateHostsInvoker {
	requestDef := GenReqDefForBatchCreateHosts()
	return &BatchCreateHostsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHost 删除主机资产
//
// 删除租户的主机资产
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsInspectorClient) DeleteHost(request *model.DeleteHostRequest) (*model.DeleteHostResponse, error) {
	requestDef := GenReqDefForDeleteHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHostResponse), nil
	}
}

// DeleteHostInvoker 删除主机资产
func (c *CodeArtsInspectorClient) DeleteHostInvoker(request *model.DeleteHostRequest) *DeleteHostInvoker {
	requestDef := GenReqDefForDeleteHost()
	return &DeleteHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHosts 获取主机资产
//
// 获取租户的主机资产列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsInspectorClient) ListHosts(request *model.ListHostsRequest) (*model.ListHostsResponse, error) {
	requestDef := GenReqDefForListHosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostsResponse), nil
	}
}

// ListHostsInvoker 获取主机资产
func (c *CodeArtsInspectorClient) ListHostsInvoker(request *model.ListHostsRequest) *ListHostsInvoker {
	requestDef := GenReqDefForListHosts()
	return &ListHostsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadTaskReport 下载网站扫描报告
//
// 下载网站扫描任务PDF报告
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsInspectorClient) DownloadTaskReport(request *model.DownloadTaskReportRequest) (*model.DownloadTaskReportResponse, error) {
	requestDef := GenReqDefForDownloadTaskReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadTaskReportResponse), nil
	}
}

// DownloadTaskReportInvoker 下载网站扫描报告
func (c *CodeArtsInspectorClient) DownloadTaskReportInvoker(request *model.DownloadTaskReportRequest) *DownloadTaskReportInvoker {
	requestDef := GenReqDefForDownloadTaskReport()
	return &DownloadTaskReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteGenerateReport 生成网站扫描报告
//
// 生成网站扫描PDF报告
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsInspectorClient) ExecuteGenerateReport(request *model.ExecuteGenerateReportRequest) (*model.ExecuteGenerateReportResponse, error) {
	requestDef := GenReqDefForExecuteGenerateReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteGenerateReportResponse), nil
	}
}

// ExecuteGenerateReportInvoker 生成网站扫描报告
func (c *CodeArtsInspectorClient) ExecuteGenerateReportInvoker(request *model.ExecuteGenerateReportRequest) *ExecuteGenerateReportInvoker {
	requestDef := GenReqDefForExecuteGenerateReport()
	return &ExecuteGenerateReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBusinessRisks 获取网站业务风险扫描结果
//
// 获取网站业务风险扫描结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsInspectorClient) ListBusinessRisks(request *model.ListBusinessRisksRequest) (*model.ListBusinessRisksResponse, error) {
	requestDef := GenReqDefForListBusinessRisks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBusinessRisksResponse), nil
	}
}

// ListBusinessRisksInvoker 获取网站业务风险扫描结果
func (c *CodeArtsInspectorClient) ListBusinessRisksInvoker(request *model.ListBusinessRisksRequest) *ListBusinessRisksInvoker {
	requestDef := GenReqDefForListBusinessRisks()
	return &ListBusinessRisksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPortResults 获取网站端口扫描结果
//
// 获取网站端口扫描结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsInspectorClient) ListPortResults(request *model.ListPortResultsRequest) (*model.ListPortResultsResponse, error) {
	requestDef := GenReqDefForListPortResults()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPortResultsResponse), nil
	}
}

// ListPortResultsInvoker 获取网站端口扫描结果
func (c *CodeArtsInspectorClient) ListPortResultsInvoker(request *model.ListPortResultsRequest) *ListPortResultsInvoker {
	requestDef := GenReqDefForListPortResults()
	return &ListPortResultsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowReportStatus 获取网站扫描报告状态
//
// 获取网站扫描PDF报告生成状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsInspectorClient) ShowReportStatus(request *model.ShowReportStatusRequest) (*model.ShowReportStatusResponse, error) {
	requestDef := GenReqDefForShowReportStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReportStatusResponse), nil
	}
}

// ShowReportStatusInvoker 获取网站扫描报告状态
func (c *CodeArtsInspectorClient) ShowReportStatusInvoker(request *model.ShowReportStatusRequest) *ShowReportStatusInvoker {
	requestDef := GenReqDefForShowReportStatus()
	return &ShowReportStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResults 获取网站扫描结果
//
// 获取网站漏洞扫描结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsInspectorClient) ShowResults(request *model.ShowResultsRequest) (*model.ShowResultsResponse, error) {
	requestDef := GenReqDefForShowResults()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResultsResponse), nil
	}
}

// ShowResultsInvoker 获取网站扫描结果
func (c *CodeArtsInspectorClient) ShowResultsInvoker(request *model.ShowResultsRequest) *ShowResultsInvoker {
	requestDef := GenReqDefForShowResults()
	return &ShowResultsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFalsePositive 更新网站漏洞的误报状态
//
// 更新网站扫描漏洞的误报状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsInspectorClient) UpdateFalsePositive(request *model.UpdateFalsePositiveRequest) (*model.UpdateFalsePositiveResponse, error) {
	requestDef := GenReqDefForUpdateFalsePositive()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFalsePositiveResponse), nil
	}
}

// UpdateFalsePositiveInvoker 更新网站漏洞的误报状态
func (c *CodeArtsInspectorClient) UpdateFalsePositiveInvoker(request *model.UpdateFalsePositiveRequest) *UpdateFalsePositiveInvoker {
	requestDef := GenReqDefForUpdateFalsePositive()
	return &UpdateFalsePositiveInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelTasks 取消或重启网站扫描任务
//
// 取消或重启网站漏洞扫描任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsInspectorClient) CancelTasks(request *model.CancelTasksRequest) (*model.CancelTasksResponse, error) {
	requestDef := GenReqDefForCancelTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelTasksResponse), nil
	}
}

// CancelTasksInvoker 取消或重启网站扫描任务
func (c *CodeArtsInspectorClient) CancelTasksInvoker(request *model.CancelTasksRequest) *CancelTasksInvoker {
	requestDef := GenReqDefForCancelTasks()
	return &CancelTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTasks 创建网站扫描任务并启动
//
// 创建网站漏洞扫描任务并启动
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsInspectorClient) CreateTasks(request *model.CreateTasksRequest) (*model.CreateTasksResponse, error) {
	requestDef := GenReqDefForCreateTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTasksResponse), nil
	}
}

// CreateTasksInvoker 创建网站扫描任务并启动
func (c *CodeArtsInspectorClient) CreateTasksInvoker(request *model.CreateTasksRequest) *CreateTasksInvoker {
	requestDef := GenReqDefForCreateTasks()
	return &CreateTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTaskHistories 获取网站的历史扫描任务
//
// 获取网站漏洞扫描的历史扫描任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsInspectorClient) ListTaskHistories(request *model.ListTaskHistoriesRequest) (*model.ListTaskHistoriesResponse, error) {
	requestDef := GenReqDefForListTaskHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTaskHistoriesResponse), nil
	}
}

// ListTaskHistoriesInvoker 获取网站的历史扫描任务
func (c *CodeArtsInspectorClient) ListTaskHistoriesInvoker(request *model.ListTaskHistoriesRequest) *ListTaskHistoriesInvoker {
	requestDef := GenReqDefForListTaskHistories()
	return &ListTaskHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTasks 获取网站扫描任务详情
//
// 获取网站漏洞扫描任务详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsInspectorClient) ShowTasks(request *model.ShowTasksRequest) (*model.ShowTasksResponse, error) {
	requestDef := GenReqDefForShowTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTasksResponse), nil
	}
}

// ShowTasksInvoker 获取网站扫描任务详情
func (c *CodeArtsInspectorClient) ShowTasksInvoker(request *model.ShowTasksRequest) *ShowTasksInvoker {
	requestDef := GenReqDefForShowTasks()
	return &ShowTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AuthorizeDomains 认证网站资产
//
// 认证租户的网站资产
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsInspectorClient) AuthorizeDomains(request *model.AuthorizeDomainsRequest) (*model.AuthorizeDomainsResponse, error) {
	requestDef := GenReqDefForAuthorizeDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AuthorizeDomainsResponse), nil
	}
}

// AuthorizeDomainsInvoker 认证网站资产
func (c *CodeArtsInspectorClient) AuthorizeDomainsInvoker(request *model.AuthorizeDomainsRequest) *AuthorizeDomainsInvoker {
	requestDef := GenReqDefForAuthorizeDomains()
	return &AuthorizeDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDomains 创建网站资产
//
// 创建租户的网站资产
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsInspectorClient) CreateDomains(request *model.CreateDomainsRequest) (*model.CreateDomainsResponse, error) {
	requestDef := GenReqDefForCreateDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDomainsResponse), nil
	}
}

// CreateDomainsInvoker 创建网站资产
func (c *CodeArtsInspectorClient) CreateDomainsInvoker(request *model.CreateDomainsRequest) *CreateDomainsInvoker {
	requestDef := GenReqDefForCreateDomains()
	return &CreateDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDomains 删除网站资产
//
// 删除租户的网站资产
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsInspectorClient) DeleteDomains(request *model.DeleteDomainsRequest) (*model.DeleteDomainsResponse, error) {
	requestDef := GenReqDefForDeleteDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDomainsResponse), nil
	}
}

// DeleteDomainsInvoker 删除网站资产
func (c *CodeArtsInspectorClient) DeleteDomainsInvoker(request *model.DeleteDomainsRequest) *DeleteDomainsInvoker {
	requestDef := GenReqDefForDeleteDomains()
	return &DeleteDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDomains 获取网站资产
//
// 获取租户的所有网站资产
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsInspectorClient) ListDomains(request *model.ListDomainsRequest) (*model.ListDomainsResponse, error) {
	requestDef := GenReqDefForListDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainsResponse), nil
	}
}

// ListDomainsInvoker 获取网站资产
func (c *CodeArtsInspectorClient) ListDomainsInvoker(request *model.ListDomainsRequest) *ListDomainsInvoker {
	requestDef := GenReqDefForListDomains()
	return &ListDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDomainSettings 获取网站配置
//
// 获取网站登录配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsInspectorClient) ShowDomainSettings(request *model.ShowDomainSettingsRequest) (*model.ShowDomainSettingsResponse, error) {
	requestDef := GenReqDefForShowDomainSettings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainSettingsResponse), nil
	}
}

// ShowDomainSettingsInvoker 获取网站配置
func (c *CodeArtsInspectorClient) ShowDomainSettingsInvoker(request *model.ShowDomainSettingsRequest) *ShowDomainSettingsInvoker {
	requestDef := GenReqDefForShowDomainSettings()
	return &ShowDomainSettingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDomainSettings 更新网站配置
//
// 更新网站登录配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsInspectorClient) UpdateDomainSettings(request *model.UpdateDomainSettingsRequest) (*model.UpdateDomainSettingsResponse, error) {
	requestDef := GenReqDefForUpdateDomainSettings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDomainSettingsResponse), nil
	}
}

// UpdateDomainSettingsInvoker 更新网站配置
func (c *CodeArtsInspectorClient) UpdateDomainSettingsInvoker(request *model.UpdateDomainSettingsRequest) *UpdateDomainSettingsInvoker {
	requestDef := GenReqDefForUpdateDomainSettings()
	return &UpdateDomainSettingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
