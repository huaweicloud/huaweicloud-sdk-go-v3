package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StarRocksInstanceInfoTagsInfoSysTags struct {

	// 标签键。最大长度36个unicode字符。key不能为空或者空字符串，不能为空格。  字符集：A-Z，a-z ，0-9，‘-’，‘_’，UNICODE字符（\\u4E00-\\u9FFF）。
	Key *string `json:"key,omitempty"`

	// 标签值。最大长度43个unicode字符。可以为空字符串。  字符集：A-Z，a-z ，0-9，‘.’，‘-’，‘_’，UNICODE字符（\\u4E00-\\u9FFF）。
	Value *string `json:"value,omitempty"`
}

func (o StarRocksInstanceInfoTagsInfoSysTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StarRocksInstanceInfoTagsInfoSysTags struct{}"
	}

	return strings.Join([]string{"StarRocksInstanceInfoTagsInfoSysTags", string(data)}, " ")
}
