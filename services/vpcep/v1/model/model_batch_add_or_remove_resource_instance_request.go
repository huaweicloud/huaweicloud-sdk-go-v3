package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchAddOrRemoveResourceInstanceRequest struct {
	// 资源类型，值为： endpoint_service或 endpoint。

	ResourceType string `json:"resource_type"`
	// 资源ID，Endpoint Service ID或Endpoint ID。

	ResourceId string `json:"resource_id"`
	// 发送的实体的MIME类型。推荐用户默认使用application/json，如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。

	ContentType string `json:"Content-Type"`

	Body *BatchAddOrRemoveResourceInstanceBody `json:"body,omitempty"`
}

func (o BatchAddOrRemoveResourceInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchAddOrRemoveResourceInstanceRequest struct{}"
	}

	return strings.Join([]string{"BatchAddOrRemoveResourceInstanceRequest", string(data)}, " ")
}
