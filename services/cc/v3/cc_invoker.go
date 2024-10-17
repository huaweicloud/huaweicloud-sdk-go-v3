package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cc/v3/model"
)

type CreateAuthorisationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAuthorisationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAuthorisationInvoker) Invoke() (*model.CreateAuthorisationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAuthorisationResponse), nil
	}
}

type DeleteAuthorisationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAuthorisationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAuthorisationInvoker) Invoke() (*model.DeleteAuthorisationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAuthorisationResponse), nil
	}
}

type ListAuthorisationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuthorisationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuthorisationsInvoker) Invoke() (*model.ListAuthorisationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuthorisationsResponse), nil
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

type UpdateAuthorisationInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAuthorisationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAuthorisationInvoker) Invoke() (*model.UpdateAuthorisationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAuthorisationResponse), nil
	}
}

type AssociateBandwidthPackageInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateBandwidthPackageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateBandwidthPackageInvoker) Invoke() (*model.AssociateBandwidthPackageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateBandwidthPackageResponse), nil
	}
}

type CreateBandwidthPackageInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateBandwidthPackageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateBandwidthPackageInvoker) Invoke() (*model.CreateBandwidthPackageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateBandwidthPackageResponse), nil
	}
}

type DeleteBandwidthPackageInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBandwidthPackageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteBandwidthPackageInvoker) Invoke() (*model.DeleteBandwidthPackageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBandwidthPackageResponse), nil
	}
}

type DisassociateBandwidthPackageInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateBandwidthPackageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisassociateBandwidthPackageInvoker) Invoke() (*model.DisassociateBandwidthPackageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateBandwidthPackageResponse), nil
	}
}

type ListBandwidthPackageTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBandwidthPackageTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBandwidthPackageTagsInvoker) Invoke() (*model.ListBandwidthPackageTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBandwidthPackageTagsResponse), nil
	}
}

type ListBandwidthPackagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBandwidthPackagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBandwidthPackagesInvoker) Invoke() (*model.ListBandwidthPackagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBandwidthPackagesResponse), nil
	}
}

type ListBandwidthPackagesByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBandwidthPackagesByTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBandwidthPackagesByTagsInvoker) Invoke() (*model.ListBandwidthPackagesByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBandwidthPackagesByTagsResponse), nil
	}
}

type ShowBandwidthPackageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBandwidthPackageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBandwidthPackageInvoker) Invoke() (*model.ShowBandwidthPackageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBandwidthPackageResponse), nil
	}
}

type TagBandwidthPackageInvoker struct {
	*invoker.BaseInvoker
}

func (i *TagBandwidthPackageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *TagBandwidthPackageInvoker) Invoke() (*model.TagBandwidthPackageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.TagBandwidthPackageResponse), nil
	}
}

type UntagBandwidthPackageInvoker struct {
	*invoker.BaseInvoker
}

func (i *UntagBandwidthPackageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UntagBandwidthPackageInvoker) Invoke() (*model.UntagBandwidthPackageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UntagBandwidthPackageResponse), nil
	}
}

type UpdateBandwidthPackageInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateBandwidthPackageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateBandwidthPackageInvoker) Invoke() (*model.UpdateBandwidthPackageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateBandwidthPackageResponse), nil
	}
}

type ApplyCentralNetworkPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ApplyCentralNetworkPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ApplyCentralNetworkPolicyInvoker) Invoke() (*model.ApplyCentralNetworkPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApplyCentralNetworkPolicyResponse), nil
	}
}

type CreateCentralNetworkInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCentralNetworkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCentralNetworkInvoker) Invoke() (*model.CreateCentralNetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCentralNetworkResponse), nil
	}
}

type CreateCentralNetworkPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCentralNetworkPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCentralNetworkPolicyInvoker) Invoke() (*model.CreateCentralNetworkPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCentralNetworkPolicyResponse), nil
	}
}

type DeleteCentralNetworkInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCentralNetworkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCentralNetworkInvoker) Invoke() (*model.DeleteCentralNetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCentralNetworkResponse), nil
	}
}

type DeleteCentralNetworkPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCentralNetworkPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCentralNetworkPolicyInvoker) Invoke() (*model.DeleteCentralNetworkPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCentralNetworkPolicyResponse), nil
	}
}

type ListCentralNetworkPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCentralNetworkPoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCentralNetworkPoliciesInvoker) Invoke() (*model.ListCentralNetworkPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCentralNetworkPoliciesResponse), nil
	}
}

type ListCentralNetworkPolicyChangeSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCentralNetworkPolicyChangeSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCentralNetworkPolicyChangeSetInvoker) Invoke() (*model.ListCentralNetworkPolicyChangeSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCentralNetworkPolicyChangeSetResponse), nil
	}
}

type ListCentralNetworkTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCentralNetworkTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCentralNetworkTagsInvoker) Invoke() (*model.ListCentralNetworkTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCentralNetworkTagsResponse), nil
	}
}

type ListCentralNetworksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCentralNetworksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCentralNetworksInvoker) Invoke() (*model.ListCentralNetworksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCentralNetworksResponse), nil
	}
}

type ListCentralNetworksByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCentralNetworksByTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCentralNetworksByTagsInvoker) Invoke() (*model.ListCentralNetworksByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCentralNetworksByTagsResponse), nil
	}
}

type ShowCentralNetworkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCentralNetworkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCentralNetworkInvoker) Invoke() (*model.ShowCentralNetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCentralNetworkResponse), nil
	}
}

type TagCentralNetworkInvoker struct {
	*invoker.BaseInvoker
}

func (i *TagCentralNetworkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *TagCentralNetworkInvoker) Invoke() (*model.TagCentralNetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.TagCentralNetworkResponse), nil
	}
}

type UntagCentralNetworkInvoker struct {
	*invoker.BaseInvoker
}

func (i *UntagCentralNetworkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UntagCentralNetworkInvoker) Invoke() (*model.UntagCentralNetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UntagCentralNetworkResponse), nil
	}
}

type UpdateCentralNetworkInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCentralNetworkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCentralNetworkInvoker) Invoke() (*model.UpdateCentralNetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCentralNetworkResponse), nil
	}
}

type CreateCentralNetworkErRouteTableAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCentralNetworkErRouteTableAttachmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCentralNetworkErRouteTableAttachmentInvoker) Invoke() (*model.CreateCentralNetworkErRouteTableAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCentralNetworkErRouteTableAttachmentResponse), nil
	}
}

type CreateCentralNetworkGdgwAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCentralNetworkGdgwAttachmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCentralNetworkGdgwAttachmentInvoker) Invoke() (*model.CreateCentralNetworkGdgwAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCentralNetworkGdgwAttachmentResponse), nil
	}
}

type DeleteCentralNetworkAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCentralNetworkAttachmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCentralNetworkAttachmentInvoker) Invoke() (*model.DeleteCentralNetworkAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCentralNetworkAttachmentResponse), nil
	}
}

type ListCentralNetworkAttachmentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCentralNetworkAttachmentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCentralNetworkAttachmentsInvoker) Invoke() (*model.ListCentralNetworkAttachmentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCentralNetworkAttachmentsResponse), nil
	}
}

type ListCentralNetworkErRouteTableAttachmentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCentralNetworkErRouteTableAttachmentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCentralNetworkErRouteTableAttachmentsInvoker) Invoke() (*model.ListCentralNetworkErRouteTableAttachmentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCentralNetworkErRouteTableAttachmentsResponse), nil
	}
}

type ListCentralNetworkGdgwAttachmentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCentralNetworkGdgwAttachmentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCentralNetworkGdgwAttachmentsInvoker) Invoke() (*model.ListCentralNetworkGdgwAttachmentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCentralNetworkGdgwAttachmentsResponse), nil
	}
}

type ShowCentralNetworkErRouteTableAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCentralNetworkErRouteTableAttachmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCentralNetworkErRouteTableAttachmentInvoker) Invoke() (*model.ShowCentralNetworkErRouteTableAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCentralNetworkErRouteTableAttachmentResponse), nil
	}
}

type ShowCentralNetworkGdgwAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCentralNetworkGdgwAttachmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCentralNetworkGdgwAttachmentInvoker) Invoke() (*model.ShowCentralNetworkGdgwAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCentralNetworkGdgwAttachmentResponse), nil
	}
}

type UpdateCentralNetworkErRouteTableAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCentralNetworkErRouteTableAttachmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCentralNetworkErRouteTableAttachmentInvoker) Invoke() (*model.UpdateCentralNetworkErRouteTableAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCentralNetworkErRouteTableAttachmentResponse), nil
	}
}

type UpdateCentralNetworkGdgwAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCentralNetworkGdgwAttachmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCentralNetworkGdgwAttachmentInvoker) Invoke() (*model.UpdateCentralNetworkGdgwAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCentralNetworkGdgwAttachmentResponse), nil
	}
}

type ListCentralNetworkCapabilitiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCentralNetworkCapabilitiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCentralNetworkCapabilitiesInvoker) Invoke() (*model.ListCentralNetworkCapabilitiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCentralNetworkCapabilitiesResponse), nil
	}
}

type ListCentralNetworkConnectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCentralNetworkConnectionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCentralNetworkConnectionsInvoker) Invoke() (*model.ListCentralNetworkConnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCentralNetworkConnectionsResponse), nil
	}
}

type UpdateCentralNetworkConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCentralNetworkConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCentralNetworkConnectionInvoker) Invoke() (*model.UpdateCentralNetworkConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCentralNetworkConnectionResponse), nil
	}
}

type ListCentralNetworkQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCentralNetworkQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCentralNetworkQuotasInvoker) Invoke() (*model.ListCentralNetworkQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCentralNetworkQuotasResponse), nil
	}
}

type CreateCloudConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCloudConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCloudConnectionInvoker) Invoke() (*model.CreateCloudConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCloudConnectionResponse), nil
	}
}

type DeleteCloudConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCloudConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCloudConnectionInvoker) Invoke() (*model.DeleteCloudConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCloudConnectionResponse), nil
	}
}

type ListCloudConnectionTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCloudConnectionTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCloudConnectionTagsInvoker) Invoke() (*model.ListCloudConnectionTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCloudConnectionTagsResponse), nil
	}
}

type ListCloudConnectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCloudConnectionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCloudConnectionsInvoker) Invoke() (*model.ListCloudConnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCloudConnectionsResponse), nil
	}
}

type ListCloudConnectionsByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCloudConnectionsByTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCloudConnectionsByTagsInvoker) Invoke() (*model.ListCloudConnectionsByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCloudConnectionsByTagsResponse), nil
	}
}

type ShowCloudConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCloudConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCloudConnectionInvoker) Invoke() (*model.ShowCloudConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCloudConnectionResponse), nil
	}
}

type TagCloudConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *TagCloudConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *TagCloudConnectionInvoker) Invoke() (*model.TagCloudConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.TagCloudConnectionResponse), nil
	}
}

type UntagCloudConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UntagCloudConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UntagCloudConnectionInvoker) Invoke() (*model.UntagCloudConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UntagCloudConnectionResponse), nil
	}
}

type UpdateCloudConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCloudConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCloudConnectionInvoker) Invoke() (*model.UpdateCloudConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCloudConnectionResponse), nil
	}
}

type ListCloudConnectionQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCloudConnectionQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCloudConnectionQuotasInvoker) Invoke() (*model.ListCloudConnectionQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCloudConnectionQuotasResponse), nil
	}
}

type ListCloudConnectionRoutesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCloudConnectionRoutesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCloudConnectionRoutesInvoker) Invoke() (*model.ListCloudConnectionRoutesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCloudConnectionRoutesResponse), nil
	}
}

type ShowCloudConnectionRoutesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCloudConnectionRoutesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCloudConnectionRoutesInvoker) Invoke() (*model.ShowCloudConnectionRoutesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCloudConnectionRoutesResponse), nil
	}
}

type BatchCreateGcbResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateGcbResourceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateGcbResourceTagsInvoker) Invoke() (*model.BatchCreateGcbResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateGcbResourceTagsResponse), nil
	}
}

type BatchDeleteGcbResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteGcbResourceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteGcbResourceTagsInvoker) Invoke() (*model.BatchDeleteGcbResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteGcbResourceTagsResponse), nil
	}
}

type CountGcbResourceByTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountGcbResourceByTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountGcbResourceByTagInvoker) Invoke() (*model.CountGcbResourceByTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountGcbResourceByTagResponse), nil
	}
}

type CreateGcbResourceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGcbResourceTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateGcbResourceTagInvoker) Invoke() (*model.CreateGcbResourceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGcbResourceTagResponse), nil
	}
}

type DeleteGcbResourceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGcbResourceTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteGcbResourceTagInvoker) Invoke() (*model.DeleteGcbResourceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGcbResourceTagResponse), nil
	}
}

type ListGcbResourceByTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGcbResourceByTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGcbResourceByTagInvoker) Invoke() (*model.ListGcbResourceByTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGcbResourceByTagResponse), nil
	}
}

type ListGcbResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGcbResourceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGcbResourceTagsInvoker) Invoke() (*model.ListGcbResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGcbResourceTagsResponse), nil
	}
}

type ListGcbTenantTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGcbTenantTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGcbTenantTagsInvoker) Invoke() (*model.ListGcbTenantTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGcbTenantTagsResponse), nil
	}
}

type AssociateGlobalConnectionBandwidthInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateGlobalConnectionBandwidthInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateGlobalConnectionBandwidthInstanceInvoker) Invoke() (*model.AssociateGlobalConnectionBandwidthInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateGlobalConnectionBandwidthInstanceResponse), nil
	}
}

type CreateGlobalConnectionBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGlobalConnectionBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateGlobalConnectionBandwidthInvoker) Invoke() (*model.CreateGlobalConnectionBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGlobalConnectionBandwidthResponse), nil
	}
}

type DeleteGlobalConnectionBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGlobalConnectionBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteGlobalConnectionBandwidthInvoker) Invoke() (*model.DeleteGlobalConnectionBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGlobalConnectionBandwidthResponse), nil
	}
}

type DisassociateGlobalConnectionBandwidthInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateGlobalConnectionBandwidthInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisassociateGlobalConnectionBandwidthInstanceInvoker) Invoke() (*model.DisassociateGlobalConnectionBandwidthInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateGlobalConnectionBandwidthInstanceResponse), nil
	}
}

type ListGlobalConnectionBandwidthConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGlobalConnectionBandwidthConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGlobalConnectionBandwidthConfigsInvoker) Invoke() (*model.ListGlobalConnectionBandwidthConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGlobalConnectionBandwidthConfigsResponse), nil
	}
}

type ListGlobalConnectionBandwidthLineLevelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGlobalConnectionBandwidthLineLevelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGlobalConnectionBandwidthLineLevelsInvoker) Invoke() (*model.ListGlobalConnectionBandwidthLineLevelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGlobalConnectionBandwidthLineLevelsResponse), nil
	}
}

type ListGlobalConnectionBandwidthSitesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGlobalConnectionBandwidthSitesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGlobalConnectionBandwidthSitesInvoker) Invoke() (*model.ListGlobalConnectionBandwidthSitesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGlobalConnectionBandwidthSitesResponse), nil
	}
}

type ListGlobalConnectionBandwidthSpecCodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGlobalConnectionBandwidthSpecCodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGlobalConnectionBandwidthSpecCodesInvoker) Invoke() (*model.ListGlobalConnectionBandwidthSpecCodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGlobalConnectionBandwidthSpecCodesResponse), nil
	}
}

type ListGlobalConnectionBandwidthsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGlobalConnectionBandwidthsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGlobalConnectionBandwidthsInvoker) Invoke() (*model.ListGlobalConnectionBandwidthsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGlobalConnectionBandwidthsResponse), nil
	}
}

type ListSupportBindingConnectionBandwidthsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSupportBindingConnectionBandwidthsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSupportBindingConnectionBandwidthsInvoker) Invoke() (*model.ListSupportBindingConnectionBandwidthsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSupportBindingConnectionBandwidthsResponse), nil
	}
}

type ShowGlobalConnectionBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGlobalConnectionBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGlobalConnectionBandwidthInvoker) Invoke() (*model.ShowGlobalConnectionBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGlobalConnectionBandwidthResponse), nil
	}
}

type UpdateGlobalConnectionBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGlobalConnectionBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateGlobalConnectionBandwidthInvoker) Invoke() (*model.UpdateGlobalConnectionBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGlobalConnectionBandwidthResponse), nil
	}
}

type CreateInterRegionBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInterRegionBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInterRegionBandwidthInvoker) Invoke() (*model.CreateInterRegionBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInterRegionBandwidthResponse), nil
	}
}

type DeleteInterRegionBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInterRegionBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInterRegionBandwidthInvoker) Invoke() (*model.DeleteInterRegionBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInterRegionBandwidthResponse), nil
	}
}

type ListInterRegionBandwidthsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInterRegionBandwidthsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInterRegionBandwidthsInvoker) Invoke() (*model.ListInterRegionBandwidthsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInterRegionBandwidthsResponse), nil
	}
}

type ShowInterRegionBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInterRegionBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInterRegionBandwidthInvoker) Invoke() (*model.ShowInterRegionBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInterRegionBandwidthResponse), nil
	}
}

type UpdateInterRegionBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInterRegionBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInterRegionBandwidthInvoker) Invoke() (*model.UpdateInterRegionBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInterRegionBandwidthResponse), nil
	}
}

type CreateNetworkInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNetworkInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateNetworkInstanceInvoker) Invoke() (*model.CreateNetworkInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNetworkInstanceResponse), nil
	}
}

type DeleteNetworkInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNetworkInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteNetworkInstanceInvoker) Invoke() (*model.DeleteNetworkInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNetworkInstanceResponse), nil
	}
}

type ListNetworkInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNetworkInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNetworkInstancesInvoker) Invoke() (*model.ListNetworkInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNetworkInstancesResponse), nil
	}
}

type ShowNetworkInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNetworkInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNetworkInstanceInvoker) Invoke() (*model.ShowNetworkInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNetworkInstanceResponse), nil
	}
}

type UpdateNetworkInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNetworkInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateNetworkInstanceInvoker) Invoke() (*model.UpdateNetworkInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNetworkInstanceResponse), nil
	}
}
