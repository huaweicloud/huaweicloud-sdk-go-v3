package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ListEventsResponse struct {

	// API版本。
	ApiVersion *string `json:"api_version,omitempty"`

	// 资源种类。
	Kind *ListEventsResponseKind `json:"kind,omitempty"`

	// 事件项。
	Items          *[]EventItem `json:"items,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventsResponse struct{}"
	}

	return strings.Join([]string{"ListEventsResponse", string(data)}, " ")
}

type ListEventsResponseKind struct {
	value string
}

type ListEventsResponseKindEnum struct {
	COMPONENT_EVENT ListEventsResponseKind
}

func GetListEventsResponseKindEnum() ListEventsResponseKindEnum {
	return ListEventsResponseKindEnum{
		COMPONENT_EVENT: ListEventsResponseKind{
			value: "ComponentEvent",
		},
	}
}

func (c ListEventsResponseKind) Value() string {
	return c.value
}

func (c ListEventsResponseKind) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventsResponseKind) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
