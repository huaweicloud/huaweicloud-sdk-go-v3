package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainDetailRequest Request Object
type ShowDomainDetailRequest struct {

	// 域名ID
	DomainId string `json:"domain_id"`
}

func (o ShowDomainDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainDetailRequest", string(data)}, " ")
}
