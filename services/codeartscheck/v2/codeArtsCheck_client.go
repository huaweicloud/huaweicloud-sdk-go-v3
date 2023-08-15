package v2

import (
	http_client "github.com/dysodeng/huaweicloud-sdk-go-v3/core"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/codeartscheck/v2/model"
)

type CodeArtsCheckClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCodeArtsCheckClient(hcClient *http_client.HcHttpClient) *CodeArtsCheckClient {
	return &CodeArtsCheckClient{HcClient: hcClient}
}

func CodeArtsCheckClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CheckParameters 查询任务规则集的检查参数
//
// 查询任务规则集的检查参数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) CheckParameters(request *model.CheckParametersRequest) (*model.CheckParametersResponse, error) {
	requestDef := GenReqDefForCheckParameters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckParametersResponse), nil
	}
}

// CheckParametersInvoker 查询任务规则集的检查参数
func (c *CodeArtsCheckClient) CheckParametersInvoker(request *model.CheckParametersRequest) *CheckParametersInvoker {
	requestDef := GenReqDefForCheckParameters()
	return &CheckParametersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckRecord 历史扫描结果查询
//
// 提供每次扫描的问题数量统计
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) CheckRecord(request *model.CheckRecordRequest) (*model.CheckRecordResponse, error) {
	requestDef := GenReqDefForCheckRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckRecordResponse), nil
	}
}

// CheckRecordInvoker 历史扫描结果查询
func (c *CodeArtsCheckClient) CheckRecordInvoker(request *model.CheckRecordRequest) *CheckRecordInvoker {
	requestDef := GenReqDefForCheckRecord()
	return &CheckRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckRulesetParameters 查询任务规则集的检查参数
//
// 查询任务规则集的检查参数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) CheckRulesetParameters(request *model.CheckRulesetParametersRequest) (*model.CheckRulesetParametersResponse, error) {
	requestDef := GenReqDefForCheckRulesetParameters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckRulesetParametersResponse), nil
	}
}

// CheckRulesetParametersInvoker 查询任务规则集的检查参数
func (c *CodeArtsCheckClient) CheckRulesetParametersInvoker(request *model.CheckRulesetParametersRequest) *CheckRulesetParametersInvoker {
	requestDef := GenReqDefForCheckRulesetParameters()
	return &CheckRulesetParametersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRuleset 创建自定义规则集
//
// 可根据需求灵活的组合规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) CreateRuleset(request *model.CreateRulesetRequest) (*model.CreateRulesetResponse, error) {
	requestDef := GenReqDefForCreateRuleset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRulesetResponse), nil
	}
}

// CreateRulesetInvoker 创建自定义规则集
func (c *CodeArtsCheckClient) CreateRulesetInvoker(request *model.CreateRulesetRequest) *CreateRulesetInvoker {
	requestDef := GenReqDefForCreateRuleset()
	return &CreateRulesetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTask 新建检查任务
//
// 新建检查任务但是不执行。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) CreateTask(request *model.CreateTaskRequest) (*model.CreateTaskResponse, error) {
	requestDef := GenReqDefForCreateTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTaskResponse), nil
	}
}

// CreateTaskInvoker 新建检查任务
func (c *CodeArtsCheckClient) CreateTaskInvoker(request *model.CreateTaskRequest) *CreateTaskInvoker {
	requestDef := GenReqDefForCreateTask()
	return &CreateTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRuleset 删除自定义规则集
//
// 删除自定义规则集，正在使用中的或默认规则集不能删除
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) DeleteRuleset(request *model.DeleteRulesetRequest) (*model.DeleteRulesetResponse, error) {
	requestDef := GenReqDefForDeleteRuleset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRulesetResponse), nil
	}
}

// DeleteRulesetInvoker 删除自定义规则集
func (c *CodeArtsCheckClient) DeleteRulesetInvoker(request *model.DeleteRulesetRequest) *DeleteRulesetInvoker {
	requestDef := GenReqDefForDeleteRuleset()
	return &DeleteRulesetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTask 删除检查任务
//
// 删除检查任务，执行中的任务删除无法再查看
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) DeleteTask(request *model.DeleteTaskRequest) (*model.DeleteTaskResponse, error) {
	requestDef := GenReqDefForDeleteTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTaskResponse), nil
	}
}

// DeleteTaskInvoker 删除检查任务
func (c *CodeArtsCheckClient) DeleteTaskInvoker(request *model.DeleteTaskRequest) *DeleteTaskInvoker {
	requestDef := GenReqDefForDeleteTask()
	return &DeleteTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRules 获取规则列表接口
//
// 根据语言、问题级别等条件查询规则列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) ListRules(request *model.ListRulesRequest) (*model.ListRulesResponse, error) {
	requestDef := GenReqDefForListRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRulesResponse), nil
	}
}

