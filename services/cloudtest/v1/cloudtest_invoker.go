package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cloudtest/v1/model"
)

type BatchAddRelationsByOneCaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAddRelationsByOneCaseInvoker) Invoke() (*model.BatchAddRelationsByOneCaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAddRelationsByOneCaseResponse), nil
	}
}

type BatchDeleteTestCaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteTestCaseInvoker) Invoke() (*model.BatchDeleteTestCaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteTestCaseResponse), nil
	}
}

type BatchDeleteTestReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteTestReportInvoker) Invoke() (*model.BatchDeleteTestReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteTestReportResponse), nil
	}
}

type CheckPermissionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckPermissionInvoker) Invoke() (*model.CheckPermissionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckPermissionResponse), nil
	}
}

type CreatePlanInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePlanInvoker) Invoke() (*model.CreatePlanResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePlanResponse), nil
	}
}

type CreateRelationsByOneCaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRelationsByOneCaseInvoker) Invoke() (*model.CreateRelationsByOneCaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRelationsByOneCaseResponse), nil
	}
}

type CreateResourceUriInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateResourceUriInvoker) Invoke() (*model.CreateResourceUriResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateResourceUriResponse), nil
	}
}

type CreateServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateServiceInvoker) Invoke() (*model.CreateServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateServiceResponse), nil
	}
}

type CreateTestCaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTestCaseInvoker) Invoke() (*model.CreateTestCaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTestCaseResponse), nil
	}
}

type CreateTestCaseInPlanInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTestCaseInPlanInvoker) Invoke() (*model.CreateTestCaseInPlanResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTestCaseInPlanResponse), nil
	}
}

type DeleteRelationsByOneCaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRelationsByOneCaseInvoker) Invoke() (*model.DeleteRelationsByOneCaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRelationsByOneCaseResponse), nil
	}
}

type DeleteServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteServiceInvoker) Invoke() (*model.DeleteServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteServiceResponse), nil
	}
}

type ListAllBranchesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllBranchesInvoker) Invoke() (*model.ListAllBranchesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllBranchesResponse), nil
	}
}

type ListAllIteratorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllIteratorsInvoker) Invoke() (*model.ListAllIteratorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllIteratorsResponse), nil
	}
}

type ListAttachmentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAttachmentsInvoker) Invoke() (*model.ListAttachmentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAttachmentsResponse), nil
	}
}

type ListBranchesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBranchesInvoker) Invoke() (*model.ListBranchesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBranchesResponse), nil
	}
}

type ListIssueTreeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIssueTreeInvoker) Invoke() (*model.ListIssueTreeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIssueTreeResponse), nil
	}
}

type ListProjectFieldConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectFieldConfigsInvoker) Invoke() (*model.ListProjectFieldConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectFieldConfigsResponse), nil
	}
}

type ListProjectTestCaseFieldsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectTestCaseFieldsInvoker) Invoke() (*model.ListProjectTestCaseFieldsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectTestCaseFieldsResponse), nil
	}
}

type ListReportsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListReportsInvoker) Invoke() (*model.ListReportsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListReportsResponse), nil
	}
}

type ListResourcePoolsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourcePoolsInvoker) Invoke() (*model.ListResourcePoolsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourcePoolsResponse), nil
	}
}

type ListTaskTestCasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTaskTestCasesInvoker) Invoke() (*model.ListTaskTestCasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTaskTestCasesResponse), nil
	}
}

type ListTestCaseHistoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTestCaseHistoriesInvoker) Invoke() (*model.ListTestCaseHistoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTestCaseHistoriesResponse), nil
	}
}

type ListTestCasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTestCasesInvoker) Invoke() (*model.ListTestCasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTestCasesResponse), nil
	}
}

type ListTestReportsByConditionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTestReportsByConditionInvoker) Invoke() (*model.ListTestReportsByConditionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTestReportsByConditionResponse), nil
	}
}

type ListTestTypesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTestTypesInvoker) Invoke() (*model.ListTestTypesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTestTypesResponse), nil
	}
}

type ListTestcasesByProjectIssuesRelationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTestcasesByProjectIssuesRelationInvoker) Invoke() (*model.ListTestcasesByProjectIssuesRelationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTestcasesByProjectIssuesRelationResponse), nil
	}
}

type ListUsageInfosInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUsageInfosInvoker) Invoke() (*model.ListUsageInfosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUsageInfosResponse), nil
	}
}

type ListUserPackageUsageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserPackageUsageInvoker) Invoke() (*model.ListUserPackageUsageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserPackageUsageResponse), nil
	}
}

type ListUserPopupInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserPopupInfoInvoker) Invoke() (*model.ListUserPopupInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserPopupInfoResponse), nil
	}
}

type RunTestCaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunTestCaseInvoker) Invoke() (*model.RunTestCaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunTestCaseResponse), nil
	}
}

type ShowAllFeatureChildrenInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAllFeatureChildrenInvoker) Invoke() (*model.ShowAllFeatureChildrenResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAllFeatureChildrenResponse), nil
	}
}

type ShowApiTestcaseHistoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApiTestcaseHistoriesInvoker) Invoke() (*model.ShowApiTestcaseHistoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApiTestcaseHistoriesResponse), nil
	}
}

type ShowBackgroundInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBackgroundInfoInvoker) Invoke() (*model.ShowBackgroundInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBackgroundInfoResponse), nil
	}
}

type ShowDisclaimerRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDisclaimerRecordInvoker) Invoke() (*model.ShowDisclaimerRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDisclaimerRecordResponse), nil
	}
}

type ShowDomainInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDomainInfoInvoker) Invoke() (*model.ShowDomainInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainInfoResponse), nil
	}
}

type ShowFeatureChildrenInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFeatureChildrenInvoker) Invoke() (*model.ShowFeatureChildrenResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFeatureChildrenResponse), nil
	}
}

type ShowFreeDeclarationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFreeDeclarationInvoker) Invoke() (*model.ShowFreeDeclarationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFreeDeclarationResponse), nil
	}
}

type ShowIssuesByPlanIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIssuesByPlanIdInvoker) Invoke() (*model.ShowIssuesByPlanIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIssuesByPlanIdResponse), nil
	}
}

type ShowIteratorByDefectInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIteratorByDefectInvoker) Invoke() (*model.ShowIteratorByDefectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIteratorByDefectResponse), nil
	}
}

type ShowMindmapByPageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMindmapByPageInvoker) Invoke() (*model.ShowMindmapByPageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMindmapByPageResponse), nil
	}
}

type ShowMindmapCreatorNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMindmapCreatorNameInvoker) Invoke() (*model.ShowMindmapCreatorNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMindmapCreatorNameResponse), nil
	}
}

type ShowPlanJournalsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPlanJournalsInvoker) Invoke() (*model.ShowPlanJournalsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPlanJournalsResponse), nil
	}
}

type ShowPlanListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPlanListInvoker) Invoke() (*model.ShowPlanListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPlanListResponse), nil
	}
}

type ShowPlansInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPlansInvoker) Invoke() (*model.ShowPlansResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPlansResponse), nil
	}
}

type ShowProgressInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProgressInvoker) Invoke() (*model.ShowProgressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProgressResponse), nil
	}
}

type ShowProjectDataDashboardInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectDataDashboardInvoker) Invoke() (*model.ShowProjectDataDashboardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectDataDashboardResponse), nil
	}
}

type ShowRegisterServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRegisterServiceInvoker) Invoke() (*model.ShowRegisterServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRegisterServiceResponse), nil
	}
}

type ShowReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowReportInvoker) Invoke() (*model.ShowReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowReportResponse), nil
	}
}

type ShowRequirementsOverviewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRequirementsOverviewInvoker) Invoke() (*model.ShowRequirementsOverviewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRequirementsOverviewResponse), nil
	}
}

type ShowSystemConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSystemConfigsInvoker) Invoke() (*model.ShowSystemConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSystemConfigsResponse), nil
	}
}

type ShowTestCaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTestCaseInvoker) Invoke() (*model.ShowTestCaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTestCaseResponse), nil
	}
}

type ShowTestCaseAndDefectInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTestCaseAndDefectInfoInvoker) Invoke() (*model.ShowTestCaseAndDefectInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTestCaseAndDefectInfoResponse), nil
	}
}

type ShowTestCaseDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTestCaseDetailInvoker) Invoke() (*model.ShowTestCaseDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTestCaseDetailResponse), nil
	}
}

type ShowTestCaseDetailV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTestCaseDetailV2Invoker) Invoke() (*model.ShowTestCaseDetailV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTestCaseDetailV2Response), nil
	}
}

type ShowUserAccessInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUserAccessInfoInvoker) Invoke() (*model.ShowUserAccessInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUserAccessInfoResponse), nil
	}
}

type ShowUserExecuteTestCaseInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUserExecuteTestCaseInfoInvoker) Invoke() (*model.ShowUserExecuteTestCaseInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUserExecuteTestCaseInfoResponse), nil
	}
}

type UpdateServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateServiceInvoker) Invoke() (*model.UpdateServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateServiceResponse), nil
	}
}

type UpdateTestCaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTestCaseInvoker) Invoke() (*model.UpdateTestCaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTestCaseResponse), nil
	}
}

type UpdateTestCaseResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTestCaseResultInvoker) Invoke() (*model.UpdateTestCaseResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTestCaseResultResponse), nil
	}
}

type ListAlertGroupsByConditionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlertGroupsByConditionInvoker) Invoke() (*model.ListAlertGroupsByConditionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlertGroupsByConditionResponse), nil
	}
}

type ListAlertTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlertTemplatesInvoker) Invoke() (*model.ListAlertTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlertTemplatesResponse), nil
	}
}

type ListAllConfigItemByTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllConfigItemByTypeInvoker) Invoke() (*model.ListAllConfigItemByTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllConfigItemByTypeResponse), nil
	}
}

type SaveTaskSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *SaveTaskSettingInvoker) Invoke() (*model.SaveTaskSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SaveTaskSettingResponse), nil
	}
}

type ShowAllConfigValueByTypeAndKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAllConfigValueByTypeAndKeyInvoker) Invoke() (*model.ShowAllConfigValueByTypeAndKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAllConfigValueByTypeAndKeyResponse), nil
	}
}

type ShowIfTaskNameRepeatInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIfTaskNameRepeatInvoker) Invoke() (*model.ShowIfTaskNameRepeatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIfTaskNameRepeatResponse), nil
	}
}

type ShowIfUserNameRepeatInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIfUserNameRepeatInvoker) Invoke() (*model.ShowIfUserNameRepeatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIfUserNameRepeatResponse), nil
	}
}

type CreateApiTestSuiteByRepoFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateApiTestSuiteByRepoFileInvoker) Invoke() (*model.CreateApiTestSuiteByRepoFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateApiTestSuiteByRepoFileResponse), nil
	}
}

type ListEnvironmentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEnvironmentsInvoker) Invoke() (*model.ListEnvironmentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEnvironmentsResponse), nil
	}
}

type ListBasicAwInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBasicAwInvoker) Invoke() (*model.ListBasicAwResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBasicAwResponse), nil
	}
}

type ListPublicLibAndAwsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPublicLibAndAwsInvoker) Invoke() (*model.ListPublicLibAndAwsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPublicLibAndAwsResponse), nil
	}
}

type ListUserDnsMappingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserDnsMappingInvoker) Invoke() (*model.ListUserDnsMappingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserDnsMappingResponse), nil
	}
}

type ListVariablesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVariablesInvoker) Invoke() (*model.ListVariablesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVariablesResponse), nil
	}
}
