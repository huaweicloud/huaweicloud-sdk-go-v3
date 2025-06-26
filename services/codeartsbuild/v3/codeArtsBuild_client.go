package v3

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codeartsbuild/v3/model"
)

type CodeArtsBuildClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCodeArtsBuildClient(hcClient *httpclient.HcHttpClient) *CodeArtsBuildClient {
	return &CodeArtsBuildClient{HcClient: hcClient}
}

func CodeArtsBuildClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateBuildJob 创建构建任务
//
// 创建构建任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) CreateBuildJob(request *model.CreateBuildJobRequest) (*model.CreateBuildJobResponse, error) {
	requestDef := GenReqDefForCreateBuildJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBuildJobResponse), nil
	}
}

// CreateBuildJobInvoker 创建构建任务
func (c *CodeArtsBuildClient) CreateBuildJobInvoker(request *model.CreateBuildJobRequest) *CreateBuildJobInvoker {
	requestDef := GenReqDefForCreateBuildJob()
	return &CreateBuildJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTemplates 创建构建模板
//
// 创建构建模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) CreateTemplates(request *model.CreateTemplatesRequest) (*model.CreateTemplatesResponse, error) {
	requestDef := GenReqDefForCreateTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTemplatesResponse), nil
	}
}

// CreateTemplatesInvoker 创建构建模板
func (c *CodeArtsBuildClient) CreateTemplatesInvoker(request *model.CreateTemplatesRequest) *CreateTemplatesInvoker {
	requestDef := GenReqDefForCreateTemplates()
	return &CreateTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBuildJob 删除构建任务
//
// 删除构建任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) DeleteBuildJob(request *model.DeleteBuildJobRequest) (*model.DeleteBuildJobResponse, error) {
	requestDef := GenReqDefForDeleteBuildJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBuildJobResponse), nil
	}
}

// DeleteBuildJobInvoker 删除构建任务
func (c *CodeArtsBuildClient) DeleteBuildJobInvoker(request *model.DeleteBuildJobRequest) *DeleteBuildJobInvoker {
	requestDef := GenReqDefForDeleteBuildJob()
	return &DeleteBuildJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTemplates 删除构建模板
//
// 删除构建模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) DeleteTemplates(request *model.DeleteTemplatesRequest) (*model.DeleteTemplatesResponse, error) {
	requestDef := GenReqDefForDeleteTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTemplatesResponse), nil
	}
}

// DeleteTemplatesInvoker 删除构建模板
func (c *CodeArtsBuildClient) DeleteTemplatesInvoker(request *model.DeleteTemplatesRequest) *DeleteTemplatesInvoker {
	requestDef := GenReqDefForDeleteTemplates()
	return &DeleteTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableBuildJob 禁用构建任务
//
// 禁用构建任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) DisableBuildJob(request *model.DisableBuildJobRequest) (*model.DisableBuildJobResponse, error) {
	requestDef := GenReqDefForDisableBuildJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableBuildJobResponse), nil
	}
}

// DisableBuildJobInvoker 禁用构建任务
func (c *CodeArtsBuildClient) DisableBuildJobInvoker(request *model.DisableBuildJobRequest) *DisableBuildJobInvoker {
	requestDef := GenReqDefForDisableBuildJob()
	return &DisableBuildJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableNotice 取消通知
//
// 取消通知
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) DisableNotice(request *model.DisableNoticeRequest) (*model.DisableNoticeResponse, error) {
	requestDef := GenReqDefForDisableNotice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableNoticeResponse), nil
	}
}

// DisableNoticeInvoker 取消通知
func (c *CodeArtsBuildClient) DisableNoticeInvoker(request *model.DisableNoticeRequest) *DisableNoticeInvoker {
	requestDef := GenReqDefForDisableNotice()
	return &DisableNoticeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadBuildLog 下载全量构建日志
//
// 下载全量构建日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) DownloadBuildLog(request *model.DownloadBuildLogRequest) (*model.DownloadBuildLogResponse, error) {
	requestDef := GenReqDefForDownloadBuildLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadBuildLogResponse), nil
	}
}

// DownloadBuildLogInvoker 下载全量构建日志
func (c *CodeArtsBuildClient) DownloadBuildLogInvoker(request *model.DownloadBuildLogRequest) *DownloadBuildLogInvoker {
	requestDef := GenReqDefForDownloadBuildLog()
	return &DownloadBuildLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadKeystore KeyStore文件下载
//
// 下载指定租户下的KeyStore文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) DownloadKeystore(request *model.DownloadKeystoreRequest) (*model.DownloadKeystoreResponse, error) {
	requestDef := GenReqDefForDownloadKeystore()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadKeystoreResponse), nil
	}
}

// DownloadKeystoreInvoker KeyStore文件下载
func (c *CodeArtsBuildClient) DownloadKeystoreInvoker(request *model.DownloadKeystoreRequest) *DownloadKeystoreInvoker {
	requestDef := GenReqDefForDownloadKeystore()
	return &DownloadKeystoreInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadRealTimeLog 下载构建实时日志
//
// 下载构建实时日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) DownloadRealTimeLog(request *model.DownloadRealTimeLogRequest) (*model.DownloadRealTimeLogResponse, error) {
	requestDef := GenReqDefForDownloadRealTimeLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadRealTimeLogResponse), nil
	}
}

// DownloadRealTimeLogInvoker 下载构建实时日志
func (c *CodeArtsBuildClient) DownloadRealTimeLogInvoker(request *model.DownloadRealTimeLogRequest) *DownloadRealTimeLogInvoker {
	requestDef := GenReqDefForDownloadRealTimeLog()
	return &DownloadRealTimeLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadTaskLog 下载构建步骤日志
//
// 下载构建步骤日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) DownloadTaskLog(request *model.DownloadTaskLogRequest) (*model.DownloadTaskLogResponse, error) {
	requestDef := GenReqDefForDownloadTaskLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadTaskLogResponse), nil
	}
}

// DownloadTaskLogInvoker 下载构建步骤日志
func (c *CodeArtsBuildClient) DownloadTaskLogInvoker(request *model.DownloadTaskLogRequest) *DownloadTaskLogInvoker {
	requestDef := GenReqDefForDownloadTaskLog()
	return &DownloadTaskLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableBuildJob 恢复构建任务
//
// 恢复构建任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) EnableBuildJob(request *model.EnableBuildJobRequest) (*model.EnableBuildJobResponse, error) {
	requestDef := GenReqDefForEnableBuildJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableBuildJobResponse), nil
	}
}

// EnableBuildJobInvoker 恢复构建任务
func (c *CodeArtsBuildClient) EnableBuildJobInvoker(request *model.EnableBuildJobRequest) *EnableBuildJobInvoker {
	requestDef := GenReqDefForEnableBuildJob()
	return &EnableBuildJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBuildInfoRecord 获取任务构建记录列表
//
// 获取任务构建记录列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ListBuildInfoRecord(request *model.ListBuildInfoRecordRequest) (*model.ListBuildInfoRecordResponse, error) {
	requestDef := GenReqDefForListBuildInfoRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBuildInfoRecordResponse), nil
	}
}

// ListBuildInfoRecordInvoker 获取任务构建记录列表
func (c *CodeArtsBuildClient) ListBuildInfoRecordInvoker(request *model.ListBuildInfoRecordRequest) *ListBuildInfoRecordInvoker {
	requestDef := GenReqDefForListBuildInfoRecord()
	return &ListBuildInfoRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListJobConfig 获取构建任务详情
//
// 获取构建任务详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ListJobConfig(request *model.ListJobConfigRequest) (*model.ListJobConfigResponse, error) {
	requestDef := GenReqDefForListJobConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJobConfigResponse), nil
	}
}

// ListJobConfigInvoker 获取构建任务详情
func (c *CodeArtsBuildClient) ListJobConfigInvoker(request *model.ListJobConfigRequest) *ListJobConfigInvoker {
	requestDef := GenReqDefForListJobConfig()
	return &ListJobConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNotice 查询通知
//
// 查询通知
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ListNotice(request *model.ListNoticeRequest) (*model.ListNoticeResponse, error) {
	requestDef := GenReqDefForListNotice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNoticeResponse), nil
	}
}

