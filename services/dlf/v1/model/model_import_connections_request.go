package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportConnectionsRequest Request Object
type ImportConnectionsRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	Body *ImportConnectionReq `json:"body,omitempty"`
}

func (o ImportConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ImportConnectionsRequest", string(data)}, " ")
}
