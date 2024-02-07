package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateV2RequestBodySysTags struct {

	// 标签键，最大长度128个unicode字符，格式为大小写字母，数字，中划线“-”，下划线“_”，中文。
	Key string `json:"key"`

	// 标签值，最大长度255个unicode字符，格式为大小写字母，数字，中划线“-”，下划线“_”，点“.”，中文。
	Value string `json:"value"`
}

func (o BatchCreateV2RequestBodySysTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateV2RequestBodySysTags struct{}"
	}

	return strings.Join([]string{"BatchCreateV2RequestBodySysTags", string(data)}, " ")
}
