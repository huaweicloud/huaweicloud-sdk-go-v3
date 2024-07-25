package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagInfo struct {

	// 标识 中文、字母、数字、_或者-，且长度范围[2, 36]
	Key string `json:"key"`

	// 内容 中文、字母、数字、_或者-，且长度范围[2, 36]
	Value string `json:"value"`
}

func (o TagInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagInfo struct{}"
	}

	return strings.Join([]string{"TagInfo", string(data)}, " ")
}
