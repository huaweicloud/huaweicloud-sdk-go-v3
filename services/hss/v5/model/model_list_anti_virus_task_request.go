package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAntiVirusTaskRequest Request Object
type ListAntiVirusTaskRequest struct {

	// **参数解释**: 主机所属的企业项目ID。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。 **约束限制**: 开通企业项目功能后才需要配置企业项目。 **取值范围**: 字符长度1-256位 **默认取值**: 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 任务名称
	TaskName *string `json:"task_name,omitempty"`

	// 偏移量：指定返回记录的开始位置
	Offset int32 `json:"offset"`

	// 每页显示个数
	Limit int32 `json:"limit"`

	// 查询时间范围天数，与自定义查询时间begin_time，end_time互斥
	LastDays *int32 `json:"last_days,omitempty"`

	// 自定义查询时间，与查询时间范围天数互斥，查询时间段的起始时间，毫秒级时间戳，end_time减去begin_time小于等于2天，与查询时间范围天数互斥
	BeginTime *string `json:"begin_time,omitempty"`

	// 自定义时间，查询时间段的终止时间，毫秒级时间戳，end_time减去begin_time小于等于2天，与查询时间范围天数互斥
	EndTime *string `json:"end_time,omitempty"`

	// 任务状态，包含如下4种   - scanning ：扫描中   - cancel ：已取消   - fail ：扫描失败   - finish ：已完成
	TaskStatus *string `json:"task_status,omitempty"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 服务器私有IP
	PrivateIp *string `json:"private_ip,omitempty"`

	// 服务器公网IP
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
