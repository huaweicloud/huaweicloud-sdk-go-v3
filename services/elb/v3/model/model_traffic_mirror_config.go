package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TrafficMirrorConfig **参数解释**：流量镜像的配置
type TrafficMirrorConfig struct {

	// **参数解释**：流量镜像的目的后端服务器组ID。  **取值范围**：不涉及
	TargetIds *[]string `json:"target_ids,omitempty"`

	// **参数解释**：镜像请求是否携带请求体，默认true。  **取值范围**：不涉及
	MirrorRequestBodyEnable *bool `json:"mirror_request_body_enable,omitempty"`
}

func (o TrafficMirrorConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrafficMirrorConfig struct{}"
	}

	return strings.Join([]string{"TrafficMirrorConfig", string(data)}, " ")
}
