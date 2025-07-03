package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateTicketRequest Request Object
type CreateTicketRequest struct {

	// **参数解释：** 需要创建的工单类型，如需创建变更单传值change，如需创建问题单传值issues_mgmt。 **约束限制：** 不涉及 **取值范围：** 枚举值：change/issues_mgmt 当前支持的工单类型：变更单传值change，问题单传值issues_mgmt **默认取值：** incident
	TicketType CreateTicketRequestTicketType `json:"ticket_type"`

	Body *CreateCocTicketRequestBody `json:"body,omitempty"`
}

func (o CreateTicketRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTicketRequest struct{}"
	}

	return strings.Join([]string{"CreateTicketRequest", string(data)}, " ")
}

type CreateTicketRequestTicketType struct {
	value string
}

type CreateTicketRequestTicketTypeEnum struct {
	CHANGE      CreateTicketRequestTicketType
	ISSUES_MGMT CreateTicketRequestTicketType
}

func GetCreateTicketRequestTicketTypeEnum() CreateTicketRequestTicketTypeEnum {
	return CreateTicketRequestTicketTypeEnum{
		CHANGE: CreateTicketRequestTicketType{
			value: "change",
		},
		ISSUES_MGMT: CreateTicketRequestTicketType{
			value: "issues_mgmt",
		},
	}
}

func (c CreateTicketRequestTicketType) Value() string {
	return c.value
}

func (c CreateTicketRequestTicketType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTicketRequestTicketType) UnmarshalJSON(b []byte) error {
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
