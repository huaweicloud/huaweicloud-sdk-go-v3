package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLlmConfigReq 创建大语言模型配置请求。
type CreateLlmConfigReq struct {

	// 大语言模型配置名称。
	Name string `json:"name"`

	// 大语言模型地址。
	LlmUrl string `json:"llm_url"`

	// 密钥。
	ApiKey string `json:"api_key"`
}

func (o CreateLlmConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLlmConfigReq struct{}"
	}

	return strings.Join([]string{"CreateLlmConfigReq", string(data)}, " ")
}
