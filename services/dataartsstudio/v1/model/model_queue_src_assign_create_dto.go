package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueueSrcAssignCreateDto struct {

	// 队列资源服务(当前只支持mrs、dli)。
	SourceType *string `json:"source_type,omitempty"`

	// 队列名称。
	QueueName *[]string `json:"queue_name,omitempty"`

	// 数据连接id。
	ConnId *string `json:"conn_id,omitempty"`

	// 集群id。
	ClusterId *string `json:"cluster_id,omitempty"`

	// 当前空间分配资源附加的描述信息。
	Description *string `json:"description,omitempty"`
}

func (o QueueSrcAssignCreateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueueSrcAssignCreateDto struct{}"
	}

	return strings.Join([]string{"QueueSrcAssignCreateDto", string(data)}, " ")
}
