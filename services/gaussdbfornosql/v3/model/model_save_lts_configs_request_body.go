package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SaveLtsConfigsRequestBody struct {

	// 需要建立关联的实例ID列表。
	InstanceIds []string `json:"instance_ids"`

	// 日志类型。slow_log表示慢日志。
	LogType SaveLtsConfigsRequestBodyLogType `json:"log_type"`

	// LTS日志组ID。可通过云日志服务-“查询账号下所有日志组”API接口获取。
	LtsGroupId string `json:"lts_group_id"`

	// LTS日志流ID。可通过云日志服务-“查询指定日志组下的所有日志流”API接口获取。
	LtsStreamId string `json:"lts_stream_id"`
}

func (o SaveLtsConfigsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveLtsConfigsRequestBody struct{}"
	}

	return strings.Join([]string{"SaveLtsConfigsRequestBody", string(data)}, " ")
}

type SaveLtsConfigsRequestBodyLogType struct {
	value string
}

type SaveLtsConfigsRequestBodyLogTypeEnum struct {
	SLOW_LOG SaveLtsConfigsRequestBodyLogType
}

func GetSaveLtsConfigsRequestBodyLogTypeEnum() SaveLtsConfigsRequestBodyLogTypeEnum {
	return SaveLtsConfigsRequestBodyLogTypeEnum{
		SLOW_LOG: SaveLtsConfigsRequestBodyLogType{
			value: "slow_log",
		},
	}
}

func (c SaveLtsConfigsRequestBodyLogType) Value() string {
	return c.value
}

func (c SaveLtsConfigsRequestBodyLogType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SaveLtsConfigsRequestBodyLogType) UnmarshalJSON(b []byte) error {
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
