package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowHistoryTaskDetailsRequest struct {
	// 当用户开启企业项目功能时，该参数生效，表示资源所属企业项目，不传表示默认项目。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 刷新任务ID。

	HistoryTasksId string `json:"history_tasks_id"`
	// 刷新预热的urls所显示单页最大数量，取值范围为1-10000。page_size和page_number必须同时传值。默认值30。

	PageSize *int32 `json:"page_size,omitempty"`
	// 刷新预热的urls当前查询为第几页，取值范围为1-65535。默认值1。

	PageNumber *int32 `json:"page_number,omitempty"`
	// url的状态 processing 处理中，succeed 完成，failed 失败，waiting 等待，refreshing 刷新中，preheating 预热中。

	Status *string `json:"status,omitempty"`
	// url的地址。

	Url *string `json:"url,omitempty"`
}

func (o ShowHistoryTaskDetailsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowHistoryTaskDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowHistoryTaskDetailsRequest", string(data)}, " ")
}
