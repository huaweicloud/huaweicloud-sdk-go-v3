package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cc/v3/model"
)

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

type BatchCreateDeleteTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateDeleteTagsInvoker) Invoke() (*model.BatchCreateDeleteTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateDeleteTagsResponse), nil
	}
}

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

type ListDomainTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDomainTagsInvoker) Invoke() (*model.ListDomainTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDomainTagsResponse), nil
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

type ListQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQuotasInvoker) Invoke() (*model.ListQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotasResponse), nil
	}
}

type ListResourceByFilterTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceByFilterTagInvoker) Invoke() (*model.ListResourceByFilterTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceByFilterTagResponse), nil
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
