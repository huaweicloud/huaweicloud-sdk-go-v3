package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RoutetableInfoErrorDetial struct {
	// 路由表ID。

	Id *string `json:"id,omitempty"`
	// 详细错误信息。

	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o RoutetableInfoErrorDetial) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoutetableInfoErrorDetial struct{}"
	}

	return strings.Join([]string{"RoutetableInfoErrorDetial", string(data)}, " ")
}
