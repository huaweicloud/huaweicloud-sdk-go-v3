package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateAccessPolicyReq struct {

	// 操作类别。 * ADD_IP： 添加IP * DELETE_IP： 删除IP
	OperationType *UpdateAccessPolicyReqOperationType `json:"operation_type,omitempty"`

	// 策略的ip列表。
	IpWhiteList *[]IpInfo `json:"ip_white_list,omitempty"`
}

func (o UpdateAccessPolicyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccessPolicyReq struct{}"
	}

	return strings.Join([]string{"UpdateAccessPolicyReq", string(data)}, " ")
}

type UpdateAccessPolicyReqOperationType struct {
	value string
}

type UpdateAccessPolicyReqOperationTypeEnum struct {
	ADD_IP    UpdateAccessPolicyReqOperationType
	DELETE_IP UpdateAccessPolicyReqOperationType
}

func GetUpdateAccessPolicyReqOperationTypeEnum() UpdateAccessPolicyReqOperationTypeEnum {
	return UpdateAccessPolicyReqOperationTypeEnum{
		ADD_IP: UpdateAccessPolicyReqOperationType{
			value: "ADD_IP",
		},
		DELETE_IP: UpdateAccessPolicyReqOperationType{
			value: "DELETE_IP",
		},
	}
}

func (c UpdateAccessPolicyReqOperationType) Value() string {
	return c.value
}

func (c UpdateAccessPolicyReqOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAccessPolicyReqOperationType) UnmarshalJSON(b []byte) error {
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
