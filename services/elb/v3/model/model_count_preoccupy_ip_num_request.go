package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountPreoccupyIpNumRequest Request Object
type CountPreoccupyIpNumRequest struct {

	// 负载均衡器七层规格的ID。传入该字段表示计算创建该规格的LB，或变更LB的原七层规格到该规格所需要的预占IP。  适用场景：创建负LB，变更LB规格。  [不支持传入l7_flavor_id](tag:hcso,hk_vdf,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b)
	L7FlavorId *string `json:"l7_flavor_id,omitempty"`

	// 是否开启跨VPC转发。  取值true表示计算创建或变更为开启跨VPC转发的LB的预占IP。  取值false表示计算创建或变更为不开启跨VPC转发的LB的预占IP。不传等价false。  适用场景：创建LB，变更LB规格。  [荷兰region不支持该字段，请勿使用。](tag:dt,dt_test)
	IpTargetEnable *bool `json:"ip_target_enable,omitempty"`

	// 负载均衡器IP地址类型，取值4，6 。  取值4表示计算创建支持IPv4地址的LB的预占IP。  取值6表示计算创建支持IPv6地址的LB的预占IP。  适用场景：创建LB。  [不支持IPv6，请勿设置为6。](tag:dt,dt_test)
	IpVersion *int32 `json:"ip_version,omitempty"`

	// 负载均衡器ID。计算LB规格变更或创建LB中的第一个七层监听器的预占IP。  适用场景：变更LB规格，创建LB中的第一个七层监听器。
	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`

	// 计算创建AZ列表为availability_zone_id的负载局衡器实例的预占IP。  适用场景：创建LB。  使用说明：传入loadbalancer_id时，该参数无效。
	AvailabilityZoneId *[]string `json:"availability_zone_id,omitempty"`
}

func (o CountPreoccupyIpNumRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountPreoccupyIpNumRequest struct{}"
	}

	return strings.Join([]string{"CountPreoccupyIpNumRequest", string(data)}, " ")
}
