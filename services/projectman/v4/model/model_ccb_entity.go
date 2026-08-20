package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CcbEntity struct {

	// 审批时间。
	ApprovalTime *string `json:"approval_time,omitempty"`

	// 工作项类型，审批对象固定为CCB。
	Category *CcbEntityCategory `json:"category,omitempty"`

	// 审批对象关联的评审单ID。
	Ccb2review *string `json:"ccb2review,omitempty"`

	// 关联的变更对象ID。
	CoId *string `json:"co_id,omitempty"`

	// 审批对象ID。
	Id *string `json:"id,omitempty"`

	Owner *UserEntity `json:"owner,omitempty"`

	// 审批意见。
	ApprovalComments *string `json:"approval_comments,omitempty"`
}

func (o CcbEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CcbEntity struct{}"
	}

	return strings.Join([]string{"CcbEntity", string(data)}, " ")
}

type CcbEntityCategory struct {
	value string
}

type CcbEntityCategoryEnum struct {
	CCB CcbEntityCategory
}

func GetCcbEntityCategoryEnum() CcbEntityCategoryEnum {
	return CcbEntityCategoryEnum{
		CCB: CcbEntityCategory{
			value: "CCB",
		},
	}
}

func (c CcbEntityCategory) Value() string {
	return c.value
}

func (c CcbEntityCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CcbEntityCategory) UnmarshalJSON(b []byte) error {
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
