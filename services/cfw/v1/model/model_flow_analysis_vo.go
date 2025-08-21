package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FlowAnalysisVo struct {

	// **参数解释**： 应用统计 **取值范围**： 不涉及
	AppCount *int64 `json:"app_count,omitempty"`

	// **参数解释**： 字节数 **取值范围**： 不涉及
	Bytes *float64 `json:"bytes,omitempty"`

	// **参数解释**： 目的IP计数 **取值范围**： 不涉及
	DstIpCount *int64 `json:"dst_ip_count,omitempty"`

	// **参数**： 目的端口计数 **取值范围**： 不涉及
	DstPortCount *int64 `json:"dst_port_count,omitempty"`

	// **参数解释**： 结束时间 **取值范围**： 不涉及
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**： TOP会话详情 **取值范围**： 不涉及
	Records *[]SessionVo `json:"records,omitempty"`

	// **参数解释**： 请求字节数 **取值范围**： 不涉及
	RequestByte *float64 `json:"request_byte,omitempty"`

	// **参数解释**： 响应字节数 **取值范围**： 不涉及
	ResponseByte *float64 `json:"response_byte,omitempty"`

	// **参数解释**： 会话次数 **取值范围**： 不涉及
	Sessions *int64 `json:"sessions,omitempty"`

	// **参数解释**： 源IP计数 **取值范围**： 不涉及
	SrcIpCount *int64 `json:"src_ip_count,omitempty"`

	// **参数解释**： 开始时间 **取值范围**： 不涉及
	StartTime *int64 `json:"start_time,omitempty"`
}

func (o FlowAnalysisVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlowAnalysisVo struct{}"
	}

	return strings.Join([]string{"FlowAnalysisVo", string(data)}, " ")
}
