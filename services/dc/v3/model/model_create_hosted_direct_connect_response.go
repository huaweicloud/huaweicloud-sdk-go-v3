package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateHostedDirectConnectResponse struct {

	// 操作请求ID
	RequestId *string `json:"request_id,omitempty"`

	HostedConnect  *HostedDirectConnect `json:"hosted_connect,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o CreateHostedDirectConnectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHostedDirectConnectResponse struct{}"
	}

	return strings.Join([]string{"CreateHostedDirectConnectResponse", string(data)}, " ")
}
