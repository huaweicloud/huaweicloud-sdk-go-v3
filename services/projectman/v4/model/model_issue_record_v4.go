package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 历史记录
type IssueRecordV4 struct {

	// 操作记录id
	Id *int32 `json:"id,omitempty" xml:"id"`

	// 操作记录创建时间
	CreatedTime *int64 `json:"created_time,omitempty" xml:"created_time"`

	User *IssueRecordV4User `json:"user,omitempty" xml:"user"`

	// 操作的记录
	Details *[]IssueRecordV4Details `json:"details,omitempty" xml:"details"`
}

func (o IssueRecordV4) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueRecordV4 struct{}"
	}

	return strings.Join([]string{"IssueRecordV4", string(data)}, " ")
}
