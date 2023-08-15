package v2

import (
	http_client "github.com/dysodeng/huaweicloud-sdk-go-v3/core"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/bcs/v2/model"
)

type BcsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewBcsClient(hcClient *http_client.HcHttpClient) *BcsClient {
	return &BcsClient{HcClient: hcClient}
}

func BcsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// BatchAddPeersToChannel peer节点加入通道
//
// peer节点加入通道,目前仅支持往一个通道中加入peer
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) BatchAddPeersToChannel(request *model.BatchAddPeersToChannelRequest) (*model.BatchAddPeersToChannelResponse, error) {
	requestDef := GenReqDefForBatchAddPeersToChannel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddPeersToChannelResponse), nil
	}
}

// BatchAddPeersToChannelInvoker peer节点加入通道
func (c *BcsClient) BatchAddPeersToChannelInvoker(request *model.BatchAddPeersToChannelRequest) *BatchAddPeersToChannelInvoker {
	requestDef := GenReqDefForBatchAddPeersToChannel()
	return &BatchAddPeersToChannelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateChannels 创建通道
//
// 创建通道
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) BatchCreateChannels(request *model.BatchCreateChannelsRequest) (*model.BatchCreateChannelsResponse, error) {
	requestDef := GenReqDefForBatchCreateChannels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateChannelsResponse), nil
	}
}

// BatchCreateChannelsInvoker 创建通道
func (c *BcsClient) BatchCreateChannelsInvoker(request *model.BatchCreateChannelsRequest) *BatchCreateChannelsInvoker {
	requestDef := GenReqDefForBatchCreateChannels()
	return &BatchCreateChannelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchInviteMembersToChannel 邀请联盟成员
//
// 批量邀请联盟成员加入通道，此操作会向被邀请方发出邀请通知
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) BatchInviteMembersToChannel(request *model.BatchInviteMembersToChannelRequest) (*model.BatchInviteMembersToChannelResponse, error) {
	requestDef := GenReqDefForBatchInviteMembersToChannel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchInviteMembersToChannelResponse), nil
	}
}

// BatchInviteMembersToChannelInvoker 邀请联盟成员
func (c *BcsClient) BatchInviteMembersToChannelInvoker(request *model.BatchInviteMembersToChannelRequest) *BatchInviteMembersToChannelInvoker {
	requestDef := GenReqDefForBatchInviteMembersToChannel()
	return &BatchInviteMembersToChannelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchRemoveOrgsFromChannel BCS组织退出某通道
//
// 该接口用于BCS组织退出某通道。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) BatchRemoveOrgsFromChannel(request *model.BatchRemoveOrgsFromChannelRequest) (*model.BatchRemoveOrgsFromChannelResponse, error) {
	requestDef := GenReqDefForBatchRemoveOrgsFromChannel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRemoveOrgsFromChannelResponse), nil
	}
}

// BatchRemoveOrgsFromChannelInvoker BCS组织退出某通道
func (c *BcsClient) BatchRemoveOrgsFromChannelInvoker(request *model.BatchRemoveOrgsFromChannelRequest) *BatchRemoveOrgsFromChannelInvoker {
	requestDef := GenReqDefForBatchRemoveOrgsFromChannel()
	return &BatchRemoveOrgsFromChannelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchRemovePeersFromChannel BCS某个组织中的节点退出某通道
//
// 该接口用于BCS某个组织中的节点退出某通道。当节点为通道中最后一个节点时，需要使用组织退通道的接口来将通道中的最后一个节点退出。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) BatchRemovePeersFromChannel(request *model.BatchRemovePeersFromChannelRequest) (*model.BatchRemovePeersFromChannelResponse, error) {
	requestDef := GenReqDefForBatchRemovePeersFromChannel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRemovePeersFromChannelResponse), nil
	}
}

// BatchRemovePeersFromChannelInvoker BCS某个组织中的节点退出某通道
func (c *BcsClient) BatchRemovePeersFromChannelInvoker(request *model.BatchRemovePeersFromChannelRequest) *BatchRemovePeersFromChannelInvoker {
	requestDef := GenReqDefForBatchRemovePeersFromChannel()
	return &BatchRemovePeersFromChannelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateBlockchainCertByUserName 生成用户证书
//
// 通过用户名生成指定服务实例组织用户证书
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) CreateBlockchainCertByUserName(request *model.CreateBlockchainCertByUserNameRequest) (*model.CreateBlockchainCertByUserNameResponse, error) {
	requestDef := GenReqDefForCreateBlockchainCertByUserName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBlockchainCertByUserNameResponse), nil
	}
}

