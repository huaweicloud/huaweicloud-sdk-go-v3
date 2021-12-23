package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListEndpointsRequest struct {
	// 发送的实体的MIME类型。推荐用户默认使用application/json，如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。

	ContentType string `json:"Content-Type"`
	// 终端节点服务的名称，支持大小 写，前后模糊匹配。

	EndpointServiceName *string `json:"endpoint_service_name,omitempty"`
	// 终端节点所在的VPC的ID。

	VpcId *string `json:"vpc_id,omitempty"`
	// 终端节点的ID，唯一标识。

	Id *string `json:"id,omitempty"`
	// 查询返回终端节点的数量限制，即 每页返回的资源个数。 取值范围：0~1000，取值一般为 10，20或者50，默认为10。

	Limit *int32 `json:"limit,omitempty"`
	// 偏移量。 偏移量为一个大于0小于终端节点 服务总个数的整数，表示从偏移量 后面的终端节点服务开始查询。

	Offset *int32 `json:"offset,omitempty"`
	// 查询结果中终端节点列表的排序字 段，取值为： ● create_at：终端节点的创建时 间 ● update_at：终端节点的更新时 间 默认值为create_at。

	SortKey *string `json:"sort_key,omitempty"`
	// 查询结果中终端节点列表的排序方 式，取值为： ● desc：降序排序 ● asc：升序排序 默认值为desc。

	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListEndpointsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointsRequest struct{}"
	}

	return strings.Join([]string{"ListEndpointsRequest", string(data)}, " ")
}
