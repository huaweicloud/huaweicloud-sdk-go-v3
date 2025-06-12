package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListRecycleBinServersRequest Request Object
type ListRecycleBinServersRequest struct {

	// 所有租户 管理员字段 1: 返回所有租户的VM 0: 返回当前租户的VM，如果为0，与不设置该查询字段的效果一致
	AllTenants *string `json:"all_tenants,omitempty"`

	AvailabilityZone *string `json:"availability_zone,omitempty"`

	ExpectFields *ListRecycleBinServersRequestExpectFields `json:"expect-fields,omitempty"`

	IpAddress *string `json:"ip_address,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Name *string `json:"name,omitempty"`

	Offset *string `json:"offset,omitempty"`

	Tags *[]string `json:"tags,omitempty"`

	TagsKey *[]string `json:"tags_key,omitempty"`
}

func (o ListRecycleBinServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecycleBinServersRequest struct{}"
	}

	return strings.Join([]string{"ListRecycleBinServersRequest", string(data)}, " ")
}

type ListRecycleBinServersRequestExpectFields struct {
	value string
}

type ListRecycleBinServersRequestExpectFieldsEnum struct {
	AVAILABLE_VALUES__POWER_STATE ListRecycleBinServersRequestExpectFields
	HOST_STATUS                   ListRecycleBinServersRequestExpectFields
	HOSTNAME                      ListRecycleBinServersRequestExpectFields
	HYPERVISOR_HOSTNAME           ListRecycleBinServersRequestExpectFields
	USER_DATA                     ListRecycleBinServersRequestExpectFields
	KEY_NAME                      ListRecycleBinServersRequestExpectFields
	ROOT_DEVICE_NAME              ListRecycleBinServersRequestExpectFields
	VOLUMES_ATTACHED              ListRecycleBinServersRequestExpectFields
	SECURITY_GROUPS               ListRecycleBinServersRequestExpectFields
	ADDRESSES                     ListRecycleBinServersRequestExpectFields
	IMAGE                         ListRecycleBinServersRequestExpectFields
	METADATA                      ListRecycleBinServersRequestExpectFields
	TAGS                          ListRecycleBinServersRequestExpectFields
	SYSTEM_TAGS                   ListRecycleBinServersRequestExpectFields
	DEDICATED_HOST_ID             ListRecycleBinServersRequestExpectFields
	ENTERPRISE_PROJECT_ID         ListRecycleBinServersRequestExpectFields
	CPU_OPTIONS                   ListRecycleBinServersRequestExpectFields
}

func GetListRecycleBinServersRequestExpectFieldsEnum() ListRecycleBinServersRequestExpectFieldsEnum {
	return ListRecycleBinServersRequestExpectFieldsEnum{
		AVAILABLE_VALUES__POWER_STATE: ListRecycleBinServersRequestExpectFields{
			value: "Available values : power_state",
		},
		HOST_STATUS: ListRecycleBinServersRequestExpectFields{
			value: "host_status",
		},
		HOSTNAME: ListRecycleBinServersRequestExpectFields{
			value: "hostname",
		},
		HYPERVISOR_HOSTNAME: ListRecycleBinServersRequestExpectFields{
			value: "hypervisor_hostname",
		},
		USER_DATA: ListRecycleBinServersRequestExpectFields{
			value: "user_data",
		},
		KEY_NAME: ListRecycleBinServersRequestExpectFields{
			value: "key_name",
		},
		ROOT_DEVICE_NAME: ListRecycleBinServersRequestExpectFields{
			value: "root_device_name",
		},
		VOLUMES_ATTACHED: ListRecycleBinServersRequestExpectFields{
			value: "volumes_attached",
		},
		SECURITY_GROUPS: ListRecycleBinServersRequestExpectFields{
			value: "security_groups",
		},
		ADDRESSES: ListRecycleBinServersRequestExpectFields{
			value: "addresses",
		},
		IMAGE: ListRecycleBinServersRequestExpectFields{
			value: "image",
		},
		METADATA: ListRecycleBinServersRequestExpectFields{
			value: "metadata",
		},
		TAGS: ListRecycleBinServersRequestExpectFields{
			value: "tags",
		},
		SYSTEM_TAGS: ListRecycleBinServersRequestExpectFields{
			value: "system_tags",
		},
		DEDICATED_HOST_ID: ListRecycleBinServersRequestExpectFields{
			value: "dedicated_host_id",
		},
		ENTERPRISE_PROJECT_ID: ListRecycleBinServersRequestExpectFields{
			value: "enterprise_project_id",
		},
		CPU_OPTIONS: ListRecycleBinServersRequestExpectFields{
			value: "cpu_options",
		},
	}
}

func (c ListRecycleBinServersRequestExpectFields) Value() string {
	return c.value
}

func (c ListRecycleBinServersRequestExpectFields) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRecycleBinServersRequestExpectFields) UnmarshalJSON(b []byte) error {
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
