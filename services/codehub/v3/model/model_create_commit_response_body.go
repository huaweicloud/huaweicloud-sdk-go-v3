package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateCommitResponseBody struct {

	// 提交对应的SHA id
	Id *string `json:"id,omitempty" xml:"id"`

	// 提交对应的短SHA id
	ShortId *string `json:"short_id,omitempty" xml:"short_id"`

	// 提交标题
	Title *string `json:"title,omitempty" xml:"title"`

	// 作者
	AuthorName *string `json:"author_name,omitempty" xml:"author_name"`

	// 作者邮箱
	AuthorEmail *string `json:"author_email,omitempty" xml:"author_email"`

	// 提交作者
	CommitterName *string `json:"committer_name,omitempty" xml:"committer_name"`

	// 提交作者邮箱
	CommitterEmail *string `json:"committer_email,omitempty" xml:"committer_email"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 提交信息
	Message *string `json:"message,omitempty" xml:"message"`

	// 父提交id
	ParentIds *[]string `json:"parent_ids,omitempty" xml:"parent_ids"`

	// 提交时间
	CommittedDate *sdktime.SdkTime `json:"committed_date,omitempty" xml:"committed_date"`

	// 作者提交时间
	AuthoredDate *sdktime.SdkTime `json:"authored_date,omitempty" xml:"authored_date"`

	Stats *CreateCommitResponseBodyStats `json:"stats,omitempty" xml:"stats"`
}

func (o CreateCommitResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCommitResponseBody struct{}"
	}

	return strings.Join([]string{"CreateCommitResponseBody", string(data)}, " ")
}
