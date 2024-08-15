package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ScriptListModel **遗留问题** 1. 脚本入参限制：参数个数、参数长度、参数合法范围 2. 字段长度和范围待定义
type ScriptListModel struct {

	// 脚本自增id
	Id int64 `json:"id"`

	// 脚本uuid
	ScriptUuid string `json:"script_uuid"`

	// 企业项目ID，默认为：0
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 脚本名称
	Name string `json:"name"`

	// 脚本类型 SHELL:shell脚本 PYTHON:python脚本 BAT:bat脚本
	Type ScriptListModelType `json:"type"`

	// 脚本状态 PENDING_APPROVE:待审批 APPROVED：正常（审批通过） REJECTED：驳回（审批人，驳回该脚本
	Status ScriptListModelStatus `json:"status"`

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
}

func (o ScriptListModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScriptListModel struct{}"
	}

	return strings.Join([]string{"ScriptListModel", string(data)}, " ")
}

type ScriptListModelType struct {
	value string
}

type ScriptListModelTypeEnum struct {
	SHELL  ScriptListModelType
	PYTHON ScriptListModelType
	BAT    ScriptListModelType
}

func GetScriptListModelTypeEnum() ScriptListModelTypeEnum {
	return ScriptListModelTypeEnum{
		SHELL: ScriptListModelType{
			value: "SHELL",
		},
		PYTHON: ScriptListModelType{
			value: "PYTHON",
		},
		BAT: ScriptListModelType{
			value: "BAT",
		},
	}
}

func (c ScriptListModelType) Value() string {
	return c.value
}

func (c ScriptListModelType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScriptListModelType) UnmarshalJSON(b []byte) error {
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

type ScriptListModelStatus struct {
	value string
}

type ScriptListModelStatusEnum struct {
	PENDING_APPROVE ScriptListModelStatus
	APPROVED        ScriptListModelStatus
	REJECTED        ScriptListModelStatus
}

func GetScriptListModelStatusEnum() ScriptListModelStatusEnum {
	return ScriptListModelStatusEnum{
		PENDING_APPROVE: ScriptListModelStatus{
			value: "PENDING_APPROVE",
		},
		APPROVED: ScriptListModelStatus{
			value: "APPROVED",
		},
		REJECTED: ScriptListModelStatus{
			value: "REJECTED",
		},
	}
}

func (c ScriptListModelStatus) Value() string {
	return c.value
}

func (c ScriptListModelStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScriptListModelStatus) UnmarshalJSON(b []byte) error {
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
