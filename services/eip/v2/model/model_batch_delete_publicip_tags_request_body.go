/*
 * EIP
 *
 * 云服务接口
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// 批量操作资源标签的请求体
type BatchDeletePublicipTagsRequestBody struct {
	// 标签列表
	Tags []ResourceTagOption `json:"tags"`
	// 操作标识  delete：删除  action为delete时，value可选
	Action BatchDeletePublicipTagsRequestBodyAction `json:"action"`
}

func (o BatchDeletePublicipTagsRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchDeletePublicipTagsRequestBody", string(data)}, " ")
}

type BatchDeletePublicipTagsRequestBodyAction struct {
	value string
}

type BatchDeletePublicipTagsRequestBodyActionEnum struct {
	DELETE BatchDeletePublicipTagsRequestBodyAction
}

func GetBatchDeletePublicipTagsRequestBodyActionEnum() BatchDeletePublicipTagsRequestBodyActionEnum {
	return BatchDeletePublicipTagsRequestBodyActionEnum{
		DELETE: BatchDeletePublicipTagsRequestBodyAction{
			value: "delete",
		},
	}
}

func (c BatchDeletePublicipTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *BatchDeletePublicipTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
