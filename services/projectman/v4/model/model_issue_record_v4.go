package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueRecordV4 历史记录
type IssueRecordV4 struct {

	// 操作记录id
	Id *int32 `json:"id,omitempty"`

	// 操作记录创建时间
	CreatedTime *int64 `json:"created_time,omitempty"`

	User *IssueRecordV4User `json:"user,omitempty"`

	// 操作的记录
	Details *[]IssueRecordV4Details `json:"details,omitempty"`
}

func (o IssueRecordV4) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueRecordV4 struct{}"
	}

	return strings.Join([]string{"IssueRecordV4", string(data)}, " ")
}
