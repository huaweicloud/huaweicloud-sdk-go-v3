package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBinlogFilesRequest Request Object
type ListBinlogFilesRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	Body *ListBinlogFilesRequestBody `json:"body,omitempty"`
}

func (o ListBinlogFilesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBinlogFilesRequest struct{}"
	}

	return strings.Join([]string{"ListBinlogFilesRequest", string(data)}, " ")
}
