package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type WebHookBranchCfgDto struct {

	// **参数解释：** 分支类型。 **取值范围：** - 0，文本。 - 1，通配符。 - 2，正则。
	BranchType *WebHookBranchCfgDtoBranchType `json:"branch_type,omitempty"`

	// **参数解释：** 分支名配置。 **取值范围：** 最小1个字节，最大255字节
	Branch *string `json:"branch,omitempty"`

	// **参数解释：** 仓库名类型。 **取值范围：** - 0，文本。 - 1，通配符。 - 2，正则。
	ProjectType *WebHookBranchCfgDtoProjectType `json:"project_type,omitempty"`

	// **参数解释：** 仓库名配置。 **取值范围：** 最小1个字节，最大255字节
	Project *string `json:"project,omitempty"`
}

func (o WebHookBranchCfgDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebHookBranchCfgDto struct{}"
	}

	return strings.Join([]string{"WebHookBranchCfgDto", string(data)}, " ")
}

type WebHookBranchCfgDtoBranchType struct {
	value int32
}

type WebHookBranchCfgDtoBranchTypeEnum struct {
	E_0 WebHookBranchCfgDtoBranchType
	E_1 WebHookBranchCfgDtoBranchType
	E_2 WebHookBranchCfgDtoBranchType
}

func GetWebHookBranchCfgDtoBranchTypeEnum() WebHookBranchCfgDtoBranchTypeEnum {
	return WebHookBranchCfgDtoBranchTypeEnum{
		E_0: WebHookBranchCfgDtoBranchType{
			value: 0,
		}, E_1: WebHookBranchCfgDtoBranchType{
			value: 1,
		}, E_2: WebHookBranchCfgDtoBranchType{
			value: 2,
		},
	}
}

func (c WebHookBranchCfgDtoBranchType) Value() int32 {
	return c.value
}

func (c WebHookBranchCfgDtoBranchType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WebHookBranchCfgDtoBranchType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type WebHookBranchCfgDtoProjectType struct {
	value int32
}

type WebHookBranchCfgDtoProjectTypeEnum struct {
	E_0 WebHookBranchCfgDtoProjectType
	E_1 WebHookBranchCfgDtoProjectType
	E_2 WebHookBranchCfgDtoProjectType
}

func GetWebHookBranchCfgDtoProjectTypeEnum() WebHookBranchCfgDtoProjectTypeEnum {
	return WebHookBranchCfgDtoProjectTypeEnum{
		E_0: WebHookBranchCfgDtoProjectType{
			value: 0,
		}, E_1: WebHookBranchCfgDtoProjectType{
			value: 1,
		}, E_2: WebHookBranchCfgDtoProjectType{
			value: 2,
		},
	}
}

func (c WebHookBranchCfgDtoProjectType) Value() int32 {
	return c.value
}

func (c WebHookBranchCfgDtoProjectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WebHookBranchCfgDtoProjectType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
