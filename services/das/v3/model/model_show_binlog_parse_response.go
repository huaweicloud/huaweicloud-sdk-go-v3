package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBinlogParseResponse Response Object
type ShowBinlogParseResponse struct {

	// binlog事件概览信息
	EventList      *[]EventEventsDto `json:"event_list,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowBinlogParseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBinlogParseResponse struct{}"
	}

	return strings.Join([]string{"ShowBinlogParseResponse", string(data)}, " ")
}
