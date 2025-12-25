package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/metastudio/v1/model"
)

type MetaStudioClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewMetaStudioClient(hcClient *httpclient.HcHttpClient) *MetaStudioClient {
	return &MetaStudioClient{HcClient: hcClient}
}

func MetaStudioClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateActiveCode 创建激活码
//
// 该接口用于创建激活码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateActiveCode(request *model.CreateActiveCodeRequest) (*model.CreateActiveCodeResponse, error) {
	requestDef := GenReqDefForCreateActiveCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateActiveCodeResponse), nil
	}
}

// CreateActiveCodeInvoker 创建激活码
func (c *MetaStudioClient) CreateActiveCodeInvoker(request *model.CreateActiveCodeRequest) *CreateActiveCodeInvoker {
	requestDef := GenReqDefForCreateActiveCode()
	return &CreateActiveCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteActiveCode 删除激活码
//
// 该接口用于删除激活码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteActiveCode(request *model.DeleteActiveCodeRequest) (*model.DeleteActiveCodeResponse, error) {
	requestDef := GenReqDefForDeleteActiveCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteActiveCodeResponse), nil
	}
}

// DeleteActiveCodeInvoker 删除激活码
func (c *MetaStudioClient) DeleteActiveCodeInvoker(request *model.DeleteActiveCodeRequest) *DeleteActiveCodeInvoker {
	requestDef := GenReqDefForDeleteActiveCode()
	return &DeleteActiveCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListActiveCode 查询激活码列表
//
// 该接口用于查询激活码列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListActiveCode(request *model.ListActiveCodeRequest) (*model.ListActiveCodeResponse, error) {
	requestDef := GenReqDefForListActiveCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListActiveCodeResponse), nil
	}
}

// ListActiveCodeInvoker 查询激活码列表
func (c *MetaStudioClient) ListActiveCodeInvoker(request *model.ListActiveCodeRequest) *ListActiveCodeInvoker {
	requestDef := GenReqDefForListActiveCode()
	return &ListActiveCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetActiveCode 重置激活码
//
// 该接口用于重置激活码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ResetActiveCode(request *model.ResetActiveCodeRequest) (*model.ResetActiveCodeResponse, error) {
	requestDef := GenReqDefForResetActiveCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetActiveCodeResponse), nil
	}
}

// ResetActiveCodeInvoker 重置激活码
func (c *MetaStudioClient) ResetActiveCodeInvoker(request *model.ResetActiveCodeRequest) *ResetActiveCodeInvoker {
	requestDef := GenReqDefForResetActiveCode()
	return &ResetActiveCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowActiveCode 查询激活码详情
//
// 该接口用于查询激活码详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowActiveCode(request *model.ShowActiveCodeRequest) (*model.ShowActiveCodeResponse, error) {
	requestDef := GenReqDefForShowActiveCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowActiveCodeResponse), nil
	}
}

// ShowActiveCodeInvoker 查询激活码详情
func (c *MetaStudioClient) ShowActiveCodeInvoker(request *model.ShowActiveCodeRequest) *ShowActiveCodeInvoker {
	requestDef := GenReqDefForShowActiveCode()
	return &ShowActiveCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateActiveCode 修改激活码
//
// 该接口用于修改激活码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateActiveCode(request *model.UpdateActiveCodeRequest) (*model.UpdateActiveCodeResponse, error) {
	requestDef := GenReqDefForUpdateActiveCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateActiveCodeResponse), nil
	}
}

// UpdateActiveCodeInvoker 修改激活码
func (c *MetaStudioClient) UpdateActiveCodeInvoker(request *model.UpdateActiveCodeRequest) *UpdateActiveCodeInvoker {
	requestDef := GenReqDefForUpdateActiveCode()
	return &UpdateActiveCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAgencyWithRoleType 创建委托
//
// 该接口用于创建委托。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateAgencyWithRoleType(request *model.CreateAgencyWithRoleTypeRequest) (*model.CreateAgencyWithRoleTypeResponse, error) {
	requestDef := GenReqDefForCreateAgencyWithRoleType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAgencyWithRoleTypeResponse), nil
	}
}

// CreateAgencyWithRoleTypeInvoker 创建委托
func (c *MetaStudioClient) CreateAgencyWithRoleTypeInvoker(request *model.CreateAgencyWithRoleTypeRequest) *CreateAgencyWithRoleTypeInvoker {
	requestDef := GenReqDefForCreateAgencyWithRoleType()
	return &CreateAgencyWithRoleTypeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAgencyWithRoleType 删除委托
//
// 该接口用于删除项目下委托。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteAgencyWithRoleType(request *model.DeleteAgencyWithRoleTypeRequest) (*model.DeleteAgencyWithRoleTypeResponse, error) {
	requestDef := GenReqDefForDeleteAgencyWithRoleType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAgencyWithRoleTypeResponse), nil
	}
}

// DeleteAgencyWithRoleTypeInvoker 删除委托
func (c *MetaStudioClient) DeleteAgencyWithRoleTypeInvoker(request *model.DeleteAgencyWithRoleTypeRequest) *DeleteAgencyWithRoleTypeInvoker {
	requestDef := GenReqDefForDeleteAgencyWithRoleType()
	return &DeleteAgencyWithRoleTypeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAgency 查询委托
//
// 该接口用于查询项目下委托
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowAgency(request *model.ShowAgencyRequest) (*model.ShowAgencyResponse, error) {
	requestDef := GenReqDefForShowAgency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAgencyResponse), nil
	}
}

// ShowAgencyInvoker 查询委托
func (c *MetaStudioClient) ShowAgencyInvoker(request *model.ShowAgencyRequest) *ShowAgencyInvoker {
	requestDef := GenReqDefForShowAgency()
	return &ShowAgencyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAudioRecordConfig 创建语音录制配置
//
// 该接口用于创建语音录制配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateAudioRecordConfig(request *model.CreateAudioRecordConfigRequest) (*model.CreateAudioRecordConfigResponse, error) {
	requestDef := GenReqDefForCreateAudioRecordConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAudioRecordConfigResponse), nil
	}
}

// CreateAudioRecordConfigInvoker 创建语音录制配置
func (c *MetaStudioClient) CreateAudioRecordConfigInvoker(request *model.CreateAudioRecordConfigRequest) *CreateAudioRecordConfigInvoker {
	requestDef := GenReqDefForCreateAudioRecordConfig()
	return &CreateAudioRecordConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAudioRecordConfig 删除语音录制配置
//
// 该接口用于删除语音录制配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteAudioRecordConfig(request *model.DeleteAudioRecordConfigRequest) (*model.DeleteAudioRecordConfigResponse, error) {
	requestDef := GenReqDefForDeleteAudioRecordConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAudioRecordConfigResponse), nil
	}
}

// DeleteAudioRecordConfigInvoker 删除语音录制配置
func (c *MetaStudioClient) DeleteAudioRecordConfigInvoker(request *model.DeleteAudioRecordConfigRequest) *DeleteAudioRecordConfigInvoker {
	requestDef := GenReqDefForDeleteAudioRecordConfig()
	return &DeleteAudioRecordConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAudioRecordConfig 查询语音录制配置
//
// 该接口用于查询语音录制配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowAudioRecordConfig(request *model.ShowAudioRecordConfigRequest) (*model.ShowAudioRecordConfigResponse, error) {
	requestDef := GenReqDefForShowAudioRecordConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAudioRecordConfigResponse), nil
	}
}

// ShowAudioRecordConfigInvoker 查询语音录制配置
func (c *MetaStudioClient) ShowAudioRecordConfigInvoker(request *model.ShowAudioRecordConfigRequest) *ShowAudioRecordConfigInvoker {
	requestDef := GenReqDefForShowAudioRecordConfig()
	return &ShowAudioRecordConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAudioRecordConfig 修改语音录制配置
//
// 该接口用于修改语音录制配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateAudioRecordConfig(request *model.UpdateAudioRecordConfigRequest) (*model.UpdateAudioRecordConfigResponse, error) {
	requestDef := GenReqDefForUpdateAudioRecordConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAudioRecordConfigResponse), nil
	}
}

// UpdateAudioRecordConfigInvoker 修改语音录制配置
func (c *MetaStudioClient) UpdateAudioRecordConfigInvoker(request *model.UpdateAudioRecordConfigRequest) *UpdateAudioRecordConfigInvoker {
	requestDef := GenReqDefForUpdateAudioRecordConfig()
	return &UpdateAudioRecordConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTask 删除导入导出任务
//
// 删除导入导出任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteTask(request *model.DeleteTaskRequest) (*model.DeleteTaskResponse, error) {
	requestDef := GenReqDefForDeleteTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTaskResponse), nil
	}
}

// DeleteTaskInvoker 删除导入导出任务
func (c *MetaStudioClient) DeleteTaskInvoker(request *model.DeleteTaskRequest) *DeleteTaskInvoker {
	requestDef := GenReqDefForDeleteTask()
	return &DeleteTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadResultFile 下载导入或导出的结果文件
//
// 下载导入或者导出结果文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DownloadResultFile(request *model.DownloadResultFileRequest) (*model.DownloadResultFileResponse, error) {
	requestDef := GenReqDefForDownloadResultFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadResultFileResponse), nil
	}
}

// DownloadResultFileInvoker 下载导入或导出的结果文件
func (c *MetaStudioClient) DownloadResultFileInvoker(request *model.DownloadResultFileRequest) *DownloadResultFileInvoker {
	requestDef := GenReqDefForDownloadResultFile()
	return &DownloadResultFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadTemplate 下载信息导入模板
//
// 下载导入模板，返回导入模板文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DownloadTemplate(request *model.DownloadTemplateRequest) (*model.DownloadTemplateResponse, error) {
	requestDef := GenReqDefForDownloadTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadTemplateResponse), nil
	}
}

// DownloadTemplateInvoker 下载信息导入模板
func (c *MetaStudioClient) DownloadTemplateInvoker(request *model.DownloadTemplateRequest) *DownloadTemplateInvoker {
	requestDef := GenReqDefForDownloadTemplate()
	return &DownloadTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportResource 导出文件
//
// 导出文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ExportResource(request *model.ExportResourceRequest) (*model.ExportResourceResponse, error) {
	requestDef := GenReqDefForExportResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportResourceResponse), nil
	}
}

// ExportResourceInvoker 导出文件
func (c *MetaStudioClient) ExportResourceInvoker(request *model.ExportResourceRequest) *ExportResourceInvoker {
	requestDef := GenReqDefForExportResource()
	return &ExportResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportResource 导入文件
//
// 导入文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ImportResource(request *model.ImportResourceRequest) (*model.ImportResourceResponse, error) {
	requestDef := GenReqDefForImportResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportResourceResponse), nil
	}
}

// ImportResourceInvoker 导入文件
func (c *MetaStudioClient) ImportResourceInvoker(request *model.ImportResourceRequest) *ImportResourceInvoker {
	requestDef := GenReqDefForImportResource()
	return &ImportResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchTask 分页查询导入导出任务列表
//
// 分页查询导入导出任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) SearchTask(request *model.SearchTaskRequest) (*model.SearchTaskResponse, error) {
	requestDef := GenReqDefForSearchTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchTaskResponse), nil
	}
}

// SearchTaskInvoker 分页查询导入导出任务列表
func (c *MetaStudioClient) SearchTaskInvoker(request *model.SearchTaskRequest) *SearchTaskInvoker {
	requestDef := GenReqDefForSearchTask()
	return &SearchTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTask 查询导入导出任务详情
//
// 查询导入导出任务详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowTask(request *model.ShowTaskRequest) (*model.ShowTaskResponse, error) {
	requestDef := GenReqDefForShowTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskResponse), nil
	}
}

// ShowTaskInvoker 查询导入导出任务详情
func (c *MetaStudioClient) ShowTaskInvoker(request *model.ShowTaskRequest) *ShowTaskInvoker {
	requestDef := GenReqDefForShowTask()
	return &ShowTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSummaryUsageData 查询用户数据概览
//
// 获取周期内用户的资源使用情况概览
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowSummaryUsageData(request *model.ShowSummaryUsageDataRequest) (*model.ShowSummaryUsageDataResponse, error) {
	requestDef := GenReqDefForShowSummaryUsageData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSummaryUsageDataResponse), nil
	}
}

// ShowSummaryUsageDataInvoker 查询用户数据概览
func (c *MetaStudioClient) ShowSummaryUsageDataInvoker(request *model.ShowSummaryUsageDataRequest) *ShowSummaryUsageDataInvoker {
	requestDef := GenReqDefForShowSummaryUsageData()
	return &ShowSummaryUsageDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUsageData 查询用户数据详情
//
// 获取周期内用户的资源使用情况
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowUsageData(request *model.ShowUsageDataRequest) (*model.ShowUsageDataResponse, error) {
	requestDef := GenReqDefForShowUsageData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUsageDataResponse), nil
	}
}

// ShowUsageDataInvoker 查询用户数据详情
func (c *MetaStudioClient) ShowUsageDataInvoker(request *model.ShowUsageDataRequest) *ShowUsageDataInvoker {
	requestDef := GenReqDefForShowUsageData()
	return &ShowUsageDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDialogUrl 创建对话链接
//
// 该接口用于创建对话链接。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateDialogUrl(request *model.CreateDialogUrlRequest) (*model.CreateDialogUrlResponse, error) {
	requestDef := GenReqDefForCreateDialogUrl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDialogUrlResponse), nil
	}
}

// CreateDialogUrlInvoker 创建对话链接
func (c *MetaStudioClient) CreateDialogUrlInvoker(request *model.CreateDialogUrlRequest) *CreateDialogUrlInvoker {
	requestDef := GenReqDefForCreateDialogUrl()
	return &CreateDialogUrlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSmartChatJob 查询数字人智能交互任务列表
//
// 该接口用于查询数字人智能交互任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListSmartChatJob(request *model.ListSmartChatJobRequest) (*model.ListSmartChatJobResponse, error) {
	requestDef := GenReqDefForListSmartChatJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSmartChatJobResponse), nil
	}
}

// ListSmartChatJobInvoker 查询数字人智能交互任务列表
func (c *MetaStudioClient) ListSmartChatJobInvoker(request *model.ListSmartChatJobRequest) *ListSmartChatJobInvoker {
	requestDef := GenReqDefForListSmartChatJob()
	return &ListSmartChatJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSmartChatJob 查询数字人智能交互任务
//
// 该接口用于查询数字人智能交互任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowSmartChatJob(request *model.ShowSmartChatJobRequest) (*model.ShowSmartChatJobResponse, error) {
	requestDef := GenReqDefForShowSmartChatJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSmartChatJobResponse), nil
	}
}

// ShowSmartChatJobInvoker 查询数字人智能交互任务
func (c *MetaStudioClient) ShowSmartChatJobInvoker(request *model.ShowSmartChatJobRequest) *ShowSmartChatJobInvoker {
	requestDef := GenReqDefForShowSmartChatJob()
	return &ShowSmartChatJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartSmartChatJob 启动数字人智能交互任务
//
// 该接口用于启动数字人智能交互任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) StartSmartChatJob(request *model.StartSmartChatJobRequest) (*model.StartSmartChatJobResponse, error) {
	requestDef := GenReqDefForStartSmartChatJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartSmartChatJobResponse), nil
	}
}

// StartSmartChatJobInvoker 启动数字人智能交互任务
func (c *MetaStudioClient) StartSmartChatJobInvoker(request *model.StartSmartChatJobRequest) *StartSmartChatJobInvoker {
	requestDef := GenReqDefForStartSmartChatJob()
	return &StartSmartChatJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopSmartChatJob 结束数字人智能交互任务
//
// 该接口用于结束数字人智能交互任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) StopSmartChatJob(request *model.StopSmartChatJobRequest) (*model.StopSmartChatJobResponse, error) {
	requestDef := GenReqDefForStopSmartChatJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopSmartChatJobResponse), nil
	}
}

// StopSmartChatJobInvoker 结束数字人智能交互任务
func (c *MetaStudioClient) StopSmartChatJobInvoker(request *model.StopSmartChatJobRequest) *StopSmartChatJobInvoker {
	requestDef := GenReqDefForStopSmartChatJob()
	return &StopSmartChatJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDialogReportConfig 创建对话结果上报配置
//
// 该接口用于创建对话结果上报配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateDialogReportConfig(request *model.CreateDialogReportConfigRequest) (*model.CreateDialogReportConfigResponse, error) {
	requestDef := GenReqDefForCreateDialogReportConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDialogReportConfigResponse), nil
	}
}

// CreateDialogReportConfigInvoker 创建对话结果上报配置
func (c *MetaStudioClient) CreateDialogReportConfigInvoker(request *model.CreateDialogReportConfigRequest) *CreateDialogReportConfigInvoker {
	requestDef := GenReqDefForCreateDialogReportConfig()
	return &CreateDialogReportConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDialogReportConfig 删除对话结果上报配置
//
// 该接口用于删除对话结果上报配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteDialogReportConfig(request *model.DeleteDialogReportConfigRequest) (*model.DeleteDialogReportConfigResponse, error) {
	requestDef := GenReqDefForDeleteDialogReportConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDialogReportConfigResponse), nil
	}
}

// DeleteDialogReportConfigInvoker 删除对话结果上报配置
func (c *MetaStudioClient) DeleteDialogReportConfigInvoker(request *model.DeleteDialogReportConfigRequest) *DeleteDialogReportConfigInvoker {
	requestDef := GenReqDefForDeleteDialogReportConfig()
	return &DeleteDialogReportConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDialogReportConfig 查询对话结果上报配置
//
// 该接口用于查询对话结果上报配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowDialogReportConfig(request *model.ShowDialogReportConfigRequest) (*model.ShowDialogReportConfigResponse, error) {
	requestDef := GenReqDefForShowDialogReportConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDialogReportConfigResponse), nil
	}
}

// ShowDialogReportConfigInvoker 查询对话结果上报配置
func (c *MetaStudioClient) ShowDialogReportConfigInvoker(request *model.ShowDialogReportConfigRequest) *ShowDialogReportConfigInvoker {
	requestDef := GenReqDefForShowDialogReportConfig()
	return &ShowDialogReportConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDialogReportConfig 修改对话结果上报配置
//
// 该接口用于修改对话结果上报配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateDialogReportConfig(request *model.UpdateDialogReportConfigRequest) (*model.UpdateDialogReportConfigResponse, error) {
	requestDef := GenReqDefForUpdateDialogReportConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDialogReportConfigResponse), nil
	}
}

// UpdateDialogReportConfigInvoker 修改对话结果上报配置
func (c *MetaStudioClient) UpdateDialogReportConfigInvoker(request *model.UpdateDialogReportConfigRequest) *UpdateDialogReportConfigInvoker {
	requestDef := GenReqDefForUpdateDialogReportConfig()
	return &UpdateDialogReportConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchExecuteAssetAction 批量资产操作
//
// 该接口用批量资产操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) BatchExecuteAssetAction(request *model.BatchExecuteAssetActionRequest) (*model.BatchExecuteAssetActionResponse, error) {
	requestDef := GenReqDefForBatchExecuteAssetAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchExecuteAssetActionResponse), nil
	}
}

// BatchExecuteAssetActionInvoker 批量资产操作
func (c *MetaStudioClient) BatchExecuteAssetActionInvoker(request *model.BatchExecuteAssetActionRequest) *BatchExecuteAssetActionInvoker {
	requestDef := GenReqDefForBatchExecuteAssetAction()
	return &BatchExecuteAssetActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAssetByReplicationInfo 复制资产
//
// 该接口用于在Region B复制Region A的指定资产。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateAssetByReplicationInfo(request *model.CreateAssetByReplicationInfoRequest) (*model.CreateAssetByReplicationInfoResponse, error) {
	requestDef := GenReqDefForCreateAssetByReplicationInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAssetByReplicationInfoResponse), nil
	}
}

// CreateAssetByReplicationInfoInvoker 复制资产
func (c *MetaStudioClient) CreateAssetByReplicationInfoInvoker(request *model.CreateAssetByReplicationInfoRequest) *CreateAssetByReplicationInfoInvoker {
	requestDef := GenReqDefForCreateAssetByReplicationInfo()
	return &CreateAssetByReplicationInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDigitalAsset 创建资产
//
// 该接口用于在资产库中添加上传新的媒体资产。可上传的资产类型包括：分身数字人模型、背景图片、素材图片、素材视频、PPT等。
// &gt; 上传的图片、视频和背景图片，如果需要在视频制作素材中可见，需要设置system_properties。
// &gt; - 资产类型是IMAGE时，通过system_properties来区分背景图片（BACKGROUND_IMG）、素材图片（MATERIAL_IMG）。
// &gt; - 资产类型是VIDEO时，通过system_properties来区分素材视频（MATERIAL_VIDEO）、名片视频（BUSSINESS_CARD_VIDEO）。
// &gt; MetaStudio平台生成的视频，system_properties带CREATED_BY_PLATFORM。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateDigitalAsset(request *model.CreateDigitalAssetRequest) (*model.CreateDigitalAssetResponse, error) {
	requestDef := GenReqDefForCreateDigitalAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDigitalAssetResponse), nil
	}
}

// CreateDigitalAssetInvoker 创建资产
func (c *MetaStudioClient) CreateDigitalAssetInvoker(request *model.CreateDigitalAssetRequest) *CreateDigitalAssetInvoker {
	requestDef := GenReqDefForCreateDigitalAsset()
	return &CreateDigitalAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAsset 删除资产
//
// 该接口用于删除资产库中的媒体资产。调用该接口删除媒体资产时，媒体资产会放入回收站中，不会彻底删除。如需彻底删除资产，需增加“mode&#x3D;force”参数配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteAsset(request *model.DeleteAssetRequest) (*model.DeleteAssetResponse, error) {
	requestDef := GenReqDefForDeleteAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAssetResponse), nil
	}
}

// DeleteAssetInvoker 删除资产
func (c *MetaStudioClient) DeleteAssetInvoker(request *model.DeleteAssetRequest) *DeleteAssetInvoker {
	requestDef := GenReqDefForDeleteAsset()
	return &DeleteAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteTransferAssetAction 转移资产任务控制
//
// 转移资产任务控制
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ExecuteTransferAssetAction(request *model.ExecuteTransferAssetActionRequest) (*model.ExecuteTransferAssetActionResponse, error) {
	requestDef := GenReqDefForExecuteTransferAssetAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteTransferAssetActionResponse), nil
	}
}

// ExecuteTransferAssetActionInvoker 转移资产任务控制
func (c *MetaStudioClient) ExecuteTransferAssetActionInvoker(request *model.ExecuteTransferAssetActionRequest) *ExecuteTransferAssetActionInvoker {
	requestDef := GenReqDefForExecuteTransferAssetAction()
	return &ExecuteTransferAssetActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAssetSummary 查询资产概要
//
// 该接口用于查询媒体资产库中指定的多个资产的概要信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListAssetSummary(request *model.ListAssetSummaryRequest) (*model.ListAssetSummaryResponse, error) {
	requestDef := GenReqDefForListAssetSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAssetSummaryResponse), nil
	}
}

// ListAssetSummaryInvoker 查询资产概要
func (c *MetaStudioClient) ListAssetSummaryInvoker(request *model.ListAssetSummaryRequest) *ListAssetSummaryInvoker {
	requestDef := GenReqDefForListAssetSummary()
	return &ListAssetSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAssets 查询资产列表
//
// 该接口用于查询资产库中的媒体资产列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListAssets(request *model.ListAssetsRequest) (*model.ListAssetsResponse, error) {
	requestDef := GenReqDefForListAssets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAssetsResponse), nil
	}
}

// ListAssetsInvoker 查询资产列表
func (c *MetaStudioClient) ListAssetsInvoker(request *model.ListAssetsRequest) *ListAssetsInvoker {
	requestDef := GenReqDefForListAssets()
	return &ListAssetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTransferAssetJobs 资产转移任务列表
//
// 资产转移任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListTransferAssetJobs(request *model.ListTransferAssetJobsRequest) (*model.ListTransferAssetJobsResponse, error) {
	requestDef := GenReqDefForListTransferAssetJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTransferAssetJobsResponse), nil
	}
}

// ListTransferAssetJobsInvoker 资产转移任务列表
func (c *MetaStudioClient) ListTransferAssetJobsInvoker(request *model.ListTransferAssetJobsRequest) *ListTransferAssetJobsInvoker {
	requestDef := GenReqDefForListTransferAssetJobs()
	return &ListTransferAssetJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestoreAsset 恢复被删除的资产
//
// 该接口用于恢复被删除至回收站的媒体资产。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) RestoreAsset(request *model.RestoreAssetRequest) (*model.RestoreAssetResponse, error) {
	requestDef := GenReqDefForRestoreAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreAssetResponse), nil
	}
}

// RestoreAssetInvoker 恢复被删除的资产
func (c *MetaStudioClient) RestoreAssetInvoker(request *model.RestoreAssetRequest) *RestoreAssetInvoker {
	requestDef := GenReqDefForRestoreAsset()
	return &RestoreAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAsset 查询资产详情
//
// 该接口用于查询资产库中指定媒体资产的详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowAsset(request *model.ShowAssetRequest) (*model.ShowAssetResponse, error) {
	requestDef := GenReqDefForShowAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAssetResponse), nil
	}
}

