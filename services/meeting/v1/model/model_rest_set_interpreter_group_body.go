package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestSetInterpreterGroupBody 设置传译组请求体
type RestSetInterpreterGroupBody struct {

	// 传译组列表信息
	InterpreterGroups []InterpreterGroupInfo `json:"interpreterGroups"`
}

func (o RestSetInterpreterGroupBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestSetInterpreterGroupBody struct{}"
	}

	return strings.Join([]string{"RestSetInterpreterGroupBody", string(data)}, " ")
}
