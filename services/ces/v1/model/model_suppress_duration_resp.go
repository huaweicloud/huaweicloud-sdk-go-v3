package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SuppressDurationResp **参数解释**： 告警抑制时间，单位为秒，对应页面上创建告警规则时告警策略最后一个字段，该字段主要为解决告警频繁的问题。 **取值范围**： 枚举值，只能为0、300、600、900、1800、3600、10800、21600、43200、86400。 - 0：对于指标类告警，0代表告警一次。对于事件类告警，在立即触发场景中，0代表不抑制；在累计触发场景，0代表只告警一次。 - 300代表满足告警触发条件后每5分钟告警一次。 - 600代表满足告警触发条件后每10分钟告警一次。 - 900代表满足告警触发条件后每15分钟告警一次。 - 1800代表满足告警触发条件后每30分钟告警一次。 - 3600代表满足告警触发条件后每60分钟告警一次。 - 10800代表满足告警触发条件后每3小时告警一次。 - 21600代表满足告警触发条件后每6小时告警一次。 - 43200代表满足告警触发条件后每12小时告警一次。 - 86000代表满足告警触发条件后每一天告警一次。
type SuppressDurationResp struct {
	value int32
}

type SuppressDurationRespEnum struct {
	E_0     SuppressDurationResp
	E_300   SuppressDurationResp
	E_600   SuppressDurationResp
	E_900   SuppressDurationResp
	E_1800  SuppressDurationResp
	E_3600  SuppressDurationResp
	E_10800 SuppressDurationResp
	E_21600 SuppressDurationResp
	E_43200 SuppressDurationResp
	E_86400 SuppressDurationResp
}

func GetSuppressDurationRespEnum() SuppressDurationRespEnum {
	return SuppressDurationRespEnum{
		E_0: SuppressDurationResp{
			value: 0,
		}, E_300: SuppressDurationResp{
			value: 300,
		}, E_600: SuppressDurationResp{
			value: 600,
		}, E_900: SuppressDurationResp{
			value: 900,
		}, E_1800: SuppressDurationResp{
			value: 1800,
		}, E_3600: SuppressDurationResp{
			value: 3600,
		}, E_10800: SuppressDurationResp{
			value: 10800,
		}, E_21600: SuppressDurationResp{
			value: 21600,
		}, E_43200: SuppressDurationResp{
			value: 43200,
		}, E_86400: SuppressDurationResp{
			value: 86400,
		},
	}
}

func (c SuppressDurationResp) Value() int32 {
	return c.value
}

func (c SuppressDurationResp) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SuppressDurationResp) UnmarshalJSON(b []byte) error {
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
