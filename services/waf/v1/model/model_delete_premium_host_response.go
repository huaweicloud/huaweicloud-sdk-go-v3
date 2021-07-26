package model

import (
	"encoding/json"

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
	// 防护状态

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
	data, err := json.Marshal(o)
	if err != nil {
		return "DeletePremiumHostResponse struct{}"
	}

	return strings.Join([]string{"DeletePremiumHostResponse", string(data)}, " ")
}
