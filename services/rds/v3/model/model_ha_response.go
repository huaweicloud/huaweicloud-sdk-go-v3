package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// HA配置参数，创建HA实例时使用。
type HaResponse struct {
	// 备机同步参数。实例主备模式为Ha时有效。 取值： - MySQL为“async”或“semisync”。 - PostgreSQL为“async”或“sync”。 - Microsoft SQL Server为“sync”。

	ReplicationMode HaResponseReplicationMode `json:"replication_mode"`
}

func (o HaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HaResponse struct{}"
	}

	return strings.Join([]string{"HaResponse", string(data)}, " ")
}

type HaResponseReplicationMode struct {
	value string
}

type HaResponseReplicationModeEnum struct {
	ASYNC    HaResponseReplicationMode
	SEMISYNC HaResponseReplicationMode
	SYNC     HaResponseReplicationMode
}

func GetHaResponseReplicationModeEnum() HaResponseReplicationModeEnum {
	return HaResponseReplicationModeEnum{
		ASYNC: HaResponseReplicationMode{
			value: "async",
		},
		SEMISYNC: HaResponseReplicationMode{
			value: "semisync",
		},
		SYNC: HaResponseReplicationMode{
			value: "sync",
		},
	}
}

func (c HaResponseReplicationMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HaResponseReplicationMode) UnmarshalJSON(b []byte) error {
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
