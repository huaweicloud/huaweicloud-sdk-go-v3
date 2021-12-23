package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListQpsTimelineRequest struct {
	// 通过企业项目管理服务的查询企业项目列表接口ListEnterpriseProject查询通过企业项目管理服务的查询企业项目列表接口ListEnterpriseProject查询企业项目ID

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 起始时间（13位毫秒时间戳），需要和to同时使用

	From *int64 `json:"from,omitempty"`
	// 结束时间（13位毫秒时间戳），需要和from同时使用

	To *int64 `json:"to,omitempty"`
	// 域名id（通过ListHost接口查询）

	Hosts *string `json:"hosts,omitempty"`
	// 独享实例实例id（仅实例化模式涉及）

	Instances *string `json:"instances,omitempty"`
	// 展示维度，需要按天展示时传\"DAY\"

	GroupBy *string `json:"group_by,omitempty"`
}

func (o ListQpsTimelineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQpsTimelineRequest struct{}"
	}

	return strings.Join([]string{"ListQpsTimelineRequest", string(data)}, " ")
}
