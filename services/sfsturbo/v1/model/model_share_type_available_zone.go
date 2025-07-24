package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ShareTypeAvailableZone struct {

	// 可用区
	AvailableZone *string `json:"available_zone,omitempty"`

	// 在可用区的状态
	Status *ShareTypeAvailableZoneStatus `json:"status,omitempty"`
}

func (o ShareTypeAvailableZone) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShareTypeAvailableZone struct{}"
	}

	return strings.Join([]string{"ShareTypeAvailableZone", string(data)}, " ")
}

type ShareTypeAvailableZoneStatus struct {
	value string
}

type ShareTypeAvailableZoneStatusEnum struct {
	NORMAL  ShareTypeAvailableZoneStatus
	SELLOUT ShareTypeAvailableZoneStatus
}

func GetShareTypeAvailableZoneStatusEnum() ShareTypeAvailableZoneStatusEnum {
	return ShareTypeAvailableZoneStatusEnum{
		NORMAL: ShareTypeAvailableZoneStatus{
			value: "normal",
		},
		SELLOUT: ShareTypeAvailableZoneStatus{
			value: "sellout",
		},
	}
}

func (c ShareTypeAvailableZoneStatus) Value() string {
	return c.value
}

func (c ShareTypeAvailableZoneStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShareTypeAvailableZoneStatus) UnmarshalJSON(b []byte) error {
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
