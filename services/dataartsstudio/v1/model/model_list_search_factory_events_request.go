package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSearchFactoryEventsRequest Request Object
type ListSearchFactoryEventsRequest struct {

	// DataArts Studio实例ID。
	InstanceId string `json:"instance_id"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	WorkspaceId string `json:"workspace_id"`

	// 基线任务名称。
	Name *string `json:"name,omitempty"`

	// 作业名称
	TaskName *string `json:"task_name,omitempty"`

	// 责任人名称。
	OwnerName *string `json:"owner_name,omitempty"`

	// 事件类型:  - ERROR: 出错 - SLOW_DOWN: 变慢  默认查询全部事件类型。
	Type *string `json:"type,omitempty"`

	// 事件状态: - NEW_DISCOVERY: 新发现 - PROCESSING: 处理中 - RESTORED: 已恢复 - IGNORED: 已忽略
	Status *string `json:"status,omitempty"`

	// 排序规则，示例 happen_time_ms asc asc，表示按照优先级升序排序，有如下取值： - happen_time_ms asc asc: 按照触发事件升序。 - happen_time_ms asc desc: 按照触发事件降序。  默认不排序。
	OrderBy *string `json:"order_by,omitempty"`

	// 创建时间检索区间，起始时间戳，单位毫秒。默认为当天0点0分0秒，当end_time有值时，该值不能为空。
	StartTime *int64 `json:"start_time,omitempty"`

	// 创建时间检索区间，终止时间戳，单位毫秒。默认为当天23点59分59秒，当start_time有值时，该值不能为空，最大查询范围为180天。
	EndTime *int64 `json:"end_time,omitempty"`

	// 分页列表的页数，默认值为1。取值范围大于等于1。
	Offset *int32 `json:"offset,omitempty"`

	// 分页返回结果，指定每页最大记录数。范围[1,100] 默认值：10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSearchFactoryEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSearchFactoryEventsRequest struct{}"
	}

	return strings.Join([]string{"ListSearchFactoryEventsRequest", string(data)}, " ")
}
