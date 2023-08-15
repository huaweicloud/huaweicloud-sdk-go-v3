package v1

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/cloudtest/v1/model"
)

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
