package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDrugModelResponse Response Object
type ListDrugModelResponse struct {

	// 模型列表
	Models *[]ModelDto `json:"models,omitempty"`

	// 模型总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDrugModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDrugModelResponse struct{}"
	}

	return strings.Join([]string{"ListDrugModelResponse", string(data)}, " ")
}
