package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateHostedDirectConnectRequest struct {

	// 托管专线连接ID。
	HostedConnectId string `json:"hosted_connect_id"`

	Body *UpdateHostedDirectConnectRequestBody `json:"body,omitempty"`
}

func (o UpdateHostedDirectConnectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostedDirectConnectRequest struct{}"
	}

	return strings.Join([]string{"UpdateHostedDirectConnectRequest", string(data)}, " ")
}
