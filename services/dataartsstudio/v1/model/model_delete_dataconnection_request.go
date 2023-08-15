package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDataconnectionRequest Request Object
type DeleteDataconnectionRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// 数据连接ID
	DataConnectionId string `json:"data_connection_id"`
}

func (o DeleteDataconnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDataconnectionRequest struct{}"
	}

	return strings.Join([]string{"DeleteDataconnectionRequest", string(data)}, " ")
}
