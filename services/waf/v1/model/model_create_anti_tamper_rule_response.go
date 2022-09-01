package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateAntiTamperRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty" xml:"id"`

	// 策略id
	Policyid *string `json:"policyid,omitempty" xml:"policyid"`

	// 防篡改的域名
	Hostname *string `json:"hostname,omitempty" xml:"hostname"`

	// 防篡改的url，
	Url *string `json:"url,omitempty" xml:"url"`

	// 创建规则的时间戳
	Description *string `json:"description,omitempty" xml:"description"`

	// 规则状态，0：关闭，1：开启
	Status         *int32 `json:"status,omitempty" xml:"status"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateAntiTamperRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAntiTamperRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateAntiTamperRuleResponse", string(data)}, " ")
}
