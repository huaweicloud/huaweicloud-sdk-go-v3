package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SearchQueryResultHistogram struct {

	// 总数
	Count int64 `json:"count"`

	// 毫秒时间戳
	From int64 `json:"from"`

	// 毫秒时间戳
	To int64 `json:"to"`
}

func (o SearchQueryResultHistogram) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchQueryResultHistogram struct{}"
	}

	return strings.Join([]string{"SearchQueryResultHistogram", string(data)}, " ")
}
