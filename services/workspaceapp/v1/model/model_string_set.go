package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StringSet API批量唯一标识列表  **⚠ 警告: 此Model不生成java代码，允许其它Model中引用，请勿直接作为Model使用**
type StringSet struct {

	// 批量唯一标识请求列表，一次请求数量区间 [1, 20]
	Items []string `json:"items"`
}

func (o StringSet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StringSet struct{}"
	}

	return strings.Join([]string{"StringSet", string(data)}, " ")
}
