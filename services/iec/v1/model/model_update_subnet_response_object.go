package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 更新子网响应对象
type UpdateSubnetResponseObject struct {
	// 子网ID

	Id *string `json:"id,omitempty"`
	// 子网的状态  取值范围： - ACTIVE：表示子网已挂载到ROUTER上 - UNKNOWN：表示子网还未挂载到ROUTER上 - ERROR：表示子网状态故障

	Status *UpdateSubnetResponseObjectStatus `json:"status,omitempty"`
}

func (o UpdateSubnetResponseObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubnetResponseObject struct{}"
	}

	return strings.Join([]string{"UpdateSubnetResponseObject", string(data)}, " ")
}

type UpdateSubnetResponseObjectStatus struct {
	value string
}

type UpdateSubnetResponseObjectStatusEnum struct {
	ACTIVE  UpdateSubnetResponseObjectStatus
	UNKNOWN UpdateSubnetResponseObjectStatus
	ERROR_  UpdateSubnetResponseObjectStatus
}

func GetUpdateSubnetResponseObjectStatusEnum() UpdateSubnetResponseObjectStatusEnum {
	return UpdateSubnetResponseObjectStatusEnum{
		ACTIVE: UpdateSubnetResponseObjectStatus{
			value: "ACTIVE",
		},
		UNKNOWN: UpdateSubnetResponseObjectStatus{
			value: "UNKNOWN",
		},
		ERROR_: UpdateSubnetResponseObjectStatus{
			value: "ERROR  ",
		},
	}
}

func (c UpdateSubnetResponseObjectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSubnetResponseObjectStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
