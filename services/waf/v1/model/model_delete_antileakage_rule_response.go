package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteAntileakageRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 规则应用的url
	Url *string `json:"url,omitempty"`

	// 类别（响应码：code，敏感信息：sensitive）
	Category *string `json:"category,omitempty"`

	// 内容
	Contents *[]string `json:"contents,omitempty"`

	// 创建规则时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 规则状态，0：关闭，1：开启
	Status *int32 `json:"status,omitempty"`

	// 规则描述
	Description    *string `json:"description,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAntileakageRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAntileakageRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteAntileakageRuleResponse", string(data)}, " ")
}
