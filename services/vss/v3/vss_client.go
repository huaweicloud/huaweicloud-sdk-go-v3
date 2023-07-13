package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vss/v3/model"
)

type VssClient struct {
	HcClient *http_client.HcHttpClient
}

func NewVssClient(hcClient *http_client.HcHttpClient) *VssClient {
	return &VssClient{HcClient: hcClient}
}

func VssClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// DownloadTaskReport 下载网站扫描报告
//
// 下载网站扫描任务PDF报告
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VssClient) DownloadTaskReport(request *model.DownloadTaskReportRequest) (*model.DownloadTaskReportResponse, error) {
	requestDef := GenReqDefForDownloadTaskReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadTaskReportResponse), nil
	}
}

// DownloadTaskReportInvoker 下载网站扫描报告
func (c *VssClient) DownloadTaskReportInvoker(request *model.DownloadTaskReportRequest) *DownloadTaskReportInvoker {
	requestDef := GenReqDefForDownloadTaskReport()
	return &DownloadTaskReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteGenerateReport 生成网站扫描报告
//
// 生成网站扫描PDF报告
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VssClient) ExecuteGenerateReport(request *model.ExecuteGenerateReportRequest) (*model.ExecuteGenerateReportResponse, error) {
	requestDef := GenReqDefForExecuteGenerateReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteGenerateReportResponse), nil
	}
}

// ExecuteGenerateReportInvoker 生成网站扫描报告
func (c *VssClient) ExecuteGenerateReportInvoker(request *model.ExecuteGenerateReportRequest) *ExecuteGenerateReportInvoker {
	requestDef := GenReqDefForExecuteGenerateReport()
	return &ExecuteGenerateReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBusinessRisks 获取网站业务风险扫描结果
//
// 获取网站业务风险扫描结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VssClient) ListBusinessRisks(request *model.ListBusinessRisksRequest) (*model.ListBusinessRisksResponse, error) {
	requestDef := GenReqDefForListBusinessRisks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBusinessRisksResponse), nil
	}
}

// ListBusinessRisksInvoker 获取网站业务风险扫描结果
func (c *VssClient) ListBusinessRisksInvoker(request *model.ListBusinessRisksRequest) *ListBusinessRisksInvoker {
	requestDef := GenReqDefForListBusinessRisks()
	return &ListBusinessRisksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPortResults 获取网站端口扫描结果
//
// 获取网站端口扫描结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VssClient) ListPortResults(request *model.ListPortResultsRequest) (*model.ListPortResultsResponse, error) {
	requestDef := GenReqDefForListPortResults()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPortResultsResponse), nil
	}
}

// ListPortResultsInvoker 获取网站端口扫描结果
func (c *VssClient) ListPortResultsInvoker(request *model.ListPortResultsRequest) *ListPortResultsInvoker {
	requestDef := GenReqDefForListPortResults()
	return &ListPortResultsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowReportStatus 获取网站扫描报告状态
//
// 获取网站扫描PDF报告生成状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VssClient) ShowReportStatus(request *model.ShowReportStatusRequest) (*model.ShowReportStatusResponse, error) {
	requestDef := GenReqDefForShowReportStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReportStatusResponse), nil
	}
}

// ShowReportStatusInvoker 获取网站扫描报告状态
func (c *VssClient) ShowReportStatusInvoker(request *model.ShowReportStatusRequest) *ShowReportStatusInvoker {
	requestDef := GenReqDefForShowReportStatus()
	return &ShowReportStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResults 获取网站扫描结果
//
// 获取网站漏洞扫描结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VssClient) ShowResults(request *model.ShowResultsRequest) (*model.ShowResultsResponse, error) {
	requestDef := GenReqDefForShowResults()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResultsResponse), nil
	}
}

// ShowResultsInvoker 获取网站扫描结果
func (c *VssClient) ShowResultsInvoker(request *model.ShowResultsRequest) *ShowResultsInvoker {
	requestDef := GenReqDefForShowResults()
	return &ShowResultsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFalsePositive 更新网站漏洞的误报状态
//
// 更新网站扫描漏洞的误报状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VssClient) UpdateFalsePositive(request *model.UpdateFalsePositiveRequest) (*model.UpdateFalsePositiveResponse, error) {
	requestDef := GenReqDefForUpdateFalsePositive()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFalsePositiveResponse), nil
	}
}

// UpdateFalsePositiveInvoker 更新网站漏洞的误报状态
func (c *VssClient) UpdateFalsePositiveInvoker(request *model.UpdateFalsePositiveRequest) *UpdateFalsePositiveInvoker {
	requestDef := GenReqDefForUpdateFalsePositive()
	return &UpdateFalsePositiveInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelTasks 取消或重启网站扫描任务
//
// 取消或重启网站漏洞扫描任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VssClient) CancelTasks(request *model.CancelTasksRequest) (*model.CancelTasksResponse, error) {
	requestDef := GenReqDefForCancelTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelTasksResponse), nil
	}
}

