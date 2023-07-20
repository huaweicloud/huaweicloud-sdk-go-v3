package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDatasourceConnectionRequest Request Object
type DeleteDatasourceConnectionRequest struct {

	// 连接ID，用于标识跨源连接的UUID。
	ConnectionId string `json:"connection_id"`
}

func (o DeleteDatasourceConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatasourceConnectionRequest struct{}"
	}

	return strings.Join([]string{"DeleteDatasourceConnectionRequest", string(data)}, " ")
}
