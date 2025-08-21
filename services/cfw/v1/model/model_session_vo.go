package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SessionVo struct {

	// **参数解释**： 应用 **取值范围**： 不涉及
	App *string `json:"app,omitempty"`

	// **参数解释**： 字节数 **取值范围**： 不涉及
	Bytes *float64 `json:"bytes,omitempty"`

	// **参数解释**： 目的IP关联资产类型 **取值范围**： 不涉及
	DstAssociateInstanceType *string `json:"dst_associate_instance_type,omitempty"`

	// **参数解释**： 目的IP关联资产名称 **取值范围**： 不涉及
	DstDeviceName *string `json:"dst_device_name,omitempty"`

	// **参数解释**： 目的IP **取值范围**： 不涉及
	DstIp *string `json:"dst_ip,omitempty"`

	// **参数解释**： 目的端口 **取值范围**： 不涉及
	DstPort *string `json:"dst_port,omitempty"`

	// **参数解释**： 目的域名 **取值范围**： 不涉及
	DstHost *string `json:"dst_host,omitempty"`

	// **参数解释**： 目的regionID **取值范围**： 不涉及
	DstRegionId *string `json:"dst_region_id,omitempty"`

	// **参数解释**： 目的地区 **取值范围**： 不涉及
	DstRegionName *string `json:"dst_region_name,omitempty"`

	// **参数解释**： 结束时间 **取值范围**： 不涉及
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**： 协议 **取值范围**： 不涉及
	Protocol *string `json:"protocol,omitempty"`

	// **参数解释**： 请求字节数 **取值范围**： 不涉及
	RequestByte *float64 `json:"request_byte,omitempty"`

	// **参数解释**： 响应字节数 **取值范围**： 不涉及
	ResponseByte *float64 `json:"response_byte,omitempty"`

	// **参数解释**： 会话数量 **取值范围**： 不涉及
	Sessions *int64 `json:"sessions,omitempty"`

	// **参数解释**： 源IP关联资产类型 **取值范围**： 不涉及
	SrcAssociateInstanceType *string `json:"src_associate_instance_type,omitempty"`

	// **参数解释**： 源IP关联资产名称 **取值范围**： 不涉及
	SrcDeviceName *string `json:"src_device_name,omitempty"`

	// **参数解释**： 源IP **取值范围**： 不涉及
	SrcIp *string `json:"src_ip,omitempty"`

	// **参数解释**： 源地区 ID **取值范围**： 不涉及
	SrcRegionId *string `json:"src_region_id,omitempty"`

	// **参数解释**： 源地区 **取值范围**： 不涉及
	SrcRegionName *string `json:"src_region_name,omitempty"`

	// **参数解释**： 会话开始时间 **取值范围**： 不涉及
	StartTime *int64 `json:"start_time,omitempty"`
}

func (o SessionVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SessionVo struct{}"
	}

	return strings.Join([]string{"SessionVo", string(data)}, " ")
}
