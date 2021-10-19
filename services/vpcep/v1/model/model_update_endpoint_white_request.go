package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateEndpointWhiteRequest struct {
	// 终端节点的ID。

	VpcEndpointId string `json:"vpc_endpoint_id"`
	// 发送的实体的MIME类型。推荐用户默认使用application/json，如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。

	ContentType string `json:"Content-Type"`

	Body *UpdateEndpointWhiteRequestBody `json:"body,omitempty"`
}

func (o UpdateEndpointWhiteRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateEndpointWhiteRequest struct{}"
	}

	return strings.Join([]string{"UpdateEndpointWhiteRequest", string(data)}, " ")
}