// CreateBlockchainCertByUserNameInvoker 生成用户证书
func (c *BcsClient) CreateBlockchainCertByUserNameInvoker(request *model.CreateBlockchainCertByUserNameRequest) *CreateBlockchainCertByUserNameInvoker {
	requestDef := GenReqDefForCreateBlockchainCertByUserName()
	return &CreateBlockchainCertByUserNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateNewBlockchain 创建服务实例
//
// 创建BCS服务实例,只支持按需创建
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) CreateNewBlockchain(request *model.CreateNewBlockchainRequest) (*model.CreateNewBlockchainResponse, error) {
	requestDef := GenReqDefForCreateNewBlockchain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNewBlockchainResponse), nil
	}
}

// CreateNewBlockchainInvoker 创建服务实例
func (c *BcsClient) CreateNewBlockchainInvoker(request *model.CreateNewBlockchainRequest) *CreateNewBlockchainInvoker {
	requestDef := GenReqDefForCreateNewBlockchain()
	return &CreateNewBlockchainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBlockchain 删除服务实例
//
// 删除bcs实例。包周期实例不支持
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) DeleteBlockchain(request *model.DeleteBlockchainRequest) (*model.DeleteBlockchainResponse, error) {
	requestDef := GenReqDefForDeleteBlockchain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBlockchainResponse), nil
	}
}

// DeleteBlockchainInvoker 删除服务实例
func (c *BcsClient) DeleteBlockchainInvoker(request *model.DeleteBlockchainRequest) *DeleteBlockchainInvoker {
	requestDef := GenReqDefForDeleteBlockchain()
	return &DeleteBlockchainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteChannel BCS删除某个通道
//
// 该接口用于BCS删除某个通道。仅支持删除空通道
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) DeleteChannel(request *model.DeleteChannelRequest) (*model.DeleteChannelResponse, error) {
	requestDef := GenReqDefForDeleteChannel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteChannelResponse), nil
	}
}

// DeleteChannelInvoker BCS删除某个通道
func (c *BcsClient) DeleteChannelInvoker(request *model.DeleteChannelRequest) *DeleteChannelInvoker {
	requestDef := GenReqDefForDeleteChannel()
	return &DeleteChannelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMemberInvite 删除邀请成员信息
//
// 可通过此接口批量取消邀请或删除对已退出或拒绝加入或解散的成员邀请信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) DeleteMemberInvite(request *model.DeleteMemberInviteRequest) (*model.DeleteMemberInviteResponse, error) {
	requestDef := GenReqDefForDeleteMemberInvite()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMemberInviteResponse), nil
	}
}

// DeleteMemberInviteInvoker 删除邀请成员信息
func (c *BcsClient) DeleteMemberInviteInvoker(request *model.DeleteMemberInviteRequest) *DeleteMemberInviteInvoker {
	requestDef := GenReqDefForDeleteMemberInvite()
	return &DeleteMemberInviteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadBlockchainCert 下载证书
//
// 下载指定服务实例相关证书
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) DownloadBlockchainCert(request *model.DownloadBlockchainCertRequest) (*model.DownloadBlockchainCertResponse, error) {
	requestDef := GenReqDefForDownloadBlockchainCert()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadBlockchainCertResponse), nil
	}
}

// DownloadBlockchainCertInvoker 下载证书
func (c *BcsClient) DownloadBlockchainCertInvoker(request *model.DownloadBlockchainCertRequest) *DownloadBlockchainCertInvoker {
	requestDef := GenReqDefForDownloadBlockchainCert()
	return &DownloadBlockchainCertInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadBlockchainSdkConfig 下载SDK配置
//
// 下载指定服务实例SDK配置文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) DownloadBlockchainSdkConfig(request *model.DownloadBlockchainSdkConfigRequest) (*model.DownloadBlockchainSdkConfigResponse, error) {
	requestDef := GenReqDefForDownloadBlockchainSdkConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadBlockchainSdkConfigResponse), nil
	}
}

