package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowAppResponse struct {

	// app名称
	AppName *string `json:"app_name,omitempty" xml:"app_name"`

	// 应用id
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	State *AppState `json:"state,omitempty" xml:"state"`

	// RTC覆盖范围。  取值如下：    - DOMESTIC：国内范围。   - OVERSEA：海外范围。   - GLOBAL：全球范围。
	Scope *ShowAppResponseScope `json:"scope,omitempty" xml:"scope"`

	// 账号名
	TenantName *string `json:"tenant_name,omitempty" xml:"tenant_name"`

	// 域名，App对应域名
	Domain *string `json:"domain,omitempty" xml:"domain"`

	// 创建时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	Authentication *AppAuth `json:"authentication,omitempty" xml:"authentication"`

	Callbacks *AppCallbacks `json:"callbacks,omitempty" xml:"callbacks"`

	AutoRecordMode *AppAutoRecordMode `json:"auto_record_mode,omitempty" xml:"auto_record_mode"`

	XRequestId     *string `json:"X-request-Id,omitempty" xml:"X-request-Id"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppResponse struct{}"
	}

	return strings.Join([]string{"ShowAppResponse", string(data)}, " ")
}

type ShowAppResponseScope struct {
	value string
}

type ShowAppResponseScopeEnum struct {
	DOMESTIC ShowAppResponseScope
	OVERSEA  ShowAppResponseScope
	GLOBAL   ShowAppResponseScope
}

func GetShowAppResponseScopeEnum() ShowAppResponseScopeEnum {
	return ShowAppResponseScopeEnum{
		DOMESTIC: ShowAppResponseScope{
			value: "DOMESTIC",
		},
		OVERSEA: ShowAppResponseScope{
			value: "OVERSEA",
		},
		GLOBAL: ShowAppResponseScope{
			value: "GLOBAL",
		},
	}
}

func (c ShowAppResponseScope) Value() string {
	return c.value
}

func (c ShowAppResponseScope) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAppResponseScope) UnmarshalJSON(b []byte) error {
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
