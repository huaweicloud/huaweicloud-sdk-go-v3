package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateProjectOption
type UpdateProjectOption struct {

	// 项目的状态信息，参数的值为\"suspended\"或\"normal\"。 - status值为\"suspended\"时，会将项目设置为冻结状态。 - status值为\"normal\"时，会将项目设置为正常（解冻）状态。
	Status UpdateProjectOptionStatus `json:"status"`
}

func (o UpdateProjectOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProjectOption struct{}"
	}

	return strings.Join([]string{"UpdateProjectOption", string(data)}, " ")
}

type UpdateProjectOptionStatus struct {
	value string
}

type UpdateProjectOptionStatusEnum struct {
	SUSPENDED UpdateProjectOptionStatus
	NORMAL    UpdateProjectOptionStatus
}

func GetUpdateProjectOptionStatusEnum() UpdateProjectOptionStatusEnum {
	return UpdateProjectOptionStatusEnum{
		SUSPENDED: UpdateProjectOptionStatus{
			value: "suspended",
		},
		NORMAL: UpdateProjectOptionStatus{
			value: "normal",
		},
	}
}

func (c UpdateProjectOptionStatus) Value() string {
	return c.value
}

func (c UpdateProjectOptionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateProjectOptionStatus) UnmarshalJSON(b []byte) error {
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
