package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 获取分布式实例时返回。
type ListHa struct {
	// 数据库一致性类型。取值为“strong”、“eventual”，分别表示强一致性、最终一致性。

	Consistency ListHaConsistency `json:"consistency"`
	// 备机同步参数。  取值：非空。  GaussDB(for openGauss)为 “sync” 说明： “sync”为同步模式。

	ReplicationMode string `json:"replication_mode"`
}

func (o ListHa) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListHa struct{}"
	}

	return strings.Join([]string{"ListHa", string(data)}, " ")
}

type ListHaConsistency struct {
	value string
}

type ListHaConsistencyEnum struct {
	STRONG   ListHaConsistency
	EVENTUAL ListHaConsistency
}

func GetListHaConsistencyEnum() ListHaConsistencyEnum {
	return ListHaConsistencyEnum{
		STRONG: ListHaConsistency{
			value: "strong",
		},
		EVENTUAL: ListHaConsistency{
			value: "eventual",
		},
	}
}

func (c ListHaConsistency) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListHaConsistency) UnmarshalJSON(b []byte) error {
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
