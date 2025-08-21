package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FlowDetailsVo struct {

	// **参数解释**： 应用统计 **取值范围**： 不涉及
	Apps *[]ItemVo `json:"apps,omitempty"`

	// **参数解释**： 关联资产类型 **取值范围**： 不涉及
	AssociateInstanceType *string `json:"associate_instance_type,omitempty"`

	// **参数解释**： 关联资产名称 **取值范围**： 不涉及
	DeviceName *string `json:"device_name,omitempty"`

	// **参数解释**： 聚合项 **取值范围**： 不涉及
	Item *string `json:"item,omitempty"`

	// **参数解释**： 最近访问时间 **取值范围**： 不涉及
	LastTime *int64 `json:"last_time,omitempty"`

	// **参数解释**： 端口统计 **取值范围**： 不涉及
	Ports *[]ItemVo `json:"ports,omitempty"`

	// **参数解释**： 地区 **取值范围**： 不涉及
	Region *string `json:"region,omitempty"`

	// **参数解释**： 请求字节数 **取值范围**： 不涉及
	RequestByte *float64 `json:"request_byte,omitempty"`

	// **参数解释**： 响应字节数 **取值范围**： 不涉及
	ResponseByte *float64 `json:"response_byte,omitempty"`

	// **参数解释**： 会话数量 **取值范围**： 不涉及
	Sessions *int64 `json:"sessions,omitempty"`

	// **参数解释**： 标签 **取值范围**： 不涉及
	Tags *string `json:"tags,omitempty"`

	// **参数解释**： 源IP **取值范围**： 不涉及
	SrcIp *[]ItemVo `json:"src_ip,omitempty"`

	// **参数解释**： 目的IP **取值范围**： 不涉及
	DstIp *[]ItemVo `json:"dst_ip,omitempty"`

	// **参数解释**： 协议 **取值范围**： 不涉及
	Protocol *string `json:"protocol,omitempty"`
}

func (o FlowDetailsVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlowDetailsVo struct{}"
	}

	return strings.Join([]string{"FlowDetailsVo", string(data)}, " ")
}
