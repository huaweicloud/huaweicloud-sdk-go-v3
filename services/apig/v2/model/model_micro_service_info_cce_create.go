package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type MicroServiceInfoCceCreate struct {

	// 云容器引擎集群编号
	ClusterId string `json:"cluster_id"`

	// 命名空间
	Namespace string `json:"namespace"`

	// 工作负载类型  - deployment：无状态负载  - statefulset：有状态负载  - daemonset：守护进程集
	WorkloadType MicroServiceInfoCceCreateWorkloadType `json:"workload_type"`

	// APP名称
	AppName string `json:"app_name"`

	// 工作负载的版本
	Version string `json:"version"`

	// 工作负载的监听端口号
	Port int32 `json:"port"`
}

func (o MicroServiceInfoCceCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MicroServiceInfoCceCreate struct{}"
	}

	return strings.Join([]string{"MicroServiceInfoCceCreate", string(data)}, " ")
}

type MicroServiceInfoCceCreateWorkloadType struct {
	value string
}

type MicroServiceInfoCceCreateWorkloadTypeEnum struct {
	DEPLOYMENT  MicroServiceInfoCceCreateWorkloadType
	STATEFULSET MicroServiceInfoCceCreateWorkloadType
	DAEMONSET   MicroServiceInfoCceCreateWorkloadType
}

func GetMicroServiceInfoCceCreateWorkloadTypeEnum() MicroServiceInfoCceCreateWorkloadTypeEnum {
	return MicroServiceInfoCceCreateWorkloadTypeEnum{
		DEPLOYMENT: MicroServiceInfoCceCreateWorkloadType{
			value: "deployment",
		},
		STATEFULSET: MicroServiceInfoCceCreateWorkloadType{
			value: "statefulset",
		},
		DAEMONSET: MicroServiceInfoCceCreateWorkloadType{
			value: "daemonset",
		},
	}
}

func (c MicroServiceInfoCceCreateWorkloadType) Value() string {
	return c.value
}

func (c MicroServiceInfoCceCreateWorkloadType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MicroServiceInfoCceCreateWorkloadType) UnmarshalJSON(b []byte) error {
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
