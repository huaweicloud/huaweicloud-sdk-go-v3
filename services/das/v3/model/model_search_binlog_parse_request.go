package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchBinlogParseRequest Request Object
type SearchBinlogParseRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	Body *SearchBinlogParseRequestBody `json:"body,omitempty"`
}

func (o SearchBinlogParseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchBinlogParseRequest struct{}"
	}

	return strings.Join([]string{"SearchBinlogParseRequest", string(data)}, " ")
}
