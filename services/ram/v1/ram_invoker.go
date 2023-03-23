package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ram/v1/model"
)

type AssociateResourceSharePermissionInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListResourceSharePermissionsInvoker) Invoke() (*model.ListResourceSharePermissionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceSharePermissionsResponse), nil
	}
}

type DisableOrganizationShareInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowOrganizationShareInvoker) Invoke() (*model.ShowOrganizationShareResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOrganizationShareResponse), nil
	}
}

type ListPermissionsInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowPermissionInvoker) Invoke() (*model.ShowPermissionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPermissionResponse), nil
	}
}

type SearchDistinctPrincipalsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchDistinctPrincipalsInvoker) Invoke() (*model.SearchDistinctPrincipalsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchDistinctPrincipalsResponse), nil
	}
}

type SearchSharedPrincipalsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchSharedPrincipalsInvoker) Invoke() (*model.SearchSharedPrincipalsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchSharedPrincipalsResponse), nil
	}
}

type SearchDistinctSharedResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchDistinctSharedResourcesInvoker) Invoke() (*model.SearchDistinctSharedResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchDistinctSharedResourcesResponse), nil
	}
}

type SearchSharedResourcesInvoker struct {
	*invoker.BaseInvoker
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

func (i *SearchResourceShareInvitationInvoker) Invoke() (*model.SearchResourceShareInvitationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchResourceShareInvitationResponse), nil
	}
}
