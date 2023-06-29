package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type OperateUserReq struct {

	// 操作类型，可选值为： - LOCK：锁定用户。 - UNLOCK：解锁用户。 - RESET_PWD：重置用户密码。
	OpType OperateUserReqOpType `json:"op_type"`
}

func (o OperateUserReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateUserReq struct{}"
	}

	return strings.Join([]string{"OperateUserReq", string(data)}, " ")
}

type OperateUserReqOpType struct {
	value string
}

type OperateUserReqOpTypeEnum struct {
	LOCK      OperateUserReqOpType
	UNLOCK    OperateUserReqOpType
	RESET_PWD OperateUserReqOpType
}

func GetOperateUserReqOpTypeEnum() OperateUserReqOpTypeEnum {
	return OperateUserReqOpTypeEnum{
		LOCK: OperateUserReqOpType{
			value: "LOCK",
		},
		UNLOCK: OperateUserReqOpType{
			value: "UNLOCK",
		},
		RESET_PWD: OperateUserReqOpType{
			value: "RESET_PWD",
		},
	}
}

func (c OperateUserReqOpType) Value() string {
	return c.value
}

func (c OperateUserReqOpType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OperateUserReqOpType) UnmarshalJSON(b []byte) error {
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
