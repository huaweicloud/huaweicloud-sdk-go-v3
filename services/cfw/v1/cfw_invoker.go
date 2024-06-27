package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cfw/v1/model"
)

type AddAddressItemInvoker struct {
	*invoker.BaseInvoker
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

func (i *BatchDeleteAddressItemsInvoker) Invoke() (*model.BatchDeleteAddressItemsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteAddressItemsResponse), nil
	}
}

type BatchDeleteServiceItemsInvoker struct {
	*invoker.BaseInvoker
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

func (i *CancelCaptureTaskInvoker) Invoke() (*model.CancelCaptureTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelCaptureTaskResponse), nil
	}
}

type ChangeEastWestFirewallStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeEastWestFirewallStatusInvoker) Invoke() (*model.ChangeEastWestFirewallStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeEastWestFirewallStatusResponse), nil
	}
}

type CreateCaptureTaskInvoker struct {
	*invoker.BaseInvoker
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

func (i *DeleteFirewallInvoker) Invoke() (*model.DeleteFirewallResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFirewallResponse), nil
	}
}

type DeleteServiceItemInvoker struct {
	*invoker.BaseInvoker
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

func (i *DeleteTagInvoker) Invoke() (*model.DeleteTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTagResponse), nil
	}
}

type ListAccessControlLogsInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListDomainParseDetailInvoker) Invoke() (*model.ListDomainParseDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDomainParseDetailResponse), nil
	}
}

type ListDomainSetsInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListFlowLogsInvoker) Invoke() (*model.ListFlowLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlowLogsResponse), nil
	}
}

type ListJobInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListLogConfigInvoker) Invoke() (*model.ListLogConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLogConfigResponse), nil
	}
}

type ListProtectedVpcsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProtectedVpcsInvoker) Invoke() (*model.ListProtectedVpcsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProtectedVpcsResponse), nil
	}
}

type ListServiceItemsInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListServiceSetsInvoker) Invoke() (*model.ListServiceSetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServiceSetsResponse), nil
	}
}

type UpdateAddressSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAddressSetInvoker) Invoke() (*model.UpdateAddressSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAddressSetResponse), nil
	}
}

type UpdateBlackWhiteListInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdateLogConfigInvoker) Invoke() (*model.UpdateLogConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLogConfigResponse), nil
	}
}

type UpdateServiceSetInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListAclRulesInvoker) Invoke() (*model.ListAclRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAclRulesResponse), nil
	}
}

type ListRuleAclTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRuleAclTagsInvoker) Invoke() (*model.ListRuleAclTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRuleAclTagsResponse), nil
	}
}

type UpdateAclRuleInvoker struct {
	*invoker.BaseInvoker
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

func (i *ChangeEipStatusInvoker) Invoke() (*model.ChangeEipStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeEipStatusResponse), nil
	}
}

type ListEipCountInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListEipsInvoker) Invoke() (*model.ListEipsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEipsResponse), nil
	}
}

type ChangeIpsProtectModeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeIpsProtectModeInvoker) Invoke() (*model.ChangeIpsProtectModeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeIpsProtectModeResponse), nil
	}
}

type ChangeIpsSwitchStatusInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListIpsProtectModeInvoker) Invoke() (*model.ListIpsProtectModeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIpsProtectModeResponse), nil
	}
}

type ListIpsSwitchStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIpsSwitchStatusInvoker) Invoke() (*model.ListIpsSwitchStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIpsSwitchStatusResponse), nil
	}
}
