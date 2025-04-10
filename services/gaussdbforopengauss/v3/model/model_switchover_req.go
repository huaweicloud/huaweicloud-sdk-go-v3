package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SwitchoverReq 容灾switchover请求参数。
type SwitchoverReq struct {

	// 是否支持倒换失败自愈，为空时默认不自愈。
	PostProcessConfig *SwitchoverReqPostProcessConfig `json:"post_process_config,omitempty"`

	// 容灾类型。
	DisasterType SwitchoverReqDisasterType `json:"disaster_type"`
}

func (o SwitchoverReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchoverReq struct{}"
	}

	return strings.Join([]string{"SwitchoverReq", string(data)}, " ")
}

type SwitchoverReqPostProcessConfig struct {
	value string
}

type SwitchoverReqPostProcessConfigEnum struct {
	AUTO   SwitchoverReqPostProcessConfig
	MANUAL SwitchoverReqPostProcessConfig
}

func GetSwitchoverReqPostProcessConfigEnum() SwitchoverReqPostProcessConfigEnum {
	return SwitchoverReqPostProcessConfigEnum{
		AUTO: SwitchoverReqPostProcessConfig{
			value: "AUTO",
		},
		MANUAL: SwitchoverReqPostProcessConfig{
			value: "MANUAL",
		},
	}
}

func (c SwitchoverReqPostProcessConfig) Value() string {
	return c.value
}

func (c SwitchoverReqPostProcessConfig) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SwitchoverReqPostProcessConfig) UnmarshalJSON(b []byte) error {
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

type SwitchoverReqDisasterType struct {
	value string
}

type SwitchoverReqDisasterTypeEnum struct {
	STREAM SwitchoverReqDisasterType
	DORADO SwitchoverReqDisasterType
}

func GetSwitchoverReqDisasterTypeEnum() SwitchoverReqDisasterTypeEnum {
	return SwitchoverReqDisasterTypeEnum{
		STREAM: SwitchoverReqDisasterType{
			value: "stream",
		},
		DORADO: SwitchoverReqDisasterType{
			value: "dorado",
		},
	}
}

func (c SwitchoverReqDisasterType) Value() string {
	return c.value
}

func (c SwitchoverReqDisasterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SwitchoverReqDisasterType) UnmarshalJSON(b []byte) error {
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
