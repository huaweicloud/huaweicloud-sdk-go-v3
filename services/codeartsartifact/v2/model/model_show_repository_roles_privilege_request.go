package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowRepositoryRolesPrivilegeRequest Request Object
type ShowRepositoryRolesPrivilegeRequest struct {

	// **参数解释**: 项目ID，可以从调用API处获取，也可以从控制台获取。获取方式请参考[获取项目ID](CloudArtifact_api_0015.xml)。 **约束限制**: 只能由英文字母、数字组成，且长度为32个字符。 **取值范围**: 不涉及。 **默认取值**: 无。
	ProjectId string `json:"project_id"`

	// **参数解释**: 仓库id。可在私有库仓库**概览**界面查看。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 无。
	RepoId string `json:"repo_id"`

	// **参数解释**: 语言类型。 **约束限制**: 不涉及。 **取值范围**: 可选值：zh-cn,en-us。 **默认取值**: zh-cn。
	XLanguage *ShowRepositoryRolesPrivilegeRequestXLanguage `json:"x-language,omitempty"`
}

func (o ShowRepositoryRolesPrivilegeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryRolesPrivilegeRequest struct{}"
	}

	return strings.Join([]string{"ShowRepositoryRolesPrivilegeRequest", string(data)}, " ")
}

type ShowRepositoryRolesPrivilegeRequestXLanguage struct {
	value string
}

type ShowRepositoryRolesPrivilegeRequestXLanguageEnum struct {
	ZH_CN ShowRepositoryRolesPrivilegeRequestXLanguage
	EN_US ShowRepositoryRolesPrivilegeRequestXLanguage
}

func GetShowRepositoryRolesPrivilegeRequestXLanguageEnum() ShowRepositoryRolesPrivilegeRequestXLanguageEnum {
	return ShowRepositoryRolesPrivilegeRequestXLanguageEnum{
		ZH_CN: ShowRepositoryRolesPrivilegeRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowRepositoryRolesPrivilegeRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowRepositoryRolesPrivilegeRequestXLanguage) Value() string {
	return c.value
}

func (c ShowRepositoryRolesPrivilegeRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowRepositoryRolesPrivilegeRequestXLanguage) UnmarshalJSON(b []byte) error {
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