// ListNoticeInvoker 查询通知
func (c *CodeArtsBuildClient) ListNoticeInvoker(request *model.ListNoticeRequest) *ListNoticeInvoker {
	requestDef := GenReqDefForListNotice()
	return &ListNoticeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTemplates 查询构建模板
//
// 查询构建模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ListTemplates(request *model.ListTemplatesRequest) (*model.ListTemplatesResponse, error) {
	requestDef := GenReqDefForListTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplatesResponse), nil
	}
}

// ListTemplatesInvoker 查询构建模板
func (c *CodeArtsBuildClient) ListTemplatesInvoker(request *model.ListTemplatesRequest) *ListTemplatesInvoker {
	requestDef := GenReqDefForListTemplates()
	return &ListTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunJob 执行构建任务
//
// 执行构建任务,可传自定义参数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) RunJob(request *model.RunJobRequest) (*model.RunJobResponse, error) {
	requestDef := GenReqDefForRunJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunJobResponse), nil
	}
}

// RunJobInvoker 执行构建任务
func (c *CodeArtsBuildClient) RunJobInvoker(request *model.RunJobRequest) *RunJobInvoker {
	requestDef := GenReqDefForRunJob()
	return &RunJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHistoryDetails 获取构建历史详情信息接口
//
// 获取构建历史详情信息接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowHistoryDetails(request *model.ShowHistoryDetailsRequest) (*model.ShowHistoryDetailsResponse, error) {
	requestDef := GenReqDefForShowHistoryDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHistoryDetailsResponse), nil
	}
}

// ShowHistoryDetailsInvoker 获取构建历史详情信息接口
func (c *CodeArtsBuildClient) ShowHistoryDetailsInvoker(request *model.ShowHistoryDetailsRequest) *ShowHistoryDetailsInvoker {
	requestDef := GenReqDefForShowHistoryDetails()
	return &ShowHistoryDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobListByProjectId 查看项目下用户的构建任务列表
//
// 查看项目下用户的构建任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowJobListByProjectId(request *model.ShowJobListByProjectIdRequest) (*model.ShowJobListByProjectIdResponse, error) {
	requestDef := GenReqDefForShowJobListByProjectId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobListByProjectIdResponse), nil
	}
}

// ShowJobListByProjectIdInvoker 查看项目下用户的构建任务列表
func (c *CodeArtsBuildClient) ShowJobListByProjectIdInvoker(request *model.ShowJobListByProjectIdRequest) *ShowJobListByProjectIdInvoker {
	requestDef := GenReqDefForShowJobListByProjectId()
	return &ShowJobListByProjectIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobStatus 查看任务运行状态
//
// 查看任务运行状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowJobStatus(request *model.ShowJobStatusRequest) (*model.ShowJobStatusResponse, error) {
	requestDef := GenReqDefForShowJobStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobStatusResponse), nil
	}
}

// ShowJobStatusInvoker 查看任务运行状态
func (c *CodeArtsBuildClient) ShowJobStatusInvoker(request *model.ShowJobStatusRequest) *ShowJobStatusInvoker {
	requestDef := GenReqDefForShowJobStatus()
	return &ShowJobStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobSuccessRatio 根据开始时间和结束时间查看构建任务的构建成功率
//
// 根据开始时间和结束时间查看构建任务的构建成功率
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowJobSuccessRatio(request *model.ShowJobSuccessRatioRequest) (*model.ShowJobSuccessRatioResponse, error) {
	requestDef := GenReqDefForShowJobSuccessRatio()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobSuccessRatioResponse), nil
	}
}

// ShowJobSuccessRatioInvoker 根据开始时间和结束时间查看构建任务的构建成功率
func (c *CodeArtsBuildClient) ShowJobSuccessRatioInvoker(request *model.ShowJobSuccessRatioRequest) *ShowJobSuccessRatioInvoker {
	requestDef := GenReqDefForShowJobSuccessRatio()
	return &ShowJobSuccessRatioInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLastHistory 查询指定代码仓库最近一次成功的构建历史
//
// 查询指定代码仓库最近一次成功的构建历史
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowLastHistory(request *model.ShowLastHistoryRequest) (*model.ShowLastHistoryResponse, error) {
	requestDef := GenReqDefForShowLastHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLastHistoryResponse), nil
	}
}

// ShowLastHistoryInvoker 查询指定代码仓库最近一次成功的构建历史
func (c *CodeArtsBuildClient) ShowLastHistoryInvoker(request *model.ShowLastHistoryRequest) *ShowLastHistoryInvoker {
	requestDef := GenReqDefForShowLastHistory()
	return &ShowLastHistoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowListHistory 查看构建任务的构建历史列表
//
// 查看构建任务的构建历史列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowListHistory(request *model.ShowListHistoryRequest) (*model.ShowListHistoryResponse, error) {
	requestDef := GenReqDefForShowListHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowListHistoryResponse), nil
	}
}

// ShowListHistoryInvoker 查看构建任务的构建历史列表
func (c *CodeArtsBuildClient) ShowListHistoryInvoker(request *model.ShowListHistoryRequest) *ShowListHistoryInvoker {
	requestDef := GenReqDefForShowListHistory()
	return &ShowListHistoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowListPeriodHistory 根据开始时间和结束时间查看构建任务的构建历史列表
//
// 根据开始时间和结束时间查看构建任务的构建历史列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowListPeriodHistory(request *model.ShowListPeriodHistoryRequest) (*model.ShowListPeriodHistoryResponse, error) {
	requestDef := GenReqDefForShowListPeriodHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowListPeriodHistoryResponse), nil
	}
}

// ShowListPeriodHistoryInvoker 根据开始时间和结束时间查看构建任务的构建历史列表
func (c *CodeArtsBuildClient) ShowListPeriodHistoryInvoker(request *model.ShowListPeriodHistoryRequest) *ShowListPeriodHistoryInvoker {
	requestDef := GenReqDefForShowListPeriodHistory()
	return &ShowListPeriodHistoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOutputInfo 获取构建产物详情信息
//
// 获取构建产物详情信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowOutputInfo(request *model.ShowOutputInfoRequest) (*model.ShowOutputInfoResponse, error) {
	requestDef := GenReqDefForShowOutputInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOutputInfoResponse), nil
	}
}

// ShowOutputInfoInvoker 获取构建产物详情信息
func (c *CodeArtsBuildClient) ShowOutputInfoInvoker(request *model.ShowOutputInfoRequest) *ShowOutputInfoInvoker {
	requestDef := GenReqDefForShowOutputInfo()
	return &ShowOutputInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRecordDetail 获取构建记录信息
//
// 获取构建记录信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowRecordDetail(request *model.ShowRecordDetailRequest) (*model.ShowRecordDetailResponse, error) {
	requestDef := GenReqDefForShowRecordDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecordDetailResponse), nil
	}
}

// ShowRecordDetailInvoker 获取构建记录信息
func (c *CodeArtsBuildClient) ShowRecordDetailInvoker(request *model.ShowRecordDetailRequest) *ShowRecordDetailInvoker {
	requestDef := GenReqDefForShowRecordDetail()
	return &ShowRecordDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopBuildJob 停止构建任务
//
// 停止构建任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) StopBuildJob(request *model.StopBuildJobRequest) (*model.StopBuildJobResponse, error) {
	requestDef := GenReqDefForStopBuildJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopBuildJobResponse), nil
	}
}

// StopBuildJobInvoker 停止构建任务
func (c *CodeArtsBuildClient) StopBuildJobInvoker(request *model.StopBuildJobRequest) *StopBuildJobInvoker {
	requestDef := GenReqDefForStopBuildJob()
	return &StopBuildJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateBuildJob 更新构建任务
//
// 更新构建任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) UpdateBuildJob(request *model.UpdateBuildJobRequest) (*model.UpdateBuildJobResponse, error) {
	requestDef := GenReqDefForUpdateBuildJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBuildJobResponse), nil
	}
}

