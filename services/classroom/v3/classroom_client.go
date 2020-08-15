package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/classroom/v3/model"
)

type ClassroomClient struct {
	hcClient *http_client.HcHttpClient
}

func NewClassroomClient(hcClient *http_client.HcHttpClient) *ClassroomClient {
	return &ClassroomClient{hcClient: hcClient}
}

func ClassroomClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//根据课堂ID获取指定课堂的课堂成员列表，支持分页，搜索字段默认同时匹配姓名，学号，用户名，班级。
func (c *ClassroomClient) ListClassroomMembers(request *model.ListClassroomMembersRequest) (*model.ListClassroomMembersResponse, error) {
	requestDef := GenReqDefForListClassroomMembers(request)
	resp, responseDef := GenRespForListClassroomMembers()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//获取当前用户的课堂列表，课堂课表分为我创建的课堂，我加入的课堂以及所有课堂，支持分页查询。
func (c *ClassroomClient) ListClassrooms(request *model.ListClassroomsRequest) (*model.ListClassroomsResponse, error) {
	requestDef := GenReqDefForListClassrooms(request)
	resp, responseDef := GenRespForListClassrooms()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//根据课堂ID获取指定课堂的详细信息
func (c *ClassroomClient) ShowClassroomDetail(request *model.ShowClassroomDetailRequest) (*model.ShowClassroomDetailResponse, error) {
	requestDef := GenReqDefForShowClassroomDetail(request)
	resp, responseDef := GenRespForShowClassroomDetail()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询课堂下指定成员的作业信息
func (c *ClassroomClient) ListClassroomMemberJobs(request *model.ListClassroomMemberJobsRequest) (*model.ListClassroomMemberJobsResponse, error) {
	requestDef := GenReqDefForListClassroomMemberJobs(request)
	resp, responseDef := GenRespForListClassroomMemberJobs()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询指定课堂下的作业列表信息，支持分页查询。
func (c *ClassroomClient) ListJobs(request *model.ListJobsRequest) (*model.ListJobsResponse, error) {
	requestDef := GenReqDefForListJobs(request)
	resp, responseDef := GenRespForListJobs()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询学生指定作业的习题提交记录信息(针对函数习题)
func (c *ClassroomClient) ListMemberJobRecords(request *model.ListMemberJobRecordsRequest) (*model.ListMemberJobRecordsResponse, error) {
	requestDef := GenReqDefForListMemberJobRecords(request)
	resp, responseDef := GenRespForListMemberJobRecords()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//根据作业ID，查询指定作业的信息
func (c *ClassroomClient) ShowJobDetail(request *model.ShowJobDetailRequest) (*model.ShowJobDetailResponse, error) {
	requestDef := GenReqDefForShowJobDetail(request)
	resp, responseDef := GenRespForShowJobDetail()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询指定作业下的习题信息
func (c *ClassroomClient) ShowJobExercises(request *model.ShowJobExercisesRequest) (*model.ShowJobExercisesResponse, error) {
	requestDef := GenReqDefForShowJobExercises(request)
	resp, responseDef := GenRespForShowJobExercises()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}
