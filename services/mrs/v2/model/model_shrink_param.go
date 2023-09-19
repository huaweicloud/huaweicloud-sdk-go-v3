package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShrinkParam 缩容参数
type ShrinkParam struct {

	// 节点组名称
	NodeGroupName string `json:"node_group_name"`

	// 缩容节点数量，如果是指定节点缩容，则该参数可以不填。
	Count *int32 `json:"count,omitempty"`

	// 缩容节点时指定待删除节点的资源ID列表。 resource_ids为空时，按照系统规则自动选择删除节点。 仅支持删除状态异常的ecs节点。 会针对指定节点进行强制删除。 可通过查询主机接口获取resource_id。
	ResourceIds *[]string `json:"resource_ids,omitempty"`

	// 启用该参数时，如果缩容失败或者节点为异常状态，会进行强制删除主机操作。如果resource_id不为空，该值默认为true。 如果resource_ids为空，则默认为false。
	ForceDelete *bool `json:"force_delete,omitempty"`

	// 支持跨AZ场景下的扩缩容，可以传入多个AZ，逗号分隔，如果为空或“”表示非跨AZ场景，不进行跨AZ扩缩容校验。
	AzCode *string `json:"az_code,omitempty"`
}

func (o ShrinkParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShrinkParam struct{}"
	}

	return strings.Join([]string{"ShrinkParam", string(data)}, " ")
}
