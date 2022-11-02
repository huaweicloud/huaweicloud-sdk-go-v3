package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CCE云容器引擎详细信息，service_type为CCE时必填
type MicroServiceInfoCceBase struct {

	// 云容器引擎集群编号
	ClusterId string `json:"cluster_id"`

	// 命名空间
	Namespace string `json:"namespace"`

	// 工作负载类型  - deployment：无状态负载  - statefulset：有状态负载  - daemonset：守护进程集
	WorkloadType MicroServiceInfoCceBaseWorkloadType `json:"workload_type"`

	// APP名称
	AppName string `json:"app_name"`
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
