package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 获取分布式/主备版实例时返回。
type ListHaResult struct {

	// 数据库一致性类型，分布式模式实例仅有。取值为“strong”、“eventual”，分别表示强一致性、最终一致性。
	Consistency ListHaResultConsistency `json:"consistency"`

	// 备机同步参数。  取值：非空。  GaussDB为 “sync” 说明： “sync”为同步模式。
	ReplicationMode string `json:"replication_mode"`
}

func (o ListHaResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHaResult struct{}"
	}

	return strings.Join([]string{"ListHaResult", string(data)}, " ")
}

type ListHaResultConsistency struct {
	value string
}

type ListHaResultConsistencyEnum struct {
	STRONG   ListHaResultConsistency
	EVENTUAL ListHaResultConsistency
}

func GetListHaResultConsistencyEnum() ListHaResultConsistencyEnum {
	return ListHaResultConsistencyEnum{
		STRONG: ListHaResultConsistency{
			value: "strong",
		},
		EVENTUAL: ListHaResultConsistency{
			value: "eventual",
		},
	}
}

func (c ListHaResultConsistency) Value() string {
	return c.value
}

func (c ListHaResultConsistency) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHaResultConsistency) UnmarshalJSON(b []byte) error {
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
