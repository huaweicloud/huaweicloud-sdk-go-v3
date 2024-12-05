package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBindingGeip Returns geip and its binding status
type ListBindingGeip struct {

	// geip的id
	GlobalEipId *string `json:"global_eip_id,omitempty"`

	// 网段geip的id
	GlobalEipSegmentId *string `json:"global_eip_segment_id,omitempty"`

	// geip的绑定状态
	Status *string `json:"status,omitempty"`

	// geip类型：IP_ADDRESS/IP_SEGMENT
	Type *string `json:"type,omitempty"`

	// geip绑定失败的原因
	ErrorMessage *string `json:"error_message,omitempty"`

	// geip的地址ip/mask
	Cidr *string `json:"cidr,omitempty"`

	// geip的地址簇
	AddressFamily *string `json:"address_family,omitempty"`

	// IES的集群vtepIp
	IeVtepIp *string `json:"ie_vtep_ip,omitempty"`

	// geip绑定时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 带宽包的id
	GcbId *string `json:"gcb_id,omitempty"`
}

func (o ListBindingGeip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBindingGeip struct{}"
	}

	return strings.Join([]string{"ListBindingGeip", string(data)}, " ")
}
