package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlertConfigRequest Request Object
type UpdateAlertConfigRequest struct {
	Body *UpdateAlertConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateAlertConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlertConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateAlertConfigRequest", string(data)}, " ")
}
