package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ThirdPartyModelConfig 第三方语言模型配置
type ThirdPartyModelConfig struct {

	// 第三方语言模型应用ID。
	AppId *string `json:"app_id,omitempty"`

	// 第三方语言模型应用密钥。
	AppKey *string `json:"app_key,omitempty"`

	// 第三方语言模型地址。
	LlmUrl *string `json:"llm_url,omitempty"`

	// 是否采用流式响应。
	IsStream *bool `json:"is_stream,omitempty"`

	// 支持的多轮对话数量，取值大于1时，请求第三方语言模型时将携带历史对话信息。
	ChatRounds *int32 `json:"chat_rounds,omitempty"`

	// SIS所在区域
	SisRegion *int32 `json:"sis_region,omitempty"`

	// SIS所在区域的projectId
	SisProjectId *string `json:"sis_project_id,omitempty"`

	// 是否开启热词（asr_type选择EI_SIS时生效）
	EnableHotWords *bool `json:"enable_hot_words,omitempty"`

	AsrType *AsrTypeEnum `json:"asr_type,omitempty"`

	// ASR帐号（asr_type选择EI_SIS时不需要配置）。
	AsrAccount *string `json:"asr_account,omitempty"`

	// ASR密钥（asr_type选择EI_SIS时不需要配置）。
	AsrSecret *string `json:"asr_secret,omitempty"`
}

func (o ThirdPartyModelConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThirdPartyModelConfig struct{}"
	}

	return strings.Join([]string{"ThirdPartyModelConfig", string(data)}, " ")
}
