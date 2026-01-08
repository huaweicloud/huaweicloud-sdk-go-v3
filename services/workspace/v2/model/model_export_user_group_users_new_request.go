package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExportUserGroupUsersNewRequest Request Object
type ExportUserGroupUsersNewRequest struct {

	// 桌面用户组ID。
	GroupId string `json:"group_id"`

	// 用户名支持模糊查询。
	UserName *string `json:"user_name,omitempty"`

	// 用户描述支持模糊查询。
	Description *string `json:"description,omitempty"`

	// 激活类型。 - USER_ACTIVATE：用户激活 - ADMIN_ACTIVATE：管理员激活
	ActiveType *ExportUserGroupUsersNewRequestActiveType `json:"active_type,omitempty"`

	// 用户组名。
	GroupName *string `json:"group_name,omitempty"`

	// 语言。 - zh_CN：中文 - en_US：英文
	Language *ExportUserGroupUsersNewRequestLanguage `json:"language,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ExportUserGroupUsersNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportUserGroupUsersNewRequest struct{}"
	}

	return strings.Join([]string{"ExportUserGroupUsersNewRequest", string(data)}, " ")
}

type ExportUserGroupUsersNewRequestActiveType struct {
	value string
}

type ExportUserGroupUsersNewRequestActiveTypeEnum struct {
	USER_ACTIVATE  ExportUserGroupUsersNewRequestActiveType
	ADMIN_ACTIVATE ExportUserGroupUsersNewRequestActiveType
}

func GetExportUserGroupUsersNewRequestActiveTypeEnum() ExportUserGroupUsersNewRequestActiveTypeEnum {
	return ExportUserGroupUsersNewRequestActiveTypeEnum{
		USER_ACTIVATE: ExportUserGroupUsersNewRequestActiveType{
			value: "USER_ACTIVATE",
		},
		ADMIN_ACTIVATE: ExportUserGroupUsersNewRequestActiveType{
			value: "ADMIN_ACTIVATE",
		},
	}
}

func (c ExportUserGroupUsersNewRequestActiveType) Value() string {
	return c.value
}

func (c ExportUserGroupUsersNewRequestActiveType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportUserGroupUsersNewRequestActiveType) UnmarshalJSON(b []byte) error {
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

type ExportUserGroupUsersNewRequestLanguage struct {
	value string
}

type ExportUserGroupUsersNewRequestLanguageEnum struct {
	ZH_CN ExportUserGroupUsersNewRequestLanguage
	EN_US ExportUserGroupUsersNewRequestLanguage
}

func GetExportUserGroupUsersNewRequestLanguageEnum() ExportUserGroupUsersNewRequestLanguageEnum {
	return ExportUserGroupUsersNewRequestLanguageEnum{
		ZH_CN: ExportUserGroupUsersNewRequestLanguage{
			value: "zh_CN",
		},
		EN_US: ExportUserGroupUsersNewRequestLanguage{
			value: "en_US",
		},
	}
}

func (c ExportUserGroupUsersNewRequestLanguage) Value() string {
	return c.value
}

func (c ExportUserGroupUsersNewRequestLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportUserGroupUsersNewRequestLanguage) UnmarshalJSON(b []byte) error {
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
