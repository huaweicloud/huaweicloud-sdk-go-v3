package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTenantGeipSupportInstancesRequest Request Object
type ListTenantGeipSupportInstancesRequest struct {
	AccessSite string `json:"access_site"`

	Fields *[]ListTenantGeipSupportInstancesRequestFields `json:"fields,omitempty"`
}

func (o ListTenantGeipSupportInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantGeipSupportInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListTenantGeipSupportInstancesRequest", string(data)}, " ")
}

type ListTenantGeipSupportInstancesRequestFields struct {
	value string
}

type ListTenantGeipSupportInstancesRequestFieldsEnum struct {
	ID                  ListTenantGeipSupportInstancesRequestFields
	INSTANCE_TYPE       ListTenantGeipSupportInstancesRequestFields
	REGION_ID           ListTenantGeipSupportInstancesRequestFields
	PUBLIC_BORDER_GROUP ListTenantGeipSupportInstancesRequestFields
	STATUS              ListTenantGeipSupportInstancesRequestFields
	CREATED_AT          ListTenantGeipSupportInstancesRequestFields
	UPDATED_AT          ListTenantGeipSupportInstancesRequestFields
}

func GetListTenantGeipSupportInstancesRequestFieldsEnum() ListTenantGeipSupportInstancesRequestFieldsEnum {
	return ListTenantGeipSupportInstancesRequestFieldsEnum{
		ID: ListTenantGeipSupportInstancesRequestFields{
			value: "id",
		},
		INSTANCE_TYPE: ListTenantGeipSupportInstancesRequestFields{
			value: "instance_type",
		},
		REGION_ID: ListTenantGeipSupportInstancesRequestFields{
			value: "region_id",
		},
		PUBLIC_BORDER_GROUP: ListTenantGeipSupportInstancesRequestFields{
			value: "public_border_group",
		},
		STATUS: ListTenantGeipSupportInstancesRequestFields{
			value: "status",
		},
		CREATED_AT: ListTenantGeipSupportInstancesRequestFields{
			value: "created_at",
		},
		UPDATED_AT: ListTenantGeipSupportInstancesRequestFields{
			value: "updated_at",
		},
	}
}

func (c ListTenantGeipSupportInstancesRequestFields) Value() string {
	return c.value
}

func (c ListTenantGeipSupportInstancesRequestFields) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTenantGeipSupportInstancesRequestFields) UnmarshalJSON(b []byte) error {
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
