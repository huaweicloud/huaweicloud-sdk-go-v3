package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FileUpdate struct {

	// 提交描述。
	Message string `json:"message"`

	// 经base64编码的文件内容。
	Content string `json:"content"`

	// 文件的sha值。
	Sha string `json:"sha"`
}

func (o FileUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileUpdate struct{}"
	}

	return strings.Join([]string{"FileUpdate", string(data)}, " ")
}
