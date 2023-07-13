package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEnterpriseRouterRequestBody This is a auto create Body Object
type CreateEnterpriseRouterRequestBody struct {
	Instance *CreateEnterpriseRouter `json:"instance"`
}

func (o CreateEnterpriseRouterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnterpriseRouterRequestBody struct{}"
	}

	return strings.Join([]string{"CreateEnterpriseRouterRequestBody", string(data)}, " ")
}
