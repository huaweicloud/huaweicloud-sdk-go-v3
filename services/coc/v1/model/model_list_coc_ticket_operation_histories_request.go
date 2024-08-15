package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListCocTicketOperationHistoriesRequest Request Object
type ListCocTicketOperationHistoriesRequest struct {

	// 工单类型:incident,issues_mgmt
	TicketType ListCocTicketOperationHistoriesRequestTicketType `json:"ticket_type"`

	Body *ListTicketParams `json:"body,omitempty"`
}

func (o ListCocTicketOperationHistoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCocTicketOperationHistoriesRequest struct{}"
	}

	return strings.Join([]string{"ListCocTicketOperationHistoriesRequest", string(data)}, " ")
}

type ListCocTicketOperationHistoriesRequestTicketType struct {
	value string
}

type ListCocTicketOperationHistoriesRequestTicketTypeEnum struct {
	INCIDENT    ListCocTicketOperationHistoriesRequestTicketType
	ISSUES_MGMT ListCocTicketOperationHistoriesRequestTicketType
}

func GetListCocTicketOperationHistoriesRequestTicketTypeEnum() ListCocTicketOperationHistoriesRequestTicketTypeEnum {
	return ListCocTicketOperationHistoriesRequestTicketTypeEnum{
		INCIDENT: ListCocTicketOperationHistoriesRequestTicketType{
			value: "incident",
		},
		ISSUES_MGMT: ListCocTicketOperationHistoriesRequestTicketType{
			value: "issues_mgmt",
		},
	}
}

func (c ListCocTicketOperationHistoriesRequestTicketType) Value() string {
	return c.value
}

func (c ListCocTicketOperationHistoriesRequestTicketType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCocTicketOperationHistoriesRequestTicketType) UnmarshalJSON(b []byte) error {
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
