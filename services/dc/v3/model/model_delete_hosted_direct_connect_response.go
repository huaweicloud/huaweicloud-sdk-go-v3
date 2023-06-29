package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHostedDirectConnectResponse Response Object
type DeleteHostedDirectConnectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteHostedDirectConnectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHostedDirectConnectResponse struct{}"
	}

	return strings.Join([]string{"DeleteHostedDirectConnectResponse", string(data)}, " ")
}
