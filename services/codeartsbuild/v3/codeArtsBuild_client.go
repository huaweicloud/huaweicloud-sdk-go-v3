package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codeartsbuild/v3/model"
)

type CodeArtsBuildClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCodeArtsBuildClient(hcClient *http_client.HcHttpClient) *CodeArtsBuildClient {
	return &CodeArtsBuildClient{HcClient: hcClient}
}

func CodeArtsBuildClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
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
