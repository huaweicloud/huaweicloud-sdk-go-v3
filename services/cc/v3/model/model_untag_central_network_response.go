package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UntagCentralNetworkResponse Response Object
type UntagCentralNetworkResponse struct {
	XRequestId     *string `json:"x-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UntagCentralNetworkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UntagCentralNetworkResponse struct{}"
	}

	return strings.Join([]string{"UntagCentralNetworkResponse", string(data)}, " ")
}
