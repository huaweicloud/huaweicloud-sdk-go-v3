package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddEndpointServicePermissionsRequest Request Object
type BatchAddEndpointServicePermissionsRequest struct {

	// 发送的实体的MIME类型。推荐用户默认使用application/json， 如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。
	ContentType *string `json:"Content-Type,omitempty"`

	// 终端节点服务的ID。
	VpcEndpointServiceId string `json:"vpc_endpoint_service_id"`

	Body *BatchAddPermissionRequest `json:"body,omitempty"`
}

func (o BatchAddEndpointServicePermissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddEndpointServicePermissionsRequest struct{}"
	}

	return strings.Join([]string{"BatchAddEndpointServicePermissionsRequest", string(data)}, " ")
}
