package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeEndpointServiceRequest Request Object
type UpgradeEndpointServiceRequest struct {

	// 发送的实体的MIME类型。推荐用户默认使用application/json， 如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。
	ContentType *string `json:"Content-Type,omitempty"`

	// 终端节点服务的ID。
	VpcEndpointServiceId string `json:"vpc_endpoint_service_id"`
}

func (o UpgradeEndpointServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeEndpointServiceRequest struct{}"
	}

	return strings.Join([]string{"UpgradeEndpointServiceRequest", string(data)}, " ")
}
