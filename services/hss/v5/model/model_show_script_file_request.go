package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowScriptFileRequest Request Object
type ShowScriptFileRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 类型
	Type ShowScriptFileRequestType `json:"type"`
}

func (o ShowScriptFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScriptFileRequest struct{}"
	}

	return strings.Join([]string{"ShowScriptFileRequest", string(data)}, " ")
}

type ShowScriptFileRequestType struct {
	value string
}

type ShowScriptFileRequestTypeEnum struct {
	AGENT_BATCH_INSTALL_WINDOWS                    ShowScriptFileRequestType
	AGENT_SINGLE_INSTALL_WINDOWS                   ShowScriptFileRequestType
	AGENT_THIRD_PARTY_CLOUD_BATCH_INSTALL_WINDOWS  ShowScriptFileRequestType
	AGENT_THIRD_PARTY_CLOUD_SINGLE_INSTALL_WINDOWS ShowScriptFileRequestType
	AGENT_OTHER_ACCOUNTS_BATCH_INSTALL_WINDOWS     ShowScriptFileRequestType
	AGENT_OTHER_ACCOUNTS_SINGLE_INSTALL_WINDOWS    ShowScriptFileRequestType
}

func GetShowScriptFileRequestTypeEnum() ShowScriptFileRequestTypeEnum {
	return ShowScriptFileRequestTypeEnum{
		AGENT_BATCH_INSTALL_WINDOWS: ShowScriptFileRequestType{
			value: "agent_batch_install_windows",
		},
		AGENT_SINGLE_INSTALL_WINDOWS: ShowScriptFileRequestType{
			value: "agent_single_install_windows",
		},
		AGENT_THIRD_PARTY_CLOUD_BATCH_INSTALL_WINDOWS: ShowScriptFileRequestType{
			value: "agent_third_party_cloud_batch_install_windows",
		},
		AGENT_THIRD_PARTY_CLOUD_SINGLE_INSTALL_WINDOWS: ShowScriptFileRequestType{
			value: "agent_third_party_cloud_single_install_windows",
		},
		AGENT_OTHER_ACCOUNTS_BATCH_INSTALL_WINDOWS: ShowScriptFileRequestType{
			value: "agent_other_accounts_batch_install_windows",
		},
		AGENT_OTHER_ACCOUNTS_SINGLE_INSTALL_WINDOWS: ShowScriptFileRequestType{
			value: "agent_other_accounts_single_install_windows",
		},
	}
}

func (c ShowScriptFileRequestType) Value() string {
	return c.value
}

func (c ShowScriptFileRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowScriptFileRequestType) UnmarshalJSON(b []byte) error {
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
