package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeletePremiumHostResponse struct {

	// 域名id
	Id *string `json:"id,omitempty"`

	// 域名
	Hostname *string `json:"hostname,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 区域id
	Region *string `json:"region,omitempty"`

	// 域名防护状态：  - -1：bypass，该域名的请求直接到达其后端服务器，不再经过WAF  - 0：暂停防护，WAF只转发该域名的请求，不做攻击检测  - 1：开启防护，WAF根据您配置的策略进行攻击检测
	ProtectStatus *int32 `json:"protect_status,omitempty"`

	// 接入状态
	AccessStatus *int32 `json:"access_status,omitempty"`

	// 特殊标识
	Flag map[string]string `json:"flag,omitempty"`

	// 特殊模式独享引擎的标识（如elb）
	Mode *string `json:"mode,omitempty"`

	// 特殊模式域名所属独享引擎组
	PoolIds        *[]string `json:"pool_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o DeletePremiumHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePremiumHostResponse struct{}"
	}

	return strings.Join([]string{"DeletePremiumHostResponse", string(data)}, " ")
}
