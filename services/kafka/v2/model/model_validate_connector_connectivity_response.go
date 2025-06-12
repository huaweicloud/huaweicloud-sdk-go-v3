package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateConnectorConnectivityResponse Response Object
type ValidateConnectorConnectivityResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ValidateConnectorConnectivityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateConnectorConnectivityResponse struct{}"
	}

	return strings.Join([]string{"ValidateConnectorConnectivityResponse", string(data)}, " ")
}
