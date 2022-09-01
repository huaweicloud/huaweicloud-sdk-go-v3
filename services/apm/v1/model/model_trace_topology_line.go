package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 拓扑图上的组件调用指向线
type TraceTopologyLine struct {

	// 开始节点id
	StartNodeId *int64 `json:"start_node_id,omitempty" xml:"start_node_id"`

	// 结束节点id
	EndNodeId *int64 `json:"end_node_id,omitempty" xml:"end_node_id"`

	// 调用跨度id
	SpanId *string `json:"span_id,omitempty" xml:"span_id"`

	ClientInfo *TraceTopologyLineInfo `json:"client_info,omitempty" xml:"client_info"`

	ServerInfo *TraceTopologyLineInfo `json:"server_info,omitempty" xml:"server_info"`

	// id
	Id *string `json:"id,omitempty" xml:"id"`

	// 获取一条线的提示信息
	Hint *string `json:"hint,omitempty" xml:"hint"`
}

func (o TraceTopologyLine) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TraceTopologyLine struct{}"
	}

	return strings.Join([]string{"TraceTopologyLine", string(data)}, " ")
}
