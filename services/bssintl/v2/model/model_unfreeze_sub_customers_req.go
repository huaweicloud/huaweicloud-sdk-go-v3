package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UnfreezeSubCustomersReq struct {

	// 需要解冻的客户账号ID列表。 您可以调用查询客户列表接口获取customer_id。
	CustomerIds []string `json:"customer_ids"`

	// 解冻原因。
	Reason string `json:"reason"`

	// 云经销商ID。获取方法请参见查询云经销商列表。如果需要查询云经销商的子客户列表，必须携带该字段。除此之外，此参数不做处理。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`

	// |参数名称：解冻类型| |参数的约束及描述：该参数非必填，解冻类型，支持枚举| |ACCOUNT：解冻账号，ACCOUNT_AND_RESOURCE：解冻账号与资源|
	UnfreezeType *UnfreezeSubCustomersReqUnfreezeType `json:"unfreeze_type,omitempty"`
}

func (o UnfreezeSubCustomersReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnfreezeSubCustomersReq struct{}"
	}

	return strings.Join([]string{"UnfreezeSubCustomersReq", string(data)}, " ")
}

type UnfreezeSubCustomersReqUnfreezeType struct {
	value string
}

type UnfreezeSubCustomersReqUnfreezeTypeEnum struct {
	ACCOUNT              UnfreezeSubCustomersReqUnfreezeType
	ACCOUNT_AND_RESOURCE UnfreezeSubCustomersReqUnfreezeType
}

func GetUnfreezeSubCustomersReqUnfreezeTypeEnum() UnfreezeSubCustomersReqUnfreezeTypeEnum {
	return UnfreezeSubCustomersReqUnfreezeTypeEnum{
		ACCOUNT: UnfreezeSubCustomersReqUnfreezeType{
			value: "ACCOUNT",
		},
		ACCOUNT_AND_RESOURCE: UnfreezeSubCustomersReqUnfreezeType{
			value: "ACCOUNT_AND_RESOURCE",
		},
	}
}

func (c UnfreezeSubCustomersReqUnfreezeType) Value() string {
	return c.value
}

func (c UnfreezeSubCustomersReqUnfreezeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UnfreezeSubCustomersReqUnfreezeType) UnmarshalJSON(b []byte) error {
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
