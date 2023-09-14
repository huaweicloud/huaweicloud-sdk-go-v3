package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainSettingsRequest Request Object
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
