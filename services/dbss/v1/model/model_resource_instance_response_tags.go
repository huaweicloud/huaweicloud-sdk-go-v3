package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceInstanceResponseTags struct {

	// 键
	Key *string `json:"key,omitempty"`

	// 值
	Value *string `json:"value,omitempty"`
}

func (o ResourceInstanceResponseTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceInstanceResponseTags struct{}"
	}

	return strings.Join([]string{"ResourceInstanceResponseTags", string(data)}, " ")
}
