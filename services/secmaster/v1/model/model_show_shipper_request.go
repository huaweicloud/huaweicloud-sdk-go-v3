package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowShipperRequest Request Object
type ShowShipperRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 投递ID
	ShipperId string `json:"shipper_id"`
}

func (o ShowShipperRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowShipperRequest struct{}"
	}

	return strings.Join([]string{"ShowShipperRequest", string(data)}, " ")
}
