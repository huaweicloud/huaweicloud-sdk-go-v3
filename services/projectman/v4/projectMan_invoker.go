package v4

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/projectman/v4/model"
)

type AddApplyJoinProjectForAgcInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddApplyJoinProjectForAgcInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *AddMemberV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *BatchAddMembersV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *BatchDeleteMembersV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *BatchUpdateChildNickNamesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CheckProjectNameV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateProjectV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteProjectV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListDomainNotAddedProjectsV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListProjectBugStaticsV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListProjectDemandStaticV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListProjectMembersV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListProjectsV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectsV4Invoker) Invoke() (*model.ListProjectsV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectsV4Response), nil
	}
}

type ListTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTemplatesInvoker) Invoke() (*model.ListTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTemplatesResponse), nil
	}
}

type ListWorkitemStatusRecordsV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkitemStatusRecordsV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkitemStatusRecordsV4Invoker) Invoke() (*model.ListWorkitemStatusRecordsV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkitemStatusRecordsV4Response), nil
	}
}

type ListWorkitemsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkitemsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkitemsInvoker) Invoke() (*model.ListWorkitemsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkitemsResponse), nil
	}
}

type RemoveProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *RemoveProjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowBugDensityV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowBugsPerDeveloperInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowCompletionRateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowCurUserInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowCurUserRoleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowProjectInfoV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowProjectSummaryV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProjectSummaryV4Invoker) Invoke() (*model.ShowProjectSummaryV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectSummaryV4Response), nil
	}
}

type ShowWorkItemWrokflowConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkItemWrokflowConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWorkItemWrokflowConfigInvoker) Invoke() (*model.ShowWorkItemWrokflowConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkItemWrokflowConfigResponse), nil
	}
}

type UpdateMembesRoleV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMembesRoleV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateNickNameV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateProjectV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateProjectV4Invoker) Invoke() (*model.UpdateProjectV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProjectV4Response), nil
	}
}

type CreateIpdProjectIssueInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateIpdProjectIssueInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateIpdProjectIssueInvoker) Invoke() (*model.CreateIpdProjectIssueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateIpdProjectIssueResponse), nil
	}
}

type CreateIpdProjectIssueAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateIpdProjectIssueAttachmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateIpdProjectIssueAttachmentInvoker) Invoke() (*model.CreateIpdProjectIssueAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateIpdProjectIssueAttachmentResponse), nil
	}
}

type ListIpdProjectIssuesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIpdProjectIssuesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIpdProjectIssuesInvoker) Invoke() (*model.ListIpdProjectIssuesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIpdProjectIssuesResponse), nil
	}
}

type ListIssueFieldsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIssueFieldsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIssueFieldsInvoker) Invoke() (*model.ListIssueFieldsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIssueFieldsResponse), nil
	}
}

type ListIssueStatuesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIssueStatuesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIssueStatuesInvoker) Invoke() (*model.ListIssueStatuesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIssueStatuesResponse), nil
	}
}

type ShowIssueConfigFieldsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIssueConfigFieldsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowIssueConfigFieldsInvoker) Invoke() (*model.ShowIssueConfigFieldsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIssueConfigFieldsResponse), nil
	}
}

type ShowIssueDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIssueDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowIssueDetailInvoker) Invoke() (*model.ShowIssueDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIssueDetailResponse), nil
	}
}

type ShowWorkflowTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkflowTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWorkflowTemplateInvoker) Invoke() (*model.ShowWorkflowTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkflowTemplateResponse), nil
	}
}

type TransferWorkItemFlowInvoker struct {
	*invoker.BaseInvoker
}

func (i *TransferWorkItemFlowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *TransferWorkItemFlowInvoker) Invoke() (*model.TransferWorkItemFlowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.TransferWorkItemFlowResponse), nil
	}
}

type DownloadIpdIssueAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadIpdIssueAttachmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadIpdIssueAttachmentInvoker) Invoke() (*model.DownloadIpdIssueAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadIpdIssueAttachmentResponse), nil
	}
}

type ShowIpdAttachmentByWorkItemIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIpdAttachmentByWorkItemIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowIpdAttachmentByWorkItemIdInvoker) Invoke() (*model.ShowIpdAttachmentByWorkItemIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIpdAttachmentByWorkItemIdResponse), nil
	}
}

type CreateScrumPlanToProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateScrumPlanToProjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateScrumPlanToProjectInvoker) Invoke() (*model.CreateScrumPlanToProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateScrumPlanToProjectResponse), nil
	}
}

type DeleteScrumPlanInProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteScrumPlanInProjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteScrumPlanInProjectInvoker) Invoke() (*model.DeleteScrumPlanInProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteScrumPlanInProjectResponse), nil
	}
}

type ShowScrumPlansByConditionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowScrumPlansByConditionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowScrumPlansByConditionInvoker) Invoke() (*model.ShowScrumPlansByConditionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowScrumPlansByConditionResponse), nil
	}
}

type UpdateScrumPlanInProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateScrumPlanInProjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateScrumPlanInProjectInvoker) Invoke() (*model.UpdateScrumPlanInProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateScrumPlanInProjectResponse), nil
	}
}

type AddIssueWorkHoursInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddIssueWorkHoursInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddIssueWorkHoursInvoker) Invoke() (*model.AddIssueWorkHoursResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddIssueWorkHoursResponse), nil
	}
}

type BatchDeleteIssuesV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteIssuesV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *BatchDeleteIterationsV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteIterationsV4Invoker) Invoke() (*model.BatchDeleteIterationsV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteIterationsV4Response), nil
	}
}

type BatchListAssociatedIssuesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchListAssociatedIssuesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchListAssociatedIssuesInvoker) Invoke() (*model.BatchListAssociatedIssuesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchListAssociatedIssuesResponse), nil
	}
}

type CancelProjectDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelProjectDomainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelProjectDomainInvoker) Invoke() (*model.CancelProjectDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelProjectDomainResponse), nil
	}
}

type CreateCustomfieldsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCustomfieldsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateIssueV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateIterationV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateIterationV4Invoker) Invoke() (*model.CreateIterationV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateIterationV4Response), nil
	}
}

type CreateProjectDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateProjectDomainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateProjectDomainInvoker) Invoke() (*model.CreateProjectDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateProjectDomainResponse), nil
	}
}

type CreateProjectModuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateProjectModuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateProjectModuleInvoker) Invoke() (*model.CreateProjectModuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateProjectModuleResponse), nil
	}
}

type CreateSystemIssueV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSystemIssueV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSystemIssueV4Invoker) Invoke() (*model.CreateSystemIssueV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSystemIssueV4Response), nil
	}
}

type DeleteAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAttachmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAttachmentInvoker) Invoke() (*model.DeleteAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAttachmentResponse), nil
	}
}

type DeleteIssueV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteIssueV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteIterationV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteIterationV4Invoker) Invoke() (*model.DeleteIterationV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteIterationV4Response), nil
	}
}

type DeleteProjectModuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteProjectModuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteProjectModuleInvoker) Invoke() (*model.DeleteProjectModuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteProjectModuleResponse), nil
	}
}

type DownloadAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadAttachmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadAttachmentInvoker) Invoke() (*model.DownloadAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadAttachmentResponse), nil
	}
}

type DownloadImageFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadImageFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadImageFileInvoker) Invoke() (*model.DownloadImageFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadImageFileResponse), nil
	}
}

type ListAssociatedIssuesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAssociatedIssuesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAssociatedIssuesInvoker) Invoke() (*model.ListAssociatedIssuesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAssociatedIssuesResponse), nil
	}
}

type ListAssociatedTestCasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAssociatedTestCasesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAssociatedTestCasesInvoker) Invoke() (*model.ListAssociatedTestCasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAssociatedTestCasesResponse), nil
	}
}

type ListAssociatedWikisInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAssociatedWikisInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAssociatedWikisInvoker) Invoke() (*model.ListAssociatedWikisResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAssociatedWikisResponse), nil
	}
}

type ListChildIssuesV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListChildIssuesV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListChildIssuesV4Invoker) Invoke() (*model.ListChildIssuesV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListChildIssuesV4Response), nil
	}
}

type ListIssueAssociatedCommitsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIssueAssociatedCommitsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIssueAssociatedCommitsInvoker) Invoke() (*model.ListIssueAssociatedCommitsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIssueAssociatedCommitsResponse), nil
	}
}

type ListIssueCommentsV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListIssueCommentsV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIssueCommentsV4Invoker) Invoke() (*model.ListIssueCommentsV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIssueCommentsV4Response), nil
	}
}

type ListIssueCustomFieldsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIssueCustomFieldsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIssueCustomFieldsInvoker) Invoke() (*model.ListIssueCustomFieldsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIssueCustomFieldsResponse), nil
	}
}

type ListIssueRecordsV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListIssueRecordsV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIssueRecordsV4Invoker) Invoke() (*model.ListIssueRecordsV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIssueRecordsV4Response), nil
	}
}

type ListIssuesSfV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListIssuesSfV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIssuesSfV4Invoker) Invoke() (*model.ListIssuesSfV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIssuesSfV4Response), nil
	}
}

type ListIssuesV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListIssuesV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListIterationHistoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIterationHistoriesInvoker) Invoke() (*model.ListIterationHistoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIterationHistoriesResponse), nil
	}
}

type ListProjectDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectDomainsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectDomainsInvoker) Invoke() (*model.ListProjectDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectDomainsResponse), nil
	}
}

type ListProjectIssuesRecordsV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectIssuesRecordsV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectIssuesRecordsV4Invoker) Invoke() (*model.ListProjectIssuesRecordsV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectIssuesRecordsV4Response), nil
	}
}

type ListProjectIterationsV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectIterationsV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectIterationsV4Invoker) Invoke() (*model.ListProjectIterationsV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectIterationsV4Response), nil
	}
}

type ListProjectModulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectModulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectModulesInvoker) Invoke() (*model.ListProjectModulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectModulesResponse), nil
	}
}

type ListProjectWorkHoursInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectWorkHoursInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectWorkHoursInvoker) Invoke() (*model.ListProjectWorkHoursResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectWorkHoursResponse), nil
	}
}

type ListProjectWorkHoursTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectWorkHoursTypeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectWorkHoursTypeInvoker) Invoke() (*model.ListProjectWorkHoursTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectWorkHoursTypeResponse), nil
	}
}

type ListScrumProjectStatusesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScrumProjectStatusesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScrumProjectStatusesInvoker) Invoke() (*model.ListScrumProjectStatusesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScrumProjectStatusesResponse), nil
	}
}

type ListSpecIssueStayTimesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSpecIssueStayTimesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSpecIssueStayTimesInvoker) Invoke() (*model.ListSpecIssueStayTimesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSpecIssueStayTimesResponse), nil
	}
}

type ListStatusStatisticInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStatusStatisticInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStatusStatisticInvoker) Invoke() (*model.ListStatusStatisticResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStatusStatisticResponse), nil
	}
}

type SearchIssuesInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchIssuesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchIssuesInvoker) Invoke() (*model.SearchIssuesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchIssuesResponse), nil
	}
}

type ShowIssueCompletionRateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIssueCompletionRateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowIssueV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowIssueV4Invoker) Invoke() (*model.ShowIssueV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIssueV4Response), nil
	}
}

type ShowIssuesWrokFlowConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIssuesWrokFlowConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowIssuesWrokFlowConfigInvoker) Invoke() (*model.ShowIssuesWrokFlowConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIssuesWrokFlowConfigResponse), nil
	}
}

type ShowIterationV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIterationV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowProjectWorkHoursInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateIssueV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateIterationV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateIterationV4Invoker) Invoke() (*model.UpdateIterationV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateIterationV4Response), nil
	}
}

type UpdateProjectDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProjectDomainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateProjectDomainInvoker) Invoke() (*model.UpdateProjectDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProjectDomainResponse), nil
	}
}

type UpdateProjectModuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProjectModuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateProjectModuleInvoker) Invoke() (*model.UpdateProjectModuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProjectModuleResponse), nil
	}
}

type UploadAttachmentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadAttachmentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UploadAttachmentsInvoker) Invoke() (*model.UploadAttachmentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadAttachmentsResponse), nil
	}
}

type UploadIssueImgInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadIssueImgInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UploadIssueImgInvoker) Invoke() (*model.UploadIssueImgResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadIssueImgResponse), nil
	}
}

type ShowScrumIssueSeveritiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowScrumIssueSeveritiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowScrumIssueSeveritiesInvoker) Invoke() (*model.ShowScrumIssueSeveritiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowScrumIssueSeveritiesResponse), nil
	}
}
