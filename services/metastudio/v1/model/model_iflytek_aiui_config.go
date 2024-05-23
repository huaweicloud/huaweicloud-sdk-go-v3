package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IflytekAiuiConfig 科大讯飞AIUI应用配置
type IflytekAiuiConfig struct {

	// AIUI应用ID。
	AppId *string `json:"app_id,omitempty"`

	// AIUI应用密钥。
	AppKey *string `json:"app_key,omitempty"`

	// AIUI API密钥。
	ApiSecret *string `json:"api_secret,omitempty"`

	// 是否为正式环境
	IsProduction *bool `json:"is_production,omitempty"`
}

func (o IflytekAiuiConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IflytekAiuiConfig struct{}"
	}

	return strings.Join([]string{"IflytekAiuiConfig", string(data)}, " ")
}
