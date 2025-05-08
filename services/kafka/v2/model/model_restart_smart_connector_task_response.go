package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartSmartConnectorTaskResponse Response Object
type RestartSmartConnectorTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RestartSmartConnectorTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartSmartConnectorTaskResponse struct{}"
	}

	return strings.Join([]string{"RestartSmartConnectorTaskResponse", string(data)}, " ")
}
