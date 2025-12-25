package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateAopWorkflowVersionApprovelInfo 创建流程审核请求
type CreateAopWorkflowVersionApprovelInfo struct {

	// **参数解释**: 流程的审核意见 **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	Content string `json:"content"`

	// **参数解释**: 流程的审核结果 - PASS 通过 - UN_PASS 拒绝  **约束限制**: 不涉及         **取值范围**: - PASS - UN_PASS  **默认值**:  不涉及
	Result CreateAopWorkflowVersionApprovelInfoResult `json:"result"`
}

func (o CreateAopWorkflowVersionApprovelInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAopWorkflowVersionApprovelInfo struct{}"
	}

	return strings.Join([]string{"CreateAopWorkflowVersionApprovelInfo", string(data)}, " ")
}

type CreateAopWorkflowVersionApprovelInfoResult struct {
	value string
}

type CreateAopWorkflowVersionApprovelInfoResultEnum struct {
	PASS    CreateAopWorkflowVersionApprovelInfoResult
	UN_PASS CreateAopWorkflowVersionApprovelInfoResult
}

func GetCreateAopWorkflowVersionApprovelInfoResultEnum() CreateAopWorkflowVersionApprovelInfoResultEnum {
	return CreateAopWorkflowVersionApprovelInfoResultEnum{
		PASS: CreateAopWorkflowVersionApprovelInfoResult{
			value: "PASS",
		},
		UN_PASS: CreateAopWorkflowVersionApprovelInfoResult{
			value: "UN_PASS",
		},
	}
}

func (c CreateAopWorkflowVersionApprovelInfoResult) Value() string {
	return c.value
}

func (c CreateAopWorkflowVersionApprovelInfoResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAopWorkflowVersionApprovelInfoResult) UnmarshalJSON(b []byte) error {
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
