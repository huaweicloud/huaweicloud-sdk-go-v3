package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDatasourceConnectionRequest Request Object
type ShowDatasourceConnectionRequest struct {

	// 连接ID，用于标识跨源连接的UUID。
	ConnectionId string `json:"connection_id"`
}

func (o ShowDatasourceConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatasourceConnectionRequest struct{}"
	}

	return strings.Join([]string{"ShowDatasourceConnectionRequest", string(data)}, " ")
}
