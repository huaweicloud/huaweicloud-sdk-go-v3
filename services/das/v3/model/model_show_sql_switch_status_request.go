package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowSqlSwitchStatusRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 开关类型。取值DAS SQL Explorer和DAS Slow Query Log，分别表示DAS收集全量SQL开关和DAS收集慢SQL开关。
	Type string `json:"type" xml:"type"`

	// 数据库类型。当前全量SQL支持的数据库类型包括MySQL和GaussDB(for MySQL)，慢SQL支持的类型：MySQL、GaussDB(for MySQL)、PostgreSQL。
	DatastoreType string `json:"datastore_type" xml:"datastore_type"`

	// 请求语言类型。
	XLanguage *ShowSqlSwitchStatusRequestXLanguage `json:"X-Language,omitempty" xml:"X-Language"`
}

func (o ShowSqlSwitchStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlSwitchStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlSwitchStatusRequest", string(data)}, " ")
}

type ShowSqlSwitchStatusRequestXLanguage struct {
	value string
}

type ShowSqlSwitchStatusRequestXLanguageEnum struct {
	EN_US ShowSqlSwitchStatusRequestXLanguage
	ZH_CN ShowSqlSwitchStatusRequestXLanguage
}

func GetShowSqlSwitchStatusRequestXLanguageEnum() ShowSqlSwitchStatusRequestXLanguageEnum {
	return ShowSqlSwitchStatusRequestXLanguageEnum{
		EN_US: ShowSqlSwitchStatusRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowSqlSwitchStatusRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowSqlSwitchStatusRequestXLanguage) Value() string {
	return c.value
}

func (c ShowSqlSwitchStatusRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSqlSwitchStatusRequestXLanguage) UnmarshalJSON(b []byte) error {
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
