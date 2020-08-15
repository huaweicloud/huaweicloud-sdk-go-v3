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
type BatchDeleteServerTagsRequestBody struct {
	// 操作标识（仅支持小写）：delete（删除）。
	Action BatchDeleteServerTagsRequestBodyAction `json:"action"`
	// 标签列表。
	Tags []ServerTag `json:"tags"`
}

func (o BatchDeleteServerTagsRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchDeleteServerTagsRequestBody", string(data)}, " ")
}

type BatchDeleteServerTagsRequestBodyAction struct {
	value string
}

type BatchDeleteServerTagsRequestBodyActionEnum struct {
	DELETE BatchDeleteServerTagsRequestBodyAction
}

func GetBatchDeleteServerTagsRequestBodyActionEnum() BatchDeleteServerTagsRequestBodyActionEnum {
	return BatchDeleteServerTagsRequestBodyActionEnum{
		DELETE: BatchDeleteServerTagsRequestBodyAction{
			value: "delete",
		},
	}
}

func (c BatchDeleteServerTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *BatchDeleteServerTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
