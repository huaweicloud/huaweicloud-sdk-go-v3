package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ScriptDetailModel 脚本详情
type ScriptDetailModel struct {

	// 脚本uuid
	ScriptUuid string `json:"script_uuid"`

	// 脚本名称
	Name string `json:"name"`

	// 脚本版本号 约束条件  后期废除，不建议使用
	Version *string `json:"version,omitempty"`

	// 脚本描述
	Description string `json:"description"`

	// 脚本类型 SHELL:shell脚本， PYTHON:Python脚本， BAT:Bat脚本，
	Type ScriptDetailModelType `json:"type"`

	// 脚本内容
	Content string `json:"content"`

	// 脚本入参
	ScriptParams *[]ScriptParamDefine `json:"script_params,omitempty"`

	// 脚本状态 PENDING_APPROVE:待审批 APPROVED：正常（审批通过） REJECTED：驳回（审批人，驳回该脚本）
	Status ScriptDetailModelStatus `json:"status"`

	// 创建时间
	GmtCreated int64 `json:"gmt_created"`

	// 修改时间
	GmtModified *int64 `json:"gmt_modified,omitempty"`

	// 创建人
	Creator string `json:"creator"`

	// 创建人Id
	CreatorId string `json:"creator_id"`

	// 修改人
	Operator *string `json:"operator,omitempty"`

	Properties *ScriptPropertiesModel `json:"properties"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ScriptDetailModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScriptDetailModel struct{}"
	}

	return strings.Join([]string{"ScriptDetailModel", string(data)}, " ")
}

type ScriptDetailModelType struct {
	value string
}

type ScriptDetailModelTypeEnum struct {
	SHELL  ScriptDetailModelType
	PYTHON ScriptDetailModelType
	BAT    ScriptDetailModelType
}

func GetScriptDetailModelTypeEnum() ScriptDetailModelTypeEnum {
	return ScriptDetailModelTypeEnum{
		SHELL: ScriptDetailModelType{
			value: "SHELL",
		},
		PYTHON: ScriptDetailModelType{
			value: "PYTHON",
		},
		BAT: ScriptDetailModelType{
			value: "BAT",
		},
	}
}

func (c ScriptDetailModelType) Value() string {
	return c.value
}

func (c ScriptDetailModelType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScriptDetailModelType) UnmarshalJSON(b []byte) error {
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

type ScriptDetailModelStatus struct {
	value string
}

type ScriptDetailModelStatusEnum struct {
	PENDING_APPROVE ScriptDetailModelStatus
	APPROVED        ScriptDetailModelStatus
	REJECTED        ScriptDetailModelStatus
}

func GetScriptDetailModelStatusEnum() ScriptDetailModelStatusEnum {
	return ScriptDetailModelStatusEnum{
		PENDING_APPROVE: ScriptDetailModelStatus{
			value: "PENDING_APPROVE",
		},
		APPROVED: ScriptDetailModelStatus{
			value: "APPROVED",
		},
		REJECTED: ScriptDetailModelStatus{
			value: "REJECTED",
		},
	}
}

func (c ScriptDetailModelStatus) Value() string {
	return c.value
}

func (c ScriptDetailModelStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScriptDetailModelStatus) UnmarshalJSON(b []byte) error {
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
