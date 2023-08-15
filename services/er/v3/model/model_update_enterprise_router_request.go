package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEnterpriseRouterRequest Request Object
type UpdateEnterpriseRouterRequest struct {

	// 企业路由器实例ID
	ErId string `json:"er_id"`

	Body *UpdateEnterpriseRouterRequestBody `json:"body,omitempty"`
}

func (o UpdateEnterpriseRouterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEnterpriseRouterRequest struct{}"
	}

	return strings.Join([]string{"UpdateEnterpriseRouterRequest", string(data)}, " ")
}
