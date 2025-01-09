package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MapObject struct {

	// 配置项的键
	Key *string `json:"key,omitempty"`

	// 配置项对应的值
	Value *string `json:"value,omitempty"`
}

func (o MapObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MapObject struct{}"
	}

	return strings.Join([]string{"MapObject", string(data)}, " ")
}
