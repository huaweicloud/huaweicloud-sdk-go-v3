package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagCentralNetworkResponse Response Object
type TagCentralNetworkResponse struct {
	XRequestId     *string `json:"x-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o TagCentralNetworkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagCentralNetworkResponse struct{}"
	}

	return strings.Join([]string{"TagCentralNetworkResponse", string(data)}, " ")
}
