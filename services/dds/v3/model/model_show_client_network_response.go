package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClientNetworkResponse Response Object
type ShowClientNetworkResponse struct {

	// 实例ID，可以调用“[查询实例列表和详情](x-wc://file=zh-cn_topic_0000001369935045.xml)”接口获取。如果未申请实例，可以调用“[创建实例](x-wc://file=zh-cn_topic_0000001369734929.xml)”接口创建。
	InstanceId *string `json:"instance_id,omitempty"`

	// 客户端所在网段。 > - [跨网段访问配置只有在客户端与副本集实例部署在不同网段的情况下才允许配置，例如访问副本集的客户端所在网段为192.168.0.0/16，副本集所在的网段为172.16.0.0/24，则需要添加跨网段配置192.168.0.0/16才能正常访问。只有副本集有该功能。](tag:ccs,cmcc,ctc,dt,dt_test,fcs,fcs_dt,g42,hic,hk_g42,hk_sbc,hc,hws_ocb,hws_sbc,ocb,tlf,tm,hk,hws_eu)
	ClientNetworkRanges *[]string `json:"client_network_ranges,omitempty"`
	HttpStatusCode      int       `json:"-"`
}

func (o ShowClientNetworkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClientNetworkResponse struct{}"
	}

	return strings.Join([]string{"ShowClientNetworkResponse", string(data)}, " ")
}
