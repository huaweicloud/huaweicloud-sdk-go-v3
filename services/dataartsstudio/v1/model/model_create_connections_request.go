package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConnectionsRequest Request Object
type CreateConnectionsRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	Body *ApigDataSourcesVo `json:"body,omitempty"`
}

func (o CreateConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectionsRequest struct{}"
	}

	return strings.Join([]string{"CreateConnectionsRequest", string(data)}, " ")
}
