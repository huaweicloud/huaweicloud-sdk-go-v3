package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Spec 微服务引擎的CCE规格
type Spec struct {

	// 微服务引擎CCE规格ID
	Id *int64 `json:"id,omitempty"`

	// 微服务引擎ID
	EngineId *string `json:"engineId,omitempty"`

	// 微服务引擎的集群部署类型
	SpecType *SpecSpecType `json:"specType,omitempty"`

	// 微服务引擎的CCE集群信息，目前为null
	Cluster *string `json:"cluster,omitempty"`

	// 微服务引擎的CCE集群ID
	ClusterId *string `json:"clusterId,omitempty"`

	ClusterNodes *SpecClusterNode `json:"clusterNodes,omitempty"`

	// 微服务引擎的CCE集群规格
	Flavor *string `json:"flavor,omitempty"`

	// 微服务引擎的CCE集群所在region
	Region *string `json:"region,omitempty"`

	// 微服务引擎的CCE集群版本
	Version *string `json:"version,omitempty"`

	// 微服务引擎的CCE集群附加参数
	ExtendParam *string `json:"extendParam,omitempty"`
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
	CSE2          SpecSpecType
	NACOS2        SpecSpecType
	MICRO_GATEWAY SpecSpecType
}

func GetSpecSpecTypeEnum() SpecSpecTypeEnum {
	return SpecSpecTypeEnum{
		CSE2: SpecSpecType{
			value: "CSE2",
		},
		NACOS2: SpecSpecType{
			value: "Nacos2",
		},
		MICRO_GATEWAY: SpecSpecType{
			value: "MicroGateway",
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