// UpdateBuildJobInvoker 更新构建任务
func (c *CodeArtsBuildClient) UpdateBuildJobInvoker(request *model.UpdateBuildJobRequest) *UpdateBuildJobInvoker {
	requestDef := GenReqDefForUpdateBuildJob()
	return &UpdateBuildJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNotice 更新通知
//
// 更新通知
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) UpdateNotice(request *model.UpdateNoticeRequest) (*model.UpdateNoticeResponse, error) {
	requestDef := GenReqDefForUpdateNotice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNoticeResponse), nil
	}
}

// UpdateNoticeInvoker 更新通知
func (c *CodeArtsBuildClient) UpdateNoticeInvoker(request *model.UpdateNoticeRequest) *UpdateNoticeInvoker {
	requestDef := GenReqDefForUpdateNotice()
	return &UpdateNoticeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRelatedProjectInfo 获取项目列表
//
// 获取项目列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ListRelatedProjectInfo(request *model.ListRelatedProjectInfoRequest) (*model.ListRelatedProjectInfoResponse, error) {
	requestDef := GenReqDefForListRelatedProjectInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRelatedProjectInfoResponse), nil
	}
}

// ListRelatedProjectInfoInvoker 获取项目列表
func (c *CodeArtsBuildClient) ListRelatedProjectInfoInvoker(request *model.ListRelatedProjectInfoRequest) *ListRelatedProjectInfoInvoker {
	requestDef := GenReqDefForListRelatedProjectInfo()
	return &ListRelatedProjectInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectPermission 获取用户权限
//
// 获取用户权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowProjectPermission(request *model.ShowProjectPermissionRequest) (*model.ShowProjectPermissionResponse, error) {
	requestDef := GenReqDefForShowProjectPermission()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectPermissionResponse), nil
	}
}

// ShowProjectPermissionInvoker 获取用户权限
func (c *CodeArtsBuildClient) ShowProjectPermissionInvoker(request *model.ShowProjectPermissionRequest) *ShowProjectPermissionInvoker {
	requestDef := GenReqDefForShowProjectPermission()
	return &ShowProjectPermissionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRelatedProject 获取当前用户的项目信息列表
//
// 获取当前用户的项目信息列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowRelatedProject(request *model.ShowRelatedProjectRequest) (*model.ShowRelatedProjectResponse, error) {
	requestDef := GenReqDefForShowRelatedProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRelatedProjectResponse), nil
	}
}

// ShowRelatedProjectInvoker 获取当前用户的项目信息列表
func (c *CodeArtsBuildClient) ShowRelatedProjectInvoker(request *model.ShowRelatedProjectRequest) *ShowRelatedProjectInvoker {
	requestDef := GenReqDefForShowRelatedProject()
	return &ShowRelatedProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSummaryBuildJobInfo 获取租户任务总数和成功率接口
//
// 获取租户任务总数和成功率接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowSummaryBuildJobInfo(request *model.ShowSummaryBuildJobInfoRequest) (*model.ShowSummaryBuildJobInfoResponse, error) {
	requestDef := GenReqDefForShowSummaryBuildJobInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSummaryBuildJobInfoResponse), nil
	}
}

// ShowSummaryBuildJobInfoInvoker 获取租户任务总数和成功率接口
func (c *CodeArtsBuildClient) ShowSummaryBuildJobInfoInvoker(request *model.ShowSummaryBuildJobInfoRequest) *ShowSummaryBuildJobInfoInvoker {
	requestDef := GenReqDefForShowSummaryBuildJobInfo()
	return &ShowSummaryBuildJobInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUserOverPackageQuota 当前用户所在项目所属租户的包周期每月时长是否超额
//
// 当前用户所在项目所属租户的包周期每月时长是否超额
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowUserOverPackageQuota(request *model.ShowUserOverPackageQuotaRequest) (*model.ShowUserOverPackageQuotaResponse, error) {
	requestDef := GenReqDefForShowUserOverPackageQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserOverPackageQuotaResponse), nil
	}
}

// ShowUserOverPackageQuotaInvoker 当前用户所在项目所属租户的包周期每月时长是否超额
func (c *CodeArtsBuildClient) ShowUserOverPackageQuotaInvoker(request *model.ShowUserOverPackageQuotaRequest) *ShowUserOverPackageQuotaInvoker {
	requestDef := GenReqDefForShowUserOverPackageQuota()
	return &ShowUserOverPackageQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDockerfileTemplate 获取dockerfileTemplate
//
// 获取dockerfileTemplate
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowDockerfileTemplate(request *model.ShowDockerfileTemplateRequest) (*model.ShowDockerfileTemplateResponse, error) {
	requestDef := GenReqDefForShowDockerfileTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDockerfileTemplateResponse), nil
	}
}

// ShowDockerfileTemplateInvoker 获取dockerfileTemplate
func (c *CodeArtsBuildClient) ShowDockerfileTemplateInvoker(request *model.ShowDockerfileTemplateRequest) *ShowDockerfileTemplateInvoker {
	requestDef := GenReqDefForShowDockerfileTemplate()
	return &ShowDockerfileTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowImageTemplateList 获取镜像模板列表
//
// 获取镜像模板列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowImageTemplateList(request *model.ShowImageTemplateListRequest) (*model.ShowImageTemplateListResponse, error) {
	requestDef := GenReqDefForShowImageTemplateList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowImageTemplateListResponse), nil
	}
}

// ShowImageTemplateListInvoker 获取镜像模板列表
func (c *CodeArtsBuildClient) ShowImageTemplateListInvoker(request *model.ShowImageTemplateListRequest) *ShowImageTemplateListInvoker {
	requestDef := GenReqDefForShowImageTemplateList()
	return &ShowImageTemplateListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckJobCountIsTopLimit 检查任务数量是否上限
//
// 检查任务数量是否上限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) CheckJobCountIsTopLimit(request *model.CheckJobCountIsTopLimitRequest) (*model.CheckJobCountIsTopLimitResponse, error) {
	requestDef := GenReqDefForCheckJobCountIsTopLimit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckJobCountIsTopLimitResponse), nil
	}
}

// CheckJobCountIsTopLimitInvoker 检查任务数量是否上限
func (c *CodeArtsBuildClient) CheckJobCountIsTopLimitInvoker(request *model.CheckJobCountIsTopLimitRequest) *CheckJobCountIsTopLimitInvoker {
	requestDef := GenReqDefForCheckJobCountIsTopLimit()
	return &CheckJobCountIsTopLimitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckJobNameIsExists 查看项目下任务名是否存在
//
// 查看项目下任务名是否存在
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) CheckJobNameIsExists(request *model.CheckJobNameIsExistsRequest) (*model.CheckJobNameIsExistsResponse, error) {
	requestDef := GenReqDefForCheckJobNameIsExists()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckJobNameIsExistsResponse), nil
	}
}

// CheckJobNameIsExistsInvoker 查看项目下任务名是否存在
func (c *CodeArtsBuildClient) CheckJobNameIsExistsInvoker(request *model.CheckJobNameIsExistsRequest) *CheckJobNameIsExistsInvoker {
	requestDef := GenReqDefForCheckJobNameIsExists()
	return &CheckJobNameIsExistsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckWebhookUrl 检查webhook地址参数
//
// 检查webhook地址参数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) CheckWebhookUrl(request *model.CheckWebhookUrlRequest) (*model.CheckWebhookUrlResponse, error) {
	requestDef := GenReqDefForCheckWebhookUrl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckWebhookUrlResponse), nil
	}
}

// CheckWebhookUrlInvoker 检查webhook地址参数
func (c *CodeArtsBuildClient) CheckWebhookUrlInvoker(request *model.CheckWebhookUrlRequest) *CheckWebhookUrlInvoker {
	requestDef := GenReqDefForCheckWebhookUrl()
	return &CheckWebhookUrlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ClearRecyclingJobs 清空回收站中的任务
//
// 清空回收站中的任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ClearRecyclingJobs(request *model.ClearRecyclingJobsRequest) (*model.ClearRecyclingJobsResponse, error) {
	requestDef := GenReqDefForClearRecyclingJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ClearRecyclingJobsResponse), nil
	}
}

