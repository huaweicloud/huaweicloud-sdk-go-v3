package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceInstanceResponseSysTags struct {

	// 键
	Key *string `json:"key,omitempty"`

	// 值
	Value *string `json:"value,omitempty"`
}

func (o ResourceInstanceResponseSysTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceInstanceResponseSysTags struct{}"
	}

	return strings.Join([]string{"ResourceInstanceResponseSysTags", string(data)}, " ")
}
