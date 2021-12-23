package v2

import (
	http_client "github.com/RandolphCYG/huaweicloud-sdk-go-v3/core"

	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/services/codecheck/v2/model"
)

type CodeCheckClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCodeCheckClient(hcClient *http_client.HcHttpClient) *CodeCheckClient {
	return &CodeCheckClient{HcClient: hcClient}
}

func CodeCheckClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//可根据需求灵活的组合规则。
func (c *CodeCheckClient) CreateRuleset(request *model.CreateRulesetRequest) (*model.CreateRulesetResponse, error) {
	requestDef := GenReqDefForCreateRuleset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRulesetResponse), nil
	}
}

//新建检查任务但是不执行。
func (c *CodeCheckClient) CreateTask(request *model.CreateTaskRequest) (*model.CreateTaskResponse, error) {
	requestDef := GenReqDefForCreateTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTaskResponse), nil
	}
}

//删除检查任务，执行中的任务删除无法再查看
func (c *CodeCheckClient) DeleteTask(request *model.DeleteTaskRequest) (*model.DeleteTaskResponse, error) {
	requestDef := GenReqDefForDeleteTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTaskResponse), nil
	}
}

//根据语言、问题级别等条件查询规则列表。
func (c *CodeCheckClient) ListRules(request *model.ListRulesRequest) (*model.ListRulesResponse, error) {
	requestDef := GenReqDefForListRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRulesResponse), nil
	}
}

//根据项目ID、语言等条件查询规则集列表。
func (c *CodeCheckClient) ListRulesets(request *model.ListRulesetsRequest) (*model.ListRulesetsResponse, error) {
	requestDef := GenReqDefForListRulesets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRulesetsResponse), nil
	}
}

//执行检查任务。
func (c *CodeCheckClient) RunTask(request *model.RunTaskRequest) (*model.RunTaskResponse, error) {
	requestDef := GenReqDefForRunTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunTaskResponse), nil
	}
}

//根据任务ID查询任务执行状态。任务状态：0表示检查中，1表示检查失败，2表示检查成功，3表示任务中止。只有正在检查中才有进度的详细信息。
func (c *CodeCheckClient) ShowProgressDetail(request *model.ShowProgressDetailRequest) (*model.ShowProgressDetailResponse, error) {
	requestDef := GenReqDefForShowProgressDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProgressDetailResponse), nil
	}
}

//根据检查任务ID分页查询缺陷结果详情。
func (c *CodeCheckClient) ShowTaskDefects(request *model.ShowTaskDefectsRequest) (*model.ShowTaskDefectsResponse, error) {
	requestDef := GenReqDefForShowTaskDefects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskDefectsResponse), nil
	}
}

//根据检查任务ID查询缺陷详情的统计
func (c *CodeCheckClient) ShowTaskDefectsStatistic(request *model.ShowTaskDefectsStatisticRequest) (*model.ShowTaskDefectsStatisticResponse, error) {
	requestDef := GenReqDefForShowTaskDefectsStatistic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskDefectsStatisticResponse), nil
	}
}

//根据检查任务ID查询缺陷结果的概要。包括问题概述、问题状态、圈复杂度、代码重复率等。
func (c *CodeCheckClient) ShowTaskDetail(request *model.ShowTaskDetailRequest) (*model.ShowTaskDetailResponse, error) {
	requestDef := GenReqDefForShowTaskDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskDetailResponse), nil
	}
}

//根据DEVCLOUD_PROJECT_UUID查询该项目下的任务列表。
func (c *CodeCheckClient) ShowTaskListByProjectId(request *model.ShowTaskListByProjectIdRequest) (*model.ShowTaskListByProjectIdResponse, error) {
	requestDef := GenReqDefForShowTaskListByProjectId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskListByProjectIdResponse), nil
	}
}

//根据任务ID终止检查任务。
func (c *CodeCheckClient) StopTaskById(request *model.StopTaskByIdRequest) (*model.StopTaskByIdResponse, error) {
	requestDef := GenReqDefForStopTaskById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopTaskByIdResponse), nil
	}
}

//修改检查出的缺陷的状态为已解决、已忽略
func (c *CodeCheckClient) UpdateDefectStatus(request *model.UpdateDefectStatusRequest) (*model.UpdateDefectStatusResponse, error) {
	requestDef := GenReqDefForUpdateDefectStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDefectStatusResponse), nil
	}
}

//修改任务规则集。
func (c *CodeCheckClient) UpdateTaskRuleset(request *model.UpdateTaskRulesetRequest) (*model.UpdateTaskRulesetResponse, error) {
	requestDef := GenReqDefForUpdateTaskRuleset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTaskRulesetResponse), nil
	}
}
