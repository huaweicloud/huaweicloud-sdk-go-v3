package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ModifyJobResp struct {

	// **参数解释**：  任务ID，用于唯一标识一个数据复制任务。  **约束限制**：  不涉及。  **取值范围**：  UUID格式字符串。  **默认取值**：  不涉及
	Id string `json:"id"`

	// **参数解释**：  任务修改操作的执行状态。  **约束限制**：  不涉及。  **取值范围**：  - success：成功 - failed：失败  **默认取值**：  不涉及
	Status ModifyJobRespStatus `json:"status"`

	// **参数解释**：  错误码，当任务修改失败时返回的错误代码。  **约束限制**：  不涉及。  **取值范围**：  格式为DRS.XXXXXX。  **默认取值**：  不涉及
	ErrorCode *string `json:"error_code,omitempty"`

	// **参数解释**：  错误详细信息，当任务修改失败时返回的具体错误描述。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o ModifyJobResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyJobResp struct{}"
	}

	return strings.Join([]string{"ModifyJobResp", string(data)}, " ")
}

type ModifyJobRespStatus struct {
	value string
}

type ModifyJobRespStatusEnum struct {
	SUCCESS ModifyJobRespStatus
	FAILED  ModifyJobRespStatus
}

func GetModifyJobRespStatusEnum() ModifyJobRespStatusEnum {
	return ModifyJobRespStatusEnum{
		SUCCESS: ModifyJobRespStatus{
			value: "success",
		},
		FAILED: ModifyJobRespStatus{
			value: "failed",
		},
	}
}

func (c ModifyJobRespStatus) Value() string {
	return c.value
}

func (c ModifyJobRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyJobRespStatus) UnmarshalJSON(b []byte) error {
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
