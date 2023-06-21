package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/classroom/v3/model"
)

type ClassroomClient struct {
	HcClient *http_client.HcHttpClient
}

func NewClassroomClient(hcClient *http_client.HcHttpClient) *ClassroomClient {
	return &ClassroomClient{HcClient: hcClient}
}

func ClassroomClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// ApplyJudgement 下发判题任务
//
// 下发判题任务，根据回调地址、代码来源、源代码文本、语言类型、超时时长、输出类型，触发后台代码编译运行和判题
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ClassroomClient) ApplyJudgement(request *model.ApplyJudgementRequest) (*model.ApplyJudgementResponse, error) {
	requestDef := GenReqDefForApplyJudgement()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ApplyJudgementResponse), nil
	}
}

// ApplyJudgementInvoker 下发判题任务
func (c *ClassroomClient) ApplyJudgementInvoker(request *model.ApplyJudgementRequest) *ApplyJudgementInvoker {
	requestDef := GenReqDefForApplyJudgement()
	return &ApplyJudgementInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJudgementDetail 获取判题结果详情
//
// 根据判题任务ID获取判题结果详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ClassroomClient) ShowJudgementDetail(request *model.ShowJudgementDetailRequest) (*model.ShowJudgementDetailResponse, error) {
	requestDef := GenReqDefForShowJudgementDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJudgementDetailResponse), nil
	}
}

// ShowJudgementDetailInvoker 获取判题结果详情
func (c *ClassroomClient) ShowJudgementDetailInvoker(request *model.ShowJudgementDetailRequest) *ShowJudgementDetailInvoker {
	requestDef := GenReqDefForShowJudgementDetail()
	return &ShowJudgementDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJudgementFile 下载判题结果文件
//
// 根据文件id或图片id下载输出结果文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ClassroomClient) ShowJudgementFile(request *model.ShowJudgementFileRequest) (*model.ShowJudgementFileResponse, error) {
	requestDef := GenReqDefForShowJudgementFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJudgementFileResponse), nil
	}
}

// ShowJudgementFileInvoker 下载判题结果文件
func (c *ClassroomClient) ShowJudgementFileInvoker(request *model.ShowJudgementFileRequest) *ShowJudgementFileInvoker {
	requestDef := GenReqDefForShowJudgementFile()
	return &ShowJudgementFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteExercise 习题判题
//
// 习题判题
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ClassroomClient) ExecuteExercise(request *model.ExecuteExerciseRequest) (*model.ExecuteExerciseResponse, error) {
	requestDef := GenReqDefForExecuteExercise()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteExerciseResponse), nil
	}
}

// ExecuteExerciseInvoker 习题判题
func (c *ClassroomClient) ExecuteExerciseInvoker(request *model.ExecuteExerciseRequest) *ExecuteExerciseInvoker {
	requestDef := GenReqDefForExecuteExercise()
	return &ExecuteExerciseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListExercises 查询习题库下习题列表
//
// 查询习题库下习题列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ClassroomClient) ListExercises(request *model.ListExercisesRequest) (*model.ListExercisesResponse, error) {
	requestDef := GenReqDefForListExercises()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListExercisesResponse), nil
	}
}

// ListExercisesInvoker 查询习题库下习题列表
func (c *ClassroomClient) ListExercisesInvoker(request *model.ListExercisesRequest) *ListExercisesInvoker {
	requestDef := GenReqDefForListExercises()
	return &ListExercisesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPackages 查询当前租户的习题库列表
//
// 查询当前租户的习题库列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ClassroomClient) ListPackages(request *model.ListPackagesRequest) (*model.ListPackagesResponse, error) {
	requestDef := GenReqDefForListPackages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPackagesResponse), nil
	}
}

// ListPackagesInvoker 查询当前租户的习题库列表
func (c *ClassroomClient) ListPackagesInvoker(request *model.ListPackagesRequest) *ListPackagesInvoker {
	requestDef := GenReqDefForListPackages()
	return &ListPackagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowExerciseDetail 查询单个习题详情
//
// 查询单个习题详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ClassroomClient) ShowExerciseDetail(request *model.ShowExerciseDetailRequest) (*model.ShowExerciseDetailResponse, error) {
	requestDef := GenReqDefForShowExerciseDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowExerciseDetailResponse), nil
	}
}

