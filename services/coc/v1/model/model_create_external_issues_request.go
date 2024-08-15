package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateExternalIssuesRequest CreateExternalIssuesRequest 创单请求参数
type CreateExternalIssuesRequest struct {

	// 创建人id
	Creator *string `json:"creator,omitempty"`

	// 标题
	Title string `json:"title"`

	// 描述
	Description string `json:"description"`

	// 区域Code,最大100个
	Regions *[]string `json:"regions,omitempty"`

	// 企业项目id
	EnterpriseProject *string `json:"enterprise_project,omitempty"`

	// 问题来源 issues_source_1000 事件 issues_source_2000 Warroom issues_source_3000 告警
	Source *string `json:"source,omitempty"`

	// 问题来源id
	SourceId *string `json:"source_id,omitempty"`

	// 发现时间
	FountTime *int64 `json:"fount_time,omitempty"`

	// 影响应用ID，最多10条
	ImpactedCloudServices []string `json:"impacted_cloud_services"`

	// 问题级别 issues_level_1000 致命 issues_level_2000 严重 issues_level_3000 一般
	Level CreateExternalIssuesRequestLevel `json:"level"`

	// 问题类型 issues_type_1000  功能性问题 issues_type_2000  性能问题 issues_type_3000  可靠性问题 issues_type_4000  兼容性问题 issues_type_5000  用户体验问题 issues_type_6000  可维护性问题 issues_type_7000  变更类问题 issues_type_8000  安全问题 issues_type_9000  工程实施类 issues_type_10000 交付部署问题 issues_type_11000 LLD规划问题 issues_type_12000 供用商问题 issues_type_13000 咨询类问题 issues_type_14000 需求类问题 issues_type_15000 其他问题
	TicketType CreateExternalIssuesRequestTicketType `json:"ticket_type"`

	// 重现概率 issues_reproduce_probability_1000 有条件必现 issues_reproduce_probability_2000 有条件概率重现 issues_reproduce_probability_3000 无规律重现 issues_reproduce_probability_4000 很难重现
	ReproduceProbability CreateExternalIssuesRequestReproduceProbability `json:"reproduce_probability"`

	// 责任应用ID，必填,限1条
	RootCauseCloudService []string `json:"root_cause_cloud_service"`

	// 排班类型 参考：枚举 issues_mgmt_virtual_schedule_type_1000 排班,schedule_scenes排班场景,virtual_schedule_role排班角色必填,指定排班 issues_mgmt_virtual_schedule_type_2000 个人,issue_contact_person字段必填,指定责任人
	VirtualScheduleType CreateExternalIssuesRequestVirtualScheduleType `json:"virtual_schedule_type"`

	// 排班场景id
	ScheduleScenes *string `json:"schedule_scenes,omitempty"`

	// 排班角色id
	VirtualScheduleRole *string `json:"virtual_schedule_role,omitempty"`

	// 问题责任人id
	IssueContactPerson *string `json:"issue_contact_person,omitempty"`
}

func (o CreateExternalIssuesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExternalIssuesRequest struct{}"
	}

	return strings.Join([]string{"CreateExternalIssuesRequest", string(data)}, " ")
}

type CreateExternalIssuesRequestLevel struct {
	value string
}

type CreateExternalIssuesRequestLevelEnum struct {
	ISSUES_LEVEL_1000 CreateExternalIssuesRequestLevel
	ISSUES_LEVEL_2000 CreateExternalIssuesRequestLevel
	ISSUES_LEVEL_3000 CreateExternalIssuesRequestLevel
}

func GetCreateExternalIssuesRequestLevelEnum() CreateExternalIssuesRequestLevelEnum {
	return CreateExternalIssuesRequestLevelEnum{
		ISSUES_LEVEL_1000: CreateExternalIssuesRequestLevel{
			value: "issues_level_1000",
		},
		ISSUES_LEVEL_2000: CreateExternalIssuesRequestLevel{
			value: "issues_level_2000",
		},
		ISSUES_LEVEL_3000: CreateExternalIssuesRequestLevel{
			value: "issues_level_3000",
		},
	}
}

func (c CreateExternalIssuesRequestLevel) Value() string {
	return c.value
}

func (c CreateExternalIssuesRequestLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateExternalIssuesRequestLevel) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CreateExternalIssuesRequestTicketType struct {
	value string
}

type CreateExternalIssuesRequestTicketTypeEnum struct {
	ISSUES_TYPE_1000  CreateExternalIssuesRequestTicketType
	ISSUES_TYPE_2000  CreateExternalIssuesRequestTicketType
	ISSUES_TYPE_3000  CreateExternalIssuesRequestTicketType
	ISSUES_TYPE_4000  CreateExternalIssuesRequestTicketType
	ISSUES_TYPE_5000  CreateExternalIssuesRequestTicketType
	ISSUES_TYPE_6000  CreateExternalIssuesRequestTicketType
	ISSUES_TYPE_7000  CreateExternalIssuesRequestTicketType
	ISSUES_TYPE_8000  CreateExternalIssuesRequestTicketType
	ISSUES_TYPE_9000  CreateExternalIssuesRequestTicketType
	ISSUES_TYPE_10000 CreateExternalIssuesRequestTicketType
	ISSUES_TYPE_11000 CreateExternalIssuesRequestTicketType
	ISSUES_TYPE_12000 CreateExternalIssuesRequestTicketType
	ISSUES_TYPE_13000 CreateExternalIssuesRequestTicketType
	ISSUES_TYPE_14000 CreateExternalIssuesRequestTicketType
	ISSUES_TYPE_15000 CreateExternalIssuesRequestTicketType
}

