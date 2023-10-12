package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MicroServiceInfoCce CCE微服务详细信息（暂不支持）
type MicroServiceInfoCce struct {

	// 云容器引擎集群编号
	ClusterId string `json:"cluster_id"`

	// 命名空间
	Namespace string `json:"namespace"`

	// 工作负载类型  - deployment：无状态负载  - statefulset：有状态负载  - daemonset：守护进程集
	WorkloadType MicroServiceInfoCceWorkloadType `json:"workload_type"`

	// APP名称
	AppName *string `json:"app_name,omitempty"`

	// 标签名
	LabelKey *string `json:"label_key,omitempty"`

	// 标签值
	LabelValue *string `json:"label_value,omitempty"`

	// 云容器引擎集群名称
	ClusterName *string `json:"cluster_name,omitempty"`
}

func (o MicroServiceInfoCce) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MicroServiceInfoCce struct{}"
	}

	return strings.Join([]string{"MicroServiceInfoCce", string(data)}, " ")
}

type MicroServiceInfoCceWorkloadType struct {
	value string
}

type MicroServiceInfoCceWorkloadTypeEnum struct {
	DEPLOYMENT  MicroServiceInfoCceWorkloadType
	STATEFULSET MicroServiceInfoCceWorkloadType
	DAEMONSET   MicroServiceInfoCceWorkloadType
}

func GetMicroServiceInfoCceWorkloadTypeEnum() MicroServiceInfoCceWorkloadTypeEnum {
	return MicroServiceInfoCceWorkloadTypeEnum{
		DEPLOYMENT: MicroServiceInfoCceWorkloadType{
			value: "deployment",
		},
		STATEFULSET: MicroServiceInfoCceWorkloadType{
			value: "statefulset",
		},
		DAEMONSET: MicroServiceInfoCceWorkloadType{
			value: "daemonset",
		},
	}
}

func (c MicroServiceInfoCceWorkloadType) Value() string {
	return c.value
}

func (c MicroServiceInfoCceWorkloadType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MicroServiceInfoCceWorkloadType) UnmarshalJSON(b []byte) error {
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
