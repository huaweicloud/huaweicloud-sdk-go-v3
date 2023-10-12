package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MicroServiceInfoCceBase CCE云容器引擎详细信息，service_type为CCE时必填。app_name或（label_key、label_value）至少填一个，只填app_name时，相当于（label_key=‘app’、label_value=app_name值）
type MicroServiceInfoCceBase struct {

	// 云容器引擎集群编号
	ClusterId string `json:"cluster_id"`

	// 命名空间
	Namespace string `json:"namespace"`

	// 工作负载类型  - deployment：无状态负载  - statefulset：有状态负载  - daemonset：守护进程集
	WorkloadType MicroServiceInfoCceBaseWorkloadType `json:"workload_type"`

	// APP名称
	AppName *string `json:"app_name,omitempty"`

	// 标签名
	LabelKey *string `json:"label_key,omitempty"`

	// 标签值
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
