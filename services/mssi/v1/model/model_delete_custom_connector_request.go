package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCustomConnectorRequest Request Object
type DeleteCustomConnectorRequest struct {

	// ID of CustomConnector
	ConnectorId string `json:"connector_id"`
}

func (o DeleteCustomConnectorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCustomConnectorRequest struct{}"
	}

	return strings.Join([]string{"DeleteCustomConnectorRequest", string(data)}, " ")
}
