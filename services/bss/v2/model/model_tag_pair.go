package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagPair struct {

	// 标签key。
	TagKey *string `json:"tag_key,omitempty"`

	// 标签value。
	TagValue *string `json:"tag_value,omitempty"`
}

func (o TagPair) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagPair struct{}"
	}

	return strings.Join([]string{"TagPair", string(data)}, " ")
}
