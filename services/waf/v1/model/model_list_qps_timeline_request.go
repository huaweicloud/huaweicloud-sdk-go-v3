package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListQpsTimelineRequest struct {

	// 通过企业项目管理服务的查询企业项目列表接口ListEnterpriseProject查询通过企业项目管理服务的查询企业项目列表接口ListEnterpriseProject查询企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 起始时间（13位毫秒时间戳），需要和to同时使用
	From int64 `json:"from" xml:"from"`

	// 结束时间（13位毫秒时间戳），需要和from同时使用
	To int64 `json:"to" xml:"to"`

	// 域名id，通过查询云模式防护域名列表（ListHost）获取域名id或者通过独享模式域名列表（ListPremiumHost）获取域名id
	Hosts *string `json:"hosts,omitempty" xml:"hosts"`

	// 要查询引擎实例id（仅独享或者ELB实例化模式涉及）
	Instances *string `json:"instances,omitempty" xml:"instances"`

	// 展示维度，按天展示时传\"DAY\"；默认不传，按照分钟展示
	GroupBy *string `json:"group_by,omitempty" xml:"group_by"`
}

func (o ListQpsTimelineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQpsTimelineRequest struct{}"
	}

	return strings.Join([]string{"ListQpsTimelineRequest", string(data)}, " ")
}
