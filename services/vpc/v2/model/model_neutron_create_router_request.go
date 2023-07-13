package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronCreateRouterRequest Request Object
type NeutronCreateRouterRequest struct {
	Body *NeutronCreateRouterRequestBody `json:"body,omitempty"`
}

func (o NeutronCreateRouterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateRouterRequest struct{}"
	}

	return strings.Join([]string{"NeutronCreateRouterRequest", string(data)}, " ")
}
