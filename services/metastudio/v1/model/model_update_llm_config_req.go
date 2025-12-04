package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLlmConfigReq 修改大语言模型配置请求。
type UpdateLlmConfigReq struct {

	// 大语言模型配置名称。
	Name *string `json:"name,omitempty"`

	// 大语言模型地址。
	LlmUrl *string `json:"llm_url,omitempty"`

	// 密钥。
	ApiKey *string `json:"api_key,omitempty"`

	// model参数
	Model *string `json:"model,omitempty"`
}

func (o UpdateLlmConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLlmConfigReq struct{}"
	}

	return strings.Join([]string{"UpdateLlmConfigReq", string(data)}, " ")
}
