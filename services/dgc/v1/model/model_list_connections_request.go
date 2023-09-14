package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConnectionsRequest Request Object
type ListConnectionsRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 分页参数：每页限定数量。范围[1,100]
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数：页数
	Offset *int32 `json:"offset,omitempty"`

	// 连接名称.
	ConnectionName *string `json:"connectionName,omitempty"`
}

func (o ListConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ListConnectionsRequest", string(data)}, " ")
}