// ShowAssetInvoker 查询资产详情
func (c *MetaStudioClient) ShowAssetInvoker(request *model.ShowAssetRequest) *ShowAssetInvoker {
	requestDef := GenReqDefForShowAsset()
	return &ShowAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAssetReplicationInfo 查询资产复制信息
//
// 当需要将资产从A Region复制到B Region时，先要在A Region调用该接口用于查询资产复制信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowAssetReplicationInfo(request *model.ShowAssetReplicationInfoRequest) (*model.ShowAssetReplicationInfoResponse, error) {
	requestDef := GenReqDefForShowAssetReplicationInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAssetReplicationInfoResponse), nil
	}
}

// ShowAssetReplicationInfoInvoker 查询资产复制信息
func (c *MetaStudioClient) ShowAssetReplicationInfoInvoker(request *model.ShowAssetReplicationInfoRequest) *ShowAssetReplicationInfoInvoker {
	requestDef := GenReqDefForShowAssetReplicationInfo()
	return &ShowAssetReplicationInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTransferAssetJob 查询转移资产任务详情
//
// 查询转移资产任务详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowTransferAssetJob(request *model.ShowTransferAssetJobRequest) (*model.ShowTransferAssetJobResponse, error) {
	requestDef := GenReqDefForShowTransferAssetJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTransferAssetJobResponse), nil
	}
}

// ShowTransferAssetJobInvoker 查询转移资产任务详情
func (c *MetaStudioClient) ShowTransferAssetJobInvoker(request *model.ShowTransferAssetJobRequest) *ShowTransferAssetJobInvoker {
	requestDef := GenReqDefForShowTransferAssetJob()
	return &ShowTransferAssetJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// TransferAsset 转移资产给其他用户
//
// 转移资产给其他用户
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) TransferAsset(request *model.TransferAssetRequest) (*model.TransferAssetResponse, error) {
	requestDef := GenReqDefForTransferAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.TransferAssetResponse), nil
	}
}

// TransferAssetInvoker 转移资产给其他用户
func (c *MetaStudioClient) TransferAssetInvoker(request *model.TransferAssetRequest) *TransferAssetInvoker {
	requestDef := GenReqDefForTransferAsset()
	return &TransferAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDigitalAsset 更新资产
//
// 该接口用于更新资产库中的媒体资产信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateDigitalAsset(request *model.UpdateDigitalAssetRequest) (*model.UpdateDigitalAssetResponse, error) {
	requestDef := GenReqDefForUpdateDigitalAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDigitalAssetResponse), nil
	}
}

// UpdateDigitalAssetInvoker 更新资产
func (c *MetaStudioClient) UpdateDigitalAssetInvoker(request *model.UpdateDigitalAssetRequest) *UpdateDigitalAssetInvoker {
	requestDef := GenReqDefForUpdateDigitalAsset()
	return &UpdateDigitalAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDigitalHumanBusinessCard 创建数字人名片制作
//
// 该接口用于数字人名片制作任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateDigitalHumanBusinessCard(request *model.CreateDigitalHumanBusinessCardRequest) (*model.CreateDigitalHumanBusinessCardResponse, error) {
	requestDef := GenReqDefForCreateDigitalHumanBusinessCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDigitalHumanBusinessCardResponse), nil
	}
}

// CreateDigitalHumanBusinessCardInvoker 创建数字人名片制作
func (c *MetaStudioClient) CreateDigitalHumanBusinessCardInvoker(request *model.CreateDigitalHumanBusinessCardRequest) *CreateDigitalHumanBusinessCardInvoker {
	requestDef := GenReqDefForCreateDigitalHumanBusinessCard()
	return &CreateDigitalHumanBusinessCardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDigitalHumanBusinessCard 删除数字人名片制作任务
//
// 该接口用于删除数字人名片制作任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteDigitalHumanBusinessCard(request *model.DeleteDigitalHumanBusinessCardRequest) (*model.DeleteDigitalHumanBusinessCardResponse, error) {
	requestDef := GenReqDefForDeleteDigitalHumanBusinessCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDigitalHumanBusinessCardResponse), nil
	}
}

// DeleteDigitalHumanBusinessCardInvoker 删除数字人名片制作任务
func (c *MetaStudioClient) DeleteDigitalHumanBusinessCardInvoker(request *model.DeleteDigitalHumanBusinessCardRequest) *DeleteDigitalHumanBusinessCardInvoker {
	requestDef := GenReqDefForDeleteDigitalHumanBusinessCard()
	return &DeleteDigitalHumanBusinessCardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDigitalHumanBusinessCard 查询数字人名片制作任务列表
//
// 该接口用于查询数字人名片制作任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListDigitalHumanBusinessCard(request *model.ListDigitalHumanBusinessCardRequest) (*model.ListDigitalHumanBusinessCardResponse, error) {
	requestDef := GenReqDefForListDigitalHumanBusinessCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDigitalHumanBusinessCardResponse), nil
	}
}

// ListDigitalHumanBusinessCardInvoker 查询数字人名片制作任务列表
func (c *MetaStudioClient) ListDigitalHumanBusinessCardInvoker(request *model.ListDigitalHumanBusinessCardRequest) *ListDigitalHumanBusinessCardInvoker {
	requestDef := GenReqDefForListDigitalHumanBusinessCard()
	return &ListDigitalHumanBusinessCardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDigitalHumanBusinessCard 查询数字人名片制作任务详情
//
// 该接口用于查询数字人名片制作任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowDigitalHumanBusinessCard(request *model.ShowDigitalHumanBusinessCardRequest) (*model.ShowDigitalHumanBusinessCardResponse, error) {
	requestDef := GenReqDefForShowDigitalHumanBusinessCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDigitalHumanBusinessCardResponse), nil
	}
}

// ShowDigitalHumanBusinessCardInvoker 查询数字人名片制作任务详情
func (c *MetaStudioClient) ShowDigitalHumanBusinessCardInvoker(request *model.ShowDigitalHumanBusinessCardRequest) *ShowDigitalHumanBusinessCardInvoker {
	requestDef := GenReqDefForShowDigitalHumanBusinessCard()
	return &ShowDigitalHumanBusinessCardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDigitalHumanBusinessCard 更新数字人名片制作
//
// 该接口用于更新数字人名片制作任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateDigitalHumanBusinessCard(request *model.UpdateDigitalHumanBusinessCardRequest) (*model.UpdateDigitalHumanBusinessCardResponse, error) {
	requestDef := GenReqDefForUpdateDigitalHumanBusinessCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDigitalHumanBusinessCardResponse), nil
	}
}

// UpdateDigitalHumanBusinessCardInvoker 更新数字人名片制作
func (c *MetaStudioClient) UpdateDigitalHumanBusinessCardInvoker(request *model.UpdateDigitalHumanBusinessCardRequest) *UpdateDigitalHumanBusinessCardInvoker {
	requestDef := GenReqDefForUpdateDigitalHumanBusinessCard()
	return &UpdateDigitalHumanBusinessCardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDigitalHumanVideo 查询视频制作任务列表
//
// 该接口用于查询视频制作任务列表。可查询分身数字人视频制作列表，照片数字人视频制作列表等。
// &gt; - 默认查询最近一个月任务记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListDigitalHumanVideo(request *model.ListDigitalHumanVideoRequest) (*model.ListDigitalHumanVideoResponse, error) {
	requestDef := GenReqDefForListDigitalHumanVideo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDigitalHumanVideoResponse), nil
	}
}

// ListDigitalHumanVideoInvoker 查询视频制作任务列表
func (c *MetaStudioClient) ListDigitalHumanVideoInvoker(request *model.ListDigitalHumanVideoRequest) *ListDigitalHumanVideoInvoker {
	requestDef := GenReqDefForListDigitalHumanVideo()
	return &ListDigitalHumanVideoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Cancel2DDigitalHumanVideo 取消等待中的分身数字人视频制作任务
//
// 该接口用于取消等待中的分身数字人视频制作任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) Cancel2DDigitalHumanVideo(request *model.Cancel2DDigitalHumanVideoRequest) (*model.Cancel2DDigitalHumanVideoResponse, error) {
	requestDef := GenReqDefForCancel2DDigitalHumanVideo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.Cancel2DDigitalHumanVideoResponse), nil
	}
}

// Cancel2DDigitalHumanVideoInvoker 取消等待中的分身数字人视频制作任务
func (c *MetaStudioClient) Cancel2DDigitalHumanVideoInvoker(request *model.Cancel2DDigitalHumanVideoRequest) *Cancel2DDigitalHumanVideoInvoker {
	requestDef := GenReqDefForCancel2DDigitalHumanVideo()
	return &Cancel2DDigitalHumanVideoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Create2DDigitalHumanVideo 创建分身数字人视频制作任务
//
// 该接口用于创建分身数字人视频制作任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) Create2DDigitalHumanVideo(request *model.Create2DDigitalHumanVideoRequest) (*model.Create2DDigitalHumanVideoResponse, error) {
	requestDef := GenReqDefForCreate2DDigitalHumanVideo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.Create2DDigitalHumanVideoResponse), nil
	}
}

// Create2DDigitalHumanVideoInvoker 创建分身数字人视频制作任务
func (c *MetaStudioClient) Create2DDigitalHumanVideoInvoker(request *model.Create2DDigitalHumanVideoRequest) *Create2DDigitalHumanVideoInvoker {
	requestDef := GenReqDefForCreate2DDigitalHumanVideo()
	return &Create2DDigitalHumanVideoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Show2DDigitalHumanVideo 查询分身数字人视频制作任务详情
//
// 该接口用于查询分身数字人视频制作任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) Show2DDigitalHumanVideo(request *model.Show2DDigitalHumanVideoRequest) (*model.Show2DDigitalHumanVideoResponse, error) {
	requestDef := GenReqDefForShow2DDigitalHumanVideo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.Show2DDigitalHumanVideoResponse), nil
	}
}

// Show2DDigitalHumanVideoInvoker 查询分身数字人视频制作任务详情
func (c *MetaStudioClient) Show2DDigitalHumanVideoInvoker(request *model.Show2DDigitalHumanVideoRequest) *Show2DDigitalHumanVideoInvoker {
	requestDef := GenReqDefForShow2DDigitalHumanVideo()
	return &Show2DDigitalHumanVideoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelPhotoDigitalHumanVideo 取消等待中的照片分身数字人视频制作任务
//
// 该接口用于取消等待中的照片分身数字人视频制作任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CancelPhotoDigitalHumanVideo(request *model.CancelPhotoDigitalHumanVideoRequest) (*model.CancelPhotoDigitalHumanVideoResponse, error) {
	requestDef := GenReqDefForCancelPhotoDigitalHumanVideo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelPhotoDigitalHumanVideoResponse), nil
	}
}

// CancelPhotoDigitalHumanVideoInvoker 取消等待中的照片分身数字人视频制作任务
func (c *MetaStudioClient) CancelPhotoDigitalHumanVideoInvoker(request *model.CancelPhotoDigitalHumanVideoRequest) *CancelPhotoDigitalHumanVideoInvoker {
	requestDef := GenReqDefForCancelPhotoDigitalHumanVideo()
	return &CancelPhotoDigitalHumanVideoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePhotoDetection 创建照片检测任务
//
// 该接口用于创建照片检测任务，检测照片是否满足制作照片数字人的要求。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreatePhotoDetection(request *model.CreatePhotoDetectionRequest) (*model.CreatePhotoDetectionResponse, error) {
	requestDef := GenReqDefForCreatePhotoDetection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePhotoDetectionResponse), nil
	}
}

// CreatePhotoDetectionInvoker 创建照片检测任务
func (c *MetaStudioClient) CreatePhotoDetectionInvoker(request *model.CreatePhotoDetectionRequest) *CreatePhotoDetectionInvoker {
	requestDef := GenReqDefForCreatePhotoDetection()
	return &CreatePhotoDetectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePhotoDigitalHumanVideo 创建照片分身数字人视频制作任务
//
// 该接口用于创建照片分身数字人视频制作任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreatePhotoDigitalHumanVideo(request *model.CreatePhotoDigitalHumanVideoRequest) (*model.CreatePhotoDigitalHumanVideoResponse, error) {
	requestDef := GenReqDefForCreatePhotoDigitalHumanVideo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePhotoDigitalHumanVideoResponse), nil
	}
}

// CreatePhotoDigitalHumanVideoInvoker 创建照片分身数字人视频制作任务
func (c *MetaStudioClient) CreatePhotoDigitalHumanVideoInvoker(request *model.CreatePhotoDigitalHumanVideoRequest) *CreatePhotoDigitalHumanVideoInvoker {
	requestDef := GenReqDefForCreatePhotoDigitalHumanVideo()
	return &CreatePhotoDigitalHumanVideoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPhotoDetection 查询照片检测任务详情
//
// 该接口用于查询照片检测任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowPhotoDetection(request *model.ShowPhotoDetectionRequest) (*model.ShowPhotoDetectionResponse, error) {
	requestDef := GenReqDefForShowPhotoDetection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPhotoDetectionResponse), nil
	}
}

// ShowPhotoDetectionInvoker 查询照片检测任务详情
func (c *MetaStudioClient) ShowPhotoDetectionInvoker(request *model.ShowPhotoDetectionRequest) *ShowPhotoDetectionInvoker {
	requestDef := GenReqDefForShowPhotoDetection()
	return &ShowPhotoDetectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPhotoDigitalHumanVideo 查询照片分身数字人视频制作任务详情
//
// 该接口用于查询照片分身数字人视频制作任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowPhotoDigitalHumanVideo(request *model.ShowPhotoDigitalHumanVideoRequest) (*model.ShowPhotoDigitalHumanVideoResponse, error) {
	requestDef := GenReqDefForShowPhotoDigitalHumanVideo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPhotoDigitalHumanVideoResponse), nil
	}
}

// ShowPhotoDigitalHumanVideoInvoker 查询照片分身数字人视频制作任务详情
func (c *MetaStudioClient) ShowPhotoDigitalHumanVideoInvoker(request *model.ShowPhotoDigitalHumanVideoRequest) *ShowPhotoDigitalHumanVideoInvoker {
	requestDef := GenReqDefForShowPhotoDigitalHumanVideo()
	return &ShowPhotoDigitalHumanVideoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDocument 上传文档
//
// 该接口用于上传文档。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateDocument(request *model.CreateDocumentRequest) (*model.CreateDocumentResponse, error) {
	requestDef := GenReqDefForCreateDocument()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDocumentResponse), nil
	}
}

// CreateDocumentInvoker 上传文档
func (c *MetaStudioClient) CreateDocumentInvoker(request *model.CreateDocumentRequest) *CreateDocumentInvoker {
	requestDef := GenReqDefForCreateDocument()
	return &CreateDocumentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDocument 批量删除文档
//
// 该接口用于批量删除文档。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteDocument(request *model.DeleteDocumentRequest) (*model.DeleteDocumentResponse, error) {
	requestDef := GenReqDefForDeleteDocument()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDocumentResponse), nil
	}
}

// DeleteDocumentInvoker 批量删除文档
func (c *MetaStudioClient) DeleteDocumentInvoker(request *model.DeleteDocumentRequest) *DeleteDocumentInvoker {
	requestDef := GenReqDefForDeleteDocument()
	return &DeleteDocumentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadDocument 下载文档
//
// 该接口用于下载文档。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DownloadDocument(request *model.DownloadDocumentRequest) (*model.DownloadDocumentResponse, error) {
	requestDef := GenReqDefForDownloadDocument()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadDocumentResponse), nil
	}
}

// DownloadDocumentInvoker 下载文档
func (c *MetaStudioClient) DownloadDocumentInvoker(request *model.DownloadDocumentRequest) *DownloadDocumentInvoker {
	requestDef := GenReqDefForDownloadDocument()
	return &DownloadDocumentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDocumentInfo 查询文档列表
//
// 该接口用于分页查询文档列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListDocumentInfo(request *model.ListDocumentInfoRequest) (*model.ListDocumentInfoResponse, error) {
	requestDef := GenReqDefForListDocumentInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDocumentInfoResponse), nil
	}
}

// ListDocumentInfoInvoker 查询文档列表
func (c *MetaStudioClient) ListDocumentInfoInvoker(request *model.ListDocumentInfoRequest) *ListDocumentInfoInvoker {
	requestDef := GenReqDefForListDocumentInfo()
	return &ListDocumentInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDocumentInfo 查询文档详情
//
// 该接口用于查询文档详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowDocumentInfo(request *model.ShowDocumentInfoRequest) (*model.ShowDocumentInfoResponse, error) {
	requestDef := GenReqDefForShowDocumentInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDocumentInfoResponse), nil
	}
}

// ShowDocumentInfoInvoker 查询文档详情
func (c *MetaStudioClient) ShowDocumentInfoInvoker(request *model.ShowDocumentInfoRequest) *ShowDocumentInfoInvoker {
	requestDef := GenReqDefForShowDocumentInfo()
	return &ShowDocumentInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDocument 修改文档
//
// 该接口用于修改文档
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateDocument(request *model.UpdateDocumentRequest) (*model.UpdateDocumentResponse, error) {
	requestDef := GenReqDefForUpdateDocument()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDocumentResponse), nil
	}
}

// UpdateDocumentInvoker 修改文档
func (c *MetaStudioClient) UpdateDocumentInvoker(request *model.UpdateDocumentRequest) *UpdateDocumentInvoker {
	requestDef := GenReqDefForUpdateDocument()
	return &UpdateDocumentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDocumentSegment 分页查询文档分段信息
//
// 该接口用于分页查询文档分段信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListDocumentSegment(request *model.ListDocumentSegmentRequest) (*model.ListDocumentSegmentResponse, error) {
	requestDef := GenReqDefForListDocumentSegment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDocumentSegmentResponse), nil
	}
}

// ListDocumentSegmentInvoker 分页查询文档分段信息
func (c *MetaStudioClient) ListDocumentSegmentInvoker(request *model.ListDocumentSegmentRequest) *ListDocumentSegmentInvoker {
	requestDef := GenReqDefForListDocumentSegment()
	return &ListDocumentSegmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PreviewDocumentSegment 文档分段效果预览
//
// 该接口用于文档分段效果预览。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) PreviewDocumentSegment(request *model.PreviewDocumentSegmentRequest) (*model.PreviewDocumentSegmentResponse, error) {
	requestDef := GenReqDefForPreviewDocumentSegment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PreviewDocumentSegmentResponse), nil
	}
}

// PreviewDocumentSegmentInvoker 文档分段效果预览
func (c *MetaStudioClient) PreviewDocumentSegmentInvoker(request *model.PreviewDocumentSegmentRequest) *PreviewDocumentSegmentInvoker {
	requestDef := GenReqDefForPreviewDocumentSegment()
	return &PreviewDocumentSegmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartDocumentSegment 开始文档分段
//
// 该接口用于开始文档分段任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) StartDocumentSegment(request *model.StartDocumentSegmentRequest) (*model.StartDocumentSegmentResponse, error) {
	requestDef := GenReqDefForStartDocumentSegment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartDocumentSegmentResponse), nil
	}
}

// StartDocumentSegmentInvoker 开始文档分段
func (c *MetaStudioClient) StartDocumentSegmentInvoker(request *model.StartDocumentSegmentRequest) *StartDocumentSegmentInvoker {
	requestDef := GenReqDefForStartDocumentSegment()
	return &StartDocumentSegmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDocumentSegmentInfo 修改文档分段内容
//
// 该接口用于文档分段内容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateDocumentSegmentInfo(request *model.UpdateDocumentSegmentInfoRequest) (*model.UpdateDocumentSegmentInfoResponse, error) {
	requestDef := GenReqDefForUpdateDocumentSegmentInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDocumentSegmentInfoResponse), nil
	}
}

// UpdateDocumentSegmentInfoInvoker 修改文档分段内容
func (c *MetaStudioClient) UpdateDocumentSegmentInfoInvoker(request *model.UpdateDocumentSegmentInfoRequest) *UpdateDocumentSegmentInfoInvoker {
	requestDef := GenReqDefForUpdateDocumentSegmentInfo()
	return &UpdateDocumentSegmentInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDocumentSegmentParam 更新文档分段配置
//
// 该接口用于更新文档分段配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateDocumentSegmentParam(request *model.UpdateDocumentSegmentParamRequest) (*model.UpdateDocumentSegmentParamResponse, error) {
	requestDef := GenReqDefForUpdateDocumentSegmentParam()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDocumentSegmentParamResponse), nil
	}
}

// UpdateDocumentSegmentParamInvoker 更新文档分段配置
func (c *MetaStudioClient) UpdateDocumentSegmentParamInvoker(request *model.UpdateDocumentSegmentParamRequest) *UpdateDocumentSegmentParamInvoker {
	requestDef := GenReqDefForUpdateDocumentSegmentParam()
	return &UpdateDocumentSegmentParamInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ConfirmFileUpload 确认文件已上传
//
// 资产文件上传完毕后，通过该接口确认上传完成。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ConfirmFileUpload(request *model.ConfirmFileUploadRequest) (*model.ConfirmFileUploadResponse, error) {
	requestDef := GenReqDefForConfirmFileUpload()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ConfirmFileUploadResponse), nil
	}
}

// ConfirmFileUploadInvoker 确认文件已上传
func (c *MetaStudioClient) ConfirmFileUploadInvoker(request *model.ConfirmFileUploadRequest) *ConfirmFileUploadInvoker {
	requestDef := GenReqDefForConfirmFileUpload()
	return &ConfirmFileUploadInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFile 创建文件并获取上传URL
//
// 该接口用于创建文件并获取上传URL。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateFile(request *model.CreateFileRequest) (*model.CreateFileResponse, error) {
	requestDef := GenReqDefForCreateFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFileResponse), nil
	}
}

// CreateFileInvoker 创建文件并获取上传URL
func (c *MetaStudioClient) CreateFileInvoker(request *model.CreateFileRequest) *CreateFileInvoker {
	requestDef := GenReqDefForCreateFile()
	return &CreateFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLargeFile 创建大文件
//
// 该接口用于创建大文件（超过5G），获取分段上传URL。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateLargeFile(request *model.CreateLargeFileRequest) (*model.CreateLargeFileResponse, error) {
	requestDef := GenReqDefForCreateLargeFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLargeFileResponse), nil
	}
}

// CreateLargeFileInvoker 创建大文件
func (c *MetaStudioClient) CreateLargeFileInvoker(request *model.CreateLargeFileRequest) *CreateLargeFileInvoker {
	requestDef := GenReqDefForCreateLargeFile()
	return &CreateLargeFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFile 删除文件
//
// 该接口用于删除媒体资产库中指定的文件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteFile(request *model.DeleteFileRequest) (*model.DeleteFileResponse, error) {
	requestDef := GenReqDefForDeleteFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFileResponse), nil
	}
}

// DeleteFileInvoker 删除文件
func (c *MetaStudioClient) DeleteFileInvoker(request *model.DeleteFileRequest) *DeleteFileInvoker {
	requestDef := GenReqDefForDeleteFile()
	return &DeleteFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateHotQuestion 创建热点问题
//
// 该接口用于创建热点问题。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateHotQuestion(request *model.CreateHotQuestionRequest) (*model.CreateHotQuestionResponse, error) {
	requestDef := GenReqDefForCreateHotQuestion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHotQuestionResponse), nil
	}
}

// CreateHotQuestionInvoker 创建热点问题
func (c *MetaStudioClient) CreateHotQuestionInvoker(request *model.CreateHotQuestionRequest) *CreateHotQuestionInvoker {
	requestDef := GenReqDefForCreateHotQuestion()
	return &CreateHotQuestionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHotQuestion 删除热点问题
//
// 该接口用于删除热点问题。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteHotQuestion(request *model.DeleteHotQuestionRequest) (*model.DeleteHotQuestionResponse, error) {
	requestDef := GenReqDefForDeleteHotQuestion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHotQuestionResponse), nil
	}
}

// DeleteHotQuestionInvoker 删除热点问题
func (c *MetaStudioClient) DeleteHotQuestionInvoker(request *model.DeleteHotQuestionRequest) *DeleteHotQuestionInvoker {
	requestDef := GenReqDefForDeleteHotQuestion()
	return &DeleteHotQuestionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHotQuestion 查询热点问题列表
//
// 该接口用于查询热点问题列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListHotQuestion(request *model.ListHotQuestionRequest) (*model.ListHotQuestionResponse, error) {
	requestDef := GenReqDefForListHotQuestion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHotQuestionResponse), nil
	}
}

// ListHotQuestionInvoker 查询热点问题列表
func (c *MetaStudioClient) ListHotQuestionInvoker(request *model.ListHotQuestionRequest) *ListHotQuestionInvoker {
	requestDef := GenReqDefForListHotQuestion()
	return &ListHotQuestionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHotQuestion 查询热点问题详情
//
// 该接口用于查询热点问题详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowHotQuestion(request *model.ShowHotQuestionRequest) (*model.ShowHotQuestionResponse, error) {
	requestDef := GenReqDefForShowHotQuestion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHotQuestionResponse), nil
	}
}

// ShowHotQuestionInvoker 查询热点问题详情
func (c *MetaStudioClient) ShowHotQuestionInvoker(request *model.ShowHotQuestionRequest) *ShowHotQuestionInvoker {
	requestDef := GenReqDefForShowHotQuestion()
	return &ShowHotQuestionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHotQuestion 修改热点问题
//
// 该接口用于修改热点问题。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateHotQuestion(request *model.UpdateHotQuestionRequest) (*model.UpdateHotQuestionResponse, error) {
	requestDef := GenReqDefForUpdateHotQuestion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHotQuestionResponse), nil
	}
}

