package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateTrafficMirrorFilterRulesRequestBody
type BatchCreateTrafficMirrorFilterRulesRequestBody struct {

	// 流量镜像筛选条件规则对象列表
	TrafficMirrorFilterRules []BatchCreateTrafficMirrorFilterRuleOption `json:"traffic_mirror_filter_rules"`
}

func (o BatchCreateTrafficMirrorFilterRulesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateTrafficMirrorFilterRulesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateTrafficMirrorFilterRulesRequestBody", string(data)}, " ")
}
