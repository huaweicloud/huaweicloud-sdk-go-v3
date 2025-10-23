package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ReqSettingDto **参数解释：** 可集成系统CodeArts Req设置信息。
type ReqSettingDto struct {

	// **参数解释：** 是否开启集成CodeArts Req **取值范围：** true：开启集成CodeArts Req，false：未开启集成CodeArts Req。
	Active *bool `json:"active,omitempty"`

	// **参数解释：** 应用排除状态和可关联类别限制的分支列表，该分支指代合并请求的目标分支，可支持多个分支用英文逗号拼接或者正则表达式，在项目层级和代码组层级只支持配置正则表达式。
	Branches *string `json:"branches,omitempty"`

	// **参数解释：** 分支的类型，文本或者正则表达式。 **取值范围：** plain代表文本，regex代表正则表达式。
	BranchesType *ReqSettingDtoBranchesType `json:"branches_type,omitempty"`

	// **参数解释：** 关联的CodeArts Req项目类型。 **取值范围：** scrum代表scrum类型项目，ipd代表IPD类型项目, xboard为看板类型项目。
	ProjectType *ReqSettingDtoProjectType `json:"project_type,omitempty"`

	// **参数解释：** 可关联类型列表，逗号分隔。
	Categories *string `json:"categories,omitempty"`

	// **参数解释：** 可关联类型编码列表，逗号分隔。
	CategoryCodes *string `json:"category_codes,omitempty"`

	// **参数解释：** 排除状态列表，逗号分隔。
	ExcludeStatuses *string `json:"exclude_statuses,omitempty"`

	// **参数解释：** 排除状态编码列表，逗号分隔。
	ExcludeStatusCodes *string `json:"exclude_status_codes,omitempty"`
}

func (o ReqSettingDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReqSettingDto struct{}"
	}

	return strings.Join([]string{"ReqSettingDto", string(data)}, " ")
}

type ReqSettingDtoBranchesType struct {
	value string
}

type ReqSettingDtoBranchesTypeEnum struct {
	PLAIN ReqSettingDtoBranchesType
	REGEX ReqSettingDtoBranchesType
}

func GetReqSettingDtoBranchesTypeEnum() ReqSettingDtoBranchesTypeEnum {
	return ReqSettingDtoBranchesTypeEnum{
		PLAIN: ReqSettingDtoBranchesType{
			value: "plain",
		},
		REGEX: ReqSettingDtoBranchesType{
			value: "regex",
		},
	}
}

func (c ReqSettingDtoBranchesType) Value() string {
	return c.value
}

func (c ReqSettingDtoBranchesType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReqSettingDtoBranchesType) UnmarshalJSON(b []byte) error {
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

type ReqSettingDtoProjectType struct {
	value string
}

type ReqSettingDtoProjectTypeEnum struct {
	SCRUM  ReqSettingDtoProjectType
	IPD    ReqSettingDtoProjectType
	XBOARD ReqSettingDtoProjectType
}

func GetReqSettingDtoProjectTypeEnum() ReqSettingDtoProjectTypeEnum {
	return ReqSettingDtoProjectTypeEnum{
		SCRUM: ReqSettingDtoProjectType{
			value: "scrum",
		},
		IPD: ReqSettingDtoProjectType{
			value: "ipd",
		},
		XBOARD: ReqSettingDtoProjectType{
			value: "xboard",
		},
	}
}

func (c ReqSettingDtoProjectType) Value() string {
	return c.value
}

func (c ReqSettingDtoProjectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReqSettingDtoProjectType) UnmarshalJSON(b []byte) error {
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