// UpdateHotQuestionInvoker 修改热点问题
func (c *MetaStudioClient) UpdateHotQuestionInvoker(request *model.UpdateHotQuestionRequest) *UpdateHotQuestionInvoker {
	requestDef := GenReqDefForUpdateHotQuestion()
	return &UpdateHotQuestionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateHotWords 创建热词记录
//
// 该接口用于创建热词记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateHotWords(request *model.CreateHotWordsRequest) (*model.CreateHotWordsResponse, error) {
	requestDef := GenReqDefForCreateHotWords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHotWordsResponse), nil
	}
}

// CreateHotWordsInvoker 创建热词记录
func (c *MetaStudioClient) CreateHotWordsInvoker(request *model.CreateHotWordsRequest) *CreateHotWordsInvoker {
	requestDef := GenReqDefForCreateHotWords()
	return &CreateHotWordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHotWords 删除热词记录
//
// 该接口用于删除热词记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteHotWords(request *model.DeleteHotWordsRequest) (*model.DeleteHotWordsResponse, error) {
	requestDef := GenReqDefForDeleteHotWords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHotWordsResponse), nil
	}
}

// DeleteHotWordsInvoker 删除热词记录
func (c *MetaStudioClient) DeleteHotWordsInvoker(request *model.DeleteHotWordsRequest) *DeleteHotWordsInvoker {
	requestDef := GenReqDefForDeleteHotWords()
	return &DeleteHotWordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHotWords 查询热词记录列表
//
// 该接口用于查询热词记录列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListHotWords(request *model.ListHotWordsRequest) (*model.ListHotWordsResponse, error) {
	requestDef := GenReqDefForListHotWords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHotWordsResponse), nil
	}
}

// ListHotWordsInvoker 查询热词记录列表
func (c *MetaStudioClient) ListHotWordsInvoker(request *model.ListHotWordsRequest) *ListHotWordsInvoker {
	requestDef := GenReqDefForListHotWords()
	return &ListHotWordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHotWords 查询配置热词记录详情
//
// 该接口用于查询热词记录详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowHotWords(request *model.ShowHotWordsRequest) (*model.ShowHotWordsResponse, error) {
	requestDef := GenReqDefForShowHotWords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHotWordsResponse), nil
	}
}

// ShowHotWordsInvoker 查询配置热词记录详情
func (c *MetaStudioClient) ShowHotWordsInvoker(request *model.ShowHotWordsRequest) *ShowHotWordsInvoker {
	requestDef := GenReqDefForShowHotWords()
	return &ShowHotWordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHotWordsSwitch 查询热词功能开关
//
// 该接口用于查询热词功能开关。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowHotWordsSwitch(request *model.ShowHotWordsSwitchRequest) (*model.ShowHotWordsSwitchResponse, error) {
	requestDef := GenReqDefForShowHotWordsSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHotWordsSwitchResponse), nil
	}
}

// ShowHotWordsSwitchInvoker 查询热词功能开关
func (c *MetaStudioClient) ShowHotWordsSwitchInvoker(request *model.ShowHotWordsSwitchRequest) *ShowHotWordsSwitchInvoker {
	requestDef := GenReqDefForShowHotWordsSwitch()
	return &ShowHotWordsSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHotWords 修改热词记录
//
// 该接口用于修改热词记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateHotWords(request *model.UpdateHotWordsRequest) (*model.UpdateHotWordsResponse, error) {
	requestDef := GenReqDefForUpdateHotWords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHotWordsResponse), nil
	}
}

// UpdateHotWordsInvoker 修改热词记录
func (c *MetaStudioClient) UpdateHotWordsInvoker(request *model.UpdateHotWordsRequest) *UpdateHotWordsInvoker {
	requestDef := GenReqDefForUpdateHotWords()
	return &UpdateHotWordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHotWordsSwitch 修改热词功能开关
//
// 该接口用于修改热词功能开关。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateHotWordsSwitch(request *model.UpdateHotWordsSwitchRequest) (*model.UpdateHotWordsSwitchResponse, error) {
	requestDef := GenReqDefForUpdateHotWordsSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHotWordsSwitchResponse), nil
	}
}

// UpdateHotWordsSwitchInvoker 修改热词功能开关
func (c *MetaStudioClient) UpdateHotWordsSwitchInvoker(request *model.UpdateHotWordsSwitchRequest) *UpdateHotWordsSwitchInvoker {
	requestDef := GenReqDefForUpdateHotWordsSwitch()
	return &UpdateHotWordsSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstructionLibrary 创建指令集
//
// 该接口用于创建指令集。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateInstructionLibrary(request *model.CreateInstructionLibraryRequest) (*model.CreateInstructionLibraryResponse, error) {
	requestDef := GenReqDefForCreateInstructionLibrary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstructionLibraryResponse), nil
	}
}

// CreateInstructionLibraryInvoker 创建指令集
func (c *MetaStudioClient) CreateInstructionLibraryInvoker(request *model.CreateInstructionLibraryRequest) *CreateInstructionLibraryInvoker {
	requestDef := GenReqDefForCreateInstructionLibrary()
	return &CreateInstructionLibraryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstructionLibrary 删除指令集
//
// 该接口用于删除指令集。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteInstructionLibrary(request *model.DeleteInstructionLibraryRequest) (*model.DeleteInstructionLibraryResponse, error) {
	requestDef := GenReqDefForDeleteInstructionLibrary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstructionLibraryResponse), nil
	}
}

// DeleteInstructionLibraryInvoker 删除指令集
func (c *MetaStudioClient) DeleteInstructionLibraryInvoker(request *model.DeleteInstructionLibraryRequest) *DeleteInstructionLibraryInvoker {
	requestDef := GenReqDefForDeleteInstructionLibrary()
	return &DeleteInstructionLibraryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstructionLibrary 查询指令集列表
//
// 该接口用于查询指令集列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListInstructionLibrary(request *model.ListInstructionLibraryRequest) (*model.ListInstructionLibraryResponse, error) {
	requestDef := GenReqDefForListInstructionLibrary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstructionLibraryResponse), nil
	}
}

// ListInstructionLibraryInvoker 查询指令集列表
func (c *MetaStudioClient) ListInstructionLibraryInvoker(request *model.ListInstructionLibraryRequest) *ListInstructionLibraryInvoker {
	requestDef := GenReqDefForListInstructionLibrary()
	return &ListInstructionLibraryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstructionLibrary 查询指令集详情
//
// 该接口用于查询指令集详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowInstructionLibrary(request *model.ShowInstructionLibraryRequest) (*model.ShowInstructionLibraryResponse, error) {
	requestDef := GenReqDefForShowInstructionLibrary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstructionLibraryResponse), nil
	}
}

// ShowInstructionLibraryInvoker 查询指令集详情
func (c *MetaStudioClient) ShowInstructionLibraryInvoker(request *model.ShowInstructionLibraryRequest) *ShowInstructionLibraryInvoker {
	requestDef := GenReqDefForShowInstructionLibrary()
	return &ShowInstructionLibraryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstructionLibrary 修改指令集
//
// 该接口用于修改指令集。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateInstructionLibrary(request *model.UpdateInstructionLibraryRequest) (*model.UpdateInstructionLibraryResponse, error) {
	requestDef := GenReqDefForUpdateInstructionLibrary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstructionLibraryResponse), nil
	}
}

// UpdateInstructionLibraryInvoker 修改指令集
func (c *MetaStudioClient) UpdateInstructionLibraryInvoker(request *model.UpdateInstructionLibraryRequest) *UpdateInstructionLibraryInvoker {
	requestDef := GenReqDefForUpdateInstructionLibrary()
	return &UpdateInstructionLibraryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstruction 创建指令
//
// 该接口用于创建指令。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateInstruction(request *model.CreateInstructionRequest) (*model.CreateInstructionResponse, error) {
	requestDef := GenReqDefForCreateInstruction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstructionResponse), nil
	}
}

// CreateInstructionInvoker 创建指令
func (c *MetaStudioClient) CreateInstructionInvoker(request *model.CreateInstructionRequest) *CreateInstructionInvoker {
	requestDef := GenReqDefForCreateInstruction()
	return &CreateInstructionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstruction 删除指令
//
// 该接口用于删除指令。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteInstruction(request *model.DeleteInstructionRequest) (*model.DeleteInstructionResponse, error) {
	requestDef := GenReqDefForDeleteInstruction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstructionResponse), nil
	}
}

// DeleteInstructionInvoker 删除指令
func (c *MetaStudioClient) DeleteInstructionInvoker(request *model.DeleteInstructionRequest) *DeleteInstructionInvoker {
	requestDef := GenReqDefForDeleteInstruction()
	return &DeleteInstructionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstruction 查询指令列表
//
// 该接口用于查询指令列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListInstruction(request *model.ListInstructionRequest) (*model.ListInstructionResponse, error) {
	requestDef := GenReqDefForListInstruction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstructionResponse), nil
	}
}

// ListInstructionInvoker 查询指令列表
func (c *MetaStudioClient) ListInstructionInvoker(request *model.ListInstructionRequest) *ListInstructionInvoker {
	requestDef := GenReqDefForListInstruction()
	return &ListInstructionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstruction 查询指令详情
//
// 该接口用于查询指令详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowInstruction(request *model.ShowInstructionRequest) (*model.ShowInstructionResponse, error) {
	requestDef := GenReqDefForShowInstruction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstructionResponse), nil
	}
}

// ShowInstructionInvoker 查询指令详情
func (c *MetaStudioClient) ShowInstructionInvoker(request *model.ShowInstructionRequest) *ShowInstructionInvoker {
	requestDef := GenReqDefForShowInstruction()
	return &ShowInstructionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstruction 修改指令
//
// 该接口用于修改指令。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateInstruction(request *model.UpdateInstructionRequest) (*model.UpdateInstructionResponse, error) {
	requestDef := GenReqDefForUpdateInstruction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstructionResponse), nil
	}
}

// UpdateInstructionInvoker 修改指令
func (c *MetaStudioClient) UpdateInstructionInvoker(request *model.UpdateInstructionRequest) *UpdateInstructionInvoker {
	requestDef := GenReqDefForUpdateInstruction()
	return &UpdateInstructionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInteractionRule 互动规则库增加规则
//
// 该接口用于互动规则库增加规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateInteractionRule(request *model.CreateInteractionRuleRequest) (*model.CreateInteractionRuleResponse, error) {
	requestDef := GenReqDefForCreateInteractionRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInteractionRuleResponse), nil
	}
}

// CreateInteractionRuleInvoker 互动规则库增加规则
func (c *MetaStudioClient) CreateInteractionRuleInvoker(request *model.CreateInteractionRuleRequest) *CreateInteractionRuleInvoker {
	requestDef := GenReqDefForCreateInteractionRule()
	return &CreateInteractionRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInteractionRuleGroup 创建智能直播间互动规则库
//
// 该接口用于创建智能直播间互动规则库。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateInteractionRuleGroup(request *model.CreateInteractionRuleGroupRequest) (*model.CreateInteractionRuleGroupResponse, error) {
	requestDef := GenReqDefForCreateInteractionRuleGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInteractionRuleGroupResponse), nil
	}
}

// CreateInteractionRuleGroupInvoker 创建智能直播间互动规则库
func (c *MetaStudioClient) CreateInteractionRuleGroupInvoker(request *model.CreateInteractionRuleGroupRequest) *CreateInteractionRuleGroupInvoker {
	requestDef := GenReqDefForCreateInteractionRuleGroup()
	return &CreateInteractionRuleGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInteractionRule 互动规则库删除某条规则
//
// 该接口用于互动规则库修改删除规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteInteractionRule(request *model.DeleteInteractionRuleRequest) (*model.DeleteInteractionRuleResponse, error) {
	requestDef := GenReqDefForDeleteInteractionRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInteractionRuleResponse), nil
	}
}

// DeleteInteractionRuleInvoker 互动规则库删除某条规则
func (c *MetaStudioClient) DeleteInteractionRuleInvoker(request *model.DeleteInteractionRuleRequest) *DeleteInteractionRuleInvoker {
	requestDef := GenReqDefForDeleteInteractionRule()
	return &DeleteInteractionRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInteractionRuleGroup 删除智能直播间互动规则库
//
// 该接口用于删除智能直播间互动规则库。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteInteractionRuleGroup(request *model.DeleteInteractionRuleGroupRequest) (*model.DeleteInteractionRuleGroupResponse, error) {
	requestDef := GenReqDefForDeleteInteractionRuleGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInteractionRuleGroupResponse), nil
	}
}

// DeleteInteractionRuleGroupInvoker 删除智能直播间互动规则库
func (c *MetaStudioClient) DeleteInteractionRuleGroupInvoker(request *model.DeleteInteractionRuleGroupRequest) *DeleteInteractionRuleGroupInvoker {
	requestDef := GenReqDefForDeleteInteractionRuleGroup()
	return &DeleteInteractionRuleGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInteractionRuleGroups 查询智能直播间互动规则库列表
//
// 该接口用于智能直播间互动规则库列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListInteractionRuleGroups(request *model.ListInteractionRuleGroupsRequest) (*model.ListInteractionRuleGroupsResponse, error) {
	requestDef := GenReqDefForListInteractionRuleGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInteractionRuleGroupsResponse), nil
	}
}

// ListInteractionRuleGroupsInvoker 查询智能直播间互动规则库列表
func (c *MetaStudioClient) ListInteractionRuleGroupsInvoker(request *model.ListInteractionRuleGroupsRequest) *ListInteractionRuleGroupsInvoker {
	requestDef := GenReqDefForListInteractionRuleGroups()
	return &ListInteractionRuleGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInteractionRuleGroupsSummary 查询智能直播间互动规则库概要列表
//
// 该接口用于智能直播间互动规则库概要列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListInteractionRuleGroupsSummary(request *model.ListInteractionRuleGroupsSummaryRequest) (*model.ListInteractionRuleGroupsSummaryResponse, error) {
	requestDef := GenReqDefForListInteractionRuleGroupsSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInteractionRuleGroupsSummaryResponse), nil
	}
}

// ListInteractionRuleGroupsSummaryInvoker 查询智能直播间互动规则库概要列表
func (c *MetaStudioClient) ListInteractionRuleGroupsSummaryInvoker(request *model.ListInteractionRuleGroupsSummaryRequest) *ListInteractionRuleGroupsSummaryInvoker {
	requestDef := GenReqDefForListInteractionRuleGroupsSummary()
	return &ListInteractionRuleGroupsSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInteractionRuleGroup 查询智能直播间互动规则库详情
//
// 该接口用于查询智能直播间互动规则库详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowInteractionRuleGroup(request *model.ShowInteractionRuleGroupRequest) (*model.ShowInteractionRuleGroupResponse, error) {
	requestDef := GenReqDefForShowInteractionRuleGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInteractionRuleGroupResponse), nil
	}
}

// ShowInteractionRuleGroupInvoker 查询智能直播间互动规则库详情
func (c *MetaStudioClient) ShowInteractionRuleGroupInvoker(request *model.ShowInteractionRuleGroupRequest) *ShowInteractionRuleGroupInvoker {
	requestDef := GenReqDefForShowInteractionRuleGroup()
	return &ShowInteractionRuleGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInteractionRule 互动规则库修改某条规则
//
// 该接口用于互动规则库修改某条规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateInteractionRule(request *model.UpdateInteractionRuleRequest) (*model.UpdateInteractionRuleResponse, error) {
	requestDef := GenReqDefForUpdateInteractionRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInteractionRuleResponse), nil
	}
}

// UpdateInteractionRuleInvoker 互动规则库修改某条规则
func (c *MetaStudioClient) UpdateInteractionRuleInvoker(request *model.UpdateInteractionRuleRequest) *UpdateInteractionRuleInvoker {
	requestDef := GenReqDefForUpdateInteractionRule()
	return &UpdateInteractionRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInteractionRuleGroup 更新智能直播间互动规则库
//
// 该接口用于更新智能直播间互动规则库。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateInteractionRuleGroup(request *model.UpdateInteractionRuleGroupRequest) (*model.UpdateInteractionRuleGroupResponse, error) {
	requestDef := GenReqDefForUpdateInteractionRuleGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInteractionRuleGroupResponse), nil
	}
}

// UpdateInteractionRuleGroupInvoker 更新智能直播间互动规则库
func (c *MetaStudioClient) UpdateInteractionRuleGroupInvoker(request *model.UpdateInteractionRuleGroupRequest) *UpdateInteractionRuleGroupInvoker {
	requestDef := GenReqDefForUpdateInteractionRuleGroup()
	return &UpdateInteractionRuleGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInteractiveChat 交互助手对话
//
// 该接口用于交互助手对话。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateInteractiveChat(request *model.CreateInteractiveChatRequest) (*model.CreateInteractiveChatResponse, error) {
	requestDef := GenReqDefForCreateInteractiveChat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInteractiveChatResponse), nil
	}
}

// CreateInteractiveChatInvoker 交互助手对话
func (c *MetaStudioClient) CreateInteractiveChatInvoker(request *model.CreateInteractiveChatRequest) *CreateInteractiveChatInvoker {
	requestDef := GenReqDefForCreateInteractiveChat()
	return &CreateInteractiveChatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateIntentAndQuestion 创建知识库意图和问法
//
// 该接口用于创建知识库意图和问法。一个意图包含一个主题，一个答案，若干个问法等。接口使用限制详见[API使用限制](metastudio_02_0000.xml)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateIntentAndQuestion(request *model.CreateIntentAndQuestionRequest) (*model.CreateIntentAndQuestionResponse, error) {
	requestDef := GenReqDefForCreateIntentAndQuestion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateIntentAndQuestionResponse), nil
	}
}

// CreateIntentAndQuestionInvoker 创建知识库意图和问法
func (c *MetaStudioClient) CreateIntentAndQuestionInvoker(request *model.CreateIntentAndQuestionRequest) *CreateIntentAndQuestionInvoker {
	requestDef := GenReqDefForCreateIntentAndQuestion()
	return &CreateIntentAndQuestionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateKnowledgeIntent 创建知识库意图
//
// 该接口用于创建知识库意图。一个意图包含一个主题，一个答案，若干个问法等。接口使用限制详见[API使用限制](metastudio_02_0000.xml)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateKnowledgeIntent(request *model.CreateKnowledgeIntentRequest) (*model.CreateKnowledgeIntentResponse, error) {
	requestDef := GenReqDefForCreateKnowledgeIntent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateKnowledgeIntentResponse), nil
	}
}

// CreateKnowledgeIntentInvoker 创建知识库意图
func (c *MetaStudioClient) CreateKnowledgeIntentInvoker(request *model.CreateKnowledgeIntentRequest) *CreateKnowledgeIntentInvoker {
	requestDef := GenReqDefForCreateKnowledgeIntent()
	return &CreateKnowledgeIntentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteKnowledgeIntent 删除知识库意图
//
// 该接口用于删除知识库意图。接口使用限制详见[API使用限制](metastudio_02_0000.xml)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteKnowledgeIntent(request *model.DeleteKnowledgeIntentRequest) (*model.DeleteKnowledgeIntentResponse, error) {
	requestDef := GenReqDefForDeleteKnowledgeIntent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteKnowledgeIntentResponse), nil
	}
}

// DeleteKnowledgeIntentInvoker 删除知识库意图
func (c *MetaStudioClient) DeleteKnowledgeIntentInvoker(request *model.DeleteKnowledgeIntentRequest) *DeleteKnowledgeIntentInvoker {
	requestDef := GenReqDefForDeleteKnowledgeIntent()
	return &DeleteKnowledgeIntentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListKnowledgeIntent 查询知识库意图列表
//
// 该接口用于查询知识库意图列表。接口使用限制详见[API使用限制](metastudio_02_0000.xml)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListKnowledgeIntent(request *model.ListKnowledgeIntentRequest) (*model.ListKnowledgeIntentResponse, error) {
	requestDef := GenReqDefForListKnowledgeIntent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKnowledgeIntentResponse), nil
	}
}

// ListKnowledgeIntentInvoker 查询知识库意图列表
func (c *MetaStudioClient) ListKnowledgeIntentInvoker(request *model.ListKnowledgeIntentRequest) *ListKnowledgeIntentInvoker {
	requestDef := GenReqDefForListKnowledgeIntent()
	return &ListKnowledgeIntentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowKnowledgeIntent 查询知识库意图详情
//
// 该接口用于查询知识库意图详情。接口使用限制详见[API使用限制](metastudio_02_0000.xml)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowKnowledgeIntent(request *model.ShowKnowledgeIntentRequest) (*model.ShowKnowledgeIntentResponse, error) {
	requestDef := GenReqDefForShowKnowledgeIntent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKnowledgeIntentResponse), nil
	}
}

// ShowKnowledgeIntentInvoker 查询知识库意图详情
func (c *MetaStudioClient) ShowKnowledgeIntentInvoker(request *model.ShowKnowledgeIntentRequest) *ShowKnowledgeIntentInvoker {
	requestDef := GenReqDefForShowKnowledgeIntent()
	return &ShowKnowledgeIntentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateKnowledgeIntent 修改知识库意图
//
// 该接口用于修改知识库意图。接口使用限制详见[API使用限制](metastudio_02_0000.xml)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateKnowledgeIntent(request *model.UpdateKnowledgeIntentRequest) (*model.UpdateKnowledgeIntentResponse, error) {
	requestDef := GenReqDefForUpdateKnowledgeIntent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateKnowledgeIntentResponse), nil
	}
}

// UpdateKnowledgeIntentInvoker 修改知识库意图
func (c *MetaStudioClient) UpdateKnowledgeIntentInvoker(request *model.UpdateKnowledgeIntentRequest) *UpdateKnowledgeIntentInvoker {
	requestDef := GenReqDefForUpdateKnowledgeIntent()
	return &UpdateKnowledgeIntentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckRecallKnowledgeLibrary 知识库召回测试
//
// 该接口用于知识库召回测试。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CheckRecallKnowledgeLibrary(request *model.CheckRecallKnowledgeLibraryRequest) (*model.CheckRecallKnowledgeLibraryResponse, error) {
	requestDef := GenReqDefForCheckRecallKnowledgeLibrary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckRecallKnowledgeLibraryResponse), nil
	}
}

// CheckRecallKnowledgeLibraryInvoker 知识库召回测试
func (c *MetaStudioClient) CheckRecallKnowledgeLibraryInvoker(request *model.CheckRecallKnowledgeLibraryRequest) *CheckRecallKnowledgeLibraryInvoker {
	requestDef := GenReqDefForCheckRecallKnowledgeLibrary()
	return &CheckRecallKnowledgeLibraryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateKnowledgeLibrary 创建知识库
//
// 该接口用于创建知识库。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateKnowledgeLibrary(request *model.CreateKnowledgeLibraryRequest) (*model.CreateKnowledgeLibraryResponse, error) {
	requestDef := GenReqDefForCreateKnowledgeLibrary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateKnowledgeLibraryResponse), nil
	}
}

// CreateKnowledgeLibraryInvoker 创建知识库
func (c *MetaStudioClient) CreateKnowledgeLibraryInvoker(request *model.CreateKnowledgeLibraryRequest) *CreateKnowledgeLibraryInvoker {
	requestDef := GenReqDefForCreateKnowledgeLibrary()
	return &CreateKnowledgeLibraryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteKnowledgeLibrary 删除知识库
//
// 该接口用于删除知识库。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteKnowledgeLibrary(request *model.DeleteKnowledgeLibraryRequest) (*model.DeleteKnowledgeLibraryResponse, error) {
	requestDef := GenReqDefForDeleteKnowledgeLibrary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteKnowledgeLibraryResponse), nil
	}
}

// DeleteKnowledgeLibraryInvoker 删除知识库
func (c *MetaStudioClient) DeleteKnowledgeLibraryInvoker(request *model.DeleteKnowledgeLibraryRequest) *DeleteKnowledgeLibraryInvoker {
	requestDef := GenReqDefForDeleteKnowledgeLibrary()
	return &DeleteKnowledgeLibraryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListKnowledgeLibrary 查询知识库列表
//
// 该接口用于查询知识库列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListKnowledgeLibrary(request *model.ListKnowledgeLibraryRequest) (*model.ListKnowledgeLibraryResponse, error) {
	requestDef := GenReqDefForListKnowledgeLibrary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKnowledgeLibraryResponse), nil
	}
}

// ListKnowledgeLibraryInvoker 查询知识库列表
func (c *MetaStudioClient) ListKnowledgeLibraryInvoker(request *model.ListKnowledgeLibraryRequest) *ListKnowledgeLibraryInvoker {
	requestDef := GenReqDefForListKnowledgeLibrary()
	return &ListKnowledgeLibraryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowKnowledgeLibrary 查询知识库详情
//
// 该接口用于查询知识库详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowKnowledgeLibrary(request *model.ShowKnowledgeLibraryRequest) (*model.ShowKnowledgeLibraryResponse, error) {
	requestDef := GenReqDefForShowKnowledgeLibrary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKnowledgeLibraryResponse), nil
	}
}

// ShowKnowledgeLibraryInvoker 查询知识库详情
func (c *MetaStudioClient) ShowKnowledgeLibraryInvoker(request *model.ShowKnowledgeLibraryRequest) *ShowKnowledgeLibraryInvoker {
	requestDef := GenReqDefForShowKnowledgeLibrary()
	return &ShowKnowledgeLibraryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateKnowledgeLibrary 修改知识库
//
// 该接口用于修改知识库。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateKnowledgeLibrary(request *model.UpdateKnowledgeLibraryRequest) (*model.UpdateKnowledgeLibraryResponse, error) {
	requestDef := GenReqDefForUpdateKnowledgeLibrary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateKnowledgeLibraryResponse), nil
	}
}