// ClearRecyclingJobsInvoker 清空回收站中的任务
func (c *CodeArtsBuildClient) ClearRecyclingJobsInvoker(request *model.ClearRecyclingJobsRequest) *ClearRecyclingJobsInvoker {
	requestDef := GenReqDefForClearRecyclingJobs()
	return &ClearRecyclingJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CopyJob 复制构建任务
//
// 复制构建任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) CopyJob(request *model.CopyJobRequest) (*model.CopyJobResponse, error) {
	requestDef := GenReqDefForCopyJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyJobResponse), nil
	}
}

// CopyJobInvoker 复制构建任务
func (c *CodeArtsBuildClient) CopyJobInvoker(request *model.CopyJobRequest) *CopyJobInvoker {
	requestDef := GenReqDefForCopyJob()
	return &CopyJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateNewJob 创建构建任务
//
// 创建构建任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) CreateNewJob(request *model.CreateNewJobRequest) (*model.CreateNewJobResponse, error) {
	requestDef := GenReqDefForCreateNewJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNewJobResponse), nil
	}
}

// CreateNewJobInvoker 创建构建任务
func (c *CodeArtsBuildClient) CreateNewJobInvoker(request *model.CreateNewJobRequest) *CreateNewJobInvoker {
	requestDef := GenReqDefForCreateNewJob()
	return &CreateNewJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRecyclingJobs 删除回收站中的任务
//
// 删除回收站中的任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) DeleteRecyclingJobs(request *model.DeleteRecyclingJobsRequest) (*model.DeleteRecyclingJobsResponse, error) {
	requestDef := GenReqDefForDeleteRecyclingJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRecyclingJobsResponse), nil
	}
}

// DeleteRecyclingJobsInvoker 删除回收站中的任务
func (c *CodeArtsBuildClient) DeleteRecyclingJobsInvoker(request *model.DeleteRecyclingJobsRequest) *DeleteRecyclingJobsInvoker {
	requestDef := GenReqDefForDeleteRecyclingJobs()
	return &DeleteRecyclingJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTheJob 删除任务
//
// 删除任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) DeleteTheJob(request *model.DeleteTheJobRequest) (*model.DeleteTheJobResponse, error) {
	requestDef := GenReqDefForDeleteTheJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTheJobResponse), nil
	}
}

// DeleteTheJobInvoker 删除任务
func (c *CodeArtsBuildClient) DeleteTheJobInvoker(request *model.DeleteTheJobRequest) *DeleteTheJobInvoker {
	requestDef := GenReqDefForDeleteTheJob()
	return &DeleteTheJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableTheJob 禁用任务
//
// 禁用任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) DisableTheJob(request *model.DisableTheJobRequest) (*model.DisableTheJobResponse, error) {
	requestDef := GenReqDefForDisableTheJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableTheJobResponse), nil
	}
}

// DisableTheJobInvoker 禁用任务
func (c *CodeArtsBuildClient) DisableTheJobInvoker(request *model.DisableTheJobRequest) *DisableTheJobInvoker {
	requestDef := GenReqDefForDisableTheJob()
	return &DisableTheJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteJob 执行构建
//
// 执行构建任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ExecuteJob(request *model.ExecuteJobRequest) (*model.ExecuteJobResponse, error) {
	requestDef := GenReqDefForExecuteJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteJobResponse), nil
	}
}

// ExecuteJobInvoker 执行构建
func (c *CodeArtsBuildClient) ExecuteJobInvoker(request *model.ExecuteJobRequest) *ExecuteJobInvoker {
	requestDef := GenReqDefForExecuteJob()
	return &ExecuteJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBuildParameter 详情页获取构建参数
//
// 详情页获取构建参数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ListBuildParameter(request *model.ListBuildParameterRequest) (*model.ListBuildParameterResponse, error) {
	requestDef := GenReqDefForListBuildParameter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBuildParameterResponse), nil
	}
}

// ListBuildParameterInvoker 详情页获取构建参数
func (c *CodeArtsBuildClient) ListBuildParameterInvoker(request *model.ListBuildParameterRequest) *ListBuildParameterInvoker {
	requestDef := GenReqDefForListBuildParameter()
	return &ListBuildParameterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListJob 查看用户全部的构建任务列表
//
// 查看用户全部的构建任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ListJob(request *model.ListJobRequest) (*model.ListJobResponse, error) {
	requestDef := GenReqDefForListJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJobResponse), nil
	}
}

// ListJobInvoker 查看用户全部的构建任务列表
func (c *CodeArtsBuildClient) ListJobInvoker(request *model.ListJobRequest) *ListJobInvoker {
	requestDef := GenReqDefForListJob()
	return &ListJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectJobs 查询项目任务列表
//
// 查询项目任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ListProjectJobs(request *model.ListProjectJobsRequest) (*model.ListProjectJobsResponse, error) {
	requestDef := GenReqDefForListProjectJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectJobsResponse), nil
	}
}

// ListProjectJobsInvoker 查询项目任务列表
func (c *CodeArtsBuildClient) ListProjectJobsInvoker(request *model.ListProjectJobsRequest) *ListProjectJobsInvoker {
	requestDef := GenReqDefForListProjectJobs()
	return &ListProjectJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRecyclingJob 查看回收站中删除的构建任务列表
//
// 查看回收站中删除的构建任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ListRecyclingJob(request *model.ListRecyclingJobRequest) (*model.ListRecyclingJobResponse, error) {
	requestDef := GenReqDefForListRecyclingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecyclingJobResponse), nil
	}
}

// ListRecyclingJobInvoker 查看回收站中删除的构建任务列表
func (c *CodeArtsBuildClient) ListRecyclingJobInvoker(request *model.ListRecyclingJobRequest) *ListRecyclingJobInvoker {
	requestDef := GenReqDefForListRecyclingJob()
	return &ListRecyclingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUpdateJobHistory 获取修改历史
//
// 获取修改历史
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ListUpdateJobHistory(request *model.ListUpdateJobHistoryRequest) (*model.ListUpdateJobHistoryResponse, error) {
	requestDef := GenReqDefForListUpdateJobHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUpdateJobHistoryResponse), nil
	}
}

// ListUpdateJobHistoryInvoker 获取修改历史
func (c *CodeArtsBuildClient) ListUpdateJobHistoryInvoker(request *model.ListUpdateJobHistoryRequest) *ListUpdateJobHistoryInvoker {
	requestDef := GenReqDefForListUpdateJobHistory()
	return &ListUpdateJobHistoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestoreRecyclingJobs 恢复回收站中的任务
//
// 恢复回收站中的任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) RestoreRecyclingJobs(request *model.RestoreRecyclingJobsRequest) (*model.RestoreRecyclingJobsResponse, error) {
	requestDef := GenReqDefForRestoreRecyclingJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreRecyclingJobsResponse), nil
	}
}

// RestoreRecyclingJobsInvoker 恢复回收站中的任务
func (c *CodeArtsBuildClient) RestoreRecyclingJobsInvoker(request *model.RestoreRecyclingJobsRequest) *RestoreRecyclingJobsInvoker {
	requestDef := GenReqDefForRestoreRecyclingJobs()
	return &RestoreRecyclingJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetKeepTime 设置回收站中的任务保留时间
//
// 设置回收站中的任务保留时间,该接口需要租户账号才能访问，租户子账号无权限访问。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) SetKeepTime(request *model.SetKeepTimeRequest) (*model.SetKeepTimeResponse, error) {
	requestDef := GenReqDefForSetKeepTime()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetKeepTimeResponse), nil
	}
}

// SetKeepTimeInvoker 设置回收站中的任务保留时间
func (c *CodeArtsBuildClient) SetKeepTimeInvoker(request *model.SetKeepTimeRequest) *SetKeepTimeInvoker {
	requestDef := GenReqDefForSetKeepTime()
	return &SetKeepTimeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBuildParamsList 编辑页获取参数类型的接口
//
// 编辑页获取参数类型的接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowBuildParamsList(request *model.ShowBuildParamsListRequest) (*model.ShowBuildParamsListResponse, error) {
	requestDef := GenReqDefForShowBuildParamsList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBuildParamsListResponse), nil
	}
}

