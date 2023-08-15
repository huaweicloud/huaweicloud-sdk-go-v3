package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SearchCriteriasBody struct {

	// 单个日志流的快速查询
	Criterias []GetQuerySearchCriteriasBody `json:"criterias"`

	// 日志流id
	LogStreamId *string `json:"log_stream_id,omitempty"`

	// 日志流名称
	LogStreamName *string `json:"log_stream_name,omitempty"`
}

func (o SearchCriteriasBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCriteriasBody struct{}"
	}

	return strings.Join([]string{"SearchCriteriasBody", string(data)}, " ")
}
