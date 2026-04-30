package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchTokenResponseTokenApp 应用信息
type SwitchTokenResponseTokenApp struct {

	// 应用名称
	AppName *string `json:"appName,omitempty"`

	// 应用ID
	AppId *string `json:"appId,omitempty"`

	// 应用状态
	Status *int32 `json:"status,omitempty"`
}

func (o SwitchTokenResponseTokenApp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchTokenResponseTokenApp struct{}"
	}

	return strings.Join([]string{"SwitchTokenResponseTokenApp", string(data)}, " ")
}
