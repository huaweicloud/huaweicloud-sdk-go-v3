package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// VpcEgressKindObj API类型，固定值“VpcEgress”，该值不可修改。
type VpcEgressKindObj struct {
	value string
}

type VpcEgressKindObjEnum struct {
	VPC_EGRESS VpcEgressKindObj
}

func GetVpcEgressKindObjEnum() VpcEgressKindObjEnum {
	return VpcEgressKindObjEnum{
		VPC_EGRESS: VpcEgressKindObj{
			value: "VpcEgress",
		},
	}
}

func (c VpcEgressKindObj) Value() string {
	return c.value
}

func (c VpcEgressKindObj) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VpcEgressKindObj) UnmarshalJSON(b []byte) error {
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
