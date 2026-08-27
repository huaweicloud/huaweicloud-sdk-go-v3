package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateModelResponse Response Object
type BatchCreateModelResponse struct {

	// 已创建的模型列表。
	CreatedModels *[]ModelItemResp `json:"created_models,omitempty"`

	// 创建总数。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchCreateModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateModelResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateModelResponse", string(data)}, " ")
}