// DownloadBlockchainSdkConfigInvoker 下载SDK配置
func (c *BcsClient) DownloadBlockchainSdkConfigInvoker(request *model.DownloadBlockchainSdkConfigRequest) *DownloadBlockchainSdkConfigInvoker {
	requestDef := GenReqDefForDownloadBlockchainSdkConfig()
	return &DownloadBlockchainSdkConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// FreezeCert 冻结用户证书
//
// 冻结指定服务实例组织用户证书，冻结后需等待半分钟到一分钟左右生效
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) FreezeCert(request *model.FreezeCertRequest) (*model.FreezeCertResponse, error) {
	requestDef := GenReqDefForFreezeCert()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.FreezeCertResponse), nil
	}
}

// FreezeCertInvoker 冻结用户证书
func (c *BcsClient) FreezeCertInvoker(request *model.FreezeCertRequest) *FreezeCertInvoker {
	requestDef := GenReqDefForFreezeCert()
	return &FreezeCertInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// HandleNotification 处理联盟邀请
//
// 处理联盟邀请
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) HandleNotification(request *model.HandleNotificationRequest) (*model.HandleNotificationResponse, error) {
	requestDef := GenReqDefForHandleNotification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.HandleNotificationResponse), nil
	}
}

// HandleNotificationInvoker 处理联盟邀请
func (c *BcsClient) HandleNotificationInvoker(request *model.HandleNotificationRequest) *HandleNotificationInvoker {
	requestDef := GenReqDefForHandleNotification()
	return &HandleNotificationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// HandleUnionMemberQuitList 被邀请方退出指定联盟
//
// 被邀请方退出联盟
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) HandleUnionMemberQuitList(request *model.HandleUnionMemberQuitListRequest) (*model.HandleUnionMemberQuitListResponse, error) {
	requestDef := GenReqDefForHandleUnionMemberQuitList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.HandleUnionMemberQuitListResponse), nil
	}
}

// HandleUnionMemberQuitListInvoker 被邀请方退出指定联盟
func (c *BcsClient) HandleUnionMemberQuitListInvoker(request *model.HandleUnionMemberQuitListRequest) *HandleUnionMemberQuitListInvoker {
	requestDef := GenReqDefForHandleUnionMemberQuitList()
	return &HandleUnionMemberQuitListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBcsEvents 查询服务实例告警信息
//
// 该接口用于查询BCS服务的事件、告警数据，可以指定查询时的过滤条件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) ListBcsEvents(request *model.ListBcsEventsRequest) (*model.ListBcsEventsResponse, error) {
	requestDef := GenReqDefForListBcsEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBcsEventsResponse), nil
	}
}

// ListBcsEventsInvoker 查询服务实例告警信息
func (c *BcsClient) ListBcsEventsInvoker(request *model.ListBcsEventsRequest) *ListBcsEventsInvoker {
	requestDef := GenReqDefForListBcsEvents()
	return &ListBcsEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBcsEventsStatistic 查询服务实例告警统计接口
//
// 该接口用于查询BCS服务的分段事件、告警统计数据，可以指定查询时的过滤条件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) ListBcsEventsStatistic(request *model.ListBcsEventsStatisticRequest) (*model.ListBcsEventsStatisticResponse, error) {
	requestDef := GenReqDefForListBcsEventsStatistic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBcsEventsStatisticResponse), nil
	}
}

// ListBcsEventsStatisticInvoker 查询服务实例告警统计接口
func (c *BcsClient) ListBcsEventsStatisticInvoker(request *model.ListBcsEventsStatisticRequest) *ListBcsEventsStatisticInvoker {
	requestDef := GenReqDefForListBcsEventsStatistic()
	return &ListBcsEventsStatisticInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBcsMetric 查询服务实例监控数据
//
// 该接口用于查询BCS服务的监控数据，可以指定相应的指标名称。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) ListBcsMetric(request *model.ListBcsMetricRequest) (*model.ListBcsMetricResponse, error) {
	requestDef := GenReqDefForListBcsMetric()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBcsMetricResponse), nil
	}
}

