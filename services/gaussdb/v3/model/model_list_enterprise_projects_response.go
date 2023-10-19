package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnterpriseProjectsResponse Response Object
type ListEnterpriseProjectsResponse struct {

	// 企业项目信息列表。
	EnterpriseProjects *[]EnterpriseProjectItem `json:"enterprise_projects,omitempty"`

	// 总数。
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
