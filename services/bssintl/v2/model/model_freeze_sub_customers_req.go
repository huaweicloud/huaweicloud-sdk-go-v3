package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type FreezeSubCustomersReq struct {

	// 需要冻结的客户账号ID列表。 您可以调用查询客户列表接口获取customer_id。
	CustomerIds []string `json:"customer_ids"`

	// 冻结原因。
	Reason string `json:"reason"`

	// 云经销商ID。获取方法请参见查询云经销商列表。如果需要查询云经销商的子客户列表，必须携带该字段。除此之外，此参数不做处理。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`

	// |参数名称：冻结类型| |参数的约束及描述：该参数非必填，冻结类型，支持枚举| |ACCOUNT：冻结账号，ACCOUNT_AND_RESOURCE：冻结账号与资源|
	FreezeType *FreezeSubCustomersReqFreezeType `json:"freeze_type,omitempty"`
}

func (o FreezeSubCustomersReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FreezeSubCustomersReq struct{}"
	}

	return strings.Join([]string{"FreezeSubCustomersReq", string(data)}, " ")
}

type FreezeSubCustomersReqFreezeType struct {
	value string
}

type FreezeSubCustomersReqFreezeTypeEnum struct {
	ACCOUNT              FreezeSubCustomersReqFreezeType
	ACCOUNT_AND_RESOURCE FreezeSubCustomersReqFreezeType
}

func GetFreezeSubCustomersReqFreezeTypeEnum() FreezeSubCustomersReqFreezeTypeEnum {
	return FreezeSubCustomersReqFreezeTypeEnum{
		ACCOUNT: FreezeSubCustomersReqFreezeType{
			value: "ACCOUNT",
		},
		ACCOUNT_AND_RESOURCE: FreezeSubCustomersReqFreezeType{
			value: "ACCOUNT_AND_RESOURCE",
		},
	}
}

func (c FreezeSubCustomersReqFreezeType) Value() string {
	return c.value
}

func (c FreezeSubCustomersReqFreezeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FreezeSubCustomersReqFreezeType) UnmarshalJSON(b []byte) error {
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
