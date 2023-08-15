package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagsBody 标签字段信息
type TagsBody struct {

	// 标签键
	Key *string `json:"key,omitempty"`

	// 标签值
	Value *string `json:"value,omitempty"`
}

func (o TagsBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsBody struct{}"
	}

	return strings.Join([]string{"TagsBody", string(data)}, " ")
}
