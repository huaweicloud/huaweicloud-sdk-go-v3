package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetDisasterRecoverySettingsRequest Request Object
type SetDisasterRecoverySettingsRequest struct {
	Body *SetDisasterRecoverySettingsRequestBody `json:"body,omitempty"`
}

func (o SetDisasterRecoverySettingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetDisasterRecoverySettingsRequest struct{}"
	}

	return strings.Join([]string{"SetDisasterRecoverySettingsRequest", string(data)}, " ")
}
