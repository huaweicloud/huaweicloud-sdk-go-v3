package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulScanTaskRequest Request Object
type ListVulScanTaskRequest struct {

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 每页显示个数
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量：指定返回记录的开始位置
	Offset *int32 `json:"offset,omitempty"`

	// 扫描任务的类型，包含如下：   -manual : 手动扫描任务   -schedule : 定时扫描任务
	ScanType *string `json:"scan_type,omitempty"`

	// 扫描任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 扫描任务开始时间的最小值
	MinStartTime *int64 `json:"min_start_time,omitempty"`

	// 扫描任务开始时间的最大值
	MaxStartTime *int64 `json:"max_start_time,omitempty"`
}

func (o ListVulScanTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulScanTaskRequest struct{}"
	}

	return strings.Join([]string{"ListVulScanTaskRequest", string(data)}, " ")
}
