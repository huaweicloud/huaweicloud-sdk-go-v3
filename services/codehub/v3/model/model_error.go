package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Error struct {

	// 错误码
	Code *string `json:"code,omitempty" xml:"code"`

	// 错误信息
	Message *string `json:"message,omitempty" xml:"message"`
}

func (o Error) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Error struct{}"
	}

	return strings.Join([]string{"Error", string(data)}, " ")
}
