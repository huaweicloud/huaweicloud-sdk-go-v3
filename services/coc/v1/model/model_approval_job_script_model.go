package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ApprovalJobScriptModel 审批脚本  comments 长度未做限制
type ApprovalJobScriptModel struct {

	// 审批状态 APPROVED:审批通过 REJECTED:审批不通过
	Status ApprovalJobScriptModelStatus `json:"status"`

	// 审批意见
	Comments *string `json:"comments,omitempty"`
}

func (o ApprovalJobScriptModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApprovalJobScriptModel struct{}"
	}

	return strings.Join([]string{"ApprovalJobScriptModel", string(data)}, " ")
}

type ApprovalJobScriptModelStatus struct {
	value string
}

type ApprovalJobScriptModelStatusEnum struct {
	APPROVED ApprovalJobScriptModelStatus
	REJECTED ApprovalJobScriptModelStatus
}

func GetApprovalJobScriptModelStatusEnum() ApprovalJobScriptModelStatusEnum {
	return ApprovalJobScriptModelStatusEnum{
		APPROVED: ApprovalJobScriptModelStatus{
			value: "APPROVED",
		},
		REJECTED: ApprovalJobScriptModelStatus{
			value: "REJECTED",
		},
	}
}

func (c ApprovalJobScriptModelStatus) Value() string {
	return c.value
}

func (c ApprovalJobScriptModelStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApprovalJobScriptModelStatus) UnmarshalJSON(b []byte) error {
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
