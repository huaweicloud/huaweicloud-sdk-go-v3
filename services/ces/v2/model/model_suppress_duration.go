package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SuppressDuration 发送告警的周期，值可为0, 300, 600, 900, 1800, 3600, 10800, 21600, 43200, 86400；0表示只告警一次，300表示每5分钟告警一次，600表示每10分钟告警一次，900表示每15分钟告警一次，1800表示每30分钟告警一次，3600表示每1小时告警一次，10800表示每3小时告警一次，21600表示每6小时告警一次，43200表示每12小时告警一次，86400表示每1天告警一次。
type SuppressDuration struct {
	value int32
}

type SuppressDurationEnum struct {
	E_0     SuppressDuration
	E_300   SuppressDuration
	E_600   SuppressDuration
	E_900   SuppressDuration
	E_1800  SuppressDuration
	E_3600  SuppressDuration
	E_10800 SuppressDuration
	E_21600 SuppressDuration
	E_43200 SuppressDuration
	E_86400 SuppressDuration
}

func GetSuppressDurationEnum() SuppressDurationEnum {
	return SuppressDurationEnum{
		E_0: SuppressDuration{
			value: 0,
		}, E_300: SuppressDuration{
			value: 300,
		}, E_600: SuppressDuration{
			value: 600,
		}, E_900: SuppressDuration{
			value: 900,
		}, E_1800: SuppressDuration{
			value: 1800,
		}, E_3600: SuppressDuration{
			value: 3600,
		}, E_10800: SuppressDuration{
			value: 10800,
		}, E_21600: SuppressDuration{
			value: 21600,
		}, E_43200: SuppressDuration{
			value: 43200,
		}, E_86400: SuppressDuration{
			value: 86400,
		},
	}
}

func (c SuppressDuration) Value() int32 {
	return c.value
}

func (c SuppressDuration) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SuppressDuration) UnmarshalJSON(b []byte) error {
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
