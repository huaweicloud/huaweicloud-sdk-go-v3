package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCustomConnectorRequest Request Object
type ShowCustomConnectorRequest struct {

	// 连接器ID
	ConnectorId string `json:"connector_id"`
}

func (o ShowCustomConnectorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCustomConnectorRequest struct{}"
	}

	return strings.Join([]string{"ShowCustomConnectorRequest", string(data)}, " ")
}
