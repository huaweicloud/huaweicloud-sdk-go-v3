package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAvailableBuildDrInstanceRequest Request Object
type ShowAvailableBuildDrInstanceRequest struct {

	// 要查询的实例类型 master：主实例。 slave：灾备实例。
	Type ShowAvailableBuildDrInstanceRequestType `json:"type"`

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ShowAvailableBuildDrInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvailableBuildDrInstanceRequest struct{}"
	}

	return strings.Join([]string{"ShowAvailableBuildDrInstanceRequest", string(data)}, " ")
}

type ShowAvailableBuildDrInstanceRequestType struct {
	value string
}

type ShowAvailableBuildDrInstanceRequestTypeEnum struct {
	MASTER ShowAvailableBuildDrInstanceRequestType
	SLAVE  ShowAvailableBuildDrInstanceRequestType
}

func GetShowAvailableBuildDrInstanceRequestTypeEnum() ShowAvailableBuildDrInstanceRequestTypeEnum {
	return ShowAvailableBuildDrInstanceRequestTypeEnum{
		MASTER: ShowAvailableBuildDrInstanceRequestType{
			value: "master",
		},
		SLAVE: ShowAvailableBuildDrInstanceRequestType{
			value: "slave",
		},
	}
}

func (c ShowAvailableBuildDrInstanceRequestType) Value() string {
	return c.value
}

func (c ShowAvailableBuildDrInstanceRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAvailableBuildDrInstanceRequestType) UnmarshalJSON(b []byte) error {
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
