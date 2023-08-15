package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceOrderRequest Request Object
type CreateInstanceOrderRequest struct {
	Body *CreateInstanceOrder `json:"body,omitempty"`
}

func (o CreateInstanceOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceOrderRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceOrderRequest", string(data)}, " ")
}
