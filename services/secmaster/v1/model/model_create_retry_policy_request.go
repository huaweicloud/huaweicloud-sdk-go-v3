package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateRetryPolicyRequest Request Object
type CreateRetryPolicyRequest struct {

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"Content-Type"`

	// 服务版本，例如25.5.0
	XSecmasterVersion string `json:"X-Secmaster-Version"`

	// **参数解释：** 工作空间id。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	WorkspaceId string `json:"workspace_id"`

	// 操作类型：create创建，retry重试
	ActionType CreateRetryPolicyRequestActionType `json:"action_type"`

	Body *CreateRetryPolicyRequestBody `json:"body,omitempty"`
}

func (o CreateRetryPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRetryPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateRetryPolicyRequest", string(data)}, " ")
}

type CreateRetryPolicyRequestActionType struct {
	value string
}

type CreateRetryPolicyRequestActionTypeEnum struct {
	CREATE CreateRetryPolicyRequestActionType
	RETRY  CreateRetryPolicyRequestActionType
}

func GetCreateRetryPolicyRequestActionTypeEnum() CreateRetryPolicyRequestActionTypeEnum {
	return CreateRetryPolicyRequestActionTypeEnum{
		CREATE: CreateRetryPolicyRequestActionType{
			value: "create",
		},
		RETRY: CreateRetryPolicyRequestActionType{
			value: "retry",
		},
	}
}

func (c CreateRetryPolicyRequestActionType) Value() string {
	return c.value
}

func (c CreateRetryPolicyRequestActionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRetryPolicyRequestActionType) UnmarshalJSON(b []byte) error {
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
