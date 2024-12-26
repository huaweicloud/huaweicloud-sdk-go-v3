package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DiscoverEventSchemaFromDataRequest Request Object
type DiscoverEventSchemaFromDataRequest struct {

	// {   \"description\": \"通过事件数据发现事件模型的请求\",   \"type\": \"object\",   \"required\": [     \"event\"   ],   \"properties\": {     \"event\": {       \"description\": \"事件数据内容\",       \"type\": \"string\",       \"maxLength\": 1024,       \"example\": \"{\\\"fileName\\\": \\\"one.jpg\\\", \\\"fileSize\\\": 1048576}\"     },     \"format\": {       \"description\": \"事件模型格式类型\",       \"type\": \"string\",       \"default\": \"JSON_SCHEMA_DRAFT_6\",       \"enum\": [         \"JSON_SCHEMA_DRAFT_6\"       ]     }   } }
	Body *interface{} `json:"body,omitempty"`
}

func (o DiscoverEventSchemaFromDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiscoverEventSchemaFromDataRequest struct{}"
	}

	return strings.Join([]string{"DiscoverEventSchemaFromDataRequest", string(data)}, " ")
}
