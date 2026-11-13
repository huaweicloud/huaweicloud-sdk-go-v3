package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateTrafficMirrorFilterRulesRequest Request Object
type BatchCreateTrafficMirrorFilterRulesRequest struct {

	// 流量镜像筛选条件ID
	TrafficMirrorFilterId string `json:"traffic_mirror_filter_id"`

	Body *BatchCreateTrafficMirrorFilterRulesRequestBody `json:"body,omitempty"`
}

func (o BatchCreateTrafficMirrorFilterRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateTrafficMirrorFilterRulesRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateTrafficMirrorFilterRulesRequest", string(data)}, " ")
}
