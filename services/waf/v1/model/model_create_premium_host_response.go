package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreatePremiumHostResponse struct {

	// 域名id
	Id *string `json:"id,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 域名id
	Hostname *string `json:"hostname,omitempty"`

	// 租户id
	Domainid *string `json:"domainid,omitempty"`

	// 项目projectid
	Projectid *string `json:"projectid,omitempty"`

	// http协议
	Protocol       *string `json:"protocol,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePremiumHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePremiumHostResponse struct{}"
	}

	return strings.Join([]string{"CreatePremiumHostResponse", string(data)}, " ")
}
