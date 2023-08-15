package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHostedDirectConnectResponse Response Object
type UpdateHostedDirectConnectResponse struct {

	// 操作请求ID
	RequestId *string `json:"request_id,omitempty"`

	HostedConnect  *HostedDirectConnect `json:"hosted_connect,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o UpdateHostedDirectConnectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostedDirectConnectResponse struct{}"
	}

	return strings.Join([]string{"UpdateHostedDirectConnectResponse", string(data)}, " ")
}
