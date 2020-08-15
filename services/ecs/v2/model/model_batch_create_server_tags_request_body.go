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

// This is a auto create Body Object
type BatchCreateServerTagsRequestBody struct {
	// 操作标识（仅支持小写）：create（创建）。
	Action BatchCreateServerTagsRequestBodyAction `json:"action"`
	// 标签列表。
	Tags []ServerTag `json:"tags"`
}

func (o BatchCreateServerTagsRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchCreateServerTagsRequestBody", string(data)}, " ")
}

type BatchCreateServerTagsRequestBodyAction struct {
	value string
}

type BatchCreateServerTagsRequestBodyActionEnum struct {
	CREATE BatchCreateServerTagsRequestBodyAction
}

func GetBatchCreateServerTagsRequestBodyActionEnum() BatchCreateServerTagsRequestBodyActionEnum {
	return BatchCreateServerTagsRequestBodyActionEnum{
		CREATE: BatchCreateServerTagsRequestBodyAction{
			value: "create",
		},
	}
}

func (c BatchCreateServerTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *BatchCreateServerTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
