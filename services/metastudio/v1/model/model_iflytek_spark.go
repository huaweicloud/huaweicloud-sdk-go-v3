package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IflytekSpark 科大讯飞星火大模型应用配置
type IflytekSpark struct {

	// 星火大模型应用ID。
	AppId *string `json:"app_id,omitempty"`

	// 星火大模型应用密钥。
	AppKey *string `json:"app_key,omitempty"`

	// 星火大模型API密钥。
	ApiSecret *string `json:"api_secret,omitempty"`

	// 是否为正式环境
	IsProduction *bool `json:"is_production,omitempty"`
}

func (o IflytekSpark) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IflytekSpark struct{}"
	}

	return strings.Join([]string{"IflytekSpark", string(data)}, " ")
}
