package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExportConnectionsRequest struct {
}

func (o ExportConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ExportConnectionsRequest", string(data)}, " ")
}
