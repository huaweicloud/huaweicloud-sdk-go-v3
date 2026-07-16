package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateWorkspaceReq struct {

	// 训练作业使用的数据集。不可与data_url或dataset_id/dataset_version_id同时使用。
	Grants *[]ViewWorkspaceResponseGrants `json:"grants,omitempty"`

	// 授权类型。可选值有PUBLIC、PRIVATE、INTERNAL。默认值为PUBLIC。 - PUBLIC：租户内部公开访问。 - PRIVATE：仅创建者和主账号可访问。 - INTERNAL：创建者、主账号、指定IAM子账号可访问，需要与grants参数配合使用。
	AuthType *UpdateWorkspaceReqAuthType `json:"auth_type,omitempty"`

	// 工作空间名称。长度限制为4-64字符[，支持中文、大小写字母、数字、中划线和下划线](tag:hc,hk)。同时'default'为系统预留的默认工作空间名称，用户无法自己创建名为'default'的工作空间。
	Name *string `json:"name,omitempty"`

	// 工作空间描述，默认为空。长度限制为0-256字符。
	Description *string `json:"description,omitempty"`
}

func (o UpdateWorkspaceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkspaceReq struct{}"
	}

	return strings.Join([]string{"UpdateWorkspaceReq", string(data)}, " ")
}

type UpdateWorkspaceReqAuthType struct {
	value string
}

type UpdateWorkspaceReqAuthTypeEnum struct {
	PUBLIC   UpdateWorkspaceReqAuthType
	PRIVATE  UpdateWorkspaceReqAuthType
	INTERNAL UpdateWorkspaceReqAuthType
}

func GetUpdateWorkspaceReqAuthTypeEnum() UpdateWorkspaceReqAuthTypeEnum {
	return UpdateWorkspaceReqAuthTypeEnum{
		PUBLIC: UpdateWorkspaceReqAuthType{
			value: "PUBLIC",
		},
		PRIVATE: UpdateWorkspaceReqAuthType{
			value: "PRIVATE",
		},
		INTERNAL: UpdateWorkspaceReqAuthType{
			value: "INTERNAL",
		},
	}
}

func (c UpdateWorkspaceReqAuthType) Value() string {
	return c.value
}

func (c UpdateWorkspaceReqAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateWorkspaceReqAuthType) UnmarshalJSON(b []byte) error {
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
