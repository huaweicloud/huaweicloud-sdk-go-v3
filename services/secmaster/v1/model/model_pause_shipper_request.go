package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PauseShipperRequest Request Object
type PauseShipperRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 投递ID
	ShipperId string `json:"shipper_id"`
}

func (o PauseShipperRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PauseShipperRequest struct{}"
	}

	return strings.Join([]string{"PauseShipperRequest", string(data)}, " ")
}
