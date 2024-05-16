package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityMemberSyncTasksResponse Response Object
type ListSecurityMemberSyncTasksResponse struct {

	// 用户同步任务总条数。
	Total *int64 `json:"total,omitempty"`

	// 用户同步任务列表。
	Tasks          *[]MemberSyncTask `json:"tasks,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListSecurityMemberSyncTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityMemberSyncTasksResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityMemberSyncTasksResponse", string(data)}, " ")
}
