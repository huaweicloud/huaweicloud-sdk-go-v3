package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChkFeatureInfo 检测特性规则
type ChkFeatureInfo struct {

	// 检测特性规则ID
	ChkFeatureId int32 `json:"chk_feature_id"`

	// 防护动作 1 检测  2 检测并阻断/拦截
	ProtectiveAction int32 `json:"protective_action"`

	// 开启状态 0 关闭 1 开启
	Enabled int32 `json:"enabled"`

	// 检测特性规则配置信息
	FeatureConfigure string `json:"feature_configure"`
}

func (o ChkFeatureInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChkFeatureInfo struct{}"
	}

	return strings.Join([]string{"ChkFeatureInfo", string(data)}, " ")
}
