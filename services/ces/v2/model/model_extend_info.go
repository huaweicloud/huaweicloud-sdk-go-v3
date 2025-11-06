package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExtendInfo 看板相关拓展信息
type ExtendInfo struct {

	// 表示指标聚合方式，average表示平均值，min表示最小值，max表示最大值，sum表示求合
	Filter *ExtendInfoFilter `json:"filter,omitempty"`

	// '表示指标聚合周期，{1:表示原始值，60:表示一分钟，300:表示5分钟，1200:表示20分钟，3600:表示1小时，14400:表示4小时，86400:表示1天}''
	Period *string `json:"period,omitempty"`

	// 展示时间，0表示使用自定义时间展示， 5分钟，15分钟，30分钟，1小时，2小时，3小时，12小时，24小时，7天，30天
	DisplayTime *ExtendInfoDisplayTime `json:"display_time,omitempty"`

	// 刷新时间 0秒表示不刷新,10秒，1分钟，5分钟，20分钟
	RefreshTime *ExtendInfoRefreshTime `json:"refresh_time,omitempty"`

	// 开始时间
	From *int64 `json:"from,omitempty"`

	// 结束时间
	To *int64 `json:"to,omitempty"`

	// 监控大屏背景颜色
	ScreenColor *string `json:"screen_color,omitempty"`

	// 监控大屏是否自动切换
	EnableScreenAutoPlay *bool `json:"enable_screen_auto_play,omitempty"`

	// 监控大屏自动切换时间间隔，10000代表10s，30000代表30s，60000代表1min
	TimeInterval *ExtendInfoTimeInterval `json:"time_interval,omitempty"`

	// 是否开启图例
	EnableLegend *bool `json:"enable_legend,omitempty"`

	// 大屏展示视图数量, 可以取的值必须与console页面可选值保持一致
	FullScreenWidgetNum *int32 `json:"full_screen_widget_num,omitempty"`
}

func (o ExtendInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtendInfo struct{}"
	}

	return strings.Join([]string{"ExtendInfo", string(data)}, " ")
}

type ExtendInfoFilter struct {
	value string
}

type ExtendInfoFilterEnum struct {
	AVERAGE ExtendInfoFilter
	MIN     ExtendInfoFilter
	MAX     ExtendInfoFilter
	SUM     ExtendInfoFilter
}

func GetExtendInfoFilterEnum() ExtendInfoFilterEnum {
	return ExtendInfoFilterEnum{
		AVERAGE: ExtendInfoFilter{
			value: "average",
		},
		MIN: ExtendInfoFilter{
			value: "min",
		},
		MAX: ExtendInfoFilter{
			value: "max",
		},
		SUM: ExtendInfoFilter{
			value: "sum",
		},
	}
}

func (c ExtendInfoFilter) Value() string {
	return c.value
}

func (c ExtendInfoFilter) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExtendInfoFilter) UnmarshalJSON(b []byte) error {
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

type ExtendInfoDisplayTime struct {
	value int32
}

type ExtendInfoDisplayTimeEnum struct {
	E_0     ExtendInfoDisplayTime
	E_5     ExtendInfoDisplayTime
	E_15    ExtendInfoDisplayTime
	E_30    ExtendInfoDisplayTime
	E_60    ExtendInfoDisplayTime
	E_120   ExtendInfoDisplayTime
	E_180   ExtendInfoDisplayTime
	E_720   ExtendInfoDisplayTime
	E_1440  ExtendInfoDisplayTime
	E_10080 ExtendInfoDisplayTime
	E_43200 ExtendInfoDisplayTime
}

func GetExtendInfoDisplayTimeEnum() ExtendInfoDisplayTimeEnum {
	return ExtendInfoDisplayTimeEnum{
		E_0: ExtendInfoDisplayTime{
			value: 0,
		}, E_5: ExtendInfoDisplayTime{
			value: 5,
		}, E_15: ExtendInfoDisplayTime{
			value: 15,
		}, E_30: ExtendInfoDisplayTime{
			value: 30,
		}, E_60: ExtendInfoDisplayTime{
			value: 60,
		}, E_120: ExtendInfoDisplayTime{
			value: 120,
		}, E_180: ExtendInfoDisplayTime{
			value: 180,
		}, E_720: ExtendInfoDisplayTime{
			value: 720,
		}, E_1440: ExtendInfoDisplayTime{
			value: 1440,
		}, E_10080: ExtendInfoDisplayTime{
			value: 10080,
		}, E_43200: ExtendInfoDisplayTime{
			value: 43200,
		},
	}
}

func (c ExtendInfoDisplayTime) Value() int32 {
	return c.value
}

func (c ExtendInfoDisplayTime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExtendInfoDisplayTime) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type ExtendInfoRefreshTime struct {
	value int32
}

type ExtendInfoRefreshTimeEnum struct {
	E_0    ExtendInfoRefreshTime
	E_10   ExtendInfoRefreshTime
	E_60   ExtendInfoRefreshTime
	E_300  ExtendInfoRefreshTime
	E_1200 ExtendInfoRefreshTime
}

func GetExtendInfoRefreshTimeEnum() ExtendInfoRefreshTimeEnum {
	return ExtendInfoRefreshTimeEnum{
		E_0: ExtendInfoRefreshTime{
			value: 0,
		}, E_10: ExtendInfoRefreshTime{
			value: 10,
		}, E_60: ExtendInfoRefreshTime{
			value: 60,
		}, E_300: ExtendInfoRefreshTime{
			value: 300,
		}, E_1200: ExtendInfoRefreshTime{
			value: 1200,
		},
	}
}

func (c ExtendInfoRefreshTime) Value() int32 {
	return c.value
}

func (c ExtendInfoRefreshTime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExtendInfoRefreshTime) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type ExtendInfoTimeInterval struct {
	value int32
}

type ExtendInfoTimeIntervalEnum struct {
	E_10000 ExtendInfoTimeInterval
	E_30000 ExtendInfoTimeInterval
	E_60000 ExtendInfoTimeInterval
}

func GetExtendInfoTimeIntervalEnum() ExtendInfoTimeIntervalEnum {
	return ExtendInfoTimeIntervalEnum{
		E_10000: ExtendInfoTimeInterval{
			value: 10000,
		}, E_30000: ExtendInfoTimeInterval{
			value: 30000,
		}, E_60000: ExtendInfoTimeInterval{
			value: 60000,
		},
	}
}

func (c ExtendInfoTimeInterval) Value() int32 {
	return c.value
}

func (c ExtendInfoTimeInterval) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExtendInfoTimeInterval) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
