package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateKeyRotationIntervalRequest struct {
	Body *UpdateKeyRotationIntervalRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateKeyRotationIntervalRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeyRotationIntervalRequest struct{}"
	}

	return strings.Join([]string{"UpdateKeyRotationIntervalRequest", string(data)}, " ")
}
