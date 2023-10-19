package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScalingConfigOption 创建伸缩配置请求
type CreateScalingConfigOption struct {

	// 伸缩配置名称(1-64个字符)，只能包含中文、字母、数字、下划线或中划线。
	ScalingConfigurationName *string `json:"scaling_configuration_name,omitempty"`

	InstanceConfig *InstanceConfig `json:"instance_config,omitempty"`

	// 源伸缩配置ID，通过该ID查询已有伸缩配置信息与instance_config中参数进行结合，创建新的伸缩配置。 说明：  - 若传入instance_config中的instance_id，则优先使用instance_id相关实例配置创建新的伸缩配置，source_scaling_configuration_id参数不生效。  - 若未传入instance_config中的instance_id，则使用source_scaling_configuration_id与instance_config中的参数相结合创建伸缩配置。         - 若instance_config中的参数值为null，则创建新伸缩配置时该字段不产生修改。         - 若instance_config中的参数值不为null，则创建新伸缩配置时该字段将覆盖原有值，其中值为空时，该字段会被清空。  - 若不指定source_scaling_configuration_id创建伸缩配置，则scaling_configuration_name和instance_config为必选。
	SourceScalingConfigurationId *string `json:"source_scaling_configuration_id,omitempty"`
}

func (o CreateScalingConfigOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingConfigOption struct{}"
	}

	return strings.Join([]string{"CreateScalingConfigOption", string(data)}, " ")
}
