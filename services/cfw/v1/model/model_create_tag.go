package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTag struct {

	// 标签键
	Key *string `json:"key,omitempty"`

	// 标签值
	Value *string `json:"value,omitempty"`
}

func (o CreateTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTag struct{}"
	}

	return strings.Join([]string{"CreateTag", string(data)}, " ")
}
