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
type BatchRebootSeversOption struct {
	// 云服务器ID列表。
	Servers []ServerId `json:"servers"`
	// 重启类型：  - SOFT：普通重启。 - HARD：强制重启。
	Type BatchRebootSeversOptionType `json:"type"`
}

func (o BatchRebootSeversOption) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchRebootSeversOption", string(data)}, " ")
}

type BatchRebootSeversOptionType struct {
	value string
}

type BatchRebootSeversOptionTypeEnum struct {
	SOFT BatchRebootSeversOptionType
	HARD BatchRebootSeversOptionType
}

func GetBatchRebootSeversOptionTypeEnum() BatchRebootSeversOptionTypeEnum {
	return BatchRebootSeversOptionTypeEnum{
		SOFT: BatchRebootSeversOptionType{
			value: "SOFT",
		},
		HARD: BatchRebootSeversOptionType{
			value: "HARD",
		},
	}
}

func (c BatchRebootSeversOptionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *BatchRebootSeversOptionType) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
