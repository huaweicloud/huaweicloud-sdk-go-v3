package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// databases参数说明
type CreateDatabaseDetail struct {
	// 逻辑库名称，需要满足以下条件：  - 长度为2-24个字符。 - 必须以字母开头，且不区分大小写。 - 可以包含字母、数字、下划线，不能包含其它特殊字符。 - 禁用关键字：  \"information_schema\"、\"mysql\"、\"performance_schema\"、\"sys\"。

	Name string `json:"name"`
	// 逻辑库的拆分模式。 - cluster表示逻辑库是拆分模式。 - single表示逻辑库是非拆分模式。

	ShardMode CreateDatabaseDetailShardMode `json:"shard_mode"`
	// 同一种工作模式下逻辑库分片的数量，shard_unit与关联rds数量的乘积。

	ShardNumber int32 `json:"shard_number"`
	// 单个RDS上的逻辑库分片数。  - 非拆分逻辑库，固定为1。 - 拆分逻辑库缺省为8，可以根据需要配置为8、16。

	ShardUnit int32 `json:"shard_unit"`
	// 逻辑库关联的RDS。

	UsedRds []DatabaseInstabcesParam `json:"used_rds"`
}

func (o CreateDatabaseDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseDetail struct{}"
	}

	return strings.Join([]string{"CreateDatabaseDetail", string(data)}, " ")
}

type CreateDatabaseDetailShardMode struct {
	value string
}

type CreateDatabaseDetailShardModeEnum struct {
	CLUSTER CreateDatabaseDetailShardMode
	SINGLE  CreateDatabaseDetailShardMode
}

func GetCreateDatabaseDetailShardModeEnum() CreateDatabaseDetailShardModeEnum {
	return CreateDatabaseDetailShardModeEnum{
		CLUSTER: CreateDatabaseDetailShardMode{
			value: "cluster",
		},
		SINGLE: CreateDatabaseDetailShardMode{
			value: "single",
		},
	}
}

func (c CreateDatabaseDetailShardMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDatabaseDetailShardMode) UnmarshalJSON(b []byte) error {
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