// UpdateKnowledgeLibraryInvoker 修改知识库
func (c *MetaStudioClient) UpdateKnowledgeLibraryInvoker(request *model.UpdateKnowledgeLibraryRequest) *UpdateKnowledgeLibraryInvoker {
	requestDef := GenReqDefForUpdateKnowledgeLibrary()
	return &UpdateKnowledgeLibraryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateBatchKnowledgeQuestion 批量创建知识库问法
//
// 该接口用于批量创建知识库问法。接口使用限制详见[API使用限制](metastudio_02_0000.xml)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateBatchKnowledgeQuestion(request *model.CreateBatchKnowledgeQuestionRequest) (*model.CreateBatchKnowledgeQuestionResponse, error) {
	requestDef := GenReqDefForCreateBatchKnowledgeQuestion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBatchKnowledgeQuestionResponse), nil
	}
}

// CreateBatchKnowledgeQuestionInvoker 批量创建知识库问法
func (c *MetaStudioClient) CreateBatchKnowledgeQuestionInvoker(request *model.CreateBatchKnowledgeQuestionRequest) *CreateBatchKnowledgeQuestionInvoker {
	requestDef := GenReqDefForCreateBatchKnowledgeQuestion()
	return &CreateBatchKnowledgeQuestionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateKnowledgeQuestion 创建知识库问法
//
// 该接口用于创建知识库问法。接口使用限制详见[API使用限制](metastudio_02_0000.xml)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateKnowledgeQuestion(request *model.CreateKnowledgeQuestionRequest) (*model.CreateKnowledgeQuestionResponse, error) {
	requestDef := GenReqDefForCreateKnowledgeQuestion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateKnowledgeQuestionResponse), nil
	}
}

// CreateKnowledgeQuestionInvoker 创建知识库问法
func (c *MetaStudioClient) CreateKnowledgeQuestionInvoker(request *model.CreateKnowledgeQuestionRequest) *CreateKnowledgeQuestionInvoker {
	requestDef := GenReqDefForCreateKnowledgeQuestion()
	return &CreateKnowledgeQuestionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteKnowledgeQuestion 删除知识库问法
//
// 该接口用于删除知识库问法。接口使用限制详见[API使用限制](metastudio_02_0000.xml)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteKnowledgeQuestion(request *model.DeleteKnowledgeQuestionRequest) (*model.DeleteKnowledgeQuestionResponse, error) {
	requestDef := GenReqDefForDeleteKnowledgeQuestion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteKnowledgeQuestionResponse), nil
	}
}

// DeleteKnowledgeQuestionInvoker 删除知识库问法
func (c *MetaStudioClient) DeleteKnowledgeQuestionInvoker(request *model.DeleteKnowledgeQuestionRequest) *DeleteKnowledgeQuestionInvoker {
	requestDef := GenReqDefForDeleteKnowledgeQuestion()
	return &DeleteKnowledgeQuestionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListKnowledgeQuestion 查询知识库问法列表
//
// 该接口用于查询知识库问法列表。接口使用限制详见[API使用限制](metastudio_02_0000.xml)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListKnowledgeQuestion(request *model.ListKnowledgeQuestionRequest) (*model.ListKnowledgeQuestionResponse, error) {
	requestDef := GenReqDefForListKnowledgeQuestion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKnowledgeQuestionResponse), nil
	}
}

// ListKnowledgeQuestionInvoker 查询知识库问法列表
func (c *MetaStudioClient) ListKnowledgeQuestionInvoker(request *model.ListKnowledgeQuestionRequest) *ListKnowledgeQuestionInvoker {
	requestDef := GenReqDefForListKnowledgeQuestion()
	return &ListKnowledgeQuestionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowKnowledgeQuestion 查询知识库问法详情
//
// 该接口用于查询知识库问法详情。接口使用限制详见[API使用限制](metastudio_02_0000.xml)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowKnowledgeQuestion(request *model.ShowKnowledgeQuestionRequest) (*model.ShowKnowledgeQuestionResponse, error) {
	requestDef := GenReqDefForShowKnowledgeQuestion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKnowledgeQuestionResponse), nil
	}
}

// ShowKnowledgeQuestionInvoker 查询知识库问法详情
func (c *MetaStudioClient) ShowKnowledgeQuestionInvoker(request *model.ShowKnowledgeQuestionRequest) *ShowKnowledgeQuestionInvoker {
	requestDef := GenReqDefForShowKnowledgeQuestion()
	return &ShowKnowledgeQuestionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateBatchKnowledgeQuestion 批量修改知识库问法
//
// 该接口用于批量修改知识库问法。接口使用限制详见[API使用限制](metastudio_02_0000.xml)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateBatchKnowledgeQuestion(request *model.UpdateBatchKnowledgeQuestionRequest) (*model.UpdateBatchKnowledgeQuestionResponse, error) {
	requestDef := GenReqDefForUpdateBatchKnowledgeQuestion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBatchKnowledgeQuestionResponse), nil
	}
}

// UpdateBatchKnowledgeQuestionInvoker 批量修改知识库问法
func (c *MetaStudioClient) UpdateBatchKnowledgeQuestionInvoker(request *model.UpdateBatchKnowledgeQuestionRequest) *UpdateBatchKnowledgeQuestionInvoker {
	requestDef := GenReqDefForUpdateBatchKnowledgeQuestion()
	return &UpdateBatchKnowledgeQuestionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateKnowledgeQuestion 修改知识库问法
//
// 该接口用于修改知识库问法。接口使用限制详见[API使用限制](metastudio_02_0000.xml)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateKnowledgeQuestion(request *model.UpdateKnowledgeQuestionRequest) (*model.UpdateKnowledgeQuestionResponse, error) {
	requestDef := GenReqDefForUpdateKnowledgeQuestion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateKnowledgeQuestionResponse), nil
	}
}

// UpdateKnowledgeQuestionInvoker 修改知识库问法
func (c *MetaStudioClient) UpdateKnowledgeQuestionInvoker(request *model.UpdateKnowledgeQuestionRequest) *UpdateKnowledgeQuestionInvoker {
	requestDef := GenReqDefForUpdateKnowledgeQuestion()
	return &UpdateKnowledgeQuestionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateKnowledgeSkill 创建知识库技能
//
// 该接口用于创建知识库技能。一个技能用于特定场景的交互问答，包含若干个意图等。接口使用限制详见[API使用限制](metastudio_02_0000.xml)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateKnowledgeSkill(request *model.CreateKnowledgeSkillRequest) (*model.CreateKnowledgeSkillResponse, error) {
	requestDef := GenReqDefForCreateKnowledgeSkill()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateKnowledgeSkillResponse), nil
	}
}

// CreateKnowledgeSkillInvoker 创建知识库技能
func (c *MetaStudioClient) CreateKnowledgeSkillInvoker(request *model.CreateKnowledgeSkillRequest) *CreateKnowledgeSkillInvoker {
	requestDef := GenReqDefForCreateKnowledgeSkill()
	return &CreateKnowledgeSkillInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteKnowledgeSkill 删除知识库技能
//
// 该接口用于删除知识库技能。接口使用限制详见[API使用限制](metastudio_02_0000.xml)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteKnowledgeSkill(request *model.DeleteKnowledgeSkillRequest) (*model.DeleteKnowledgeSkillResponse, error) {
	requestDef := GenReqDefForDeleteKnowledgeSkill()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteKnowledgeSkillResponse), nil
	}
}

// DeleteKnowledgeSkillInvoker 删除知识库技能
func (c *MetaStudioClient) DeleteKnowledgeSkillInvoker(request *model.DeleteKnowledgeSkillRequest) *DeleteKnowledgeSkillInvoker {
	requestDef := GenReqDefForDeleteKnowledgeSkill()
	return &DeleteKnowledgeSkillInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportKnowledgeSkill 导出知识库技能
//
// 该接口用于导出知识库技能。接口使用限制详见[API使用限制](metastudio_02_0000.xml)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ExportKnowledgeSkill(request *model.ExportKnowledgeSkillRequest) (*model.ExportKnowledgeSkillResponse, error) {
	requestDef := GenReqDefForExportKnowledgeSkill()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportKnowledgeSkillResponse), nil
	}
}

// ExportKnowledgeSkillInvoker 导出知识库技能
func (c *MetaStudioClient) ExportKnowledgeSkillInvoker(request *model.ExportKnowledgeSkillRequest) *ExportKnowledgeSkillInvoker {
	requestDef := GenReqDefForExportKnowledgeSkill()
	return &ExportKnowledgeSkillInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListKnowledgeSkill 查询知识库技能列表
//
// 该接口用于查询知识库技能列表。接口使用限制详见[API使用限制](metastudio_02_0000.xml)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListKnowledgeSkill(request *model.ListKnowledgeSkillRequest) (*model.ListKnowledgeSkillResponse, error) {
	requestDef := GenReqDefForListKnowledgeSkill()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKnowledgeSkillResponse), nil
	}
}

// ListKnowledgeSkillInvoker 查询知识库技能列表
func (c *MetaStudioClient) ListKnowledgeSkillInvoker(request *model.ListKnowledgeSkillRequest) *ListKnowledgeSkillInvoker {
	requestDef := GenReqDefForListKnowledgeSkill()
	return &ListKnowledgeSkillInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowKnowledgeSkill 查询知识库技能详情
//
// 该接口用于查询知识库技能详情。接口使用限制详见[API使用限制](metastudio_02_0000.xml)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowKnowledgeSkill(request *model.ShowKnowledgeSkillRequest) (*model.ShowKnowledgeSkillResponse, error) {
	requestDef := GenReqDefForShowKnowledgeSkill()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKnowledgeSkillResponse), nil
	}
}

// ShowKnowledgeSkillInvoker 查询知识库技能详情
func (c *MetaStudioClient) ShowKnowledgeSkillInvoker(request *model.ShowKnowledgeSkillRequest) *ShowKnowledgeSkillInvoker {
	requestDef := GenReqDefForShowKnowledgeSkill()
	return &ShowKnowledgeSkillInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateKnowledgeSkill 修改知识库技能
//
// 该接口用于修改知识库技能。接口使用限制详见[API使用限制](metastudio_02_0000.xml)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateKnowledgeSkill(request *model.UpdateKnowledgeSkillRequest) (*model.UpdateKnowledgeSkillResponse, error) {
	requestDef := GenReqDefForUpdateKnowledgeSkill()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateKnowledgeSkillResponse), nil
	}
}

// UpdateKnowledgeSkillInvoker 修改知识库技能
func (c *MetaStudioClient) UpdateKnowledgeSkillInvoker(request *model.UpdateKnowledgeSkillRequest) *UpdateKnowledgeSkillInvoker {
	requestDef := GenReqDefForUpdateKnowledgeSkill()
	return &UpdateKnowledgeSkillInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLivePlatform 创建第三方直播平台
//
// 该接口用于创建第三方直播平台。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateLivePlatform(request *model.CreateLivePlatformRequest) (*model.CreateLivePlatformResponse, error) {
	requestDef := GenReqDefForCreateLivePlatform()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLivePlatformResponse), nil
	}
}

// CreateLivePlatformInvoker 创建第三方直播平台
func (c *MetaStudioClient) CreateLivePlatformInvoker(request *model.CreateLivePlatformRequest) *CreateLivePlatformInvoker {
	requestDef := GenReqDefForCreateLivePlatform()
	return &CreateLivePlatformInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLivePlatform 删除第三方直播平台信息
//
// 该接口用于删除第三方直播平台信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteLivePlatform(request *model.DeleteLivePlatformRequest) (*model.DeleteLivePlatformResponse, error) {
	requestDef := GenReqDefForDeleteLivePlatform()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLivePlatformResponse), nil
	}
}

// DeleteLivePlatformInvoker 删除第三方直播平台信息
func (c *MetaStudioClient) DeleteLivePlatformInvoker(request *model.DeleteLivePlatformRequest) *DeleteLivePlatformInvoker {
	requestDef := GenReqDefForDeleteLivePlatform()
	return &DeleteLivePlatformInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLivePlatformProducts 查询第三方直播平台商品列表
//
// 该接口用于查询第三方直播平台商品列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListLivePlatformProducts(request *model.ListLivePlatformProductsRequest) (*model.ListLivePlatformProductsResponse, error) {
	requestDef := GenReqDefForListLivePlatformProducts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLivePlatformProductsResponse), nil
	}
}

// ListLivePlatformProductsInvoker 查询第三方直播平台商品列表
func (c *MetaStudioClient) ListLivePlatformProductsInvoker(request *model.ListLivePlatformProductsRequest) *ListLivePlatformProductsInvoker {
	requestDef := GenReqDefForListLivePlatformProducts()
	return &ListLivePlatformProductsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLivePlatforms 查询直播平台列表
//
// 该接口用于查询直播平台列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListLivePlatforms(request *model.ListLivePlatformsRequest) (*model.ListLivePlatformsResponse, error) {
	requestDef := GenReqDefForListLivePlatforms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLivePlatformsResponse), nil
	}
}

// ListLivePlatformsInvoker 查询直播平台列表
func (c *MetaStudioClient) ListLivePlatformsInvoker(request *model.ListLivePlatformsRequest) *ListLivePlatformsInvoker {
	requestDef := GenReqDefForListLivePlatforms()
	return &ListLivePlatformsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLivePlatform 查询第三方直播平台信息
//
// 该接口用于查询第三方直播平台信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowLivePlatform(request *model.ShowLivePlatformRequest) (*model.ShowLivePlatformResponse, error) {
	requestDef := GenReqDefForShowLivePlatform()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLivePlatformResponse), nil
	}
}

// ShowLivePlatformInvoker 查询第三方直播平台信息
func (c *MetaStudioClient) ShowLivePlatformInvoker(request *model.ShowLivePlatformRequest) *ShowLivePlatformInvoker {
	requestDef := GenReqDefForShowLivePlatform()
	return &ShowLivePlatformInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLivePlatformAccessType 查询直播平台对接方式
//
// 该接口用于直播平台对接方式。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowLivePlatformAccessType(request *model.ShowLivePlatformAccessTypeRequest) (*model.ShowLivePlatformAccessTypeResponse, error) {
	requestDef := GenReqDefForShowLivePlatformAccessType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLivePlatformAccessTypeResponse), nil
	}
}

// ShowLivePlatformAccessTypeInvoker 查询直播平台对接方式
func (c *MetaStudioClient) ShowLivePlatformAccessTypeInvoker(request *model.ShowLivePlatformAccessTypeRequest) *ShowLivePlatformAccessTypeInvoker {
	requestDef := GenReqDefForShowLivePlatformAccessType()
	return &ShowLivePlatformAccessTypeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLivePlatform 更新第三方直播平台信息
//
// 该接口用于更新第三方直播平台信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateLivePlatform(request *model.UpdateLivePlatformRequest) (*model.UpdateLivePlatformResponse, error) {
	requestDef := GenReqDefForUpdateLivePlatform()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLivePlatformResponse), nil
	}
}

// UpdateLivePlatformInvoker 更新第三方直播平台信息
func (c *MetaStudioClient) UpdateLivePlatformInvoker(request *model.UpdateLivePlatformRequest) *UpdateLivePlatformInvoker {
	requestDef := GenReqDefForUpdateLivePlatform()
	return &UpdateLivePlatformInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLlmConfig 创建大语言模型配置
//
// 该接口用于创建大语言模型配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateLlmConfig(request *model.CreateLlmConfigRequest) (*model.CreateLlmConfigResponse, error) {
	requestDef := GenReqDefForCreateLlmConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLlmConfigResponse), nil
	}
}

// CreateLlmConfigInvoker 创建大语言模型配置
func (c *MetaStudioClient) CreateLlmConfigInvoker(request *model.CreateLlmConfigRequest) *CreateLlmConfigInvoker {
	requestDef := GenReqDefForCreateLlmConfig()
	return &CreateLlmConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLlmConfig 删除大语言模型配置
//
// 该接口用于删除大语言模型配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteLlmConfig(request *model.DeleteLlmConfigRequest) (*model.DeleteLlmConfigResponse, error) {
	requestDef := GenReqDefForDeleteLlmConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLlmConfigResponse), nil
	}
}

// DeleteLlmConfigInvoker 删除大语言模型配置
func (c *MetaStudioClient) DeleteLlmConfigInvoker(request *model.DeleteLlmConfigRequest) *DeleteLlmConfigInvoker {
	requestDef := GenReqDefForDeleteLlmConfig()
	return &DeleteLlmConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLlmConfig 查询大语言模型配置列表
//
// 该接口用于查询大语言模型配置列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListLlmConfig(request *model.ListLlmConfigRequest) (*model.ListLlmConfigResponse, error) {
	requestDef := GenReqDefForListLlmConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLlmConfigResponse), nil
	}
}

// ListLlmConfigInvoker 查询大语言模型配置列表
func (c *MetaStudioClient) ListLlmConfigInvoker(request *model.ListLlmConfigRequest) *ListLlmConfigInvoker {
	requestDef := GenReqDefForListLlmConfig()
	return &ListLlmConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLlmConfig 查询大语言模型配置详情
//
// 该接口用于查询大语言模型配置详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowLlmConfig(request *model.ShowLlmConfigRequest) (*model.ShowLlmConfigResponse, error) {
	requestDef := GenReqDefForShowLlmConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLlmConfigResponse), nil
	}
}

// ShowLlmConfigInvoker 查询大语言模型配置详情
func (c *MetaStudioClient) ShowLlmConfigInvoker(request *model.ShowLlmConfigRequest) *ShowLlmConfigInvoker {
	requestDef := GenReqDefForShowLlmConfig()
	return &ShowLlmConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLlmConfig 修改大语言模型配置
//
// 该接口用于修改大语言模型配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateLlmConfig(request *model.UpdateLlmConfigRequest) (*model.UpdateLlmConfigResponse, error) {
	requestDef := GenReqDefForUpdateLlmConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLlmConfigResponse), nil
	}
}

// UpdateLlmConfigInvoker 修改大语言模型配置
func (c *MetaStudioClient) UpdateLlmConfigInvoker(request *model.UpdateLlmConfigRequest) *UpdateLlmConfigInvoker {
	requestDef := GenReqDefForUpdateLlmConfig()
	return &UpdateLlmConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMcpServer 创建MCP服务端对接配置
//
// 该接口用于创建MCP服务端对接配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateMcpServer(request *model.CreateMcpServerRequest) (*model.CreateMcpServerResponse, error) {
	requestDef := GenReqDefForCreateMcpServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMcpServerResponse), nil
	}
}

// CreateMcpServerInvoker 创建MCP服务端对接配置
func (c *MetaStudioClient) CreateMcpServerInvoker(request *model.CreateMcpServerRequest) *CreateMcpServerInvoker {
	requestDef := GenReqDefForCreateMcpServer()
	return &CreateMcpServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMcpServer 删除MCP服务端对接配置
//
// 该接口用于删除MCP服务端对接配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteMcpServer(request *model.DeleteMcpServerRequest) (*model.DeleteMcpServerResponse, error) {
	requestDef := GenReqDefForDeleteMcpServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMcpServerResponse), nil
	}
}

// DeleteMcpServerInvoker 删除MCP服务端对接配置
func (c *MetaStudioClient) DeleteMcpServerInvoker(request *model.DeleteMcpServerRequest) *DeleteMcpServerInvoker {
	requestDef := GenReqDefForDeleteMcpServer()
	return &DeleteMcpServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMcpServer 查询MCP服务端对接配置列表
//
// 该接口用于查询MCP服务端对接配置列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListMcpServer(request *model.ListMcpServerRequest) (*model.ListMcpServerResponse, error) {
	requestDef := GenReqDefForListMcpServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMcpServerResponse), nil
	}
}

// ListMcpServerInvoker 查询MCP服务端对接配置列表
func (c *MetaStudioClient) ListMcpServerInvoker(request *model.ListMcpServerRequest) *ListMcpServerInvoker {
	requestDef := GenReqDefForListMcpServer()
	return &ListMcpServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMcpServer 查询MCP服务端对接配置详情
//
// 该接口用于查询MCP服务端对接配置详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowMcpServer(request *model.ShowMcpServerRequest) (*model.ShowMcpServerResponse, error) {
	requestDef := GenReqDefForShowMcpServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMcpServerResponse), nil
	}
}

// ShowMcpServerInvoker 查询MCP服务端对接配置详情
func (c *MetaStudioClient) ShowMcpServerInvoker(request *model.ShowMcpServerRequest) *ShowMcpServerInvoker {
	requestDef := GenReqDefForShowMcpServer()
	return &ShowMcpServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMcpServer 修改MCP服务端对接配置
//
// 该接口用于修改MCP服务端对接配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateMcpServer(request *model.UpdateMcpServerRequest) (*model.UpdateMcpServerResponse, error) {
	requestDef := GenReqDefForUpdateMcpServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMcpServerResponse), nil
	}
}

// UpdateMcpServerInvoker 修改MCP服务端对接配置
func (c *MetaStudioClient) UpdateMcpServerInvoker(request *model.UpdateMcpServerRequest) *UpdateMcpServerInvoker {
	requestDef := GenReqDefForUpdateMcpServer()
	return &UpdateMcpServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateOnceCode 创建一次性鉴权码
//
// 该接口用于创建一次性鉴权码，有效期5分钟，鉴权码只能使用一次，每次使用后需要重新获取。
// &gt; 接口只能通过第三方后台调用，不能在浏览器前台直接调用，否则会有跨域问题。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateOnceCode(request *model.CreateOnceCodeRequest) (*model.CreateOnceCodeResponse, error) {
	requestDef := GenReqDefForCreateOnceCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOnceCodeResponse), nil
	}
}

// CreateOnceCodeInvoker 创建一次性鉴权码
func (c *MetaStudioClient) CreateOnceCodeInvoker(request *model.CreateOnceCodeRequest) *CreateOnceCodeInvoker {
	requestDef := GenReqDefForCreateOnceCode()
	return &CreateOnceCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMetaStudioOrders 订购metastudio云服务产品
//
// 该接口用于订购MetaStudio服务的包周期,一次性,按需套餐包产品
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateMetaStudioOrders(request *model.CreateMetaStudioOrdersRequest) (*model.CreateMetaStudioOrdersResponse, error) {
	requestDef := GenReqDefForCreateMetaStudioOrders()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMetaStudioOrdersResponse), nil
	}
}

// CreateMetaStudioOrdersInvoker 订购metastudio云服务产品
func (c *MetaStudioClient) CreateMetaStudioOrdersInvoker(request *model.CreateMetaStudioOrdersRequest) *CreateMetaStudioOrdersInvoker {
	requestDef := GenReqDefForCreateMetaStudioOrders()
	return &CreateMetaStudioOrdersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeletePacifyWords 批量删除安抚话术
//
// 该接口用于批量删除安抚话术。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) BatchDeletePacifyWords(request *model.BatchDeletePacifyWordsRequest) (*model.BatchDeletePacifyWordsResponse, error) {
	requestDef := GenReqDefForBatchDeletePacifyWords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeletePacifyWordsResponse), nil
	}
}

// BatchDeletePacifyWordsInvoker 批量删除安抚话术
func (c *MetaStudioClient) BatchDeletePacifyWordsInvoker(request *model.BatchDeletePacifyWordsRequest) *BatchDeletePacifyWordsInvoker {
	requestDef := GenReqDefForBatchDeletePacifyWords()
	return &BatchDeletePacifyWordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePacifyWords 创建安抚话术
//
// 该接口用于创建安抚话术。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreatePacifyWords(request *model.CreatePacifyWordsRequest) (*model.CreatePacifyWordsResponse, error) {
	requestDef := GenReqDefForCreatePacifyWords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePacifyWordsResponse), nil
	}
}

// CreatePacifyWordsInvoker 创建安抚话术
func (c *MetaStudioClient) CreatePacifyWordsInvoker(request *model.CreatePacifyWordsRequest) *CreatePacifyWordsInvoker {
	requestDef := GenReqDefForCreatePacifyWords()
	return &CreatePacifyWordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePacifyWords 删除安抚话术
//
// 该接口用于删除安抚话术。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeletePacifyWords(request *model.DeletePacifyWordsRequest) (*model.DeletePacifyWordsResponse, error) {
	requestDef := GenReqDefForDeletePacifyWords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePacifyWordsResponse), nil
	}
}

// DeletePacifyWordsInvoker 删除安抚话术
func (c *MetaStudioClient) DeletePacifyWordsInvoker(request *model.DeletePacifyWordsRequest) *DeletePacifyWordsInvoker {
	requestDef := GenReqDefForDeletePacifyWords()
	return &DeletePacifyWordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPacifyWords 查询安抚话术列表
//
// 该接口用于查询安抚话术列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListPacifyWords(request *model.ListPacifyWordsRequest) (*model.ListPacifyWordsResponse, error) {
	requestDef := GenReqDefForListPacifyWords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPacifyWordsResponse), nil
	}
}

