package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBandwidthTimelineRequest Request Object
type ListBandwidthTimelineRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 查询的带宽统计数据的起始时间（13位毫秒时间戳），需要和to同时使用
	From int64 `json:"from"`

	// 查询的带宽统计数据的结束时间（13位毫秒时间戳），需要和from同时使用
	To int64 `json:"to"`

	// 域名id，用于查询指定的防护域名在from到to这段时间内的带宽数据。通过查询云模式防护域名列表（ListHost）获取域名id或者通过独享模式域名列表（ListPremiumHost）获取域名id
	Hosts *string `json:"hosts,omitempty"`

	// 引擎实例id，用于查询指定的独享引擎实例所防护的域名在from到to这段时间内的带宽数据。
	Instances *string `json:"instances,omitempty"`

	// 展示维度，按天展示时传\"DAY\"；默认不传，按照分钟展示。
	GroupBy *string `json:"group_by,omitempty"`

	// 发送/接受字节数，查看峰值请输入1，查看平均值请输入0
	DisplayOption *int32 `json:"display_option,omitempty"`
}

func (o ListBandwidthTimelineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthTimelineRequest struct{}"
	}

	return strings.Join([]string{"ListBandwidthTimelineRequest", string(data)}, " ")
}
