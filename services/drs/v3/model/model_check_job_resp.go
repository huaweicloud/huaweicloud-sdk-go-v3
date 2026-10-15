package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CheckJobResp 响应体
type CheckJobResp struct {

	// **参数解释**：  任务ID，用于唯一标识一个数据复制任务。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Id string `json:"id"`

	// **参数解释**：  测试结果状态。  **约束限制**：  不涉及。  **取值范围**：  - true：成功 - false：失败  **默认取值**：  不涉及
	Status CheckJobRespStatus `json:"status"`

	// **参数解释**：  错误码，当任务报错时返回的错误代码。  **约束限制**：  不涉及。  **取值范围**：  格式为DRS.XXXXXX。  **默认取值**：  不涉及
	ErrorCode *string `json:"error_code,omitempty"`

	// **参数解释**：  错误详细信息，当任务执行失败时返回的具体错误描述。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	ErrorMsg *string `json:"error_msg,omitempty"`

	// **参数解释**：  操作是否成功，标识测试连接操作的执行结果。  **约束限制**：  不涉及。  **取值范围**：  - true：操作成功 - false：操作失败  **默认取值**：  不涉及
	Success *bool `json:"success,omitempty"`
}

func (o CheckJobResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckJobResp struct{}"
	}

	return strings.Join([]string{"CheckJobResp", string(data)}, " ")
}

type CheckJobRespStatus struct {
	value string
}

type CheckJobRespStatusEnum struct {
	TRUE  CheckJobRespStatus
	FALSE CheckJobRespStatus
}

func GetCheckJobRespStatusEnum() CheckJobRespStatusEnum {
	return CheckJobRespStatusEnum{
		TRUE: CheckJobRespStatus{
			value: "true",
		},
		FALSE: CheckJobRespStatus{
			value: "false",
		},
	}
}

func (c CheckJobRespStatus) Value() string {
	return c.value
}

func (c CheckJobRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CheckJobRespStatus) UnmarshalJSON(b []byte) error {
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
