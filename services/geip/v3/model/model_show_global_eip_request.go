package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowGlobalEipRequest Request Object
type ShowGlobalEipRequest struct {
	GlobalEipId string `json:"global_eip_id"`

	Fields *[]ShowGlobalEipRequestFields `json:"fields,omitempty"`
}

func (o ShowGlobalEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGlobalEipRequest struct{}"
	}

	return strings.Join([]string{"ShowGlobalEipRequest", string(data)}, " ")
}

type ShowGlobalEipRequestFields struct {
	value string
}

type ShowGlobalEipRequestFieldsEnum struct {
	ID                               ShowGlobalEipRequestFields
	NAME                             ShowGlobalEipRequestFields
	DESCRIPTION                      ShowGlobalEipRequestFields
	DOMAIN_ID                        ShowGlobalEipRequestFields
	ACCESS_SITE                      ShowGlobalEipRequestFields
	GEIP_POOL_NAME                   ShowGlobalEipRequestFields
	ISP                              ShowGlobalEipRequestFields
	IP_VERSION                       ShowGlobalEipRequestFields
	IP_ADDRESS                       ShowGlobalEipRequestFields
	IPV6_ADDRESS                     ShowGlobalEipRequestFields
	FREEZEN                          ShowGlobalEipRequestFields
	FREEZEN_INFO                     ShowGlobalEipRequestFields
	POLLUTED                         ShowGlobalEipRequestFields
	STATUS                           ShowGlobalEipRequestFields
	CREATED_AT                       ShowGlobalEipRequestFields
	UPDATED_AT                       ShowGlobalEipRequestFields
	INTERNET_BANDWIDTH_INFO          ShowGlobalEipRequestFields
	GLOBAL_CONNECTION_BANDWIDTH_INFO ShowGlobalEipRequestFields
	ASSOCIATE_INSTANCE_INFO          ShowGlobalEipRequestFields
	IS_PRE_PAID                      ShowGlobalEipRequestFields
	TAGS                             ShowGlobalEipRequestFields
	SYS_TAGS                         ShowGlobalEipRequestFields
	ENTERPRISE_PROJECT_ID            ShowGlobalEipRequestFields
}

func GetShowGlobalEipRequestFieldsEnum() ShowGlobalEipRequestFieldsEnum {
	return ShowGlobalEipRequestFieldsEnum{
		ID: ShowGlobalEipRequestFields{
			value: "id",
		},
		NAME: ShowGlobalEipRequestFields{
			value: "name",
		},
		DESCRIPTION: ShowGlobalEipRequestFields{
			value: "description",
		},
		DOMAIN_ID: ShowGlobalEipRequestFields{
			value: "domain_id",
		},
		ACCESS_SITE: ShowGlobalEipRequestFields{
			value: "access_site",
		},
		GEIP_POOL_NAME: ShowGlobalEipRequestFields{
			value: "geip_pool_name",
		},
		ISP: ShowGlobalEipRequestFields{
			value: "isp",
		},
		IP_VERSION: ShowGlobalEipRequestFields{
			value: "ip_version",
		},
		IP_ADDRESS: ShowGlobalEipRequestFields{
			value: "ip_address",
		},
		IPV6_ADDRESS: ShowGlobalEipRequestFields{
			value: "ipv6_address",
		},
		FREEZEN: ShowGlobalEipRequestFields{
			value: "freezen",
		},
		FREEZEN_INFO: ShowGlobalEipRequestFields{
			value: "freezen_info",
		},
		POLLUTED: ShowGlobalEipRequestFields{
			value: "polluted",
		},
		STATUS: ShowGlobalEipRequestFields{
			value: "status",
		},
		CREATED_AT: ShowGlobalEipRequestFields{
			value: "created_at",
		},
		UPDATED_AT: ShowGlobalEipRequestFields{
			value: "updated_at",
		},
		INTERNET_BANDWIDTH_INFO: ShowGlobalEipRequestFields{
			value: "internet_bandwidth_info",
		},
		GLOBAL_CONNECTION_BANDWIDTH_INFO: ShowGlobalEipRequestFields{
			value: "global_connection_bandwidth_info",
		},
		ASSOCIATE_INSTANCE_INFO: ShowGlobalEipRequestFields{
			value: "associate_instance_info",
		},
		IS_PRE_PAID: ShowGlobalEipRequestFields{
			value: "is_pre_paid",
		},
		TAGS: ShowGlobalEipRequestFields{
			value: "tags",
		},
		SYS_TAGS: ShowGlobalEipRequestFields{
			value: "sys_tags",
		},
		ENTERPRISE_PROJECT_ID: ShowGlobalEipRequestFields{
			value: "enterprise_project_id",
		},
	}
}

func (c ShowGlobalEipRequestFields) Value() string {
	return c.value
}

func (c ShowGlobalEipRequestFields) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowGlobalEipRequestFields) UnmarshalJSON(b []byte) error {
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
