package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PostPreCheckResp
type PostPreCheckResp struct {

	// **参数解释**：  任务ID，用于唯一标识一个数据复制任务，在查询、修改、删除等操作中作为任务的识别凭证。  **约束限制**：  不涉及。  **取值范围**：  UUID格式字符串。  **默认取值**：  不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**：  预检查后返回的ID，可用于查询预检查结果。  **约束限制**：  不涉及。  **取值范围**：  UUID格式字符串。  **默认取值**：  不涉及
	PrecheckId *string `json:"precheck_id,omitempty"`

	// **参数解释**：  预检查结果状态，表示预检查是否通过。  **约束限制**：  不涉及。  **取值范围**：  - success：预检查成功 - failed：预检查失败  **默认取值**：  不涉及
	Status *PostPreCheckRespStatus `json:"status,omitempty"`

	// **参数解释**：  错误码，当预检查失败时返回的错误代码。  **约束限制**：  不涉及。  **取值范围**：  格式为DRS.XXXXXX。  **默认取值**：  不涉及
	ErrorCode *string `json:"error_code,omitempty"`

	// **参数解释**：  错误详细信息，当预检查失败时返回的具体错误描述。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o PostPreCheckResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostPreCheckResp struct{}"
	}

	return strings.Join([]string{"PostPreCheckResp", string(data)}, " ")
}

type PostPreCheckRespStatus struct {
	value string
}

type PostPreCheckRespStatusEnum struct {
	SUCCESS PostPreCheckRespStatus
	FAILED  PostPreCheckRespStatus
}

func GetPostPreCheckRespStatusEnum() PostPreCheckRespStatusEnum {
	return PostPreCheckRespStatusEnum{
		SUCCESS: PostPreCheckRespStatus{
			value: "success",
		},
		FAILED: PostPreCheckRespStatus{
			value: "failed",
		},
	}
}

func (c PostPreCheckRespStatus) Value() string {
	return c.value
}

func (c PostPreCheckRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostPreCheckRespStatus) UnmarshalJSON(b []byte) error {
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
