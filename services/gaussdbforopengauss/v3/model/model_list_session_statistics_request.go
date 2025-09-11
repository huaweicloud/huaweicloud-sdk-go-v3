package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSessionStatisticsRequest Request Object
type ListSessionStatisticsRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *ListSessionStatisticsRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**  统计维度。 **约束限制**: 不涉及。 **取值范围** - usename：用户名。 - client_addr：访问来源。 - datname：数据库名。  **默认取值**: 不涉及。
	Dimension ListSessionStatisticsRequestDimension `json:"dimension"`

	// **参数解释** 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。 **约束限制**: 不涉及。 **取值范围** [0, 2^31-1] **默认取值** 默认为0（偏移0条数据，表示从第一条数据开始查询）。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释** 查询记录数。 **约束限制**: 不涉及。 **取值范围** [1, 100] **默认取值** 默认为100。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释** 实时会话统计排序字段。 **约束限制**: 不涉及。 **取值范围** - active_num：活跃会话数。 - total_num：总会话数。  **默认取值**: 不涉及。
	OrderField *ListSessionStatisticsRequestOrderField `json:"order_field,omitempty"`

	// **参数解释** 实时会话统计排序方式。 **约束限制**: 不涉及。 **取值范围** - ASC：根据order_field值升序。 - DESC：根据order_field值降序。  **默认取值** ASC
	Order *ListSessionStatisticsRequestOrder `json:"order,omitempty"`
}

func (o ListSessionStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSessionStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListSessionStatisticsRequest", string(data)}, " ")
}

type ListSessionStatisticsRequestXLanguage struct {
	value string
}

type ListSessionStatisticsRequestXLanguageEnum struct {
	ZH_CN ListSessionStatisticsRequestXLanguage
	EN_US ListSessionStatisticsRequestXLanguage
}

func GetListSessionStatisticsRequestXLanguageEnum() ListSessionStatisticsRequestXLanguageEnum {
	return ListSessionStatisticsRequestXLanguageEnum{
		ZH_CN: ListSessionStatisticsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListSessionStatisticsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListSessionStatisticsRequestXLanguage) Value() string {
	return c.value
}

func (c ListSessionStatisticsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSessionStatisticsRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ListSessionStatisticsRequestDimension struct {
	value string
}

type ListSessionStatisticsRequestDimensionEnum struct {
	USENAME     ListSessionStatisticsRequestDimension
	CLIENT_ADDR ListSessionStatisticsRequestDimension
	DATNAME     ListSessionStatisticsRequestDimension
}

func GetListSessionStatisticsRequestDimensionEnum() ListSessionStatisticsRequestDimensionEnum {
	return ListSessionStatisticsRequestDimensionEnum{
		USENAME: ListSessionStatisticsRequestDimension{
			value: "usename",
		},
		CLIENT_ADDR: ListSessionStatisticsRequestDimension{
			value: "client_addr",
		},
		DATNAME: ListSessionStatisticsRequestDimension{
			value: "datname",
		},
	}
}

func (c ListSessionStatisticsRequestDimension) Value() string {
	return c.value
}

func (c ListSessionStatisticsRequestDimension) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSessionStatisticsRequestDimension) UnmarshalJSON(b []byte) error {
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

type ListSessionStatisticsRequestOrderField struct {
	value string
}

type ListSessionStatisticsRequestOrderFieldEnum struct {
	ACTIVE_NUM ListSessionStatisticsRequestOrderField
	TOTAL_NUM  ListSessionStatisticsRequestOrderField
}

func GetListSessionStatisticsRequestOrderFieldEnum() ListSessionStatisticsRequestOrderFieldEnum {
	return ListSessionStatisticsRequestOrderFieldEnum{
		ACTIVE_NUM: ListSessionStatisticsRequestOrderField{
			value: "active_num",
		},
		TOTAL_NUM: ListSessionStatisticsRequestOrderField{
			value: "total_num",
		},
	}
}

func (c ListSessionStatisticsRequestOrderField) Value() string {
	return c.value
}

func (c ListSessionStatisticsRequestOrderField) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSessionStatisticsRequestOrderField) UnmarshalJSON(b []byte) error {
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

type ListSessionStatisticsRequestOrder struct {
	value string
}

type ListSessionStatisticsRequestOrderEnum struct {
	ASC  ListSessionStatisticsRequestOrder
	DESC ListSessionStatisticsRequestOrder
}

func GetListSessionStatisticsRequestOrderEnum() ListSessionStatisticsRequestOrderEnum {
	return ListSessionStatisticsRequestOrderEnum{
		ASC: ListSessionStatisticsRequestOrder{
			value: "ASC",
		},
		DESC: ListSessionStatisticsRequestOrder{
			value: "DESC",
		},
	}
}

func (c ListSessionStatisticsRequestOrder) Value() string {
	return c.value
}

func (c ListSessionStatisticsRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSessionStatisticsRequestOrder) UnmarshalJSON(b []byte) error {
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
