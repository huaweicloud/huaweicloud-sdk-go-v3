package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ChangeInstanceOrderRequest struct {

	// 实例ID
	ServerId string `json:"server_id"`

	// 实例的Instance_id
	InstanceKey string `json:"instance_key"`
}

func (o ChangeInstanceOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeInstanceOrderRequest struct{}"
	}

	return strings.Join([]string{"ChangeInstanceOrderRequest", string(data)}, " ")
}
