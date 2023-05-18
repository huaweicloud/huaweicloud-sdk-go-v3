package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 父工作项信息
type WorkTableIssuseListResponseBodyParentIssue struct {

	// 父工作项id
	Id *int32 `json:"id,omitempty"`

	// 父工作项标题
	Subject *string `json:"subject,omitempty"`
}

func (o WorkTableIssuseListResponseBodyParentIssue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkTableIssuseListResponseBodyParentIssue struct{}"
	}

	return strings.Join([]string{"WorkTableIssuseListResponseBodyParentIssue", string(data)}, " ")
}
