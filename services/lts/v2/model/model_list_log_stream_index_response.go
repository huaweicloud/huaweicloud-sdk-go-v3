package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogStreamIndexResponse Response Object
type ListLogStreamIndexResponse struct {
	FullTextIndex *LtsFullTextIndexInfo `json:"fullTextIndex,omitempty"`

	// 字段索引配置
	Fields *[]LtsFieldsInfo `json:"fields,omitempty"`

	// 日志流id
	LogStreamId    *string `json:"logStreamId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListLogStreamIndexResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogStreamIndexResponse struct{}"
	}

	return strings.Join([]string{"ListLogStreamIndexResponse", string(data)}, " ")
}
