package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MergeRequestTemplateFromDto 合并请求模板来源
type MergeRequestTemplateFromDto struct {

	// **参数解释：** 设置来源的url链接，点击可跳转到项目、代码组或仓库的和并请求模板设置。
	Path *string `json:"path,omitempty"`

	// **参数解释：** repository: 设置来自于仓库 group: 设置继承自代码组 project: 设置继承自项目
	Type *MergeRequestTemplateFromDtoType `json:"type,omitempty"`

	// **参数解释：** 仓库id，不来源于仓库时为null。
	RepositoryId *int32 `json:"repository_id,omitempty"`

	// **参数解释：** 代码组id，不来源于代码组时为null。
	GroupId *int32 `json:"group_id,omitempty"`

	// **参数解释：** 项目id，不来源于项目时为null。
	ProjectId *string `json:"project_id,omitempty"`
}

func (o MergeRequestTemplateFromDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeRequestTemplateFromDto struct{}"
	}

	return strings.Join([]string{"MergeRequestTemplateFromDto", string(data)}, " ")
}

type MergeRequestTemplateFromDtoType struct {
	value string
}

type MergeRequestTemplateFromDtoTypeEnum struct {
	REPOSITORY MergeRequestTemplateFromDtoType
	GROUP      MergeRequestTemplateFromDtoType
	PROJECT    MergeRequestTemplateFromDtoType
}

func GetMergeRequestTemplateFromDtoTypeEnum() MergeRequestTemplateFromDtoTypeEnum {
	return MergeRequestTemplateFromDtoTypeEnum{
		REPOSITORY: MergeRequestTemplateFromDtoType{
			value: "repository",
		},
		GROUP: MergeRequestTemplateFromDtoType{
			value: "group",
		},
		PROJECT: MergeRequestTemplateFromDtoType{
			value: "project",
		},
	}
}

func (c MergeRequestTemplateFromDtoType) Value() string {
	return c.value
}

func (c MergeRequestTemplateFromDtoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MergeRequestTemplateFromDtoType) UnmarshalJSON(b []byte) error {
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
