package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddPolicyGroupRequestInfo 复制的策略组信息
type AddPolicyGroupRequestInfo struct {

	// **参数解释**: 策略组的描述信息 **约束限制**: 不涉及 **取值范围**: 字符长度1-64，只能由中文字符、英文字母、数字、逗号、句号、空格及\"_\"、\"-\"组成 **默认取值**: 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**: 策略组名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-32，只能由中文字符、英文字母、数字、逗号、句号、空格及\"_\"、\"-\"组成，且不能以default_policy_group开头 **默认取值**: 不涉及
	Name string `json:"name"`

	// **参数解释**: 需要被复制的策略组的策略组ID，仅旗舰版和容器版策略组支持复制 **约束限制**: 需要使用 ListPolicyGroup 接口查询旗舰版和容器版策略组，在 ListPolicyGroup 接口的响应体中，support_version 等于 hss.version.container.enterprise 或 hss.version.premium 的 group_id 是可以被复制的策略组ID **取值范围**: 字符长度36-64 **默认取值**: 不涉及
	GroupId string `json:"group_id"`
}

func (o AddPolicyGroupRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddPolicyGroupRequestInfo struct{}"
	}

	return strings.Join([]string{"AddPolicyGroupRequestInfo", string(data)}, " ")
}
