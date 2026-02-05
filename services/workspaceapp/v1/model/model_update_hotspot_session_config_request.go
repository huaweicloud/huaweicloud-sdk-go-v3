package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHotspotSessionConfigRequest Request Object
type UpdateHotspotSessionConfigRequest struct {

	// 热点会话迁移配置ID。
	ConfigId string `json:"config_id"`

	Body *UpdateHotspotSessionConfigReq `json:"body,omitempty"`
}

func (o UpdateHotspotSessionConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHotspotSessionConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateHotspotSessionConfigRequest", string(data)}, " ")
}
