/*
 * ecs
 *
 * ECS Open API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

//
type BatchStopServersOption struct {
	// 标记为启动云服务器操作。
	Servers []ServerId `json:"servers"`
	// 关机类型，默认为SOFT：  - SOFT：普通关机（默认）。 - HARD：强制关机。
	Type BatchStopServersOptionType `json:"type,omitempty"`
}

func (o BatchStopServersOption) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchStopServersOption", string(data)}, " ")
}

type BatchStopServersOptionType struct {
	value string
}

type BatchStopServersOptionTypeEnum struct {
	SOFT BatchStopServersOptionType
	HARD BatchStopServersOptionType
}

func GetBatchStopServersOptionTypeEnum() BatchStopServersOptionTypeEnum {
	return BatchStopServersOptionTypeEnum{
		SOFT: BatchStopServersOptionType{
			value: "SOFT",
		},
		HARD: BatchStopServersOptionType{
			value: "HARD",
		},
	}
}

func (c BatchStopServersOptionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *BatchStopServersOptionType) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
