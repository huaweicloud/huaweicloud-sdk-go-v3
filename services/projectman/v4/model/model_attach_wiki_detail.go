package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 关联Wiki详情
type AttachWikiDetail struct {

	// 工作项ID
	IssueId *int32 `json:"issue_id,omitempty" xml:"issue_id"`

	// Wiki标题
	WikiTitle *string `json:"wiki_title,omitempty" xml:"wiki_title"`

	WikiAuthor *SimpleUser `json:"wiki_author,omitempty" xml:"wiki_author"`

	Project *SimpleProject `json:"project,omitempty" xml:"project"`

	// 创建时间
	CreatedDate *string `json:"created_date,omitempty" xml:"created_date"`

	// wiki ID
	WikiId *string `json:"wiki_id,omitempty" xml:"wiki_id"`

	// region值
	Region *string `json:"region,omitempty" xml:"region"`
}

func (o AttachWikiDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachWikiDetail struct{}"
	}

	return strings.Join([]string{"AttachWikiDetail", string(data)}, " ")
}
