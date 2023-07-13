package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateEnterpriseRouterRequestBody struct {
	Instance *UpdateEnterpriseRouter `json:"instance,omitempty"`
}

func (o UpdateEnterpriseRouterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEnterpriseRouterRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEnterpriseRouterRequestBody", string(data)}, " ")
}
