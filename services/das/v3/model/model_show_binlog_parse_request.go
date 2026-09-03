package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBinlogParseRequest Request Object
type ShowBinlogParseRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	Body *ShowBinlogParseRequestBody `json:"body,omitempty"`
}

func (o ShowBinlogParseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBinlogParseRequest struct{}"
	}

	return strings.Join([]string{"ShowBinlogParseRequest", string(data)}, " ")
}
