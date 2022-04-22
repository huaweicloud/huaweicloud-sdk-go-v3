package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommitAction struct {

	// 要执行的操作：create、delete、move、update、chmod
	Action string `json:"action"`

	// 文件的完整路径。例如lib/class.rb
	FilePath string `json:"file_path"`

	// 要移动的文件的原始完整路径。例如lib/class1.rb。仅在move操作时生效
	PreviousPath *string `json:"previous_path,omitempty"`

	// 文件内容，create和update操作时必须。未指定内容的移动操作将保留现有文件内容，内容的任何其他值将覆盖文件内容。
	Content *string `json:"content,omitempty"`

	// 文件编码：text、base64。默认为text
	Encoding *string `json:"encoding,omitempty"`

	// 最后一个已知的提交ID。仅在update、move、delete操作时生效
	LastCommitId *string `json:"last_commit_id,omitempty"`

	// 启用或者禁用文件的执行模式。仅在chmod操作时生效
	ExecuteFilemode *bool `json:"execute_filemode,omitempty"`
}

func (o CommitAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitAction struct{}"
	}

	return strings.Join([]string{"CommitAction", string(data)}, " ")
}
