package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 标签
type Tag struct {

	// 标签键
	Key *string `json:"key,omitempty" xml:"key"`

	// 标签值
	Value *string `json:"value,omitempty" xml:"value"`
}

func (o Tag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
