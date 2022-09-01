package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPodsRequest struct {

	// 指定节点查询应用实例，可选， 与group_id、deployment_id、deployment_ids四选一
	NodeId *string `json:"node_id,omitempty" xml:"node_id"`

	// 指定节点组查询应用实例，可选，与node_id、deployment_id、deployment_ids四选一
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// 指定应用部署ID查询应用实例，可选， 与node_id、group_id、deployment_ids四选一
	DeploymentId *string `json:"deployment_id,omitempty" xml:"deployment_id"`

	// 指定应用部署ID列表查询应用实例，多个ID使用逗号分隔，可选， 与node_id、group_id、deployment_id四选一
	DeploymentIds *string `json:"deployment_ids,omitempty" xml:"deployment_ids"`

	// 查询返回记录的数量限制
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 偏移量，表示查询该偏移量后面的记录
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty" xml:"ief-instance-id"`
}

func (o ListPodsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPodsRequest struct{}"
	}

	return strings.Join([]string{"ListPodsRequest", string(data)}, " ")
}
