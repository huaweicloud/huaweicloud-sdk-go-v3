package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

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

	// 扩容包年包月实例的节点数量时可指定，表示是否自动从账户中支付，此字段不影响自动续订的支付方式。 - true，表示自动从账户中支付。 - false，表示手动从账户中支付，默认为该方式。
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
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

func (c EnlargeInstanceRequestBodyType) Value() string {
	return c.value
}

func (c EnlargeInstanceRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EnlargeInstanceRequestBodyType) UnmarshalJSON(b []byte) error {
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
