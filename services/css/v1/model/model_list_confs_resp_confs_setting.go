package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListConfsRespConfsSetting struct {

	// **参数解释**： 并行执行管道的Filters+Outputs阶段的工作线程数。 **取值范围**： 不涉及
	Workers *int32 `json:"workers,omitempty"`

	// **参数解释**： 单个工作线程在尝试执行其Filters和Outputs之前将从inputs收集的最大事件数，调大该值会增加内存开销。 **取值范围**： 1-100000
	BatchSize *int32 `json:"batchSize,omitempty"`

	// **参数解释**： 每个event被pipeline调度等待的最小时间。单位毫秒。 **取值范围**： 1-600000
	BatchDelayMs *int32 `json:"batchDelayMs,omitempty"`

	// **参数解释**： 用于事件缓冲的内部队列模型。 **取值范围**： - memory：基于内存的传统队列。 - persisted：基于磁盘的ACKed持久化队列。
	QueueType string `json:"queueType"`

	// **参数解释**： 如果使用持久化队列，则表示强制执行检查点之前写入的最大事件数。 **取值范围**： 1-100000
	QueueCheckPointWrites *int32 `json:"queueCheckPointWrites,omitempty"`

	// **参数解释**： 如果使用持久化队列，则表示持久化队列的总容量（以兆字节MB为单位），确保磁盘的容量大于该值。 **取值范围**： 1-10240
	QueueMaxBytesMb *int32 `json:"queueMaxBytesMb,omitempty"`
}

func (o ListConfsRespConfsSetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfsRespConfsSetting struct{}"
	}

	return strings.Join([]string{"ListConfsRespConfsSetting", string(data)}, " ")
}
