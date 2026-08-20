package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIssueSprintSnapshotsResponse Response Object
type ListIssueSprintSnapshotsResponse struct {

	// 快照列表。
	Result *[]SnapshotsVo `json:"result,omitempty"`

	// 请求状态。
	Status *string `json:"status,omitempty"`

	// 请求失败信息。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListIssueSprintSnapshotsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssueSprintSnapshotsResponse struct{}"
	}

	return strings.Join([]string{"ListIssueSprintSnapshotsResponse", string(data)}, " ")
}
