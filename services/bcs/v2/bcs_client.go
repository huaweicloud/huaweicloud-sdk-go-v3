package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bcs/v2/model"
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

//peer节点加入通道,目前仅支持往一个通道中加入peer
func (c *BcsClient) BatchAddPeersToChannel(request *model.BatchAddPeersToChannelRequest) (*model.BatchAddPeersToChannelResponse, error) {
	requestDef := GenReqDefForBatchAddPeersToChannel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddPeersToChannelResponse), nil
	}
}

//创建通道
func (c *BcsClient) BatchCreateChannels(request *model.BatchCreateChannelsRequest) (*model.BatchCreateChannelsResponse, error) {
	requestDef := GenReqDefForBatchCreateChannels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateChannelsResponse), nil
	}
}

//批量邀请联盟成员加入通道，此操作会向被邀请方发出邀请通知
func (c *BcsClient) BatchInviteMembersToChannel(request *model.BatchInviteMembersToChannelRequest) (*model.BatchInviteMembersToChannelResponse, error) {
	requestDef := GenReqDefForBatchInviteMembersToChannel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchInviteMembersToChannelResponse), nil
	}
}

//该接口用于BCS组织退出某通道。
func (c *BcsClient) BatchRemoveOrgsFromChannel(request *model.BatchRemoveOrgsFromChannelRequest) (*model.BatchRemoveOrgsFromChannelResponse, error) {
	requestDef := GenReqDefForBatchRemoveOrgsFromChannel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRemoveOrgsFromChannelResponse), nil
	}
}

//通过用户名生成指定服务实例组织用户证书
func (c *BcsClient) CreateBlockchainCertByUserName(request *model.CreateBlockchainCertByUserNameRequest) (*model.CreateBlockchainCertByUserNameResponse, error) {
	requestDef := GenReqDefForCreateBlockchainCertByUserName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBlockchainCertByUserNameResponse), nil
	}
}

//创建BCS服务实例,只支持按需创建
func (c *BcsClient) CreateNewBlockchain(request *model.CreateNewBlockchainRequest) (*model.CreateNewBlockchainResponse, error) {
	requestDef := GenReqDefForCreateNewBlockchain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNewBlockchainResponse), nil
	}
}

//删除bcs实例。包周期实例不支持
func (c *BcsClient) DeleteBlockchain(request *model.DeleteBlockchainRequest) (*model.DeleteBlockchainResponse, error) {
	requestDef := GenReqDefForDeleteBlockchain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBlockchainResponse), nil
	}
}

//可通过此接口批量取消邀请或删除对已退出或拒绝加入或解散的成员邀请信息
func (c *BcsClient) DeleteMemberInvite(request *model.DeleteMemberInviteRequest) (*model.DeleteMemberInviteResponse, error) {
	requestDef := GenReqDefForDeleteMemberInvite()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMemberInviteResponse), nil
	}
}

//下载指定服务实例相关证书
func (c *BcsClient) DownloadBlockchainCert(request *model.DownloadBlockchainCertRequest) (*model.DownloadBlockchainCertResponse, error) {
	requestDef := GenReqDefForDownloadBlockchainCert()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadBlockchainCertResponse), nil
	}
}

//下载指定服务实例SDK配置文件
func (c *BcsClient) DownloadBlockchainSdkConfig(request *model.DownloadBlockchainSdkConfigRequest) (*model.DownloadBlockchainSdkConfigResponse, error) {
	requestDef := GenReqDefForDownloadBlockchainSdkConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadBlockchainSdkConfigResponse), nil
	}
}

//冻结指定服务实例组织用户证书，冻结后需等待半分钟到一分钟左右生效
func (c *BcsClient) FreezeCert(request *model.FreezeCertRequest) (*model.FreezeCertResponse, error) {
	requestDef := GenReqDefForFreezeCert()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.FreezeCertResponse), nil
	}
}

//处理联盟邀请
func (c *BcsClient) HandleNotification(request *model.HandleNotificationRequest) (*model.HandleNotificationResponse, error) {
	requestDef := GenReqDefForHandleNotification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.HandleNotificationResponse), nil
	}
}

//该接口用于查询BCS服务的监控数据，可以指定相应的指标名称。[目前不支持IEF节点](tag:hasief)
func (c *BcsClient) ListBcsMetric(request *model.ListBcsMetricRequest) (*model.ListBcsMetricResponse, error) {
	requestDef := GenReqDefForListBcsMetric()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBcsMetricResponse), nil
	}
}

