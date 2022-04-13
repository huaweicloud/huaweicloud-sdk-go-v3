package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 源实例信息。
type SourceInstanceBody struct {
	// Redis实例名称(source_instance信息中填写)。

	Addrs string `json:"addrs"`
	// Redis密码，如果设置了密码，则必须填写。

	Password *string `json:"password,omitempty"`
}

func (o SourceInstanceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceInstanceBody struct{}"
	}

	return strings.Join([]string{"SourceInstanceBody", string(data)}, " ")
}
