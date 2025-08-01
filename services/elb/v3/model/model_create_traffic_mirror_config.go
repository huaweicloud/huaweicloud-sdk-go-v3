package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTrafficMirrorConfig 参数解释：流量镜像的配置。
type CreateTrafficMirrorConfig struct {

	// 流量镜像的目的后端服务器组ID。
	TargetIds *[]string `json:"target_ids,omitempty"`

	// 镜像请求是否携带请求体，默认true。
	MirrorRequestBodyEnable *bool `json:"mirror_request_body_enable,omitempty"`
}

func (o CreateTrafficMirrorConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrafficMirrorConfig struct{}"
	}

	return strings.Join([]string{"CreateTrafficMirrorConfig", string(data)}, " ")
}
