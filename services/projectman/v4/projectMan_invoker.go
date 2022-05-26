package v4

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/projectman/v4/model"
)

type AddApplyJoinProjectForAgcInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddApplyJoinProjectForAgcInvoker) Invoke() (*model.AddApplyJoinProjectForAgcResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddApplyJoinProjectForAgcResponse), nil
	}
}

type AddMemberV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *AddMemberV4Invoker) Invoke() (*model.AddMemberV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddMemberV4Response), nil
	}
}

type BatchAddMembersV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAddMembersV4Invoker) Invoke() (*model.BatchAddMembersV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAddMembersV4Response), nil
	}
}

type BatchDeleteMembersV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteMembersV4Invoker) Invoke() (*model.BatchDeleteMembersV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteMembersV4Response), nil
	}
}

type BatchUpdateChildNickNamesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateChildNickNamesInvoker) Invoke() (*model.BatchUpdateChildNickNamesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateChildNickNamesResponse), nil
	}
}

type CheckProjectNameV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *CheckProjectNameV4Invoker) Invoke() (*model.CheckProjectNameV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckProjectNameV4Response), nil
	}
}

type CreateProjectV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateProjectV4Invoker) Invoke() (*model.CreateProjectV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateProjectV4Response), nil
	}
}

type DeleteProjectV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteProjectV4Invoker) Invoke() (*model.DeleteProjectV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteProjectV4Response), nil
	}
}

type ListDomainNotAddedProjectsV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListDomainNotAddedProjectsV4Invoker) Invoke() (*model.ListDomainNotAddedProjectsV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDomainNotAddedProjectsV4Response), nil
	}
}

type ListProjectBugStaticsV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectBugStaticsV4Invoker) Invoke() (*model.ListProjectBugStaticsV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectBugStaticsV4Response), nil
	}
}

type ListProjectDemandStaticV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectDemandStaticV4Invoker) Invoke() (*model.ListProjectDemandStaticV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectDemandStaticV4Response), nil
	}
}

type ListProjectMembersV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectMembersV4Invoker) Invoke() (*model.ListProjectMembersV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectMembersV4Response), nil
	}
}

type ListProjectsV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectsV4Invoker) Invoke() (*model.ListProjectsV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectsV4Response), nil
	}
}

type RemoveProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *RemoveProjectInvoker) Invoke() (*model.RemoveProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RemoveProjectResponse), nil
	}
}

type ShowBugDensityV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBugDensityV2Invoker) Invoke() (*model.ShowBugDensityV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBugDensityV2Response), nil
	}
}

type ShowBugsPerDeveloperInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBugsPerDeveloperInvoker) Invoke() (*model.ShowBugsPerDeveloperResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBugsPerDeveloperResponse), nil
	}
}

type ShowCompletionRateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCompletionRateInvoker) Invoke() (*model.ShowCompletionRateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCompletionRateResponse), nil
	}
}

type ShowCurUserInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCurUserInfoInvoker) Invoke() (*model.ShowCurUserInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCurUserInfoResponse), nil
	}
}

type ShowCurUserRoleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCurUserRoleInvoker) Invoke() (*model.ShowCurUserRoleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCurUserRoleResponse), nil
	}
}

type ShowProjectInfoV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectInfoV4Invoker) Invoke() (*model.ShowProjectInfoV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectInfoV4Response), nil
	}
}

type ShowProjectSummaryV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectSummaryV4Invoker) Invoke() (*model.ShowProjectSummaryV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectSummaryV4Response), nil
	}
}

type UpdateMembesRoleV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMembesRoleV4Invoker) Invoke() (*model.UpdateMembesRoleV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMembesRoleV4Response), nil
	}
}

type UpdateNickNameV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNickNameV4Invoker) Invoke() (*model.UpdateNickNameV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNickNameV4Response), nil
	}
}

type UpdateProjectV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProjectV4Invoker) Invoke() (*model.UpdateProjectV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProjectV4Response), nil
	}
}

type BatchDeleteIssuesV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteIssuesV4Invoker) Invoke() (*model.BatchDeleteIssuesV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteIssuesV4Response), nil
	}
}

type BatchDeleteIterationsV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteIterationsV4Invoker) Invoke() (*model.BatchDeleteIterationsV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteIterationsV4Response), nil
	}
}

type CreateCustomfieldsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCustomfieldsInvoker) Invoke() (*model.CreateCustomfieldsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCustomfieldsResponse), nil
	}
}

type CreateIssueV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateIssueV4Invoker) Invoke() (*model.CreateIssueV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateIssueV4Response), nil
	}
}

type CreateIterationV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateIterationV4Invoker) Invoke() (*model.CreateIterationV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateIterationV4Response), nil
	}
}

type CreateSystemIssueV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSystemIssueV4Invoker) Invoke() (*model.CreateSystemIssueV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSystemIssueV4Response), nil
	}
}

type DeleteIssueV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteIssueV4Invoker) Invoke() (*model.DeleteIssueV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteIssueV4Response), nil
	}
}

type DeleteIterationV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteIterationV4Invoker) Invoke() (*model.DeleteIterationV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteIterationV4Response), nil
	}
}

type ListChildIssuesV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListChildIssuesV4Invoker) Invoke() (*model.ListChildIssuesV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListChildIssuesV4Response), nil
	}
}

type ListIssueCommentsV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListIssueCommentsV4Invoker) Invoke() (*model.ListIssueCommentsV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIssueCommentsV4Response), nil
	}
}

type ListIssueRecordsV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListIssueRecordsV4Invoker) Invoke() (*model.ListIssueRecordsV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIssueRecordsV4Response), nil
	}
}

type ListIssuesV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListIssuesV4Invoker) Invoke() (*model.ListIssuesV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIssuesV4Response), nil
	}
}

type ListIterationHistoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIterationHistoriesInvoker) Invoke() (*model.ListIterationHistoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIterationHistoriesResponse), nil
	}
}

type ListProjectIterationsV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectIterationsV4Invoker) Invoke() (*model.ListProjectIterationsV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectIterationsV4Response), nil
	}
}

type ListProjectWorkHoursInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectWorkHoursInvoker) Invoke() (*model.ListProjectWorkHoursResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectWorkHoursResponse), nil
	}
}

type ShowIssueCompletionRateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIssueCompletionRateInvoker) Invoke() (*model.ShowIssueCompletionRateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIssueCompletionRateResponse), nil
	}
}

type ShowIssueV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIssueV4Invoker) Invoke() (*model.ShowIssueV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIssueV4Response), nil
	}
}

type ShowIterationV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIterationV4Invoker) Invoke() (*model.ShowIterationV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIterationV4Response), nil
	}
}

type ShowProjectWorkHoursInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectWorkHoursInvoker) Invoke() (*model.ShowProjectWorkHoursResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectWorkHoursResponse), nil
	}
}

type UpdateIssueV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateIssueV4Invoker) Invoke() (*model.UpdateIssueV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateIssueV4Response), nil
	}
}

type UpdateIterationV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateIterationV4Invoker) Invoke() (*model.UpdateIterationV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateIterationV4Response), nil
	}
}

type UploadIssueImgInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadIssueImgInvoker) Invoke() (*model.UploadIssueImgResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadIssueImgResponse), nil
	}
}
