package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHostMassageRequest Request Object
type UpdateHostMassageRequest struct {

	// 连接ID，用于标识跨源连接的UUID。
	ConnectionId string `json:"connection_id"`

	Body *UpdateHostMassageReq `json:"body,omitempty"`
}

func (o UpdateHostMassageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostMassageRequest struct{}"
	}

	return strings.Join([]string{"UpdateHostMassageRequest", string(data)}, " ")
}
