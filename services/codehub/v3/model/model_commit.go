package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/sdktime"
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Commit struct {
	// 作者邮箱

	AuthorEmail *string `json:"author_email,omitempty"`
	// 作者

	AuthorName *string `json:"author_name,omitempty"`
	// 作者提交时间

	AuthoredDate *sdktime.SdkTime `json:"authored_date,omitempty"`
	// 提交时间

	CommittedDate *sdktime.SdkTime `json:"committed_date,omitempty"`
	// 提交作者邮箱

	CommitterEmail *string `json:"committer_email,omitempty"`
	// 提交作者

	CommitterName *string `json:"committer_name,omitempty"`
	// 文件变更的详情信息，其格式由请求查询参数 stat_format 决定

	Format *interface{} `json:"format,omitempty"`
	// 提交对应的SHA id

	Id *string `json:"id,omitempty"`
	// 提交的信息

	Message *string `json:"message,omitempty"`
	// 父提交id

	ParentIds *[]string `json:"parent_ids,omitempty"`
}

func (o Commit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Commit struct{}"
	}

	return strings.Join([]string{"Commit", string(data)}, " ")
}
