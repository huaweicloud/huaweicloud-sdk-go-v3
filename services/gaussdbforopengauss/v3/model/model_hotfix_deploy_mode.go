package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type HotfixDeployMode struct {

	// **参数解释**： 补丁默认升级策略。 **取值范围**： - true：无需选择，默认升级该版本。 - false：用户可选，选择后升级该版本。
	DefaultUpgrade *bool `json:"default_upgrade,omitempty"`

	// **参数解释**： 热补丁升级策略的修改时间，UNIX时间戳格式，单位是毫秒。 **取值范围**： 不涉及。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// **参数解释**： 可升级该补丁的实例部署形态。 **取值范围**： - distributed：包括独立部署，基础版混合部署。 - centralization_standard：包括主备版。
	Mode *HotfixDeployModeMode `json:"mode,omitempty"`
}

func (o HotfixDeployMode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HotfixDeployMode struct{}"
	}

	return strings.Join([]string{"HotfixDeployMode", string(data)}, " ")
}

type HotfixDeployModeMode struct {
	value string
}

type HotfixDeployModeModeEnum struct {
	CENTRALIZATION_STANDARD HotfixDeployModeMode
	DISTRIBUTED             HotfixDeployModeMode
}

func GetHotfixDeployModeModeEnum() HotfixDeployModeModeEnum {
	return HotfixDeployModeModeEnum{
		CENTRALIZATION_STANDARD: HotfixDeployModeMode{
			value: "centralization_standard",
		},
		DISTRIBUTED: HotfixDeployModeMode{
			value: "distributed",
		},
	}
}

func (c HotfixDeployModeMode) Value() string {
	return c.value
}

func (c HotfixDeployModeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HotfixDeployModeMode) UnmarshalJSON(b []byte) error {
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