// ListRulesInvoker 获取规则列表接口
func (c *CodeArtsCheckClient) ListRulesInvoker(request *model.ListRulesRequest) *ListRulesInvoker {
	requestDef := GenReqDefForListRules()
	return &ListRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRulesets 查询规则集列表
//
// 根据项目ID、语言等条件查询规则集列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) ListRulesets(request *model.ListRulesetsRequest) (*model.ListRulesetsResponse, error) {
	requestDef := GenReqDefForListRulesets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRulesetsResponse), nil
	}
}

// ListRulesetsInvoker 查询规则集列表
func (c *CodeArtsCheckClient) ListRulesetsInvoker(request *model.ListRulesetsRequest) *ListRulesetsInvoker {
	requestDef := GenReqDefForListRulesets()
	return &ListRulesetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTaskParameter 任务配置检查参数
//
// 任务配置检查参数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) ListTaskParameter(request *model.ListTaskParameterRequest) (*model.ListTaskParameterResponse, error) {
	requestDef := GenReqDefForListTaskParameter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTaskParameterResponse), nil
	}
}

// ListTaskParameterInvoker 任务配置检查参数
func (c *CodeArtsCheckClient) ListTaskParameterInvoker(request *model.ListTaskParameterRequest) *ListTaskParameterInvoker {
	requestDef := GenReqDefForListTaskParameter()
	return &ListTaskParameterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTaskRuleset 查询任务的已选规则集列表
//
// 查询任务的已选规则集列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) ListTaskRuleset(request *model.ListTaskRulesetRequest) (*model.ListTaskRulesetResponse, error) {
	requestDef := GenReqDefForListTaskRuleset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTaskRulesetResponse), nil
	}
}

// ListTaskRulesetInvoker 查询任务的已选规则集列表
func (c *CodeArtsCheckClient) ListTaskRulesetInvoker(request *model.ListTaskRulesetRequest) *ListTaskRulesetInvoker {
	requestDef := GenReqDefForListTaskRuleset()
	return &ListTaskRulesetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTemplateRules 查看规则集的规则列表
//
// 根据项目ID、规则集ID等条件查询规则列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) ListTemplateRules(request *model.ListTemplateRulesRequest) (*model.ListTemplateRulesResponse, error) {
	requestDef := GenReqDefForListTemplateRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplateRulesResponse), nil
	}
}

// ListTemplateRulesInvoker 查看规则集的规则列表
func (c *CodeArtsCheckClient) ListTemplateRulesInvoker(request *model.ListTemplateRulesRequest) *ListTemplateRulesInvoker {
	requestDef := GenReqDefForListTemplateRules()
	return &ListTemplateRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunTask 执行检查任务
//
// 执行检查任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) RunTask(request *model.RunTaskRequest) (*model.RunTaskResponse, error) {
	requestDef := GenReqDefForRunTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunTaskResponse), nil
	}
}

// RunTaskInvoker 执行检查任务
func (c *CodeArtsCheckClient) RunTaskInvoker(request *model.RunTaskRequest) *RunTaskInvoker {
	requestDef := GenReqDefForRunTask()
	return &RunTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetDefaulTemplate 设置每个项目对应语言的默认规则集配置
//
// 设置每个项目对应语言的默认规则集配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) SetDefaulTemplate(request *model.SetDefaulTemplateRequest) (*model.SetDefaulTemplateResponse, error) {
	requestDef := GenReqDefForSetDefaulTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetDefaulTemplateResponse), nil
	}
}

// SetDefaulTemplateInvoker 设置每个项目对应语言的默认规则集配置
func (c *CodeArtsCheckClient) SetDefaulTemplateInvoker(request *model.SetDefaulTemplateRequest) *SetDefaulTemplateInvoker {
	requestDef := GenReqDefForSetDefaulTemplate()
	return &SetDefaulTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProgressDetail 查询任务执行状态
//
// 根据任务ID查询任务执行状态。任务状态：0表示检查中，1表示检查失败，2表示检查成功，3表示任务中止。只有正在检查中才有进度的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) ShowProgressDetail(request *model.ShowProgressDetailRequest) (*model.ShowProgressDetailResponse, error) {
	requestDef := GenReqDefForShowProgressDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProgressDetailResponse), nil
	}
}

