package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDirectConnectLocationResponse Response Object
type ShowDirectConnectLocationResponse struct {
	DirectConnectLocation *DirectConnectLocationEntry `json:"direct_connect_location,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDirectConnectLocationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDirectConnectLocationResponse struct{}"
	}

	return strings.Join([]string{"ShowDirectConnectLocationResponse", string(data)}, " ")
}
