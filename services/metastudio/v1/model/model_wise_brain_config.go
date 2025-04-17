package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WiseBrainConfig 交互助手应用配置
type WiseBrainConfig struct {

	// 角色ID。
	RoleId *string `json:"role_id,omitempty"`

	// SIS所在区域
	SisRegion *int32 `json:"sis_region,omitempty"`

	// SIS所在区域的projectId
	SisProjectId *string `json:"sis_project_id,omitempty"`

	// 是否开启热词
	EnableHotWords *bool `json:"enable_hot_words,omitempty"`
}

func (o WiseBrainConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WiseBrainConfig struct{}"
	}

	return strings.Join([]string{"WiseBrainConfig", string(data)}, " ")
}