// ShowBuildParamsListInvoker 编辑页获取参数类型的接口
func (c *CodeArtsBuildClient) ShowBuildParamsListInvoker(request *model.ShowBuildParamsListRequest) *ShowBuildParamsListInvoker {
	requestDef := GenReqDefForShowBuildParamsList()
	return &ShowBuildParamsListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCopyName 复制任务名
//
// 复制任务名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowCopyName(request *model.ShowCopyNameRequest) (*model.ShowCopyNameResponse, error) {
	requestDef := GenReqDefForShowCopyName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCopyNameResponse), nil
	}
}

// ShowCopyNameInvoker 复制任务名
func (c *CodeArtsBuildClient) ShowCopyNameInvoker(request *model.ShowCopyNameRequest) *ShowCopyNameInvoker {
	requestDef := GenReqDefForShowCopyName()
	return &ShowCopyNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDefaultBuildParameters 获取编译构建默认参数
//
// 获取编译构建默认参数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowDefaultBuildParameters(request *model.ShowDefaultBuildParametersRequest) (*model.ShowDefaultBuildParametersResponse, error) {
	requestDef := GenReqDefForShowDefaultBuildParameters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDefaultBuildParametersResponse), nil
	}
}

// ShowDefaultBuildParametersInvoker 获取编译构建默认参数
func (c *CodeArtsBuildClient) ShowDefaultBuildParametersInvoker(request *model.ShowDefaultBuildParametersRequest) *ShowDefaultBuildParametersInvoker {
	requestDef := GenReqDefForShowDefaultBuildParameters()
	return &ShowDefaultBuildParametersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDefaultProjectPermission 获取当前项目默认角色权限矩阵信息
//
// 获取当前项目默认角色权限矩阵信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowDefaultProjectPermission(request *model.ShowDefaultProjectPermissionRequest) (*model.ShowDefaultProjectPermissionResponse, error) {
	requestDef := GenReqDefForShowDefaultProjectPermission()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDefaultProjectPermissionResponse), nil
	}
}

// ShowDefaultProjectPermissionInvoker 获取当前项目默认角色权限矩阵信息
func (c *CodeArtsBuildClient) ShowDefaultProjectPermissionInvoker(request *model.ShowDefaultProjectPermissionRequest) *ShowDefaultProjectPermissionInvoker {
	requestDef := GenReqDefForShowDefaultProjectPermission()
	return &ShowDefaultProjectPermissionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDisable 查询任务是否已禁用
//
// 查询任务是否已禁用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowDisable(request *model.ShowDisableRequest) (*model.ShowDisableResponse, error) {
	requestDef := GenReqDefForShowDisable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDisableResponse), nil
	}
}

// ShowDisableInvoker 查询任务是否已禁用
func (c *CodeArtsBuildClient) ShowDisableInvoker(request *model.ShowDisableRequest) *ShowDisableInvoker {
	requestDef := GenReqDefForShowDisable()
	return &ShowDisableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobConfig 获取构建任务详情
//
// 获取构建任务详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowJobConfig(request *model.ShowJobConfigRequest) (*model.ShowJobConfigResponse, error) {
	requestDef := GenReqDefForShowJobConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobConfigResponse), nil
	}
}

// ShowJobConfigInvoker 获取构建任务详情
func (c *CodeArtsBuildClient) ShowJobConfigInvoker(request *model.ShowJobConfigRequest) *ShowJobConfigInvoker {
	requestDef := GenReqDefForShowJobConfig()
	return &ShowJobConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobConfigDiff 获取构建任务配置的对比差异
//
// 获取构建任务配置的对比差异
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowJobConfigDiff(request *model.ShowJobConfigDiffRequest) (*model.ShowJobConfigDiffResponse, error) {
	requestDef := GenReqDefForShowJobConfigDiff()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobConfigDiffResponse), nil
	}
}

// ShowJobConfigDiffInvoker 获取构建任务配置的对比差异
func (c *CodeArtsBuildClient) ShowJobConfigDiffInvoker(request *model.ShowJobConfigDiffRequest) *ShowJobConfigDiffInvoker {
	requestDef := GenReqDefForShowJobConfigDiff()
	return &ShowJobConfigDiffInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobInfo 查看构建任务构建信息
//
// 查看构建任务构建信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowJobInfo(request *model.ShowJobInfoRequest) (*model.ShowJobInfoResponse, error) {
	requestDef := GenReqDefForShowJobInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobInfoResponse), nil
	}
}

// ShowJobInfoInvoker 查看构建任务构建信息
func (c *CodeArtsBuildClient) ShowJobInfoInvoker(request *model.ShowJobInfoRequest) *ShowJobInfoInvoker {
	requestDef := GenReqDefForShowJobInfo()
	return &ShowJobInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobNoticeConfigInfo 获取通知信息
//
// 获取通知信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowJobNoticeConfigInfo(request *model.ShowJobNoticeConfigInfoRequest) (*model.ShowJobNoticeConfigInfoResponse, error) {
	requestDef := GenReqDefForShowJobNoticeConfigInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobNoticeConfigInfoResponse), nil
	}
}

// ShowJobNoticeConfigInfoInvoker 获取通知信息
func (c *CodeArtsBuildClient) ShowJobNoticeConfigInfoInvoker(request *model.ShowJobNoticeConfigInfoRequest) *ShowJobNoticeConfigInfoInvoker {
	requestDef := GenReqDefForShowJobNoticeConfigInfo()
	return &ShowJobNoticeConfigInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobRolePermission 获取构建任务的角色权限矩阵信息
//
// 获取构建任务的角色权限矩阵信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowJobRolePermission(request *model.ShowJobRolePermissionRequest) (*model.ShowJobRolePermissionResponse, error) {
	requestDef := GenReqDefForShowJobRolePermission()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobRolePermissionResponse), nil
	}
}

// ShowJobRolePermissionInvoker 获取构建任务的角色权限矩阵信息
func (c *CodeArtsBuildClient) ShowJobRolePermissionInvoker(request *model.ShowJobRolePermissionRequest) *ShowJobRolePermissionInvoker {
	requestDef := GenReqDefForShowJobRolePermission()
	return &ShowJobRolePermissionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobStepStatus 查询任务状态
//
// 查询任务状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowJobStepStatus(request *model.ShowJobStepStatusRequest) (*model.ShowJobStepStatusResponse, error) {
	requestDef := GenReqDefForShowJobStepStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobStepStatusResponse), nil
	}
}

// ShowJobStepStatusInvoker 查询任务状态
func (c *CodeArtsBuildClient) ShowJobStepStatusInvoker(request *model.ShowJobStepStatusRequest) *ShowJobStepStatusInvoker {
	requestDef := GenReqDefForShowJobStepStatus()
	return &ShowJobStepStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobSystemParameters 查看系统预定义参数
//
// 查看系统预定义参数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowJobSystemParameters(request *model.ShowJobSystemParametersRequest) (*model.ShowJobSystemParametersResponse, error) {
	requestDef := GenReqDefForShowJobSystemParameters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobSystemParametersResponse), nil
	}
}

// ShowJobSystemParametersInvoker 查看系统预定义参数
func (c *CodeArtsBuildClient) ShowJobSystemParametersInvoker(request *model.ShowJobSystemParametersRequest) *ShowJobSystemParametersInvoker {
	requestDef := GenReqDefForShowJobSystemParameters()
	return &ShowJobSystemParametersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectJobPermission 获取任务权限矩阵
//
// 获取任务权限矩阵
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowProjectJobPermission(request *model.ShowProjectJobPermissionRequest) (*model.ShowProjectJobPermissionResponse, error) {
	requestDef := GenReqDefForShowProjectJobPermission()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectJobPermissionResponse), nil
	}
}

