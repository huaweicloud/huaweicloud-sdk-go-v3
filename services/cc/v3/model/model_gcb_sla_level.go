package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// GcbSlaLevel 全域互联带宽的带宽等级。
type GcbSlaLevel struct {

	// 功能说明：描述网络等级，从高到低分为铂金、金、银。默认金，其余租户白名单控制。 - Pt: 铂金 - Au: 金 - Ag: 银
	SlaLevel *GcbSlaLevelSlaLevel `json:"sla_level,omitempty"`
}

func (o GcbSlaLevel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GcbSlaLevel struct{}"
	}

	return strings.Join([]string{"GcbSlaLevel", string(data)}, " ")
}

type GcbSlaLevelSlaLevel struct {
	value string
}

type GcbSlaLevelSlaLevelEnum struct {
	PT GcbSlaLevelSlaLevel
	AU GcbSlaLevelSlaLevel
	AG GcbSlaLevelSlaLevel
}

func GetGcbSlaLevelSlaLevelEnum() GcbSlaLevelSlaLevelEnum {
	return GcbSlaLevelSlaLevelEnum{
		PT: GcbSlaLevelSlaLevel{
			value: "Pt",
		},
		AU: GcbSlaLevelSlaLevel{
			value: "Au",
		},
		AG: GcbSlaLevelSlaLevel{
			value: "Ag",
		},
	}
}

func (c GcbSlaLevelSlaLevel) Value() string {
	return c.value
}

func (c GcbSlaLevelSlaLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GcbSlaLevelSlaLevel) UnmarshalJSON(b []byte) error {
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
