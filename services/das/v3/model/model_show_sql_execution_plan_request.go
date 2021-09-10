package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowSqlExecutionPlanRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 数据库用户ID

	DbUserId string `json:"db_user_id"`
	// 数据库名称

	Database string `json:"database"`
	// SQL语句

	Sql string `json:"sql"`
	// 语言

	XLanguage *ShowSqlExecutionPlanRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowSqlExecutionPlanRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowSqlExecutionPlanRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlExecutionPlanRequest", string(data)}, " ")
}

type ShowSqlExecutionPlanRequestXLanguage struct {
	value string
}

type ShowSqlExecutionPlanRequestXLanguageEnum struct {
	ZH_CN ShowSqlExecutionPlanRequestXLanguage
	EN_US ShowSqlExecutionPlanRequestXLanguage
}

func GetShowSqlExecutionPlanRequestXLanguageEnum() ShowSqlExecutionPlanRequestXLanguageEnum {
	return ShowSqlExecutionPlanRequestXLanguageEnum{
		ZH_CN: ShowSqlExecutionPlanRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowSqlExecutionPlanRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowSqlExecutionPlanRequestXLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowSqlExecutionPlanRequestXLanguage) UnmarshalJSON(b []byte) error {
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