// ShowExerciseDetailInvoker 查询单个习题详情
func (c *ClassroomClient) ShowExerciseDetailInvoker(request *model.ShowExerciseDetailRequest) *ShowExerciseDetailInvoker {
	requestDef := GenReqDefForShowExerciseDetail()
	return &ShowExerciseDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPackageDetail 查询单个习题库详情
//
// 查询单个习题库详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ClassroomClient) ShowPackageDetail(request *model.ShowPackageDetailRequest) (*model.ShowPackageDetailResponse, error) {
	requestDef := GenReqDefForShowPackageDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPackageDetailResponse), nil
	}
}

// ShowPackageDetailInvoker 查询单个习题库详情
func (c *ClassroomClient) ShowPackageDetailInvoker(request *model.ShowPackageDetailRequest) *ShowPackageDetailInvoker {
	requestDef := GenReqDefForShowPackageDetail()
	return &ShowPackageDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllDifficults 获取习题所有难度
//
// 获取习题所有难度
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ClassroomClient) ListAllDifficults(request *model.ListAllDifficultsRequest) (*model.ListAllDifficultsResponse, error) {
	requestDef := GenReqDefForListAllDifficults()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllDifficultsResponse), nil
	}
}

// ListAllDifficultsInvoker 获取习题所有难度
func (c *ClassroomClient) ListAllDifficultsInvoker(request *model.ListAllDifficultsRequest) *ListAllDifficultsInvoker {
	requestDef := GenReqDefForListAllDifficults()
	return &ListAllDifficultsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMyKnowledgePoints 获取自身习题知识点
//
// 获取自身习题知识点
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ClassroomClient) ListMyKnowledgePoints(request *model.ListMyKnowledgePointsRequest) (*model.ListMyKnowledgePointsResponse, error) {
	requestDef := GenReqDefForListMyKnowledgePoints()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMyKnowledgePointsResponse), nil
	}
}

// ListMyKnowledgePointsInvoker 获取自身习题知识点
func (c *ClassroomClient) ListMyKnowledgePointsInvoker(request *model.ListMyKnowledgePointsRequest) *ListMyKnowledgePointsInvoker {
	requestDef := GenReqDefForListMyKnowledgePoints()
	return &ListMyKnowledgePointsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClassroomMembers 根据课堂ID获取指定课堂的课堂成员列表
//
// 根据课堂ID获取指定课堂的课堂成员列表，支持分页，搜索字段默认同时匹配姓名，学号，用户名，班级。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ClassroomClient) ListClassroomMembers(request *model.ListClassroomMembersRequest) (*model.ListClassroomMembersResponse, error) {
	requestDef := GenReqDefForListClassroomMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClassroomMembersResponse), nil
	}
}

// ListClassroomMembersInvoker 根据课堂ID获取指定课堂的课堂成员列表
func (c *ClassroomClient) ListClassroomMembersInvoker(request *model.ListClassroomMembersRequest) *ListClassroomMembersInvoker {
	requestDef := GenReqDefForListClassroomMembers()
	return &ListClassroomMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClassrooms 获取当前用户的课堂列表
//
// 获取当前用户的课堂列表，课堂课表分为我创建的课堂，我加入的课堂以及所有课堂，支持分页查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ClassroomClient) ListClassrooms(request *model.ListClassroomsRequest) (*model.ListClassroomsResponse, error) {
	requestDef := GenReqDefForListClassrooms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClassroomsResponse), nil
	}
}