// CancelTasksInvoker 取消或重启网站扫描任务
func (c *VssClient) CancelTasksInvoker(request *model.CancelTasksRequest) *CancelTasksInvoker {
	requestDef := GenReqDefForCancelTasks()
	return &CancelTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTasks 创建网站扫描任务并启动
//
// 创建网站漏洞扫描任务并启动
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VssClient) CreateTasks(request *model.CreateTasksRequest) (*model.CreateTasksResponse, error) {
	requestDef := GenReqDefForCreateTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTasksResponse), nil
	}
}

// CreateTasksInvoker 创建网站扫描任务并启动
func (c *VssClient) CreateTasksInvoker(request *model.CreateTasksRequest) *CreateTasksInvoker {
	requestDef := GenReqDefForCreateTasks()
	return &CreateTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTaskHistories 获取网站的历史扫描任务
//
// 获取网站漏洞扫描的历史扫描任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VssClient) ListTaskHistories(request *model.ListTaskHistoriesRequest) (*model.ListTaskHistoriesResponse, error) {
	requestDef := GenReqDefForListTaskHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTaskHistoriesResponse), nil
	}
}

// ListTaskHistoriesInvoker 获取网站的历史扫描任务
func (c *VssClient) ListTaskHistoriesInvoker(request *model.ListTaskHistoriesRequest) *ListTaskHistoriesInvoker {
	requestDef := GenReqDefForListTaskHistories()
	return &ListTaskHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTasks 获取网站扫描任务详情
//
// 获取网站漏洞扫描任务详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VssClient) ShowTasks(request *model.ShowTasksRequest) (*model.ShowTasksResponse, error) {
	requestDef := GenReqDefForShowTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTasksResponse), nil
	}
}

// ShowTasksInvoker 获取网站扫描任务详情
func (c *VssClient) ShowTasksInvoker(request *model.ShowTasksRequest) *ShowTasksInvoker {
	requestDef := GenReqDefForShowTasks()
	return &ShowTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AuthorizeDomains 认证网站资产
//
// 认证租户的网站资产
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VssClient) AuthorizeDomains(request *model.AuthorizeDomainsRequest) (*model.AuthorizeDomainsResponse, error) {
	requestDef := GenReqDefForAuthorizeDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AuthorizeDomainsResponse), nil
	}
}

// AuthorizeDomainsInvoker 认证网站资产
func (c *VssClient) AuthorizeDomainsInvoker(request *model.AuthorizeDomainsRequest) *AuthorizeDomainsInvoker {
	requestDef := GenReqDefForAuthorizeDomains()
	return &AuthorizeDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDomains 创建网站资产
//
// 创建租户的网站资产
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VssClient) CreateDomains(request *model.CreateDomainsRequest) (*model.CreateDomainsResponse, error) {
	requestDef := GenReqDefForCreateDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDomainsResponse), nil
	}
}

// CreateDomainsInvoker 创建网站资产
func (c *VssClient) CreateDomainsInvoker(request *model.CreateDomainsRequest) *CreateDomainsInvoker {
	requestDef := GenReqDefForCreateDomains()
	return &CreateDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDomains 删除网站资产
//
// 删除租户的网站资产
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VssClient) DeleteDomains(request *model.DeleteDomainsRequest) (*model.DeleteDomainsResponse, error) {
	requestDef := GenReqDefForDeleteDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDomainsResponse), nil
	}
}

// DeleteDomainsInvoker 删除网站资产
func (c *VssClient) DeleteDomainsInvoker(request *model.DeleteDomainsRequest) *DeleteDomainsInvoker {
	requestDef := GenReqDefForDeleteDomains()
	return &DeleteDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDomains 获取网站资产
//
// 获取租户的所有网站资产
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VssClient) ListDomains(request *model.ListDomainsRequest) (*model.ListDomainsResponse, error) {
	requestDef := GenReqDefForListDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainsResponse), nil
	}
}

// ListDomainsInvoker 获取网站资产
func (c *VssClient) ListDomainsInvoker(request *model.ListDomainsRequest) *ListDomainsInvoker {
	requestDef := GenReqDefForListDomains()
	return &ListDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDomainSettings 获取网站配置
//
// 获取网站登录配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VssClient) ShowDomainSettings(request *model.ShowDomainSettingsRequest) (*model.ShowDomainSettingsResponse, error) {
	requestDef := GenReqDefForShowDomainSettings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainSettingsResponse), nil
	}
}

// ShowDomainSettingsInvoker 获取网站配置
func (c *VssClient) ShowDomainSettingsInvoker(request *model.ShowDomainSettingsRequest) *ShowDomainSettingsInvoker {
	requestDef := GenReqDefForShowDomainSettings()
	return &ShowDomainSettingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDomainSettings 更新网站配置
//
// 更新网站登录配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VssClient) UpdateDomainSettings(request *model.UpdateDomainSettingsRequest) (*model.UpdateDomainSettingsResponse, error) {
	requestDef := GenReqDefForUpdateDomainSettings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDomainSettingsResponse), nil
	}
}

// UpdateDomainSettingsInvoker 更新网站配置
func (c *VssClient) UpdateDomainSettingsInvoker(request *model.UpdateDomainSettingsRequest) *UpdateDomainSettingsInvoker {
	requestDef := GenReqDefForUpdateDomainSettings()
	return &UpdateDomainSettingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
