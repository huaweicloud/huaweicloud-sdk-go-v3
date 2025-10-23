package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagsItem 标签的返回体
type TagsItem struct {

	// 标签的key
	Key *string `json:"key,omitempty"`

	// 标签的取值
	Values *[]string `json:"values,omitempty"`
}

func (o TagsItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsItem struct{}"
	}

	return strings.Join([]string{"TagsItem", string(data)}, " ")
}
