package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagValue struct {

	// tag键
	Key *string `json:"key,omitempty"`

	// tag值
	Value *[]string `json:"value,omitempty"`
}

func (o TagValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagValue struct{}"
	}

	return strings.Join([]string{"TagValue", string(data)}, " ")
}
