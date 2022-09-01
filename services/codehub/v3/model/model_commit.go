package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Commit struct {

	// 作者邮箱
	AuthorEmail *string `json:"author_email,omitempty" xml:"author_email"`

	// 作者
	AuthorName *string `json:"author_name,omitempty" xml:"author_name"`

	// 作者提交时间
	AuthoredDate *sdktime.SdkTime `json:"authored_date,omitempty" xml:"authored_date"`

	// 提交时间
	CommittedDate *sdktime.SdkTime `json:"committed_date,omitempty" xml:"committed_date"`

	// 提交作者邮箱
	CommitterEmail *string `json:"committer_email,omitempty" xml:"committer_email"`

	// 提交作者
	CommitterName *string `json:"committer_name,omitempty" xml:"committer_name"`

	// 文件变更的详情信息，其格式由请求查询参数 stat_format 决定
	Format *interface{} `json:"format,omitempty" xml:"format"`

	// 提交对应的SHA id
	Id *string `json:"id,omitempty" xml:"id"`

	// 提交的信息
	Message *string `json:"message,omitempty" xml:"message"`

	// 父提交id
	ParentIds *[]string `json:"parent_ids,omitempty" xml:"parent_ids"`
}

func (o Commit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Commit struct{}"
	}

	return strings.Join([]string{"Commit", string(data)}, " ")
}
