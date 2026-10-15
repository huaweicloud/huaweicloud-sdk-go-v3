package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SqlSessionStatus **参数解释**：SQL会话状态。 **取值范围**：    - RUNNING：运行中。   - CLOSED：已关闭。   - WAITING：等待中。   - CREATING：创建中。   - FAIL：失败。
type SqlSessionStatus struct {
	value string
}

type SqlSessionStatusEnum struct {
	RUNNING  SqlSessionStatus
	CLOSED   SqlSessionStatus
	WAITING  SqlSessionStatus
	CREATING SqlSessionStatus
	FAIL     SqlSessionStatus
}

func GetSqlSessionStatusEnum() SqlSessionStatusEnum {
	return SqlSessionStatusEnum{
		RUNNING: SqlSessionStatus{
			value: "RUNNING",
		},
		CLOSED: SqlSessionStatus{
			value: "CLOSED",
		},
		WAITING: SqlSessionStatus{
			value: "WAITING",
		},
		CREATING: SqlSessionStatus{
			value: "CREATING",
		},
		FAIL: SqlSessionStatus{
			value: "FAIL",
		},
	}
}

func (c SqlSessionStatus) Value() string {
	return c.value
}

func (c SqlSessionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SqlSessionStatus) UnmarshalJSON(b []byte) error {
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
