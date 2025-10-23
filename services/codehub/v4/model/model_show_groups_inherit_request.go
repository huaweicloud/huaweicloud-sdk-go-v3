package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowGroupsInheritRequest Request Object
type ShowGroupsInheritRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`

	// **参数解释：** 必填项，设置类型protected_branches保护分支 protected_tags保护tag push_rules推送规则 merge_requests合并请求 mr_branch_policies合并分支 reviews检视意见 e2e_settings e2e设置 webhook_settings hook设置 deploy_keys 部署key watermark水印 repository_settings仓库设置。
	SettingType ShowGroupsInheritRequestSettingType `json:"setting_type"`
}

func (o ShowGroupsInheritRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupsInheritRequest struct{}"
	}

	return strings.Join([]string{"ShowGroupsInheritRequest", string(data)}, " ")
}

type ShowGroupsInheritRequestSettingType struct {
	value string
}

type ShowGroupsInheritRequestSettingTypeEnum struct {
	PROTECTED_BRANCHES  ShowGroupsInheritRequestSettingType
	PROTECTED_TAGS      ShowGroupsInheritRequestSettingType
	PUSH_RULES          ShowGroupsInheritRequestSettingType
	MERGE_REQUESTS      ShowGroupsInheritRequestSettingType
	MR_BRANCH_POLICIES  ShowGroupsInheritRequestSettingType
	REVIEWS             ShowGroupsInheritRequestSettingType
	E2E_SETTINGS        ShowGroupsInheritRequestSettingType
	WEBHOOK_SETTINGS    ShowGroupsInheritRequestSettingType
	DEPLOY_KEYS         ShowGroupsInheritRequestSettingType
	WATERMARK           ShowGroupsInheritRequestSettingType
	REPOSITORY_SETTINGS ShowGroupsInheritRequestSettingType
}

func GetShowGroupsInheritRequestSettingTypeEnum() ShowGroupsInheritRequestSettingTypeEnum {
	return ShowGroupsInheritRequestSettingTypeEnum{
		PROTECTED_BRANCHES: ShowGroupsInheritRequestSettingType{
			value: "protected_branches",
		},
		PROTECTED_TAGS: ShowGroupsInheritRequestSettingType{
			value: "protected_tags",
		},
		PUSH_RULES: ShowGroupsInheritRequestSettingType{
			value: "push_rules",
		},
		MERGE_REQUESTS: ShowGroupsInheritRequestSettingType{
			value: "merge_requests",
		},
		MR_BRANCH_POLICIES: ShowGroupsInheritRequestSettingType{
			value: "mr_branch_policies",
		},
		REVIEWS: ShowGroupsInheritRequestSettingType{
			value: "reviews",
		},
		E2E_SETTINGS: ShowGroupsInheritRequestSettingType{
			value: "e2e_settings",
		},
		WEBHOOK_SETTINGS: ShowGroupsInheritRequestSettingType{
			value: "webhook_settings",
		},
		DEPLOY_KEYS: ShowGroupsInheritRequestSettingType{
			value: "deploy_keys",
		},
		WATERMARK: ShowGroupsInheritRequestSettingType{
			value: "watermark",
		},
		REPOSITORY_SETTINGS: ShowGroupsInheritRequestSettingType{
			value: "repository_settings",
		},
	}
}

func (c ShowGroupsInheritRequestSettingType) Value() string {
	return c.value
}

func (c ShowGroupsInheritRequestSettingType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowGroupsInheritRequestSettingType) UnmarshalJSON(b []byte) error {
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
