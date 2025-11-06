package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LinkSettingDto **参数解释：** 可集成系统Link服务设置信息。
type LinkSettingDto struct {

	// **参数解释：** 是否开启集成Link服务 **取值范围：** true：开启集成Link服务，false：未开启集成Link服务。
	Active *bool `json:"active,omitempty"`

	// **参数解释：** Link服务的对接地址。
	Url *string `json:"url,omitempty"`

	// **参数解释：** Link服务的对接鉴权类型，ak_sk代表使用ak和sk来鉴权。
	AppAuthType *LinkSettingDtoAppAuthType `json:"app_auth_type,omitempty"`

	// **参数解释：** Link服务的对接鉴权ak。
	AppAk *string `json:"app_ak,omitempty"`

	// **参数解释：** Link服务的对接鉴权sk，作为返回值时若已配置则返回掩码，掩码格式为************。
	AppSk *string `json:"app_sk,omitempty"`

	// **参数解释：** 可关联类型列表，逗号分隔。
	Categories *string `json:"categories,omitempty"`

	// **参数解释：** 排除状态列表，逗号分隔。
	ExcludeStatuses *string `json:"exclude_statuses,omitempty"`
}

func (o LinkSettingDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LinkSettingDto struct{}"
	}

	return strings.Join([]string{"LinkSettingDto", string(data)}, " ")
}

type LinkSettingDtoAppAuthType struct {
	value string
}

type LinkSettingDtoAppAuthTypeEnum struct {
	AK_SK LinkSettingDtoAppAuthType
}

func GetLinkSettingDtoAppAuthTypeEnum() LinkSettingDtoAppAuthTypeEnum {
	return LinkSettingDtoAppAuthTypeEnum{
		AK_SK: LinkSettingDtoAppAuthType{
			value: "ak_sk",
		},
	}
}

func (c LinkSettingDtoAppAuthType) Value() string {
	return c.value
}

func (c LinkSettingDtoAppAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LinkSettingDtoAppAuthType) UnmarshalJSON(b []byte) error {
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
