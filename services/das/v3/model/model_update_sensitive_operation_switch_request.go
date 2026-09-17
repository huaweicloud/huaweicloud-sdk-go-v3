package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSensitiveOperationSwitchRequest Request Object
type UpdateSensitiveOperationSwitchRequest struct {
	Body *UpdateSensitiveOperationSwitchRequestBody `json:"body,omitempty"`
}

func (o UpdateSensitiveOperationSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSensitiveOperationSwitchRequest struct{}"
	}

	return strings.Join([]string{"UpdateSensitiveOperationSwitchRequest", string(data)}, " ")
}
