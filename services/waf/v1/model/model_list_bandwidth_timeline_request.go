package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListBandwidthTimelineRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 起始时间（13位毫秒时间戳）

	From int64 `json:"from"`
	// 结束时间（13位毫秒时间戳）

	To int64 `json:"to"`
	// 要查询域名列表（通过ListHost接口查询）

	Hosts *string `json:"hosts,omitempty"`
	// 要查询实例列表（仅实例化模式涉及）

	Instances *string `json:"instances,omitempty"`
	// 展示维度，按天展示时传\"DAY\"

	GroupBy *string `json:"group_by,omitempty"`
}

func (o ListBandwidthTimelineRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListBandwidthTimelineRequest struct{}"
	}

	return strings.Join([]string{"ListBandwidthTimelineRequest", string(data)}, " ")
}
