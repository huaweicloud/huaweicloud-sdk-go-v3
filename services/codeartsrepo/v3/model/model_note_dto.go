package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NoteDto struct {

	// note id
	Id *int32 `json:"id,omitempty"`

	// note类型
	Type *string `json:"type,omitempty"`

	// 检视意见内容
	Body *string `json:"body,omitempty"`

	// 附件
	Attachment *string `json:"attachment,omitempty"`

	Author *UserBasicDto `json:"author,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 是否是系统生成的日志
	System *bool `json:"system,omitempty"`

	// 目标id
	NoteableId *int32 `json:"noteable_id,omitempty"`

	// 目标类型
	NoteableType *string `json:"noteable_type,omitempty"`

	// 关联的提交id
	CommitId *string `json:"commit_id,omitempty"`

	// 是否可解决
	Resolvable *bool `json:"resolvable,omitempty"`

	// 是否是回复
	IsReply *bool `json:"is_reply,omitempty"`

	ResolvedBy *UserBasicDto `json:"resolved_by,omitempty"`

	// 目标iid
	NoteableIid *int32 `json:"noteable_iid,omitempty"`

	// 讨论id
	DiscussionId *string `json:"discussion_id,omitempty"`

	// 所属项目
	Project *string `json:"project,omitempty"`

	// 变更文件
	DiffFile *string `json:"diff_file,omitempty"`

	// 变更内容
	Diff *string `json:"diff,omitempty"`

	// 是否存档
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

	// 文件路径
	FilePath *string `json:"file_path,omitempty"`

	// 行号
	Line *string `json:"line,omitempty"`

	Assignee *UserBasicDto `json:"assignee,omitempty"`

	Proposer *UserBasicDto `json:"proposer,omitempty"`

	Position *PositionDto `json:"position,omitempty"`

	// 是否解决
	Resolved *bool `json:"resolved,omitempty"`

	// 是否过时
	IsOutdated *bool `json:"is_outdated,omitempty"`
}

func (o NoteDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NoteDto struct{}"
	}

	return strings.Join([]string{"NoteDto", string(data)}, " ")
}
