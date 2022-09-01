package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeletePrivacyRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty" xml:"id"`

	// 策略id
	Policyid *string `json:"policyid,omitempty" xml:"policyid"`

	// 创建规则的时间，格式为13位毫秒时间戳
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp"`

	// 规则描述，可选参数，设置该规则的备注信息。
	Description *string `json:"description,omitempty" xml:"description"`

	// 规则状态，0：关闭，1：开启
	Status *int32 `json:"status,omitempty" xml:"status"`

	// 隐私屏蔽规则防护的url，需要填写标准的url格式，例如/admin/xxx或者/admin/_*,以\"*\"号结尾代表路径前缀
	Url *string `json:"url,omitempty" xml:"url"`

	// 屏蔽字段   - Params：请求参数   - Cookie：根据Cookie区分的Web访问者   - Header：自定义HTTP首部   - Form：表单参数
	Category *string `json:"category,omitempty" xml:"category"`

	// 屏蔽字段名
	Index          *string `json:"index,omitempty" xml:"index"`
	HttpStatusCode int     `json:"-"`
}

func (o DeletePrivacyRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivacyRuleResponse struct{}"
	}

	return strings.Join([]string{"DeletePrivacyRuleResponse", string(data)}, " ")
}
