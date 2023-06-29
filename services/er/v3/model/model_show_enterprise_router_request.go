package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEnterpriseRouterRequest Request Object
type ShowEnterpriseRouterRequest struct {

	// 企业路由器实例ID
	ErId string `json:"er_id"`
}

func (o ShowEnterpriseRouterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEnterpriseRouterRequest struct{}"
	}

	return strings.Join([]string{"ShowEnterpriseRouterRequest", string(data)}, " ")
}
