package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteExportTaskNewRequest Request Object
type DeleteExportTaskNewRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	Body *DeleteExportTaskNewRequestBody `json:"body,omitempty"`
}

func (o DeleteExportTaskNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteExportTaskNewRequest struct{}"
	}

	return strings.Join([]string{"DeleteExportTaskNewRequest", string(data)}, " ")
}
