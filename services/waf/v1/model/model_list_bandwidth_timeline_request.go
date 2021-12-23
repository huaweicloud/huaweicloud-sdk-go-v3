package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListBandwidthTimelineRequest struct {
	// 您可以通过调用企业项目管理服务（EPS)的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 起始时间（13位毫秒时间戳），需要和to同时使用

	From *int64 `json:"from,omitempty"`
	// 结束时间（13位毫秒时间戳），需要和from同时使用

	To *int64 `json:"to,omitempty"`
	// 要查询域名列表（通过ListHost接口查询）

	Hosts *string `json:"hosts,omitempty"`
	// 要查询实例列表（仅实例化模式涉及）

	Instances *string `json:"instances,omitempty"`
	// 展示维度，按天展示时传\"DAY\"

	GroupBy *string `json:"group_by,omitempty"`
}

func (o ListBandwidthTimelineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthTimelineRequest struct{}"
	}

	return strings.Join([]string{"ListBandwidthTimelineRequest", string(data)}, " ")
}
