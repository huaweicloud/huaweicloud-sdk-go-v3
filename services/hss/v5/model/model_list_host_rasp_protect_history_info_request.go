package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostRaspProtectHistoryInfoRequest Request Object
type ListHostRaspProtectHistoryInfoRequest struct {

	// Region Id
	Region string `json:"region"`

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// Host Id，为空时查所有主机
	HostId *string `json:"host_id,omitempty"`

	// 起始时间(ms)
	StartTime int64 `json:"start_time"`

	// 终止时间(ms)
	EndTime int64 `json:"end_time"`

	// limit
	Limit int32 `json:"limit"`

	// 偏移量：指定返回记录的开始位置
	Offset int32 `json:"offset"`

	// 告警级别 - 1 : 低危 - 2 : 中危 - 3 : 高危 - 4 : 严重
	AlarmLevel *int32 `json:"alarm_level,omitempty"`

	// 威胁等级   - Security : 安全   - Low : 低危   - Medium : 中危   - High : 高危   - Critical : 危急
	Severity *string `json:"severity,omitempty"`

	// 防护状态   - closed : 未开启   - opened : 防护中
	ProtectStatus *string `json:"protect_status,omitempty"`
}

func (o ListHostRaspProtectHistoryInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostRaspProtectHistoryInfoRequest struct{}"
	}

	return strings.Join([]string{"ListHostRaspProtectHistoryInfoRequest", string(data)}, " ")
}
