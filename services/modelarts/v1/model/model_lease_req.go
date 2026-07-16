package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LeaseReq 实例租约配置。
type LeaseReq struct {

	// **参数解释**：定时停止，以当前时刻为起点，运行的时长（到期后自动停止）。单位：毫秒。 **约束限制**：不涉及。 **取值范围**：3600000-259200000。 **默认取值**：3600000。
	Duration *int64 `json:"duration,omitempty"`

	// **参数解释**：自动停止类别。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： - TIMING：自动停止。 - IDLE：空闲停止。  **默认取值**：TIMING。
	Type *LeaseReqType `json:"type,omitempty"`
}

func (o LeaseReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LeaseReq struct{}"
	}

	return strings.Join([]string{"LeaseReq", string(data)}, " ")
}

type LeaseReqType struct {
	value string
}

type LeaseReqTypeEnum struct {
	TIMING LeaseReqType
	IDLE   LeaseReqType
}

func GetLeaseReqTypeEnum() LeaseReqTypeEnum {
	return LeaseReqTypeEnum{
		TIMING: LeaseReqType{
			value: "timing",
		},
		IDLE: LeaseReqType{
			value: "idle",
		},
	}
}

func (c LeaseReqType) Value() string {
	return c.value
}

func (c LeaseReqType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LeaseReqType) UnmarshalJSON(b []byte) error {
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
