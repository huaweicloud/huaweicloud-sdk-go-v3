package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagDto 键值对
type TagDto struct {

	// 标签的键
	Key string `json:"key"`

	// 标签的值，可以为空字符串，但不能为null
	Value string `json:"value"`
}

func (o TagDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagDto struct{}"
	}

	return strings.Join([]string{"TagDto", string(data)}, " ")
}
