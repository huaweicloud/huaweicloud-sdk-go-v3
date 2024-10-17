package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpcep/v1/model"
)

type AcceptOrRejectEndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *AcceptOrRejectEndpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AcceptOrRejectEndpointInvoker) Invoke() (*model.AcceptOrRejectEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AcceptOrRejectEndpointResponse), nil
	}
}

type AddOrRemoveServicePermissionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddOrRemoveServicePermissionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddOrRemoveServicePermissionsInvoker) Invoke() (*model.AddOrRemoveServicePermissionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddOrRemoveServicePermissionsResponse), nil
	}
}

type BatchAddEndpointServicePermissionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAddEndpointServicePermissionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchAddEndpointServicePermissionsInvoker) Invoke() (*model.BatchAddEndpointServicePermissionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAddEndpointServicePermissionsResponse), nil
	}
}

type BatchRemoveEndpointServicePermissionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchRemoveEndpointServicePermissionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchRemoveEndpointServicePermissionsInvoker) Invoke() (*model.BatchRemoveEndpointServicePermissionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchRemoveEndpointServicePermissionsResponse), nil
	}
}

type CreateEndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEndpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEndpointInvoker) Invoke() (*model.CreateEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEndpointResponse), nil
	}
}

type CreateEndpointServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEndpointServiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEndpointServiceInvoker) Invoke() (*model.CreateEndpointServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEndpointServiceResponse), nil
	}
}

type DeleteEndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEndpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEndpointInvoker) Invoke() (*model.DeleteEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEndpointResponse), nil
	}
}

type DeleteEndpointPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEndpointPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEndpointPolicyInvoker) Invoke() (*model.DeleteEndpointPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEndpointPolicyResponse), nil
	}
}

type DeleteEndpointServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEndpointServiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEndpointServiceInvoker) Invoke() (*model.DeleteEndpointServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEndpointServiceResponse), nil
	}
}

type ListEndpointInfoDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEndpointInfoDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEndpointInfoDetailsInvoker) Invoke() (*model.ListEndpointInfoDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEndpointInfoDetailsResponse), nil
	}
}

type ListEndpointServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEndpointServiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEndpointServiceInvoker) Invoke() (*model.ListEndpointServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEndpointServiceResponse), nil
	}
}

type ListEndpointsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEndpointsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEndpointsInvoker) Invoke() (*model.ListEndpointsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEndpointsResponse), nil
	}
}

type ListQuotaDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQuotaDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQuotaDetailsInvoker) Invoke() (*model.ListQuotaDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotaDetailsResponse), nil
	}
}

type ListServiceConnectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServiceConnectionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServiceConnectionsInvoker) Invoke() (*model.ListServiceConnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServiceConnectionsResponse), nil
	}
}

type ListServiceDescribeDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServiceDescribeDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServiceDescribeDetailsInvoker) Invoke() (*model.ListServiceDescribeDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServiceDescribeDetailsResponse), nil
	}
}

type ListServiceDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServiceDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServiceDetailsInvoker) Invoke() (*model.ListServiceDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServiceDetailsResponse), nil
	}
}

type ListServicePermissionsDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServicePermissionsDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServicePermissionsDetailsInvoker) Invoke() (*model.ListServicePermissionsDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServicePermissionsDetailsResponse), nil
	}
}

type ListServicePublicDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServicePublicDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServicePublicDetailsInvoker) Invoke() (*model.ListServicePublicDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServicePublicDetailsResponse), nil
	}
}

type ListSpecifiedVersionDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSpecifiedVersionDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSpecifiedVersionDetailsInvoker) Invoke() (*model.ListSpecifiedVersionDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSpecifiedVersionDetailsResponse), nil
	}
}

type ListVersionDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVersionDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVersionDetailsInvoker) Invoke() (*model.ListVersionDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVersionDetailsResponse), nil
	}
}

type UpdateEndpointConnectionsDescInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEndpointConnectionsDescInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEndpointConnectionsDescInvoker) Invoke() (*model.UpdateEndpointConnectionsDescResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEndpointConnectionsDescResponse), nil
	}
}

type UpdateEndpointPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEndpointPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEndpointPolicyInvoker) Invoke() (*model.UpdateEndpointPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEndpointPolicyResponse), nil
	}
}

type UpdateEndpointRoutetableInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEndpointRoutetableInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEndpointRoutetableInvoker) Invoke() (*model.UpdateEndpointRoutetableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEndpointRoutetableResponse), nil
	}
}

type UpdateEndpointServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEndpointServiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEndpointServiceInvoker) Invoke() (*model.UpdateEndpointServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEndpointServiceResponse), nil
	}
}

type UpdateEndpointServiceNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEndpointServiceNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEndpointServiceNameInvoker) Invoke() (*model.UpdateEndpointServiceNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEndpointServiceNameResponse), nil
	}
}

type UpdateEndpointServicePermissionDescInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEndpointServicePermissionDescInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEndpointServicePermissionDescInvoker) Invoke() (*model.UpdateEndpointServicePermissionDescResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEndpointServicePermissionDescResponse), nil
	}
}

type UpdateEndpointWhiteInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEndpointWhiteInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEndpointWhiteInvoker) Invoke() (*model.UpdateEndpointWhiteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEndpointWhiteResponse), nil
	}
}

type BatchAddOrRemoveResourceInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAddOrRemoveResourceInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchAddOrRemoveResourceInstanceInvoker) Invoke() (*model.BatchAddOrRemoveResourceInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAddOrRemoveResourceInstanceResponse), nil
	}
}

type ListQueryProjectResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQueryProjectResourceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQueryProjectResourceTagsInvoker) Invoke() (*model.ListQueryProjectResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQueryProjectResourceTagsResponse), nil
	}
}

type ListResourceInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceInstancesInvoker) Invoke() (*model.ListResourceInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceInstancesResponse), nil
	}
}
