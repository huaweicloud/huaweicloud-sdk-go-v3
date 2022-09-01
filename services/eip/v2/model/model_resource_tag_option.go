package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 标签
type ResourceTagOption struct {

	// 标签名称。不能为空。长度不超过36个字符。由英文字母、数字、下划线、中划线、中文字符组成。同一资源的key值不能重复。
	Key string `json:"key" xml:"key"`

	// 值列表。长度不超过43个字符。由英文字母、数字、下划线、点、中划线、中文字符组成。
	Value string `json:"value" xml:"value"`
}

func (o ResourceTagOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTagOption struct{}"
	}

	return strings.Join([]string{"ResourceTagOption", string(data)}, " ")
}
