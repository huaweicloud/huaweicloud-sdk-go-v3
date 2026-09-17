package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListIpdIssueCommentsRequest Request Object
type ListIpdIssueCommentsRequest struct {

	// 项目32位ID，项目唯一标识。通过查询IPD项目列表获取，响应消息体中的id字段的值就是项目ID。
	ProjectId string `json:"project_id"`

	// 工作项唯一ID。可以通过查询工作项列表或者查询树状工作项接口获取，响应消息体中的id字段的值就是工作项ID。
	IssueId string `json:"issue_id"`

	// **参数解释**： 是否按创建日期倒序排列。 **取值范围**： - true：按创建时间倒序排列。 - false：按创建时间正序排列。 **默认取值**： 不涉及。
	DateDesc *bool `json:"date_desc,omitempty"`

	// **参数解释**： 分页索引。 **约束限制**： 不涉及 **取值范围**： 最小值1，最大值10000 **默认取值**： 1
	PageNo int32 `json:"page_no"`

	// **参数解释**： 分页大小。 **约束限制**： 不涉及 **取值范围**： 最小值5，最大值200 **默认取值**： 200
	PageSize int32 `json:"page_size"`

	// **参数解释**： 评论类型，支持多值，使用英文逗号分隔。 **取值范围**： - comment：评论 - reply：回复 - operation：系统操作。 **默认取值**： 不涉及。
	Category *ListIpdIssueCommentsRequestCategory `json:"category,omitempty"`
}

func (o ListIpdIssueCommentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpdIssueCommentsRequest struct{}"
	}

	return strings.Join([]string{"ListIpdIssueCommentsRequest", string(data)}, " ")
}

type ListIpdIssueCommentsRequestCategory struct {
	value string
}

type ListIpdIssueCommentsRequestCategoryEnum struct {
	COMMENT   ListIpdIssueCommentsRequestCategory
	REPLY     ListIpdIssueCommentsRequestCategory
	OPERATION ListIpdIssueCommentsRequestCategory
}

func GetListIpdIssueCommentsRequestCategoryEnum() ListIpdIssueCommentsRequestCategoryEnum {
	return ListIpdIssueCommentsRequestCategoryEnum{
		COMMENT: ListIpdIssueCommentsRequestCategory{
			value: "comment",
		},
		REPLY: ListIpdIssueCommentsRequestCategory{
			value: "reply",
		},
		OPERATION: ListIpdIssueCommentsRequestCategory{
			value: "operation",
		},
	}
}

func (c ListIpdIssueCommentsRequestCategory) Value() string {
	return c.value
}

func (c ListIpdIssueCommentsRequestCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListIpdIssueCommentsRequestCategory) UnmarshalJSON(b []byte) error {
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