// ListPacifyWordsInvoker 查询安抚话术列表
func (c *MetaStudioClient) ListPacifyWordsInvoker(request *model.ListPacifyWordsRequest) *ListPacifyWordsInvoker {
	requestDef := GenReqDefForListPacifyWords()
	return &ListPacifyWordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPacifyWords 查询安抚话术详情
//
// 该接口用于查询安抚话术详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowPacifyWords(request *model.ShowPacifyWordsRequest) (*model.ShowPacifyWordsResponse, error) {
	requestDef := GenReqDefForShowPacifyWords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPacifyWordsResponse), nil
	}
}

// ShowPacifyWordsInvoker 查询安抚话术详情
func (c *MetaStudioClient) ShowPacifyWordsInvoker(request *model.ShowPacifyWordsRequest) *ShowPacifyWordsInvoker {
	requestDef := GenReqDefForShowPacifyWords()
	return &ShowPacifyWordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPacifyWordsIntent 查询安抚话术意图
//
// 该接口用于查询安抚话术意图。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowPacifyWordsIntent(request *model.ShowPacifyWordsIntentRequest) (*model.ShowPacifyWordsIntentResponse, error) {
	requestDef := GenReqDefForShowPacifyWordsIntent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPacifyWordsIntentResponse), nil
	}
}

// ShowPacifyWordsIntentInvoker 查询安抚话术意图
func (c *MetaStudioClient) ShowPacifyWordsIntentInvoker(request *model.ShowPacifyWordsIntentRequest) *ShowPacifyWordsIntentInvoker {
	requestDef := GenReqDefForShowPacifyWordsIntent()
	return &ShowPacifyWordsIntentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPacifyWordsSwitch 查询安抚话术功能开关
//
// 该接口用于查询安抚话术功能开关。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowPacifyWordsSwitch(request *model.ShowPacifyWordsSwitchRequest) (*model.ShowPacifyWordsSwitchResponse, error) {
	requestDef := GenReqDefForShowPacifyWordsSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPacifyWordsSwitchResponse), nil
	}
}

// ShowPacifyWordsSwitchInvoker 查询安抚话术功能开关
func (c *MetaStudioClient) ShowPacifyWordsSwitchInvoker(request *model.ShowPacifyWordsSwitchRequest) *ShowPacifyWordsSwitchInvoker {
	requestDef := GenReqDefForShowPacifyWordsSwitch()
	return &ShowPacifyWordsSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPacifyWordsTriggerTime 查询安抚话术等待触发时长
//
// 该接口用于查询等待触发时长。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowPacifyWordsTriggerTime(request *model.ShowPacifyWordsTriggerTimeRequest) (*model.ShowPacifyWordsTriggerTimeResponse, error) {
	requestDef := GenReqDefForShowPacifyWordsTriggerTime()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPacifyWordsTriggerTimeResponse), nil
	}
}

// ShowPacifyWordsTriggerTimeInvoker 查询安抚话术等待触发时长
func (c *MetaStudioClient) ShowPacifyWordsTriggerTimeInvoker(request *model.ShowPacifyWordsTriggerTimeRequest) *ShowPacifyWordsTriggerTimeInvoker {
	requestDef := GenReqDefForShowPacifyWordsTriggerTime()
	return &ShowPacifyWordsTriggerTimeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePacifyWords 修改安抚话术
//
// 该接口用于修改安抚话术。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdatePacifyWords(request *model.UpdatePacifyWordsRequest) (*model.UpdatePacifyWordsResponse, error) {
	requestDef := GenReqDefForUpdatePacifyWords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePacifyWordsResponse), nil
	}
}

// UpdatePacifyWordsInvoker 修改安抚话术
func (c *MetaStudioClient) UpdatePacifyWordsInvoker(request *model.UpdatePacifyWordsRequest) *UpdatePacifyWordsInvoker {
	requestDef := GenReqDefForUpdatePacifyWords()
	return &UpdatePacifyWordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePacifyWordsSwitch 修改安抚话术功能开关
//
// 该接口用于修改安抚话术功能开关。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdatePacifyWordsSwitch(request *model.UpdatePacifyWordsSwitchRequest) (*model.UpdatePacifyWordsSwitchResponse, error) {
	requestDef := GenReqDefForUpdatePacifyWordsSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePacifyWordsSwitchResponse), nil
	}
}

// UpdatePacifyWordsSwitchInvoker 修改安抚话术功能开关
func (c *MetaStudioClient) UpdatePacifyWordsSwitchInvoker(request *model.UpdatePacifyWordsSwitchRequest) *UpdatePacifyWordsSwitchInvoker {
	requestDef := GenReqDefForUpdatePacifyWordsSwitch()
	return &UpdatePacifyWordsSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePacifyWordsTriggerTime 修改安抚话术等待触发时长
//
// 该接口用于修改安抚话术等待触发时长。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdatePacifyWordsTriggerTime(request *model.UpdatePacifyWordsTriggerTimeRequest) (*model.UpdatePacifyWordsTriggerTimeResponse, error) {
	requestDef := GenReqDefForUpdatePacifyWordsTriggerTime()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePacifyWordsTriggerTimeResponse), nil
	}
}

// UpdatePacifyWordsTriggerTimeInvoker 修改安抚话术等待触发时长
func (c *MetaStudioClient) UpdatePacifyWordsTriggerTimeInvoker(request *model.UpdatePacifyWordsTriggerTimeRequest) *UpdatePacifyWordsTriggerTimeInvoker {
	requestDef := GenReqDefForUpdatePacifyWordsTriggerTime()
	return &UpdatePacifyWordsTriggerTimeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePictureModelingByUrlJob 基于图片URL创建照片建模任务
//
// 该接口用于从URL中获取图片进行照片建模任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreatePictureModelingByUrlJob(request *model.CreatePictureModelingByUrlJobRequest) (*model.CreatePictureModelingByUrlJobResponse, error) {
	requestDef := GenReqDefForCreatePictureModelingByUrlJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePictureModelingByUrlJobResponse), nil
	}
}

// CreatePictureModelingByUrlJobInvoker 基于图片URL创建照片建模任务
func (c *MetaStudioClient) CreatePictureModelingByUrlJobInvoker(request *model.CreatePictureModelingByUrlJobRequest) *CreatePictureModelingByUrlJobInvoker {
	requestDef := GenReqDefForCreatePictureModelingByUrlJob()
	return &CreatePictureModelingByUrlJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePictureModelingJob 创建照片建模任务
//
// 该接口用于创建风格化照片建模任务。通过上传照片，生成风格化数字人模型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreatePictureModelingJob(request *model.CreatePictureModelingJobRequest) (*model.CreatePictureModelingJobResponse, error) {
	requestDef := GenReqDefForCreatePictureModelingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePictureModelingJobResponse), nil
	}
}

// CreatePictureModelingJobInvoker 创建照片建模任务
func (c *MetaStudioClient) CreatePictureModelingJobInvoker(request *model.CreatePictureModelingJobRequest) *CreatePictureModelingJobInvoker {
	requestDef := GenReqDefForCreatePictureModelingJob()
	return &CreatePictureModelingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPictureModelingJobs 照片建模任务列表查询
//
// 该接口用于查询风格化照片建模任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListPictureModelingJobs(request *model.ListPictureModelingJobsRequest) (*model.ListPictureModelingJobsResponse, error) {
	requestDef := GenReqDefForListPictureModelingJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPictureModelingJobsResponse), nil
	}
}

// ListPictureModelingJobsInvoker 照片建模任务列表查询
func (c *MetaStudioClient) ListPictureModelingJobsInvoker(request *model.ListPictureModelingJobsRequest) *ListPictureModelingJobsInvoker {
	requestDef := GenReqDefForListPictureModelingJobs()
	return &ListPictureModelingJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPictureModelingJob 照片建模任务详情查询
//
// 该接口用于风格化查询照片建模任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowPictureModelingJob(request *model.ShowPictureModelingJobRequest) (*model.ShowPictureModelingJobResponse, error) {
	requestDef := GenReqDefForShowPictureModelingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPictureModelingJobResponse), nil
	}
}

// ShowPictureModelingJobInvoker 照片建模任务详情查询
func (c *MetaStudioClient) ShowPictureModelingJobInvoker(request *model.ShowPictureModelingJobRequest) *ShowPictureModelingJobInvoker {
	requestDef := GenReqDefForShowPictureModelingJob()
	return &ShowPictureModelingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePluginConfig 创建插件配置
//
// 该接口用于创建插件配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreatePluginConfig(request *model.CreatePluginConfigRequest) (*model.CreatePluginConfigResponse, error) {
	requestDef := GenReqDefForCreatePluginConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePluginConfigResponse), nil
	}
}

// CreatePluginConfigInvoker 创建插件配置
func (c *MetaStudioClient) CreatePluginConfigInvoker(request *model.CreatePluginConfigRequest) *CreatePluginConfigInvoker {
	requestDef := GenReqDefForCreatePluginConfig()
	return &CreatePluginConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePluginConfig 删除插件配置
//
// 该接口用于删除插件配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeletePluginConfig(request *model.DeletePluginConfigRequest) (*model.DeletePluginConfigResponse, error) {
	requestDef := GenReqDefForDeletePluginConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePluginConfigResponse), nil
	}
}

// DeletePluginConfigInvoker 删除插件配置
func (c *MetaStudioClient) DeletePluginConfigInvoker(request *model.DeletePluginConfigRequest) *DeletePluginConfigInvoker {
	requestDef := GenReqDefForDeletePluginConfig()
	return &DeletePluginConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPluginConfig 查询插件配置列表
//
// 该接口用于查询插件配置列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListPluginConfig(request *model.ListPluginConfigRequest) (*model.ListPluginConfigResponse, error) {
	requestDef := GenReqDefForListPluginConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPluginConfigResponse), nil
	}
}

// ListPluginConfigInvoker 查询插件配置列表
func (c *MetaStudioClient) ListPluginConfigInvoker(request *model.ListPluginConfigRequest) *ListPluginConfigInvoker {
	requestDef := GenReqDefForListPluginConfig()
	return &ListPluginConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPluginConfig 查询插件配置详情
//
// 该接口用于查询插件配置详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowPluginConfig(request *model.ShowPluginConfigRequest) (*model.ShowPluginConfigResponse, error) {
	requestDef := GenReqDefForShowPluginConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPluginConfigResponse), nil
	}
}

// ShowPluginConfigInvoker 查询插件配置详情
func (c *MetaStudioClient) ShowPluginConfigInvoker(request *model.ShowPluginConfigRequest) *ShowPluginConfigInvoker {
	requestDef := GenReqDefForShowPluginConfig()
	return &ShowPluginConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPluginConfigDefaultInfo 查询插件配置默认信息
//
// 该接口用于查询插件配置默认信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowPluginConfigDefaultInfo(request *model.ShowPluginConfigDefaultInfoRequest) (*model.ShowPluginConfigDefaultInfoResponse, error) {
	requestDef := GenReqDefForShowPluginConfigDefaultInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPluginConfigDefaultInfoResponse), nil
	}
}

// ShowPluginConfigDefaultInfoInvoker 查询插件配置默认信息
func (c *MetaStudioClient) ShowPluginConfigDefaultInfoInvoker(request *model.ShowPluginConfigDefaultInfoRequest) *ShowPluginConfigDefaultInfoInvoker {
	requestDef := GenReqDefForShowPluginConfigDefaultInfo()
	return &ShowPluginConfigDefaultInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePluginConfig 修改插件配置
//
// 该接口用于修改插件配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdatePluginConfig(request *model.UpdatePluginConfigRequest) (*model.UpdatePluginConfigResponse, error) {
	requestDef := GenReqDefForUpdatePluginConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePluginConfigResponse), nil
	}
}

// UpdatePluginConfigInvoker 修改插件配置
func (c *MetaStudioClient) UpdatePluginConfigInvoker(request *model.UpdatePluginConfigRequest) *UpdatePluginConfigInvoker {
	requestDef := GenReqDefForUpdatePluginConfig()
	return &UpdatePluginConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateProduct 创建商品
//
// Create product
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateProduct(request *model.CreateProductRequest) (*model.CreateProductResponse, error) {
	requestDef := GenReqDefForCreateProduct()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProductResponse), nil
	}
}

// CreateProductInvoker 创建商品
func (c *MetaStudioClient) CreateProductInvoker(request *model.CreateProductRequest) *CreateProductInvoker {
	requestDef := GenReqDefForCreateProduct()
	return &CreateProductInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteProduct 删除商品
//
// 删除商品
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteProduct(request *model.DeleteProductRequest) (*model.DeleteProductResponse, error) {
	requestDef := GenReqDefForDeleteProduct()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteProductResponse), nil
	}
}

// DeleteProductInvoker 删除商品
func (c *MetaStudioClient) DeleteProductInvoker(request *model.DeleteProductRequest) *DeleteProductInvoker {
	requestDef := GenReqDefForDeleteProduct()
	return &DeleteProductInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProducts 查询商品列表
//
// 查询商品列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListProducts(request *model.ListProductsRequest) (*model.ListProductsResponse, error) {
	requestDef := GenReqDefForListProducts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProductsResponse), nil
	}
}

// ListProductsInvoker 查询商品列表
func (c *MetaStudioClient) ListProductsInvoker(request *model.ListProductsRequest) *ListProductsInvoker {
	requestDef := GenReqDefForListProducts()
	return &ListProductsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetProductAsset 商品资产组合配置
//
// 商品资产组合配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) SetProductAsset(request *model.SetProductAssetRequest) (*model.SetProductAssetResponse, error) {
	requestDef := GenReqDefForSetProductAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetProductAssetResponse), nil
	}
}

// SetProductAssetInvoker 商品资产组合配置
func (c *MetaStudioClient) SetProductAssetInvoker(request *model.SetProductAssetRequest) *SetProductAssetInvoker {
	requestDef := GenReqDefForSetProductAsset()
	return &SetProductAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProduct 查询商品详情
//
// Show product
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowProduct(request *model.ShowProductRequest) (*model.ShowProductResponse, error) {
	requestDef := GenReqDefForShowProduct()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProductResponse), nil
	}
}

// ShowProductInvoker 查询商品详情
func (c *MetaStudioClient) ShowProductInvoker(request *model.ShowProductRequest) *ShowProductInvoker {
	requestDef := GenReqDefForShowProduct()
	return &ShowProductInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProduct 更新商品
//
// Update product
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateProduct(request *model.UpdateProductRequest) (*model.UpdateProductResponse, error) {
	requestDef := GenReqDefForUpdateProduct()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProductResponse), nil
	}
}

// UpdateProductInvoker 更新商品
func (c *MetaStudioClient) UpdateProductInvoker(request *model.UpdateProductRequest) *UpdateProductInvoker {
	requestDef := GenReqDefForUpdateProduct()
	return &UpdateProductInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateQuestionAnswer 创建问答对
//
// 该接口用于创建问答对。一个问答对包含一个标准问题，一个答案，若干个相似问题等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateQuestionAnswer(request *model.CreateQuestionAnswerRequest) (*model.CreateQuestionAnswerResponse, error) {
	requestDef := GenReqDefForCreateQuestionAnswer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateQuestionAnswerResponse), nil
	}
}

// CreateQuestionAnswerInvoker 创建问答对
func (c *MetaStudioClient) CreateQuestionAnswerInvoker(request *model.CreateQuestionAnswerRequest) *CreateQuestionAnswerInvoker {
	requestDef := GenReqDefForCreateQuestionAnswer()
	return &CreateQuestionAnswerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteQuestionAnswer 删除问答对
//
// 该接口用于删除问答对。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteQuestionAnswer(request *model.DeleteQuestionAnswerRequest) (*model.DeleteQuestionAnswerResponse, error) {
	requestDef := GenReqDefForDeleteQuestionAnswer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteQuestionAnswerResponse), nil
	}
}

// DeleteQuestionAnswerInvoker 删除问答对
func (c *MetaStudioClient) DeleteQuestionAnswerInvoker(request *model.DeleteQuestionAnswerRequest) *DeleteQuestionAnswerInvoker {
	requestDef := GenReqDefForDeleteQuestionAnswer()
	return &DeleteQuestionAnswerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQuestionAnswer 查询问答对列表
//
// 该接口用于查询问答对列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListQuestionAnswer(request *model.ListQuestionAnswerRequest) (*model.ListQuestionAnswerResponse, error) {
	requestDef := GenReqDefForListQuestionAnswer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuestionAnswerResponse), nil
	}
}

// ListQuestionAnswerInvoker 查询问答对列表
func (c *MetaStudioClient) ListQuestionAnswerInvoker(request *model.ListQuestionAnswerRequest) *ListQuestionAnswerInvoker {
	requestDef := GenReqDefForListQuestionAnswer()
	return &ListQuestionAnswerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQuestionAnswer 查询问答对详情
//
// 该接口用于查询问答对详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowQuestionAnswer(request *model.ShowQuestionAnswerRequest) (*model.ShowQuestionAnswerResponse, error) {
	requestDef := GenReqDefForShowQuestionAnswer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuestionAnswerResponse), nil
	}
}

// ShowQuestionAnswerInvoker 查询问答对详情
func (c *MetaStudioClient) ShowQuestionAnswerInvoker(request *model.ShowQuestionAnswerRequest) *ShowQuestionAnswerInvoker {
	requestDef := GenReqDefForShowQuestionAnswer()
	return &ShowQuestionAnswerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateQuestionAnswer 修改问答对
//
// 该接口用于修改问答对。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateQuestionAnswer(request *model.UpdateQuestionAnswerRequest) (*model.UpdateQuestionAnswerResponse, error) {
	requestDef := GenReqDefForUpdateQuestionAnswer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateQuestionAnswerResponse), nil
	}
}

// UpdateQuestionAnswerInvoker 修改问答对
func (c *MetaStudioClient) UpdateQuestionAnswerInvoker(request *model.UpdateQuestionAnswerRequest) *UpdateQuestionAnswerInvoker {
	requestDef := GenReqDefForUpdateQuestionAnswer()
	return &UpdateQuestionAnswerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVoiceTrainingQuotas 查询声音训练资源
//
// 查询声音训练资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowVoiceTrainingQuotas(request *model.ShowVoiceTrainingQuotasRequest) (*model.ShowVoiceTrainingQuotasResponse, error) {
	requestDef := GenReqDefForShowVoiceTrainingQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVoiceTrainingQuotasResponse), nil
	}
}

// ShowVoiceTrainingQuotasInvoker 查询声音训练资源
func (c *MetaStudioClient) ShowVoiceTrainingQuotasInvoker(request *model.ShowVoiceTrainingQuotasRequest) *ShowVoiceTrainingQuotasInvoker {
	requestDef := GenReqDefForShowVoiceTrainingQuotas()
	return &ShowVoiceTrainingQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRobot 创建应用
//
// 该接口用于创建应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateRobot(request *model.CreateRobotRequest) (*model.CreateRobotResponse, error) {
	requestDef := GenReqDefForCreateRobot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRobotResponse), nil
	}
}

// CreateRobotInvoker 创建应用
func (c *MetaStudioClient) CreateRobotInvoker(request *model.CreateRobotRequest) *CreateRobotInvoker {
	requestDef := GenReqDefForCreateRobot()
	return &CreateRobotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRobot 删除应用
//
// 该接口用于删除应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteRobot(request *model.DeleteRobotRequest) (*model.DeleteRobotResponse, error) {
	requestDef := GenReqDefForDeleteRobot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRobotResponse), nil
	}
}

// DeleteRobotInvoker 删除应用
func (c *MetaStudioClient) DeleteRobotInvoker(request *model.DeleteRobotRequest) *DeleteRobotInvoker {
	requestDef := GenReqDefForDeleteRobot()
	return &DeleteRobotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRobot 查询应用列表
//
// 该接口用于查询应用列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListRobot(request *model.ListRobotRequest) (*model.ListRobotResponse, error) {
	requestDef := GenReqDefForListRobot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRobotResponse), nil
	}
}

// ListRobotInvoker 查询应用列表
func (c *MetaStudioClient) ListRobotInvoker(request *model.ListRobotRequest) *ListRobotInvoker {
	requestDef := GenReqDefForListRobot()
	return &ListRobotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRobot 查询应用详情
//
// 该接口用于查询应用详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowRobot(request *model.ShowRobotRequest) (*model.ShowRobotResponse, error) {
	requestDef := GenReqDefForShowRobot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRobotResponse), nil
	}
}

// ShowRobotInvoker 查询应用详情
func (c *MetaStudioClient) ShowRobotInvoker(request *model.ShowRobotRequest) *ShowRobotInvoker {
	requestDef := GenReqDefForShowRobot()
	return &ShowRobotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRobot 修改应用
//
// 该接口用于修改应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateRobot(request *model.UpdateRobotRequest) (*model.UpdateRobotResponse, error) {
	requestDef := GenReqDefForUpdateRobot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRobotResponse), nil
	}
}

// UpdateRobotInvoker 修改应用
func (c *MetaStudioClient) UpdateRobotInvoker(request *model.UpdateRobotRequest) *UpdateRobotInvoker {
	requestDef := GenReqDefForUpdateRobot()
	return &UpdateRobotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ValidateRobot 校验应用
//
// 该接口用于校验应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ValidateRobot(request *model.ValidateRobotRequest) (*model.ValidateRobotResponse, error) {
	requestDef := GenReqDefForValidateRobot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ValidateRobotResponse), nil
	}
}

// ValidateRobotInvoker 校验应用
func (c *MetaStudioClient) ValidateRobotInvoker(request *model.ValidateRobotRequest) *ValidateRobotInvoker {
	requestDef := GenReqDefForValidateRobot()
	return &ValidateRobotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRole 创建角色
//
// 该接口用于创建角色。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateRole(request *model.CreateRoleRequest) (*model.CreateRoleResponse, error) {
	requestDef := GenReqDefForCreateRole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRoleResponse), nil
	}
}

// CreateRoleInvoker 创建角色
func (c *MetaStudioClient) CreateRoleInvoker(request *model.CreateRoleRequest) *CreateRoleInvoker {
	requestDef := GenReqDefForCreateRole()
	return &CreateRoleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRole 删除角色
//
// 该接口用于删除角色。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteRole(request *model.DeleteRoleRequest) (*model.DeleteRoleResponse, error) {
	requestDef := GenReqDefForDeleteRole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRoleResponse), nil
	}
}

// DeleteRoleInvoker 删除角色
func (c *MetaStudioClient) DeleteRoleInvoker(request *model.DeleteRoleRequest) *DeleteRoleInvoker {
	requestDef := GenReqDefForDeleteRole()
	return &DeleteRoleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRole 查询角色列表
//
// 该接口用于查询角色列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListRole(request *model.ListRoleRequest) (*model.ListRoleResponse, error) {
	requestDef := GenReqDefForListRole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRoleResponse), nil
	}
}

// ListRoleInvoker 查询角色列表
func (c *MetaStudioClient) ListRoleInvoker(request *model.ListRoleRequest) *ListRoleInvoker {
	requestDef := GenReqDefForListRole()
	return &ListRoleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRole 查询角色详情
//
// 该接口用于查询角色详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowRole(request *model.ShowRoleRequest) (*model.ShowRoleResponse, error) {
	requestDef := GenReqDefForShowRole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRoleResponse), nil
	}
}

// ShowRoleInvoker 查询角色详情
func (c *MetaStudioClient) ShowRoleInvoker(request *model.ShowRoleRequest) *ShowRoleInvoker {
	requestDef := GenReqDefForShowRole()
	return &ShowRoleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRole 修改角色
//
// 该接口用于修改角色。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateRole(request *model.UpdateRoleRequest) (*model.UpdateRoleResponse, error) {
	requestDef := GenReqDefForUpdateRole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRoleResponse), nil
	}
}

// UpdateRoleInvoker 修改角色
func (c *MetaStudioClient) UpdateRoleInvoker(request *model.UpdateRoleRequest) *UpdateRoleInvoker {
	requestDef := GenReqDefForUpdateRole()
	return &UpdateRoleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSmartChatRoom 创建智能交互对话
//
// 该接口用于创建智能交互对话。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateSmartChatRoom(request *model.CreateSmartChatRoomRequest) (*model.CreateSmartChatRoomResponse, error) {
	requestDef := GenReqDefForCreateSmartChatRoom()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSmartChatRoomResponse), nil
	}
}

// CreateSmartChatRoomInvoker 创建智能交互对话
func (c *MetaStudioClient) CreateSmartChatRoomInvoker(request *model.CreateSmartChatRoomRequest) *CreateSmartChatRoomInvoker {
	requestDef := GenReqDefForCreateSmartChatRoom()
	return &CreateSmartChatRoomInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSmartChatRoom 删除智能交互对话
//
// 该接口用于删除智能交互对话。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteSmartChatRoom(request *model.DeleteSmartChatRoomRequest) (*model.DeleteSmartChatRoomResponse, error) {
	requestDef := GenReqDefForDeleteSmartChatRoom()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSmartChatRoomResponse), nil
	}
}

// DeleteSmartChatRoomInvoker 删除智能交互对话
func (c *MetaStudioClient) DeleteSmartChatRoomInvoker(request *model.DeleteSmartChatRoomRequest) *DeleteSmartChatRoomInvoker {
	requestDef := GenReqDefForDeleteSmartChatRoom()
	return &DeleteSmartChatRoomInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSmartChatRooms 查询智能交互对话列表
//
// 该接口用于智能交互对话列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListSmartChatRooms(request *model.ListSmartChatRoomsRequest) (*model.ListSmartChatRoomsResponse, error) {
	requestDef := GenReqDefForListSmartChatRooms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSmartChatRoomsResponse), nil
	}
}

