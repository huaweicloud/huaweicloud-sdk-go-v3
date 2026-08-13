package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangePaymentModeNewRequest Request Object
type ChangePaymentModeNewRequest struct {
	Body *ChangePaymentModeForConsoleBody `json:"body,omitempty"`
}

func (o ChangePaymentModeNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangePaymentModeNewRequest struct{}"
	}

	return strings.Join([]string{"ChangePaymentModeNewRequest", string(data)}, " ")
}
