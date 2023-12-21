package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InterpreterGroupInfo 传译组信息
type InterpreterGroupInfo struct {

	// 传译组序号。
	GroupID string `json:"groupID"`

	// 传译组名称。
	GroupName *string `json:"groupName,omitempty"`

	// 传译组支持的第一种语言。
	FirstLanguage string `json:"firstLanguage"`

	// 传译组支持的第二种语言。
	SecondLanguage string `json:"secondLanguage"`

	// 传译员列表。
	Interpreters *[]InterpreterInfo `json:"interpreters,omitempty"`
}

func (o InterpreterGroupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InterpreterGroupInfo struct{}"
	}

	return strings.Join([]string{"InterpreterGroupInfo", string(data)}, " ")
}
