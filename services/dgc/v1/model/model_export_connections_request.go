package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportConnectionsRequest Request Object
type ExportConnectionsRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`
}

func (o ExportConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ExportConnectionsRequest", string(data)}, " ")
}
