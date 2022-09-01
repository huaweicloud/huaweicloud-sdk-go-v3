package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// app详细信息
type App struct {

	// app名称
	AppName *string `json:"app_name,omitempty" xml:"app_name"`

	// 应用id
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	State *AppState `json:"state,omitempty" xml:"state"`

	// RTC覆盖范围。  取值如下：    - DOMESTIC：国内范围。   - OVERSEA：海外范围。   - GLOBAL：全球范围。
	Scope *AppScope `json:"scope,omitempty" xml:"scope"`

	// 账号名
	TenantName *string `json:"tenant_name,omitempty" xml:"tenant_name"`

	// 域名，App对应域名
	Domain *string `json:"domain,omitempty" xml:"domain"`

	// 创建时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	Authentication *AppAuth `json:"authentication,omitempty" xml:"authentication"`

	Callbacks *AppCallbacks `json:"callbacks,omitempty" xml:"callbacks"`

	AutoRecordMode *AppAutoRecordMode `json:"auto_record_mode,omitempty" xml:"auto_record_mode"`
}

func (o App) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "App struct{}"
	}

	return strings.Join([]string{"App", string(data)}, " ")
}

type AppScope struct {
	value string
}

type AppScopeEnum struct {
	DOMESTIC AppScope
	OVERSEA  AppScope
	GLOBAL   AppScope
}

func GetAppScopeEnum() AppScopeEnum {
	return AppScopeEnum{
		DOMESTIC: AppScope{
			value: "DOMESTIC",
		},
		OVERSEA: AppScope{
			value: "OVERSEA",
		},
		GLOBAL: AppScope{
			value: "GLOBAL",
		},
	}
}

func (c AppScope) Value() string {
	return c.value
}

func (c AppScope) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AppScope) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