// ShowProgressDetailInvoker 查询任务执行状态
func (c *CodeArtsCheckClient) ShowProgressDetailInvoker(request *model.ShowProgressDetailRequest) *ShowProgressDetailInvoker {
	requestDef := GenReqDefForShowProgressDetail()
	return &ShowProgressDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTaskCmetrics 查询cmertrics缺陷概要
//
// 根据检查任务ID查询cmertrics缺陷概要。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) ShowTaskCmetrics(request *model.ShowTaskCmetricsRequest) (*model.ShowTaskCmetricsResponse, error) {
	requestDef := GenReqDefForShowTaskCmetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskCmetricsResponse), nil
	}
}

// ShowTaskCmetricsInvoker 查询cmertrics缺陷概要
func (c *CodeArtsCheckClient) ShowTaskCmetricsInvoker(request *model.ShowTaskCmetricsRequest) *ShowTaskCmetricsInvoker {
	requestDef := GenReqDefForShowTaskCmetrics()
	return &ShowTaskCmetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTaskDefects 查询缺陷详情
//
// 根据检查任务ID分页查询缺陷结果详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) ShowTaskDefects(request *model.ShowTaskDefectsRequest) (*model.ShowTaskDefectsResponse, error) {
	requestDef := GenReqDefForShowTaskDefects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskDefectsResponse), nil
	}
}

// ShowTaskDefectsInvoker 查询缺陷详情
func (c *CodeArtsCheckClient) ShowTaskDefectsInvoker(request *model.ShowTaskDefectsRequest) *ShowTaskDefectsInvoker {
	requestDef := GenReqDefForShowTaskDefects()
	return &ShowTaskDefectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTaskDefectsStatistic 查询缺陷详情的统计
//
// 根据检查任务ID查询缺陷详情的统计
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) ShowTaskDefectsStatistic(request *model.ShowTaskDefectsStatisticRequest) (*model.ShowTaskDefectsStatisticResponse, error) {
	requestDef := GenReqDefForShowTaskDefectsStatistic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskDefectsStatisticResponse), nil
	}
}

// ShowTaskDefectsStatisticInvoker 查询缺陷详情的统计
func (c *CodeArtsCheckClient) ShowTaskDefectsStatisticInvoker(request *model.ShowTaskDefectsStatisticRequest) *ShowTaskDefectsStatisticInvoker {
	requestDef := GenReqDefForShowTaskDefectsStatistic()
	return &ShowTaskDefectsStatisticInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTaskDetail 查询缺陷概要
//
// 根据检查任务ID查询缺陷结果的概要。包括问题概述、问题状态、圈复杂度、代码重复率等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) ShowTaskDetail(request *model.ShowTaskDetailRequest) (*model.ShowTaskDetailResponse, error) {
	requestDef := GenReqDefForShowTaskDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskDetailResponse), nil
	}
}

// ShowTaskDetailInvoker 查询缺陷概要
func (c *CodeArtsCheckClient) ShowTaskDetailInvoker(request *model.ShowTaskDetailRequest) *ShowTaskDetailInvoker {
	requestDef := GenReqDefForShowTaskDetail()
	return &ShowTaskDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTaskListByProjectId 查询任务列表
//
// 根据DEVCLOUD_PROJECT_UUID查询该项目下的任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) ShowTaskListByProjectId(request *model.ShowTaskListByProjectIdRequest) (*model.ShowTaskListByProjectIdResponse, error) {
	requestDef := GenReqDefForShowTaskListByProjectId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskListByProjectIdResponse), nil
	}
}

// ShowTaskListByProjectIdInvoker 查询任务列表
func (c *CodeArtsCheckClient) ShowTaskListByProjectIdInvoker(request *model.ShowTaskListByProjectIdRequest) *ShowTaskListByProjectIdInvoker {
	requestDef := GenReqDefForShowTaskListByProjectId()
	return &ShowTaskListByProjectIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTaskPathTree 获取任务的目录树
//
// 获取任务的目录树
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) ShowTaskPathTree(request *model.ShowTaskPathTreeRequest) (*model.ShowTaskPathTreeResponse, error) {
	requestDef := GenReqDefForShowTaskPathTree()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskPathTreeResponse), nil
	}
}

// ShowTaskPathTreeInvoker 获取任务的目录树
func (c *CodeArtsCheckClient) ShowTaskPathTreeInvoker(request *model.ShowTaskPathTreeRequest) *ShowTaskPathTreeInvoker {
	requestDef := GenReqDefForShowTaskPathTree()
	return &ShowTaskPathTreeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTaskSettings 查询任务的高级选项
//
// 查询任务的高级选项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) ShowTaskSettings(request *model.ShowTaskSettingsRequest) (*model.ShowTaskSettingsResponse, error) {
	requestDef := GenReqDefForShowTaskSettings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskSettingsResponse), nil
	}
}

