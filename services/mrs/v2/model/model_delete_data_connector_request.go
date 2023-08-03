package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDataConnectorRequest Request Object
type DeleteDataConnectorRequest struct {

	// 数据连接id
	ConnectorId string `json:"connector_id"`
}

func (o DeleteDataConnectorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDataConnectorRequest struct{}"
	}

	return strings.Join([]string{"DeleteDataConnectorRequest", string(data)}, " ")
}
