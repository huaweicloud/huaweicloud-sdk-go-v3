package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ActionOnComponentSource 源信息。
type ActionOnComponentSource struct {
	Code *Repo `json:"code,omitempty"`

	// 源类型。
	Type *ActionOnComponentSourceType `json:"type,omitempty"`

	// 源子类型。 - 源类型为code时，子类型表示不同的代码仓库，如DevCloud（CodeArts)、GitLab、GitHub、Gitee、Bitbucket。 - 源类型为softwarePackage时，子类型表示不同的软件包仓库，如BinObs、BinDevCloud（CodeArts软件发布库)。
	SubType *ActionOnComponentSourceSubType `json:"sub_type,omitempty"`

	// url地址。 - 源类型为image时，url地址为镜像地址。 - 源类型为code时，url地址为git地址。 - 源类型为softwarePackage时，url地址为软件包地址。
	Url *string `json:"url,omitempty"`
}

func (o ActionOnComponentSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionOnComponentSource struct{}"
	}

	return strings.Join([]string{"ActionOnComponentSource", string(data)}, " ")
}

type ActionOnComponentSourceType struct {
	value string
}

type ActionOnComponentSourceTypeEnum struct {
	IMAGE            ActionOnComponentSourceType
	CODE             ActionOnComponentSourceType
	SOFTWARE_PACKAGE ActionOnComponentSourceType
}

func GetActionOnComponentSourceTypeEnum() ActionOnComponentSourceTypeEnum {
	return ActionOnComponentSourceTypeEnum{
		IMAGE: ActionOnComponentSourceType{
			value: "image",
		},
		CODE: ActionOnComponentSourceType{
			value: "code",
		},
		SOFTWARE_PACKAGE: ActionOnComponentSourceType{
			value: "softwarePackage",
		},
	}
}

func (c ActionOnComponentSourceType) Value() string {
	return c.value
}

func (c ActionOnComponentSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ActionOnComponentSourceType) UnmarshalJSON(b []byte) error {
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

type ActionOnComponentSourceSubType struct {
	value string
}

type ActionOnComponentSourceSubTypeEnum struct {
	BIN_OBS       ActionOnComponentSourceSubType
	BIN_DEV_CLOUD ActionOnComponentSourceSubType
	GIT_LAB       ActionOnComponentSourceSubType
	GIT_HUB       ActionOnComponentSourceSubType
	DEV_CLOUD     ActionOnComponentSourceSubType
	GITEE         ActionOnComponentSourceSubType
	BITBUCKET     ActionOnComponentSourceSubType
}

func GetActionOnComponentSourceSubTypeEnum() ActionOnComponentSourceSubTypeEnum {
	return ActionOnComponentSourceSubTypeEnum{
		BIN_OBS: ActionOnComponentSourceSubType{
			value: "BinObs",
		},
		BIN_DEV_CLOUD: ActionOnComponentSourceSubType{
			value: "BinDevCloud",
		},
		GIT_LAB: ActionOnComponentSourceSubType{
			value: "GitLab",
		},
		GIT_HUB: ActionOnComponentSourceSubType{
			value: "GitHub",
		},
		DEV_CLOUD: ActionOnComponentSourceSubType{
			value: "DevCloud",
		},
		GITEE: ActionOnComponentSourceSubType{
			value: "Gitee",
		},
		BITBUCKET: ActionOnComponentSourceSubType{
			value: "Bitbucket",
		},
	}
}

func (c ActionOnComponentSourceSubType) Value() string {
	return c.value
}

func (c ActionOnComponentSourceSubType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ActionOnComponentSourceSubType) UnmarshalJSON(b []byte) error {
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
