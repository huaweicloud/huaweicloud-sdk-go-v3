package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ImportFileReq struct {
	Path *string `json:"path,omitempty"`

	// 公共作业参数
	Params *interface{} `json:"params,omitempty"`

	SameNamePolicy *ImportFileReqSameNamePolicy `json:"sameNamePolicy,omitempty"`

	// 指定作业参数
	JobsParam *interface{} `json:"jobsParam,omitempty"`

	ExecuteUser *string `json:"executeUser,omitempty"`

	// 在开启审批开关后，需要填写该字段。表示创建作业的目标状态，有三种状态：SAVED、SUBMITTED和PRODUCTION，分别表示作业创建后是保存态，提交态，生产态
	TargetStatus *ImportFileReqTargetStatus `json:"targetStatus,omitempty"`

	// 在开启审批开关后，需要填写该字段，表示作业审批人
	Approvers *[]JobApprover `json:"approvers,omitempty"`

	// 如需替换资源，需要填写该字段，包含替换的资源名和资源类型和替换后的资源名
	Resources *[]JobResourceInfo `json:"resources,omitempty"`
}

func (o ImportFileReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportFileReq struct{}"
	}

	return strings.Join([]string{"ImportFileReq", string(data)}, " ")
}

type ImportFileReqSameNamePolicy struct {
	value string
}

type ImportFileReqSameNamePolicyEnum struct {
	SKIP      ImportFileReqSameNamePolicy
	OVERWRITE ImportFileReqSameNamePolicy
}

func GetImportFileReqSameNamePolicyEnum() ImportFileReqSameNamePolicyEnum {
	return ImportFileReqSameNamePolicyEnum{
		SKIP: ImportFileReqSameNamePolicy{
			value: "SKIP",
		},
		OVERWRITE: ImportFileReqSameNamePolicy{
			value: "OVERWRITE",
		},
	}
}

func (c ImportFileReqSameNamePolicy) Value() string {
	return c.value
}

func (c ImportFileReqSameNamePolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImportFileReqSameNamePolicy) UnmarshalJSON(b []byte) error {
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

type ImportFileReqTargetStatus struct {
	value string
}

type ImportFileReqTargetStatusEnum struct {
	SAVED      ImportFileReqTargetStatus
	SUBMITTED  ImportFileReqTargetStatus
	PRODUCTION ImportFileReqTargetStatus
}

func GetImportFileReqTargetStatusEnum() ImportFileReqTargetStatusEnum {
	return ImportFileReqTargetStatusEnum{
		SAVED: ImportFileReqTargetStatus{
			value: "SAVED",
		},
		SUBMITTED: ImportFileReqTargetStatus{
			value: "SUBMITTED",
		},
		PRODUCTION: ImportFileReqTargetStatus{
			value: "PRODUCTION",
		},
	}
}

func (c ImportFileReqTargetStatus) Value() string {
	return c.value
}

func (c ImportFileReqTargetStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImportFileReqTargetStatus) UnmarshalJSON(b []byte) error {
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