//查询指定服务实例通道信息
func (c *BcsClient) ListBlockchainChannels(request *model.ListBlockchainChannelsRequest) (*model.ListBlockchainChannelsResponse, error) {
	requestDef := GenReqDefForListBlockchainChannels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBlockchainChannelsResponse), nil
	}
}

//查询当前项目下所有服务实例的简要信息
func (c *BcsClient) ListBlockchains(request *model.ListBlockchainsRequest) (*model.ListBlockchainsResponse, error) {
	requestDef := GenReqDefForListBlockchains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBlockchainsResponse), nil
	}
}

//该接口用于查询BCS组织的监控数据列表。
func (c *BcsClient) ListEntityMetric(request *model.ListEntityMetricRequest) (*model.ListEntityMetricResponse, error) {
	requestDef := GenReqDefForListEntityMetric()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEntityMetricResponse), nil
	}
}

//该接口用于BCS组织实例监控数据详情。
func (c *BcsClient) ListInstanceMetric(request *model.ListInstanceMetricRequest) (*model.ListInstanceMetricResponse, error) {
	requestDef := GenReqDefForListInstanceMetric()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceMetricResponse), nil
	}
}

//获取联盟成员列表
func (c *BcsClient) ListMembers(request *model.ListMembersRequest) (*model.ListMembersResponse, error) {
	requestDef := GenReqDefForListMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMembersResponse), nil
	}
}

//获取全部通知
func (c *BcsClient) ListNotifications(request *model.ListNotificationsRequest) (*model.ListNotificationsResponse, error) {
	requestDef := GenReqDefForListNotifications()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNotificationsResponse), nil
	}
}

//查询异步操作结果
func (c *BcsClient) ListOpRecord(request *model.ListOpRecordRequest) (*model.ListOpRecordResponse, error) {
	requestDef := GenReqDefForListOpRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOpRecordResponse), nil
	}
}

//查询当前项目下BCS服务所有资源的配额信息
func (c *BcsClient) ListQuotas(request *model.ListQuotasRequest) (*model.ListQuotasResponse, error) {
	requestDef := GenReqDefForListQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotasResponse), nil
	}
}

//查询指定服务实例详细信息
func (c *BcsClient) ShowBlockchainDetail(request *model.ShowBlockchainDetailRequest) (*model.ShowBlockchainDetailResponse, error) {
	requestDef := GenReqDefForShowBlockchainDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBlockchainDetailResponse), nil
	}
}

//查询服务联盟链规格信息
func (c *BcsClient) ShowBlockchainFlavors(request *model.ShowBlockchainFlavorsRequest) (*model.ShowBlockchainFlavorsResponse, error) {
	requestDef := GenReqDefForShowBlockchainFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBlockchainFlavorsResponse), nil
	}
}

//查询指定服务实例节点信息
func (c *BcsClient) ShowBlockchainNodes(request *model.ShowBlockchainNodesRequest) (*model.ShowBlockchainNodesResponse, error) {
	requestDef := GenReqDefForShowBlockchainNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBlockchainNodesResponse), nil
	}
}

//查询指定服务实例创建状态
func (c *BcsClient) ShowBlockchainStatus(request *model.ShowBlockchainStatusRequest) (*model.ShowBlockchainStatusResponse, error) {
	requestDef := GenReqDefForShowBlockchainStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBlockchainStatusResponse), nil
	}
}

//解冻指定服务实例组织用户证书，解冻后需等待半分钟到一分钟左右生效
func (c *BcsClient) UnfreezeCert(request *model.UnfreezeCertRequest) (*model.UnfreezeCertResponse, error) {
	requestDef := GenReqDefForUnfreezeCert()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnfreezeCertResponse), nil
	}
}

//修改实例的节点、组织，目前仅支持添加、删除节点（IEF模式不支持添加、删除节点），添加、删除组织，共4种类型，每次操作只可以操作一种类型。此接口不支持包周期模式; 注意注册IEF节点时，IEF节点名称长度应该为4-24位的字符
func (c *BcsClient) UpdateInstance(request *model.UpdateInstanceRequest) (*model.UpdateInstanceResponse, error) {
	requestDef := GenReqDefForUpdateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceResponse), nil
	}
}
