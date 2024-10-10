package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WafCustomRule struct {

	// id
	Id *string `json:"id,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// 精准防护规则生效时间。true-自定义生效时间，false-立即生效
	Time bool `json:"time"`

	// 精准防护规则生效的起始时间戳（秒）
	Start *int64 `json:"start,omitempty"`

	// 精准防护规则生效的终止时间戳（秒）
	Terminal *int64 `json:"terminal,omitempty"`

	// 执行该规则的优先级，值越小，优先级越高。取值范围：0到1000。
	Priority int32 `json:"priority"`

	// condition
	Conditions *[]WafCustomCondition `json:"conditions,omitempty"`

	Action *WafCustomRuleAction `json:"action,omitempty"`

	// 域名
	DomainName *string `json:"domain_name,omitempty"`

	// 防护区域，0-大陆，1-海外
	OverseasType *int32 `json:"overseas_type,omitempty"`
}

func (o WafCustomRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WafCustomRule struct{}"
	}

	return strings.Join([]string{"WafCustomRule", string(data)}, " ")
}
