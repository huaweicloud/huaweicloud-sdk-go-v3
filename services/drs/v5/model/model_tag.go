package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Tag 标签信息体
type Tag struct {

	// 标签键。例如键值对{“aaa”:\"bbb\"}的key为\"aaa\"
	Key string `json:"key"`

	// 标签值。例如键值对{“aaa”:[\"bbb\"]}的values为[\"bbb\"]
	Values []string `json:"values"`
}

func (o Tag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
