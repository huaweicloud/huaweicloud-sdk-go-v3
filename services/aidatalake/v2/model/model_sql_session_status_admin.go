package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SqlSessionStatusAdmin **参数解释**：SQL会话状态。 **取值范围**：    - RUNNING：运行中。   - CLOSED：已关闭。   - WAITING：等待中。   - CREATING：创建中。   - FAIL：失败。
type SqlSessionStatusAdmin struct {
	value string
}

type SqlSessionStatusAdminEnum struct {
	RUNNING  SqlSessionStatusAdmin
	CLOSED   SqlSessionStatusAdmin
	WAITING  SqlSessionStatusAdmin
	CREATING SqlSessionStatusAdmin
	FAIL     SqlSessionStatusAdmin
}

func GetSqlSessionStatusAdminEnum() SqlSessionStatusAdminEnum {
	return SqlSessionStatusAdminEnum{
		RUNNING: SqlSessionStatusAdmin{
			value: "RUNNING",
		},
		CLOSED: SqlSessionStatusAdmin{
			value: "CLOSED",
		},
		WAITING: SqlSessionStatusAdmin{
			value: "WAITING",
		},
		CREATING: SqlSessionStatusAdmin{
			value: "CREATING",
		},
		FAIL: SqlSessionStatusAdmin{
			value: "FAIL",
		},
	}
}

func (c SqlSessionStatusAdmin) Value() string {
	return c.value
}

func (c SqlSessionStatusAdmin) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SqlSessionStatusAdmin) UnmarshalJSON(b []byte) error {
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
