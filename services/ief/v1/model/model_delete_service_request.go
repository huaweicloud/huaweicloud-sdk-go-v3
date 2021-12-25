package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteServiceRequest struct {
	// 服务ID

	ServiceId string `json:"service_id"`
	// 铂金版实例ID

	IefInstanceId string `json:"ief-instance-id"`
}

func (o DeleteServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServiceRequest struct{}"
	}

	return strings.Join([]string{"DeleteServiceRequest", string(data)}, " ")
}
