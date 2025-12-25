package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResumeShipperRequest Request Object
type ResumeShipperRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 投递ID
	ShipperId string `json:"shipper_id"`
}

func (o ResumeShipperRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResumeShipperRequest struct{}"
	}

	return strings.Join([]string{"ResumeShipperRequest", string(data)}, " ")
}
