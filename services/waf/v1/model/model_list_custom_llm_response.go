package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomLlmResponse Response Object
type ListCustomLlmResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 自定义大模型列表
	Items          *[]CustomLlmItem `json:"items,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListCustomLlmResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomLlmResponse struct{}"
	}

	return strings.Join([]string{"ListCustomLlmResponse", string(data)}, " ")
}
