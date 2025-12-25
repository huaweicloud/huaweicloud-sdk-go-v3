package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryShipperAuthorizationRequest Request Object
type RetryShipperAuthorizationRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 投递ID
	ShipperId string `json:"shipper_id"`

	// 根据param的传参决定，传类型的一级菜单id，如：(lts_group_id)是一串id
	Param *string `json:"param,omitempty"`
}

func (o RetryShipperAuthorizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryShipperAuthorizationRequest struct{}"
	}

	return strings.Join([]string{"RetryShipperAuthorizationRequest", string(data)}, " ")
}
