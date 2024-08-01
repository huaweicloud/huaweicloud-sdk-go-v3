package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAnalyzersRequest Request Object
type ListAnalyzersRequest struct {

	// 单页最大结果数。
	Limit *int32 `json:"limit,omitempty"`

	// 页面标记。
	Marker *string `json:"marker,omitempty"`

	// 分析器的类型。
	Type *ListAnalyzersRequestType `json:"type,omitempty"`
}

func (o ListAnalyzersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAnalyzersRequest struct{}"
	}

	return strings.Join([]string{"ListAnalyzersRequest", string(data)}, " ")
}

type ListAnalyzersRequestType struct {
	value string
}

type ListAnalyzersRequestTypeEnum struct {
	ACCOUNT                    ListAnalyzersRequestType
	ORGANIZATION               ListAnalyzersRequestType
	ACCOUNT_UNUSED_ACCESS      ListAnalyzersRequestType
	ORGANIZATION_UNUSED_ACCESS ListAnalyzersRequestType
}

func GetListAnalyzersRequestTypeEnum() ListAnalyzersRequestTypeEnum {
	return ListAnalyzersRequestTypeEnum{
		ACCOUNT: ListAnalyzersRequestType{
			value: "account",
		},
		ORGANIZATION: ListAnalyzersRequestType{
			value: "organization",
		},
		ACCOUNT_UNUSED_ACCESS: ListAnalyzersRequestType{
			value: "account_unused_access",
		},
		ORGANIZATION_UNUSED_ACCESS: ListAnalyzersRequestType{
			value: "organization_unused_access",
		},
	}
}

func (c ListAnalyzersRequestType) Value() string {
	return c.value
}

func (c ListAnalyzersRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAnalyzersRequestType) UnmarshalJSON(b []byte) error {
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