func GetCreateExternalIssuesRequestTicketTypeEnum() CreateExternalIssuesRequestTicketTypeEnum {
	return CreateExternalIssuesRequestTicketTypeEnum{
		ISSUES_TYPE_1000: CreateExternalIssuesRequestTicketType{
			value: "issues_type_1000",
		},
		ISSUES_TYPE_2000: CreateExternalIssuesRequestTicketType{
			value: "issues_type_2000",
		},
		ISSUES_TYPE_3000: CreateExternalIssuesRequestTicketType{
			value: "issues_type_3000",
		},
		ISSUES_TYPE_4000: CreateExternalIssuesRequestTicketType{
			value: "issues_type_4000",
		},
		ISSUES_TYPE_5000: CreateExternalIssuesRequestTicketType{
			value: "issues_type_5000",
		},
		ISSUES_TYPE_6000: CreateExternalIssuesRequestTicketType{
			value: "issues_type_6000",
		},
		ISSUES_TYPE_7000: CreateExternalIssuesRequestTicketType{
			value: "issues_type_7000",
		},
		ISSUES_TYPE_8000: CreateExternalIssuesRequestTicketType{
			value: "issues_type_8000",
		},
		ISSUES_TYPE_9000: CreateExternalIssuesRequestTicketType{
			value: "issues_type_9000",
		},
		ISSUES_TYPE_10000: CreateExternalIssuesRequestTicketType{
			value: "issues_type_10000",
		},
		ISSUES_TYPE_11000: CreateExternalIssuesRequestTicketType{
			value: "issues_type_11000",
		},
		ISSUES_TYPE_12000: CreateExternalIssuesRequestTicketType{
			value: "issues_type_12000",
		},
		ISSUES_TYPE_13000: CreateExternalIssuesRequestTicketType{
			value: "issues_type_13000",
		},
		ISSUES_TYPE_14000: CreateExternalIssuesRequestTicketType{
			value: "issues_type_14000",
		},
		ISSUES_TYPE_15000: CreateExternalIssuesRequestTicketType{
			value: "issues_type_15000",
		},
	}
}

func (c CreateExternalIssuesRequestTicketType) Value() string {
	return c.value
}

func (c CreateExternalIssuesRequestTicketType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateExternalIssuesRequestTicketType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CreateExternalIssuesRequestReproduceProbability struct {
	value string
}

type CreateExternalIssuesRequestReproduceProbabilityEnum struct {
	ISSUES_REPRODUCE_PROBABILITY_1000 CreateExternalIssuesRequestReproduceProbability
	ISSUES_REPRODUCE_PROBABILITY_2000 CreateExternalIssuesRequestReproduceProbability
	ISSUES_REPRODUCE_PROBABILITY_3000 CreateExternalIssuesRequestReproduceProbability
	ISSUES_REPRODUCE_PROBABILITY_4000 CreateExternalIssuesRequestReproduceProbability
}

func GetCreateExternalIssuesRequestReproduceProbabilityEnum() CreateExternalIssuesRequestReproduceProbabilityEnum {
	return CreateExternalIssuesRequestReproduceProbabilityEnum{
		ISSUES_REPRODUCE_PROBABILITY_1000: CreateExternalIssuesRequestReproduceProbability{
			value: "issues_reproduce_probability_1000",
		},
		ISSUES_REPRODUCE_PROBABILITY_2000: CreateExternalIssuesRequestReproduceProbability{
			value: "issues_reproduce_probability_2000",
		},
		ISSUES_REPRODUCE_PROBABILITY_3000: CreateExternalIssuesRequestReproduceProbability{
			value: "issues_reproduce_probability_3000",
		},
		ISSUES_REPRODUCE_PROBABILITY_4000: CreateExternalIssuesRequestReproduceProbability{
			value: "issues_reproduce_probability_4000",
		},
	}
}

func (c CreateExternalIssuesRequestReproduceProbability) Value() string {
	return c.value
}

func (c CreateExternalIssuesRequestReproduceProbability) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateExternalIssuesRequestReproduceProbability) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CreateExternalIssuesRequestVirtualScheduleType struct {
	value string
}

type CreateExternalIssuesRequestVirtualScheduleTypeEnum struct {
	ISSUES_MGMT_VIRTUAL_SCHEDULE_TYPE_1000 CreateExternalIssuesRequestVirtualScheduleType
	ISSUES_MGMT_VIRTUAL_SCHEDULE_TYPE_2000 CreateExternalIssuesRequestVirtualScheduleType
}

func GetCreateExternalIssuesRequestVirtualScheduleTypeEnum() CreateExternalIssuesRequestVirtualScheduleTypeEnum {
	return CreateExternalIssuesRequestVirtualScheduleTypeEnum{
		ISSUES_MGMT_VIRTUAL_SCHEDULE_TYPE_1000: CreateExternalIssuesRequestVirtualScheduleType{
			value: "issues_mgmt_virtual_schedule_type_1000",
		},
		ISSUES_MGMT_VIRTUAL_SCHEDULE_TYPE_2000: CreateExternalIssuesRequestVirtualScheduleType{
			value: "issues_mgmt_virtual_schedule_type_2000",
		},
	}
}

func (c CreateExternalIssuesRequestVirtualScheduleType) Value() string {
	return c.value
}

func (c CreateExternalIssuesRequestVirtualScheduleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateExternalIssuesRequestVirtualScheduleType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