// ShowTaskSettingsInvoker 查询任务的高级选项
func (c *CodeArtsCheckClient) ShowTaskSettingsInvoker(request *model.ShowTaskSettingsRequest) *ShowTaskSettingsInvoker {
	requestDef := GenReqDefForShowTaskSettings()
	return &ShowTaskSettingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTasklog 查询任务检查失败日志
//
// 查询任务检查失败日志，不传execute_id则查询最近一次的检查日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) ShowTasklog(request *model.ShowTasklogRequest) (*model.ShowTasklogResponse, error) {
	requestDef := GenReqDefForShowTasklog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTasklogResponse), nil
	}
}

// ShowTasklogInvoker 查询任务检查失败日志
func (c *CodeArtsCheckClient) ShowTasklogInvoker(request *model.ShowTasklogRequest) *ShowTasklogInvoker {
	requestDef := GenReqDefForShowTasklog()
	return &ShowTasklogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTasksRulesets 查询任务的已选规则集列表
//
// 查询任务的已选规则集列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) ShowTasksRulesets(request *model.ShowTasksRulesetsRequest) (*model.ShowTasksRulesetsResponse, error) {
	requestDef := GenReqDefForShowTasksRulesets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTasksRulesetsResponse), nil
	}
}

// ShowTasksRulesetsInvoker 查询任务的已选规则集列表
func (c *CodeArtsCheckClient) ShowTasksRulesetsInvoker(request *model.ShowTasksRulesetsRequest) *ShowTasksRulesetsInvoker {
	requestDef := GenReqDefForShowTasksRulesets()
	return &ShowTasksRulesetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopTaskById 终止检查任务
//
// 根据任务ID终止检查任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) StopTaskById(request *model.StopTaskByIdRequest) (*model.StopTaskByIdResponse, error) {
	requestDef := GenReqDefForStopTaskById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopTaskByIdResponse), nil
	}
}

// StopTaskByIdInvoker 终止检查任务
func (c *CodeArtsCheckClient) StopTaskByIdInvoker(request *model.StopTaskByIdRequest) *StopTaskByIdInvoker {
	requestDef := GenReqDefForStopTaskById()
	return &StopTaskByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDefectStatus 修改缺陷状态
//
// 修改检查出的缺陷的状态为已解决、已忽略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) UpdateDefectStatus(request *model.UpdateDefectStatusRequest) (*model.UpdateDefectStatusResponse, error) {
	requestDef := GenReqDefForUpdateDefectStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDefectStatusResponse), nil
	}
}

// UpdateDefectStatusInvoker 修改缺陷状态
func (c *CodeArtsCheckClient) UpdateDefectStatusInvoker(request *model.UpdateDefectStatusRequest) *UpdateDefectStatusInvoker {
	requestDef := GenReqDefForUpdateDefectStatus()
	return &UpdateDefectStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateIgnorePath 任务配置屏蔽目录
//
// 任务配置屏蔽目录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) UpdateIgnorePath(request *model.UpdateIgnorePathRequest) (*model.UpdateIgnorePathResponse, error) {
	requestDef := GenReqDefForUpdateIgnorePath()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIgnorePathResponse), nil
	}
}

// UpdateIgnorePathInvoker 任务配置屏蔽目录
func (c *CodeArtsCheckClient) UpdateIgnorePathInvoker(request *model.UpdateIgnorePathRequest) *UpdateIgnorePathInvoker {
	requestDef := GenReqDefForUpdateIgnorePath()
	return &UpdateIgnorePathInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTaskRuleset 修改任务规则集
//
// 修改任务规则集。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) UpdateTaskRuleset(request *model.UpdateTaskRulesetRequest) (*model.UpdateTaskRulesetResponse, error) {
	requestDef := GenReqDefForUpdateTaskRuleset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTaskRulesetResponse), nil
	}
}

// UpdateTaskRulesetInvoker 修改任务规则集
func (c *CodeArtsCheckClient) UpdateTaskRulesetInvoker(request *model.UpdateTaskRulesetRequest) *UpdateTaskRulesetInvoker {
	requestDef := GenReqDefForUpdateTaskRuleset()
	return &UpdateTaskRulesetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTaskSettings 任务配置高级选项
//
// 任务配置高级选项，如自定义镜像
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsCheckClient) UpdateTaskSettings(request *model.UpdateTaskSettingsRequest) (*model.UpdateTaskSettingsResponse, error) {
	requestDef := GenReqDefForUpdateTaskSettings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTaskSettingsResponse), nil
	}
}

// UpdateTaskSettingsInvoker 任务配置高级选项
func (c *CodeArtsCheckClient) UpdateTaskSettingsInvoker(request *model.UpdateTaskSettingsRequest) *UpdateTaskSettingsInvoker {
	requestDef := GenReqDefForUpdateTaskSettings()
	return &UpdateTaskSettingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
