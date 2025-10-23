package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateReportSettingBody 更新报告配置请求参数
type UpdateReportSettingBody struct {

	// 配置名称，取值长度为2到32个字符
	SettingName string `json:"setting_name"`

	// 是否启用,取值范围：true,false
	Enabled bool `json:"enabled"`

	// 参与统计的资源归属的区域列表，不设置表示全部区域，默认不设置
	RegionIds *[]string `json:"region_ids,omitempty"`

	// 参与统计的资源类型列表，不设置表示全部资源类型，默认不设置
	ResourceTypes *[]ResourceTypeEnum `json:"resource_types,omitempty"`

	// 参与统计的数据范围（生成报告时刻往前多少天内的数据参与统计，滚动计算），默认7天，取值范围：7到30天，详情类数据报告不受此项配置约束
	TimeRange int32 `json:"time_range"`

	// 报告名称的前缀，取值长度为3到128个字符
	ReportNamePrefix string `json:"report_name_prefix"`
}

func (o UpdateReportSettingBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateReportSettingBody struct{}"
	}

	return strings.Join([]string{"UpdateReportSettingBody", string(data)}, " ")
}
