package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteListenerForceRequest struct {

	// 监听器ID。
	ListenerId string `json:"listener_id"`
}

func (o DeleteListenerForceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteListenerForceRequest struct{}"
	}

	return strings.Join([]string{"DeleteListenerForceRequest", string(data)}, " ")
}
