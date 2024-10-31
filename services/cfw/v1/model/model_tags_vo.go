package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagsVo struct {

	// 规则id
	TagId *string `json:"tag_id,omitempty"`

	// 规则标签键
	TagKey *string `json:"tag_key,omitempty"`

	// 规则标签值
	TagValue *string `json:"tag_value,omitempty"`
}

func (o TagsVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsVo struct{}"
	}

	return strings.Join([]string{"TagsVo", string(data)}, " ")
}
