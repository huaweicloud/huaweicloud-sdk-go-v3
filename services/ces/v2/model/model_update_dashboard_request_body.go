package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateDashboardRequestBody struct {

	// 自定义监控看板名称
	DashboardName *string `json:"dashboard_name,omitempty"`

	// 监控看板是否标记收藏, true: 收藏, false: 未收藏
	IsFavorite *bool `json:"is_favorite,omitempty"`

	// 监控视图展示模式，0表示自定义坐标，1代表每行一个
	RowWidgetNum *int32 `json:"row_widget_num,omitempty"`

	// 表示指标聚合方式，average表示平均值，min表示最小值，max表示最大值，sum表示求合
	Filter *UpdateDashboardRequestBodyFilter `json:"filter,omitempty"`

	// 表示指标聚合周期，1表示原始值，60表示一分钟，300表示5分钟，1200表示20分钟，3600表示1小时，14400表示4小时，86400表示1天
	Period *UpdateDashboardRequestBodyPeriod `json:"period,omitempty"`

	// 展示时间，0表示使用自定义时间展示， 5分钟，15分钟，30分钟，1小时，2小时，3小时，12小时，24小时，7天，30天
	DisplayTime *UpdateDashboardRequestBodyDisplayTime `json:"display_time,omitempty"`

	// 刷新时间 0秒表示不刷新,10秒，1分钟，5分钟，20分钟
	RefreshTime *UpdateDashboardRequestBodyRefreshTime `json:"refresh_time,omitempty"`

	// 开始时间
	From *int32 `json:"from,omitempty"`

	// 结束时间
	To *int32 `json:"to,omitempty"`

	// 监控大屏背景颜色
	ScreenColor *string `json:"screen_color,omitempty"`

	// 监控大屏是否自动切换
	EnableScreenAutoPlay *bool `json:"enable_screen_auto_play,omitempty"`

	// 监控大屏自动切换时间间隔，10000代表10s，30000代表30s，60000代表1min
	TimeInterval *UpdateDashboardRequestBodyTimeInterval `json:"time_interval,omitempty"`

	// 是否开启图例
	EnableLegend *bool `json:"enable_legend,omitempty"`

	// 大屏展示视图数量, 可以取得值必须与console页面可选值保持一致
	FullScreenWidgetNum *int32 `json:"full_screen_widget_num,omitempty"`
}

func (o UpdateDashboardRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDashboardRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDashboardRequestBody", string(data)}, " ")
}

type UpdateDashboardRequestBodyFilter struct {
	value string
}

type UpdateDashboardRequestBodyFilterEnum struct {
	AVERAGE UpdateDashboardRequestBodyFilter
	MIN     UpdateDashboardRequestBodyFilter
	MAX     UpdateDashboardRequestBodyFilter
	SUM     UpdateDashboardRequestBodyFilter
}

func GetUpdateDashboardRequestBodyFilterEnum() UpdateDashboardRequestBodyFilterEnum {
	return UpdateDashboardRequestBodyFilterEnum{
		AVERAGE: UpdateDashboardRequestBodyFilter{
			value: "average",
		},
		MIN: UpdateDashboardRequestBodyFilter{
			value: "min",
		},
		MAX: UpdateDashboardRequestBodyFilter{
			value: "max",
		},
		SUM: UpdateDashboardRequestBodyFilter{
			value: "sum",
		},
	}
}

func (c UpdateDashboardRequestBodyFilter) Value() string {
	return c.value
}

