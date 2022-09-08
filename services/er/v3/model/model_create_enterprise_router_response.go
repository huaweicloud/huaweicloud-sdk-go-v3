package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateEnterpriseRouterResponse struct {
	Instance *EnterpriseRouter `json:"instance,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	XClientToken   *string `json:"X-Client-Token,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEnterpriseRouterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnterpriseRouterResponse struct{}"
	}

	return strings.Join([]string{"CreateEnterpriseRouterResponse", string(data)}, " ")
}
