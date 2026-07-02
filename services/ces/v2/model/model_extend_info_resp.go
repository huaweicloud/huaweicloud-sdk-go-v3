package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExtendInfoResp **参数解释** 看板相关拓展信息
type ExtendInfoResp struct {

	// **参数解释** 表示指标聚合方式 **取值范围** 枚举值： - average 平均值 - min 最小值 - max 最大值 - sum 求和值
	Filter *ExtendInfoRespFilter `json:"filter,omitempty"`

	// **参数解释** 表示指标聚合周期 **取值范围** - 1 原始值 - 60 一分钟 - 300 5分钟 - 1200 20分钟 - 3600 1小时 - 14400 4小时 - 86400 1天
	Period *string `json:"period,omitempty"`

	// **参数解释** 展示时间 **取值范围** 枚举值： - 0 自定义时间 - 5 5分钟 - 15 15分钟 - 30 30分钟 - 60 1小时 - 120 2小时 - 180 3小时 - 720 12小时 - 1440 24小时 - 10080 7天 - 43200 30天
	DisplayTime *ExtendInfoRespDisplayTime `json:"display_time,omitempty"`

	// **参数解释** 刷新时间 **取值范围** 枚举值： - 0 不刷新 - 10 10秒 - 60 1分钟 - 300 5分钟 - 1200 20分钟
	RefreshTime *ExtendInfoRespRefreshTime `json:"refresh_time,omitempty"`

	// **参数解释** 开始时间 **取值范围** 时间戳取值为[0,9999999999999]
	From *int64 `json:"from,omitempty"`

	// **参数解释** 结束时间 **取值范围** 时间戳取值为[0,9999999999999]
	To *int64 `json:"to,omitempty"`

	// **参数解释** 监控大屏背景颜色 **取值范围** 背景颜色长度最长为100个字符
	ScreenColor *string `json:"screen_color,omitempty"`

	// **参数解释** 监控大屏是否自动切换 **取值范围** - true 是 - false 否
	EnableScreenAutoPlay *bool `json:"enable_screen_auto_play,omitempty"`

	// **参数解释**  监控大屏自动切换时间间隔  **取值范围**  枚举值：  - 10000 10秒  - 30000 30秒  - 60000 1分钟
	TimeInterval *ExtendInfoRespTimeInterval `json:"time_interval,omitempty"`

	// **参数解释** 是否开启图例 **取值范围** - true 是 - false 否
	EnableLegend *bool `json:"enable_legend,omitempty"`

	// **参数解释** 大屏展示视图数量 **取值范围** 视图数量为[0,65535]
	FullScreenWidgetNum *int32 `json:"full_screen_widget_num,omitempty"`
}

func (o ExtendInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtendInfoResp struct{}"
	}

	return strings.Join([]string{"ExtendInfoResp", string(data)}, " ")
}

type ExtendInfoRespFilter struct {
	value string
}

type ExtendInfoRespFilterEnum struct {
	AVERAGE ExtendInfoRespFilter
	MIN     ExtendInfoRespFilter
	MAX     ExtendInfoRespFilter
	SUM     ExtendInfoRespFilter
}

func GetExtendInfoRespFilterEnum() ExtendInfoRespFilterEnum {
	return ExtendInfoRespFilterEnum{
		AVERAGE: ExtendInfoRespFilter{
			value: "average",
		},
		MIN: ExtendInfoRespFilter{
			value: "min",
		},
		MAX: ExtendInfoRespFilter{
			value: "max",
		},
		SUM: ExtendInfoRespFilter{
			value: "sum",
		},
	}
}

func (c ExtendInfoRespFilter) Value() string {
	return c.value
}

