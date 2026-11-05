package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AiguardCustomRegexSetBatchDeleteBody struct {

	// 正则库ID列表
	RegexSetIds []string `json:"regex_set_ids"`
}

func (o AiguardCustomRegexSetBatchDeleteBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiguardCustomRegexSetBatchDeleteBody struct{}"
	}

	return strings.Join([]string{"AiguardCustomRegexSetBatchDeleteBody", string(data)}, " ")
}
