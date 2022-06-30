package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Error struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误描述
	Reason *string `json:"reason,omitempty"`
}

func (o Error) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Error struct{}"
	}

	return strings.Join([]string{"Error", string(data)}, " ")
}
