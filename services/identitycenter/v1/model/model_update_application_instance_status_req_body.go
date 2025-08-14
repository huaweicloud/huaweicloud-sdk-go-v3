package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateApplicationInstanceStatusReqBody UpdateApplicationInstanceStatus的请求体
type UpdateApplicationInstanceStatusReqBody struct {

	// 应用程序状态
	Status UpdateApplicationInstanceStatusReqBodyStatus `json:"status"`
}

func (o UpdateApplicationInstanceStatusReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApplicationInstanceStatusReqBody struct{}"
	}

	return strings.Join([]string{"UpdateApplicationInstanceStatusReqBody", string(data)}, " ")
}

type UpdateApplicationInstanceStatusReqBodyStatus struct {
	value string
}

type UpdateApplicationInstanceStatusReqBodyStatusEnum struct {
	ENABLED UpdateApplicationInstanceStatusReqBodyStatus
}

func GetUpdateApplicationInstanceStatusReqBodyStatusEnum() UpdateApplicationInstanceStatusReqBodyStatusEnum {
	return UpdateApplicationInstanceStatusReqBodyStatusEnum{
		ENABLED: UpdateApplicationInstanceStatusReqBodyStatus{
			value: "ENABLED",
		},
	}
}

func (c UpdateApplicationInstanceStatusReqBodyStatus) Value() string {
	return c.value
}

func (c UpdateApplicationInstanceStatusReqBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateApplicationInstanceStatusReqBodyStatus) UnmarshalJSON(b []byte) error {
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