func (c ExtendInfoRespFilter) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExtendInfoRespFilter) UnmarshalJSON(b []byte) error {
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

type ExtendInfoRespDisplayTime struct {
	value int32
}

type ExtendInfoRespDisplayTimeEnum struct {
	E_0     ExtendInfoRespDisplayTime
	E_5     ExtendInfoRespDisplayTime
	E_15    ExtendInfoRespDisplayTime
	E_30    ExtendInfoRespDisplayTime
	E_60    ExtendInfoRespDisplayTime
	E_120   ExtendInfoRespDisplayTime
	E_180   ExtendInfoRespDisplayTime
	E_720   ExtendInfoRespDisplayTime
	E_1440  ExtendInfoRespDisplayTime
	E_10080 ExtendInfoRespDisplayTime
	E_43200 ExtendInfoRespDisplayTime
}

func GetExtendInfoRespDisplayTimeEnum() ExtendInfoRespDisplayTimeEnum {
	return ExtendInfoRespDisplayTimeEnum{
		E_0: ExtendInfoRespDisplayTime{
			value: 0,
		}, E_5: ExtendInfoRespDisplayTime{
			value: 5,
		}, E_15: ExtendInfoRespDisplayTime{
			value: 15,
		}, E_30: ExtendInfoRespDisplayTime{
			value: 30,
		}, E_60: ExtendInfoRespDisplayTime{
			value: 60,
		}, E_120: ExtendInfoRespDisplayTime{
			value: 120,
		}, E_180: ExtendInfoRespDisplayTime{
			value: 180,
		}, E_720: ExtendInfoRespDisplayTime{
			value: 720,
		}, E_1440: ExtendInfoRespDisplayTime{
			value: 1440,
		}, E_10080: ExtendInfoRespDisplayTime{
			value: 10080,
		}, E_43200: ExtendInfoRespDisplayTime{
			value: 43200,
		},
	}
}

func (c ExtendInfoRespDisplayTime) Value() int32 {
	return c.value
}

func (c ExtendInfoRespDisplayTime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExtendInfoRespDisplayTime) UnmarshalJSON(b []byte) error {
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

type ExtendInfoRespRefreshTime struct {
	value int32
}

type ExtendInfoRespRefreshTimeEnum struct {
	E_0    ExtendInfoRespRefreshTime
	E_10   ExtendInfoRespRefreshTime
	E_60   ExtendInfoRespRefreshTime
	E_300  ExtendInfoRespRefreshTime
	E_1200 ExtendInfoRespRefreshTime
}

func GetExtendInfoRespRefreshTimeEnum() ExtendInfoRespRefreshTimeEnum {
	return ExtendInfoRespRefreshTimeEnum{
		E_0: ExtendInfoRespRefreshTime{
			value: 0,
		}, E_10: ExtendInfoRespRefreshTime{
			value: 10,
		}, E_60: ExtendInfoRespRefreshTime{
			value: 60,
		}, E_300: ExtendInfoRespRefreshTime{
			value: 300,
		}, E_1200: ExtendInfoRespRefreshTime{
			value: 1200,
		},
	}
}

func (c ExtendInfoRespRefreshTime) Value() int32 {
	return c.value
}

func (c ExtendInfoRespRefreshTime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExtendInfoRespRefreshTime) UnmarshalJSON(b []byte) error {
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

type ExtendInfoRespTimeInterval struct {
	value int32
}

type ExtendInfoRespTimeIntervalEnum struct {
	E_10000 ExtendInfoRespTimeInterval
	E_30000 ExtendInfoRespTimeInterval
	E_60000 ExtendInfoRespTimeInterval
}

func GetExtendInfoRespTimeIntervalEnum() ExtendInfoRespTimeIntervalEnum {
	return ExtendInfoRespTimeIntervalEnum{
		E_10000: ExtendInfoRespTimeInterval{
			value: 10000,
		}, E_30000: ExtendInfoRespTimeInterval{
			value: 30000,
		}, E_60000: ExtendInfoRespTimeInterval{
			value: 60000,
		},
	}
}

func (c ExtendInfoRespTimeInterval) Value() int32 {
	return c.value
}

func (c ExtendInfoRespTimeInterval) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExtendInfoRespTimeInterval) UnmarshalJSON(b []byte) error {
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
