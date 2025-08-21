package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateFileBodyDto struct {

	// 用户名
	Name *string `json:"name,omitempty"`

	// 文件路径
	FilePath string `json:"file_path"`

	// 分支名
	Branch string `json:"branch"`

	// 提交信息
	CommitMessage string `json:"commit_message"`

	// 作者邮箱
	AuthorEmail *string `json:"author_email,omitempty"`

	// 作者名称
	AuthorName *string `json:"author_name,omitempty"`

	// 文件内容
	Content string `json:"content"`

	// 编码方式
	Encoding *string `json:"encoding,omitempty"`

	// 最新提交commit
	LastCommitId *string `json:"last_commit_id,omitempty"`
}

func (o UpdateFileBodyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFileBodyDto struct{}"
	}

	return strings.Join([]string{"UpdateFileBodyDto", string(data)}, " ")
}