// ListBcsMetricInvoker 查询服务实例监控数据
func (c *BcsClient) ListBcsMetricInvoker(request *model.ListBcsMetricRequest) *ListBcsMetricInvoker {
	requestDef := GenReqDefForListBcsMetric()
	return &ListBcsMetricInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBlockchainChannels 查询通道信息
//
// 查询指定服务实例通道信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) ListBlockchainChannels(request *model.ListBlockchainChannelsRequest) (*model.ListBlockchainChannelsResponse, error) {
	requestDef := GenReqDefForListBlockchainChannels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBlockchainChannelsResponse), nil
	}
}

// ListBlockchainChannelsInvoker 查询通道信息
func (c *BcsClient) ListBlockchainChannelsInvoker(request *model.ListBlockchainChannelsRequest) *ListBlockchainChannelsInvoker {
	requestDef := GenReqDefForListBlockchainChannels()
	return &ListBlockchainChannelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBlockchains 查询服务实例列表
//
// 查询当前项目下所有服务实例的简要信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) ListBlockchains(request *model.ListBlockchainsRequest) (*model.ListBlockchainsResponse, error) {
	requestDef := GenReqDefForListBlockchains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBlockchainsResponse), nil
	}
}

// ListBlockchainsInvoker 查询服务实例列表
func (c *BcsClient) ListBlockchainsInvoker(request *model.ListBlockchainsRequest) *ListBlockchainsInvoker {
	requestDef := GenReqDefForListBlockchains()
	return &ListBlockchainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEntityMetric 查询BCS组织监控数据列表
//
// 该接口用于查询BCS组织的监控数据列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) ListEntityMetric(request *model.ListEntityMetricRequest) (*model.ListEntityMetricResponse, error) {
	requestDef := GenReqDefForListEntityMetric()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEntityMetricResponse), nil
	}
}

// ListEntityMetricInvoker 查询BCS组织监控数据列表
func (c *BcsClient) ListEntityMetricInvoker(request *model.ListEntityMetricRequest) *ListEntityMetricInvoker {
	requestDef := GenReqDefForListEntityMetric()
	return &ListEntityMetricInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceMetric 查询BCS组织实例监控数据详情
//
// 该接口用于BCS组织实例监控数据详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) ListInstanceMetric(request *model.ListInstanceMetricRequest) (*model.ListInstanceMetricResponse, error) {
	requestDef := GenReqDefForListInstanceMetric()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceMetricResponse), nil
	}
}

// ListInstanceMetricInvoker 查询BCS组织实例监控数据详情
func (c *BcsClient) ListInstanceMetricInvoker(request *model.ListInstanceMetricRequest) *ListInstanceMetricInvoker {
	requestDef := GenReqDefForListInstanceMetric()
	return &ListInstanceMetricInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMembers 获取联盟成员列表
//
// 获取联盟成员列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) ListMembers(request *model.ListMembersRequest) (*model.ListMembersResponse, error) {
	requestDef := GenReqDefForListMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMembersResponse), nil
	}
}

// ListMembersInvoker 获取联盟成员列表
func (c *BcsClient) ListMembersInvoker(request *model.ListMembersRequest) *ListMembersInvoker {
	requestDef := GenReqDefForListMembers()
	return &ListMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNotifications 获取全部通知
//
// 获取全部通知
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) ListNotifications(request *model.ListNotificationsRequest) (*model.ListNotificationsResponse, error) {
	requestDef := GenReqDefForListNotifications()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNotificationsResponse), nil
	}
}

// ListNotificationsInvoker 获取全部通知
func (c *BcsClient) ListNotificationsInvoker(request *model.ListNotificationsRequest) *ListNotificationsInvoker {
	requestDef := GenReqDefForListNotifications()
	return &ListNotificationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOpRecord 查询异步操作结果
//
// 查询异步操作结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) ListOpRecord(request *model.ListOpRecordRequest) (*model.ListOpRecordResponse, error) {
	requestDef := GenReqDefForListOpRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOpRecordResponse), nil
	}
}

// ListOpRecordInvoker 查询异步操作结果
func (c *BcsClient) ListOpRecordInvoker(request *model.ListOpRecordRequest) *ListOpRecordInvoker {
	requestDef := GenReqDefForListOpRecord()
	return &ListOpRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQuotas 查询配额
//
// 查询当前项目下BCS服务所有资源的配额信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) ListQuotas(request *model.ListQuotasRequest) (*model.ListQuotasResponse, error) {
	requestDef := GenReqDefForListQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotasResponse), nil
	}
}

