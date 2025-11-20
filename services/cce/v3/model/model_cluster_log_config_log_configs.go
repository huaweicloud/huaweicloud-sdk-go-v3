package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ClusterLogConfigLogConfigs struct {

	// **参数解释**： 日志类型 **约束限制**： 必须为 **kube-apiserver**、**kube-controller-manager**、**kube-scheduler**、**audit** 或者系统插件名称 **取值范围**： - kube-apiserver：采集kube-apiserver组件日志 - kube-controller-manager：采集kube-controller-manager日志 - kube-scheduler：采集kube-scheduler日志 - audit：采集审计日志 - 系统插件名称：采集插件日志  **默认取值**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 是否采集 **约束限制**： 不涉及 **取值范围**： - true：开启日志采集 - false：关闭采集日志  **默认取值**： 不涉及
	Enable *bool `json:"enable,omitempty"`

	// **参数解释**： 组件类型 , 合法取值为control，audit，system-addon。 **约束限制**： - 仅组件类型为系统插件类型时返回该参数 - 作为**配置集群日志**接口更新参数时，如果要采集系统插件日志则必须声明为**system-addon**  **取值范围**： - control: 控制面组件日志。 - audit: 控制面审计日志。 - system-addon: 系统插件日志。  **默认取值**： 不涉及
	Type *ClusterLogConfigLogConfigsType `json:"type,omitempty"`
}

func (o ClusterLogConfigLogConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterLogConfigLogConfigs struct{}"
	}

	return strings.Join([]string{"ClusterLogConfigLogConfigs", string(data)}, " ")
}

type ClusterLogConfigLogConfigsType struct {
	value string
}

type ClusterLogConfigLogConfigsTypeEnum struct {
	CONTROL      ClusterLogConfigLogConfigsType
	AUDIT        ClusterLogConfigLogConfigsType
	SYSTEM_ADDON ClusterLogConfigLogConfigsType
}

func GetClusterLogConfigLogConfigsTypeEnum() ClusterLogConfigLogConfigsTypeEnum {
	return ClusterLogConfigLogConfigsTypeEnum{
		CONTROL: ClusterLogConfigLogConfigsType{
			value: "control",
		},
		AUDIT: ClusterLogConfigLogConfigsType{
			value: "audit",
		},
		SYSTEM_ADDON: ClusterLogConfigLogConfigsType{
			value: "system-addon",
		},
	}
}

func (c ClusterLogConfigLogConfigsType) Value() string {
	return c.value
}

func (c ClusterLogConfigLogConfigsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ClusterLogConfigLogConfigsType) UnmarshalJSON(b []byte) error {
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
