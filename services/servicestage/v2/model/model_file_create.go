package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FileCreate struct {
	// 提交描述。

	Message string `json:"message"`
	// 经base64编码的文件内容。

	Content string `json:"content"`
}

func (o FileCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileCreate struct{}"
	}

	return strings.Join([]string{"FileCreate", string(data)}, " ")
}
