package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDomainSettingsRequest struct {
	// 域名ID

	DomainId string `json:"domain_id"`
}

func (o ShowDomainSettingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainSettingsRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainSettingsRequest", string(data)}, " ")
}
