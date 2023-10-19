package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cc/v3/model"
)

type CreateAuthorisationInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListCentralNetworksInvoker) Invoke() (*model.ListCentralNetworksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCentralNetworksResponse), nil
	}
}

type ShowCentralNetworkInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdateCentralNetworkInvoker) Invoke() (*model.UpdateCentralNetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCentralNetworkResponse), nil
	}
}

type CreateCentralNetworkGdgwAttachmentInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListCentralNetworkAttachmentsInvoker) Invoke() (*model.ListCentralNetworkAttachmentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCentralNetworkAttachmentsResponse), nil
	}
}

type ListCentralNetworkGdgwAttachmentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCentralNetworkGdgwAttachmentsInvoker) Invoke() (*model.ListCentralNetworkGdgwAttachmentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCentralNetworkGdgwAttachmentsResponse), nil
	}
}

type ShowCentralNetworkGdgwAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCentralNetworkGdgwAttachmentInvoker) Invoke() (*model.ShowCentralNetworkGdgwAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCentralNetworkGdgwAttachmentResponse), nil
	}
}

type UpdateCentralNetworkGdgwAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCentralNetworkGdgwAttachmentInvoker) Invoke() (*model.UpdateCentralNetworkGdgwAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCentralNetworkGdgwAttachmentResponse), nil
	}
}

type ListCentralNetworkConnectionsInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowCloudConnectionRoutesInvoker) Invoke() (*model.ShowCloudConnectionRoutesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCloudConnectionRoutesResponse), nil
	}
}

type CreateInterRegionBandwidthInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdateNetworkInstanceInvoker) Invoke() (*model.UpdateNetworkInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNetworkInstanceResponse), nil
	}
}
