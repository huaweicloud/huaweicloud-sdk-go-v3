package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/devstar/v1/model"
)

type ShowApplicationReleaseRepositoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApplicationReleaseRepositoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowApplicationReleaseRepositoriesInvoker) Invoke() (*model.ShowApplicationReleaseRepositoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApplicationReleaseRepositoriesResponse), nil
	}
}

type ShowApplicationResDeleteStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApplicationResDeleteStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowApplicationResDeleteStatusInvoker) Invoke() (*model.ShowApplicationResDeleteStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApplicationResDeleteStatusResponse), nil
	}
}

type ShowApplicationDependentResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApplicationDependentResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowApplicationDependentResourcesInvoker) Invoke() (*model.ShowApplicationDependentResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApplicationDependentResourcesResponse), nil
	}
}

type ShowApplicationV3Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApplicationV3Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowApplicationV3Invoker) Invoke() (*model.ShowApplicationV3Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApplicationV3Response), nil
	}
}

type UpdateApplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateApplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateApplicationInvoker) Invoke() (*model.UpdateApplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateApplicationResponse), nil
	}
}

type DeleteApplicationV4Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteApplicationV4Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteApplicationV4Invoker) Invoke() (*model.DeleteApplicationV4Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteApplicationV4Response), nil
	}
}

type ListApplicationsV6Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListApplicationsV6Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApplicationsV6Invoker) Invoke() (*model.ListApplicationsV6Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApplicationsV6Response), nil
	}
}

type DownloadApplicationCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadApplicationCodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadApplicationCodeInvoker) Invoke() (*model.DownloadApplicationCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadApplicationCodeResponse), nil
	}
}

type ConfirmDeploymentJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ConfirmDeploymentJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ConfirmDeploymentJobInvoker) Invoke() (*model.ConfirmDeploymentJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ConfirmDeploymentJobResponse), nil
	}
}

type CreateDeploymentJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDeploymentJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDeploymentJobsInvoker) Invoke() (*model.CreateDeploymentJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDeploymentJobsResponse), nil
	}
}

type ShowDeploymentJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeploymentJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDeploymentJobsInvoker) Invoke() (*model.ShowDeploymentJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeploymentJobsResponse), nil
	}
}

type RunCodehubTemplateJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunCodehubTemplateJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunCodehubTemplateJobInvoker) Invoke() (*model.RunCodehubTemplateJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunCodehubTemplateJobResponse), nil
	}
}

type RunDevstarTemplateJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunDevstarTemplateJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunDevstarTemplateJobInvoker) Invoke() (*model.RunDevstarTemplateJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunDevstarTemplateJobResponse), nil
	}
}

type ShowJobDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobDetailInvoker) Invoke() (*model.ShowJobDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobDetailResponse), nil
	}
}

type ListPipelineTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPipelineTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPipelineTemplatesInvoker) Invoke() (*model.ListPipelineTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPipelineTemplatesResponse), nil
	}
}

type ShowPipelineLastStatusV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPipelineLastStatusV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPipelineLastStatusV2Invoker) Invoke() (*model.ShowPipelineLastStatusV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPipelineLastStatusV2Response), nil
	}
}

type StartPipelineInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartPipelineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartPipelineInvoker) Invoke() (*model.StartPipelineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartPipelineResponse), nil
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

type CheckRepositoryDuplicateNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckRepositoryDuplicateNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckRepositoryDuplicateNameInvoker) Invoke() (*model.CheckRepositoryDuplicateNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckRepositoryDuplicateNameResponse), nil
	}
}

type ShowRepositoryByCloudIdeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRepositoryByCloudIdeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRepositoryByCloudIdeInvoker) Invoke() (*model.ShowRepositoryByCloudIdeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRepositoryByCloudIdeResponse), nil
	}
}

type ShowRepositoryStatisticalDataV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRepositoryStatisticalDataV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRepositoryStatisticalDataV2Invoker) Invoke() (*model.ShowRepositoryStatisticalDataV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRepositoryStatisticalDataV2Response), nil
	}
}

type ShowTemplateFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTemplateFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTemplateFileInvoker) Invoke() (*model.ShowTemplateFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTemplateFileResponse), nil
	}
}

type CreateTemplateViewHistoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTemplateViewHistoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTemplateViewHistoriesInvoker) Invoke() (*model.CreateTemplateViewHistoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTemplateViewHistoriesResponse), nil
	}
}

type ListPublishedTemplatesInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListPublishedTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListPublishedTemplatesInvoker) Invoke() (*model.ListPublishedTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPublishedTemplatesResponse), nil
	}
}

type ListTemplateViewHistoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTemplateViewHistoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTemplateViewHistoriesInvoker) Invoke() (*model.ListTemplateViewHistoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTemplateViewHistoriesResponse), nil
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

type ListTemplatesV2Invoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListTemplatesV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListTemplatesV2Invoker) Invoke() (*model.ListTemplatesV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTemplatesV2Response), nil
	}
}

type ShowTemplateV3Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTemplateV3Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTemplateV3Invoker) Invoke() (*model.ShowTemplateV3Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTemplateV3Response), nil
	}
}

type ShowTemplateDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTemplateDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTemplateDetailInvoker) Invoke() (*model.ShowTemplateDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTemplateDetailResponse), nil
	}
}
