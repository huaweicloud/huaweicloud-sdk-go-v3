package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServiceDescribeDetailsRequest Request Object
type ListServiceDescribeDetailsRequest struct {

	// 发送的实体的MIME类型。推荐用户默认使用application/json， 如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。
	ContentType *string `json:"Content-Type,omitempty"`

	// 终端节点服务的名称。说明：该字段和id字段必须二选一，否则会出现错误。
	EndpointServiceName *string `json:"endpoint_service_name,omitempty"`

	// 终端节点服务的ID，唯一标识。 说明：该字段必须和endpoint_service_name字段二选一，否则会出现错误。
	Id *string `json:"id,omitempty"`
}

func (o ListServiceDescribeDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceDescribeDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListServiceDescribeDetailsRequest", string(data)}, " ")
}
