package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTopologyResponse Response Object
type ShowTopologyResponse struct {

	// 全局traceId。
	GlobalTraceId *string `json:"global_trace_id,omitempty"`

	// 组件之间调用指向线列表。
	LineList *[]TraceTopologyLine `json:"line_list,omitempty"`

	// 组件节点列表。
	NodeList       *[]TraceTopologyNode `json:"node_list,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowTopologyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopologyResponse struct{}"
	}

	return strings.Join([]string{"ShowTopologyResponse", string(data)}, " ")
}
