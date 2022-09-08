package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ChangeAvailabilityZoneRequest struct {

	// 企业路由器实例ID
	ErId string `json:"er_id"`

	Body *EnterpriseRouterAz `json:"body,omitempty"`
}

func (o ChangeAvailabilityZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAvailabilityZoneRequest struct{}"
	}

	return strings.Join([]string{"ChangeAvailabilityZoneRequest", string(data)}, " ")
}
