package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAntiVirusTaskRequest Request Object
type ListAntiVirusTaskRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 任务名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	TaskName *string `json:"task_name,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 不涉及
	Offset int32 `json:"offset"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit int32 `json:"limit"`

	// **参数解释**: 查询时间范围天数 **约束限制**: 与begin_time、end_time互斥，不可同时传参，优先按last_days筛选 **取值范围**: 最小值1，最大值90（支持查询近90天内任务） **默认取值**: 不涉及
	LastDays *int32 `json:"last_days,omitempty"`

	// **参数解释**： 自定义筛选任务的开始时间（任务启动时间≥该时间） **约束限制**： 与last_days互斥，需与end_time同时传参，格式需合法 **取值范围**： UTC时区，格式为YYYY-MM-DD HH:MM:SS **默认取值**： 不涉及
	BeginTime *string `json:"begin_time,omitempty"`

	// **参数解释**： 自定义筛选任务的结束时间（任务启动时间≤该时间） **约束限制**： 与last_days互斥，需与begin_time同时传参，且需大于等于begin_time **取值范围**： UTC时区，格式为YYYY-MM-DD HH:MM:SS **默认取值**： 不涉及
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**: 任务状态 **约束限制**: 不涉及 **取值范围**: 包含如下4种   - scanning：扫描中   - cancel：已取消   - fail：扫描失败   - finish：已完成 **默认取值**: 不涉及
	TaskStatus *string `json:"task_status,omitempty"`

	// **参数解释**: 服务器名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器私有IP **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 服务器弹性IP地址 **约束限制**: 不涉及 **取值范围**: IPv4格式（长度7-15位）、IPv6格式（长度15-39位） **默认取值**: 无
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**: 此次扫描任务是否付费 **约束限制**: 必选参数，仅支持指定布尔值 **取值范围**: true（付费任务）、false（免费任务） **默认取值**: 不涉及
	WhetherPaidTask bool `json:"whether_paid_task"`

	// 服务器扫描状态， **参数解释**： 服务器扫描状态 **约束限制**: 不涉及 **取值范围**: 包含如下4种   - scanning ：扫描中   - success ：扫描成功   - fail ：扫描失败   - cancel ：取消扫描 **默认取值**: 不涉及
	HostTaskStatus *[]string `json:"host_task_status,omitempty"`
}

func (o ListAntiVirusTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAntiVirusTaskRequest struct{}"
	}

	return strings.Join([]string{"ListAntiVirusTaskRequest", string(data)}, " ")
}
