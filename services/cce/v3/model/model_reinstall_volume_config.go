package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ReinstallVolumeConfig 节点重装场景服务器相关配置
type ReinstallVolumeConfig struct {

	// **参数解释**: Docker数据盘配置项(已废弃)。  默认配置示例如下： ``` \"lvmConfig\":\"dockerThinpool=vgpaas/90%VG;kubernetesLV=vgpaas/10%VG;diskType=evs;lvType=linear\" ```  包含如下字段：   - userLV：用户空间的大小，示例格式：vgpaas/20%VG   - userPath：用户空间挂载路径，示例格式：/home/wqt-test   - diskType：磁盘类型，目前只有evs、hdd和ssd三种格式   - lvType：逻辑卷的类型，目前支持linear和striped两种，示例格式：striped   - dockerThinpool：Docker盘的空间大小，示例格式：vgpaas/60%VG   - kubernetesLV：Kubelet空间大小，示例格式：vgpaas/20%VG  **约束限制**: 不涉及 **取值范围**: 不涉及 **默认取值**: 不涉及
	LvmConfig *string `json:"lvmConfig,omitempty"`

	Storage *Storage `json:"storage,omitempty"`

	// **参数解释**： 节点重置时磁盘数据的保留策略。 不传或该字段为空时，默认使用reset_managed_volumes策略清空由CCE管理的数据盘。 **约束限制**： 当保留自定义挂载卷时，挂载到指定目录与作为持久存储卷的高级配置不允许修改。 **取值范围**： - reset_managed_volumes：清空由CCE管理的数据盘。 - retain_custom_volumes：保留用户自定义挂载卷（包括挂载到指定目录的卷和用作本地持久卷的卷），集群版本需为v1.29.15-r90、v1.30.14-r90、v1.31.14-r50、v1.32.13-r20、v1.33.12-r0、v1.34.8-r0、v1.35.5-r0、v1.36.1-r10或以上版本。  **默认取值**： reset_managed_volumes
	VolumeResetPolicy *ReinstallVolumeConfigVolumeResetPolicy `json:"volumeResetPolicy,omitempty"`
}

func (o ReinstallVolumeConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallVolumeConfig struct{}"
	}

	return strings.Join([]string{"ReinstallVolumeConfig", string(data)}, " ")
}

type ReinstallVolumeConfigVolumeResetPolicy struct {
	value string
}

type ReinstallVolumeConfigVolumeResetPolicyEnum struct {
	RESET_MANAGED_VOLUMES ReinstallVolumeConfigVolumeResetPolicy
	RETAIN_CUSTOM_VOLUMES ReinstallVolumeConfigVolumeResetPolicy
}

func GetReinstallVolumeConfigVolumeResetPolicyEnum() ReinstallVolumeConfigVolumeResetPolicyEnum {
	return ReinstallVolumeConfigVolumeResetPolicyEnum{
		RESET_MANAGED_VOLUMES: ReinstallVolumeConfigVolumeResetPolicy{
			value: "reset_managed_volumes",
		},
		RETAIN_CUSTOM_VOLUMES: ReinstallVolumeConfigVolumeResetPolicy{
			value: "retain_custom_volumes",
		},
	}
}

func (c ReinstallVolumeConfigVolumeResetPolicy) Value() string {
	return c.value
}

func (c ReinstallVolumeConfigVolumeResetPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReinstallVolumeConfigVolumeResetPolicy) UnmarshalJSON(b []byte) error {
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
