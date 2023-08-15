package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHostedDirectConnectResponse Response Object
type ShowHostedDirectConnectResponse struct {

	// 操作请求ID
	RequestId *string `json:"request_id,omitempty"`

	HostedConnect  *HostedDirectConnect `json:"hosted_connect,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowHostedDirectConnectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHostedDirectConnectResponse struct{}"
	}

	return strings.Join([]string{"ShowHostedDirectConnectResponse", string(data)}, " ")
}
