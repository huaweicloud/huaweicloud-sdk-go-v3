package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWafPolicyRequest Request Object
type ShowWafPolicyRequest struct {

	// 域名
	DomainName string `json:"domain_name"`

	// 防护区域，0-大陆，1-海外
	OverseasType int32 `json:"overseas_type"`
}

func (o ShowWafPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWafPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowWafPolicyRequest", string(data)}, " ")
}
