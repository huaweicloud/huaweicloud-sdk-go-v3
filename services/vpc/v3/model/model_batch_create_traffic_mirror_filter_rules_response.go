package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateTrafficMirrorFilterRulesResponse Response Object
type BatchCreateTrafficMirrorFilterRulesResponse struct {

	// 流量镜像筛选条件规则对象
	TrafficMirrorFilterRules *[]TrafficMirrorFilterRule `json:"traffic_mirror_filter_rules,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 请求ID
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
