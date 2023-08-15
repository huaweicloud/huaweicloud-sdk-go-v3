package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommitRepoV2 struct {

	// 提交对应的SHA id
	Id *string `json:"id,omitempty"`

	// 提交对应的短SHA id
	ShortId *string `json:"short_id,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 提交标题
	Title *string `json:"title,omitempty"`

	// 父提交id
	ParentIds *[]string `json:"parent_ids,omitempty"`

	// 提交信息
	Message *string `json:"message,omitempty"`

	// 作者
	AuthorName *string `json:"author_name,omitempty"`

	// 提交作者
	CommitterName *string `json:"committer_name,omitempty"`

	// 提交时间
	CommittedDate *sdktime.SdkTime `json:"committed_date,omitempty"`
}

func (o CommitRepoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitRepoV2 struct{}"
	}

	return strings.Join([]string{"CommitRepoV2", string(data)}, " ")
}
