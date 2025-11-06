package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ImportMrParamsDto struct {

	// 合并请求iid
	Iid int32 `json:"iid"`

	// 导入唯一标识
	SourceUniqKey string `json:"source_uniq_key"`

	// 作者id
	AuthorId *int32 `json:"author_id,omitempty"`

	// 状态
	State ImportMrParamsDtoState `json:"state"`

	// 标题
	Title *string `json:"title,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 源分支
	SourceBranch string `json:"source_branch"`

	// 目标分支
	TargetBranch string `json:"target_branch"`

	// 目标仓库
	TargetRepositoryId string `json:"target_repository_id"`

	// 标签
	Labels *interface{} `json:"labels,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 合并时间
	MergedAt *string `json:"merged_at,omitempty"`

	// 关闭时间
	ClosedAt *string `json:"closed_at,omitempty"`

	// 审核人列表
	Approvers *[]ApproverParamDto `json:"approvers,omitempty"`

	// 合并请求变更commit列表
	DiffRefs []DiffRefsParamDto `json:"diff_refs"`

	// squash合并
	Squash *bool `json:"squash,omitempty"`

	// 合并mr后删除源分支
	RemoveSourceBranch *bool `json:"remove_source_branch,omitempty"`

	// 源分支是否被删除
	BranchIsDeleted *bool `json:"branch_is_deleted,omitempty"`

	// 源合并请求是否是fork合并请求
	Fork *bool `json:"fork,omitempty"`

	// 导入来源
	ImportSourceFrom *string `json:"import_source_from,omitempty"`
}

func (o ImportMrParamsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportMrParamsDto struct{}"
	}

	return strings.Join([]string{"ImportMrParamsDto", string(data)}, " ")
}

type ImportMrParamsDtoState struct {
	value string
}

type ImportMrParamsDtoStateEnum struct {
	OPENED ImportMrParamsDtoState
	CLOSED ImportMrParamsDtoState
	MERGED ImportMrParamsDtoState
}

func GetImportMrParamsDtoStateEnum() ImportMrParamsDtoStateEnum {
	return ImportMrParamsDtoStateEnum{
		OPENED: ImportMrParamsDtoState{
			value: "opened",
		},
		CLOSED: ImportMrParamsDtoState{
			value: "closed",
		},
		MERGED: ImportMrParamsDtoState{
			value: "merged",
		},
	}
}

func (c ImportMrParamsDtoState) Value() string {
	return c.value
}

func (c ImportMrParamsDtoState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImportMrParamsDtoState) UnmarshalJSON(b []byte) error {
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
