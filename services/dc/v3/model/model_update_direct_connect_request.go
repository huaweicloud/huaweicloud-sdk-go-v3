package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateDirectConnectRequest struct {

	// 物理专线连接ID。
	DirectConnectId string `json:"direct_connect_id"`

	Body *UpdateDirectConnectRequestBody `json:"body,omitempty"`
}

func (o UpdateDirectConnectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDirectConnectRequest struct{}"
	}

	return strings.Join([]string{"UpdateDirectConnectRequest", string(data)}, " ")
}