// ListClassroomsInvoker 获取当前用户的课堂列表
func (c *ClassroomClient) ListClassroomsInvoker(request *model.ListClassroomsRequest) *ListClassroomsInvoker {
	requestDef := GenReqDefForListClassrooms()
	return &ListClassroomsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowClassroomDetail 根据课堂ID获取指定课堂的详细信息
//
// 根据课堂ID获取指定课堂的详细信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ClassroomClient) ShowClassroomDetail(request *model.ShowClassroomDetailRequest) (*model.ShowClassroomDetailResponse, error) {
	requestDef := GenReqDefForShowClassroomDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClassroomDetailResponse), nil
	}
}

// ShowClassroomDetailInvoker 根据课堂ID获取指定课堂的详细信息
func (c *ClassroomClient) ShowClassroomDetailInvoker(request *model.ShowClassroomDetailRequest) *ShowClassroomDetailInvoker {
	requestDef := GenReqDefForShowClassroomDetail()
	return &ShowClassroomDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClassroomMemberJobs 查询课堂下指定成员的作业信息
//
// 查询课堂下指定成员的作业信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ClassroomClient) ListClassroomMemberJobs(request *model.ListClassroomMemberJobsRequest) (*model.ListClassroomMemberJobsResponse, error) {
	requestDef := GenReqDefForListClassroomMemberJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClassroomMemberJobsResponse), nil
	}
}

// ListClassroomMemberJobsInvoker 查询课堂下指定成员的作业信息
func (c *ClassroomClient) ListClassroomMemberJobsInvoker(request *model.ListClassroomMemberJobsRequest) *ListClassroomMemberJobsInvoker {
	requestDef := GenReqDefForListClassroomMemberJobs()
	return &ListClassroomMemberJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListJobs 查询指定课堂下的作业列表信息
//
// 查询指定课堂下的作业列表信息，支持分页查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ClassroomClient) ListJobs(request *model.ListJobsRequest) (*model.ListJobsResponse, error) {
	requestDef := GenReqDefForListJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJobsResponse), nil
	}
}

// ListJobsInvoker 查询指定课堂下的作业列表信息
func (c *ClassroomClient) ListJobsInvoker(request *model.ListJobsRequest) *ListJobsInvoker {
	requestDef := GenReqDefForListJobs()
	return &ListJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMemberJobRecords 查询学生函数习题提交记录信息
//
// 查询学生指定作业的习题提交记录信息(针对函数习题)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ClassroomClient) ListMemberJobRecords(request *model.ListMemberJobRecordsRequest) (*model.ListMemberJobRecordsResponse, error) {
	requestDef := GenReqDefForListMemberJobRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMemberJobRecordsResponse), nil
	}
}

// ListMemberJobRecordsInvoker 查询学生函数习题提交记录信息
func (c *ClassroomClient) ListMemberJobRecordsInvoker(request *model.ListMemberJobRecordsRequest) *ListMemberJobRecordsInvoker {
	requestDef := GenReqDefForListMemberJobRecords()
	return &ListMemberJobRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobDetail 根据作业ID，查询指定作业的信息
//
// 根据作业ID，查询指定作业的信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ClassroomClient) ShowJobDetail(request *model.ShowJobDetailRequest) (*model.ShowJobDetailResponse, error) {
	requestDef := GenReqDefForShowJobDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobDetailResponse), nil
	}
}

// ShowJobDetailInvoker 根据作业ID，查询指定作业的信息
func (c *ClassroomClient) ShowJobDetailInvoker(request *model.ShowJobDetailRequest) *ShowJobDetailInvoker {
	requestDef := GenReqDefForShowJobDetail()
	return &ShowJobDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobExercises 查询指定作业下的习题信息
//
// 查询指定作业下的习题信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ClassroomClient) ShowJobExercises(request *model.ShowJobExercisesRequest) (*model.ShowJobExercisesResponse, error) {
	requestDef := GenReqDefForShowJobExercises()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobExercisesResponse), nil
	}
}

// ShowJobExercisesInvoker 查询指定作业下的习题信息
func (c *ClassroomClient) ShowJobExercisesInvoker(request *model.ShowJobExercisesRequest) *ShowJobExercisesInvoker {
	requestDef := GenReqDefForShowJobExercises()
	return &ShowJobExercisesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
