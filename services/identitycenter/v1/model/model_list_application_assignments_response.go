package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicationAssignmentsResponse Response Object
type ListApplicationAssignmentsResponse struct {

	// 应用程序分配的用户或用户组列表
	ApplicationAssignments *[]ApplicationAssignmentDto `json:"application_assignments,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListApplicationAssignmentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationAssignmentsResponse struct{}"
	}

	return strings.Join([]string{"ListApplicationAssignmentsResponse", string(data)}, " ")
}
