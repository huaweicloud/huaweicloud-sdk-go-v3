package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 当修改终端节点子网路由表失败 时，返回错误提示信息
type RoutetableInfoError struct {
	BindFailed *RoutetableInfoErrorDetial `json:"bind_failed,omitempty"`

	UnbindFailed *RoutetableInfoErrorDetial `json:"unbind_failed,omitempty"`
}

func (o RoutetableInfoError) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoutetableInfoError struct{}"
	}

	return strings.Join([]string{"RoutetableInfoError", string(data)}, " ")
}
