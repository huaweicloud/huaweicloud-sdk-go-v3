package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MobvoiConfig 奇妙问应用配置
type MobvoiConfig struct {

	// 奇妙问应用帐号。
	AppKey *string `json:"app_key,omitempty"`

	// 奇妙问应用Secret。
	AppSecret *string `json:"app_secret,omitempty"`

	// 奇妙问角色ID。
	RoleId *string `json:"role_id,omitempty"`

	// SIS所在区域
	SisRegion *int32 `json:"sis_region,omitempty"`

	// SIS所在区域的projectId
	SisProjectId *string `json:"sis_project_id,omitempty"`
}

func (o MobvoiConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MobvoiConfig struct{}"
	}

	return strings.Join([]string{"MobvoiConfig", string(data)}, " ")
}