// ListSmartChatRoomsInvoker 查询智能交互对话列表
func (c *MetaStudioClient) ListSmartChatRoomsInvoker(request *model.ListSmartChatRoomsRequest) *ListSmartChatRoomsInvoker {
	requestDef := GenReqDefForListSmartChatRooms()
	return &ListSmartChatRoomsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSmartChatRoom 查询智能交互对话详情
//
// 该接口用于查询智能交互对话详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowSmartChatRoom(request *model.ShowSmartChatRoomRequest) (*model.ShowSmartChatRoomResponse, error) {
	requestDef := GenReqDefForShowSmartChatRoom()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSmartChatRoomResponse), nil
	}
}

// ShowSmartChatRoomInvoker 查询智能交互对话详情
func (c *MetaStudioClient) ShowSmartChatRoomInvoker(request *model.ShowSmartChatRoomRequest) *ShowSmartChatRoomInvoker {
	requestDef := GenReqDefForShowSmartChatRoom()
	return &ShowSmartChatRoomInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSmartChatRoom 更新智能交互对话信息
//
// 该接口用于智能交互对话信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateSmartChatRoom(request *model.UpdateSmartChatRoomRequest) (*model.UpdateSmartChatRoomResponse, error) {
	requestDef := GenReqDefForUpdateSmartChatRoom()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSmartChatRoomResponse), nil
	}
}

// UpdateSmartChatRoomInvoker 更新智能交互对话信息
func (c *MetaStudioClient) UpdateSmartChatRoomInvoker(request *model.UpdateSmartChatRoomRequest) *UpdateSmartChatRoomInvoker {
	requestDef := GenReqDefForUpdateSmartChatRoom()
	return &UpdateSmartChatRoomInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSmartLiveUserConfig 租户查询直播租户级配置
//
// 该接口用于租户设置直播租户级配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowSmartLiveUserConfig(request *model.ShowSmartLiveUserConfigRequest) (*model.ShowSmartLiveUserConfigResponse, error) {
	requestDef := GenReqDefForShowSmartLiveUserConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSmartLiveUserConfigResponse), nil
	}
}

// ShowSmartLiveUserConfigInvoker 租户查询直播租户级配置
func (c *MetaStudioClient) ShowSmartLiveUserConfigInvoker(request *model.ShowSmartLiveUserConfigRequest) *ShowSmartLiveUserConfigInvoker {
	requestDef := GenReqDefForShowSmartLiveUserConfig()
	return &ShowSmartLiveUserConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSmartLiveUserConfig 租户设置直播租户级配置
//
// 该接口用于租户设置直播租户级配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateSmartLiveUserConfig(request *model.UpdateSmartLiveUserConfigRequest) (*model.UpdateSmartLiveUserConfigResponse, error) {
	requestDef := GenReqDefForUpdateSmartLiveUserConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSmartLiveUserConfigResponse), nil
	}
}

// UpdateSmartLiveUserConfigInvoker 租户设置直播租户级配置
func (c *MetaStudioClient) UpdateSmartLiveUserConfigInvoker(request *model.UpdateSmartLiveUserConfigRequest) *UpdateSmartLiveUserConfigInvoker {
	requestDef := GenReqDefForUpdateSmartLiveUserConfig()
	return &UpdateSmartLiveUserConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchConfirmLiveCommands 批量确认命令
//
// 该接口用于批量确认命令列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) BatchConfirmLiveCommands(request *model.BatchConfirmLiveCommandsRequest) (*model.BatchConfirmLiveCommandsResponse, error) {
	requestDef := GenReqDefForBatchConfirmLiveCommands()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchConfirmLiveCommandsResponse), nil
	}
}

// BatchConfirmLiveCommandsInvoker 批量确认命令
func (c *MetaStudioClient) BatchConfirmLiveCommandsInvoker(request *model.BatchConfirmLiveCommandsRequest) *BatchConfirmLiveCommandsInvoker {
	requestDef := GenReqDefForBatchConfirmLiveCommands()
	return &BatchConfirmLiveCommandsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteSmartLiveCommand 控制数字人直播过程
//
// 该接口用于控制数字人直播过程。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ExecuteSmartLiveCommand(request *model.ExecuteSmartLiveCommandRequest) (*model.ExecuteSmartLiveCommandResponse, error) {
	requestDef := GenReqDefForExecuteSmartLiveCommand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteSmartLiveCommandResponse), nil
	}
}

// ExecuteSmartLiveCommandInvoker 控制数字人直播过程
func (c *MetaStudioClient) ExecuteSmartLiveCommandInvoker(request *model.ExecuteSmartLiveCommandRequest) *ExecuteSmartLiveCommandInvoker {
	requestDef := GenReqDefForExecuteSmartLiveCommand()
	return &ExecuteSmartLiveCommandInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInsertCommands 查询数字人直播插入命令列表
//
// 该接口用于查询数字人直播插入命令列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListInsertCommands(request *model.ListInsertCommandsRequest) (*model.ListInsertCommandsResponse, error) {
	requestDef := GenReqDefForListInsertCommands()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInsertCommandsResponse), nil
	}
}

// ListInsertCommandsInvoker 查询数字人直播插入命令列表
func (c *MetaStudioClient) ListInsertCommandsInvoker(request *model.ListInsertCommandsRequest) *ListInsertCommandsInvoker {
	requestDef := GenReqDefForListInsertCommands()
	return &ListInsertCommandsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSmartLive 查询某个智能直播间下直播任务列表
//
// 该接口用于查询某个智能直播间的直播任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListSmartLive(request *model.ListSmartLiveRequest) (*model.ListSmartLiveResponse, error) {
	requestDef := GenReqDefForListSmartLive()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSmartLiveResponse), nil
	}
}

// ListSmartLiveInvoker 查询某个智能直播间下直播任务列表
func (c *MetaStudioClient) ListSmartLiveInvoker(request *model.ListSmartLiveRequest) *ListSmartLiveInvoker {
	requestDef := GenReqDefForListSmartLive()
	return &ListSmartLiveInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSmartLiveJobs 查询租户所有数字人直播任务列表
//
// 该接口用于查询租户所有数字人智能直播任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListSmartLiveJobs(request *model.ListSmartLiveJobsRequest) (*model.ListSmartLiveJobsResponse, error) {
	requestDef := GenReqDefForListSmartLiveJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSmartLiveJobsResponse), nil
	}
}

// ListSmartLiveJobsInvoker 查询租户所有数字人直播任务列表
func (c *MetaStudioClient) ListSmartLiveJobsInvoker(request *model.ListSmartLiveJobsRequest) *ListSmartLiveJobsInvoker {
	requestDef := GenReqDefForListSmartLiveJobs()
	return &ListSmartLiveJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSmartLiveRuleCommands 查询租户未确认的互动规则命令列表
//
// 该接口用于查询租户未确认的互动规则命令列表，仅限于需要做二次确认的特定用户使用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListSmartLiveRuleCommands(request *model.ListSmartLiveRuleCommandsRequest) (*model.ListSmartLiveRuleCommandsResponse, error) {
	requestDef := GenReqDefForListSmartLiveRuleCommands()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSmartLiveRuleCommandsResponse), nil
	}
}

// ListSmartLiveRuleCommandsInvoker 查询租户未确认的互动规则命令列表
func (c *MetaStudioClient) ListSmartLiveRuleCommandsInvoker(request *model.ListSmartLiveRuleCommandsRequest) *ListSmartLiveRuleCommandsInvoker {
	requestDef := GenReqDefForListSmartLiveRuleCommands()
	return &ListSmartLiveRuleCommandsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSmartLiveScriptCommands 查询租户未确认的剧本命令列表
//
// 该接口用于查询租户未确认的剧本命令列表，仅限于需要做二次确认的特定用户使用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListSmartLiveScriptCommands(request *model.ListSmartLiveScriptCommandsRequest) (*model.ListSmartLiveScriptCommandsResponse, error) {
	requestDef := GenReqDefForListSmartLiveScriptCommands()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSmartLiveScriptCommandsResponse), nil
	}
}

// ListSmartLiveScriptCommandsInvoker 查询租户未确认的剧本命令列表
func (c *MetaStudioClient) ListSmartLiveScriptCommandsInvoker(request *model.ListSmartLiveScriptCommandsRequest) *ListSmartLiveScriptCommandsInvoker {
	requestDef := GenReqDefForListSmartLiveScriptCommands()
	return &ListSmartLiveScriptCommandsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// LiveEventReport 上报直播间事件
//
// 该接口用于上报直播间事件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) LiveEventReport(request *model.LiveEventReportRequest) (*model.LiveEventReportResponse, error) {
	requestDef := GenReqDefForLiveEventReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.LiveEventReportResponse), nil
	}
}

// LiveEventReportInvoker 上报直播间事件
func (c *MetaStudioClient) LiveEventReportInvoker(request *model.LiveEventReportRequest) *LiveEventReportInvoker {
	requestDef := GenReqDefForLiveEventReport()
	return &LiveEventReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSmartLive 查询数字人智能直播任务详情
//
// 该接口用于查询数字人智能直播任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowSmartLive(request *model.ShowSmartLiveRequest) (*model.ShowSmartLiveResponse, error) {
	requestDef := GenReqDefForShowSmartLive()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSmartLiveResponse), nil
	}
}

// ShowSmartLiveInvoker 查询数字人智能直播任务详情
func (c *MetaStudioClient) ShowSmartLiveInvoker(request *model.ShowSmartLiveRequest) *ShowSmartLiveInvoker {
	requestDef := GenReqDefForShowSmartLive()
	return &ShowSmartLiveInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartSmartLive 启动数字人智能直播任务
//
// 该接口用于启动数字人智能直播任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) StartSmartLive(request *model.StartSmartLiveRequest) (*model.StartSmartLiveResponse, error) {
	requestDef := GenReqDefForStartSmartLive()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartSmartLiveResponse), nil
	}
}

// StartSmartLiveInvoker 启动数字人智能直播任务
func (c *MetaStudioClient) StartSmartLiveInvoker(request *model.StartSmartLiveRequest) *StartSmartLiveInvoker {
	requestDef := GenReqDefForStartSmartLive()
	return &StartSmartLiveInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopSmartLive 结束数字人智能直播任务
//
// 该接口用于结束数字人智能直播任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) StopSmartLive(request *model.StopSmartLiveRequest) (*model.StopSmartLiveResponse, error) {
	requestDef := GenReqDefForStopSmartLive()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopSmartLiveResponse), nil
	}
}

// StopSmartLiveInvoker 结束数字人智能直播任务
func (c *MetaStudioClient) StopSmartLiveInvoker(request *model.StopSmartLiveRequest) *StopSmartLiveInvoker {
	requestDef := GenReqDefForStopSmartLive()
	return &StopSmartLiveInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ConfirmSmarLiveRoom 直播间确认
//
// 该接口用直播间二次确认
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ConfirmSmarLiveRoom(request *model.ConfirmSmarLiveRoomRequest) (*model.ConfirmSmarLiveRoomResponse, error) {
	requestDef := GenReqDefForConfirmSmarLiveRoom()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ConfirmSmarLiveRoomResponse), nil
	}
}

// ConfirmSmarLiveRoomInvoker 直播间确认
func (c *MetaStudioClient) ConfirmSmarLiveRoomInvoker(request *model.ConfirmSmarLiveRoomRequest) *ConfirmSmarLiveRoomInvoker {
	requestDef := GenReqDefForConfirmSmarLiveRoom()
	return &ConfirmSmarLiveRoomInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSmartLiveRoom 创建智能直播间
//
// 该接口用于创建智能直播间。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateSmartLiveRoom(request *model.CreateSmartLiveRoomRequest) (*model.CreateSmartLiveRoomResponse, error) {
	requestDef := GenReqDefForCreateSmartLiveRoom()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSmartLiveRoomResponse), nil
	}
}

// CreateSmartLiveRoomInvoker 创建智能直播间
func (c *MetaStudioClient) CreateSmartLiveRoomInvoker(request *model.CreateSmartLiveRoomRequest) *CreateSmartLiveRoomInvoker {
	requestDef := GenReqDefForCreateSmartLiveRoom()
	return &CreateSmartLiveRoomInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSmartLiveRoom 删除智能直播间
//
// 该接口用于删除智能直播间。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteSmartLiveRoom(request *model.DeleteSmartLiveRoomRequest) (*model.DeleteSmartLiveRoomResponse, error) {
	requestDef := GenReqDefForDeleteSmartLiveRoom()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSmartLiveRoomResponse), nil
	}
}

// DeleteSmartLiveRoomInvoker 删除智能直播间
func (c *MetaStudioClient) DeleteSmartLiveRoomInvoker(request *model.DeleteSmartLiveRoomRequest) *DeleteSmartLiveRoomInvoker {
	requestDef := GenReqDefForDeleteSmartLiveRoom()
	return &DeleteSmartLiveRoomInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSmartLiveRooms 查询智能直播间列表
//
// 该接口用于智能直播间列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListSmartLiveRooms(request *model.ListSmartLiveRoomsRequest) (*model.ListSmartLiveRoomsResponse, error) {
	requestDef := GenReqDefForListSmartLiveRooms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSmartLiveRoomsResponse), nil
	}
}

// ListSmartLiveRoomsInvoker 查询智能直播间列表
func (c *MetaStudioClient) ListSmartLiveRoomsInvoker(request *model.ListSmartLiveRoomsRequest) *ListSmartLiveRoomsInvoker {
	requestDef := GenReqDefForListSmartLiveRooms()
	return &ListSmartLiveRoomsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLiveWarningInfo 查询直播建配置风险信息
//
// 该接口用查询直播建配置风险信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowLiveWarningInfo(request *model.ShowLiveWarningInfoRequest) (*model.ShowLiveWarningInfoResponse, error) {
	requestDef := GenReqDefForShowLiveWarningInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLiveWarningInfoResponse), nil
	}
}

// ShowLiveWarningInfoInvoker 查询直播建配置风险信息
func (c *MetaStudioClient) ShowLiveWarningInfoInvoker(request *model.ShowLiveWarningInfoRequest) *ShowLiveWarningInfoInvoker {
	requestDef := GenReqDefForShowLiveWarningInfo()
	return &ShowLiveWarningInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSmartLiveRoom 查询智能直播间剧本详情
//
// 该接口用于查询智能直播间剧本详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowSmartLiveRoom(request *model.ShowSmartLiveRoomRequest) (*model.ShowSmartLiveRoomResponse, error) {
	requestDef := GenReqDefForShowSmartLiveRoom()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSmartLiveRoomResponse), nil
	}
}

// ShowSmartLiveRoomInvoker 查询智能直播间剧本详情
func (c *MetaStudioClient) ShowSmartLiveRoomInvoker(request *model.ShowSmartLiveRoomRequest) *ShowSmartLiveRoomInvoker {
	requestDef := GenReqDefForShowSmartLiveRoom()
	return &ShowSmartLiveRoomInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSmartLiveRoom 更新智能直播间信息
//
// 该接口用于智能直播间信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateSmartLiveRoom(request *model.UpdateSmartLiveRoomRequest) (*model.UpdateSmartLiveRoomResponse, error) {
	requestDef := GenReqDefForUpdateSmartLiveRoom()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSmartLiveRoomResponse), nil
	}
}

// UpdateSmartLiveRoomInvoker 更新智能直播间信息
func (c *MetaStudioClient) UpdateSmartLiveRoomInvoker(request *model.UpdateSmartLiveRoomRequest) *UpdateSmartLiveRoomInvoker {
	requestDef := GenReqDefForUpdateSmartLiveRoom()
	return &UpdateSmartLiveRoomInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStyles 查询数字人风格列表
//
// 查询数字人风格列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListStyles(request *model.ListStylesRequest) (*model.ListStylesResponse, error) {
	requestDef := GenReqDefForListStyles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStylesResponse), nil
	}
}

// ListStylesInvoker 查询数字人风格列表
func (c *MetaStudioClient) ListStylesInvoker(request *model.ListStylesRequest) *ListStylesInvoker {
	requestDef := GenReqDefForListStyles()
	return &ListStylesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSubtitleFile 创建分身数字人视频字幕文件
//
// 该接口用于创建分身数字人视频字幕文件任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateSubtitleFile(request *model.CreateSubtitleFileRequest) (*model.CreateSubtitleFileResponse, error) {
	requestDef := GenReqDefForCreateSubtitleFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSubtitleFileResponse), nil
	}
}

// CreateSubtitleFileInvoker 创建分身数字人视频字幕文件
func (c *MetaStudioClient) CreateSubtitleFileInvoker(request *model.CreateSubtitleFileRequest) *CreateSubtitleFileInvoker {
	requestDef := GenReqDefForCreateSubtitleFile()
	return &CreateSubtitleFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSubtitleFile 查询分身数字人视频字幕文件任务详情
//
// 该接口用于查询分身数字人视频字幕文件任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowSubtitleFile(request *model.ShowSubtitleFileRequest) (*model.ShowSubtitleFileResponse, error) {
	requestDef := GenReqDefForShowSubtitleFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSubtitleFileResponse), nil
	}
}

// ShowSubtitleFileInvoker 查询分身数字人视频字幕文件任务详情
func (c *MetaStudioClient) ShowSubtitleFileInvoker(request *model.ShowSubtitleFileRequest) *ShowSubtitleFileInvoker {
	requestDef := GenReqDefForShowSubtitleFile()
	return &ShowSubtitleFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BindUserAssetResource 资源绑定接口
//
// 资源绑定接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) BindUserAssetResource(request *model.BindUserAssetResourceRequest) (*model.BindUserAssetResourceResponse, error) {
	requestDef := GenReqDefForBindUserAssetResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BindUserAssetResourceResponse), nil
	}
}

// BindUserAssetResourceInvoker 资源绑定接口
func (c *MetaStudioClient) BindUserAssetResourceInvoker(request *model.BindUserAssetResourceRequest) *BindUserAssetResourceInvoker {
	requestDef := GenReqDefForBindUserAssetResource()
	return &BindUserAssetResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountTenantResources 统计时间段内过期的资源数量
//
// 统计指定时间段内即将过期的包周期与一次性资源数量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CountTenantResources(request *model.CountTenantResourcesRequest) (*model.CountTenantResourcesResponse, error) {
	requestDef := GenReqDefForCountTenantResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountTenantResourcesResponse), nil
	}
}

// CountTenantResourcesInvoker 统计时间段内过期的资源数量
func (c *MetaStudioClient) CountTenantResourcesInvoker(request *model.CountTenantResourcesRequest) *CountTenantResourcesInvoker {
	requestDef := GenReqDefForCountTenantResources()
	return &CountTenantResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateUserQuotas 创建子账户配额
//
// 创建子账户（IAM用户）配额，需要先开启子账户隔离后才能配置。 只有根账户可创建。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateUserQuotas(request *model.CreateUserQuotasRequest) (*model.CreateUserQuotasResponse, error) {
	requestDef := GenReqDefForCreateUserQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateUserQuotasResponse), nil
	}
}

// CreateUserQuotasInvoker 创建子账户配额
func (c *MetaStudioClient) CreateUserQuotasInvoker(request *model.CreateUserQuotasRequest) *CreateUserQuotasInvoker {
	requestDef := GenReqDefForCreateUserQuotas()
	return &CreateUserQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTenantUserConfiguration 删除租户个性化配置
//
// 删除租户个性化配置。由租户下用户操作设置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteTenantUserConfiguration(request *model.DeleteTenantUserConfigurationRequest) (*model.DeleteTenantUserConfigurationResponse, error) {
	requestDef := GenReqDefForDeleteTenantUserConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTenantUserConfigurationResponse), nil
	}
}

// DeleteTenantUserConfigurationInvoker 删除租户个性化配置
func (c *MetaStudioClient) DeleteTenantUserConfigurationInvoker(request *model.DeleteTenantUserConfigurationRequest) *DeleteTenantUserConfigurationInvoker {
	requestDef := GenReqDefForDeleteTenantUserConfiguration()
	return &DeleteTenantUserConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteUserQuotas 删除子账户配额
//
// 删除子账户（IAM用户）配额，需要先开启子账户隔离后才能配置。 只有根账户可删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteUserQuotas(request *model.DeleteUserQuotasRequest) (*model.DeleteUserQuotasResponse, error) {
	requestDef := GenReqDefForDeleteUserQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteUserQuotasResponse), nil
	}
}

// DeleteUserQuotasInvoker 删除子账户配额
func (c *MetaStudioClient) DeleteUserQuotasInvoker(request *model.DeleteUserQuotasRequest) *DeleteUserQuotasInvoker {
	requestDef := GenReqDefForDeleteUserQuotas()
	return &DeleteUserQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTenantResources 查看租户资源列表
//
// 查看租户资源列表。
//  &gt; 按需套餐包用量本接口无法查询，需要调用CBC接口查询，详见[按需套餐包用量查询](https://cbc.huaweicloud.com/bm/support/api-apidt/CBCInterface_0001239.html)和[查询资源包信息](https://cbc.huaweicloud.com/bm/support/api-apidt/CBCInterface_0000511.html)。
// &gt; 各种资源的计费方式请参考[计费说明](https://support.huaweicloud.com/productdesc-metastudio/metastudio_01_0006.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListTenantResources(request *model.ListTenantResourcesRequest) (*model.ListTenantResourcesResponse, error) {
	requestDef := GenReqDefForListTenantResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTenantResourcesResponse), nil
	}
}

// ListTenantResourcesInvoker 查看租户资源列表
func (c *MetaStudioClient) ListTenantResourcesInvoker(request *model.ListTenantResourcesRequest) *ListTenantResourcesInvoker {
	requestDef := GenReqDefForListTenantResources()
	return &ListTenantResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserQuotas 查询子账户配额
//
// 查询子账户（IAM用户）配额。 只有根账户可查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListUserQuotas(request *model.ListUserQuotasRequest) (*model.ListUserQuotasResponse, error) {
	requestDef := GenReqDefForListUserQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserQuotasResponse), nil
	}
}

// ListUserQuotasInvoker 查询子账户配额
func (c *MetaStudioClient) ListUserQuotasInvoker(request *model.ListUserQuotasRequest) *ListUserQuotasInvoker {
	requestDef := GenReqDefForListUserQuotas()
	return &ListUserQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetTenantNoticeConfiguration 设置租户个性化通知配置
//
// 设置租户个性化通知配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) SetTenantNoticeConfiguration(request *model.SetTenantNoticeConfigurationRequest) (*model.SetTenantNoticeConfigurationResponse, error) {
	requestDef := GenReqDefForSetTenantNoticeConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetTenantNoticeConfigurationResponse), nil
	}
}

// SetTenantNoticeConfigurationInvoker 设置租户个性化通知配置
func (c *MetaStudioClient) SetTenantNoticeConfigurationInvoker(request *model.SetTenantNoticeConfigurationRequest) *SetTenantNoticeConfigurationInvoker {
	requestDef := GenReqDefForSetTenantNoticeConfiguration()
	return &SetTenantNoticeConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetTenantUserConfiguration 设置租户个性化配置
//
// 设置租户个性化配置。由租户下用户操作设置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) SetTenantUserConfiguration(request *model.SetTenantUserConfigurationRequest) (*model.SetTenantUserConfigurationResponse, error) {
	requestDef := GenReqDefForSetTenantUserConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetTenantUserConfigurationResponse), nil
	}
}

// SetTenantUserConfigurationInvoker 设置租户个性化配置
func (c *MetaStudioClient) SetTenantUserConfigurationInvoker(request *model.SetTenantUserConfigurationRequest) *SetTenantUserConfigurationInvoker {
	requestDef := GenReqDefForSetTenantUserConfiguration()
	return &SetTenantUserConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceUsage 查看租户资源用量信息
//
// 查询租户一次性和包周期（包年/包月）资源用量信息。
// &gt; 按需套餐包用量本接口无法查询，需要调用CBC接口查询，详见[按需套餐包用量查询](https://cbc.huaweicloud.com/bm/support/api-apidt/CBCInterface_0001239.html)和[查询资源包信息](https://cbc.huaweicloud.com/bm/support/api-apidt/CBCInterface_0000511.html)。
// &gt; 各种资源的计费方式请参考[计费说明](https://support.huaweicloud.com/productdesc-metastudio/metastudio_01_0006.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowResourceUsage(request *model.ShowResourceUsageRequest) (*model.ShowResourceUsageResponse, error) {
	requestDef := GenReqDefForShowResourceUsage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceUsageResponse), nil
	}
}

// ShowResourceUsageInvoker 查看租户资源用量信息
func (c *MetaStudioClient) ShowResourceUsageInvoker(request *model.ShowResourceUsageRequest) *ShowResourceUsageInvoker {
	requestDef := GenReqDefForShowResourceUsage()
	return &ShowResourceUsageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTenantAssginRecord 查询租户下分配的资源详情
//
// 该接口用于普通租户查询租户下的资源详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowTenantAssginRecord(request *model.ShowTenantAssginRecordRequest) (*model.ShowTenantAssginRecordResponse, error) {
	requestDef := GenReqDefForShowTenantAssginRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTenantAssginRecordResponse), nil
	}
}

