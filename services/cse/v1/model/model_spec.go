package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 微服务引擎专享版的CCE规格
type Spec struct {

	// 微服务引擎专享版CCE规格ID
	Id *int64 `json:"id,omitempty"`

	// 微服务引擎专享版ID
	EngineId *string `json:"engine_id,omitempty"`

	// 微服务引擎专享版CCE集群部署类型
	SpecType *SpecSpecType `json:"spec_type,omitempty"`

	// 微服务引擎专享版CCE集群信息，目前为null
	Cluster *string `json:"cluster,omitempty"`

	// 微服务引擎专享版CCE集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	ClusterNodes *SpecClusterNode `json:"cluster_nodes,omitempty"`

	// 微服务引擎专享版CCE集群规格
	Flavor *string `json:"flavor,omitempty"`

	// 微服务引擎专享版CCE集群所在region
	Region *string `json:"region,omitempty"`

	// 微服务引擎专享版CCE集群版本
	Version *string `json:"version,omitempty"`

	// 微服务引擎专享版CCE集群附加参数
	ExtendParam *string `json:"extend_param,omitempty"`
}

func (o Spec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Spec struct{}"
	}

	return strings.Join([]string{"Spec", string(data)}, " ")
}

type SpecSpecType struct {
	value string
}

type SpecSpecTypeEnum struct {
	CCE          SpecSpecType
	CSE          SpecSpecType
	SPRING_CLOUD SpecSpecType
}

func GetSpecSpecTypeEnum() SpecSpecTypeEnum {
	return SpecSpecTypeEnum{
		CCE: SpecSpecType{
			value: "CCE",
		},
		CSE: SpecSpecType{
			value: "CSE",
		},
		SPRING_CLOUD: SpecSpecType{
			value: "SpringCloud",
		},
	}
}

func (c SpecSpecType) Value() string {
	return c.value
}

func (c SpecSpecType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SpecSpecType) UnmarshalJSON(b []byte) error {
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
