package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ugo/v1/model"
)

type CheckPermissionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckPermissionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckPermissionInvoker) Invoke() (*model.CheckPermissionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckPermissionResponse), nil
	}
}

type CommitSyntaxConversionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CommitSyntaxConversionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CommitSyntaxConversionInvoker) Invoke() (*model.CommitSyntaxConversionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommitSyntaxConversionResponse), nil
	}
}

type CommitVerificationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CommitVerificationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CommitVerificationInvoker) Invoke() (*model.CommitVerificationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommitVerificationResponse), nil
	}
}

type ConfirmTargetDbTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ConfirmTargetDbTypeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ConfirmTargetDbTypeInvoker) Invoke() (*model.ConfirmTargetDbTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ConfirmTargetDbTypeResponse), nil
	}
}

type CreateEvaluationProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEvaluationProjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEvaluationProjectInvoker) Invoke() (*model.CreateEvaluationProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEvaluationProjectResponse), nil
	}
}

type CreateMigrationProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMigrationProjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMigrationProjectInvoker) Invoke() (*model.CreateMigrationProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMigrationProjectResponse), nil
	}
}

type DeleteEvaluationProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEvaluationProjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEvaluationProjectInvoker) Invoke() (*model.DeleteEvaluationProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEvaluationProjectResponse), nil
	}
}

type DeleteMigrationProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMigrationProjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteMigrationProjectInvoker) Invoke() (*model.DeleteMigrationProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMigrationProjectResponse), nil
	}
}

type DownloadFailureReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadFailureReportInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadFailureReportInvoker) Invoke() (*model.DownloadFailureReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadFailureReportResponse), nil
	}
}

type ListEvaluationProjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEvaluationProjectsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEvaluationProjectsInvoker) Invoke() (*model.ListEvaluationProjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEvaluationProjectsResponse), nil
	}
}

type ListMigrationProjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMigrationProjectsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMigrationProjectsInvoker) Invoke() (*model.ListMigrationProjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMigrationProjectsResponse), nil
	}
}

type ListPermissionCheckResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPermissionCheckResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPermissionCheckResultInvoker) Invoke() (*model.ListPermissionCheckResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPermissionCheckResultResponse), nil
	}
}

type ListQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQuotasInvoker) Invoke() (*model.ListQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotasResponse), nil
	}
}

type ListSyntaxConversionProgressInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSyntaxConversionProgressInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSyntaxConversionProgressInvoker) Invoke() (*model.ListSyntaxConversionProgressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSyntaxConversionProgressResponse), nil
	}
}

type ListVerificationProgressInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVerificationProgressInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVerificationProgressInvoker) Invoke() (*model.ListVerificationProgressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVerificationProgressResponse), nil
	}
}

type ShowEvaluationProjectDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEvaluationProjectDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEvaluationProjectDetailInvoker) Invoke() (*model.ShowEvaluationProjectDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEvaluationProjectDetailResponse), nil
	}
}

type ShowEvaluationProjectStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEvaluationProjectStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEvaluationProjectStatusInvoker) Invoke() (*model.ShowEvaluationProjectStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEvaluationProjectStatusResponse), nil
	}
}

type ShowMigrationProjectDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMigrationProjectDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMigrationProjectDetailInvoker) Invoke() (*model.ShowMigrationProjectDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMigrationProjectDetailResponse), nil
	}
}

type ShowMigrationProjectStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMigrationProjectStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMigrationProjectStatusInvoker) Invoke() (*model.ShowMigrationProjectStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMigrationProjectStatusResponse), nil
	}
}

type ListApiVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApiVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApiVersionsInvoker) Invoke() (*model.ListApiVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApiVersionsResponse), nil
	}
}

type ShowApiVersionInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApiVersionInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowApiVersionInfoInvoker) Invoke() (*model.ShowApiVersionInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApiVersionInfoResponse), nil
	}
}

type RunSqlConversionInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunSqlConversionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunSqlConversionInvoker) Invoke() (*model.RunSqlConversionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunSqlConversionResponse), nil
	}
}
