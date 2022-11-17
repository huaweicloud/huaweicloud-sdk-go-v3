package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateDirectConnectResponse struct {

	// 操作请求ID
	RequestId *string `json:"request_id,omitempty"`

	DirectConnect  *DirectConnect `json:"direct_connect,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o UpdateDirectConnectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDirectConnectResponse struct{}"
	}

	return strings.Join([]string{"UpdateDirectConnectResponse", string(data)}, " ")
}
