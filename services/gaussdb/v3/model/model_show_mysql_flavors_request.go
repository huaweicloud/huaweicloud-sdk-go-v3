package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowMysqlFlavorsRequest struct {
	// 语言。

	XLanguage *string `json:"X-Language,omitempty"`
	// 数据库引擎名称。

	DatabaseName string `json:"database_name"`
	// 数据库版本号，目前仅支持兼容MySQL 8.0。

	VersionName *string `json:"version_name,omitempty"`
	// 规格的可用区模式，现在仅支持\"single\"、\"multi\"，不区分大小写。

	AvailabilityZoneMode ShowMysqlFlavorsRequestAvailabilityZoneMode `json:"availability_zone_mode"`
	// 规格编码。

	SpecCode *string `json:"spec_code,omitempty"`
}

func (o ShowMysqlFlavorsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMysqlFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ShowMysqlFlavorsRequest", string(data)}, " ")
}

type ShowMysqlFlavorsRequestAvailabilityZoneMode struct {
	value string
}

type ShowMysqlFlavorsRequestAvailabilityZoneModeEnum struct {
	SINGLE ShowMysqlFlavorsRequestAvailabilityZoneMode
	MULTI  ShowMysqlFlavorsRequestAvailabilityZoneMode
}

func GetShowMysqlFlavorsRequestAvailabilityZoneModeEnum() ShowMysqlFlavorsRequestAvailabilityZoneModeEnum {
	return ShowMysqlFlavorsRequestAvailabilityZoneModeEnum{
		SINGLE: ShowMysqlFlavorsRequestAvailabilityZoneMode{
			value: "single",
		},
		MULTI: ShowMysqlFlavorsRequestAvailabilityZoneMode{
			value: "multi",
		},
	}
}

func (c ShowMysqlFlavorsRequestAvailabilityZoneMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowMysqlFlavorsRequestAvailabilityZoneMode) UnmarshalJSON(b []byte) error {
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
