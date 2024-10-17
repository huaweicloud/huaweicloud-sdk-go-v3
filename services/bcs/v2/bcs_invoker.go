package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bcs/v2/model"
)

type BatchAddPeersToChannelInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAddPeersToChannelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchAddPeersToChannelInvoker) Invoke() (*model.BatchAddPeersToChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAddPeersToChannelResponse), nil
	}
}

type BatchCreateChannelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateChannelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateChannelsInvoker) Invoke() (*model.BatchCreateChannelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateChannelsResponse), nil
	}
}

type BatchInviteMembersToChannelInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchInviteMembersToChannelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchInviteMembersToChannelInvoker) Invoke() (*model.BatchInviteMembersToChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchInviteMembersToChannelResponse), nil
	}
}

type BatchRemoveOrgsFromChannelInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchRemoveOrgsFromChannelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchRemoveOrgsFromChannelInvoker) Invoke() (*model.BatchRemoveOrgsFromChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchRemoveOrgsFromChannelResponse), nil
	}
}

type BatchRemovePeersFromChannelInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchRemovePeersFromChannelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchRemovePeersFromChannelInvoker) Invoke() (*model.BatchRemovePeersFromChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchRemovePeersFromChannelResponse), nil
	}
}

type CreateBlockchainCertByUserNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateBlockchainCertByUserNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateBlockchainCertByUserNameInvoker) Invoke() (*model.CreateBlockchainCertByUserNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateBlockchainCertByUserNameResponse), nil
	}
}

type CreateNewBlockchainInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNewBlockchainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateNewBlockchainInvoker) Invoke() (*model.CreateNewBlockchainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNewBlockchainResponse), nil
	}
}

type DeleteBlockchainInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBlockchainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteBlockchainInvoker) Invoke() (*model.DeleteBlockchainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBlockchainResponse), nil
	}
}

type DeleteChannelInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteChannelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteChannelInvoker) Invoke() (*model.DeleteChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteChannelResponse), nil
	}
}

type DeleteMemberInviteInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMemberInviteInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteMemberInviteInvoker) Invoke() (*model.DeleteMemberInviteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMemberInviteResponse), nil
	}
}

type DownloadBlockchainCertInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadBlockchainCertInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadBlockchainCertInvoker) Invoke() (*model.DownloadBlockchainCertResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadBlockchainCertResponse), nil
	}
}

type DownloadBlockchainSdkConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadBlockchainSdkConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadBlockchainSdkConfigInvoker) Invoke() (*model.DownloadBlockchainSdkConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadBlockchainSdkConfigResponse), nil
	}
}

type FreezeCertInvoker struct {
	*invoker.BaseInvoker
}

func (i *FreezeCertInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *FreezeCertInvoker) Invoke() (*model.FreezeCertResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.FreezeCertResponse), nil
	}
}

type HandleNotificationInvoker struct {
	*invoker.BaseInvoker
}

func (i *HandleNotificationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *HandleNotificationInvoker) Invoke() (*model.HandleNotificationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.HandleNotificationResponse), nil
	}
}

type HandleUnionMemberQuitListInvoker struct {
	*invoker.BaseInvoker
}

func (i *HandleUnionMemberQuitListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *HandleUnionMemberQuitListInvoker) Invoke() (*model.HandleUnionMemberQuitListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.HandleUnionMemberQuitListResponse), nil
	}
}

type ListBcsEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBcsEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBcsEventsInvoker) Invoke() (*model.ListBcsEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBcsEventsResponse), nil
	}
}

type ListBcsEventsStatisticInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBcsEventsStatisticInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBcsEventsStatisticInvoker) Invoke() (*model.ListBcsEventsStatisticResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBcsEventsStatisticResponse), nil
	}
}

type ListBcsMetricInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBcsMetricInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBcsMetricInvoker) Invoke() (*model.ListBcsMetricResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBcsMetricResponse), nil
	}
}

type ListBlockchainChannelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBlockchainChannelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBlockchainChannelsInvoker) Invoke() (*model.ListBlockchainChannelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBlockchainChannelsResponse), nil
	}
}

type ListBlockchainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBlockchainsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBlockchainsInvoker) Invoke() (*model.ListBlockchainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBlockchainsResponse), nil
	}
}

type ListEntityMetricInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEntityMetricInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEntityMetricInvoker) Invoke() (*model.ListEntityMetricResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEntityMetricResponse), nil
	}
}

type ListInstanceMetricInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceMetricInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceMetricInvoker) Invoke() (*model.ListInstanceMetricResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceMetricResponse), nil
	}
}

type ListMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMembersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMembersInvoker) Invoke() (*model.ListMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMembersResponse), nil
	}
}

type ListNotificationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNotificationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNotificationsInvoker) Invoke() (*model.ListNotificationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNotificationsResponse), nil
	}
}

type ListOpRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOpRecordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListOpRecordInvoker) Invoke() (*model.ListOpRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOpRecordResponse), nil
	}
}

type ListQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQuotasInvoker) Invoke() (*model.ListQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotasResponse), nil
	}
}

type ShowBlockchainDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBlockchainDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBlockchainDetailInvoker) Invoke() (*model.ShowBlockchainDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBlockchainDetailResponse), nil
	}
}

type ShowBlockchainFlavorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBlockchainFlavorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBlockchainFlavorsInvoker) Invoke() (*model.ShowBlockchainFlavorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBlockchainFlavorsResponse), nil
	}
}

type ShowBlockchainNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBlockchainNodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBlockchainNodesInvoker) Invoke() (*model.ShowBlockchainNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBlockchainNodesResponse), nil
	}
}

type ShowBlockchainStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBlockchainStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBlockchainStatusInvoker) Invoke() (*model.ShowBlockchainStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBlockchainStatusResponse), nil
	}
}

type UnfreezeCertInvoker struct {
	*invoker.BaseInvoker
}

func (i *UnfreezeCertInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UnfreezeCertInvoker) Invoke() (*model.UnfreezeCertResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnfreezeCertResponse), nil
	}
}

type UpdateInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstanceInvoker) Invoke() (*model.UpdateInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceResponse), nil
	}
}
