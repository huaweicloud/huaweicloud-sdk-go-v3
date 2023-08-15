package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateJobReq 创建单个任务请求体。
type CreateJobReq struct {
	BaseInfo *JobBaseInfo `json:"base_info"`

	// 创建任务数据库信息体。
	SourceEndpoint []JobEndpointInfo `json:"source_endpoint"`

	// 创建任务数据库信息体。
	TargetEndpoint []JobEndpointInfo `json:"target_endpoint"`

	PeriodOrder *PeriodOrderInfo `json:"period_order,omitempty"`

	NodeInfo *JobNodeInfo `json:"node_info"`
}

func (o CreateJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateJobReq struct{}"
	}

	return strings.Join([]string{"CreateJobReq", string(data)}, " ")
}
