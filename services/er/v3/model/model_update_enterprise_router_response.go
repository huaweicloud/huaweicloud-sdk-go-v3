package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEnterpriseRouterResponse Response Object
type UpdateEnterpriseRouterResponse struct {
	Instance *EnterpriseRouter `json:"instance,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateEnterpriseRouterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEnterpriseRouterResponse struct{}"
	}

	return strings.Join([]string{"UpdateEnterpriseRouterResponse", string(data)}, " ")
}
