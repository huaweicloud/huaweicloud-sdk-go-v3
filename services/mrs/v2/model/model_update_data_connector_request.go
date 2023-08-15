package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDataConnectorRequest Request Object
type UpdateDataConnectorRequest struct {

	// 数据连接id
	ConnectorId string `json:"connector_id"`

	Body *DataConnectorReq `json:"body,omitempty"`
}

func (o UpdateDataConnectorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDataConnectorRequest struct{}"
	}

	return strings.Join([]string{"UpdateDataConnectorRequest", string(data)}, " ")
}
