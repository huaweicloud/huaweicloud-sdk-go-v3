package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群扩容状态详情
type ResizeInfo struct {
	// 扩容状态，取值如下：  GROWING：扩容中  RESIZE_FAILURE：扩容失败

	ResizeStatus *string `json:"resize_status,omitempty"`
	// 扩容开始时间，格式为ISO8601：YYYY-MM-DDThh:mm:ss

	StartTime *string `json:"start_time,omitempty"`
	// 扩容后的节点数量

	TargetNodeNum *int32 `json:"target_node_num,omitempty"`
	// 扩容前的节点数量

	OriginNodeNum *int32 `json:"origin_node_num,omitempty"`
}

func (o ResizeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInfo struct{}"
	}

	return strings.Join([]string{"ResizeInfo", string(data)}, " ")
}
