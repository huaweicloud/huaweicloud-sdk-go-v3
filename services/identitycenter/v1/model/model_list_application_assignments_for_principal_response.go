package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicationAssignmentsForPrincipalResponse Response Object
type ListApplicationAssignmentsForPrincipalResponse struct {
	ApplicationAssignments *[]ApplicationAssignmentDto `json:"application_assignments,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListApplicationAssignmentsForPrincipalResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationAssignmentsForPrincipalResponse struct{}"
	}

	return strings.Join([]string{"ListApplicationAssignmentsForPrincipalResponse", string(data)}, " ")
}
