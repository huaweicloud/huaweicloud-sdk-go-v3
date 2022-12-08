package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 注入文件
type Personality struct {

	// 注入路径
	Path string `json:"path"`

	// 注入内容，需要base64格式编码
	Content string `json:"content"`
}

func (o Personality) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Personality struct{}"
	}

	return strings.Join([]string{"Personality", string(data)}, " ")
}
