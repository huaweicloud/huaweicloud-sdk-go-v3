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

// DownloadLogByRecordId 下载构建日志
//
// 下载构建日志
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

// DownloadLogByRecordIdInvoker 下载构建日志
func (c *CodeArtsBuildClient) DownloadLogByRecordIdInvoker(request *model.DownloadLogByRecordIdRequest) *DownloadLogByRecordIdInvoker {
	requestDef := GenReqDefForDownloadLogByRecordId()
	return &DownloadLogByRecordIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResumeBuildJob 恢复构建任务
//
// 恢复构建任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsBuildClient) ResumeBuildJob(request *model.ResumeBuildJobRequest) (*model.ResumeBuildJobResponse, error) {
	requestDef := GenReqDefForResumeBuildJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResumeBuildJobResponse), nil
	}
}

// ResumeBuildJobInvoker 恢复构建任务
func (c *CodeArtsBuildClient) ResumeBuildJobInvoker(request *model.ResumeBuildJobRequest) *ResumeBuildJobInvoker {
	requestDef := GenReqDefForResumeBuildJob()
	return &ResumeBuildJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunJob 执行构建任务
//
// 执行构建任务,可传自定义参数。
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

// ShowRecordInfo 获取构建记录信息
//
// 获取构建记录信息
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

// ShowRecordInfoInvoker 获取构建记录信息
func (c *CodeArtsBuildClient) ShowRecordInfoInvoker(request *model.ShowRecordInfoRequest) *ShowRecordInfoInvoker {
	requestDef := GenReqDefForShowRecordInfo()
	return &ShowRecordInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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
