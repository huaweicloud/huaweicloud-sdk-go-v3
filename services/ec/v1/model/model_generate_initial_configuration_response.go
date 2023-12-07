package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GenerateInitialConfigurationResponse Response Object
type GenerateInitialConfigurationResponse struct {

	// 智能企业网关设备ID
	EquipmentId *string `json:"equipment_id,omitempty"`

	// 初始配置URL
	UrlConfigContent *string `json:"url_config_content,omitempty"`

	// 初始配置文件
	ScriptConfigContent *string `json:"script_config_content,omitempty"`

	// URL失效时间
	ExpireAt       *string `json:"expire_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GenerateInitialConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenerateInitialConfigurationResponse struct{}"
	}

	return strings.Join([]string{"GenerateInitialConfigurationResponse", string(data)}, " ")
}
