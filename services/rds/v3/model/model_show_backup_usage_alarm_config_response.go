package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowBackupUsageAlarmConfigResponse Response Object
type ShowBackupUsageAlarmConfigResponse struct {

	// **参数解释**：  告警开关。  **约束限制**：  不涉及。  **取值范围**：  - ON - OFF  **默认取值**：  OFF。
	AlarmEnabled *ShowBackupUsageAlarmConfigResponseAlarmEnabled `json:"alarm_enabled,omitempty"`

	// **参数解释**：  阈值百分比，占免费备份空间大小的百分比。  **约束限制**：  不涉及。  **取值范围**：  1-100。  **默认取值**：  90。
	ThresholdPercent *int32 `json:"threshold_percent,omitempty"`

	// **参数解释**：  增量百分比，占免费备份空间大小的百分比。  **约束限制**：  不涉及。  **取值范围**：  1-100。  **默认取值**：  10。
	IncrementPercent *int32 `json:"increment_percent,omitempty"`
	HttpStatusCode   int    `json:"-"`
}

func (o ShowBackupUsageAlarmConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupUsageAlarmConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowBackupUsageAlarmConfigResponse", string(data)}, " ")
}

type ShowBackupUsageAlarmConfigResponseAlarmEnabled struct {
	value string
}

type ShowBackupUsageAlarmConfigResponseAlarmEnabledEnum struct {
	ON  ShowBackupUsageAlarmConfigResponseAlarmEnabled
	OFF ShowBackupUsageAlarmConfigResponseAlarmEnabled
}

func GetShowBackupUsageAlarmConfigResponseAlarmEnabledEnum() ShowBackupUsageAlarmConfigResponseAlarmEnabledEnum {
	return ShowBackupUsageAlarmConfigResponseAlarmEnabledEnum{
		ON: ShowBackupUsageAlarmConfigResponseAlarmEnabled{
			value: "ON",
		},
		OFF: ShowBackupUsageAlarmConfigResponseAlarmEnabled{
			value: "OFF",
		},
	}
}

func (c ShowBackupUsageAlarmConfigResponseAlarmEnabled) Value() string {
	return c.value
}

func (c ShowBackupUsageAlarmConfigResponseAlarmEnabled) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowBackupUsageAlarmConfigResponseAlarmEnabled) UnmarshalJSON(b []byte) error {
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
