package v4

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/projectman/v4/model"
)

type ProjectManClient struct {
	hcClient *http_client.HcHttpClient
}

func NewProjectManClient(hcClient *http_client.HcHttpClient) *ProjectManClient {
	return &ProjectManClient{hcClient: hcClient}
}

func ProjectManClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//创建项目
func (c *ProjectManClient) CreateProjectV4(request *model.CreateProjectV4Request) (*model.CreateProjectV4Response, error) {
	requestDef := GenReqDefForCreateProjectV4(request)
	resp, responseDef := GenRespForCreateProjectV4()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//获取租户没有加入的项目
func (c *ProjectManClient) ListDomainNotAddedProjectsV4(request *model.ListDomainNotAddedProjectsV4Request) (*model.ListDomainNotAddedProjectsV4Response, error) {
	requestDef := GenReqDefForListDomainNotAddedProjectsV4(request)
	resp, responseDef := GenRespForListDomainNotAddedProjectsV4()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//获取工作项的评论
func (c *ProjectManClient) ListIssueCommentsV4(request *model.ListIssueCommentsV4Request) (*model.ListIssueCommentsV4Response, error) {
	requestDef := GenReqDefForListIssueCommentsV4(request)
	resp, responseDef := GenRespForListIssueCommentsV4()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//获取项目成员列表
func (c *ProjectManClient) ListIssueRecordsV4(request *model.ListIssueRecordsV4Request) (*model.ListIssueRecordsV4Response, error) {
	requestDef := GenReqDefForListIssueRecordsV4(request)
	resp, responseDef := GenRespForListIssueRecordsV4()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//获取项目成员列表
func (c *ProjectManClient) ListProjectMembersV4(request *model.ListProjectMembersV4Request) (*model.ListProjectMembersV4Response, error) {
	requestDef := GenReqDefForListProjectMembersV4(request)
	resp, responseDef := GenRespForListProjectMembersV4()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//获取项目迭代
func (c *ProjectManClient) ListProjectVersionsV4(request *model.ListProjectVersionsV4Request) (*model.ListProjectVersionsV4Response, error) {
	requestDef := GenReqDefForListProjectVersionsV4(request)
	resp, responseDef := GenRespForListProjectVersionsV4()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//按用户查询工时（多项目）
func (c *ProjectManClient) ListProjectWorkHours(request *model.ListProjectWorkHoursRequest) (*model.ListProjectWorkHoursResponse, error) {
	requestDef := GenReqDefForListProjectWorkHours(request)
	resp, responseDef := GenRespForListProjectWorkHours()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询项目列表
func (c *ProjectManClient) ListProjectsV4(request *model.ListProjectsV4Request) (*model.ListProjectsV4Response, error) {
	requestDef := GenReqDefForListProjectsV4(request)
	resp, responseDef := GenRespForListProjectsV4()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//项目成员主动退出项目，项目创建者不能退出
func (c *ProjectManClient) RemoveProject(request *model.RemoveProjectRequest) (*model.RemoveProjectResponse, error) {
	requestDef := GenReqDefForRemoveProject(request)
	resp, responseDef := GenRespForRemoveProject()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//获取当前用户信息
func (c *ProjectManClient) ShowCurUserInfo(request *model.ShowCurUserInfoRequest) (*model.ShowCurUserInfoResponse, error) {
	requestDef := GenReqDefForShowCurUserInfo(request)
	resp, responseDef := GenRespForShowCurUserInfo()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//获取用户在项目中的角色
func (c *ProjectManClient) ShowCurUserRole(request *model.ShowCurUserRoleRequest) (*model.ShowCurUserRoleResponse, error) {
	requestDef := GenReqDefForShowCurUserRole(request)
	resp, responseDef := GenRespForShowCurUserRole()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//按用户查询工时（单项目）
func (c *ProjectManClient) ShowProjectWorkHours(request *model.ShowProjectWorkHoursRequest) (*model.ShowProjectWorkHoursResponse, error) {
	requestDef := GenReqDefForShowProjectWorkHours(request)
	resp, responseDef := GenRespForShowProjectWorkHours()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//获取工作项的完成率
func (c *ProjectManClient) ShowtIssueCompletionRate(request *model.ShowtIssueCompletionRateRequest) (*model.ShowtIssueCompletionRateResponse, error) {
	requestDef := GenReqDefForShowtIssueCompletionRate(request)
	resp, responseDef := GenRespForShowtIssueCompletionRate()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//更新成员在项目中的角色
func (c *ProjectManClient) UpdateMembesRoleV4(request *model.UpdateMembesRoleV4Request) (*model.UpdateMembesRoleV4Response, error) {
	requestDef := GenReqDefForUpdateMembesRoleV4(request)
	resp, responseDef := GenRespForUpdateMembesRoleV4()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}
