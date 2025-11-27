package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAntiVirusTaskRequest Request Object
type ListAntiVirusTaskRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 任务名称
	TaskName *string `json:"task_name,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 不涉及
	Offset int32 `json:"offset"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit int32 `json:"limit"`

	// 查询时间范围天数，与自定义查询时间begin_time，end_time互斥
	LastDays *int32 `json:"last_days,omitempty"`

	// 自定义查询时间，开始时间
	BeginTime *string `json:"begin_time,omitempty"`

	// 自定义查询时间，结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 任务状态，包含如下4种   - scanning ：扫描中   - cancel ：已取消   - fail ：扫描失败   - finish ：已完成
	TaskStatus *string `json:"task_status,omitempty"`

	// **参数解释**: 服务器名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器私有IP **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 服务器弹性IP地址。 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 无
	PublicIp *string `json:"public_ip,omitempty"`

	// 此次扫描任务是否付费
	WhetherPaidTask bool `json:"whether_paid_task"`

	// 服务器扫描状态，包含如下4种   - scanning ：扫描中   - success ：扫描成功   - fail ：扫描失败   - cancel ：取消扫描
	HostTaskStatus *[]string `json:"host_task_status,omitempty"`
}

func (o ListAntiVirusTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAntiVirusTaskRequest struct{}"
	}

	return strings.Join([]string{"ListAntiVirusTaskRequest", string(data)}, " ")
}
