package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SearchCustomizedFieldsRequest Request Object
type SearchCustomizedFieldsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 查询条数，即查询Y条数据。默认值50，取值范围[1,100]
	Limit *int32 `json:"limit,omitempty"`

	// 查询起始坐标，即跳过X条数据，仅支持0或limit的整数倍，不满足则向下取整。默认值0
	Offset *int32 `json:"offset,omitempty"`

	Type SearchCustomizedFieldsRequestType `json:"type"`
}

func (o SearchCustomizedFieldsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCustomizedFieldsRequest struct{}"
	}

	return strings.Join([]string{"SearchCustomizedFieldsRequest", string(data)}, " ")
}

type SearchCustomizedFieldsRequestType struct {
	value string
}

type SearchCustomizedFieldsRequestTypeEnum struct {
	TABLE     SearchCustomizedFieldsRequestType
	ATTRIBUTE SearchCustomizedFieldsRequestType
	SUBJECT   SearchCustomizedFieldsRequestType
}

func GetSearchCustomizedFieldsRequestTypeEnum() SearchCustomizedFieldsRequestTypeEnum {
	return SearchCustomizedFieldsRequestTypeEnum{
		TABLE: SearchCustomizedFieldsRequestType{
			value: "TABLE",
		},
		ATTRIBUTE: SearchCustomizedFieldsRequestType{
			value: "ATTRIBUTE",
		},
		SUBJECT: SearchCustomizedFieldsRequestType{
			value: "SUBJECT",
		},
	}
}

func (c SearchCustomizedFieldsRequestType) Value() string {
	return c.value
}

func (c SearchCustomizedFieldsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchCustomizedFieldsRequestType) UnmarshalJSON(b []byte) error {
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
