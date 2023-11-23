package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type InstanceLogConfigDetail struct {

	// 日志类型。slow_log表示慢日志。
	LogType InstanceLogConfigDetailLogType `json:"log_type"`

	// 关联的LTS日志组ID，若enabled为false，表示最近一次关联的LTS日志组ID。
	LtsGroupId string `json:"lts_group_id"`

	// 关联的LTS日志流ID，若enabled为false，表示最近一次关联的LTS日志流ID。
	LtsStreamId string `json:"lts_stream_id"`

	// 关联的LTS日志流是否启用，true代表已启用，false代表未启用。
	Enabled *bool `json:"enabled,omitempty"`
}

func (o InstanceLogConfigDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceLogConfigDetail struct{}"
	}

	return strings.Join([]string{"InstanceLogConfigDetail", string(data)}, " ")
}

type InstanceLogConfigDetailLogType struct {
	value string
}

type InstanceLogConfigDetailLogTypeEnum struct {
	SLOW_LOG InstanceLogConfigDetailLogType
}

func GetInstanceLogConfigDetailLogTypeEnum() InstanceLogConfigDetailLogTypeEnum {
	return InstanceLogConfigDetailLogTypeEnum{
		SLOW_LOG: InstanceLogConfigDetailLogType{
			value: "slow_log",
		},
	}
}

func (c InstanceLogConfigDetailLogType) Value() string {
	return c.value
}

func (c InstanceLogConfigDetailLogType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceLogConfigDetailLogType) UnmarshalJSON(b []byte) error {
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
