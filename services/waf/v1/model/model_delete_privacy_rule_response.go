package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeletePrivacyRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 创建规则时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 规则状态（0：关闭，1：开启）
	Status *int32 `json:"status,omitempty"`

	// 路径
	Url *string `json:"url,omitempty"`

	// 屏蔽字段
	Category *string `json:"category,omitempty"`

	// 屏蔽字段名
	Index          *string `json:"index,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeletePrivacyRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivacyRuleResponse struct{}"
	}

	return strings.Join([]string{"DeletePrivacyRuleResponse", string(data)}, " ")
}
