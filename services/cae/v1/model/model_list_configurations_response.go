package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListConfigurationsResponse Response Object
type ListConfigurationsResponse struct {

	// API版本。
	ApiVersion *string `json:"api_version,omitempty"`

	// 资源种类。
	Kind *ListConfigurationsResponseKind `json:"kind,omitempty"`

	// 组件配置列表。
	Items          *[]Configuration `json:"items,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"ListConfigurationsResponse", string(data)}, " ")
}

type ListConfigurationsResponseKind struct {
	value string
}

type ListConfigurationsResponseKindEnum struct {
	CONFIGURATION ListConfigurationsResponseKind
}

func GetListConfigurationsResponseKindEnum() ListConfigurationsResponseKindEnum {
	return ListConfigurationsResponseKindEnum{
		CONFIGURATION: ListConfigurationsResponseKind{
			value: "Configuration",
		},
	}
}

func (c ListConfigurationsResponseKind) Value() string {
	return c.value
}

func (c ListConfigurationsResponseKind) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListConfigurationsResponseKind) UnmarshalJSON(b []byte) error {
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
