package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomLlmBatchDeleteBody struct {

	// 自定义大模型ID列表
	LlmIds []string `json:"llm_ids"`
}

func (o CustomLlmBatchDeleteBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomLlmBatchDeleteBody struct{}"
	}

	return strings.Join([]string{"CustomLlmBatchDeleteBody", string(data)}, " ")
}
