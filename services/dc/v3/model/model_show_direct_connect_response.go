package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDirectConnectResponse Response Object
type ShowDirectConnectResponse struct {

	// 操作请求ID
	RequestId *string `json:"request_id,omitempty"`

	DirectConnect  *DirectConnect `json:"direct_connect,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowDirectConnectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDirectConnectResponse struct{}"
	}

	return strings.Join([]string{"ShowDirectConnectResponse", string(data)}, " ")
}
