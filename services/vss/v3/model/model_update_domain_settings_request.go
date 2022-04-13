package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateDomainSettingsRequest struct {
	Body *UpdateDomainSettingsRequestBody `json:"body,omitempty"`
}

func (o UpdateDomainSettingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainSettingsRequest struct{}"
	}

	return strings.Join([]string{"UpdateDomainSettingsRequest", string(data)}, " ")
}
