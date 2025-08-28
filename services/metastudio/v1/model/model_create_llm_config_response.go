package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLlmConfigResponse Response Object
type CreateLlmConfigResponse struct {

	// 大语言模型配置ID。
	LlmConfigId *string `json:"llm_config_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLlmConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLlmConfigResponse struct{}"
	}

	return strings.Join([]string{"CreateLlmConfigResponse", string(data)}, " ")
}
