package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SearchQueryResult struct {

	// 时间戳
	Timestamp string `json:"timestamp"`

	// 数据源
	DataSource *interface{} `json:"data_source"`
}

func (o SearchQueryResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchQueryResult struct{}"
	}

	return strings.Join([]string{"SearchQueryResult", string(data)}, " ")
}
