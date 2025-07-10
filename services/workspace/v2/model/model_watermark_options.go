package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type WatermarkOptions struct {

	// 展示方式。取值为：FIXED：固定位置。RANDOM：随机运动。
	ShowStyle *WatermarkOptionsShowStyle `json:"show_style,omitempty"`

	// 颜色。格式：RRGGBB。默认：2a2a2a。
	Color *string `json:"color,omitempty"`

	// 字体大小。取值范围为[1-200]。默认：30。
	FontSize *int32 `json:"font_size,omitempty"`

	// 不透明度（%）。取值范围为[0-100]。默认：\"12.5\"。
	OpacitySetting *string `json:"opacity_setting,omitempty"`

	// 条目数量。取值范围为[1-30]。默认：1。
	ContentItemCount *int32 `json:"content_item_count,omitempty"`

	// 水印内容显示格式。
	DisplayFormat *string `json:"display_format,omitempty"`

	// 倾斜度。取值范围为[-90-90]。默认：-45。
	Lean *int32 `json:"lean,omitempty"`

	// 自定义内容。内容不得带有类似'>'的特殊字符。
	ContentCustomized *string `json:"content_customized,omitempty"`

	// 内容颜色。
	ContentColor *string `json:"content_color,omitempty"`

	// 安全优先开关。
	WatermarkSecurityAccess *bool `json:"watermark_security_access,omitempty"`

	// 用户扩展信息开关。false：表示关闭。true：表示开启。
	UserExtendInfoSwitch *bool `json:"user_extend_info_switch,omitempty"`

	// 用户扩展信息。
	UserExtendInfo *string `json:"user_extend_info,omitempty"`
}

func (o WatermarkOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WatermarkOptions struct{}"
	}

	return strings.Join([]string{"WatermarkOptions", string(data)}, " ")
}

type WatermarkOptionsShowStyle struct {
	value string
}

type WatermarkOptionsShowStyleEnum struct {
	FIXED  WatermarkOptionsShowStyle
	RANDOM WatermarkOptionsShowStyle
}

func GetWatermarkOptionsShowStyleEnum() WatermarkOptionsShowStyleEnum {
	return WatermarkOptionsShowStyleEnum{
		FIXED: WatermarkOptionsShowStyle{
			value: "FIXED",
		},
		RANDOM: WatermarkOptionsShowStyle{
			value: "RANDOM",
		},
	}
}

func (c WatermarkOptionsShowStyle) Value() string {
	return c.value
}

func (c WatermarkOptionsShowStyle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WatermarkOptionsShowStyle) UnmarshalJSON(b []byte) error {
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
