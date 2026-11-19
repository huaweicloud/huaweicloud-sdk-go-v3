package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateTrafficMirrorFilterRulesResponse Response Object
type BatchCreateTrafficMirrorFilterRulesResponse struct {

	// **参数解释**： 流量镜像筛选规则列表的响应体。 **取值范围**： 不涉及。
	TrafficMirrorFilterRules *[]TrafficMirrorFilterRule `json:"traffic_mirror_filter_rules,omitempty"`

	// **参数解释**： 请求ID。 **取值范围**： 不涉及。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchCreateTrafficMirrorFilterRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateTrafficMirrorFilterRulesResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateTrafficMirrorFilterRulesResponse", string(data)}, " ")
}
