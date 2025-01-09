package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDesktopPoolChangeOrderRequest Request Object
type CreateDesktopPoolChangeOrderRequest struct {
	Body *CreateBatchChangeOrderRequestBody `json:"body,omitempty"`
}

func (o CreateDesktopPoolChangeOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesktopPoolChangeOrderRequest struct{}"
	}

	return strings.Join([]string{"CreateDesktopPoolChangeOrderRequest", string(data)}, " ")
}
