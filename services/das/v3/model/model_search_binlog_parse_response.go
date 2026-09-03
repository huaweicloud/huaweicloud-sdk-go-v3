package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchBinlogParseResponse Response Object
type SearchBinlogParseResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// binlog详情信息列表
	EventList      *[]EventRowsVo `json:"event_list,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o SearchBinlogParseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchBinlogParseResponse struct{}"
	}

	return strings.Join([]string{"SearchBinlogParseResponse", string(data)}, " ")
}
