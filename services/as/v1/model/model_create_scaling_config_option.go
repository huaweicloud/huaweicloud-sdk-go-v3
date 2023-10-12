package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScalingConfigOption 创建伸缩配置请求
type CreateScalingConfigOption struct {

	// 伸缩配置名称(1-64个字符)，只能包含中文、字母、数字、下划线或中划线。
	ScalingConfigurationName string `json:"scaling_configuration_name"`

	InstanceConfig *InstanceConfig `json:"instance_config"`

	// 源伸缩配置ID，通过ID获取原有伸缩配置信息进行修改，传入需修改的配置字段若为null值不产生修改，其他任何值（包括空值）均覆盖原有值。注意：若传入instance_id则优先使用instance_id获取到的值进行修改。
	SourceScalingConfigurationId *string `json:"source_scaling_configuration_id,omitempty"`
}

func (o CreateScalingConfigOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingConfigOption struct{}"
	}

	return strings.Join([]string{"CreateScalingConfigOption", string(data)}, " ")
}
