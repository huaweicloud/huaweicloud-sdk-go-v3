package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowGlobalEipSegmentRequest Request Object
type ShowGlobalEipSegmentRequest struct {
	GlobalEipSegmentId string `json:"global_eip_segment_id"`

	Fields *[]ShowGlobalEipSegmentRequestFields `json:"fields,omitempty"`
}

func (o ShowGlobalEipSegmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGlobalEipSegmentRequest struct{}"
	}

	return strings.Join([]string{"ShowGlobalEipSegmentRequest", string(data)}, " ")
}

type ShowGlobalEipSegmentRequestFields struct {
	value string
}

type ShowGlobalEipSegmentRequestFieldsEnum struct {
	ID                    ShowGlobalEipSegmentRequestFields
	NAME                  ShowGlobalEipSegmentRequestFields
	DESCRIPTION           ShowGlobalEipSegmentRequestFields
	DOMAIN_ID             ShowGlobalEipSegmentRequestFields
	ACCESS_SITE           ShowGlobalEipSegmentRequestFields
	GEIP_POOL_NAME        ShowGlobalEipSegmentRequestFields
	ISP                   ShowGlobalEipSegmentRequestFields
	IP_VERSION            ShowGlobalEipSegmentRequestFields
	CIDR                  ShowGlobalEipSegmentRequestFields
	CIDR_V6               ShowGlobalEipSegmentRequestFields
	FREEZEN               ShowGlobalEipSegmentRequestFields
	FREEZEN_INFO          ShowGlobalEipSegmentRequestFields
	STATUS                ShowGlobalEipSegmentRequestFields
	CREATED_AT            ShowGlobalEipSegmentRequestFields
	UPDATED_AT            ShowGlobalEipSegmentRequestFields
	INTERNET_BANDWIDTH    ShowGlobalEipSegmentRequestFields
	ASSOCIATE_INSTANCE    ShowGlobalEipSegmentRequestFields
	IS_PRE_PAID           ShowGlobalEipSegmentRequestFields
	TAGS                  ShowGlobalEipSegmentRequestFields
	SYS_TAGS              ShowGlobalEipSegmentRequestFields
	ENTERPRISE_PROJECT_ID ShowGlobalEipSegmentRequestFields
}

func GetShowGlobalEipSegmentRequestFieldsEnum() ShowGlobalEipSegmentRequestFieldsEnum {
	return ShowGlobalEipSegmentRequestFieldsEnum{
		ID: ShowGlobalEipSegmentRequestFields{
			value: "id",
		},
		NAME: ShowGlobalEipSegmentRequestFields{
			value: "name",
		},
		DESCRIPTION: ShowGlobalEipSegmentRequestFields{
			value: "description",
		},
		DOMAIN_ID: ShowGlobalEipSegmentRequestFields{
			value: "domain_id",
		},
		ACCESS_SITE: ShowGlobalEipSegmentRequestFields{
			value: "access_site",
		},
		GEIP_POOL_NAME: ShowGlobalEipSegmentRequestFields{
			value: "geip_pool_name",
		},
		ISP: ShowGlobalEipSegmentRequestFields{
			value: "isp",
		},
		IP_VERSION: ShowGlobalEipSegmentRequestFields{
			value: "ip_version",
		},
		CIDR: ShowGlobalEipSegmentRequestFields{
			value: "cidr",
		},
		CIDR_V6: ShowGlobalEipSegmentRequestFields{
			value: "cidr_v6",
		},
		FREEZEN: ShowGlobalEipSegmentRequestFields{
			value: "freezen",
		},
		FREEZEN_INFO: ShowGlobalEipSegmentRequestFields{
			value: "freezen_info",
		},
		STATUS: ShowGlobalEipSegmentRequestFields{
			value: "status",
		},
		CREATED_AT: ShowGlobalEipSegmentRequestFields{
			value: "created_at",
		},
		UPDATED_AT: ShowGlobalEipSegmentRequestFields{
			value: "updated_at",
		},
		INTERNET_BANDWIDTH: ShowGlobalEipSegmentRequestFields{
			value: "internet_bandwidth",
		},
		ASSOCIATE_INSTANCE: ShowGlobalEipSegmentRequestFields{
			value: "associate_instance",
		},
		IS_PRE_PAID: ShowGlobalEipSegmentRequestFields{
			value: "is_pre_paid",
		},
		TAGS: ShowGlobalEipSegmentRequestFields{
			value: "tags",
		},
		SYS_TAGS: ShowGlobalEipSegmentRequestFields{
			value: "sys_tags",
		},
		ENTERPRISE_PROJECT_ID: ShowGlobalEipSegmentRequestFields{
			value: "enterprise_project_id",
		},
	}
}

func (c ShowGlobalEipSegmentRequestFields) Value() string {
	return c.value
}

func (c ShowGlobalEipSegmentRequestFields) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowGlobalEipSegmentRequestFields) UnmarshalJSON(b []byte) error {
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
