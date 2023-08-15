package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceInstanceTagRequestMatches struct {

	// 键
	Key *string `json:"key,omitempty"`

	// 值
	Value *string `json:"value,omitempty"`
}

func (o ResourceInstanceTagRequestMatches) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceInstanceTagRequestMatches struct{}"
	}

	return strings.Join([]string{"ResourceInstanceTagRequestMatches", string(data)}, " ")
}
