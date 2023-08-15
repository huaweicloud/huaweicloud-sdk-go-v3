package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainSettingsRequest Request Object
type ShowDomainSettingsRequest struct {

	// 网站域名ID
	DomainId string `json:"domain_id"`
}

func (o ShowDomainSettingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainSettingsRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainSettingsRequest", string(data)}, " ")
}
