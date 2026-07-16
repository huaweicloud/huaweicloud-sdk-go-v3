package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSearchAlgorithmsParams struct {

	// 超参搜索算法的参数名称。
	Key *string `json:"key,omitempty"`

	// 超参搜索算法的参数取值。
	Value *string `json:"value,omitempty"`

	// 超参搜索算法的参数类型。
	Type *string `json:"type,omitempty"`
}

func (o ListSearchAlgorithmsParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSearchAlgorithmsParams struct{}"
	}

	return strings.Join([]string{"ListSearchAlgorithmsParams", string(data)}, " ")
}
