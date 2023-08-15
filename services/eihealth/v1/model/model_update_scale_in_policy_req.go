package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateScaleInPolicyReq struct {

	// 空置时间
	IdleTime *int32 `json:"idle_time,omitempty"`

	// 缩容阈值
	Threshold *int32 `json:"threshold,omitempty"`

	// 扩容后多久再次判断缩容
	DelayAfterAdd *int32 `json:"delay_after_add,omitempty"`

	// 节点删除后多久再次判断缩容
	DelayAfterDelete *int32 `json:"delay_after_delete,omitempty"`

	// 缩容失败后多久再次判断缩容
	DelayAfterFailure *int32 `json:"delay_after_failure,omitempty"`

	// 缩容并发数
	MaxNodesBatchDeletion *int32 `json:"max_nodes_batch_deletion,omitempty"`

	// 检查间隔
	CheckInterval *int32 `json:"check_interval,omitempty"`
}

func (o UpdateScaleInPolicyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScaleInPolicyReq struct{}"
	}

	return strings.Join([]string{"UpdateScaleInPolicyReq", string(data)}, " ")
}