func (c UpdateDashboardRequestBodyFilter) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDashboardRequestBodyFilter) UnmarshalJSON(b []byte) error {
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

type UpdateDashboardRequestBodyPeriod struct {
	value int32
}

type UpdateDashboardRequestBodyPeriodEnum struct {
	E_1     UpdateDashboardRequestBodyPeriod
	E_60    UpdateDashboardRequestBodyPeriod
	E_300   UpdateDashboardRequestBodyPeriod
	E_1200  UpdateDashboardRequestBodyPeriod
	E_3600  UpdateDashboardRequestBodyPeriod
	E_14400 UpdateDashboardRequestBodyPeriod
	E_86400 UpdateDashboardRequestBodyPeriod
}

func GetUpdateDashboardRequestBodyPeriodEnum() UpdateDashboardRequestBodyPeriodEnum {
	return UpdateDashboardRequestBodyPeriodEnum{
		E_1: UpdateDashboardRequestBodyPeriod{
			value: 1,
		}, E_60: UpdateDashboardRequestBodyPeriod{
			value: 60,
		}, E_300: UpdateDashboardRequestBodyPeriod{
			value: 300,
		}, E_1200: UpdateDashboardRequestBodyPeriod{
			value: 1200,
		}, E_3600: UpdateDashboardRequestBodyPeriod{
			value: 3600,
		}, E_14400: UpdateDashboardRequestBodyPeriod{
			value: 14400,
		}, E_86400: UpdateDashboardRequestBodyPeriod{
			value: 86400,
		},
	}
}

func (c UpdateDashboardRequestBodyPeriod) Value() int32 {
	return c.value
}

func (c UpdateDashboardRequestBodyPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDashboardRequestBodyPeriod) UnmarshalJSON(b []byte) error {
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

type UpdateDashboardRequestBodyDisplayTime struct {
	value int32
}

type UpdateDashboardRequestBodyDisplayTimeEnum struct {
	E_0     UpdateDashboardRequestBodyDisplayTime
	E_5     UpdateDashboardRequestBodyDisplayTime
	E_15    UpdateDashboardRequestBodyDisplayTime
	E_30    UpdateDashboardRequestBodyDisplayTime
	E_60    UpdateDashboardRequestBodyDisplayTime
	E_120   UpdateDashboardRequestBodyDisplayTime
	E_180   UpdateDashboardRequestBodyDisplayTime
	E_720   UpdateDashboardRequestBodyDisplayTime
	E_1440  UpdateDashboardRequestBodyDisplayTime
	E_10080 UpdateDashboardRequestBodyDisplayTime
	E_43200 UpdateDashboardRequestBodyDisplayTime
}

func GetUpdateDashboardRequestBodyDisplayTimeEnum() UpdateDashboardRequestBodyDisplayTimeEnum {
	return UpdateDashboardRequestBodyDisplayTimeEnum{
		E_0: UpdateDashboardRequestBodyDisplayTime{
			value: 0,
		}, E_5: UpdateDashboardRequestBodyDisplayTime{
			value: 5,
		}, E_15: UpdateDashboardRequestBodyDisplayTime{
			value: 15,
		}, E_30: UpdateDashboardRequestBodyDisplayTime{
			value: 30,
		}, E_60: UpdateDashboardRequestBodyDisplayTime{
			value: 60,
		}, E_120: UpdateDashboardRequestBodyDisplayTime{
			value: 120,
		}, E_180: UpdateDashboardRequestBodyDisplayTime{
			value: 180,
		}, E_720: UpdateDashboardRequestBodyDisplayTime{
			value: 720,
		}, E_1440: UpdateDashboardRequestBodyDisplayTime{
			value: 1440,
		}, E_10080: UpdateDashboardRequestBodyDisplayTime{
			value: 10080,
		}, E_43200: UpdateDashboardRequestBodyDisplayTime{
			value: 43200,
		},
	}
}

func (c UpdateDashboardRequestBodyDisplayTime) Value() int32 {
	return c.value
}

func (c UpdateDashboardRequestBodyDisplayTime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDashboardRequestBodyDisplayTime) UnmarshalJSON(b []byte) error {
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

type UpdateDashboardRequestBodyRefreshTime struct {
	value int32
}

type UpdateDashboardRequestBodyRefreshTimeEnum struct {
	E_0    UpdateDashboardRequestBodyRefreshTime
	E_10   UpdateDashboardRequestBodyRefreshTime
	E_60   UpdateDashboardRequestBodyRefreshTime
	E_300  UpdateDashboardRequestBodyRefreshTime
	E_1200 UpdateDashboardRequestBodyRefreshTime
}

func GetUpdateDashboardRequestBodyRefreshTimeEnum() UpdateDashboardRequestBodyRefreshTimeEnum {
	return UpdateDashboardRequestBodyRefreshTimeEnum{
		E_0: UpdateDashboardRequestBodyRefreshTime{
			value: 0,
		}, E_10: UpdateDashboardRequestBodyRefreshTime{
			value: 10,
		}, E_60: UpdateDashboardRequestBodyRefreshTime{
			value: 60,
		}, E_300: UpdateDashboardRequestBodyRefreshTime{
			value: 300,
		}, E_1200: UpdateDashboardRequestBodyRefreshTime{
			value: 1200,
		},
	}
}

func (c UpdateDashboardRequestBodyRefreshTime) Value() int32 {
	return c.value
}

func (c UpdateDashboardRequestBodyRefreshTime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDashboardRequestBodyRefreshTime) UnmarshalJSON(b []byte) error {
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

type UpdateDashboardRequestBodyTimeInterval struct {
	value int32
}

type UpdateDashboardRequestBodyTimeIntervalEnum struct {
	E_10000 UpdateDashboardRequestBodyTimeInterval
	E_30000 UpdateDashboardRequestBodyTimeInterval
	E_60000 UpdateDashboardRequestBodyTimeInterval
}

func GetUpdateDashboardRequestBodyTimeIntervalEnum() UpdateDashboardRequestBodyTimeIntervalEnum {
	return UpdateDashboardRequestBodyTimeIntervalEnum{
		E_10000: UpdateDashboardRequestBodyTimeInterval{
			value: 10000,
		}, E_30000: UpdateDashboardRequestBodyTimeInterval{
			value: 30000,
		}, E_60000: UpdateDashboardRequestBodyTimeInterval{
			value: 60000,
		},
	}
}

func (c UpdateDashboardRequestBodyTimeInterval) Value() int32 {
	return c.value
}

func (c UpdateDashboardRequestBodyTimeInterval) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDashboardRequestBodyTimeInterval) UnmarshalJSON(b []byte) error {
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
