package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchCreateOrDeleteResourceTagsRequest struct {

	// 资源类型，目前有:  smn_topic，主题  smn_sms，短信  smn_application，移动推送
	ResourceType string `json:"resource_type" xml:"resource_type"`

	// 资源ID。  获取resource_id的方法：  当resource_type为“smn_topic”时， 手动添加请求消息头“X-SMN-RESOURCEID-TYPE=name”，资源ID即为topic名称。 不添加请求消息头，通过“查询资源实例”，获取资源ID。 当resource_type为“smn_sms”时，resource_id为签名ID。您可在控制台获取。
	ResourceId string `json:"resource_id" xml:"resource_id"`

	Body *BatchCreateOrDeleteResourceTagsRequestBody `json:"body,omitempty" xml:"body"`
}

func (o BatchCreateOrDeleteResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteResourceTagsRequest", string(data)}, " ")
}
