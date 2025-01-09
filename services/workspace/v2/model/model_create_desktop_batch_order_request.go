package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDesktopBatchOrderRequest Request Object
type CreateDesktopBatchOrderRequest struct {
	Body *CreateBatchChangeOrderRequestBody `json:"body,omitempty"`
}

func (o CreateDesktopBatchOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesktopBatchOrderRequest struct{}"
	}

	return strings.Join([]string{"CreateDesktopBatchOrderRequest", string(data)}, " ")
}
