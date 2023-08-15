package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BigkeysBody 大key记录结构体
type BigkeysBody struct {

	// key名称
	Name *string `json:"name,omitempty"`

	// key类型
	Type *BigkeysBodyType `json:"type,omitempty"`

	// 大key所在的分片，仅在实例类型为集群时支持,格式为ip:port
	Shard *string `json:"shard,omitempty"`

	// 大key所在的db
	Db *int32 `json:"db,omitempty"`

	// key的value大小。
	Size *int64 `json:"size,omitempty"`

	// key大小的单位。type为string时，单位是：byte；type为list/set/zset/hash时，单位是：count
	Unit *string `json:"unit,omitempty"`
}

func (o BigkeysBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BigkeysBody struct{}"
	}

	return strings.Join([]string{"BigkeysBody", string(data)}, " ")
}

type BigkeysBodyType struct {
	value string
}

type BigkeysBodyTypeEnum struct {
	STRING BigkeysBodyType
	LIST   BigkeysBodyType
	SET    BigkeysBodyType
	ZSET   BigkeysBodyType
	HASH   BigkeysBodyType
}

func GetBigkeysBodyTypeEnum() BigkeysBodyTypeEnum {
	return BigkeysBodyTypeEnum{
		STRING: BigkeysBodyType{
			value: "string",
		},
		LIST: BigkeysBodyType{
			value: "list",
		},
		SET: BigkeysBodyType{
			value: "set",
		},
		ZSET: BigkeysBodyType{
			value: "zset",
		},
		HASH: BigkeysBodyType{
			value: "hash",
		},
	}
}

func (c BigkeysBodyType) Value() string {
	return c.value
}

func (c BigkeysBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BigkeysBodyType) UnmarshalJSON(b []byte) error {
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
