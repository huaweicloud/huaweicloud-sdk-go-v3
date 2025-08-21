package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cfw/v1/model"
)

type AddAddressItemInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddAddressItemInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddAddressItemInvoker) Invoke() (*model.AddAddressItemResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddAddressItemResponse), nil
	}
}

type AddAddressSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddAddressSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddAddressSetInvoker) Invoke() (*model.AddAddressSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddAddressSetResponse), nil
	}
}

type AddBlackWhiteListInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddBlackWhiteListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddBlackWhiteListInvoker) Invoke() (*model.AddBlackWhiteListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddBlackWhiteListResponse), nil
	}
}

type AddDomainSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddDomainSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddDomainSetInvoker) Invoke() (*model.AddDomainSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddDomainSetResponse), nil
	}
}

type AddDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddDomainsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddDomainsInvoker) Invoke() (*model.AddDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddDomainsResponse), nil
	}
}

type AddLogConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddLogConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddLogConfigInvoker) Invoke() (*model.AddLogConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddLogConfigResponse), nil
	}
}

type AddServiceItemsInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddServiceItemsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddServiceItemsInvoker) Invoke() (*model.AddServiceItemsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddServiceItemsResponse), nil
	}
}

type AddServiceSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddServiceSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddServiceSetInvoker) Invoke() (*model.AddServiceSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddServiceSetResponse), nil
	}
}

type BatchDeleteAddressItemsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteAddressItemsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteAddressItemsInvoker) Invoke() (*model.BatchDeleteAddressItemsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteAddressItemsResponse), nil
	}
}

type BatchDeleteDomainSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteDomainSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteDomainSetInvoker) Invoke() (*model.BatchDeleteDomainSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteDomainSetResponse), nil
	}
}

type BatchDeleteServiceItemsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteServiceItemsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteServiceItemsInvoker) Invoke() (*model.BatchDeleteServiceItemsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteServiceItemsResponse), nil
	}
}

type CancelCaptureTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelCaptureTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelCaptureTaskInvoker) Invoke() (*model.CancelCaptureTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelCaptureTaskResponse), nil
	}
}

type CreateCaptureTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCaptureTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCaptureTaskInvoker) Invoke() (*model.CreateCaptureTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCaptureTaskResponse), nil
	}
}

type CreateEastWestFirewallInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEastWestFirewallInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEastWestFirewallInvoker) Invoke() (*model.CreateEastWestFirewallResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEastWestFirewallResponse), nil
	}
}

type CreateFirewallInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFirewallInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFirewallInvoker) Invoke() (*model.CreateFirewallResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFirewallResponse), nil
	}
}

type CreateTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTagInvoker) Invoke() (*model.CreateTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTagResponse), nil
	}
}

type DeleteAddressItemInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAddressItemInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAddressItemInvoker) Invoke() (*model.DeleteAddressItemResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAddressItemResponse), nil
	}
}

type DeleteAddressSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAddressSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAddressSetInvoker) Invoke() (*model.DeleteAddressSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAddressSetResponse), nil
	}
}

type DeleteBlackWhiteListInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBlackWhiteListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteBlackWhiteListInvoker) Invoke() (*model.DeleteBlackWhiteListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBlackWhiteListResponse), nil
	}
}

type DeleteCaptureTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCaptureTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCaptureTaskInvoker) Invoke() (*model.DeleteCaptureTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCaptureTaskResponse), nil
	}
}

type DeleteDomainSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDomainSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDomainSetInvoker) Invoke() (*model.DeleteDomainSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDomainSetResponse), nil
	}
}

type DeleteDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDomainsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDomainsInvoker) Invoke() (*model.DeleteDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDomainsResponse), nil
	}
}

type DeleteFirewallInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFirewallInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteFirewallInvoker) Invoke() (*model.DeleteFirewallResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFirewallResponse), nil
	}
}

type DeleteIpBlacklistInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteIpBlacklistInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteIpBlacklistInvoker) Invoke() (*model.DeleteIpBlacklistResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteIpBlacklistResponse), nil
	}
}

type DeleteServiceItemInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteServiceItemInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteServiceItemInvoker) Invoke() (*model.DeleteServiceItemResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteServiceItemResponse), nil
	}
}

type DeleteServiceSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteServiceSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteServiceSetInvoker) Invoke() (*model.DeleteServiceSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteServiceSetResponse), nil
	}
}

type DeleteTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTagInvoker) Invoke() (*model.DeleteTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTagResponse), nil
	}
}

type EnableIpBlacklistInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableIpBlacklistInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableIpBlacklistInvoker) Invoke() (*model.EnableIpBlacklistResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableIpBlacklistResponse), nil
	}
}

type ExportIpBlacklistInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportIpBlacklistInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportIpBlacklistInvoker) Invoke() (*model.ExportIpBlacklistResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportIpBlacklistResponse), nil
	}
}

type ImportIpBlacklistInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportIpBlacklistInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportIpBlacklistInvoker) Invoke() (*model.ImportIpBlacklistResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportIpBlacklistResponse), nil
	}
}

type ListAccessControlLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccessControlLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAccessControlLogsInvoker) Invoke() (*model.ListAccessControlLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccessControlLogsResponse), nil
	}
}

type ListAddressItemsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAddressItemsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAddressItemsInvoker) Invoke() (*model.ListAddressItemsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAddressItemsResponse), nil
	}
}

type ListAddressSetDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAddressSetDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAddressSetDetailInvoker) Invoke() (*model.ListAddressSetDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAddressSetDetailResponse), nil
	}
}

type ListAddressSetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAddressSetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAddressSetsInvoker) Invoke() (*model.ListAddressSetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAddressSetsResponse), nil
	}
}

type ListAttackLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAttackLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAttackLogsInvoker) Invoke() (*model.ListAttackLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAttackLogsResponse), nil
	}
}

type ListBlackWhiteListsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBlackWhiteListsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBlackWhiteListsInvoker) Invoke() (*model.ListBlackWhiteListsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBlackWhiteListsResponse), nil
	}
}

type ListCaptureResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCaptureResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCaptureResultInvoker) Invoke() (*model.ListCaptureResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCaptureResultResponse), nil
	}
}

type ListCaptureTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCaptureTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCaptureTaskInvoker) Invoke() (*model.ListCaptureTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCaptureTaskResponse), nil
	}
}

type ListDnsServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDnsServersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDnsServersInvoker) Invoke() (*model.ListDnsServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDnsServersResponse), nil
	}
}

type ListDomainParseDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDomainParseDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDomainParseDetailInvoker) Invoke() (*model.ListDomainParseDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDomainParseDetailResponse), nil
	}
}

type ListDomainParseIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDomainParseIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDomainParseIpInvoker) Invoke() (*model.ListDomainParseIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDomainParseIpResponse), nil
	}
}

type ListDomainSetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDomainSetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDomainSetsInvoker) Invoke() (*model.ListDomainSetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDomainSetsResponse), nil
	}
}

type ListDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDomainsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDomainsInvoker) Invoke() (*model.ListDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDomainsResponse), nil
	}
}

type ListEastWestFirewallInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEastWestFirewallInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEastWestFirewallInvoker) Invoke() (*model.ListEastWestFirewallResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEastWestFirewallResponse), nil
	}
}

type ListFirewallDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFirewallDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFirewallDetailInvoker) Invoke() (*model.ListFirewallDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFirewallDetailResponse), nil
	}
}

type ListFirewallListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFirewallListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFirewallListInvoker) Invoke() (*model.ListFirewallListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFirewallListResponse), nil
	}
}

type ListFlowLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlowLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFlowLogsInvoker) Invoke() (*model.ListFlowLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlowLogsResponse), nil
	}
}

type ListIpBlacklistInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIpBlacklistInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIpBlacklistInvoker) Invoke() (*model.ListIpBlacklistResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIpBlacklistResponse), nil
	}
}

type ListIpBlacklistSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIpBlacklistSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIpBlacklistSwitchInvoker) Invoke() (*model.ListIpBlacklistSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIpBlacklistSwitchResponse), nil
	}
}

type ListJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListJobInvoker) Invoke() (*model.ListJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJobResponse), nil
	}
}

type ListLogConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLogConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLogConfigInvoker) Invoke() (*model.ListLogConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLogConfigResponse), nil
	}
}

type ListProjectTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectTagsInvoker) Invoke() (*model.ListProjectTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectTagsResponse), nil
	}
}

type ListProtectedVpcsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProtectedVpcsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProtectedVpcsInvoker) Invoke() (*model.ListProtectedVpcsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProtectedVpcsResponse), nil
	}
}

type ListResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceTagsInvoker) Invoke() (*model.ListResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceTagsResponse), nil
	}
}

type ListServiceItemsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServiceItemsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServiceItemsInvoker) Invoke() (*model.ListServiceItemsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServiceItemsResponse), nil
	}
}

type ListServiceSetDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServiceSetDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServiceSetDetailInvoker) Invoke() (*model.ListServiceSetDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServiceSetDetailResponse), nil
	}
}

type ListServiceSetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServiceSetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServiceSetsInvoker) Invoke() (*model.ListServiceSetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServiceSetsResponse), nil
	}
}

type RetryIpBlacklistInvoker struct {
	*invoker.BaseInvoker
}

func (i *RetryIpBlacklistInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RetryIpBlacklistInvoker) Invoke() (*model.RetryIpBlacklistResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetryIpBlacklistResponse), nil
	}
}

type SaveTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SaveTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SaveTagsInvoker) Invoke() (*model.SaveTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SaveTagsResponse), nil
	}
}

type ShowAlarmConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAlarmConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAlarmConfigInvoker) Invoke() (*model.ShowAlarmConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAlarmConfigResponse), nil
	}
}

type ShowAntiVirusRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAntiVirusRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAntiVirusRuleInvoker) Invoke() (*model.ShowAntiVirusRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAntiVirusRuleResponse), nil
	}
}

type ShowAntiVirusSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAntiVirusSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAntiVirusSwitchInvoker) Invoke() (*model.ShowAntiVirusSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAntiVirusSwitchResponse), nil
	}
}

type ShowDomainSetDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDomainSetDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDomainSetDetailInvoker) Invoke() (*model.ShowDomainSetDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainSetDetailResponse), nil
	}
}

type UpdateAddressSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAddressSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAddressSetInvoker) Invoke() (*model.UpdateAddressSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAddressSetResponse), nil
	}
}

type UpdateAlarmConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAlarmConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAlarmConfigInvoker) Invoke() (*model.UpdateAlarmConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAlarmConfigResponse), nil
	}
}

type UpdateAntiVirusRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAntiVirusRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAntiVirusRuleInvoker) Invoke() (*model.UpdateAntiVirusRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAntiVirusRuleResponse), nil
	}
}

type UpdateAntiVirusSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAntiVirusSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAntiVirusSwitchInvoker) Invoke() (*model.UpdateAntiVirusSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAntiVirusSwitchResponse), nil
	}
}

type UpdateBlackWhiteListInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateBlackWhiteListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateBlackWhiteListInvoker) Invoke() (*model.UpdateBlackWhiteListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateBlackWhiteListResponse), nil
	}
}

type UpdateDnsServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDnsServersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDnsServersInvoker) Invoke() (*model.UpdateDnsServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDnsServersResponse), nil
	}
}

type UpdateDomainSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDomainSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDomainSetInvoker) Invoke() (*model.UpdateDomainSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDomainSetResponse), nil
	}
}

type UpdateLogConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLogConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateLogConfigInvoker) Invoke() (*model.UpdateLogConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLogConfigResponse), nil
	}
}

type UpdateObjectConfigDescInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateObjectConfigDescInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateObjectConfigDescInvoker) Invoke() (*model.UpdateObjectConfigDescResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateObjectConfigDescResponse), nil
	}
}

type UpdateServiceSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateServiceSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateServiceSetInvoker) Invoke() (*model.UpdateServiceSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateServiceSetResponse), nil
	}
}

type AddAclRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddAclRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddAclRuleInvoker) Invoke() (*model.AddAclRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddAclRuleResponse), nil
	}
}

type BatchDeleteAclRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteAclRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteAclRulesInvoker) Invoke() (*model.BatchDeleteAclRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteAclRulesResponse), nil
	}
}

type BatchUpdateAclRuleActionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateAclRuleActionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpdateAclRuleActionsInvoker) Invoke() (*model.BatchUpdateAclRuleActionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateAclRuleActionsResponse), nil
	}
}

type DeleteAclRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAclRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAclRuleInvoker) Invoke() (*model.DeleteAclRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAclRuleResponse), nil
	}
}

type DeleteAclRuleHitCountInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAclRuleHitCountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAclRuleHitCountInvoker) Invoke() (*model.DeleteAclRuleHitCountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAclRuleHitCountResponse), nil
	}
}

type ListAclRuleHitCountInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAclRuleHitCountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAclRuleHitCountInvoker) Invoke() (*model.ListAclRuleHitCountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAclRuleHitCountResponse), nil
	}
}

type ListAclRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAclRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAclRulesInvoker) Invoke() (*model.ListAclRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAclRulesResponse), nil
	}
}

type ListRegionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRegionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRegionsInvoker) Invoke() (*model.ListRegionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRegionsResponse), nil
	}
}

type ListRuleAclTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRuleAclTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRuleAclTagsInvoker) Invoke() (*model.ListRuleAclTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRuleAclTagsResponse), nil
	}
}

type ShowImportStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowImportStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowImportStatusInvoker) Invoke() (*model.ShowImportStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowImportStatusResponse), nil
	}
}

type UpdateAclRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAclRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAclRuleInvoker) Invoke() (*model.UpdateAclRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAclRuleResponse), nil
	}
}

type UpdateAclRuleOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAclRuleOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAclRuleOrderInvoker) Invoke() (*model.UpdateAclRuleOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAclRuleOrderResponse), nil
	}
}

type ChangeEipStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeEipStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeEipStatusInvoker) Invoke() (*model.ChangeEipStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeEipStatusResponse), nil
	}
}

type ListAlarmWhitelistInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlarmWhitelistInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAlarmWhitelistInvoker) Invoke() (*model.ListAlarmWhitelistResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlarmWhitelistResponse), nil
	}
}

type ListEipCountInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEipCountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEipCountInvoker) Invoke() (*model.ListEipCountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEipCountResponse), nil
	}
}

type ListEipsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEipsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEipsInvoker) Invoke() (*model.ListEipsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEipsResponse), nil
	}
}

type ShowAutoProtectStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAutoProtectStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAutoProtectStatusInvoker) Invoke() (*model.ShowAutoProtectStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAutoProtectStatusResponse), nil
	}
}

type SwitchAutoProtectStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchAutoProtectStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchAutoProtectStatusInvoker) Invoke() (*model.SwitchAutoProtectStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchAutoProtectStatusResponse), nil
	}
}

type ListCustomerIpsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCustomerIpsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCustomerIpsInvoker) Invoke() (*model.ListCustomerIpsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCustomerIpsResponse), nil
	}
}

type ShowCustomerIpsInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCustomerIpsInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCustomerIpsInfoInvoker) Invoke() (*model.ShowCustomerIpsInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCustomerIpsInfoResponse), nil
	}
}

type UpdateCustomerIpsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCustomerIpsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCustomerIpsInvoker) Invoke() (*model.UpdateCustomerIpsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCustomerIpsResponse), nil
	}
}

type ChangeIpsProtectModeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeIpsProtectModeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeIpsProtectModeInvoker) Invoke() (*model.ChangeIpsProtectModeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeIpsProtectModeResponse), nil
	}
}

type ChangeIpsRuleModeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeIpsRuleModeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeIpsRuleModeInvoker) Invoke() (*model.ChangeIpsRuleModeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeIpsRuleModeResponse), nil
	}
}

type ChangeIpsSwitchStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeIpsSwitchStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeIpsSwitchStatusInvoker) Invoke() (*model.ChangeIpsSwitchStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeIpsSwitchStatusResponse), nil
	}
}

type ListIpsProtectModeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIpsProtectModeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIpsProtectModeInvoker) Invoke() (*model.ListIpsProtectModeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIpsProtectModeResponse), nil
	}
}

type ListIpsRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIpsRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIpsRulesInvoker) Invoke() (*model.ListIpsRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIpsRulesResponse), nil
	}
}

type ListIpsRules1Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListIpsRules1Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIpsRules1Invoker) Invoke() (*model.ListIpsRules1Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIpsRules1Response), nil
	}
}

type ListIpsSwitchStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIpsSwitchStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIpsSwitchStatusInvoker) Invoke() (*model.ListIpsSwitchStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIpsSwitchStatusResponse), nil
	}
}

type ShowIpsUpdateTimeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIpsUpdateTimeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowIpsUpdateTimeInvoker) Invoke() (*model.ShowIpsUpdateTimeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIpsUpdateTimeResponse), nil
	}
}

type UpdateAdvancedIpsRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAdvancedIpsRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAdvancedIpsRuleInvoker) Invoke() (*model.UpdateAdvancedIpsRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAdvancedIpsRuleResponse), nil
	}
}

type ListAttackStatisticInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAttackStatisticInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAttackStatisticInvoker) Invoke() (*model.ListAttackStatisticResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAttackStatisticResponse), nil
	}
}

type ListFlowStatisticInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlowStatisticInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFlowStatisticInvoker) Invoke() (*model.ListFlowStatisticResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlowStatisticResponse), nil
	}
}

type ShowAccessDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAccessDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAccessDetailInvoker) Invoke() (*model.ShowAccessDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAccessDetailResponse), nil
	}
}

type ShowAccessTopInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAccessTopInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAccessTopInvoker) Invoke() (*model.ShowAccessTopResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAccessTopResponse), nil
	}
}

type ShowAttackDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAttackDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAttackDetailInvoker) Invoke() (*model.ShowAttackDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAttackDetailResponse), nil
	}
}

type ShowAttackTopInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAttackTopInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAttackTopInvoker) Invoke() (*model.ShowAttackTopResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAttackTopResponse), nil
	}
}

type ShowAttackTotalInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAttackTotalInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAttackTotalInvoker) Invoke() (*model.ShowAttackTotalResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAttackTotalResponse), nil
	}
}

type ShowAttackTrendInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAttackTrendInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAttackTrendInvoker) Invoke() (*model.ShowAttackTrendResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAttackTrendResponse), nil
	}
}

type ShowFlowDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFlowDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFlowDetailInvoker) Invoke() (*model.ShowFlowDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFlowDetailResponse), nil
	}
}

type ShowFlowTopInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFlowTopInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFlowTopInvoker) Invoke() (*model.ShowFlowTopResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFlowTopResponse), nil
	}
}

type ShowFlowTrendInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFlowTrendInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFlowTrendInvoker) Invoke() (*model.ShowFlowTrendResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFlowTrendResponse), nil
	}
}

type ShowLogsCountInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLogsCountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLogsCountInvoker) Invoke() (*model.ShowLogsCountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLogsCountResponse), nil
	}
}

type ShowTrafficTrendInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTrafficTrendInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTrafficTrendInvoker) Invoke() (*model.ShowTrafficTrendResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTrafficTrendResponse), nil
	}
}

type ExportLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportLogsInvoker) Invoke() (*model.ExportLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportLogsResponse), nil
	}
}

type ListLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLogsInvoker) Invoke() (*model.ListLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLogsResponse), nil
	}
}

type EnableMultiAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableMultiAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableMultiAccountInvoker) Invoke() (*model.EnableMultiAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableMultiAccountResponse), nil
	}
}

type ListAccountsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccountsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAccountsInvoker) Invoke() (*model.ListAccountsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccountsResponse), nil
	}
}

type ListOrganizationAccountsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOrganizationAccountsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListOrganizationAccountsInvoker) Invoke() (*model.ListOrganizationAccountsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOrganizationAccountsResponse), nil
	}
}

type ListOrganizationTreeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOrganizationTreeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListOrganizationTreeInvoker) Invoke() (*model.ListOrganizationTreeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOrganizationTreeResponse), nil
	}
}

type CreateReportProfileInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateReportProfileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateReportProfileInvoker) Invoke() (*model.CreateReportProfileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateReportProfileResponse), nil
	}
}

type DeleteReportProfileInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteReportProfileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteReportProfileInvoker) Invoke() (*model.DeleteReportProfileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteReportProfileResponse), nil
	}
}

type ListReportProfilesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListReportProfilesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListReportProfilesInvoker) Invoke() (*model.ListReportProfilesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListReportProfilesResponse), nil
	}
}

type ShowFirewallReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFirewallReportInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFirewallReportInvoker) Invoke() (*model.ShowFirewallReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFirewallReportResponse), nil
	}
}

type ShowReportProfileInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowReportProfileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowReportProfileInvoker) Invoke() (*model.ShowReportProfileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowReportProfileResponse), nil
	}
}

type UpdateReportProfileInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateReportProfileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateReportProfileInvoker) Invoke() (*model.UpdateReportProfileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateReportProfileResponse), nil
	}
}

type ChangeEastWestFirewallStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeEastWestFirewallStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeEastWestFirewallStatusInvoker) Invoke() (*model.ChangeEastWestFirewallStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeEastWestFirewallStatusResponse), nil
	}
}

type ShowEwAssociatedErInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEwAssociatedErInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEwAssociatedErInvoker) Invoke() (*model.ShowEwAssociatedErResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEwAssociatedErResponse), nil
	}
}

type ShowEwAssociatedVpcInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEwAssociatedVpcInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEwAssociatedVpcInvoker) Invoke() (*model.ShowEwAssociatedVpcResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEwAssociatedVpcResponse), nil
	}
}
