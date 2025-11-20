package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BillResourceType 资源类型。 当前支持的形象资源类型如下： * 2D_DIGITAL_HUMAN_BASIC：形象制作基础版 * 2D_DIGITAL_HUMAN_ADVANCED：形象制作高级版 * 2D_DIGITAL_HUMAN_FLEXUS：形象制作FLEXUS版  当前支持的声音资源类型如下： * VOICE_BASIC: 声音制作基础版 * VOICE_MIDDLE: 声音制作进阶版 * VOICE_ADVANCE：声音制作高级版 * VOICE_THIRD_PARTY：声音制作第三方出门问问 * VOICE_THIRD_PARTY_LJZN: 声音制作第三方逻辑智能 * VOICE_FLEXUS: 声音制作Flexus版资源
type BillResourceType struct {
	value string
}

type BillResourceTypeEnum struct {
	E_2_D_DIGITAL_HUMAN_BASIC    BillResourceType
	E_2_D_DIGITAL_HUMAN_ADVANCED BillResourceType
	E_2_D_DIGITAL_HUMAN_FLEXUS   BillResourceType
	VOICE_BASIC                  BillResourceType
	VOICE_MIDDLE                 BillResourceType
	VOICE_ADVANCE                BillResourceType
	VOICE_THIRD_PARTY            BillResourceType
	VOICE_THIRD_PARTY_LJZN       BillResourceType
	VOICE_FLEXUS                 BillResourceType
}

func GetBillResourceTypeEnum() BillResourceTypeEnum {
	return BillResourceTypeEnum{
		E_2_D_DIGITAL_HUMAN_BASIC: BillResourceType{
			value: "2D_DIGITAL_HUMAN_BASIC",
		},
		E_2_D_DIGITAL_HUMAN_ADVANCED: BillResourceType{
			value: "2D_DIGITAL_HUMAN_ADVANCED",
		},
		E_2_D_DIGITAL_HUMAN_FLEXUS: BillResourceType{
			value: "2D_DIGITAL_HUMAN_FLEXUS",
		},
		VOICE_BASIC: BillResourceType{
			value: "VOICE_BASIC",
		},
		VOICE_MIDDLE: BillResourceType{
			value: "VOICE_MIDDLE",
		},
		VOICE_ADVANCE: BillResourceType{
			value: "VOICE_ADVANCE",
		},
		VOICE_THIRD_PARTY: BillResourceType{
			value: "VOICE_THIRD_PARTY",
		},
		VOICE_THIRD_PARTY_LJZN: BillResourceType{
			value: "VOICE_THIRD_PARTY_LJZN",
		},
		VOICE_FLEXUS: BillResourceType{
			value: "VOICE_FLEXUS",
		},
	}
}

func (c BillResourceType) Value() string {
	return c.value
}

func (c BillResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillResourceType) UnmarshalJSON(b []byte) error {
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
