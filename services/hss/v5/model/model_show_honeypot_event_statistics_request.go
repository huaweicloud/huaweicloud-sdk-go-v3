package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHoneypotEventStatisticsRequest Request Object
type ShowHoneypotEventStatisticsRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 开始时间，13位时间戳
	BeginTime *int64 `json:"begin_time,omitempty"`

	// 结束时间，13位时间戳
	EndTime *int64 `json:"end_time,omitempty"`

	// 查询时间范围天数，与自定义查询时间begin_time，end_time互斥
	LastDays *int32 `json:"last_days,omitempty"`
}

func (o ShowHoneypotEventStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHoneypotEventStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ShowHoneypotEventStatisticsRequest", string(data)}, " ")
}
