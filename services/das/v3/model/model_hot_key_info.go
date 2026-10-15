package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type HotKeyInfo struct {

	// key名称
	Name *string `json:"name,omitempty"`

	// key类型。取值：string、hash、list、zset、set、stream、exhash
	Type *HotKeyInfoType `json:"type,omitempty"`

	// 最新执行命令
	Command *string `json:"command,omitempty"`

	// QPS值
	Qps *int64 `json:"qps,omitempty"`

	// 入口流量（字节/秒）
	TrafficIn *int64 `json:"traffic_in,omitempty"`

	// 出口流量（字节/秒）
	TrafficOut *int64 `json:"traffic_out,omitempty"`

	// DB ID
	DbId *string `json:"db_id,omitempty"`
}

func (o HotKeyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HotKeyInfo struct{}"
	}

	return strings.Join([]string{"HotKeyInfo", string(data)}, " ")
}

type HotKeyInfoType struct {
	value string
}

type HotKeyInfoTypeEnum struct {
	STRING HotKeyInfoType
	HASH   HotKeyInfoType
	LIST   HotKeyInfoType
	ZSET   HotKeyInfoType
	SET    HotKeyInfoType
	STREAM HotKeyInfoType
	EXHASH HotKeyInfoType
}

func GetHotKeyInfoTypeEnum() HotKeyInfoTypeEnum {
	return HotKeyInfoTypeEnum{
		STRING: HotKeyInfoType{
			value: "string",
		},
		HASH: HotKeyInfoType{
			value: "hash",
		},
		LIST: HotKeyInfoType{
			value: "list",
		},
		ZSET: HotKeyInfoType{
			value: "zset",
		},
		SET: HotKeyInfoType{
			value: "set",
		},
		STREAM: HotKeyInfoType{
			value: "stream",
		},
		EXHASH: HotKeyInfoType{
			value: "exhash",
		},
	}
}

func (c HotKeyInfoType) Value() string {
	return c.value
}

func (c HotKeyInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HotKeyInfoType) UnmarshalJSON(b []byte) error {
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
