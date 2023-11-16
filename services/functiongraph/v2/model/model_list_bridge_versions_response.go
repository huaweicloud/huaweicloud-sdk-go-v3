package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBridgeVersionsResponse Response Object
type ListBridgeVersionsResponse struct {
	Body           *[]ServiceBridgeVersion `json:"body,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListBridgeVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBridgeVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListBridgeVersionsResponse", string(data)}, " ")
}
