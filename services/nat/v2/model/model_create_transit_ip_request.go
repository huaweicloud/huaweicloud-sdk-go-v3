package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTransitIpRequest Request Object
type CreateTransitIpRequest struct {
	Body *CreateTransitIpRequestBody `json:"body,omitempty"`
}

func (o CreateTransitIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransitIpRequest struct{}"
	}

	return strings.Join([]string{"CreateTransitIpRequest", string(data)}, " ")
}
