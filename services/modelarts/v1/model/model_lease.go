package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Lease 实例自动停止的倒计时信息。
type Lease struct {

	// **参数解释**：实例创建的时间，UTC毫秒。 **取值范围**：不涉及。
	CreateAt *int64 `json:"create_at,omitempty"`

	// **参数解释**：实例运行时长，以创建时间为起点计算，即“创建时间+duration > 当前时刻”时，系统会自动停止实例。 **取值范围**：不涉及。
	Duration *int64 `json:"duration,omitempty"`

	// **参数解释**：是否启用自动停止功能。 **取值范围**：布尔类型： - true：启动自动停止功能。 - false：关闭自动停止功能。
	Enable *bool `json:"enable,omitempty"`

	// **参数解释**：自动停止类别。 **取值范围**：枚举类型，取值如下： - TIMING：自动停止。 - IDLE：空闲停止。
	Type *LeaseType `json:"type,omitempty"`

	// **参数解释**：实例最后更新（不包括探活心跳）的时间，UTC毫秒。 **取值范围**：不涉及。
	UpdateAt *int64 `json:"update_at,omitempty"`
}

func (o Lease) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Lease struct{}"
	}

	return strings.Join([]string{"Lease", string(data)}, " ")
}

type LeaseType struct {
	value string
}

type LeaseTypeEnum struct {
	TIMING LeaseType
	IDLE   LeaseType
}

func GetLeaseTypeEnum() LeaseTypeEnum {
	return LeaseTypeEnum{
		TIMING: LeaseType{
			value: "timing",
		},
		IDLE: LeaseType{
			value: "idle",
		},
	}
}

func (c LeaseType) Value() string {
	return c.value
}

func (c LeaseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LeaseType) UnmarshalJSON(b []byte) error {
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
