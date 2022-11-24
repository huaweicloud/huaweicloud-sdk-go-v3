package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateEnterpriseRouterRequest struct {

	// 幂等性标识
	XClientToken *string `json:"X-Client-Token,omitempty"`

	Body *CreateEnterpriseRouterRequestBody `json:"body,omitempty"`
}

func (o CreateEnterpriseRouterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnterpriseRouterRequest struct{}"
	}

	return strings.Join([]string{"CreateEnterpriseRouterRequest", string(data)}, " ")
}
