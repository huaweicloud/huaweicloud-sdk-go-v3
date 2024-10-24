package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LlmModelConfig struct {

	// 类型请从ListBaseModels列举基模型接口响应中获取
	BaseModelType string `json:"base_model_type"`

	// 模型文件路径
	ModelPath string `json:"model_path"`
}

func (o LlmModelConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LlmModelConfig struct{}"
	}

	return strings.Join([]string{"LlmModelConfig", string(data)}, " ")
}
