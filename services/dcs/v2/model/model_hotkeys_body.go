package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 热key记录结构体
type HotkeysBody struct {

	// key名称
	Name *string `json:"name,omitempty" xml:"name"`

	// key类型
	Type *HotkeysBodyType `json:"type,omitempty" xml:"type"`

	// 热key所在的分片，仅在实例类型为集群时支持,格式为ip:port
	Shard *string `json:"shard,omitempty" xml:"shard"`

	// 热key所在的db
	Db *int32 `json:"db,omitempty" xml:"db"`

	// key的value大小。
	Size *int32 `json:"size,omitempty" xml:"size"`

	// key大小的单位。type为string时，单位是：byte；type为list/set/zset/hash时，单位是：count
	Unit *string `json:"unit,omitempty" xml:"unit"`

	// 表示某个key在一段时间的访问频度，会随着访问的频率而变化。  该值并不是简单的访问频率值，而是一个基于概率的对数计数器结果，最大为255(可表示100万次访问)，超过255后如果继续频繁访问该值并不会继续增大，同时默认如果每过一分钟没有访问，该值会衰减1。
	Freq *int32 `json:"freq,omitempty" xml:"freq"`
}

func (o HotkeysBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HotkeysBody struct{}"
	}

	return strings.Join([]string{"HotkeysBody", string(data)}, " ")
}

type HotkeysBodyType struct {
	value string
}

type HotkeysBodyTypeEnum struct {
	STRING HotkeysBodyType
	LIST   HotkeysBodyType
	SET    HotkeysBodyType
	ZSET   HotkeysBodyType
	HASH   HotkeysBodyType
}

func GetHotkeysBodyTypeEnum() HotkeysBodyTypeEnum {
	return HotkeysBodyTypeEnum{
		STRING: HotkeysBodyType{
			value: "string",
		},
		LIST: HotkeysBodyType{
			value: "list",
		},
		SET: HotkeysBodyType{
			value: "set",
		},
		ZSET: HotkeysBodyType{
			value: "zset",
		},
		HASH: HotkeysBodyType{
			value: "hash",
		},
	}
}

func (c HotkeysBodyType) Value() string {
	return c.value
}

func (c HotkeysBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HotkeysBodyType) UnmarshalJSON(b []byte) error {
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
