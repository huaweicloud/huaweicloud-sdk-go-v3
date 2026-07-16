package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateWorkspaceReq struct {

	// 授权用户列表，默认为空。需要与“auth_type”参数配合使用，且仅当授权类型为“INTERNAL”时才会生效。
	Grants *[]CreateWorkspaceReqGrants `json:"grants,omitempty"`

	// 授权类型。可选值有PUBLIC、PRIVATE、INTERNAL。默认值为PUBLIC。 - PUBLIC：租户内部公开访问。 - PRIVATE：仅创建者和主账号可访问。 - INTERNAL：创建者、主账号、指定IAM子账号可访问，需要与grants参数配合使用。
	AuthType *CreateWorkspaceReqAuthType `json:"auth_type,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 工作空间名称。长度限制为4-64字符[，支持中文、大小写字母、数字、中划线和下划线](tag:hc,hk)。同时'default'为系统预留的默认工作空间名称，用户无法自己创建名为'default'的工作空间。
	Name string `json:"name"`

	// 工作空间描述，默认为空。长度限制为0-256字符。
	Description *string `json:"description,omitempty"`
}

func (o CreateWorkspaceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkspaceReq struct{}"
	}

	return strings.Join([]string{"CreateWorkspaceReq", string(data)}, " ")
}

type CreateWorkspaceReqAuthType struct {
	value string
}

type CreateWorkspaceReqAuthTypeEnum struct {
	PUBLIC   CreateWorkspaceReqAuthType
	PRIVATE  CreateWorkspaceReqAuthType
	INTERNAL CreateWorkspaceReqAuthType
}

func GetCreateWorkspaceReqAuthTypeEnum() CreateWorkspaceReqAuthTypeEnum {
	return CreateWorkspaceReqAuthTypeEnum{
		PUBLIC: CreateWorkspaceReqAuthType{
			value: "PUBLIC",
		},
		PRIVATE: CreateWorkspaceReqAuthType{
			value: "PRIVATE",
		},
		INTERNAL: CreateWorkspaceReqAuthType{
			value: "INTERNAL",
		},
	}
}

func (c CreateWorkspaceReqAuthType) Value() string {
	return c.value
}

func (c CreateWorkspaceReqAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateWorkspaceReqAuthType) UnmarshalJSON(b []byte) error {
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
