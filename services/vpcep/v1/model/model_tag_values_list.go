package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagValuesList struct {
	// 键

	Key string `json:"key"`
	// 值列表

	Values []string `json:"values"`
}

func (o TagValuesList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagValuesList struct{}"
	}

	return strings.Join([]string{"TagValuesList", string(data)}, " ")
}
