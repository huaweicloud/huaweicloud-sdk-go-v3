package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MergeRequestBasicDiscussionDto struct {

	// 评论id
	Id *string `json:"id,omitempty"`

	// individual_note
	IndividualNote *bool `json:"individual_note,omitempty"`

	// 主评和回复列表
	Notes *[]NoteDto `json:"notes,omitempty"`

	// 仓库id
	ProjectId *int32 `json:"project_id,omitempty"`

	// 目标类型
	NoteableType *string `json:"noteable_type,omitempty"`

	// 关联的提交id
	CommitId *string `json:"commit_id,omitempty"`

	// 仓库路径
	ProjectFullPath *string `json:"project_full_path,omitempty"`

	// 变更前文件模式
	AMode *string `json:"a_mode,omitempty"`

	// 变更后文件模式
	BMode *string `json:"b_mode,omitempty"`

	// 此次变更是否删除文件
	DeletedFile *bool `json:"deleted_file,omitempty"`

	// 此次变更是否新增文件
	NewFile *bool `json:"new_file,omitempty"`

	// 检视意见是否解决
	Resolved *bool `json:"resolved,omitempty"`

	// 检视意见是否存档
	Archived *bool `json:"archived,omitempty"`

	// 检视意见分类
	ReviewCategories *string `json:"review_categories,omitempty"`

	// 检视意见分类中文名
	ReviewCategoriesCn *string `json:"review_categories_cn,omitempty"`

	// 检视意见分类英文名
	ReviewCategoriesEn *string `json:"review_categories_en,omitempty"`

	// 检视意见模块
	ReviewModules *string `json:"review_modules,omitempty"`

	// 严重程度
	Severity *string `json:"severity,omitempty"`

	// 严重程度中文名
	SeverityCn *string `json:"severity_cn,omitempty"`

	// 严重程度英文名
	SeverityEn *string `json:"severity_en,omitempty"`

	Assignee *UserBasicDto `json:"assignee,omitempty"`

	Proposer *UserBasicDto `json:"proposer,omitempty"`

	MergeRequestVersionParams *MergeRequestVersionParamsDto `json:"merge_request_version_params,omitempty"`
}

func (o MergeRequestBasicDiscussionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeRequestBasicDiscussionDto struct{}"
	}

	return strings.Join([]string{"MergeRequestBasicDiscussionDto", string(data)}, " ")
}
