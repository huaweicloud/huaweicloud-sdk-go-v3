package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDataConnectorResponse Response Object
type DeleteDataConnectorResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDataConnectorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDataConnectorResponse struct{}"
	}

	return strings.Join([]string{"DeleteDataConnectorResponse", string(data)}, " ")
}
