package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDesktopOrderRequest Request Object
type CreateDesktopOrderRequest struct {
	Body *CreateDesktopOrderRequestBody `json:"body,omitempty"`
}

func (o CreateDesktopOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesktopOrderRequest struct{}"
	}

	return strings.Join([]string{"CreateDesktopOrderRequest", string(data)}, " ")
}
