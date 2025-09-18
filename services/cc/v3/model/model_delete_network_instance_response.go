package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNetworkInstanceResponse Response Object
type DeleteNetworkInstanceResponse struct {
	XRequestId     *string `json:"x-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteNetworkInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNetworkInstanceResponse struct{}"
	}

	return strings.Join([]string{"DeleteNetworkInstanceResponse", string(data)}, " ")
}
