package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Tag 自定义标签信息。
type Tag struct {

	// 标签键，可以包含任意语种字母、数字、空格以及\"_\"、\".\"、\":\"、\"=\"、\"+\"、\"-\"、\"@\"符号的任意组合，但是首尾不能包含空格以及不能使用\"\\_sys\\_\"为开头，长度范围[1,64]。
	TagKey string `json:"tag_key"`

	// 标签值，可以包含任意语种字母、数字、空格以及\"_\"、\".\"、\":\"、\"/\"、\"=\"、\"+\"、\"-\"、\"@\"符号的任意组合，可以是空字符串，长度范围[0,128]。
	TagValue string `json:"tag_value"`
}

func (o Tag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
