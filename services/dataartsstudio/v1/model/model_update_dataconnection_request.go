package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDataconnectionRequest Request Object
type UpdateDataconnectionRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// 数据连接ID
	DataConnectionId string `json:"data_connection_id"`

	Body *ApigDataSourcesVo `json:"body,omitempty"`
}

func (o UpdateDataconnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDataconnectionRequest struct{}"
	}

	return strings.Join([]string{"UpdateDataconnectionRequest", string(data)}, " ")
}
