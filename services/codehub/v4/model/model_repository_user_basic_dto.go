package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RepositoryUserBasicDto struct {

	// **参数解释：**  用户id **约束限制：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：**  用户名称 **约束限制：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：**  用户名 **约束限制：** 不涉及。
	Username *string `json:"username,omitempty"`

	// **参数解释：** 用户状态。 **取值范围：** - active，激活。 - blocked，禁用。
	State *RepositoryUserBasicDtoState `json:"state,omitempty"`

	// **参数解释：**  服务级权限状态 0：停用 1：启用 **约束限制：** 不涉及。
	ServiceLicenseStatus *int32 `json:"service_license_status,omitempty"`

	// **参数解释：**  用户中文名称 **约束限制：** 不涉及。
	NameCn *string `json:"name_cn,omitempty"`

	// **参数解释：**  用户昵称 **约束限制：** 不涉及。
	NickName *string `json:"nick_name,omitempty"`

	// **参数解释：**  租户名称 **约束限制：** 不涉及。
	TenantName *string `json:"tenant_name,omitempty"`
}

func (o RepositoryUserBasicDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryUserBasicDto struct{}"
	}

	return strings.Join([]string{"RepositoryUserBasicDto", string(data)}, " ")
}

type RepositoryUserBasicDtoState struct {
	value string
}

type RepositoryUserBasicDtoStateEnum struct {
	ACTIVE  RepositoryUserBasicDtoState
	BLOCKED RepositoryUserBasicDtoState
}

func GetRepositoryUserBasicDtoStateEnum() RepositoryUserBasicDtoStateEnum {
	return RepositoryUserBasicDtoStateEnum{
		ACTIVE: RepositoryUserBasicDtoState{
			value: "active",
		},
		BLOCKED: RepositoryUserBasicDtoState{
			value: "blocked",
		},
	}
}

func (c RepositoryUserBasicDtoState) Value() string {
	return c.value
}

func (c RepositoryUserBasicDtoState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RepositoryUserBasicDtoState) UnmarshalJSON(b []byte) error {
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
