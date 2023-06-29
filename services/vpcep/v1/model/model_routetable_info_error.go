package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RoutetableInfoError 当修改终端节点子网路由表失败时，返回错误提示信息
type RoutetableInfoError struct {

	// 绑定终端节点子网路由表失败信息。
	BindFailed *[]RoutetableInfoErrorDetial `json:"bind_failed,omitempty"`

	// 解绑终端节点子网路由表失败信息。
	UnbindFailed *[]RoutetableInfoErrorDetial `json:"unbind_failed,omitempty"`
}

func (o RoutetableInfoError) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoutetableInfoError struct{}"
	}

	return strings.Join([]string{"RoutetableInfoError", string(data)}, " ")
}
