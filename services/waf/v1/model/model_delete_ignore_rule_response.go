package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteIgnoreRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 创建规则的时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 规则状态，0：关闭，1：开启
	Status *int32 `json:"status,omitempty"`

	// 误报规则屏蔽路径，仅在mode为0的状态下有该字段
	Url *string `json:"url,omitempty"`

	// 屏蔽的内置规则id（内置规则id通常可以在Web应用防火墙控制台的防护策略->策略名称->Web基础防护->防护规则中查询，也可以再防护事件的事件详情中查询内置规则id）
	Rule *string `json:"rule,omitempty"`

	// 版本号，0代表v1旧版本，1代表v2新版本；mode为0时，不存在conditions字段，存在url和url_logic字段；mode为1时，不存在url和url_logic字段，存在conditions字段
	Mode *int32 `json:"mode,omitempty"`

	// url匹配逻辑
	UrlLogic *string `json:"url_logic,omitempty"`

	// 条件
	Conditions *[]Condition `json:"conditions,omitempty"`

	// 防护域名或防护网站
	Domains        *[]string `json:"domains,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o DeleteIgnoreRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIgnoreRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteIgnoreRuleResponse", string(data)}, " ")
}
