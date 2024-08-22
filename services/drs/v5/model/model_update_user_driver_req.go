package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateUserDriverReq struct {

	// JDBC驱动文件名称，name的长度5-64，结尾以.jar结尾。
	DriverName string `json:"driver_name"`

	// 指定待同步的驱动文件类型。取值范围： - db2：DB2 for LUW - informix：Informix
	DriverType UpdateUserDriverReqDriverType `json:"driver_type"`
}

func (o UpdateUserDriverReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserDriverReq struct{}"
	}

	return strings.Join([]string{"UpdateUserDriverReq", string(data)}, " ")
}

type UpdateUserDriverReqDriverType struct {
	value string
}

type UpdateUserDriverReqDriverTypeEnum struct {
	DB2      UpdateUserDriverReqDriverType
	INFORMIX UpdateUserDriverReqDriverType
}

func GetUpdateUserDriverReqDriverTypeEnum() UpdateUserDriverReqDriverTypeEnum {
	return UpdateUserDriverReqDriverTypeEnum{
		DB2: UpdateUserDriverReqDriverType{
			value: "db2",
		},
		INFORMIX: UpdateUserDriverReqDriverType{
			value: "informix",
		},
	}
}

func (c UpdateUserDriverReqDriverType) Value() string {
	return c.value
}

func (c UpdateUserDriverReqDriverType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateUserDriverReqDriverType) UnmarshalJSON(b []byte) error {
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
