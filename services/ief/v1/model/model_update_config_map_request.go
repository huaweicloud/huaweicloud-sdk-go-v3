package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateConfigMapRequest struct {
	// 铂金版实例ID，专业版实例为空值

	IefInstanceId *string `json:"ief-instance-id,omitempty"`
	// 配置项ID

	ConfigmapId string `json:"configmap_id"`

	Body *UpdateConfigMaps `json:"body,omitempty"`
}

func (o UpdateConfigMapRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigMapRequest struct{}"
	}

	return strings.Join([]string{"UpdateConfigMapRequest", string(data)}, " ")
}
