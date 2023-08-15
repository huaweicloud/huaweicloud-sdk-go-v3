package v1

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/cfw/v1/model"
)

type AddAddressItemsUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddAddressItemsUsingPostInvoker) Invoke() (*model.AddAddressItemsUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddAddressItemsUsingPostResponse), nil
	}
}

type AddAddressSetInfoUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddAddressSetInfoUsingPostInvoker) Invoke() (*model.AddAddressSetInfoUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddAddressSetInfoUsingPostResponse), nil
	}
}

type AddBlackWhiteListUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddBlackWhiteListUsingPostInvoker) Invoke() (*model.AddBlackWhiteListUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddBlackWhiteListUsingPostResponse), nil
	}
}

type AddServiceItemsUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddServiceItemsUsingPostInvoker) Invoke() (*model.AddServiceItemsUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddServiceItemsUsingPostResponse), nil
	}
}

type AddServiceSetUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddServiceSetUsingPostInvoker) Invoke() (*model.AddServiceSetUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddServiceSetUsingPostResponse), nil
	}
}

type ChangeEwProtectStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeEwProtectStatusInvoker) Invoke() (*model.ChangeEwProtectStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeEwProtectStatusResponse), nil
	}
}

type ChangeIpsProtectModeUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeIpsProtectModeUsingPostInvoker) Invoke() (*model.ChangeIpsProtectModeUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeIpsProtectModeUsingPostResponse), nil
	}
}

type DeleteAclRuleCountInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAclRuleCountInvoker) Invoke() (*model.DeleteAclRuleCountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAclRuleCountResponse), nil
	}
}

type DeleteAddressItemUsingDeleteInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAddressItemUsingDeleteInvoker) Invoke() (*model.DeleteAddressItemUsingDeleteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAddressItemUsingDeleteResponse), nil
	}
}

type DeleteAddressSetInfoUsingDeleteInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAddressSetInfoUsingDeleteInvoker) Invoke() (*model.DeleteAddressSetInfoUsingDeleteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAddressSetInfoUsingDeleteResponse), nil
	}
}

type DeleteBlackWhiteListUsingDeleteInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBlackWhiteListUsingDeleteInvoker) Invoke() (*model.DeleteBlackWhiteListUsingDeleteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBlackWhiteListUsingDeleteResponse), nil
	}
}

type DeleteServiceItemUsingDeleteInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteServiceItemUsingDeleteInvoker) Invoke() (*model.DeleteServiceItemUsingDeleteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteServiceItemUsingDeleteResponse), nil
	}
}

type DeleteServiceSetUsingDeleteInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteServiceSetUsingDeleteInvoker) Invoke() (*model.DeleteServiceSetUsingDeleteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteServiceSetUsingDeleteResponse), nil
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

type ListAddressItemsUsingGetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAddressItemsUsingGetInvoker) Invoke() (*model.ListAddressItemsUsingGetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAddressItemsUsingGetResponse), nil
	}
}

type ListAddressSetDetailUsingGetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAddressSetDetailUsingGetInvoker) Invoke() (*model.ListAddressSetDetailUsingGetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAddressSetDetailUsingGetResponse), nil
	}
}

type ListAddressSetListUsingGetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAddressSetListUsingGetInvoker) Invoke() (*model.ListAddressSetListUsingGetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAddressSetListUsingGetResponse), nil
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

type ListBlackWhiteListsUsingGetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBlackWhiteListsUsingGetInvoker) Invoke() (*model.ListBlackWhiteListsUsingGetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBlackWhiteListsUsingGetResponse), nil
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

type ListFirewallUsingGetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFirewallUsingGetInvoker) Invoke() (*model.ListFirewallUsingGetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFirewallUsingGetResponse), nil
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

type ListIpsProtectModeUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIpsProtectModeUsingPostInvoker) Invoke() (*model.ListIpsProtectModeUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIpsProtectModeUsingPostResponse), nil
	}
}

type ListParseDomainDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListParseDomainDetailsInvoker) Invoke() (*model.ListParseDomainDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListParseDomainDetailsResponse), nil
	}
}

type ListRuleHitCountInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRuleHitCountInvoker) Invoke() (*model.ListRuleHitCountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRuleHitCountResponse), nil
	}
}

type ListServiceItemsDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServiceItemsDetailsInvoker) Invoke() (*model.ListServiceItemsDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServiceItemsDetailsResponse), nil
	}
}

type ListServiceSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServiceSetInvoker) Invoke() (*model.ListServiceSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServiceSetResponse), nil
	}
}

type ListServiceSetDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServiceSetDetailsInvoker) Invoke() (*model.ListServiceSetDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServiceSetDetailsResponse), nil
	}
}

type UpdateAddressSetInfoUsingPutInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAddressSetInfoUsingPutInvoker) Invoke() (*model.UpdateAddressSetInfoUsingPutResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAddressSetInfoUsingPutResponse), nil
	}
}

type UpdateBlackWhiteListUsingPutInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateBlackWhiteListUsingPutInvoker) Invoke() (*model.UpdateBlackWhiteListUsingPutResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateBlackWhiteListUsingPutResponse), nil
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

type UpdateServiceSetUsingPutInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateServiceSetUsingPutInvoker) Invoke() (*model.UpdateServiceSetUsingPutResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateServiceSetUsingPutResponse), nil
	}
}

type AddRuleAclUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddRuleAclUsingPostInvoker) Invoke() (*model.AddRuleAclUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddRuleAclUsingPostResponse), nil
	}
}

type DeleteRuleAclUsingDeleteInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRuleAclUsingDeleteInvoker) Invoke() (*model.DeleteRuleAclUsingDeleteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRuleAclUsingDeleteResponse), nil
	}
}

type ListRuleAclUsingPutInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRuleAclUsingPutInvoker) Invoke() (*model.ListRuleAclUsingPutResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRuleAclUsingPutResponse), nil
	}
}

type ListRuleAclsUsingGetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRuleAclsUsingGetInvoker) Invoke() (*model.ListRuleAclsUsingGetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRuleAclsUsingGetResponse), nil
	}
}

type UpdateRuleAclUsingPutInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRuleAclUsingPutInvoker) Invoke() (*model.UpdateRuleAclUsingPutResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRuleAclUsingPutResponse), nil
	}
}

type ChangeProtectEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeProtectEipInvoker) Invoke() (*model.ChangeProtectEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeProtectEipResponse), nil
	}
}

type CountEipsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountEipsInvoker) Invoke() (*model.CountEipsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountEipsResponse), nil
	}
}

type ListEipResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEipResourcesInvoker) Invoke() (*model.ListEipResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEipResourcesResponse), nil
	}
}

type ChangeIpsSwitchUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeIpsSwitchUsingPostInvoker) Invoke() (*model.ChangeIpsSwitchUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeIpsSwitchUsingPostResponse), nil
	}
}

type ListIpsSwitchStatusUsingGetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIpsSwitchStatusUsingGetInvoker) Invoke() (*model.ListIpsSwitchStatusUsingGetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIpsSwitchStatusUsingGetResponse), nil
	}
}

type ListVpcProtectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVpcProtectsInvoker) Invoke() (*model.ListVpcProtectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVpcProtectsResponse), nil
	}
}
