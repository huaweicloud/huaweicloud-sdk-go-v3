package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CoEntity struct {

	// 变更对象修改后内容。
	AfterChange *string `json:"after_change,omitempty"`

	// 变更对象评审完成时间。
	ReviewCompleteTime *string `json:"review_complete_time,omitempty"`

	// 变更对象评审阶段结果。
	ReviewPhaseResult *string `json:"review_phase_result,omitempty"`

	// 变更对象评审时间。
	ReviewTime *string `json:"review_time,omitempty"`

	// 变更对象工作项修改前内容。
	BeforeChange *string `json:"before_change,omitempty"`

	// 变更对象工作项类型，此处固定为CO。
	Category *CoEntityCategory `json:"category,omitempty"`

	CcbInfo *CcbEntity `json:"ccb_info,omitempty"`

	// 变更对象决策人列表，列表中只有一个元素。
	Ccbs *[]UserEntity `json:"ccbs,omitempty"`

	// 变更类型。
	ChangeType *string `json:"change_type,omitempty"`

	// 变更对象关联的评审单ID。
	Co2review *string `json:"co2review,omitempty"`

	// 变更对象的创建人ID。
	CreatedBy *string `json:"created_by,omitempty"`

	// 变更对象创建时间。
	CreatedDate *string `json:"created_date,omitempty"`

	// 变更对象描述信息。
	Description *string `json:"description,omitempty"`

	// 变更对象ID。
	Id *string `json:"id,omitempty"`

	// 变更对象关联的工作项ID。
	IssueId *string `json:"issue_id,omitempty"`

	// 变更对象关联的工作项编号。
	IssueNumber *string `json:"issue_number,omitempty"`

	// 变更对象关联的工作项类型。
	IssueCategory *string `json:"issue_category,omitempty"`

	// 变更对象最后修改人ID。
	ModifiedBy *string `json:"modified_by,omitempty"`

	// 变更对象最后修改时间。
	ModifiedDate *string `json:"modified_date,omitempty"`

	// 变更对象评审专家Id列表（创建变更评审时使用）。
	Opinions *[]UserEntity `json:"opinions,omitempty"`

	// 变更对象评审意见。
	OpinionComments *[]ReviewOpinionEntity `json:"opinion_comments,omitempty"`

	// 变更对象评审意见（评审更新时使用）。
	ReviewComments *[]ReviewCommentEntity `json:"review_comments,omitempty"`

	// 变更对象决策意见（决策更新时使用）。
	ApprovalComments *[]ReviewCommentEntity `json:"approval_comments,omitempty"`

	// 变更对象评审专家Id列表。
	Reviewer *[]string `json:"reviewer,omitempty"`

	// 变更对象决策人ID数组。
	Approver *[]string `json:"approver,omitempty"`

	// 变更对象状态。
	Status *string `json:"status,omitempty"`
}

func (o CoEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CoEntity struct{}"
	}

	return strings.Join([]string{"CoEntity", string(data)}, " ")
}

type CoEntityCategory struct {
	value string
}

type CoEntityCategoryEnum struct {
	CO CoEntityCategory
}

func GetCoEntityCategoryEnum() CoEntityCategoryEnum {
	return CoEntityCategoryEnum{
		CO: CoEntityCategory{
			value: "CO",
		},
	}
}

func (c CoEntityCategory) Value() string {
	return c.value
}

func (c CoEntityCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CoEntityCategory) UnmarshalJSON(b []byte) error {
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
