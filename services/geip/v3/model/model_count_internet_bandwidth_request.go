package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CountInternetBandwidthRequest Request Object
type CountInternetBandwidthRequest struct {
	Id *[]string `json:"id,omitempty"`

	Size *[]int32 `json:"size,omitempty"`

	Name *[]string `json:"name,omitempty"`

	NameLike *string `json:"name_like,omitempty"`

	AccessSite *[]string `json:"access_site,omitempty"`

	Status *[]CountInternetBandwidthRequestStatus `json:"status,omitempty"`

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	Tags *[]string `json:"tags,omitempty"`
}

func (o CountInternetBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountInternetBandwidthRequest struct{}"
	}

	return strings.Join([]string{"CountInternetBandwidthRequest", string(data)}, " ")
}

type CountInternetBandwidthRequestStatus struct {
	value string
}

type CountInternetBandwidthRequestStatusEnum struct {
	FREEZED CountInternetBandwidthRequestStatus
	NORMAL  CountInternetBandwidthRequestStatus
}

func GetCountInternetBandwidthRequestStatusEnum() CountInternetBandwidthRequestStatusEnum {
	return CountInternetBandwidthRequestStatusEnum{
		FREEZED: CountInternetBandwidthRequestStatus{
			value: "freezed",
		},
		NORMAL: CountInternetBandwidthRequestStatus{
			value: "normal",
		},
	}
}

func (c CountInternetBandwidthRequestStatus) Value() string {
	return c.value
}

func (c CountInternetBandwidthRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CountInternetBandwidthRequestStatus) UnmarshalJSON(b []byte) error {
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
