package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDataconnectionRequest Request Object
type ShowDataconnectionRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// 数据连接ID
	DataConnectionId string `json:"data_connection_id"`
}

func (o ShowDataconnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataconnectionRequest struct{}"
	}

	return strings.Join([]string{"ShowDataconnectionRequest", string(data)}, " ")
}
