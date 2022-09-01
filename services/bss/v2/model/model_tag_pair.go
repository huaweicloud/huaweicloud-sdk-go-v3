package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagPair struct {

	// 标签key。
	TagKey *string `json:"tag_key,omitempty" xml:"tag_key"`

	// 标签value。
	TagValue *string `json:"tag_value,omitempty" xml:"tag_value"`
}

func (o TagPair) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagPair struct{}"
	}

	return strings.Join([]string{"TagPair", string(data)}, " ")
}
