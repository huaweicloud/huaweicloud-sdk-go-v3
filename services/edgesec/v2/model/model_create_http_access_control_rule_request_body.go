package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateHttpAccessControlRuleRequestBody struct {

	// 规则名称
	Name string `json:"name"`

	// 规则描述，最长512字符
	Description *string `json:"description,omitempty"`

	// 精准防护规则生效时间:  - “false”：表示该规则立即生效。   - “true”：表示自定义生效时间。
	Time bool `json:"time"`

	// 精准防护规则生效的起始时间戳（秒）。当time=true，才需要填写该参数。
	Start *int64 `json:"start,omitempty"`

	// 精准防护规则生效的终止时间戳（秒）。当time=true，才需要填写该参数。
	Terminal *int64 `json:"terminal,omitempty"`

	// 执行该规则的优先级，值越小，优先级越高，值相同时，规则创建时间早，优先级越高。取值范围：0到1000。
	Priority int32 `json:"priority"`

	// 命中条件
	Conditions []HttpAccessControlRuleCondition `json:"conditions"`

	Action *HttpRuleAction `json:"action"`
}

func (o CreateHttpAccessControlRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHttpAccessControlRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateHttpAccessControlRuleRequestBody", string(data)}, " ")
}
