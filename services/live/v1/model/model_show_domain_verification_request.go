package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainVerificationRequest Request Object
type ShowDomainVerificationRequest struct {

	// 直播域名
	Domain string `json:"domain"`
}

func (o ShowDomainVerificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainVerificationRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainVerificationRequest", string(data)}, " ")
}
