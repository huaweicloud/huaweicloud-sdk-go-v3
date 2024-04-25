package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiParam struct {

	// API名称。
	Name *string `json:"name,omitempty"`

	// API创建人名称。
	CreateUser *string `json:"create_user,omitempty"`

	// API描述。
	Description *string `json:"description,omitempty"`

	// API标签列表。
	Tags *[]string `json:"tags,omitempty"`

	// API所使用到的数据库表名。
	TableName *string `json:"table_name,omitempty"`

	// API发布状态。
	PublishStatusType *ApiParamPublishStatusType `json:"publish_status_type,omitempty"`

	// API取数方式。
	ApiSpecificTypeStr *ApiParamApiSpecificTypeStr `json:"api_specific_type_str,omitempty"`

	// API创建开始时间。
	StartTime *string `json:"start_time,omitempty"`

	// API创建结束时间。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ApiParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiParam struct{}"
	}

	return strings.Join([]string{"ApiParam", string(data)}, " ")
}

type ApiParamPublishStatusType struct {
	value string
}

type ApiParamPublishStatusTypeEnum struct {
	PUBLISHED     ApiParamPublishStatusType
	NOT_PUBLISHED ApiParamPublishStatusType
}

func GetApiParamPublishStatusTypeEnum() ApiParamPublishStatusTypeEnum {
	return ApiParamPublishStatusTypeEnum{
		PUBLISHED: ApiParamPublishStatusType{
			value: "PUBLISHED",
		},
		NOT_PUBLISHED: ApiParamPublishStatusType{
			value: "NOT_PUBLISHED",
		},
	}
}

func (c ApiParamPublishStatusType) Value() string {
	return c.value
}

func (c ApiParamPublishStatusType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiParamPublishStatusType) UnmarshalJSON(b []byte) error {
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

type ApiParamApiSpecificTypeStr struct {
	value string
}

type ApiParamApiSpecificTypeStrEnum struct {
	API_SPECIFIC_TYPE_CONFIGURATION ApiParamApiSpecificTypeStr
	API_SPECIFIC_TYPE_SCRIPT        ApiParamApiSpecificTypeStr
	API_SPECIFIC_TYPE_MYBATIS       ApiParamApiSpecificTypeStr
	API_SPECIFIC_TYPE_GROOVY        ApiParamApiSpecificTypeStr
}

func GetApiParamApiSpecificTypeStrEnum() ApiParamApiSpecificTypeStrEnum {
	return ApiParamApiSpecificTypeStrEnum{
		API_SPECIFIC_TYPE_CONFIGURATION: ApiParamApiSpecificTypeStr{
			value: "API_SPECIFIC_TYPE_CONFIGURATION",
		},
		API_SPECIFIC_TYPE_SCRIPT: ApiParamApiSpecificTypeStr{
			value: "API_SPECIFIC_TYPE_SCRIPT",
		},
		API_SPECIFIC_TYPE_MYBATIS: ApiParamApiSpecificTypeStr{
			value: "API_SPECIFIC_TYPE_MYBATIS",
		},
		API_SPECIFIC_TYPE_GROOVY: ApiParamApiSpecificTypeStr{
			value: "API_SPECIFIC_TYPE_GROOVY",
		},
	}
}

func (c ApiParamApiSpecificTypeStr) Value() string {
	return c.value
}

func (c ApiParamApiSpecificTypeStr) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiParamApiSpecificTypeStr) UnmarshalJSON(b []byte) error {
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
