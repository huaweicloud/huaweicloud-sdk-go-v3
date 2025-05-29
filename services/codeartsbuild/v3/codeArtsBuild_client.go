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