// ShowProjectJobPermissionInvoker 获取任务权限矩阵
func (c *CodeArtsBuildClient) ShowProjectJobPermissionInvoker(request *model.ShowProjectJobPermissionRequest) *ShowProjectJobPermissionInvoker {
	requestDef := GenReqDefForShowProjectJobPermission()
	return &ShowProjectJobPermissionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRunningStatus 查看任务是否在构建
//
// 查看任务是否在构建
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowRunningStatus(request *model.ShowRunningStatusRequest) (*model.ShowRunningStatusResponse, error) {
	requestDef := GenReqDefForShowRunningStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRunningStatusResponse), nil
	}
}

// ShowRunningStatusInvoker 查看任务是否在构建
func (c *CodeArtsBuildClient) ShowRunningStatusInvoker(request *model.ShowRunningStatusRequest) *ShowRunningStatusInvoker {
	requestDef := GenReqDefForShowRunningStatus()
	return &ShowRunningStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNewJob 更新构建任务
//
// 更新构建任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) UpdateNewJob(request *model.UpdateNewJobRequest) (*model.UpdateNewJobResponse, error) {
	requestDef := GenReqDefForUpdateNewJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNewJobResponse), nil
	}
}

// UpdateNewJobInvoker 更新构建任务
func (c *CodeArtsBuildClient) UpdateNewJobInvoker(request *model.UpdateNewJobRequest) *UpdateNewJobInvoker {
	requestDef := GenReqDefForUpdateNewJob()
	return &UpdateNewJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteKeystore 删除文件管理文件
//
// 删除文件管理文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) DeleteKeystore(request *model.DeleteKeystoreRequest) (*model.DeleteKeystoreResponse, error) {
	requestDef := GenReqDefForDeleteKeystore()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteKeystoreResponse), nil
	}
}

// DeleteKeystoreInvoker 删除文件管理文件
func (c *CodeArtsBuildClient) DeleteKeystoreInvoker(request *model.DeleteKeystoreRequest) *DeleteKeystoreInvoker {
	requestDef := GenReqDefForDeleteKeystore()
	return &DeleteKeystoreInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteKeystorePermission 文件管理删除权限
//
// 文件管理删除权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) DeleteKeystorePermission(request *model.DeleteKeystorePermissionRequest) (*model.DeleteKeystorePermissionResponse, error) {
	requestDef := GenReqDefForDeleteKeystorePermission()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteKeystorePermissionResponse), nil
	}
}

// DeleteKeystorePermissionInvoker 文件管理删除权限
func (c *CodeArtsBuildClient) DeleteKeystorePermissionInvoker(request *model.DeleteKeystorePermissionRequest) *DeleteKeystorePermissionInvoker {
	requestDef := GenReqDefForDeleteKeystorePermission()
	return &DeleteKeystorePermissionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadKeystoreByName 文件管理文件下载
//
// 文件管理文件下载
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) DownloadKeystoreByName(request *model.DownloadKeystoreByNameRequest) (*model.DownloadKeystoreByNameResponse, error) {
	requestDef := GenReqDefForDownloadKeystoreByName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadKeystoreByNameResponse), nil
	}
}

// DownloadKeystoreByNameInvoker 文件管理文件下载
func (c *CodeArtsBuildClient) DownloadKeystoreByNameInvoker(request *model.DownloadKeystoreByNameRequest) *DownloadKeystoreByNameInvoker {
	requestDef := GenReqDefForDownloadKeystoreByName()
	return &DownloadKeystoreByNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListKeystore 查询用户可使用文件
//
// 查询用户可使用文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ListKeystore(request *model.ListKeystoreRequest) (*model.ListKeystoreResponse, error) {
	requestDef := GenReqDefForListKeystore()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKeystoreResponse), nil
	}
}

// ListKeystoreInvoker 查询用户可使用文件
func (c *CodeArtsBuildClient) ListKeystoreInvoker(request *model.ListKeystoreRequest) *ListKeystoreInvoker {
	requestDef := GenReqDefForListKeystore()
	return &ListKeystoreInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListKeystoreSearch 查询租户下文件列表
//
// 查询租户下文件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ListKeystoreSearch(request *model.ListKeystoreSearchRequest) (*model.ListKeystoreSearchResponse, error) {
	requestDef := GenReqDefForListKeystoreSearch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKeystoreSearchResponse), nil
	}
}

// ListKeystoreSearchInvoker 查询租户下文件列表
func (c *CodeArtsBuildClient) ListKeystoreSearchInvoker(request *model.ListKeystoreSearchRequest) *ListKeystoreSearchInvoker {
	requestDef := GenReqDefForListKeystoreSearch()
	return &ListKeystoreSearchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowKeystorePermission 文件管理查询权限
//
// 文件管理查询权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowKeystorePermission(request *model.ShowKeystorePermissionRequest) (*model.ShowKeystorePermissionResponse, error) {
	requestDef := GenReqDefForShowKeystorePermission()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKeystorePermissionResponse), nil
	}
}

// ShowKeystorePermissionInvoker 文件管理查询权限
func (c *CodeArtsBuildClient) ShowKeystorePermissionInvoker(request *model.ShowKeystorePermissionRequest) *ShowKeystorePermissionInvoker {
	requestDef := GenReqDefForShowKeystorePermission()
	return &ShowKeystorePermissionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateKeystore 更新文件信息
//
// 更新文件信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) UpdateKeystore(request *model.UpdateKeystoreRequest) (*model.UpdateKeystoreResponse, error) {
	requestDef := GenReqDefForUpdateKeystore()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateKeystoreResponse), nil
	}
}

// UpdateKeystoreInvoker 更新文件信息
func (c *CodeArtsBuildClient) UpdateKeystoreInvoker(request *model.UpdateKeystoreRequest) *UpdateKeystoreInvoker {
	requestDef := GenReqDefForUpdateKeystore()
	return &UpdateKeystoreInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadKeystore 上传文件
//
// 上传文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) UploadKeystore(request *model.UploadKeystoreRequest) (*model.UploadKeystoreResponse, error) {
	requestDef := GenReqDefForUploadKeystore()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadKeystoreResponse), nil
	}
}

// UploadKeystoreInvoker 上传文件
func (c *CodeArtsBuildClient) UploadKeystoreInvoker(request *model.UploadKeystoreRequest) *UploadKeystoreInvoker {
	requestDef := GenReqDefForUploadKeystore()
	return &UploadKeystoreInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadLogByRecordId 下载构建日志(待下线)
//
// 下载构建日志(待下线)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) DownloadLogByRecordId(request *model.DownloadLogByRecordIdRequest) (*model.DownloadLogByRecordIdResponse, error) {
	requestDef := GenReqDefForDownloadLogByRecordId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadLogByRecordIdResponse), nil
	}
}

// DownloadLogByRecordIdInvoker 下载构建日志(待下线)
func (c *CodeArtsBuildClient) DownloadLogByRecordIdInvoker(request *model.DownloadLogByRecordIdRequest) *DownloadLogByRecordIdInvoker {
	requestDef := GenReqDefForDownloadLogByRecordId()
	return &DownloadLogByRecordIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFlowGraph 获取构建记录的有向无环图(待下线)
//
// 获取构建记录的有向无环图(待下线)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowFlowGraph(request *model.ShowFlowGraphRequest) (*model.ShowFlowGraphResponse, error) {
	requestDef := GenReqDefForShowFlowGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFlowGraphResponse), nil
	}
}

// ShowFlowGraphInvoker 获取构建记录的有向无环图(待下线)
func (c *CodeArtsBuildClient) ShowFlowGraphInvoker(request *model.ShowFlowGraphRequest) *ShowFlowGraphInvoker {
	requestDef := GenReqDefForShowFlowGraph()
	return &ShowFlowGraphInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRecordInfo 获取构建记录信息(待下线)
//
// 获取构建记录信息(待下线)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowRecordInfo(request *model.ShowRecordInfoRequest) (*model.ShowRecordInfoResponse, error) {
	requestDef := GenReqDefForShowRecordInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecordInfoResponse), nil
	}
}

// ShowRecordInfoInvoker 获取构建记录信息(待下线)
func (c *CodeArtsBuildClient) ShowRecordInfoInvoker(request *model.ShowRecordInfoRequest) *ShowRecordInfoInvoker {
	requestDef := GenReqDefForShowRecordInfo()
	return &ShowRecordInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopJob 停止构建任务(待下线)
//
// 停止构建任务(待下线)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) StopJob(request *model.StopJobRequest) (*model.StopJobResponse, error) {
	requestDef := GenReqDefForStopJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopJobResponse), nil
	}
}

