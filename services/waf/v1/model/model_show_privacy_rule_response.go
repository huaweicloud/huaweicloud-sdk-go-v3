package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPrivacyRuleResponse Response Object
type ShowPrivacyRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 创建规则的时间，格式为13位毫秒时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 规则描述，可选参数，设置该规则的备注信息。
	Description *string `json:"description,omitempty"`

	// 规则状态，0：关闭，1：开启
	Status *int32 `json:"status,omitempty"`

	// 隐私屏蔽规则防护的url，需要填写标准的url格式，例如/admin/xxx或者/admin/_*,以\"*\"号结尾代表路径前缀
	Url *string `json:"url,omitempty"`

	// 屏蔽字段   - Params：请求参数   - Cookie：根据Cookie区分的Web访问者   - Header：自定义HTTP首部   - Form：表单参数
	Category *string `json:"category,omitempty"`

	// 屏蔽字段名
	Index          *string `json:"index,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPrivacyRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivacyRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowPrivacyRuleResponse", string(data)}, " ")
}
