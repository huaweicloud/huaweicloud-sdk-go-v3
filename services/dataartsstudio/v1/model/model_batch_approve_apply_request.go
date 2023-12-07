package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchApproveApplyRequest Request Object
type BatchApproveApplyRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// dlm版本类型
	DlmType *BatchApproveApplyRequestDlmType `json:"Dlm-Type,omitempty"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	Body *OpenApplyIdsForApproveApply `json:"body,omitempty"`
}

func (o BatchApproveApplyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchApproveApplyRequest struct{}"
	}

	return strings.Join([]string{"BatchApproveApplyRequest", string(data)}, " ")
}

type BatchApproveApplyRequestDlmType struct {
	value string
}

type BatchApproveApplyRequestDlmTypeEnum struct {
	SHARED    BatchApproveApplyRequestDlmType
	EXCLUSIVE BatchApproveApplyRequestDlmType
}

func GetBatchApproveApplyRequestDlmTypeEnum() BatchApproveApplyRequestDlmTypeEnum {
	return BatchApproveApplyRequestDlmTypeEnum{
		SHARED: BatchApproveApplyRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: BatchApproveApplyRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c BatchApproveApplyRequestDlmType) Value() string {
	return c.value
}

func (c BatchApproveApplyRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchApproveApplyRequestDlmType) UnmarshalJSON(b []byte) error {
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
