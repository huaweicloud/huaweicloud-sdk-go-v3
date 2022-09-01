package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 大key记录结构体
type BigkeysBody struct {

	// key名称
	Name *string `json:"name,omitempty" xml:"name"`

	// key类型
	Type *BigkeysBodyType `json:"type,omitempty" xml:"type"`

	// 大key所在的分片，仅在实例类型为集群时支持,格式为ip:port
	Shard *string `json:"shard,omitempty" xml:"shard"`

	// 大key所在的db
	Db *int32 `json:"db,omitempty" xml:"db"`

	// key的value大小。
	Size *int32 `json:"size,omitempty" xml:"size"`

	// key大小的单位。type为string时，单位是：byte；type为list/set/zset/hash时，单位是：count
	Unit *string `json:"unit,omitempty" xml:"unit"`
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
