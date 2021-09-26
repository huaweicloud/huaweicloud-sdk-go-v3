package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListStatisticsRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 起始时间

	From int64 `json:"from"`
	// 结束时间

	To int64 `json:"to"`
	// 要查询域名列表

	Hosts *string `json:"hosts,omitempty"`
	// 要查询实例列表

	Instances *string `json:"instances,omitempty"`
}

func (o ListStatisticsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListStatisticsRequest", string(data)}, " ")
}
