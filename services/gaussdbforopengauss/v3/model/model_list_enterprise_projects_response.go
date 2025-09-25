package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnterpriseProjectsResponse Response Object
type ListEnterpriseProjectsResponse struct {

	// **参数解释**: 企业项目列表。
	EnterpriseProjects *[]EnterpriseProjectInfoResult `json:"enterprise_projects,omitempty"`

	// **参数解释**: 企业项目总数。 **取值范围**: 不涉及。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEnterpriseProjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnterpriseProjectsResponse struct{}"
	}

	return strings.Join([]string{"ListEnterpriseProjectsResponse", string(data)}, " ")
}
