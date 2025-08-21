package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MergeRequestStatisticDto 合并请求统计数据
type MergeRequestStatisticDto struct {

	// 合并请求ID
	Id *int32 `json:"id,omitempty"`

	// 合并请求序号
	Iid *int32 `json:"iid,omitempty"`

	// 合并请求标题
	Title *string `json:"title,omitempty"`

	// 合并请求状态
	State *string `json:"state,omitempty"`

	// 合并请求提交数
	CommitsCount *int32 `json:"commits_count,omitempty"`

	// 合并请求文件变数
	ChangedFilesCount *string `json:"changed_files_count,omitempty"`

	NotesCount *NotesCountDto `json:"notes_count,omitempty"`

	ChangedLinesCount *MergeRequestLineChange `json:"changed_lines_count,omitempty"`

	// 合并请求合入异常信息
	MergeError *string `json:"merge_error,omitempty"`

	// 合并请求合入异常信息
	JsonMergeError *interface{} `json:"json_merge_error,omitempty"`

	// 合并请求评分数
	Votes *int32 `json:"votes,omitempty"`
}

func (o MergeRequestStatisticDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeRequestStatisticDto struct{}"
	}

	return strings.Join([]string{"MergeRequestStatisticDto", string(data)}, " ")
}
