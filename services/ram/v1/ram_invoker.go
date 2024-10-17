package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ram/v1/model"
)

type AssociateResourceSharePermissionInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateResourceSharePermissionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateResourceSharePermissionInvoker) Invoke() (*model.AssociateResourceSharePermissionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateResourceSharePermissionResponse), nil
	}
}

type DisassociateResourceSharePermissionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateResourceSharePermissionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisassociateResourceSharePermissionInvoker) Invoke() (*model.DisassociateResourceSharePermissionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateResourceSharePermissionResponse), nil
	}
}

type ListResourceSharePermissionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceSharePermissionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceSharePermissionsInvoker) Invoke() (*model.ListResourceSharePermissionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceSharePermissionsResponse), nil
	}
}

type ListQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQuotaInvoker) Invoke() (*model.ListQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotaResponse), nil
	}
}

type ListResourceTypesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceTypesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceTypesInvoker) Invoke() (*model.ListResourceTypesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceTypesResponse), nil
	}
}

type DisableOrganizationShareInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableOrganizationShareInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisableOrganizationShareInvoker) Invoke() (*model.DisableOrganizationShareResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableOrganizationShareResponse), nil
	}
}

type EnableOrganizationShareInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableOrganizationShareInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableOrganizationShareInvoker) Invoke() (*model.EnableOrganizationShareResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableOrganizationShareResponse), nil
	}
}

type ShowOrganizationShareInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOrganizationShareInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowOrganizationShareInvoker) Invoke() (*model.ShowOrganizationShareResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOrganizationShareResponse), nil
	}
}

type ListPermissionVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPermissionVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPermissionVersionsInvoker) Invoke() (*model.ListPermissionVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPermissionVersionsResponse), nil
	}
}

type ListPermissionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPermissionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPermissionsInvoker) Invoke() (*model.ListPermissionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPermissionsResponse), nil
	}
}

type ShowPermissionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPermissionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPermissionInvoker) Invoke() (*model.ShowPermissionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPermissionResponse), nil
	}
}

type SearchSharedPrincipalsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchSharedPrincipalsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchSharedPrincipalsInvoker) Invoke() (*model.SearchSharedPrincipalsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchSharedPrincipalsResponse), nil
	}
}

type SearchSharedResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchSharedResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchSharedResourcesInvoker) Invoke() (*model.SearchSharedResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchSharedResourcesResponse), nil
	}
}

type CreateResourceShareInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateResourceShareInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateResourceShareInvoker) Invoke() (*model.CreateResourceShareResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateResourceShareResponse), nil
	}
}

type DeleteResourceShareInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteResourceShareInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteResourceShareInvoker) Invoke() (*model.DeleteResourceShareResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteResourceShareResponse), nil
	}
}

type SearchResourceSharesInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchResourceSharesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchResourceSharesInvoker) Invoke() (*model.SearchResourceSharesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchResourceSharesResponse), nil
	}
}

type UpdateResourceShareInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateResourceShareInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateResourceShareInvoker) Invoke() (*model.UpdateResourceShareResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateResourceShareResponse), nil
	}
}

type AssociateResourceShareInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateResourceShareInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateResourceShareInvoker) Invoke() (*model.AssociateResourceShareResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateResourceShareResponse), nil
	}
}

type DisassociateResourceShareInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateResourceShareInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisassociateResourceShareInvoker) Invoke() (*model.DisassociateResourceShareResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateResourceShareResponse), nil
	}
}

type SearchResourceShareAssociationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchResourceShareAssociationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchResourceShareAssociationsInvoker) Invoke() (*model.SearchResourceShareAssociationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchResourceShareAssociationsResponse), nil
	}
}

type AcceptResourceShareInvitationInvoker struct {
	*invoker.BaseInvoker
}

func (i *AcceptResourceShareInvitationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AcceptResourceShareInvitationInvoker) Invoke() (*model.AcceptResourceShareInvitationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AcceptResourceShareInvitationResponse), nil
	}
}

type RejectResourceShareInvitationInvoker struct {
	*invoker.BaseInvoker
}

func (i *RejectResourceShareInvitationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RejectResourceShareInvitationInvoker) Invoke() (*model.RejectResourceShareInvitationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RejectResourceShareInvitationResponse), nil
	}
}

type SearchResourceShareInvitationInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchResourceShareInvitationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchResourceShareInvitationInvoker) Invoke() (*model.SearchResourceShareInvitationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchResourceShareInvitationResponse), nil
	}
}

type BatchCreateResourceShareTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateResourceShareTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateResourceShareTagsInvoker) Invoke() (*model.BatchCreateResourceShareTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateResourceShareTagsResponse), nil
	}
}

type BatchDeleteResourceShareTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteResourceShareTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteResourceShareTagsInvoker) Invoke() (*model.BatchDeleteResourceShareTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteResourceShareTagsResponse), nil
	}
}

type ListResourceShareTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceShareTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceShareTagsInvoker) Invoke() (*model.ListResourceShareTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceShareTagsResponse), nil
	}
}

type ListResourceSharesByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceSharesByTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceSharesByTagsInvoker) Invoke() (*model.ListResourceSharesByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceSharesByTagsResponse), nil
	}
}

type SearchResourceShareCountByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchResourceShareCountByTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchResourceShareCountByTagsInvoker) Invoke() (*model.SearchResourceShareCountByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchResourceShareCountByTagsResponse), nil
	}
}
