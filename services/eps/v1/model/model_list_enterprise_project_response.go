package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnterpriseProjectResponse Response Object
type ListEnterpriseProjectResponse struct {

	// 企业项目列表
	EnterpriseProjects *[]EpDetail `json:"enterprise_projects,omitempty"`

	// 企业项目总数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEnterpriseProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnterpriseProjectResponse struct{}"
	}

	return strings.Join([]string{"ListEnterpriseProjectResponse", string(data)}, " ")
}