// StopJobInvoker 停止构建任务(待下线)
func (c *CodeArtsBuildClient) StopJobInvoker(request *model.StopJobRequest) *StopJobInvoker {
	requestDef := GenReqDefForStopJob()
	return &StopJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBriefRecord 获取指定工程的简要构建信息
//
// 获取指定工程的简要构建信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ListBriefRecord(request *model.ListBriefRecordRequest) (*model.ListBriefRecordResponse, error) {
	requestDef := GenReqDefForListBriefRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBriefRecordResponse), nil
	}
}

// ListBriefRecordInvoker 获取指定工程的简要构建信息
func (c *CodeArtsBuildClient) ListBriefRecordInvoker(request *model.ListBriefRecordRequest) *ListBriefRecordInvoker {
	requestDef := GenReqDefForListBriefRecord()
	return &ListBriefRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBuildInfoRecordByJobId 获取任务构建记录列表v1
//
// 获取任务构建记录列表v1
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ListBuildInfoRecordByJobId(request *model.ListBuildInfoRecordByJobIdRequest) (*model.ListBuildInfoRecordByJobIdResponse, error) {
	requestDef := GenReqDefForListBuildInfoRecordByJobId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBuildInfoRecordByJobIdResponse), nil
	}
}

// ListBuildInfoRecordByJobIdInvoker 获取任务构建记录列表v1
func (c *CodeArtsBuildClient) ListBuildInfoRecordByJobIdInvoker(request *model.ListBuildInfoRecordByJobIdRequest) *ListBuildInfoRecordByJobIdInvoker {
	requestDef := GenReqDefForListBuildInfoRecordByJobId()
	return &ListBuildInfoRecordByJobIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRecords 获取指定工程的构建记录列表
//
// 获取指定工程的构建记录列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ListRecords(request *model.ListRecordsRequest) (*model.ListRecordsResponse, error) {
	requestDef := GenReqDefForListRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecordsResponse), nil
	}
}

// ListRecordsInvoker 获取指定工程的构建记录列表
func (c *CodeArtsBuildClient) ListRecordsInvoker(request *model.ListRecordsRequest) *ListRecordsInvoker {
	requestDef := GenReqDefForListRecords()
	return &ListRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBuildInfoRecord 获取任务构建记录列表
//
// 获取任务构建记录列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowBuildInfoRecord(request *model.ShowBuildInfoRecordRequest) (*model.ShowBuildInfoRecordResponse, error) {
	requestDef := GenReqDefForShowBuildInfoRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBuildInfoRecordResponse), nil
	}
}

// ShowBuildInfoRecordInvoker 获取任务构建记录列表
func (c *CodeArtsBuildClient) ShowBuildInfoRecordInvoker(request *model.ShowBuildInfoRecordRequest) *ShowBuildInfoRecordInvoker {
	requestDef := GenReqDefForShowBuildInfoRecord()
	return &ShowBuildInfoRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBuildRecord 查询指定构建记录详情
//
// 查询指定构建记录详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowBuildRecord(request *model.ShowBuildRecordRequest) (*model.ShowBuildRecordResponse, error) {
	requestDef := GenReqDefForShowBuildRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBuildRecordResponse), nil
	}
}

// ShowBuildRecordInvoker 查询指定构建记录详情
func (c *CodeArtsBuildClient) ShowBuildRecordInvoker(request *model.ShowBuildRecordRequest) *ShowBuildRecordInvoker {
	requestDef := GenReqDefForShowBuildRecord()
	return &ShowBuildRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBuildRecordBuildScript 获取构建记录的构建脚本
//
// 获取构建记录的构建脚本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowBuildRecordBuildScript(request *model.ShowBuildRecordBuildScriptRequest) (*model.ShowBuildRecordBuildScriptResponse, error) {
	requestDef := GenReqDefForShowBuildRecordBuildScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBuildRecordBuildScriptResponse), nil
	}
}

// ShowBuildRecordBuildScriptInvoker 获取构建记录的构建脚本
func (c *CodeArtsBuildClient) ShowBuildRecordBuildScriptInvoker(request *model.ShowBuildRecordBuildScriptRequest) *ShowBuildRecordBuildScriptInvoker {
	requestDef := GenReqDefForShowBuildRecordBuildScript()
	return &ShowBuildRecordBuildScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBuildRecordFlowGraph 获取构建记录的有向无环图
//
// 获取构建记录的有向无环图
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowBuildRecordFlowGraph(request *model.ShowBuildRecordFlowGraphRequest) (*model.ShowBuildRecordFlowGraphResponse, error) {
	requestDef := GenReqDefForShowBuildRecordFlowGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBuildRecordFlowGraphResponse), nil
	}
}

// ShowBuildRecordFlowGraphInvoker 获取构建记录的有向无环图
func (c *CodeArtsBuildClient) ShowBuildRecordFlowGraphInvoker(request *model.ShowBuildRecordFlowGraphRequest) *ShowBuildRecordFlowGraphInvoker {
	requestDef := GenReqDefForShowBuildRecordFlowGraph()
	return &ShowBuildRecordFlowGraphInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBuildRecordFullStages 获取任务各阶段信息
//
// 获取任务各阶段信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowBuildRecordFullStages(request *model.ShowBuildRecordFullStagesRequest) (*model.ShowBuildRecordFullStagesResponse, error) {
	requestDef := GenReqDefForShowBuildRecordFullStages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBuildRecordFullStagesResponse), nil
	}
}

// ShowBuildRecordFullStagesInvoker 获取任务各阶段信息
func (c *CodeArtsBuildClient) ShowBuildRecordFullStagesInvoker(request *model.ShowBuildRecordFullStagesRequest) *ShowBuildRecordFullStagesInvoker {
	requestDef := GenReqDefForShowBuildRecordFullStages()
	return &ShowBuildRecordFullStagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobBuildRecordDetail 获取构建记录信息
//
// 获取构建记录信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowJobBuildRecordDetail(request *model.ShowJobBuildRecordDetailRequest) (*model.ShowJobBuildRecordDetailResponse, error) {
	requestDef := GenReqDefForShowJobBuildRecordDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobBuildRecordDetailResponse), nil
	}
}

// ShowJobBuildRecordDetailInvoker 获取构建记录信息
func (c *CodeArtsBuildClient) ShowJobBuildRecordDetailInvoker(request *model.ShowJobBuildRecordDetailRequest) *ShowJobBuildRecordDetailInvoker {
	requestDef := GenReqDefForShowJobBuildRecordDetail()
	return &ShowJobBuildRecordDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobTotal 构建历史页获取构建次数
//
// 构建历史页获取构建次数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowJobTotal(request *model.ShowJobTotalRequest) (*model.ShowJobTotalResponse, error) {
	requestDef := GenReqDefForShowJobTotal()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobTotalResponse), nil
	}
}

// ShowJobTotalInvoker 构建历史页获取构建次数
func (c *CodeArtsBuildClient) ShowJobTotalInvoker(request *model.ShowJobTotalRequest) *ShowJobTotalInvoker {
	requestDef := GenReqDefForShowJobTotal()
	return &ShowJobTotalInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadJunitCoverageZip 获取单元测试覆盖率报告压缩包
//
// 获取单元测试覆盖率报告压缩包
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) DownloadJunitCoverageZip(request *model.DownloadJunitCoverageZipRequest) (*model.DownloadJunitCoverageZipResponse, error) {
	requestDef := GenReqDefForDownloadJunitCoverageZip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadJunitCoverageZipResponse), nil
	}
}

// DownloadJunitCoverageZipInvoker 获取单元测试覆盖率报告压缩包
func (c *CodeArtsBuildClient) DownloadJunitCoverageZipInvoker(request *model.DownloadJunitCoverageZipRequest) *DownloadJunitCoverageZipInvoker {
	requestDef := GenReqDefForDownloadJunitCoverageZip()
	return &DownloadJunitCoverageZipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListJunitCoverageSummary 获取单元测试覆盖率报告列表
//
// 获取单元测试覆盖率报告列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ListJunitCoverageSummary(request *model.ListJunitCoverageSummaryRequest) (*model.ListJunitCoverageSummaryResponse, error) {
	requestDef := GenReqDefForListJunitCoverageSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJunitCoverageSummaryResponse), nil
	}
}

