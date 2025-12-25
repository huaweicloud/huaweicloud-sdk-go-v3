package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCollectorParserTemplateResponse Response Object
type ListCollectorParserTemplateResponse struct {

	// 计数
	Count *int64 `json:"count,omitempty"`

	// 相关描述信息
	Records        *[]ParserTemplate `json:"records,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListCollectorParserTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectorParserTemplateResponse struct{}"
	}

	return strings.Join([]string{"ListCollectorParserTemplateResponse", string(data)}, " ")
}
