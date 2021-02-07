package model

import (
	"encoding/json"

	"strings"
)

type AutoScalingPolicy struct {
	// 当前自动伸缩规则是否开启。
	AutoScalingEnable bool `json:"auto_scaling_enable"`
	// 指定该节点组的最小保留节点数。  取值范围：[0～500]
	MinCapacity int32 `json:"min_capacity"`
	// 指定该节点组的最大节点数。  取值范围：[0～500]
	MaxCapacity int32 `json:"max_capacity"`
	// 资源计划列表，详细说明请参见[表8](https://support.huaweicloud.com/api-mrs/mrs_02_0028.html#mrs_02_0028__table10281451162111)。若该参数为空表示不启用资源计划。  当启用弹性伸缩时，资源计划与自动伸缩规则需至少配置其中一种。  MRS 1.6.0及以后版本支持该参数。
	ResourcesPlans *[]ResourcesPlan `json:"resources_plans,omitempty"`
	// 自动伸缩的规则列表，详细说明请参见[表10](https://support.huaweicloud.com/api-mrs/mrs_02_0028.html#mrs_02_0028__t4c9e3e169631470d81d260543affb7e1https://support.huaweicloud.com/api-mrs/mrs_02_0028.html#mrs_02_0028__t4c9e3e169631470d81d260543affb7e1)。  当启用弹性伸缩时，资源计划与自动伸缩规则需至少配置其中一种。
	Rules *[]Rules `json:"rules,omitempty"`
	// 弹性伸缩自定义自动化脚本列表。详细说明请参见[表9](https://support.huaweicloud.com/api-mrs/mrs_02_0028.html#mrs_02_0028__table1921110172216)。若该参数为空表示不启用钩子脚本。  MRS 1.7.1及以后版本支持该参数。
	ExecScripts *[]ScaleScript `json:"exec_scripts,omitempty"`
}

func (o AutoScalingPolicy) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AutoScalingPolicy struct{}"
	}

	return strings.Join([]string{"AutoScalingPolicy", string(data)}, " ")
}
