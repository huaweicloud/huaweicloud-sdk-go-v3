package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Tag 标签信息
type Tag struct {

	// 标签key
	TagKey *string `json:"tagKey,omitempty"`

	// 标签value
	TagValue *string `json:"tagValue,omitempty"`
}

func (o Tag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
