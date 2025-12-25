package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SearchResult struct {

	// 原始日志内容
	DataSource *interface{} `json:"data_source"`

	// 数据接收时间
	Timestamp int64 `json:"timestamp"`
}

func (o SearchResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchResult struct{}"
	}

	return strings.Join([]string{"SearchResult", string(data)}, " ")
}
