package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HeaderParameter struct {

	// 是否加密
	IsValueSecret *bool `json:"is_value_secret,omitempty"`

	// header的key值
	Key *string `json:"key,omitempty"`

	// deader的value值
	Value *string `json:"value,omitempty"`
}

func (o HeaderParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HeaderParameter struct{}"
	}

	return strings.Join([]string{"HeaderParameter", string(data)}, " ")
}
