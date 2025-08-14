package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FeatureRuleInfo struct {

	// 检测特性规则ID
	ChkFeatureId *int32 `json:"chk_feature_id,omitempty"`

	// 检测特性规则标识
	ChkFeatureName *string `json:"chk_feature_name,omitempty"`

	// 检测特性规则描述
	ChkFeatureDesc *string `json:"chk_feature_desc,omitempty"`

	// 操作系统类型
	OsType *string `json:"os_type,omitempty"`

	// 检测特性规则配置信息
	FeatureConfigure *string `json:"feature_configure,omitempty"`

	// 可选防护动作，包含如下 -1 检测   -2 检测并阻断/拦截   -3 都可以
	OptionalProtectiveAction *int32 `json:"optional_protective_action,omitempty"`

	// 默认防护动作，包含如下 -1 检测   -2 检测并阻断/拦截
	ProtectiveAction *int32 `json:"protective_action,omitempty"`

	// 是否可编辑配置信息，包含如下 -0 否   -1 是
	Editable *int32 `json:"editable,omitempty"`
}

func (o FeatureRuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FeatureRuleInfo struct{}"
	}

	return strings.Join([]string{"FeatureRuleInfo", string(data)}, " ")
}