// ListJunitCoverageSummaryInvoker 获取单元测试覆盖率报告列表
func (c *CodeArtsBuildClient) ListJunitCoverageSummaryInvoker(request *model.ListJunitCoverageSummaryRequest) *ListJunitCoverageSummaryInvoker {
	requestDef := GenReqDefForListJunitCoverageSummary()
	return &ListJunitCoverageSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepoBranch 获取该任务所有分支信息
//
// 获取该任务所有分支信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ListRepoBranch(request *model.ListRepoBranchRequest) (*model.ListRepoBranchResponse, error) {
	requestDef := GenReqDefForListRepoBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepoBranchResponse), nil
	}
}

// ListRepoBranchInvoker 获取该任务所有分支信息
func (c *CodeArtsBuildClient) ListRepoBranchInvoker(request *model.ListRepoBranchRequest) *ListRepoBranchInvoker {
	requestDef := GenReqDefForListRepoBranch()
	return &ListRepoBranchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepository 查看仓库
//
// 查看仓库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ListRepository(request *model.ListRepositoryRequest) (*model.ListRepositoryResponse, error) {
	requestDef := GenReqDefForListRepository()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryResponse), nil
	}
}

// ListRepositoryInvoker 查看仓库
func (c *CodeArtsBuildClient) ListRepositoryInvoker(request *model.ListRepositoryRequest) *ListRepositoryInvoker {
	requestDef := GenReqDefForListRepository()
	return &ListRepositoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobBuildSuccessRatio 查询构建成功率
//
// 查询构建成功率
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowJobBuildSuccessRatio(request *model.ShowJobBuildSuccessRatioRequest) (*model.ShowJobBuildSuccessRatioResponse, error) {
	requestDef := GenReqDefForShowJobBuildSuccessRatio()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobBuildSuccessRatioResponse), nil
	}
}

// ShowJobBuildSuccessRatioInvoker 查询构建成功率
func (c *CodeArtsBuildClient) ShowJobBuildSuccessRatioInvoker(request *model.ShowJobBuildSuccessRatioRequest) *ShowJobBuildSuccessRatioInvoker {
	requestDef := GenReqDefForShowJobBuildSuccessRatio()
	return &ShowJobBuildSuccessRatioInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobBuildTime 洞察构建时长
//
// 洞察构建时长
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowJobBuildTime(request *model.ShowJobBuildTimeRequest) (*model.ShowJobBuildTimeResponse, error) {
	requestDef := GenReqDefForShowJobBuildTime()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobBuildTimeResponse), nil
	}
}

// ShowJobBuildTimeInvoker 洞察构建时长
func (c *CodeArtsBuildClient) ShowJobBuildTimeInvoker(request *model.ShowJobBuildTimeRequest) *ShowJobBuildTimeInvoker {
	requestDef := GenReqDefForShowJobBuildTime()
	return &ShowJobBuildTimeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowReportSummary 获取覆盖率接口
//
// 获取覆盖率接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowReportSummary(request *model.ShowReportSummaryRequest) (*model.ShowReportSummaryResponse, error) {
	requestDef := GenReqDefForShowReportSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReportSummaryResponse), nil
	}
}

// ShowReportSummaryInvoker 获取覆盖率接口
func (c *CodeArtsBuildClient) ShowReportSummaryInvoker(request *model.ShowReportSummaryRequest) *ShowReportSummaryInvoker {
	requestDef := GenReqDefForShowReportSummary()
	return &ShowReportSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTemplate 创建构建模板
//
// 创建构建模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) CreateTemplate(request *model.CreateTemplateRequest) (*model.CreateTemplateResponse, error) {
	requestDef := GenReqDefForCreateTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTemplateResponse), nil
	}
}

// CreateTemplateInvoker 创建构建模板
func (c *CodeArtsBuildClient) CreateTemplateInvoker(request *model.CreateTemplateRequest) *CreateTemplateInvoker {
	requestDef := GenReqDefForCreateTemplate()
	return &CreateTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTemplate 删除构建模板
//
// 删除构建模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) DeleteTemplate(request *model.DeleteTemplateRequest) (*model.DeleteTemplateResponse, error) {
	requestDef := GenReqDefForDeleteTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTemplateResponse), nil
	}
}

// DeleteTemplateInvoker 删除构建模板
func (c *CodeArtsBuildClient) DeleteTemplateInvoker(request *model.DeleteTemplateRequest) *DeleteTemplateInvoker {
	requestDef := GenReqDefForDeleteTemplate()
	return &DeleteTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCustomTemplate 根据条件查询特定模板
//
// 根据条件查询特定模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ListCustomTemplate(request *model.ListCustomTemplateRequest) (*model.ListCustomTemplateResponse, error) {
	requestDef := GenReqDefForListCustomTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomTemplateResponse), nil
	}
}

// ListCustomTemplateInvoker 根据条件查询特定模板
func (c *CodeArtsBuildClient) ListCustomTemplateInvoker(request *model.ListCustomTemplateRequest) *ListCustomTemplateInvoker {
	requestDef := GenReqDefForListCustomTemplate()
	return &ListCustomTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOfficialTemplate 查询官方模版
//
// 查询官方模版
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ListOfficialTemplate(request *model.ListOfficialTemplateRequest) (*model.ListOfficialTemplateResponse, error) {
	requestDef := GenReqDefForListOfficialTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOfficialTemplateResponse), nil
	}
}

// ListOfficialTemplateInvoker 查询官方模版
func (c *CodeArtsBuildClient) ListOfficialTemplateInvoker(request *model.ListOfficialTemplateRequest) *ListOfficialTemplateInvoker {
	requestDef := GenReqDefForListOfficialTemplate()
	return &ListOfficialTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRecommendOfficialTemplate 获取官方推荐模板
//
// 获取官方推荐模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ListRecommendOfficialTemplate(request *model.ListRecommendOfficialTemplateRequest) (*model.ListRecommendOfficialTemplateResponse, error) {
	requestDef := GenReqDefForListRecommendOfficialTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecommendOfficialTemplateResponse), nil
	}
}

// ListRecommendOfficialTemplateInvoker 获取官方推荐模板
func (c *CodeArtsBuildClient) ListRecommendOfficialTemplateInvoker(request *model.ListRecommendOfficialTemplateRequest) *ListRecommendOfficialTemplateInvoker {
	requestDef := GenReqDefForListRecommendOfficialTemplate()
	return &ListRecommendOfficialTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SaveTemplateUsedInfo 保存模板使用记录
//
// 保存模板使用记录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) SaveTemplateUsedInfo(request *model.SaveTemplateUsedInfoRequest) (*model.SaveTemplateUsedInfoResponse, error) {
	requestDef := GenReqDefForSaveTemplateUsedInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SaveTemplateUsedInfoResponse), nil
	}
}

// SaveTemplateUsedInfoInvoker 保存模板使用记录
func (c *CodeArtsBuildClient) SaveTemplateUsedInfoInvoker(request *model.SaveTemplateUsedInfoRequest) *SaveTemplateUsedInfoInvoker {
	requestDef := GenReqDefForSaveTemplateUsedInfo()
	return &SaveTemplateUsedInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowYamlTemplate 获取代码化构建默认模板
//
// 获取代码化构建默认模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ShowYamlTemplate(request *model.ShowYamlTemplateRequest) (*model.ShowYamlTemplateResponse, error) {
	requestDef := GenReqDefForShowYamlTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowYamlTemplateResponse), nil
	}
}

// ShowYamlTemplateInvoker 获取代码化构建默认模板
func (c *CodeArtsBuildClient) ShowYamlTemplateInvoker(request *model.ShowYamlTemplateRequest) *ShowYamlTemplateInvoker {
	requestDef := GenReqDefForShowYamlTemplate()
	return &ShowYamlTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
