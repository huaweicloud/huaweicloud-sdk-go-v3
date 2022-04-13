package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AcceptOrRejectEndpointRequest struct {
	// 发送的实体的MIME类型。推荐用户默认使用application/json，如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。

	ContentType string `json:"Content-Type"`
	// 终端节点服务的ID。

	VpcEndpointServiceId string `json:"vpc_endpoint_service_id"`

	Body *AcceptOrRejectEndpointRequestBody `json:"body,omitempty"`
}

func (o AcceptOrRejectEndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptOrRejectEndpointRequest struct{}"
	}

	return strings.Join([]string{"AcceptOrRejectEndpointRequest", string(data)}, " ")
}