// ListQuotasInvoker 查询配额
func (c *BcsClient) ListQuotasInvoker(request *model.ListQuotasRequest) *ListQuotasInvoker {
	requestDef := GenReqDefForListQuotas()
	return &ListQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBlockchainDetail 查询实例信息
//
// 查询指定服务实例详细信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) ShowBlockchainDetail(request *model.ShowBlockchainDetailRequest) (*model.ShowBlockchainDetailResponse, error) {
	requestDef := GenReqDefForShowBlockchainDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBlockchainDetailResponse), nil
	}
}

// ShowBlockchainDetailInvoker 查询实例信息
func (c *BcsClient) ShowBlockchainDetailInvoker(request *model.ShowBlockchainDetailRequest) *ShowBlockchainDetailInvoker {
	requestDef := GenReqDefForShowBlockchainDetail()
	return &ShowBlockchainDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBlockchainFlavors 查询规格
//
// 查询服务联盟链规格信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) ShowBlockchainFlavors(request *model.ShowBlockchainFlavorsRequest) (*model.ShowBlockchainFlavorsResponse, error) {
	requestDef := GenReqDefForShowBlockchainFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBlockchainFlavorsResponse), nil
	}
}

// ShowBlockchainFlavorsInvoker 查询规格
func (c *BcsClient) ShowBlockchainFlavorsInvoker(request *model.ShowBlockchainFlavorsRequest) *ShowBlockchainFlavorsInvoker {
	requestDef := GenReqDefForShowBlockchainFlavors()
	return &ShowBlockchainFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBlockchainNodes 查询节点信息
//
// 查询指定服务实例节点信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) ShowBlockchainNodes(request *model.ShowBlockchainNodesRequest) (*model.ShowBlockchainNodesResponse, error) {
	requestDef := GenReqDefForShowBlockchainNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBlockchainNodesResponse), nil
	}
}

// ShowBlockchainNodesInvoker 查询节点信息
func (c *BcsClient) ShowBlockchainNodesInvoker(request *model.ShowBlockchainNodesRequest) *ShowBlockchainNodesInvoker {
	requestDef := GenReqDefForShowBlockchainNodes()
	return &ShowBlockchainNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBlockchainStatus 查询创建状态
//
// 查询指定服务实例创建状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) ShowBlockchainStatus(request *model.ShowBlockchainStatusRequest) (*model.ShowBlockchainStatusResponse, error) {
	requestDef := GenReqDefForShowBlockchainStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBlockchainStatusResponse), nil
	}
}

// ShowBlockchainStatusInvoker 查询创建状态
func (c *BcsClient) ShowBlockchainStatusInvoker(request *model.ShowBlockchainStatusRequest) *ShowBlockchainStatusInvoker {
	requestDef := GenReqDefForShowBlockchainStatus()
	return &ShowBlockchainStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UnfreezeCert 解冻用户证书
//
// 解冻指定服务实例组织用户证书，解冻后需等待半分钟到一分钟左右生效
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) UnfreezeCert(request *model.UnfreezeCertRequest) (*model.UnfreezeCertResponse, error) {
	requestDef := GenReqDefForUnfreezeCert()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnfreezeCertResponse), nil
	}
}

// UnfreezeCertInvoker 解冻用户证书
func (c *BcsClient) UnfreezeCertInvoker(request *model.UnfreezeCertRequest) *UnfreezeCertInvoker {
	requestDef := GenReqDefForUnfreezeCert()
	return &UnfreezeCertInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstance 修改服务实例
//
// 修改实例的节点、组织，目前仅支持添加、删除节点（IEF模式不支持添加、删除节点），添加、删除组织，共4种类型，每次操作只可以操作一种类型。此接口不支持包周期模式; 注意注册IEF节点时，IEF节点名称长度应该为4-24位的字符
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BcsClient) UpdateInstance(request *model.UpdateInstanceRequest) (*model.UpdateInstanceResponse, error) {
	requestDef := GenReqDefForUpdateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceResponse), nil
	}
}

// UpdateInstanceInvoker 修改服务实例
func (c *BcsClient) UpdateInstanceInvoker(request *model.UpdateInstanceRequest) *UpdateInstanceInvoker {
	requestDef := GenReqDefForUpdateInstance()
	return &UpdateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
