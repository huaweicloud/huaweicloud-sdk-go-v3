package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteEnterpriseRouterRequest struct {

	// 企业路由器实例ID
	ErId string `json:"er_id"`
}

func (o DeleteEnterpriseRouterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEnterpriseRouterRequest struct{}"
	}

	return strings.Join([]string{"DeleteEnterpriseRouterRequest", string(data)}, " ")
}