// ShowTenantAssginRecordInvoker 查询租户下分配的资源详情
func (c *MetaStudioClient) ShowTenantAssginRecordInvoker(request *model.ShowTenantAssginRecordRequest) *ShowTenantAssginRecordInvoker {
	requestDef := GenReqDefForShowTenantAssginRecord()
	return &ShowTenantAssginRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTenantNoticeConfiguration 查询租户个性化通知配置
//
// 查询租户个性化通知配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowTenantNoticeConfiguration(request *model.ShowTenantNoticeConfigurationRequest) (*model.ShowTenantNoticeConfigurationResponse, error) {
	requestDef := GenReqDefForShowTenantNoticeConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTenantNoticeConfigurationResponse), nil
	}
}

// ShowTenantNoticeConfigurationInvoker 查询租户个性化通知配置
func (c *MetaStudioClient) ShowTenantNoticeConfigurationInvoker(request *model.ShowTenantNoticeConfigurationRequest) *ShowTenantNoticeConfigurationInvoker {
	requestDef := GenReqDefForShowTenantNoticeConfiguration()
	return &ShowTenantNoticeConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTenantServiceConfigs 查看租户服务业务配置
//
// 查看租户服务业务配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowTenantServiceConfigs(request *model.ShowTenantServiceConfigsRequest) (*model.ShowTenantServiceConfigsResponse, error) {
	requestDef := GenReqDefForShowTenantServiceConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTenantServiceConfigsResponse), nil
	}
}

// ShowTenantServiceConfigsInvoker 查看租户服务业务配置
func (c *MetaStudioClient) ShowTenantServiceConfigsInvoker(request *model.ShowTenantServiceConfigsRequest) *ShowTenantServiceConfigsInvoker {
	requestDef := GenReqDefForShowTenantServiceConfigs()
	return &ShowTenantServiceConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTenantUserConfiguration 查询租户个性化配置
//
// 查询租户个性化配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowTenantUserConfiguration(request *model.ShowTenantUserConfigurationRequest) (*model.ShowTenantUserConfigurationResponse, error) {
	requestDef := GenReqDefForShowTenantUserConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTenantUserConfigurationResponse), nil
	}
}

// ShowTenantUserConfigurationInvoker 查询租户个性化配置
func (c *MetaStudioClient) ShowTenantUserConfigurationInvoker(request *model.ShowTenantUserConfigurationRequest) *ShowTenantUserConfigurationInvoker {
	requestDef := GenReqDefForShowTenantUserConfiguration()
	return &ShowTenantUserConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SignAgreement 签署云服务声明
//
// 签署云服务声明。调用此接口前请知悉[[metastudio隐私协议](https://www.huaweicloud.com/declaration/tsa_metastudio.html)](tag:hws)[[metastudio隐私协议](https://www.huaweicloud.com/intl/en-us/declaration-sg/tsa-metastudio.html)](tag:hws_hk)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) SignAgreement(request *model.SignAgreementRequest) (*model.SignAgreementResponse, error) {
	requestDef := GenReqDefForSignAgreement()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SignAgreementResponse), nil
	}
}

// SignAgreementInvoker 签署云服务声明
func (c *MetaStudioClient) SignAgreementInvoker(request *model.SignAgreementRequest) *SignAgreementInvoker {
	requestDef := GenReqDefForSignAgreement()
	return &SignAgreementInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SignSpecialAgreement 签署特殊云服务声明
//
// 签署特殊云服务声明,目前可签署自动支付协议。开启自动支付协议之后,调用下单接口时，华为云将进行自动扣费。若因账户余额不足导致扣费失败，系统会生成待支付订单，您可前往费用中心-我的订单查看，届时您需手动完成支付。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) SignSpecialAgreement(request *model.SignSpecialAgreementRequest) (*model.SignSpecialAgreementResponse, error) {
	requestDef := GenReqDefForSignSpecialAgreement()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SignSpecialAgreementResponse), nil
	}
}

// SignSpecialAgreementInvoker 签署特殊云服务声明
func (c *MetaStudioClient) SignSpecialAgreementInvoker(request *model.SignSpecialAgreementRequest) *SignSpecialAgreementInvoker {
	requestDef := GenReqDefForSignSpecialAgreement()
	return &SignSpecialAgreementInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTenantServiceConfigs 设置租户服务配置
//
// 设置租户服务业务配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateTenantServiceConfigs(request *model.UpdateTenantServiceConfigsRequest) (*model.UpdateTenantServiceConfigsResponse, error) {
	requestDef := GenReqDefForUpdateTenantServiceConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTenantServiceConfigsResponse), nil
	}
}

// UpdateTenantServiceConfigsInvoker 设置租户服务配置
func (c *MetaStudioClient) UpdateTenantServiceConfigsInvoker(request *model.UpdateTenantServiceConfigsRequest) *UpdateTenantServiceConfigsInvoker {
	requestDef := GenReqDefForUpdateTenantServiceConfigs()
	return &UpdateTenantServiceConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUserQuotas 设置子账户配额
//
// 设置子账户（IAM用户）配额，需要先开启子账户隔离后才能配置。 只有根账户可修改。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateUserQuotas(request *model.UpdateUserQuotasRequest) (*model.UpdateUserQuotasResponse, error) {
	requestDef := GenReqDefForUpdateUserQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUserQuotasResponse), nil
	}
}

// UpdateUserQuotasInvoker 设置子账户配额
func (c *MetaStudioClient) UpdateUserQuotasInvoker(request *model.UpdateUserQuotasRequest) *UpdateUserQuotasInvoker {
	requestDef := GenReqDefForUpdateUserQuotas()
	return &UpdateUserQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CommitShortJob 提交短任务
//
// 提交短任务，执行该接口后，任务会正式开始执行。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CommitShortJob(request *model.CommitShortJobRequest) (*model.CommitShortJobResponse, error) {
	requestDef := GenReqDefForCommitShortJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CommitShortJobResponse), nil
	}
}

// CommitShortJobInvoker 提交短任务
func (c *MetaStudioClient) CommitShortJobInvoker(request *model.CommitShortJobRequest) *CommitShortJobInvoker {
	requestDef := GenReqDefForCommitShortJob()
	return &CommitShortJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CommitVoiceTrainingJob 提交语音训练任务
//
// 提交训练任务,执行该接口后,任务会进入审核状态,审核完成后会等待训练。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CommitVoiceTrainingJob(request *model.CommitVoiceTrainingJobRequest) (*model.CommitVoiceTrainingJobResponse, error) {
	requestDef := GenReqDefForCommitVoiceTrainingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CommitVoiceTrainingJobResponse), nil
	}
}

// CommitVoiceTrainingJobInvoker 提交语音训练任务
func (c *MetaStudioClient) CommitVoiceTrainingJobInvoker(request *model.CommitVoiceTrainingJobRequest) *CommitVoiceTrainingJobInvoker {
	requestDef := GenReqDefForCommitVoiceTrainingJob()
	return &CommitVoiceTrainingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ConfirmTrainingSegment 确认在线录音结果
//
// 确认在线录音结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ConfirmTrainingSegment(request *model.ConfirmTrainingSegmentRequest) (*model.ConfirmTrainingSegmentResponse, error) {
	requestDef := GenReqDefForConfirmTrainingSegment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ConfirmTrainingSegmentResponse), nil
	}
}

// ConfirmTrainingSegmentInvoker 确认在线录音结果
func (c *MetaStudioClient) ConfirmTrainingSegmentInvoker(request *model.ConfirmTrainingSegmentRequest) *ConfirmTrainingSegmentInvoker {
	requestDef := GenReqDefForConfirmTrainingSegment()
	return &ConfirmTrainingSegmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateShortJob 创建短任务
//
// 用户创建短任务（音频质量检测等），该接口会返回一个obs上传地址，用于上传语音文件。
// 文件上传后，调用“提交短任务”接口，启动短任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateShortJob(request *model.CreateShortJobRequest) (*model.CreateShortJobResponse, error) {
	requestDef := GenReqDefForCreateShortJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateShortJobResponse), nil
	}
}

// CreateShortJobInvoker 创建短任务
func (c *MetaStudioClient) CreateShortJobInvoker(request *model.CreateShortJobRequest) *CreateShortJobInvoker {
	requestDef := GenReqDefForCreateShortJob()
	return &CreateShortJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTrainingAdvanceJob 创建高级版语音训练任务
//
// 用户创建语音训练高级版任务，该接口会返回一个obs上传地址，用于上传语音文件。
//
// 语音文件为一段WAV格式的长音频文件，仅支持将语音文件打包成zip压缩格式上传。
//
// 文件上传后，调用“提交语音训练任务”接口，启动审核和训练。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateTrainingAdvanceJob(request *model.CreateTrainingAdvanceJobRequest) (*model.CreateTrainingAdvanceJobResponse, error) {
	requestDef := GenReqDefForCreateTrainingAdvanceJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTrainingAdvanceJobResponse), nil
	}
}

// CreateTrainingAdvanceJobInvoker 创建高级版语音训练任务
func (c *MetaStudioClient) CreateTrainingAdvanceJobInvoker(request *model.CreateTrainingAdvanceJobRequest) *CreateTrainingAdvanceJobInvoker {
	requestDef := GenReqDefForCreateTrainingAdvanceJob()
	return &CreateTrainingAdvanceJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTrainingBasicJob 创建基础版语音训练任务
//
// 用户创建语音训练基础版任务，该接口会返回一个obs上传地址，用于上传语音文件。
//
// 支持2种方式上传语音文件：
// * 语音文件和文本文件打包成zip上传：语音文件已经切分成20个wav文件，每个语音文件对应一个txt文本文件，所有文件打包成zip文件。语音文件命名规则：0.wav~19.wav；文本文件命名规则：0.txt~19.txt。
// * 语音文件和文本文件逐句上传：每次上传一句语料的语音文件和文本文件，再调用“确认在线录音结果”接口确认语音和文本内容是否一致。确认成功后再上传和确认下一句。
//
// 文件上传后，调用“提交语音训练任务”接口，启动审核和训练。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateTrainingBasicJob(request *model.CreateTrainingBasicJobRequest) (*model.CreateTrainingBasicJobResponse, error) {
	requestDef := GenReqDefForCreateTrainingBasicJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTrainingBasicJobResponse), nil
	}
}

// CreateTrainingBasicJobInvoker 创建基础版语音训练任务
func (c *MetaStudioClient) CreateTrainingBasicJobInvoker(request *model.CreateTrainingBasicJobRequest) *CreateTrainingBasicJobInvoker {
	requestDef := GenReqDefForCreateTrainingBasicJob()
	return &CreateTrainingBasicJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTrainingMiddleJob 创建进阶版语音训练任务
//
// 用户创建语音训练进阶版任务，该接口会返回一个obs上传地址，用于上传语音文件。
//
// 支持2种方式上传语音文件：
// * 语音文件和文本文件打包成zip上传：语音文件已经切分成100个wav文件，每个语音文件对应一个txt文本文件，所有文件打包成zip文件。语音文件命名规则：0.wav~99.wav；文本文件命名规则：0.txt~99.txt。
// * 语音文件和文本文件逐句上传：每次上传一句语料的语音文件和文本文件，再调用“确认在线录音结果”接口确认语音和文本内容是否一致。确认成功后再上传和确认下一句。
//
// 文件上传后，调用“提交语音训练任务”接口，启动审核和训练。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateTrainingMiddleJob(request *model.CreateTrainingMiddleJobRequest) (*model.CreateTrainingMiddleJobResponse, error) {
	requestDef := GenReqDefForCreateTrainingMiddleJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTrainingMiddleJobResponse), nil
	}
}

// CreateTrainingMiddleJobInvoker 创建进阶版语音训练任务
func (c *MetaStudioClient) CreateTrainingMiddleJobInvoker(request *model.CreateTrainingMiddleJobRequest) *CreateTrainingMiddleJobInvoker {
	requestDef := GenReqDefForCreateTrainingMiddleJob()
	return &CreateTrainingMiddleJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTrainingThirdPartyJob 创建第三方平台语音训练任务
//
// 用户创建第三方平台语音训练任务,该接口会返回一个obs上传地址，用于上传语音文件。
// 仅支持zip包方式上传语音文件：
// * 语音文件打包成zip上传：上传的训练数据为一个zip格式压缩文件,其中包含一段wav格式的长音频文件。
//
// &gt; * 文件上传后，调用“提交语音训练任务”接口，启动审核和训练。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateTrainingThirdPartyJob(request *model.CreateTrainingThirdPartyJobRequest) (*model.CreateTrainingThirdPartyJobResponse, error) {
	requestDef := GenReqDefForCreateTrainingThirdPartyJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTrainingThirdPartyJobResponse), nil
	}
}

// CreateTrainingThirdPartyJobInvoker 创建第三方平台语音训练任务
func (c *MetaStudioClient) CreateTrainingThirdPartyJobInvoker(request *model.CreateTrainingThirdPartyJobRequest) *CreateTrainingThirdPartyJobInvoker {
	requestDef := GenReqDefForCreateTrainingThirdPartyJob()
	return &CreateTrainingThirdPartyJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVoiceTrainingJob 删除语音训练任务
//
// 删除语音训练任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteVoiceTrainingJob(request *model.DeleteVoiceTrainingJobRequest) (*model.DeleteVoiceTrainingJobResponse, error) {
	requestDef := GenReqDefForDeleteVoiceTrainingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVoiceTrainingJobResponse), nil
	}
}

// DeleteVoiceTrainingJobInvoker 删除语音训练任务
func (c *MetaStudioClient) DeleteVoiceTrainingJobInvoker(request *model.DeleteVoiceTrainingJobRequest) *DeleteVoiceTrainingJobInvoker {
	requestDef := GenReqDefForDeleteVoiceTrainingJob()
	return &DeleteVoiceTrainingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListJobOperationLog 查询任务操作日志
//
// 查询任务操作日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListJobOperationLog(request *model.ListJobOperationLogRequest) (*model.ListJobOperationLogResponse, error) {
	requestDef := GenReqDefForListJobOperationLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJobOperationLogResponse), nil
	}
}

// ListJobOperationLogInvoker 查询任务操作日志
func (c *MetaStudioClient) ListJobOperationLogInvoker(request *model.ListJobOperationLogRequest) *ListJobOperationLogInvoker {
	requestDef := GenReqDefForListJobOperationLog()
	return &ListJobOperationLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVoiceTrainingJob 查询语音训练任务列表
//
// 查询语音训练任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListVoiceTrainingJob(request *model.ListVoiceTrainingJobRequest) (*model.ListVoiceTrainingJobResponse, error) {
	requestDef := GenReqDefForListVoiceTrainingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVoiceTrainingJobResponse), nil
	}
}

// ListVoiceTrainingJobInvoker 查询语音训练任务列表
func (c *MetaStudioClient) ListVoiceTrainingJobInvoker(request *model.ListVoiceTrainingJobRequest) *ListVoiceTrainingJobInvoker {
	requestDef := GenReqDefForListVoiceTrainingJob()
	return &ListVoiceTrainingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetJobBatchName 设置任务批次
//
// 用户设置任务批次，该接口用于批量任务管理场景，设置任务的批次
// * 需要开通NA租户权限后才能正常调用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) SetJobBatchName(request *model.SetJobBatchNameRequest) (*model.SetJobBatchNameResponse, error) {
	requestDef := GenReqDefForSetJobBatchName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetJobBatchNameResponse), nil
	}
}

// SetJobBatchNameInvoker 设置任务批次
func (c *MetaStudioClient) SetJobBatchNameInvoker(request *model.SetJobBatchNameRequest) *SetJobBatchNameInvoker {
	requestDef := GenReqDefForSetJobBatchName()
	return &SetJobBatchNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEncryptFile 下载加密文件
//
// 下载加密文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowEncryptFile(request *model.ShowEncryptFileRequest) (*model.ShowEncryptFileResponse, error) {
	requestDef := GenReqDefForShowEncryptFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEncryptFileResponse), nil
	}
}

// ShowEncryptFileInvoker 下载加密文件
func (c *MetaStudioClient) ShowEncryptFileInvoker(request *model.ShowEncryptFileRequest) *ShowEncryptFileInvoker {
	requestDef := GenReqDefForShowEncryptFile()
	return &ShowEncryptFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobAuditResult 获取语音训练任务审核结果
//
// 获取语音训练任务审核结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowJobAuditResult(request *model.ShowJobAuditResultRequest) (*model.ShowJobAuditResultResponse, error) {
	requestDef := GenReqDefForShowJobAuditResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobAuditResultResponse), nil
	}
}

// ShowJobAuditResultInvoker 获取语音训练任务审核结果
func (c *MetaStudioClient) ShowJobAuditResultInvoker(request *model.ShowJobAuditResultRequest) *ShowJobAuditResultInvoker {
	requestDef := GenReqDefForShowJobAuditResult()
	return &ShowJobAuditResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobUploadingAddress 获取语音文件上传地址
//
// 获取语音文件上传地址
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowJobUploadingAddress(request *model.ShowJobUploadingAddressRequest) (*model.ShowJobUploadingAddressResponse, error) {
	requestDef := GenReqDefForShowJobUploadingAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobUploadingAddressResponse), nil
	}
}

// ShowJobUploadingAddressInvoker 获取语音文件上传地址
func (c *MetaStudioClient) ShowJobUploadingAddressInvoker(request *model.ShowJobUploadingAddressRequest) *ShowJobUploadingAddressInvoker {
	requestDef := GenReqDefForShowJobUploadingAddress()
	return &ShowJobUploadingAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowShortJob 查询短任务详情
//
// 查询短任务详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowShortJob(request *model.ShowShortJobRequest) (*model.ShowShortJobResponse, error) {
	requestDef := GenReqDefForShowShortJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowShortJobResponse), nil
	}
}

// ShowShortJobInvoker 查询短任务详情
func (c *MetaStudioClient) ShowShortJobInvoker(request *model.ShowShortJobRequest) *ShowShortJobInvoker {
	requestDef := GenReqDefForShowShortJob()
	return &ShowShortJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTenantDurationCfg 查询用户配置的个性化音频时长
//
// 查询用户配置的个性化音频时长
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowTenantDurationCfg(request *model.ShowTenantDurationCfgRequest) (*model.ShowTenantDurationCfgResponse, error) {
	requestDef := GenReqDefForShowTenantDurationCfg()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTenantDurationCfgResponse), nil
	}
}

// ShowTenantDurationCfgInvoker 查询用户配置的个性化音频时长
func (c *MetaStudioClient) ShowTenantDurationCfgInvoker(request *model.ShowTenantDurationCfgRequest) *ShowTenantDurationCfgInvoker {
	requestDef := GenReqDefForShowTenantDurationCfg()
	return &ShowTenantDurationCfgInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTrainingSegmentInfo 获取在线录音确认结果
//
// 获取在线录音确认结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowTrainingSegmentInfo(request *model.ShowTrainingSegmentInfoRequest) (*model.ShowTrainingSegmentInfoResponse, error) {
	requestDef := GenReqDefForShowTrainingSegmentInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTrainingSegmentInfoResponse), nil
	}
}

// ShowTrainingSegmentInfoInvoker 获取在线录音确认结果
func (c *MetaStudioClient) ShowTrainingSegmentInfoInvoker(request *model.ShowTrainingSegmentInfoRequest) *ShowTrainingSegmentInfoInvoker {
	requestDef := GenReqDefForShowTrainingSegmentInfo()
	return &ShowTrainingSegmentInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUserReviewAttachmentUploadingAddress 用户获取附件上传url
//
// 用户获取附件上传url
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowUserReviewAttachmentUploadingAddress(request *model.ShowUserReviewAttachmentUploadingAddressRequest) (*model.ShowUserReviewAttachmentUploadingAddressResponse, error) {
	requestDef := GenReqDefForShowUserReviewAttachmentUploadingAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserReviewAttachmentUploadingAddressResponse), nil
	}
}

// ShowUserReviewAttachmentUploadingAddressInvoker 用户获取附件上传url
func (c *MetaStudioClient) ShowUserReviewAttachmentUploadingAddressInvoker(request *model.ShowUserReviewAttachmentUploadingAddressRequest) *ShowUserReviewAttachmentUploadingAddressInvoker {
	requestDef := GenReqDefForShowUserReviewAttachmentUploadingAddress()
	return &ShowUserReviewAttachmentUploadingAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVoiceTrainingJob 查询语音训练任务详情
//
// 查询语音训练任务详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowVoiceTrainingJob(request *model.ShowVoiceTrainingJobRequest) (*model.ShowVoiceTrainingJobResponse, error) {
	requestDef := GenReqDefForShowVoiceTrainingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVoiceTrainingJobResponse), nil
	}
}

// ShowVoiceTrainingJobInvoker 查询语音训练任务详情
func (c *MetaStudioClient) ShowVoiceTrainingJobInvoker(request *model.ShowVoiceTrainingJobRequest) *ShowVoiceTrainingJobInvoker {
	requestDef := GenReqDefForShowVoiceTrainingJob()
	return &ShowVoiceTrainingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Create2dModelTrainingJob 创建分身数字人模型训练任务
//
// 该接口用于创建分身数字人模型训练任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) Create2dModelTrainingJob(request *model.Create2dModelTrainingJobRequest) (*model.Create2dModelTrainingJobResponse, error) {
	requestDef := GenReqDefForCreate2dModelTrainingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.Create2dModelTrainingJobResponse), nil
	}
}

// Create2dModelTrainingJobInvoker 创建分身数字人模型训练任务
func (c *MetaStudioClient) Create2dModelTrainingJobInvoker(request *model.Create2dModelTrainingJobRequest) *Create2dModelTrainingJobInvoker {
	requestDef := GenReqDefForCreate2dModelTrainingJob()
	return &Create2dModelTrainingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Delete2dModelTrainingJob 删除分身数字人模型训练任务
//
// 该接口用于删除分身数字人模型训练任务。同时需要删除训练任务相关的训练视频、身份证照片、授权文件、模型资产等。
// &gt; * 该接口应当在任务处于以下状态时调用：WAIT_FILE_UPLOAD、AUTO_VERIFY_FAILED、MANUAL_VERIFYING、MANUAL_VERIFY_FAILED、TRAINING_DATA_PREPROCESS_FAILED、TRAIN_FAILED、INFERENCE_DATA_PREPROCESS_FAILED、JOB_SUCCESS、WAIT_USER_CONFIRM、JOB_REJECT、JOB_FINISH
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) Delete2dModelTrainingJob(request *model.Delete2dModelTrainingJobRequest) (*model.Delete2dModelTrainingJobResponse, error) {
	requestDef := GenReqDefForDelete2dModelTrainingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.Delete2dModelTrainingJobResponse), nil
	}
}

// Delete2dModelTrainingJobInvoker 删除分身数字人模型训练任务
func (c *MetaStudioClient) Delete2dModelTrainingJobInvoker(request *model.Delete2dModelTrainingJobRequest) *Delete2dModelTrainingJobInvoker {
	requestDef := GenReqDefForDelete2dModelTrainingJob()
	return &Delete2dModelTrainingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Download2dModelTraningEncryptFile 下载加密文件
//
// 下载加密文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) Download2dModelTraningEncryptFile(request *model.Download2dModelTraningEncryptFileRequest) (*model.Download2dModelTraningEncryptFileResponse, error) {
	requestDef := GenReqDefForDownload2dModelTraningEncryptFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.Download2dModelTraningEncryptFileResponse), nil
	}
}

// Download2dModelTraningEncryptFileInvoker 下载加密文件
func (c *MetaStudioClient) Download2dModelTraningEncryptFileInvoker(request *model.Download2dModelTraningEncryptFileRequest) *Download2dModelTraningEncryptFileInvoker {
	requestDef := GenReqDefForDownload2dModelTraningEncryptFile()
	return &Download2dModelTraningEncryptFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Execute2dModelTrainingCommandByUser 租户执行分身数字人模型训练任务命令
//
// 该接口用于租户执行分身数字人模型训练任务命令，如提交训练审核等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) Execute2dModelTrainingCommandByUser(request *model.Execute2dModelTrainingCommandByUserRequest) (*model.Execute2dModelTrainingCommandByUserResponse, error) {
	requestDef := GenReqDefForExecute2dModelTrainingCommandByUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.Execute2dModelTrainingCommandByUserResponse), nil
	}
}

// Execute2dModelTrainingCommandByUserInvoker 租户执行分身数字人模型训练任务命令
func (c *MetaStudioClient) Execute2dModelTrainingCommandByUserInvoker(request *model.Execute2dModelTrainingCommandByUserRequest) *Execute2dModelTrainingCommandByUserInvoker {
	requestDef := GenReqDefForExecute2dModelTrainingCommandByUser()
	return &Execute2dModelTrainingCommandByUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// List2dModelTrainingJob 查询分身数字人模型训练任务列表
//
// 该接口用于查询分身数字人模型训练任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) List2dModelTrainingJob(request *model.List2dModelTrainingJobRequest) (*model.List2dModelTrainingJobResponse, error) {
	requestDef := GenReqDefForList2dModelTrainingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.List2dModelTrainingJobResponse), nil
	}
}

// List2dModelTrainingJobInvoker 查询分身数字人模型训练任务列表
func (c *MetaStudioClient) List2dModelTrainingJobInvoker(request *model.List2dModelTrainingJobRequest) *List2dModelTrainingJobInvoker {
	requestDef := GenReqDefForList2dModelTrainingJob()
	return &List2dModelTrainingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Show2dModelTrainingJob 查询分身数字人模型训练任务详情
//
// 该接口用于查询分身数字人模型训练任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) Show2dModelTrainingJob(request *model.Show2dModelTrainingJobRequest) (*model.Show2dModelTrainingJobResponse, error) {
	requestDef := GenReqDefForShow2dModelTrainingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.Show2dModelTrainingJobResponse), nil
	}
}

