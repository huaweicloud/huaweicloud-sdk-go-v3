package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// WorkflowApproveOpinionDetail 审核流程的详情信息
type WorkflowApproveOpinionDetail struct {

	// **参数解释**: 流程的审核结果 - PASS 通过 - UN_PASS 拒绝  **约束限制**: 不涉及         **取值范围**: - PASS - UN_PASS  **默认值**:  不涉及
	Result *WorkflowApproveOpinionDetailResult `json:"result,omitempty"`

	// **参数解释**: 流程的审核意见 **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	Content *string `json:"content,omitempty"`
}

func (o WorkflowApproveOpinionDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowApproveOpinionDetail struct{}"
	}

	return strings.Join([]string{"WorkflowApproveOpinionDetail", string(data)}, " ")
}

type WorkflowApproveOpinionDetailResult struct {
	value string
}

type WorkflowApproveOpinionDetailResultEnum struct {
	PASS    WorkflowApproveOpinionDetailResult
	UN_PASS WorkflowApproveOpinionDetailResult
}

func GetWorkflowApproveOpinionDetailResultEnum() WorkflowApproveOpinionDetailResultEnum {
	return WorkflowApproveOpinionDetailResultEnum{
		PASS: WorkflowApproveOpinionDetailResult{
			value: "PASS",
		},
		UN_PASS: WorkflowApproveOpinionDetailResult{
			value: "UN_PASS",
		},
	}
}

func (c WorkflowApproveOpinionDetailResult) Value() string {
	return c.value
}

func (c WorkflowApproveOpinionDetailResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WorkflowApproveOpinionDetailResult) UnmarshalJSON(b []byte) error {
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
