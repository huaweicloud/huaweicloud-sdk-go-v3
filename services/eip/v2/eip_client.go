package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eip/v2/model"
)

type EipClient struct {
	hcClient *http_client.HcHttpClient
}

func NewEipClient(hcClient *http_client.HcHttpClient) *EipClient {
	return &EipClient{hcClient: hcClient}
}

func EipClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//共享带宽插入弹性公网IP。
func (c *EipClient) AddPublicipsIntoSharedBandwidth(request *model.AddPublicipsIntoSharedBandwidthRequest) (*model.AddPublicipsIntoSharedBandwidthResponse, error) {
	requestDef := GenReqDefForAddPublicipsIntoSharedBandwidth(request)
	resp, responseDef := GenRespForAddPublicipsIntoSharedBandwidth()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//批量创建共享带宽。
func (c *EipClient) BatchCreateSharedBandwidths(request *model.BatchCreateSharedBandwidthsRequest) (*model.BatchCreateSharedBandwidthsResponse, error) {
	requestDef := GenReqDefForBatchCreateSharedBandwidths(request)
	resp, responseDef := GenRespForBatchCreateSharedBandwidths()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//创建共享带宽。
func (c *EipClient) CreateSharedBandwidth(request *model.CreateSharedBandwidthRequest) (*model.CreateSharedBandwidthResponse, error) {
	requestDef := GenReqDefForCreateSharedBandwidth(request)
	resp, responseDef := GenRespForCreateSharedBandwidth()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除共享带宽。
func (c *EipClient) DeleteSharedBandwidth(request *model.DeleteSharedBandwidthRequest) (*model.DeleteSharedBandwidthResponse, error) {
	requestDef := GenReqDefForDeleteSharedBandwidth(request)
	resp, responseDef := GenRespForDeleteSharedBandwidth()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询带宽列表。
func (c *EipClient) ListBandwidths(request *model.ListBandwidthsRequest) (*model.ListBandwidthsResponse, error) {
	requestDef := GenReqDefForListBandwidths(request)
	resp, responseDef := GenRespForListBandwidths()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询配额
func (c *EipClient) ListQuotas(request *model.ListQuotasRequest) (*model.ListQuotasResponse, error) {
	requestDef := GenReqDefForListQuotas(request)
	resp, responseDef := GenRespForListQuotas()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//共享带宽移除弹性公网IP。
func (c *EipClient) RemovePublicipsFromSharedBandwidth(request *model.RemovePublicipsFromSharedBandwidthRequest) (*model.RemovePublicipsFromSharedBandwidthResponse, error) {
	requestDef := GenReqDefForRemovePublicipsFromSharedBandwidth(request)
	resp, responseDef := GenRespForRemovePublicipsFromSharedBandwidth()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询带宽
func (c *EipClient) ShowBandwidth(request *model.ShowBandwidthRequest) (*model.ShowBandwidthResponse, error) {
	requestDef := GenReqDefForShowBandwidth(request)
	resp, responseDef := GenRespForShowBandwidth()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//更新带宽。
func (c *EipClient) UpdateBandwidth(request *model.UpdateBandwidthRequest) (*model.UpdateBandwidthResponse, error) {
	requestDef := GenReqDefForUpdateBandwidth(request)
	resp, responseDef := GenRespForUpdateBandwidth()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//更新带宽。
func (c *EipClient) UpdatePrePaidBandwidth(request *model.UpdatePrePaidBandwidthRequest) (*model.UpdatePrePaidBandwidthResponse, error) {
	requestDef := GenReqDefForUpdatePrePaidBandwidth(request)
	resp, responseDef := GenRespForUpdatePrePaidBandwidth()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//为指定的弹性公网IP资源实例批量添加标签。
func (c *EipClient) BatchCreatePublicipTags(request *model.BatchCreatePublicipTagsRequest) (*model.BatchCreatePublicipTagsResponse, error) {
	requestDef := GenReqDefForBatchCreatePublicipTags(request)
	resp, responseDef := GenRespForBatchCreatePublicipTags()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//为指定的弹性公网IP资源实例批量删除标签。
func (c *EipClient) BatchDeletePublicipTags(request *model.BatchDeletePublicipTagsRequest) (*model.BatchDeletePublicipTagsResponse, error) {
	requestDef := GenReqDefForBatchDeletePublicipTags(request)
	resp, responseDef := GenRespForBatchDeletePublicipTags()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//申请包年包月的弹性公网IP。
func (c *EipClient) CreatePrePaidPublicip(request *model.CreatePrePaidPublicipRequest) (*model.CreatePrePaidPublicipResponse, error) {
	requestDef := GenReqDefForCreatePrePaidPublicip(request)
	resp, responseDef := GenRespForCreatePrePaidPublicip()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//申请弹性公网IP，支持IPv4和IPv6。  弹性公网IP（Elastic IP）提供独立的公网IP资源，包括公网IP地址与公网出口带宽服务。可以与弹性云服务器、裸金属服务器、虚拟IP、弹性负载均衡、NAT网关等资源灵活地绑定及解绑。拥有多种灵活的计费方式，可以满足各种业务场景的需要。
func (c *EipClient) CreatePublicip(request *model.CreatePublicipRequest) (*model.CreatePublicipResponse, error) {
	requestDef := GenReqDefForCreatePublicip(request)
	resp, responseDef := GenRespForCreatePublicip()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//给指定弹性IP资源实例增加标签信息。
func (c *EipClient) CreatePublicipTag(request *model.CreatePublicipTagRequest) (*model.CreatePublicipTagResponse, error) {
	requestDef := GenReqDefForCreatePublicipTag(request)
	resp, responseDef := GenRespForCreatePublicipTag()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除弹性公网IP。
func (c *EipClient) DeletePublicip(request *model.DeletePublicipRequest) (*model.DeletePublicipResponse, error) {
	requestDef := GenReqDefForDeletePublicip(request)
	resp, responseDef := GenRespForDeletePublicip()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除指定弹性公网IP的标签信息。其中project_id是项目ID，publicip_id 是要操作的弹性公网IP的id。key是要删除标签的键。
func (c *EipClient) DeletePublicipTag(request *model.DeletePublicipTagRequest) (*model.DeletePublicipTagResponse, error) {
	requestDef := GenReqDefForDeletePublicipTag(request)
	resp, responseDef := GenRespForDeletePublicipTag()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询租户在指定区域和实例类型的所有标签集合。
func (c *EipClient) ListPublicipTags(request *model.ListPublicipTagsRequest) (*model.ListPublicipTagsResponse, error) {
	requestDef := GenReqDefForListPublicipTags(request)
	resp, responseDef := GenRespForListPublicipTags()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询弹性公网IP列表
func (c *EipClient) ListPublicips(request *model.ListPublicipsRequest) (*model.ListPublicipsResponse, error) {
	requestDef := GenReqDefForListPublicips(request)
	resp, responseDef := GenRespForListPublicips()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//使用标签过滤弹性公网IP。
func (c *EipClient) ListPublicipsByTags(request *model.ListPublicipsByTagsRequest) (*model.ListPublicipsByTagsResponse, error) {
	requestDef := GenReqDefForListPublicipsByTags(request)
	resp, responseDef := GenRespForListPublicipsByTags()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询指定的弹性公网IP。
func (c *EipClient) ShowPublicip(request *model.ShowPublicipRequest) (*model.ShowPublicipResponse, error) {
	requestDef := GenReqDefForShowPublicip(request)
	resp, responseDef := GenRespForShowPublicip()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询指定弹性IP实例的标签信息。
func (c *EipClient) ShowPublicipTags(request *model.ShowPublicipTagsRequest) (*model.ShowPublicipTagsResponse, error) {
	requestDef := GenReqDefForShowPublicipTags(request)
	resp, responseDef := GenRespForShowPublicipTags()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//更新弹性公网IP，将弹性公网IP跟一个网卡绑定或者解绑定，转换IP地址版本类型。
func (c *EipClient) UpdatePublicip(request *model.UpdatePublicipRequest) (*model.UpdatePublicipResponse, error) {
	requestDef := GenReqDefForUpdatePublicip(request)
	resp, responseDef := GenRespForUpdatePublicip()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//创建浮动IP的外部网络UUID，请使用GET /v2.0/networks?router:external=True或neutron net-external-list方式获取。
func (c *EipClient) NeutronCreateFloatingIp(request *model.NeutronCreateFloatingIpRequest) (*model.NeutronCreateFloatingIpResponse, error) {
	requestDef := GenReqDefForNeutronCreateFloatingIp(request)
	resp, responseDef := GenRespForNeutronCreateFloatingIp()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除指定的浮动IP。
func (c *EipClient) NeutronDeleteFloatingIp(request *model.NeutronDeleteFloatingIpRequest) (*model.NeutronDeleteFloatingIpResponse, error) {
	requestDef := GenReqDefForNeutronDeleteFloatingIp(request)
	resp, responseDef := GenRespForNeutronDeleteFloatingIp()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询提交请求的租户有权限操作的所有浮动IP地址。
func (c *EipClient) NeutronListFloatingIps(request *model.NeutronListFloatingIpsRequest) (*model.NeutronListFloatingIpsResponse, error) {
	requestDef := GenReqDefForNeutronListFloatingIps(request)
	resp, responseDef := GenRespForNeutronListFloatingIps()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询浮动IP详情，包括浮动IP状态，浮动IP所属路由器ID，浮动IP的外部网络ID等等。
func (c *EipClient) NeutronShowFloatingIp(request *model.NeutronShowFloatingIpRequest) (*model.NeutronShowFloatingIpResponse, error) {
	requestDef := GenReqDefForNeutronShowFloatingIp(request)
	resp, responseDef := GenRespForNeutronShowFloatingIp()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//更新浮动IP。  更新时需在URL中给出浮动IP地址的ID。  port_id 为空，则表示浮动IP从端口解绑。
func (c *EipClient) NeutronUpdateFloatingIp(request *model.NeutronUpdateFloatingIpRequest) (*model.NeutronUpdateFloatingIpResponse, error) {
	requestDef := GenReqDefForNeutronUpdateFloatingIp(request)
	resp, responseDef := GenRespForNeutronUpdateFloatingIp()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}
