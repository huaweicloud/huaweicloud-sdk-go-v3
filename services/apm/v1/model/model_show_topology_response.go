package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowTopologyResponse struct {

	// 全局traceID
	GlobalTraceId *string `json:"global_trace_id,omitempty" xml:"global_trace_id"`

	// 组件之间调用指向线列表
	LineList *[]TraceTopologyLine `json:"line_list,omitempty" xml:"line_list"`

	// 组件节点列表
	NodeList       *[]TraceTopologyNode `json:"node_list,omitempty" xml:"node_list"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowTopologyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopologyResponse struct{}"
	}

	return strings.Join([]string{"ShowTopologyResponse", string(data)}, " ")
}