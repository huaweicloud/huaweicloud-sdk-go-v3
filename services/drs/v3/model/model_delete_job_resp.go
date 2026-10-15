package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteJobResp 删除或者结束任务返回体。
type DeleteJobResp struct {

	// **参数解释**：  结束或删除操作的任务ID，用于标识该操作结果对应的任务。  **约束限制**：  不涉及。  **取值范围**：  UUID格式字符串。  **默认取值**：  不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**：  批量结束或删除操作的任务执行结果状态，表示该任务的操作是否成功。  **约束限制**：  不涉及。  **取值范围**：  - success：操作成功 - failed：操作失败  **默认取值**：  不涉及
	Status *DeleteJobRespStatus `json:"status,omitempty"`

	// **参数解释**：  错误码，当任务操作失败时返回的错误代码。  **约束限制**：  不涉及。  **取值范围**：  格式为DRS.XXXXXX。  **默认取值**：  不涉及
	ErrorCode *string `json:"error_code,omitempty"`

	// **参数解释**：  错误详细信息，当任务操作失败时返回的具体错误描述。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o DeleteJobResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteJobResp struct{}"
	}

	return strings.Join([]string{"DeleteJobResp", string(data)}, " ")
}

type DeleteJobRespStatus struct {
	value string
}

type DeleteJobRespStatusEnum struct {
	SUCCESS DeleteJobRespStatus
	FAILED  DeleteJobRespStatus
}

func GetDeleteJobRespStatusEnum() DeleteJobRespStatusEnum {
	return DeleteJobRespStatusEnum{
		SUCCESS: DeleteJobRespStatus{
			value: "success",
		},
		FAILED: DeleteJobRespStatus{
			value: "failed",
		},
	}
}

func (c DeleteJobRespStatus) Value() string {
	return c.value
}

func (c DeleteJobRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteJobRespStatus) UnmarshalJSON(b []byte) error {
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
