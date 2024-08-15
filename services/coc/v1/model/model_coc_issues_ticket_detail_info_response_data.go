package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CocIssuesTicketDetailInfoResponseData CocIssuesTicketDetailInfoResponseData
type CocIssuesTicketDetailInfoResponseData struct {

	// 问题类型
	TicketType *string `json:"ticket_type,omitempty"`

	// 问题等级
	Level *string `json:"level,omitempty"`

	// 影响服务
	ImpactedCloudServices *string `json:"impacted_cloud_services,omitempty"`

	// 责任服务
	RootCauseCloudService *string `json:"root_cause_cloud_service,omitempty"`

	// 根因类型
	RootCauseType *string `json:"root_cause_type,omitempty"`

	// 根因分析
	RootCauseComment *string `json:"root_cause_comment,omitempty"`

	// 解决方案
	Solution *string `json:"solution,omitempty"`

	// 问题接口人id
	IssueContactPerson *string `json:"issue_contact_person,omitempty"`

	// 重现概率
	ReproduceProbability *string `json:"reproduce_probability,omitempty"`

	// 发现问题的版本号
	IssueVersion *string `json:"issue_version,omitempty"`

	// 问题标题
	Title *string `json:"title,omitempty"`

	// 排班类型
	VirtualScheduleType *string `json:"virtual_schedule_type,omitempty"`

	// 区域
	Regions *string `json:"regions,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 发现时间
	FountTime *int64 `json:"fount_time,omitempty"`

	// 是否共性问题
	IsCommonIssue *bool `json:"is_common_issue,omitempty"`

	// 是否需要变更
	IsNeedChange *bool `json:"is_need_change,omitempty"`

	// 创建人
	Creator *string `json:"creator,omitempty"`

	// 操作人
	Operator *string `json:"operator,omitempty"`

	// 问题单id
	TicketId *string `json:"ticket_id,omitempty"`

	// 责任人
	Assignee *string `json:"assignee,omitempty"`

	// 问题单状态
	WorkFlowStatus *string `json:"work_flow_status,omitempty"`

	// 阶段
	Phase *string `json:"phase,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 是否删除
	IsDeleted *bool `json:"is_deleted,omitempty"`

	// 枚举列表
	EnumDataList *[]TicketInfoEnumData `json:"enum_data_list,omitempty"`
}

func (o CocIssuesTicketDetailInfoResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CocIssuesTicketDetailInfoResponseData struct{}"
	}

	return strings.Join([]string{"CocIssuesTicketDetailInfoResponseData", string(data)}, " ")
}
