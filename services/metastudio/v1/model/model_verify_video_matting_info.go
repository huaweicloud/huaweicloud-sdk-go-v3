package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// VerifyVideoMattingInfo 自动扣绿时，不抠图的区域。
type VerifyVideoMattingInfo struct {

	// 区域左上角像素点的X轴位置值。
	Dx *int32 `json:"dx,omitempty"`

	// 区域左上角像素点的Y轴位置值。
	Dy *int32 `json:"dy,omitempty"`

	// 区域宽度像素值。
	Width *int32 `json:"width,omitempty"`

	// 区域高度像素值。
	Height *int32 `json:"height,omitempty"`

	// 资源操作类型。 * RESERVED: 保留区域 * DELETE：删除区域 * NO_DEGREEN：无区域
	Method *VerifyVideoMattingInfoMethod `json:"method,omitempty"`
}

func (o VerifyVideoMattingInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VerifyVideoMattingInfo struct{}"
	}

	return strings.Join([]string{"VerifyVideoMattingInfo", string(data)}, " ")
}

type VerifyVideoMattingInfoMethod struct {
	value string
}

type VerifyVideoMattingInfoMethodEnum struct {
	RESERVED   VerifyVideoMattingInfoMethod
	DELETE     VerifyVideoMattingInfoMethod
	NO_DEGREEN VerifyVideoMattingInfoMethod
}

func GetVerifyVideoMattingInfoMethodEnum() VerifyVideoMattingInfoMethodEnum {
	return VerifyVideoMattingInfoMethodEnum{
		RESERVED: VerifyVideoMattingInfoMethod{
			value: "RESERVED",
		},
		DELETE: VerifyVideoMattingInfoMethod{
			value: "DELETE",
		},
		NO_DEGREEN: VerifyVideoMattingInfoMethod{
			value: "NO_DEGREEN",
		},
	}
}

func (c VerifyVideoMattingInfoMethod) Value() string {
	return c.value
}

func (c VerifyVideoMattingInfoMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VerifyVideoMattingInfoMethod) UnmarshalJSON(b []byte) error {
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
