package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLlmConfigResponse Response Object
type ShowLlmConfigResponse struct {

	// 大语言模型配置ID。
	LlmConfigId *string `json:"llm_config_id,omitempty"`

	// 大语言模型配置名称。
	Name *string `json:"name,omitempty"`

	// 大语言模型地址。
	LlmUrl *string `json:"llm_url,omitempty"`

	// model参数
	Model *string `json:"model,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowLlmConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLlmConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowLlmConfigResponse", string(data)}, " ")
}
