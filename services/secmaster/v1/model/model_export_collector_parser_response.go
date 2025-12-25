package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportCollectorParserResponse Response Object
type ExportCollectorParserResponse struct {

	// 相关描述信息
	ParserIds      *[]ExportParserResponseDto `json:"parser_ids,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ExportCollectorParserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportCollectorParserResponse struct{}"
	}

	return strings.Join([]string{"ExportCollectorParserResponse", string(data)}, " ")
}