// Show2dModelTrainingJobInvoker 查询分身数字人模型训练任务详情
func (c *MetaStudioClient) Show2dModelTrainingJobInvoker(request *model.Show2dModelTrainingJobRequest) *Show2dModelTrainingJobInvoker {
	requestDef := GenReqDefForShow2dModelTrainingJob()
	return &Show2dModelTrainingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Update2dModelTrainingJob 更新分身数字人模型训练任务
//
// 该接口用于更新分身数字人模型训练任务。用于在自动审核或者人工审核不通过情况下，更新训练视频、身份证照片等。
// &gt; * 该接口只能在AUTO_VERIFY_FAILED或者MANUAL_VERIFY_FAILED状态下调用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) Update2dModelTrainingJob(request *model.Update2dModelTrainingJobRequest) (*model.Update2dModelTrainingJobResponse, error) {
	requestDef := GenReqDefForUpdate2dModelTrainingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.Update2dModelTrainingJobResponse), nil
	}
}

// Update2dModelTrainingJobInvoker 更新分身数字人模型训练任务
func (c *MetaStudioClient) Update2dModelTrainingJobInvoker(request *model.Update2dModelTrainingJobRequest) *Update2dModelTrainingJobInvoker {
	requestDef := GenReqDefForUpdate2dModelTrainingJob()
	return &Update2dModelTrainingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckVoiceAsset 校验音色模型是否可用（自研和第三方音色）
//
// 该接口用于校验音色模型是否可用，模型可用返回模型信息，不可用返回具体不可用的原因
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CheckVoiceAsset(request *model.CheckVoiceAssetRequest) (*model.CheckVoiceAssetResponse, error) {
	requestDef := GenReqDefForCheckVoiceAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckVoiceAssetResponse), nil
	}
}

// CheckVoiceAssetInvoker 校验音色模型是否可用（自研和第三方音色）
func (c *MetaStudioClient) CheckVoiceAssetInvoker(request *model.CheckVoiceAssetRequest) *CheckVoiceAssetInvoker {
	requestDef := GenReqDefForCheckVoiceAsset()
	return &CheckVoiceAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTtsJob 获取TTS语音合成任务记录
//
// 该接口用于获取TTS语音合成任务记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowTtsJob(request *model.ShowTtsJobRequest) (*model.ShowTtsJobResponse, error) {
	requestDef := GenReqDefForShowTtsJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTtsJobResponse), nil
	}
}

// ShowTtsJobInvoker 获取TTS语音合成任务记录
func (c *MetaStudioClient) ShowTtsJobInvoker(request *model.ShowTtsJobRequest) *ShowTtsJobInvoker {
	requestDef := GenReqDefForShowTtsJob()
	return &ShowTtsJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTtsPhoneticSymbol 获取英文单词音标
//
// 根据英文单词返回对应音标列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowTtsPhoneticSymbol(request *model.ShowTtsPhoneticSymbolRequest) (*model.ShowTtsPhoneticSymbolResponse, error) {
	requestDef := GenReqDefForShowTtsPhoneticSymbol()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTtsPhoneticSymbolResponse), nil
	}
}

// ShowTtsPhoneticSymbolInvoker 获取英文单词音标
func (c *MetaStudioClient) ShowTtsPhoneticSymbolInvoker(request *model.ShowTtsPhoneticSymbolRequest) *ShowTtsPhoneticSymbolInvoker {
	requestDef := GenReqDefForShowTtsPhoneticSymbol()
	return &ShowTtsPhoneticSymbolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAsyncTtsJob 创建TTS异步任务
//
// 该接口用于对外生成音频文件。每个预置音色的计费标准详见[预置音色计费标准](metastudio_02_0060.xml)。
//
// &gt; 使用本接口前，需要在MetaStudio控制台服务概览页面，开通“声音合成”的按需计费。
// &gt; 详细操作为：单击“声音合成”卡片中的“去开通”，在弹出的“开通按需计费服务提示”对话框中，勾选同意协议。单击“确定”，开通按需计费。
// &gt; 如需使用第三方声音进行语音合成，请购买出门问问声音套餐，操作请参考《用户指南》的“购买出门问问声音套餐”章节。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateAsyncTtsJob(request *model.CreateAsyncTtsJobRequest) (*model.CreateAsyncTtsJobResponse, error) {
	requestDef := GenReqDefForCreateAsyncTtsJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAsyncTtsJobResponse), nil
	}
}

// CreateAsyncTtsJobInvoker 创建TTS异步任务
func (c *MetaStudioClient) CreateAsyncTtsJobInvoker(request *model.CreateAsyncTtsJobRequest) *CreateAsyncTtsJobInvoker {
	requestDef := GenReqDefForCreateAsyncTtsJob()
	return &CreateAsyncTtsJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTtsAudition 创建TTS试听任务
//
// 该接口用于创建生成播报内容的语音试听文件任务。
//
// [第三方音色试听需要收费，收费标准参考：https://marketplace.huaweicloud.com/product/OFFI919400645308506112#productid&#x3D;OFFI919400645308506112](tag:hc)
//
// [第三方音色试听需要收费，收费标准参考：https://marketplace.huaweicloud.com/intl/contents/939bf377-3005-472b-a4e2-383911e6102f](tag:hk)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateTtsAudition(request *model.CreateTtsAuditionRequest) (*model.CreateTtsAuditionResponse, error) {
	requestDef := GenReqDefForCreateTtsAudition()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTtsAuditionResponse), nil
	}
}

// CreateTtsAuditionInvoker 创建TTS试听任务
func (c *MetaStudioClient) CreateTtsAuditionInvoker(request *model.CreateTtsAuditionRequest) *CreateTtsAuditionInvoker {
	requestDef := GenReqDefForCreateTtsAudition()
	return &CreateTtsAuditionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAsyncTtsJob 获取TTS异步任务
//
// 该接口用于获取TTS音频文件下载链接。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowAsyncTtsJob(request *model.ShowAsyncTtsJobRequest) (*model.ShowAsyncTtsJobResponse, error) {
	requestDef := GenReqDefForShowAsyncTtsJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAsyncTtsJobResponse), nil
	}
}

// ShowAsyncTtsJobInvoker 获取TTS异步任务
func (c *MetaStudioClient) ShowAsyncTtsJobInvoker(request *model.ShowAsyncTtsJobRequest) *ShowAsyncTtsJobInvoker {
	requestDef := GenReqDefForShowAsyncTtsJob()
	return &ShowAsyncTtsJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTtsAuditionFile 获取TTS试听文件
//
// 该接口用于获取TTS试听文件下载链接，返回List中包含当前已生产的试听文件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowTtsAuditionFile(request *model.ShowTtsAuditionFileRequest) (*model.ShowTtsAuditionFileResponse, error) {
	requestDef := GenReqDefForShowTtsAuditionFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTtsAuditionFileResponse), nil
	}
}

// ShowTtsAuditionFileInvoker 获取TTS试听文件
func (c *MetaStudioClient) ShowTtsAuditionFileInvoker(request *model.ShowTtsAuditionFileRequest) *ShowTtsAuditionFileInvoker {
	requestDef := GenReqDefForShowTtsAuditionFile()
	return &ShowTtsAuditionFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTtsOnceCode 外部接口-获取TTS一次性token
//
// 该接口用于获取TTS租户级一次性token。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateTtsOnceCode(request *model.CreateTtsOnceCodeRequest) (*model.CreateTtsOnceCodeResponse, error) {
	requestDef := GenReqDefForCreateTtsOnceCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTtsOnceCodeResponse), nil
	}
}

// CreateTtsOnceCodeInvoker 外部接口-获取TTS一次性token
func (c *MetaStudioClient) CreateTtsOnceCodeInvoker(request *model.CreateTtsOnceCodeRequest) *CreateTtsOnceCodeInvoker {
	requestDef := GenReqDefForCreateTtsOnceCode()
	return &CreateTtsOnceCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTtscVocabularyConfigs 设置TTS租户级自定义读法配置
//
// 该接口用于设置TTS租户级自定义读法配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateTtscVocabularyConfigs(request *model.CreateTtscVocabularyConfigsRequest) (*model.CreateTtscVocabularyConfigsResponse, error) {
	requestDef := GenReqDefForCreateTtscVocabularyConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTtscVocabularyConfigsResponse), nil
	}
}

// CreateTtscVocabularyConfigsInvoker 设置TTS租户级自定义读法配置
func (c *MetaStudioClient) CreateTtscVocabularyConfigsInvoker(request *model.CreateTtscVocabularyConfigsRequest) *CreateTtscVocabularyConfigsInvoker {
	requestDef := GenReqDefForCreateTtscVocabularyConfigs()
	return &CreateTtscVocabularyConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTtscVocabularyGroups 设置TTS租户级词表分组配置
//
// 该接口用于设置TTS租户级词表分组配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateTtscVocabularyGroups(request *model.CreateTtscVocabularyGroupsRequest) (*model.CreateTtscVocabularyGroupsResponse, error) {
	requestDef := GenReqDefForCreateTtscVocabularyGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTtscVocabularyGroupsResponse), nil
	}
}

// CreateTtscVocabularyGroupsInvoker 设置TTS租户级词表分组配置
func (c *MetaStudioClient) CreateTtscVocabularyGroupsInvoker(request *model.CreateTtscVocabularyGroupsRequest) *CreateTtscVocabularyGroupsInvoker {
	requestDef := GenReqDefForCreateTtscVocabularyGroups()
	return &CreateTtscVocabularyGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTtscVocabularyConfigs 删除TTS租户级自定义读法配置
//
// 该接口用于删除TTS租户级自定义读法配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteTtscVocabularyConfigs(request *model.DeleteTtscVocabularyConfigsRequest) (*model.DeleteTtscVocabularyConfigsResponse, error) {
	requestDef := GenReqDefForDeleteTtscVocabularyConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTtscVocabularyConfigsResponse), nil
	}
}

// DeleteTtscVocabularyConfigsInvoker 删除TTS租户级自定义读法配置
func (c *MetaStudioClient) DeleteTtscVocabularyConfigsInvoker(request *model.DeleteTtscVocabularyConfigsRequest) *DeleteTtscVocabularyConfigsInvoker {
	requestDef := GenReqDefForDeleteTtscVocabularyConfigs()
	return &DeleteTtscVocabularyConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTtscVocabularyGroups 删除TTS租户级词表分组
//
// 该接口用于删除TTS租户级词表分组配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteTtscVocabularyGroups(request *model.DeleteTtscVocabularyGroupsRequest) (*model.DeleteTtscVocabularyGroupsResponse, error) {
	requestDef := GenReqDefForDeleteTtscVocabularyGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTtscVocabularyGroupsResponse), nil
	}
}

// DeleteTtscVocabularyGroupsInvoker 删除TTS租户级词表分组
func (c *MetaStudioClient) DeleteTtscVocabularyGroupsInvoker(request *model.DeleteTtscVocabularyGroupsRequest) *DeleteTtscVocabularyGroupsInvoker {
	requestDef := GenReqDefForDeleteTtscVocabularyGroups()
	return &DeleteTtscVocabularyGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTtscVocabularyConfigs 获取TTS租户级自定义读法配置
//
// 该接口用于获取TTS租户级自定义读法配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListTtscVocabularyConfigs(request *model.ListTtscVocabularyConfigsRequest) (*model.ListTtscVocabularyConfigsResponse, error) {
	requestDef := GenReqDefForListTtscVocabularyConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTtscVocabularyConfigsResponse), nil
	}
}

// ListTtscVocabularyConfigsInvoker 获取TTS租户级自定义读法配置
func (c *MetaStudioClient) ListTtscVocabularyConfigsInvoker(request *model.ListTtscVocabularyConfigsRequest) *ListTtscVocabularyConfigsInvoker {
	requestDef := GenReqDefForListTtscVocabularyConfigs()
	return &ListTtscVocabularyConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTtscVocabularyGroups 获取TTS租户级词表分组列表
//
// 该接口用于获取TTS租户级词表分组列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListTtscVocabularyGroups(request *model.ListTtscVocabularyGroupsRequest) (*model.ListTtscVocabularyGroupsResponse, error) {
	requestDef := GenReqDefForListTtscVocabularyGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTtscVocabularyGroupsResponse), nil
	}
}

// ListTtscVocabularyGroupsInvoker 获取TTS租户级词表分组列表
func (c *MetaStudioClient) ListTtscVocabularyGroupsInvoker(request *model.ListTtscVocabularyGroupsRequest) *ListTtscVocabularyGroupsInvoker {
	requestDef := GenReqDefForListTtscVocabularyGroups()
	return &ListTtscVocabularyGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SaveTtscTenantConfigs 设置租户级配置
//
// 该接口用于设置租户级配置，当前用于租户级自定义读法配置的全局开关。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) SaveTtscTenantConfigs(request *model.SaveTtscTenantConfigsRequest) (*model.SaveTtscTenantConfigsResponse, error) {
	requestDef := GenReqDefForSaveTtscTenantConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SaveTtscTenantConfigsResponse), nil
	}
}

// SaveTtscTenantConfigsInvoker 设置租户级配置
func (c *MetaStudioClient) SaveTtscTenantConfigsInvoker(request *model.SaveTtscTenantConfigsRequest) *SaveTtscTenantConfigsInvoker {
	requestDef := GenReqDefForSaveTtscTenantConfigs()
	return &SaveTtscTenantConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SaveTtscVocabularyConfigs 修改TTS租户级自定义读法配置
//
// 该接口用于修改TTS租户级自定义读法配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) SaveTtscVocabularyConfigs(request *model.SaveTtscVocabularyConfigsRequest) (*model.SaveTtscVocabularyConfigsResponse, error) {
	requestDef := GenReqDefForSaveTtscVocabularyConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SaveTtscVocabularyConfigsResponse), nil
	}
}

// SaveTtscVocabularyConfigsInvoker 修改TTS租户级自定义读法配置
func (c *MetaStudioClient) SaveTtscVocabularyConfigsInvoker(request *model.SaveTtscVocabularyConfigsRequest) *SaveTtscVocabularyConfigsInvoker {
	requestDef := GenReqDefForSaveTtscVocabularyConfigs()
	return &SaveTtscVocabularyConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetTtscGroupAssets 设置TTS租户级词表分组的资产列表
//
// 该接口用于设置TTS租户级词表分组的资产列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) SetTtscGroupAssets(request *model.SetTtscGroupAssetsRequest) (*model.SetTtscGroupAssetsResponse, error) {
	requestDef := GenReqDefForSetTtscGroupAssets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetTtscGroupAssetsResponse), nil
	}
}

// SetTtscGroupAssetsInvoker 设置TTS租户级词表分组的资产列表
func (c *MetaStudioClient) SetTtscGroupAssetsInvoker(request *model.SetTtscGroupAssetsRequest) *SetTtscGroupAssetsInvoker {
	requestDef := GenReqDefForSetTtscGroupAssets()
	return &SetTtscGroupAssetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVocabularySwitchConfigs 获取租户级全局配置
//
// 该接口用于获取租户级全局配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowVocabularySwitchConfigs(request *model.ShowVocabularySwitchConfigsRequest) (*model.ShowVocabularySwitchConfigsResponse, error) {
	requestDef := GenReqDefForShowVocabularySwitchConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVocabularySwitchConfigsResponse), nil
	}
}

// ShowVocabularySwitchConfigsInvoker 获取租户级全局配置
func (c *MetaStudioClient) ShowVocabularySwitchConfigsInvoker(request *model.ShowVocabularySwitchConfigsRequest) *ShowVocabularySwitchConfigsInvoker {
	requestDef := GenReqDefForShowVocabularySwitchConfigs()
	return &ShowVocabularySwitchConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTtscVocabularyGroups TTS租户级词表分组重命名
//
// 该接口用于对TTS租户级词表分组重命名。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateTtscVocabularyGroups(request *model.UpdateTtscVocabularyGroupsRequest) (*model.UpdateTtscVocabularyGroupsResponse, error) {
	requestDef := GenReqDefForUpdateTtscVocabularyGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTtscVocabularyGroupsResponse), nil
	}
}

// UpdateTtscVocabularyGroupsInvoker TTS租户级词表分组重命名
func (c *MetaStudioClient) UpdateTtscVocabularyGroupsInvoker(request *model.UpdateTtscVocabularyGroupsRequest) *UpdateTtscVocabularyGroupsInvoker {
	requestDef := GenReqDefForUpdateTtscVocabularyGroups()
	return &UpdateTtscVocabularyGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CopyVideoScripts 复制视频制作剧本
//
// 该接口用于复制视频制作剧本。
// &gt; - 复制的剧本不包含预览字幕信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CopyVideoScripts(request *model.CopyVideoScriptsRequest) (*model.CopyVideoScriptsResponse, error) {
	requestDef := GenReqDefForCopyVideoScripts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyVideoScriptsResponse), nil
	}
}

// CopyVideoScriptsInvoker 复制视频制作剧本
func (c *MetaStudioClient) CopyVideoScriptsInvoker(request *model.CopyVideoScriptsRequest) *CopyVideoScriptsInvoker {
	requestDef := GenReqDefForCopyVideoScripts()
	return &CopyVideoScriptsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVideoScripts 创建视频制作剧本
//
// 该接口用于创建视频制作剧本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateVideoScripts(request *model.CreateVideoScriptsRequest) (*model.CreateVideoScriptsResponse, error) {
	requestDef := GenReqDefForCreateVideoScripts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVideoScriptsResponse), nil
	}
}

// CreateVideoScriptsInvoker 创建视频制作剧本
func (c *MetaStudioClient) CreateVideoScriptsInvoker(request *model.CreateVideoScriptsRequest) *CreateVideoScriptsInvoker {
	requestDef := GenReqDefForCreateVideoScripts()
	return &CreateVideoScriptsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVideoScript 删除视频制作剧本
//
// 该接口用于删除视频制作剧本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteVideoScript(request *model.DeleteVideoScriptRequest) (*model.DeleteVideoScriptResponse, error) {
	requestDef := GenReqDefForDeleteVideoScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVideoScriptResponse), nil
	}
}

// DeleteVideoScriptInvoker 删除视频制作剧本
func (c *MetaStudioClient) DeleteVideoScriptInvoker(request *model.DeleteVideoScriptRequest) *DeleteVideoScriptInvoker {
	requestDef := GenReqDefForDeleteVideoScript()
	return &DeleteVideoScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVideoScripts 查询视频制作剧本列表
//
// 该接口用于查询视频制作剧本列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListVideoScripts(request *model.ListVideoScriptsRequest) (*model.ListVideoScriptsResponse, error) {
	requestDef := GenReqDefForListVideoScripts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVideoScriptsResponse), nil
	}
}

// ListVideoScriptsInvoker 查询视频制作剧本列表
func (c *MetaStudioClient) ListVideoScriptsInvoker(request *model.ListVideoScriptsRequest) *ListVideoScriptsInvoker {
	requestDef := GenReqDefForListVideoScripts()
	return &ListVideoScriptsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVideoScript 查询视频制作剧本详情
//
// 该接口用于查询视频制作剧本详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowVideoScript(request *model.ShowVideoScriptRequest) (*model.ShowVideoScriptResponse, error) {
	requestDef := GenReqDefForShowVideoScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVideoScriptResponse), nil
	}
}

// ShowVideoScriptInvoker 查询视频制作剧本详情
func (c *MetaStudioClient) ShowVideoScriptInvoker(request *model.ShowVideoScriptRequest) *ShowVideoScriptInvoker {
	requestDef := GenReqDefForShowVideoScript()
	return &ShowVideoScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVideoScript 更新视频制作剧本
//
// 该接口用于更新视频制作剧本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateVideoScript(request *model.UpdateVideoScriptRequest) (*model.UpdateVideoScriptResponse, error) {
	requestDef := GenReqDefForUpdateVideoScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVideoScriptResponse), nil
	}
}

// UpdateVideoScriptInvoker 更新视频制作剧本
func (c *MetaStudioClient) UpdateVideoScriptInvoker(request *model.UpdateVideoScriptRequest) *UpdateVideoScriptInvoker {
	requestDef := GenReqDefForUpdateVideoScript()
	return &UpdateVideoScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWelcomeSpeech 创建欢迎词
//
// 该接口用于创建欢迎词。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateWelcomeSpeech(request *model.CreateWelcomeSpeechRequest) (*model.CreateWelcomeSpeechResponse, error) {
	requestDef := GenReqDefForCreateWelcomeSpeech()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWelcomeSpeechResponse), nil
	}
}

// CreateWelcomeSpeechInvoker 创建欢迎词
func (c *MetaStudioClient) CreateWelcomeSpeechInvoker(request *model.CreateWelcomeSpeechRequest) *CreateWelcomeSpeechInvoker {
	requestDef := GenReqDefForCreateWelcomeSpeech()
	return &CreateWelcomeSpeechInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteWelcomeSpeech 删除欢迎词
//
// 该接口用于删除欢迎词。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteWelcomeSpeech(request *model.DeleteWelcomeSpeechRequest) (*model.DeleteWelcomeSpeechResponse, error) {
	requestDef := GenReqDefForDeleteWelcomeSpeech()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWelcomeSpeechResponse), nil
	}
}

// DeleteWelcomeSpeechInvoker 删除欢迎词
func (c *MetaStudioClient) DeleteWelcomeSpeechInvoker(request *model.DeleteWelcomeSpeechRequest) *DeleteWelcomeSpeechInvoker {
	requestDef := GenReqDefForDeleteWelcomeSpeech()
	return &DeleteWelcomeSpeechInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWelcomeSpeech 查询欢迎词列表
//
// 该接口用于查询欢迎词列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListWelcomeSpeech(request *model.ListWelcomeSpeechRequest) (*model.ListWelcomeSpeechResponse, error) {
	requestDef := GenReqDefForListWelcomeSpeech()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWelcomeSpeechResponse), nil
	}
}

// ListWelcomeSpeechInvoker 查询欢迎词列表
func (c *MetaStudioClient) ListWelcomeSpeechInvoker(request *model.ListWelcomeSpeechRequest) *ListWelcomeSpeechInvoker {
	requestDef := GenReqDefForListWelcomeSpeech()
	return &ListWelcomeSpeechInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWelcomeSpeech 查询欢迎词详情
//
// 该接口用于查询欢迎词详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowWelcomeSpeech(request *model.ShowWelcomeSpeechRequest) (*model.ShowWelcomeSpeechResponse, error) {
	requestDef := GenReqDefForShowWelcomeSpeech()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWelcomeSpeechResponse), nil
	}
}

// ShowWelcomeSpeechInvoker 查询欢迎词详情
func (c *MetaStudioClient) ShowWelcomeSpeechInvoker(request *model.ShowWelcomeSpeechRequest) *ShowWelcomeSpeechInvoker {
	requestDef := GenReqDefForShowWelcomeSpeech()
	return &ShowWelcomeSpeechInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWelcomeSpeechSwitch 查询欢迎词功能开关
//
// 该接口用于查询欢迎词功能开关。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowWelcomeSpeechSwitch(request *model.ShowWelcomeSpeechSwitchRequest) (*model.ShowWelcomeSpeechSwitchResponse, error) {
	requestDef := GenReqDefForShowWelcomeSpeechSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWelcomeSpeechSwitchResponse), nil
	}
}

// ShowWelcomeSpeechSwitchInvoker 查询欢迎词功能开关
func (c *MetaStudioClient) ShowWelcomeSpeechSwitchInvoker(request *model.ShowWelcomeSpeechSwitchRequest) *ShowWelcomeSpeechSwitchInvoker {
	requestDef := GenReqDefForShowWelcomeSpeechSwitch()
	return &ShowWelcomeSpeechSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWelcomeSpeech 修改欢迎词
//
// 该接口用于修改欢迎词。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateWelcomeSpeech(request *model.UpdateWelcomeSpeechRequest) (*model.UpdateWelcomeSpeechResponse, error) {
	requestDef := GenReqDefForUpdateWelcomeSpeech()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWelcomeSpeechResponse), nil
	}
}

// UpdateWelcomeSpeechInvoker 修改欢迎词
func (c *MetaStudioClient) UpdateWelcomeSpeechInvoker(request *model.UpdateWelcomeSpeechRequest) *UpdateWelcomeSpeechInvoker {
	requestDef := GenReqDefForUpdateWelcomeSpeech()
	return &UpdateWelcomeSpeechInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWelcomeSpeechSwitch 修改欢迎词功能开关
//
// 该接口用于修改欢迎词功能开关。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateWelcomeSpeechSwitch(request *model.UpdateWelcomeSpeechSwitchRequest) (*model.UpdateWelcomeSpeechSwitchResponse, error) {
	requestDef := GenReqDefForUpdateWelcomeSpeechSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWelcomeSpeechSwitchResponse), nil
	}
}

// UpdateWelcomeSpeechSwitchInvoker 修改欢迎词功能开关
func (c *MetaStudioClient) UpdateWelcomeSpeechSwitchInvoker(request *model.UpdateWelcomeSpeechSwitchRequest) *UpdateWelcomeSpeechSwitchInvoker {
	requestDef := GenReqDefForUpdateWelcomeSpeechSwitch()
	return &UpdateWelcomeSpeechSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
