package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRequestTimelineRequest Request Object
type ListRequestTimelineRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 起始时间（13位毫秒时间戳），需要和to同时使用
	From int64 `json:"from"`

	// 结束时间（13位毫秒时间戳），需要和from同时使用
	To int64 `json:"to"`

	// 域名id，通过查询云模式防护域名列表（ListHost）获取域名id或者通过独享模式域名列表（ListPremiumHost）获取域名id。默认不传，查询该项目下所有防护域名的top业务异常统计信息。
	Hosts *[]string `json:"hosts,omitempty"`

	// 要查询引擎实例id
	Instances *[]string `json:"instances,omitempty"`

	// 展示维度，按天展示时传\"DAY\"；默认不传，按照分钟展示
	GroupBy *string `json:"group_by,omitempty"`
}

func (o ListRequestTimelineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRequestTimelineRequest struct{}"
	}

	return strings.Join([]string{"ListRequestTimelineRequest", string(data)}, " ")
}
