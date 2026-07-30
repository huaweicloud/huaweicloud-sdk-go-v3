package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ErrorInfo struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误消息
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 规则的配置名称
	SettingName *string `json:"setting_name,omitempty"`

	// region的ID
	RegionId *string `json:"region_id,omitempty"`
}

func (o ErrorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorInfo struct{}"
	}

	return strings.Join([]string{"ErrorInfo", string(data)}, " ")
}
