package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// V2ImageVersionStatus **参数解释**：镜像版本状态。 **取值范围**：   - AVAILABLE：可用。   - UNAVAILABLE：不可用。   - UNKNOWN：未知。
type V2ImageVersionStatus struct {
	value string
}

type V2ImageVersionStatusEnum struct {
	AVAILABLE   V2ImageVersionStatus
	UNAVAILABLE V2ImageVersionStatus
	UNKNOWN     V2ImageVersionStatus
}

func GetV2ImageVersionStatusEnum() V2ImageVersionStatusEnum {
	return V2ImageVersionStatusEnum{
		AVAILABLE: V2ImageVersionStatus{
			value: "AVAILABLE",
		},
		UNAVAILABLE: V2ImageVersionStatus{
			value: "UNAVAILABLE",
		},
		UNKNOWN: V2ImageVersionStatus{
			value: "UNKNOWN",
		},
	}
}

func (c V2ImageVersionStatus) Value() string {
	return c.value
}

func (c V2ImageVersionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *V2ImageVersionStatus) UnmarshalJSON(b []byte) error {
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
