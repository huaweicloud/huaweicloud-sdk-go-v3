package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LtsIndexConfigInfo struct {

	// 日志流ID
	LogStreamId *string `json:"logStreamId,omitempty"`

	FullTextIndex *LtsFullTextIndexInfo `json:"fullTextIndex"`

	// 字段索引配置
	Fields *[]LtsFieldsInfo `json:"fields,omitempty"`
}

func (o LtsIndexConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsIndexConfigInfo struct{}"
	}

	return strings.Join([]string{"LtsIndexConfigInfo", string(data)}, " ")
}
