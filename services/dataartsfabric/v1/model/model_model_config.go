package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelConfig 模型配置
type ModelConfig struct {
	LlmModelConfig *LlmModelConfig `json:"llm_model_config,omitempty"`
}

func (o ModelConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelConfig struct{}"
	}

	return strings.Join([]string{"ModelConfig", string(data)}, " ")
}
