package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowInternetBandwidthRequest Request Object
type ShowInternetBandwidthRequest struct {
	InternetBandwidthId string `json:"internet_bandwidth_id"`

	Fields *[]ShowInternetBandwidthRequestFields `json:"fields,omitempty"`
}

func (o ShowInternetBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInternetBandwidthRequest struct{}"
	}

	return strings.Join([]string{"ShowInternetBandwidthRequest", string(data)}, " ")
}

type ShowInternetBandwidthRequestFields struct {
	value string
}

type ShowInternetBandwidthRequestFieldsEnum struct {
	ID                    ShowInternetBandwidthRequestFields
	NAME                  ShowInternetBandwidthRequestFields
	ISP                   ShowInternetBandwidthRequestFields
	INGRESS_SIZE          ShowInternetBandwidthRequestFields
	ACCESS_SITE           ShowInternetBandwidthRequestFields
	SIZE                  ShowInternetBandwidthRequestFields
	DESCRIPTION           ShowInternetBandwidthRequestFields
	CHARGE_MODE           ShowInternetBandwidthRequestFields
	RATIO_95PEAK          ShowInternetBandwidthRequestFields
	FREEZEN_INFO          ShowInternetBandwidthRequestFields
	DOMAIN_ID             ShowInternetBandwidthRequestFields
	STATUS                ShowInternetBandwidthRequestFields
	CREATED_AT            ShowInternetBandwidthRequestFields
	UPDATED_AT            ShowInternetBandwidthRequestFields
	IS_PRE_PAID           ShowInternetBandwidthRequestFields
	TYPE                  ShowInternetBandwidthRequestFields
	TAGS                  ShowInternetBandwidthRequestFields
	SYS_TAGS              ShowInternetBandwidthRequestFields
	ENTERPRISE_PROJECT_ID ShowInternetBandwidthRequestFields
}

func GetShowInternetBandwidthRequestFieldsEnum() ShowInternetBandwidthRequestFieldsEnum {
	return ShowInternetBandwidthRequestFieldsEnum{
		ID: ShowInternetBandwidthRequestFields{
			value: "id",
		},
		NAME: ShowInternetBandwidthRequestFields{
			value: "name",
		},
		ISP: ShowInternetBandwidthRequestFields{
			value: "isp",
		},
		INGRESS_SIZE: ShowInternetBandwidthRequestFields{
			value: "ingress_size",
		},
		ACCESS_SITE: ShowInternetBandwidthRequestFields{
			value: "access_site",
		},
		SIZE: ShowInternetBandwidthRequestFields{
			value: "size",
		},
		DESCRIPTION: ShowInternetBandwidthRequestFields{
			value: "description",
		},
		CHARGE_MODE: ShowInternetBandwidthRequestFields{
			value: "charge_mode",
		},
		RATIO_95PEAK: ShowInternetBandwidthRequestFields{
			value: "ratio_95peak",
		},
		FREEZEN_INFO: ShowInternetBandwidthRequestFields{
			value: "freezen_info",
		},
		DOMAIN_ID: ShowInternetBandwidthRequestFields{
			value: "domain_id",
		},
		STATUS: ShowInternetBandwidthRequestFields{
			value: "status",
		},
		CREATED_AT: ShowInternetBandwidthRequestFields{
			value: "created_at",
		},
		UPDATED_AT: ShowInternetBandwidthRequestFields{
			value: "updated_at",
		},
		IS_PRE_PAID: ShowInternetBandwidthRequestFields{
			value: "is_pre_paid",
		},
		TYPE: ShowInternetBandwidthRequestFields{
			value: "type",
		},
		TAGS: ShowInternetBandwidthRequestFields{
			value: "tags",
		},
		SYS_TAGS: ShowInternetBandwidthRequestFields{
			value: "sys_tags",
		},
		ENTERPRISE_PROJECT_ID: ShowInternetBandwidthRequestFields{
			value: "enterprise_project_id",
		},
	}
}

func (c ShowInternetBandwidthRequestFields) Value() string {
	return c.value
}

func (c ShowInternetBandwidthRequestFields) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInternetBandwidthRequestFields) UnmarshalJSON(b []byte) error {
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
