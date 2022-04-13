package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type EnlargeInstanceRequestBody struct {
	// 待扩容的对象类型。 - 扩容mongos节点时，取值为“mongos”。 - 扩容shard组时，取值为“shard”。

	Type EnlargeInstanceRequestBodyType `json:"type"`
	// 资源规格编码。

	SpecCode string `json:"spec_code"`
	// 一个集群实例下，最多支持16个mongos节点和16个shard组。

	Num string `json:"num"`

	Volume *AddShardingNodeVolumeOption `json:"volume,omitempty"`
}

func (o EnlargeInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnlargeInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"EnlargeInstanceRequestBody", string(data)}, " ")
}

type EnlargeInstanceRequestBodyType struct {
	value string
}

type EnlargeInstanceRequestBodyTypeEnum struct {
	MONGOS EnlargeInstanceRequestBodyType
	SHARD  EnlargeInstanceRequestBodyType
}

func GetEnlargeInstanceRequestBodyTypeEnum() EnlargeInstanceRequestBodyTypeEnum {
	return EnlargeInstanceRequestBodyTypeEnum{
		MONGOS: EnlargeInstanceRequestBodyType{
			value: "mongos",
		},
		SHARD: EnlargeInstanceRequestBodyType{
			value: "shard",
		},
	}
}

func (c EnlargeInstanceRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EnlargeInstanceRequestBodyType) UnmarshalJSON(b []byte) error {
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
