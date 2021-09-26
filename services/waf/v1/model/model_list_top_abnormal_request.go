package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListTopAbnormalRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 起始时间（13位时间戳）

	From int64 `json:"from"`
	// 结束时间（13位时间戳）

	To int64 `json:"to"`
	// 要查询的前几的结果

	Top *int32 `json:"top,omitempty"`
	// 状态码

	Code *int32 `json:"code,omitempty"`
	// 要查询域名列表（通过ListHost接口查询）

	Hosts *string `json:"hosts,omitempty"`
	// 要查询实例列表（仅独享模式涉及）

	Instances *string `json:"instances,omitempty"`
}

func (o ListTopAbnormalRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListTopAbnormalRequest struct{}"
	}

	return strings.Join([]string{"ListTopAbnormalRequest", string(data)}, " ")
}
