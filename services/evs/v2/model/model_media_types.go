package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MediaTypes 接口版本的请求消息类型信息
type MediaTypes struct {

	// 文本类型
	Base string `json:"base"`

	// 返回类型
	Type string `json:"type"`
}

func (o MediaTypes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MediaTypes struct{}"
	}

	return strings.Join([]string{"MediaTypes", string(data)}, " ")
}
