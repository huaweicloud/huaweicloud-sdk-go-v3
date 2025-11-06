package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowRepositoryInheritSettingSourceRequest Request Object
type ShowRepositoryInheritSettingSourceRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 设置名称。 **约束限制：** 不涉及。 **取值范围：** - protected_branches，保护分支。 - protected_tags，保护Tag。 **默认取值：** 不涉及。
	Name *ShowRepositoryInheritSettingSourceRequestName `json:"name,omitempty"`
}

func (o ShowRepositoryInheritSettingSourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryInheritSettingSourceRequest struct{}"
	}

	return strings.Join([]string{"ShowRepositoryInheritSettingSourceRequest", string(data)}, " ")
}

type ShowRepositoryInheritSettingSourceRequestName struct {
	value string
}

type ShowRepositoryInheritSettingSourceRequestNameEnum struct {
	PROTECTED_BRANCHES ShowRepositoryInheritSettingSourceRequestName
	PROTECTED_TAGS     ShowRepositoryInheritSettingSourceRequestName
}

func GetShowRepositoryInheritSettingSourceRequestNameEnum() ShowRepositoryInheritSettingSourceRequestNameEnum {
	return ShowRepositoryInheritSettingSourceRequestNameEnum{
		PROTECTED_BRANCHES: ShowRepositoryInheritSettingSourceRequestName{
			value: "protected_branches",
		},
		PROTECTED_TAGS: ShowRepositoryInheritSettingSourceRequestName{
			value: "protected_tags",
		},
	}
}

func (c ShowRepositoryInheritSettingSourceRequestName) Value() string {
	return c.value
}

func (c ShowRepositoryInheritSettingSourceRequestName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowRepositoryInheritSettingSourceRequestName) UnmarshalJSON(b []byte) error {
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
