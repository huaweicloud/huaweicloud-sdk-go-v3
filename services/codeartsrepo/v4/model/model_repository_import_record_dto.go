package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RepositoryImportRecordDto struct {

	// **参数解释：** 主键ID。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 规则名称。 **约束限制：** 不涉及。 **取值范围：** - finished, 导入成功 - fail, 导入失败 - importing, 导入中 **默认取值：** 不涉及。
	State *RepositoryImportRecordDtoState `json:"state,omitempty"`

	Repository *RepositorySimpleDto `json:"repository,omitempty"`

	// **参数解释：** 源仓库路径。
	OriginFullName *string `json:"origin_full_name,omitempty"`

	// **参数解释：** 源仓库地址。
	SourceUrl *string `json:"source_url,omitempty"`

	// **参数解释：** 导入来源。 **取值范围：** - gitee - self_managed_gitlab - gitlab - github - git - svn - coding - bitbucket - gerrit - codeup
	SourceType *RepositoryImportRecordDtoSourceType `json:"source_type,omitempty"`

	// **参数解释：** 导入时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 导入完成时间。
	FinishedAt *string `json:"finished_at,omitempty"`

	// **参数解释：** 源仓库容量。
	RepositorySize *float64 `json:"repository_size,omitempty"`

	// **参数解释：** 失败原因。
	ErrorMessage *string `json:"error_message,omitempty"`

	// **参数解释：** 仓库路径。
	TargetFullName *string `json:"target_full_name,omitempty"`

	// **参数解释：** 项目ID。
	TargetProjectId *string `json:"target_project_id,omitempty"`
}

func (o RepositoryImportRecordDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryImportRecordDto struct{}"
	}

	return strings.Join([]string{"RepositoryImportRecordDto", string(data)}, " ")
}

type RepositoryImportRecordDtoState struct {
	value string
}

type RepositoryImportRecordDtoStateEnum struct {
	FINISHED  RepositoryImportRecordDtoState
	FAIL      RepositoryImportRecordDtoState
	IMPORTING RepositoryImportRecordDtoState
}

func GetRepositoryImportRecordDtoStateEnum() RepositoryImportRecordDtoStateEnum {
	return RepositoryImportRecordDtoStateEnum{
		FINISHED: RepositoryImportRecordDtoState{
			value: "finished, 导入成功",
		},
		FAIL: RepositoryImportRecordDtoState{
			value: "fail, 导入失败",
		},
		IMPORTING: RepositoryImportRecordDtoState{
			value: "importing, 导入中",
		},
	}
}

func (c RepositoryImportRecordDtoState) Value() string {
	return c.value
}

func (c RepositoryImportRecordDtoState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RepositoryImportRecordDtoState) UnmarshalJSON(b []byte) error {
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

type RepositoryImportRecordDtoSourceType struct {
	value string
}

type RepositoryImportRecordDtoSourceTypeEnum struct {
	GITEE               RepositoryImportRecordDtoSourceType
	SELF_MANAGED_GITLAB RepositoryImportRecordDtoSourceType
	GITLAB              RepositoryImportRecordDtoSourceType
	GITHUB              RepositoryImportRecordDtoSourceType
	GIT                 RepositoryImportRecordDtoSourceType
	SVN                 RepositoryImportRecordDtoSourceType
	CODING              RepositoryImportRecordDtoSourceType
	BITBUCKET           RepositoryImportRecordDtoSourceType
	GERRIT              RepositoryImportRecordDtoSourceType
	CODEUP              RepositoryImportRecordDtoSourceType
}

func GetRepositoryImportRecordDtoSourceTypeEnum() RepositoryImportRecordDtoSourceTypeEnum {
	return RepositoryImportRecordDtoSourceTypeEnum{
		GITEE: RepositoryImportRecordDtoSourceType{
			value: "gitee",
		},
		SELF_MANAGED_GITLAB: RepositoryImportRecordDtoSourceType{
			value: "self_managed_gitlab",
		},
		GITLAB: RepositoryImportRecordDtoSourceType{
			value: "gitlab",
		},
		GITHUB: RepositoryImportRecordDtoSourceType{
			value: "github",
		},
		GIT: RepositoryImportRecordDtoSourceType{
			value: "git",
		},
		SVN: RepositoryImportRecordDtoSourceType{
			value: "svn",
		},
		CODING: RepositoryImportRecordDtoSourceType{
			value: "coding",
		},
		BITBUCKET: RepositoryImportRecordDtoSourceType{
			value: "bitbucket",
		},
		GERRIT: RepositoryImportRecordDtoSourceType{
			value: "gerrit",
		},
		CODEUP: RepositoryImportRecordDtoSourceType{
			value: "codeup",
		},
	}
}

func (c RepositoryImportRecordDtoSourceType) Value() string {
	return c.value
}

func (c RepositoryImportRecordDtoSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RepositoryImportRecordDtoSourceType) UnmarshalJSON(b []byte) error {
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
