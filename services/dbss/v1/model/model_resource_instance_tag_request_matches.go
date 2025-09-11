package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceInstanceTagRequestMatches struct {

	// 键，目前仅支持：resource_name
	Key *string `json:"key,omitempty"`

	// 值，需要匹配的资源名称
	Value *string `json:"value,omitempty"`
}

func (o ResourceInstanceTagRequestMatches) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceInstanceTagRequestMatches struct{}"
	}

	return strings.Join([]string{"ResourceInstanceTagRequestMatches", string(data)}, " ")
}
