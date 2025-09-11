package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SuppressDuration **参数解释**： 告警抑制时间，单位为秒，对应页面上创建告警规则时告警策略最后一个字段，该字段主要为解决告警频繁的问题。 **约束限制**： 不涉及。 **取值范围**： 枚举值，只能为0、300、600、900、1800、3600、10800、21600、43200、86400。 - 0：对于指标类告警，0代表告警一次。对于事件类告警，在立即触发场景中，0代表不抑制；在累计触发场景，0代表只告警一次。 - 300代表满足告警触发条件后每5分钟告警一次。 - 600代表满足告警触发条件后每10分钟告警一次。 - 900代表满足告警触发条件后每15分钟告警一次。 - 1800代表满足告警触发条件后每30分钟告警一次。 - 3600代表满足告警触发条件后每60分钟告警一次。 - 10800代表满足告警触发条件后每3小时告警一次。 - 21600代表满足告警触发条件后每6小时告警一次。 - 43200代表满足告警触发条件后每12小时告警一次。 - 86000代表满足告警触发条件后每一天告警一次。 **默认取值**： 不涉及。
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
