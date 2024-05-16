package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueueSrcAssignEntity 创建的队列资源。
type QueueSrcAssignEntity struct {

	// 队列资源id。
	Id *string `json:"id,omitempty"`

	// 队列资源服务名称。
	SourceType *string `json:"source_type,omitempty"`

	// 队列名称。
	QueueName *string `json:"queue_name,omitempty"`

	// 队列类型。
	QueueType *string `json:"queue_type,omitempty"`

	// 队列属性(0:默认,1:实时队列,2:离线队列), 当前只有yarn队列涉及。
	QueueAttr *int32 `json:"queue_attr,omitempty"`

	// 数据连接id。
	ConnId *string `json:"conn_id,omitempty"`

	// 数据连接名称。
	ConnName *string `json:"conn_name,omitempty"`

	// 集群id。
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称。
	ClusterName *string `json:"cluster_name,omitempty"`

	// 队列加入此空间的时间。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 队列加入此的操作人。
	CreateUser *string `json:"create_user,omitempty"`

	// 当前空间下管理的队列更新时间。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 当前空间下管理的队列更新人。
	UpdateUser *string `json:"update_user,omitempty"`

	// 项目id。
	ProjectId *string `json:"project_id,omitempty"`

	// 当前空间分配资源附加的描述信息。
	Description *string `json:"description,omitempty"`
}

func (o QueueSrcAssignEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueueSrcAssignEntity struct{}"
	}

	return strings.Join([]string{"QueueSrcAssignEntity", string(data)}, " ")
}
