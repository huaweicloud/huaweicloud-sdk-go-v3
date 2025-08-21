package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccessDetailVo struct {

	// **参数解释**： 目的IP数量 **取值范围**： 不涉及
	DstIpCount *int64 `json:"dst_ip_count,omitempty"`

	// **参数解释**： 目的端口数量 **取值范围**： 不涉及
	DstPortCount *int64 `json:"dst_port_count,omitempty"`

	// **参数解释**： 命中次数 **取值范围**： 不涉及
	HitCount *int64 `json:"hit_count,omitempty"`

	// **参数解释**： 协议数量 **取值范围**： 不涉及
	ProtocolCount *int64 `json:"protocol_count,omitempty"`

	// **参数解释**： 结束时间 **取值范围**： 不涉及
	RecentEndTime *int64 `json:"recent_end_time,omitempty"`

	// **参数解释**： 开始时间 **取值范围**： 不涉及
	RecentStartTime *int64 `json:"recent_start_time,omitempty"`

	// **参数解释**： 记录数量 **取值范围**： 不涉及
	RecordTotal *int64 `json:"record_total,omitempty"`

	// **参数解释**： 命中详情 **取值范围**： 不涉及
	Records *[]AccessLogInfo `json:"records,omitempty"`

	// **参数解释**： 命中规则数 **取值范围**： 不涉及
	RuleHitCount *int64 `json:"rule_hit_count,omitempty"`

	// **参数解释**： 源IP数量 **取值范围**： 不涉及
	SrcIpCount *int64 `json:"src_ip_count,omitempty"`
}

func (o AccessDetailVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessDetailVo struct{}"
	}

	return strings.Join([]string{"AccessDetailVo", string(data)}, " ")
}
