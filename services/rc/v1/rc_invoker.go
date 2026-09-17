package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rc/v1/model"
)

type AddResourcesToGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddResourcesToGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddResourcesToGroupInvoker) Invoke() (*model.AddResourcesToGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddResourcesToGroupResponse), nil
	}
}

type CreateResourceGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateResourceGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateResourceGroupInvoker) Invoke() (*model.CreateResourceGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateResourceGroupResponse), nil
	}
}

type DeleteResourceGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteResourceGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteResourceGroupInvoker) Invoke() (*model.DeleteResourceGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteResourceGroupResponse), nil
	}
}

type ListResourceGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceGroupsInvoker) Invoke() (*model.ListResourceGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceGroupsResponse), nil
	}
}

type RemoveResourceFromGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *RemoveResourceFromGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RemoveResourceFromGroupInvoker) Invoke() (*model.RemoveResourceFromGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RemoveResourceFromGroupResponse), nil
	}
}

type ShowResourceGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResourceGroupInvoker) Invoke() (*model.ShowResourceGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceGroupResponse), nil
	}
}

type UpdateResourceGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateResourceGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateResourceGroupInvoker) Invoke() (*model.UpdateResourceGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateResourceGroupResponse), nil
	}
}

type CreateStoredQueryInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateStoredQueryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateStoredQueryInvoker) Invoke() (*model.CreateStoredQueryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateStoredQueryResponse), nil
	}
}

type DeleteStoredQueryInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteStoredQueryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteStoredQueryInvoker) Invoke() (*model.DeleteStoredQueryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteStoredQueryResponse), nil
	}
}

type ListSchemasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSchemasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSchemasInvoker) Invoke() (*model.ListSchemasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSchemasResponse), nil
	}
}

type ListStoredQueriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStoredQueriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStoredQueriesInvoker) Invoke() (*model.ListStoredQueriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStoredQueriesResponse), nil
	}
}

type RunQueryInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunQueryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunQueryInvoker) Invoke() (*model.RunQueryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunQueryResponse), nil
	}
}

type ShowStoredQueryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStoredQueryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowStoredQueryInvoker) Invoke() (*model.ShowStoredQueryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStoredQueryResponse), nil
	}
}

type UpdateStoredQueryInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateStoredQueryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateStoredQueryInvoker) Invoke() (*model.UpdateStoredQueryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateStoredQueryResponse), nil
	}
}

type ShowResourceRelationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceRelationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResourceRelationsInvoker) Invoke() (*model.ShowResourceRelationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceRelationsResponse), nil
	}
}

type CollectAllResourcesSummaryInvoker struct {
	*invoker.BaseInvoker
}

func (i *CollectAllResourcesSummaryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CollectAllResourcesSummaryInvoker) Invoke() (*model.CollectAllResourcesSummaryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectAllResourcesSummaryResponse), nil
	}
}

type CountAllResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountAllResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountAllResourcesInvoker) Invoke() (*model.CountAllResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountAllResourcesResponse), nil
	}
}

type ListAllProvidersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllProvidersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAllProvidersInvoker) Invoke() (*model.ListAllProvidersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllProvidersResponse), nil
	}
}

type ListAllResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAllResourcesInvoker) Invoke() (*model.ListAllResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllResourcesResponse), nil
	}
}

type ListAllTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAllTagsInvoker) Invoke() (*model.ListAllTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllTagsResponse), nil
	}
}

type ListResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourcesInvoker) Invoke() (*model.ListResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourcesResponse), nil
	}
}

type ShowResourceByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResourceByIdInvoker) Invoke() (*model.ShowResourceByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceByIdResponse), nil
	}
}

type ShowResourceDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResourceDetailInvoker) Invoke() (*model.ShowResourceDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceDetailResponse), nil
	}
}
