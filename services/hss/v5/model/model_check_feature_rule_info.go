package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckFeatureRuleInfo struct {

	// 检测规则ID
	ChkFeatureId *int32 `json:"chk_feature_id,omitempty"`

	// 检测规则标识
	ChkFeatureName *string `json:"chk_feature_name,omitempty"`

	// 检测规则描述
	ChkFeatureDesc *string `json:"chk_feature_desc,omitempty"`

	// 检测特性规则配置信息
	FeatureConfigure *string `json:"feature_configure,omitempty"`

	// 防护动作，包含如下 -1 检测   -2 检测并阻断/拦截
	ProtectiveAction *int32 `json:"protective_action,omitempty"`

	// 可选防护动作，包含如下 -1 检测   -2 检测并阻断/拦截   -3 都可以
	OptionalProtectiveAction *int32 `json:"optional_protective_action,omitempty"`

	// 开启状态，包含如下 -0 开启 -1 关闭
	Enabled *int32 `json:"enabled,omitempty"`

	// 是否可编辑配置信息，包含如下 -0 否   -1 是
	Editable *int32 `json:"editable,omitempty"`
}

func (o CheckFeatureRuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckFeatureRuleInfo struct{}"
	}

	return strings.Join([]string{"CheckFeatureRuleInfo", string(data)}, " ")
}
