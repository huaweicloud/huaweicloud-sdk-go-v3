package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateCommitResponseBody struct {

	// 提交对应的SHA id
	Id *string `json:"id,omitempty"`

	// 提交对应的短SHA id
	ShortId *string `json:"short_id,omitempty"`

	// 提交标题
	Title *string `json:"title,omitempty"`

	// 作者
	AuthorName *string `json:"author_name,omitempty"`

	// 作者邮箱
	AuthorEmail *string `json:"author_email,omitempty"`

	// 提交作者
	CommitterName *string `json:"committer_name,omitempty"`

	// 提交作者邮箱
	CommitterEmail *string `json:"committer_email,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 提交信息
	Message *string `json:"message,omitempty"`

	// 父提交id
	ParentIds *[]string `json:"parent_ids,omitempty"`

	// 提交时间
	CommittedDate *sdktime.SdkTime `json:"committed_date,omitempty"`

	// 作者提交时间
	AuthoredDate *sdktime.SdkTime `json:"authored_date,omitempty"`

	Stats *CreateCommitResponseBodyStats `json:"stats,omitempty"`
}

func (o CreateCommitResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCommitResponseBody struct{}"
	}

	return strings.Join([]string{"CreateCommitResponseBody", string(data)}, " ")
}
