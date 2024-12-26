package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAgentInstallScriptRequest Request Object
type ListAgentInstallScriptRequest struct {

	// Region ID
	Region string `json:"region"`

	// 企业项目ID，查询所有企业项目时填写：all_granted_eps
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// os类型：Windows和Linux
	OsType *ListAgentInstallScriptRequestOsType `json:"os_type,omitempty"`

	// 系统架构：x86_64和aarch64；当os_type为Windows时，只能选择x86_64
	OsArch ListAgentInstallScriptRequestOsArch `json:"os_arch"`

	// 是否非华为云
	OutsideHost *bool `json:"outside_host,omitempty"`

	// 服务器组ID
	OutsideGroupId *string `json:"outside_group_id,omitempty"`

	// 是否批量安装
	BatchInstall *bool `json:"batch_install,omitempty"`

	// 类型：password和ssh_key
	Type *ListAgentInstallScriptRequestType `json:"type,omitempty"`
}

func (o ListAgentInstallScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgentInstallScriptRequest struct{}"
	}

	return strings.Join([]string{"ListAgentInstallScriptRequest", string(data)}, " ")
}

type ListAgentInstallScriptRequestOsType struct {
	value string
}

type ListAgentInstallScriptRequestOsTypeEnum struct {
	WINDOWS ListAgentInstallScriptRequestOsType
	LINUX   ListAgentInstallScriptRequestOsType
}

func GetListAgentInstallScriptRequestOsTypeEnum() ListAgentInstallScriptRequestOsTypeEnum {
	return ListAgentInstallScriptRequestOsTypeEnum{
		WINDOWS: ListAgentInstallScriptRequestOsType{
			value: "Windows",
		},
		LINUX: ListAgentInstallScriptRequestOsType{
			value: "Linux",
		},
	}
}

func (c ListAgentInstallScriptRequestOsType) Value() string {
	return c.value
}

func (c ListAgentInstallScriptRequestOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAgentInstallScriptRequestOsType) UnmarshalJSON(b []byte) error {
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

type ListAgentInstallScriptRequestOsArch struct {
	value string
}

type ListAgentInstallScriptRequestOsArchEnum struct {
	X86_64  ListAgentInstallScriptRequestOsArch
	AARCH64 ListAgentInstallScriptRequestOsArch
}

func GetListAgentInstallScriptRequestOsArchEnum() ListAgentInstallScriptRequestOsArchEnum {
	return ListAgentInstallScriptRequestOsArchEnum{
		X86_64: ListAgentInstallScriptRequestOsArch{
			value: "x86_64",
		},
		AARCH64: ListAgentInstallScriptRequestOsArch{
			value: "aarch64",
		},
	}
}

func (c ListAgentInstallScriptRequestOsArch) Value() string {
	return c.value
}

func (c ListAgentInstallScriptRequestOsArch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAgentInstallScriptRequestOsArch) UnmarshalJSON(b []byte) error {
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

type ListAgentInstallScriptRequestType struct {
	value string
}

type ListAgentInstallScriptRequestTypeEnum struct {
	PASSWORD ListAgentInstallScriptRequestType
	SSH_KEY  ListAgentInstallScriptRequestType
}

func GetListAgentInstallScriptRequestTypeEnum() ListAgentInstallScriptRequestTypeEnum {
	return ListAgentInstallScriptRequestTypeEnum{
		PASSWORD: ListAgentInstallScriptRequestType{
			value: "password",
		},
		SSH_KEY: ListAgentInstallScriptRequestType{
			value: "ssh_key",
		},
	}
}

func (c ListAgentInstallScriptRequestType) Value() string {
	return c.value
}

func (c ListAgentInstallScriptRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAgentInstallScriptRequestType) UnmarshalJSON(b []byte) error {
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
