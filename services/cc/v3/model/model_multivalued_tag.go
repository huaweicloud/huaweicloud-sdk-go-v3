package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MultivaluedTag 对于多个key相同的value，聚合成为key/values
type MultivaluedTag struct {

	// 标签键，最大长度128个unicode字符，格式为大小写字母，数字，中划线“-”，下划线“_”，中文。
	Key string `json:"key"`

	// 相同键的值列表
	Values []string `json:"values"`
}

func (o MultivaluedTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultivaluedTag struct{}"
	}

	return strings.Join([]string{"MultivaluedTag", string(data)}, " ")
}
