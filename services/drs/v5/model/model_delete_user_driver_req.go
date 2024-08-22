package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DeleteUserDriverReq struct {

	// 指定待删除的驱动文件类型。取值范围： - db2：DB2 for LUW - informix：Informix
	DriverType DeleteUserDriverReqDriverType `json:"driver_type"`

	// JDBC驱动文件列表，列表长度1-20，driver_name的长度5-64，结尾以.jar结尾。
	DriverNames []string `json:"driver_names"`
}

func (o DeleteUserDriverReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteUserDriverReq struct{}"
	}

	return strings.Join([]string{"DeleteUserDriverReq", string(data)}, " ")
}

type DeleteUserDriverReqDriverType struct {
	value string
}

type DeleteUserDriverReqDriverTypeEnum struct {
	DB2      DeleteUserDriverReqDriverType
	INFORMIX DeleteUserDriverReqDriverType
}

func GetDeleteUserDriverReqDriverTypeEnum() DeleteUserDriverReqDriverTypeEnum {
	return DeleteUserDriverReqDriverTypeEnum{
		DB2: DeleteUserDriverReqDriverType{
			value: "db2",
		},
		INFORMIX: DeleteUserDriverReqDriverType{
			value: "informix",
		},
	}
}

func (c DeleteUserDriverReqDriverType) Value() string {
	return c.value
}

func (c DeleteUserDriverReqDriverType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteUserDriverReqDriverType) UnmarshalJSON(b []byte) error {
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
