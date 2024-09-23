package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ImageAssetMeta 图片元数据
type ImageAssetMeta struct {

	// **参数解释**： 图片编码格式。 **约束限制**： 用户无需填写，系统自行提取。 **取值范围**： 字符长度0-32位。 **默认取值**： 不涉及。
	Codec *string `json:"codec,omitempty"`

	// **参数解释**： 图片宽度。 **约束限制**： 用户无需填写，系统自行提取。 **默认取值**： 不涉及。
	Width *int32 `json:"width,omitempty"`

	// **参数解释**： 图片高度。 **约束限制**： 用户无需填写，系统自行提取。 **默认取值**： 不涉及。
	Height *int32 `json:"height,omitempty"`

	// **参数解释**： 图片大小。 **约束限制**： 用户无需填写，系统自行提取。 **默认取值**： 不涉及。
	Size float32 `json:"size,omitempty"`

	// **参数解释**： 图片形态。 **约束限制**： 用户无需填写，系统自行提取。 **取值范围**： * Horizontal：横向 * Vertical：纵向  **默认取值**： 不涉及。
	Mode *ImageAssetMetaMode `json:"mode,omitempty"`

	ErrorInfo *ErrorResponse `json:"error_info,omitempty"`
}

func (o ImageAssetMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageAssetMeta struct{}"
	}

	return strings.Join([]string{"ImageAssetMeta", string(data)}, " ")
}

type ImageAssetMetaMode struct {
	value string
}

type ImageAssetMetaModeEnum struct {
	HORIZONTAL ImageAssetMetaMode
	VERTICAL   ImageAssetMetaMode
}

func GetImageAssetMetaModeEnum() ImageAssetMetaModeEnum {
	return ImageAssetMetaModeEnum{
		HORIZONTAL: ImageAssetMetaMode{
			value: "Horizontal",
		},
		VERTICAL: ImageAssetMetaMode{
			value: "Vertical",
		},
	}
}

func (c ImageAssetMetaMode) Value() string {
	return c.value
}

func (c ImageAssetMetaMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageAssetMetaMode) UnmarshalJSON(b []byte) error {
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
