package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAiGuardLtsConfigRequest Request Object
type UpdateAiGuardLtsConfigRequest struct {
	Body *UpdateAiGuardLtsConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateAiGuardLtsConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAiGuardLtsConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateAiGuardLtsConfigRequest", string(data)}, " ")
}
