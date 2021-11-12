package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListFlavorsRequest struct {
	// 产品规格编码。

	SpecCode *string `json:"spec_code,omitempty"`
	// 缓存实例类型。取值范围如下： - single：表示单机实例 - ha：表示主备实例 - cluster：表示cluster集群实例 - proxy：表示Proxy集群实例 - ha_rw_split： 表示读写分离实例

	CacheMode *string `json:"cache_mode,omitempty"`
	// 缓存引擎类型。取值范围如下： - Redis - Memcached

	Engine *string `json:"engine,omitempty"`
	// 缓存版本，当缓存引擎为Redis时，取值范围如下： - 3.0 - 4.0 - 5.0

	EngineVersion *string `json:"engine_version,omitempty"`
	// CPU架构类型。取值范围如下： - x86_64：X86架构 - aarch64：ARM架构

	CpuType *ListFlavorsRequestCpuType `json:"cpu_type,omitempty"`
	// 缓存容量（G Byte）。 - Redis3.0：单机和主备类型实例取值：2、4、8、16、32、64。Proxy集群实例规格支持64、128、256、512和1024。 - Redis4.0和Redis5.0：单机和主备类型实例取值：0.125、0.25、0.5、1、2、4、8、16、32、64。Cluster集群实例规格支持24、32、48、64、96、128、192、256、384、512、768、1024。 - Memcached：单机和主备类型实例取值：2、4、8、16、32、64。

	Capacity *string `json:"capacity,omitempty"`
}

func (o ListFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ListFlavorsRequest", string(data)}, " ")
}

type ListFlavorsRequestCpuType struct {
	value string
}

type ListFlavorsRequestCpuTypeEnum struct {
	X86_64  ListFlavorsRequestCpuType
	AARCH64 ListFlavorsRequestCpuType
}

func GetListFlavorsRequestCpuTypeEnum() ListFlavorsRequestCpuTypeEnum {
	return ListFlavorsRequestCpuTypeEnum{
		X86_64: ListFlavorsRequestCpuType{
			value: "x86_64",
		},
		AARCH64: ListFlavorsRequestCpuType{
			value: "aarch64",
		},
	}
}

func (c ListFlavorsRequestCpuType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlavorsRequestCpuType) UnmarshalJSON(b []byte) error {
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
