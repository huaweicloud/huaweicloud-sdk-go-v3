package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// VolumeConfig 节点池同步场景磁盘配置
type VolumeConfig struct {

	// **参数解释**： 节点重置时磁盘数据的保留策略。 不传或该字段为空时，默认使用reset_managed_volumes策略清空由CCE管理的数据盘。 **约束限制**： 当保留自定义挂载卷时，挂载到指定目录与作为持久存储卷的高级配置不允许修改。 **取值范围**： - reset_managed_volumes：清空由CCE管理的数据盘。 - retain_custom_volumes：保留用户自定义挂载卷（包括挂载到指定目录的卷和用作本地持久卷的卷），集群版本需为v1.29.15-r90、v1.30.14-r90、v1.31.14-r50、v1.32.13-r20、v1.33.12-r0、v1.34.8-r0、v1.35.5-r0、v1.36.1-r10或以上版本。  **默认取值**： reset_managed_volumes
	VolumeResetPolicy *VolumeConfigVolumeResetPolicy `json:"volumeResetPolicy,omitempty"`
}

func (o VolumeConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeConfig struct{}"
	}

	return strings.Join([]string{"VolumeConfig", string(data)}, " ")
}

type VolumeConfigVolumeResetPolicy struct {
	value string
}

type VolumeConfigVolumeResetPolicyEnum struct {
	RESET_MANAGED_VOLUMES VolumeConfigVolumeResetPolicy
	RETAIN_CUSTOM_VOLUMES VolumeConfigVolumeResetPolicy
}

func GetVolumeConfigVolumeResetPolicyEnum() VolumeConfigVolumeResetPolicyEnum {
	return VolumeConfigVolumeResetPolicyEnum{
		RESET_MANAGED_VOLUMES: VolumeConfigVolumeResetPolicy{
			value: "reset_managed_volumes",
		},
		RETAIN_CUSTOM_VOLUMES: VolumeConfigVolumeResetPolicy{
			value: "retain_custom_volumes",
		},
	}
}

func (c VolumeConfigVolumeResetPolicy) Value() string {
	return c.value
}

func (c VolumeConfigVolumeResetPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VolumeConfigVolumeResetPolicy) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
