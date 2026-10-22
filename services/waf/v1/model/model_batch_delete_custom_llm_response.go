package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteCustomLlmResponse Response Object
type BatchDeleteCustomLlmResponse struct {

	// 已删除的自定义大模型ID列表
	Deleted *[]string `json:"deleted,omitempty"`

	// 删除失败的自定义大模型列表
	Failed         *[]FailedLlm `json:"failed,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o BatchDeleteCustomLlmResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteCustomLlmResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteCustomLlmResponse", string(data)}, " ")
}
