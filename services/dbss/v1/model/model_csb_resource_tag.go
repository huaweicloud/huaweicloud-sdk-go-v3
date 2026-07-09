package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CsbResourceTag struct {

	// 标签KEY
	Key *string `json:"key,omitempty"`

	// 标签value
	Value *string `json:"value,omitempty"`
}

func (o CsbResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CsbResourceTag struct{}"
	}

	return strings.Join([]string{"CsbResourceTag", string(data)}, " ")
}
