package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/servicestage/v2/model"
)

type ChangeApplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeApplicationInvoker) Invoke() (*model.ChangeApplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeApplicationResponse), nil
	}
}

type ChangeApplicationConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeApplicationConfigurationInvoker) Invoke() (*model.ChangeApplicationConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeApplicationConfigurationResponse), nil
	}
}

type ChangeComponentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeComponentInvoker) Invoke() (*model.ChangeComponentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeComponentResponse), nil
	}
}

type ChangeEnvironmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeEnvironmentInvoker) Invoke() (*model.ChangeEnvironmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeEnvironmentResponse), nil
	}
}

type ChangeInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeInstanceInvoker) Invoke() (*model.ChangeInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeInstanceResponse), nil
	}
}

type ChangeResourceInEnvironmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeResourceInEnvironmentInvoker) Invoke() (*model.ChangeResourceInEnvironmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeResourceInEnvironmentResponse), nil
	}
}

type CreateApplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateApplicationInvoker) Invoke() (*model.CreateApplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateApplicationResponse), nil
	}
}

type CreateCamInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCamInstanceInvoker) Invoke() (*model.CreateCamInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCamInstanceResponse), nil
	}
}

type CreateComponentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateComponentInvoker) Invoke() (*model.CreateComponentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateComponentResponse), nil
	}
}

type CreateEnvironmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEnvironmentInvoker) Invoke() (*model.CreateEnvironmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEnvironmentResponse), nil
	}
}

type CreateInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceInvoker) Invoke() (*model.CreateInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceResponse), nil
	}
}

type CreateTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTemplateInvoker) Invoke() (*model.CreateTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTemplateResponse), nil
	}
}

type DeleteApplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteApplicationInvoker) Invoke() (*model.DeleteApplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteApplicationResponse), nil
	}
}

type DeleteApplicationConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteApplicationConfigurationInvoker) Invoke() (*model.DeleteApplicationConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteApplicationConfigurationResponse), nil
	}
}

type DeleteComponentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteComponentInvoker) Invoke() (*model.DeleteComponentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteComponentResponse), nil
	}
}

type DeleteEnvironmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEnvironmentInvoker) Invoke() (*model.DeleteEnvironmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEnvironmentResponse), nil
	}
}

type DeleteInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceInvoker) Invoke() (*model.DeleteInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceResponse), nil
	}
}

type DeleteInstanceByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceByIdInvoker) Invoke() (*model.DeleteInstanceByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceByIdResponse), nil
	}
}

type DeleteTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTemplateInvoker) Invoke() (*model.DeleteTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTemplateResponse), nil
	}
}

type DeployInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeployInstanceInvoker) Invoke() (*model.DeployInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployInstanceResponse), nil
	}
}

type ListApplicationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApplicationsInvoker) Invoke() (*model.ListApplicationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApplicationsResponse), nil
	}
}

type ListComponentOverviewsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListComponentOverviewsInvoker) Invoke() (*model.ListComponentOverviewsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListComponentOverviewsResponse), nil
	}
}

type ListComponentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListComponentsInvoker) Invoke() (*model.ListComponentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListComponentsResponse), nil
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

type ListInstanceSnapshotsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceSnapshotsInvoker) Invoke() (*model.ListInstanceSnapshotsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceSnapshotsResponse), nil
	}
}

type ListInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstancesInvoker) Invoke() (*model.ListInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstancesResponse), nil
	}
}

type ShowApplicationConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApplicationConfigurationInvoker) Invoke() (*model.ShowApplicationConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApplicationConfigurationResponse), nil
	}
}

type ShowApplicationDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApplicationDetailInvoker) Invoke() (*model.ShowApplicationDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApplicationDetailResponse), nil
	}
}

type ShowComponentDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowComponentDetailInvoker) Invoke() (*model.ShowComponentDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowComponentDetailResponse), nil
	}
}

type ShowEnvironmentDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEnvironmentDetailInvoker) Invoke() (*model.ShowEnvironmentDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEnvironmentDetailResponse), nil
	}
}

type ShowInstanceDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceDetailInvoker) Invoke() (*model.ShowInstanceDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceDetailResponse), nil
	}
}

type ShowJobDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobDetailInvoker) Invoke() (*model.ShowJobDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobDetailResponse), nil
	}
}

type UpdateInstanceActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceActionInvoker) Invoke() (*model.UpdateInstanceActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceActionResponse), nil
	}
}

type UpdateTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTemplateInvoker) Invoke() (*model.UpdateTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTemplateResponse), nil
	}
}

type CreateFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFileInvoker) Invoke() (*model.CreateFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFileResponse), nil
	}
}

type CreateHookInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHookInvoker) Invoke() (*model.CreateHookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHookResponse), nil
	}
}

type CreateOAuthInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateOAuthInvoker) Invoke() (*model.CreateOAuthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateOAuthResponse), nil
	}
}

type CreatePasswordAuthInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePasswordAuthInvoker) Invoke() (*model.CreatePasswordAuthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePasswordAuthResponse), nil
	}
}

type CreatePersonalAuthInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePersonalAuthInvoker) Invoke() (*model.CreatePersonalAuthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePersonalAuthResponse), nil
	}
}

type CreateProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateProjectInvoker) Invoke() (*model.CreateProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateProjectResponse), nil
	}
}

type CreateTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTagInvoker) Invoke() (*model.CreateTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTagResponse), nil
	}
}

type DeleteAuthorizeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAuthorizeInvoker) Invoke() (*model.DeleteAuthorizeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAuthorizeResponse), nil
	}
}

type DeleteFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFileInvoker) Invoke() (*model.DeleteFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFileResponse), nil
	}
}

type DeleteHookInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHookInvoker) Invoke() (*model.DeleteHookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHookResponse), nil
	}
}

type DeleteTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTagInvoker) Invoke() (*model.DeleteTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTagResponse), nil
	}
}

type ListAuthorizationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuthorizationsInvoker) Invoke() (*model.ListAuthorizationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuthorizationsResponse), nil
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

type ListCommitsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCommitsInvoker) Invoke() (*model.ListCommitsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCommitsResponse), nil
	}
}

type ListHooksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHooksInvoker) Invoke() (*model.ListHooksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHooksResponse), nil
	}
}

type ListNamespacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNamespacesInvoker) Invoke() (*model.ListNamespacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNamespacesResponse), nil
	}
}

type ListProjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectsInvoker) Invoke() (*model.ListProjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectsResponse), nil
	}
}

type ListTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTagsInvoker) Invoke() (*model.ListTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTagsResponse), nil
	}
}

type ListTreesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTreesInvoker) Invoke() (*model.ListTreesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTreesResponse), nil
	}
}

type ShowContentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowContentInvoker) Invoke() (*model.ShowContentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowContentResponse), nil
	}
}

type ShowProjectDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectDetailInvoker) Invoke() (*model.ShowProjectDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectDetailResponse), nil
	}
}

type ShowRedirectUrlInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRedirectUrlInvoker) Invoke() (*model.ShowRedirectUrlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRedirectUrlResponse), nil
	}
}

type UpdateFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFileInvoker) Invoke() (*model.UpdateFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFileResponse), nil
	}
}

type ListFlavorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlavorsInvoker) Invoke() (*model.ListFlavorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlavorsResponse), nil
	}
}

type ListRuntimesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRuntimesInvoker) Invoke() (*model.ListRuntimesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRuntimesResponse), nil
	}
}

type ListTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTemplatesInvoker) Invoke() (*model.ListTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTemplatesResponse), nil
	}
}
