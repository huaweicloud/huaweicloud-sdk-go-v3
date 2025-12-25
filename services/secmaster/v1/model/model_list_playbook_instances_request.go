package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPlaybookInstancesRequest Request Object
type ListPlaybookInstancesRequest struct {

	// **参数解释：** 内容类型 - application/json    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json  **默认取值：** application/json
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 剧本实例状态. (RUNNING--运行中、FINISHED--成功、FAILED--失败、RETRYING--重试中、TERMINATING--终止中、TERMINATED--已终止)
	Status *string `json:"status,omitempty"`

	// 实例名称
	Name *string `json:"name,omitempty"`

	// 剧本名称
	PlaybookName *string `json:"playbook_name,omitempty"`

	// 数据类名称
	DataclassName *string `json:"dataclass_name,omitempty"`

	// 数据对象名称
	DataobjectName *string `json:"dataobject_name,omitempty"`

	// 触发类型. TIMER--定时触发, EVENT--事件触发
	TriggerType *string `json:"trigger_type,omitempty"`

	// 查询起始时间，时间格式 yyyy-MM-dd'T'HH:mm:ss.SSS'Z'Z
	FromDate *string `json:"from_date,omitempty"`

	// 查询结束时间，时间格式 yyyy-MM-dd'T'HH:mm:ss.SSS'Z'Z
	ToDate *string `json:"to_date,omitempty"`

	// 时间类型. START--按照开始时间检索, END--按照结束时间检索
	DateType *string `json:"date_type,omitempty"`

	// 分页查询参数，用于指定一次查询最多的结果数，从1开始，limit 和 offset相加需要小于10000
	Limit int32 `json:"limit"`

	// 分页查询参数。用于指定查询结果的起始位置，从0开始, limit 和 offset相加需要小于10000
	Offset int32 `json:"offset"`
}

func (o ListPlaybookInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlaybookInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListPlaybookInstancesRequest", string(data)}, " ")
}
