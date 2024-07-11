package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cloudtest/v1/model"
)

type AddTestCaseCommentInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddTestCaseCommentInvoker) Invoke() (*model.AddTestCaseCommentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddTestCaseCommentResponse), nil
	}
}

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

type BatchAddResourcesForIteratorInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAddResourcesForIteratorInvoker) Invoke() (*model.BatchAddResourcesForIteratorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAddResourcesForIteratorResponse), nil
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

type BatchDeleteTestCasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteTestCasesInvoker) Invoke() (*model.BatchDeleteTestCasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteTestCasesResponse), nil
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

type BatchRemoveTestCasesFromIteratorInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchRemoveTestCasesFromIteratorInvoker) Invoke() (*model.BatchRemoveTestCasesFromIteratorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchRemoveTestCasesFromIteratorResponse), nil
	}
}

type BatchUpdateVersionTestCasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateVersionTestCasesInvoker) Invoke() (*model.BatchUpdateVersionTestCasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateVersionTestCasesResponse), nil
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

type CreateIteratorInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateIteratorInvoker) Invoke() (*model.CreateIteratorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateIteratorResponse), nil
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

type CreateProjectBranchInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateProjectBranchInvoker) Invoke() (*model.CreateProjectBranchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateProjectBranchResponse), nil
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

type CreateReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateReportInvoker) Invoke() (*model.CreateReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateReportResponse), nil
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

type CreateUserDefinedUrlKeyWordInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateUserDefinedUrlKeyWordInvoker) Invoke() (*model.CreateUserDefinedUrlKeyWordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateUserDefinedUrlKeyWordResponse), nil
	}
}

type CreateVersionTestCaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVersionTestCaseInvoker) Invoke() (*model.CreateVersionTestCaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVersionTestCaseResponse), nil
	}
}

type DeleteBasicAwByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBasicAwByIdInvoker) Invoke() (*model.DeleteBasicAwByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBasicAwByIdResponse), nil
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

type DeleteTestCaseCommentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTestCaseCommentInvoker) Invoke() (*model.DeleteTestCaseCommentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTestCaseCommentResponse), nil
	}
}

type ListAlarmStatisticsUsingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlarmStatisticsUsingInvoker) Invoke() (*model.ListAlarmStatisticsUsingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlarmStatisticsUsingResponse), nil
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

type ListAllTestCasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllTestCasesInvoker) Invoke() (*model.ListAllTestCasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllTestCasesResponse), nil
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

type ListAvailableConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAvailableConfigInvoker) Invoke() (*model.ListAvailableConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAvailableConfigResponse), nil
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

type ListBasicAwInfoListSupportsSearchInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBasicAwInfoListSupportsSearchInvoker) Invoke() (*model.ListBasicAwInfoListSupportsSearchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBasicAwInfoListSupportsSearchResponse), nil
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

type ListCasesStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCasesStatusInvoker) Invoke() (*model.ListCasesStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCasesStatusResponse), nil
	}
}

type ListDomainVisibleServicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDomainVisibleServicesInvoker) Invoke() (*model.ListDomainVisibleServicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDomainVisibleServicesResponse), nil
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

type ListIteratorIssueTreeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIteratorIssueTreeInvoker) Invoke() (*model.ListIteratorIssueTreeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIteratorIssueTreeResponse), nil
	}
}

type ListIteratorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIteratorsInvoker) Invoke() (*model.ListIteratorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIteratorsResponse), nil
	}
}

type ListLinesUsingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLinesUsingInvoker) Invoke() (*model.ListLinesUsingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLinesUsingResponse), nil
	}
}

type ListMsgInfosUsingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMsgInfosUsingInvoker) Invoke() (*model.ListMsgInfosUsingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMsgInfosUsingResponse), nil
	}
}

type ListOwnTestCasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOwnTestCasesInvoker) Invoke() (*model.ListOwnTestCasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOwnTestCasesResponse), nil
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

type ListScattersUsingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScattersUsingInvoker) Invoke() (*model.ListScattersUsingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScattersUsingResponse), nil
	}
}

type ListSubTaskCaseOverstockUsingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubTaskCaseOverstockUsingInvoker) Invoke() (*model.ListSubTaskCaseOverstockUsingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubTaskCaseOverstockUsingResponse), nil
	}
}

type ListTaskAssignCasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTaskAssignCasesInvoker) Invoke() (*model.ListTaskAssignCasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTaskAssignCasesResponse), nil
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

type ListTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTasksInvoker) Invoke() (*model.ListTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTasksResponse), nil
	}
}

type ListTestCaseCommentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTestCaseCommentsInvoker) Invoke() (*model.ListTestCaseCommentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTestCaseCommentsResponse), nil
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

type ListTestCaseScriptDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTestCaseScriptDetailInvoker) Invoke() (*model.ListTestCaseScriptDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTestCaseScriptDetailResponse), nil
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

type ListTestCasesByIssueInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTestCasesByIssueInvoker) Invoke() (*model.ListTestCasesByIssueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTestCasesByIssueResponse), nil
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

type ListUsingGetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUsingGetInvoker) Invoke() (*model.ListUsingGetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUsingGetResponse), nil
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

type RemoveIssuesFromIteratorInvoker struct {
	*invoker.BaseInvoker
}

func (i *RemoveIssuesFromIteratorInvoker) Invoke() (*model.RemoveIssuesFromIteratorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RemoveIssuesFromIteratorResponse), nil
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

type ShowAssetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAssetInvoker) Invoke() (*model.ShowAssetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAssetResponse), nil
	}
}

type ShowAssetTreeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAssetTreeInvoker) Invoke() (*model.ShowAssetTreeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAssetTreeResponse), nil
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

type ShowBranchInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBranchInvoker) Invoke() (*model.ShowBranchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBranchResponse), nil
	}
}

type ShowConcurrencyPackageUsingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConcurrencyPackageUsingInvoker) Invoke() (*model.ShowConcurrencyPackageUsingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConcurrencyPackageUsingResponse), nil
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

type ShowEchoTestPackageUsingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEchoTestPackageUsingInvoker) Invoke() (*model.ShowEchoTestPackageUsingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEchoTestPackageUsingResponse), nil
	}
}

type ShowFactorByAssetIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFactorByAssetIdInvoker) Invoke() (*model.ShowFactorByAssetIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFactorByAssetIdResponse), nil
	}
}

type ShowFactorByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFactorByIdInvoker) Invoke() (*model.ShowFactorByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFactorByIdResponse), nil
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

type ShowIteratorDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIteratorDetailInvoker) Invoke() (*model.ShowIteratorDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIteratorDetailResponse), nil
	}
}

type ShowMindMapByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMindMapByIdInvoker) Invoke() (*model.ShowMindMapByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMindMapByIdResponse), nil
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

type ShowOperationalDataCurrentMonthUsingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOperationalDataCurrentMonthUsingInvoker) Invoke() (*model.ShowOperationalDataCurrentMonthUsingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOperationalDataCurrentMonthUsingResponse), nil
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

type ShowReviewByPageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowReviewByPageInvoker) Invoke() (*model.ShowReviewByPageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowReviewByPageResponse), nil
	}
}

type ShowSceneByPageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSceneByPageInvoker) Invoke() (*model.ShowSceneByPageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSceneByPageResponse), nil
	}
}

type ShowStatisticByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStatisticByIdInvoker) Invoke() (*model.ShowStatisticByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStatisticByIdResponse), nil
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

type ShowTemplateByPageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTemplateByPageInvoker) Invoke() (*model.ShowTemplateByPageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTemplateByPageResponse), nil
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

type ShowTestCaseReviewsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTestCaseReviewsInvoker) Invoke() (*model.ShowTestCaseReviewsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTestCaseReviewsResponse), nil
	}
}

type ShowTestCasesChangeStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTestCasesChangeStatisticsInvoker) Invoke() (*model.ShowTestCasesChangeStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTestCasesChangeStatisticsResponse), nil
	}
}

type ShowTestcaseByPageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTestcaseByPageInvoker) Invoke() (*model.ShowTestcaseByPageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTestcaseByPageResponse), nil
	}
}

type ShowTestpointByPageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTestpointByPageInvoker) Invoke() (*model.ShowTestpointByPageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTestpointByPageResponse), nil
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

type UpdateBasicAwByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateBasicAwByIdInvoker) Invoke() (*model.UpdateBasicAwByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateBasicAwByIdResponse), nil
	}
}

type UpdateIteratorInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateIteratorInvoker) Invoke() (*model.UpdateIteratorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateIteratorResponse), nil
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

type UpdateTestCaseAndScriptInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTestCaseAndScriptInvoker) Invoke() (*model.UpdateTestCaseAndScriptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTestCaseAndScriptResponse), nil
	}
}

type UpdateTestCaseCommentInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTestCaseCommentInvoker) Invoke() (*model.UpdateTestCaseCommentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTestCaseCommentResponse), nil
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

type UpdateUserDnsMappingInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateUserDnsMappingInvoker) Invoke() (*model.UpdateUserDnsMappingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateUserDnsMappingResponse), nil
	}
}

type UpdateVersionTestCaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVersionTestCaseInvoker) Invoke() (*model.UpdateVersionTestCaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVersionTestCaseResponse), nil
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
