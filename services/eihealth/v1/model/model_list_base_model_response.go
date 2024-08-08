package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBaseModelResponse Response Object
type ListBaseModelResponse struct {

	// 模型列表
	Models *[]BaseModelDto `json:"models,omitempty"`

	// 模型总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListBaseModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBaseModelResponse struct{}"
	}

	return strings.Join([]string{"ListBaseModelResponse", string(data)}, " ")
}
