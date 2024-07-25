package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagsPojo struct {

	// 标签key
	Key *string `json:"key,omitempty"`

	// 标签value
	Value *string `json:"value,omitempty"`
}

func (o TagsPojo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsPojo struct{}"
	}

	return strings.Join([]string{"TagsPojo", string(data)}, " ")
}
