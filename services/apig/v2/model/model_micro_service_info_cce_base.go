package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MicroServiceInfoCceBase CCE云容器引擎工作负载信息，service_type为CCE时必填。app_name或（label_key、label_value）至少填一个，只填app_name时，相当于（label_key=‘app’、label_value=app_name值）
type MicroServiceInfoCceBase struct {

	// 云容器引擎集群编号
	ClusterId string `json:"cluster_id"`

	// 命名空间
	Namespace string `json:"namespace"`

	// 工作负载类型  - deployment：无状态负载  - statefulset：有状态负载  - daemonset：守护进程集
	WorkloadType MicroServiceInfoCceBaseWorkloadType `json:"workload_type"`

	// APP名称。支持汉字，英文，数字，点，中划线，下划线，且只能以英文和汉字开头，1-64字符。 > 中文字符必须为UTF-8或者unicode编码。
	AppName *string `json:"app_name,omitempty"`

	// 服务标识名。支持汉字、英文、数字、中划线、下划线、点、斜杠、中英文格式下的小括号和冒号，且只能以英文、汉字和数字开头，1-64个字符。 > 中文字符必须为UTF-8或者unicode编码。
	LabelKey *string `json:"label_key,omitempty"`

	// 服务标识值。支持汉字，英文，数字，点，中划线，下划线，且只能以英文和汉字开头，1-64字符。 > 中文字符必须为UTF-8或者unicode编码。
	LabelValue *string `json:"label_value,omitempty"`
}

func (o MicroServiceInfoCceBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MicroServiceInfoCceBase struct{}"
	}

	return strings.Join([]string{"MicroServiceInfoCceBase", string(data)}, " ")
}

type MicroServiceInfoCceBaseWorkloadType struct {
	value string
}

type MicroServiceInfoCceBaseWorkloadTypeEnum struct {
	DEPLOYMENT  MicroServiceInfoCceBaseWorkloadType
	STATEFULSET MicroServiceInfoCceBaseWorkloadType
	DAEMONSET   MicroServiceInfoCceBaseWorkloadType
}

func GetMicroServiceInfoCceBaseWorkloadTypeEnum() MicroServiceInfoCceBaseWorkloadTypeEnum {
	return MicroServiceInfoCceBaseWorkloadTypeEnum{
		DEPLOYMENT: MicroServiceInfoCceBaseWorkloadType{
			value: "deployment",
		},
		STATEFULSET: MicroServiceInfoCceBaseWorkloadType{
			value: "statefulset",
		},
		DAEMONSET: MicroServiceInfoCceBaseWorkloadType{
			value: "daemonset",
		},
	}
}

func (c MicroServiceInfoCceBaseWorkloadType) Value() string {
	return c.value
}

func (c MicroServiceInfoCceBaseWorkloadType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MicroServiceInfoCceBaseWorkloadType) UnmarshalJSON(b []byte) error {
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
