package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEnterpriseRouterResponse Response Object
type ShowEnterpriseRouterResponse struct {
	Instance *EnterpriseRouter `json:"instance,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowEnterpriseRouterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEnterpriseRouterResponse struct{}"
	}

	return strings.Join([]string{"ShowEnterpriseRouterResponse", string(data)}, " ")
}
