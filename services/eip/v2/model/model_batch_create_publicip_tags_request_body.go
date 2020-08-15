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
type BatchCreatePublicipTagsRequestBody struct {
	// 标签列表
	Tags []ResourceTagOption `json:"tags"`
	// 操作标识  create：创建  action为create时，tag的value必选
	Action BatchCreatePublicipTagsRequestBodyAction `json:"action"`
}

func (o BatchCreatePublicipTagsRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchCreatePublicipTagsRequestBody", string(data)}, " ")
}

type BatchCreatePublicipTagsRequestBodyAction struct {
	value string
}

type BatchCreatePublicipTagsRequestBodyActionEnum struct {
	CREATE BatchCreatePublicipTagsRequestBodyAction
}

func GetBatchCreatePublicipTagsRequestBodyActionEnum() BatchCreatePublicipTagsRequestBodyActionEnum {
	return BatchCreatePublicipTagsRequestBodyActionEnum{
		CREATE: BatchCreatePublicipTagsRequestBodyAction{
			value: "create",
		},
	}
}

func (c BatchCreatePublicipTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *BatchCreatePublicipTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
