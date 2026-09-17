package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportBinlogRequest Request Object
type ExportBinlogRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	Body *ExportBinlogRequestBody `json:"body,omitempty"`
}

func (o ExportBinlogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportBinlogRequest struct{}"
	}

	return strings.Join([]string{"ExportBinlogRequest", string(data)}, " ")
}
