package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUsersForEnterpriseProjectRequest Request Object
type ListUsersForEnterpriseProjectRequest struct {

	// 待查询企业项目ID。
	EnterpriseProjectId string `json:"enterprise_project_id"`
}

func (o ListUsersForEnterpriseProjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsersForEnterpriseProjectRequest struct{}"
	}

	return strings.Join([]string{"ListUsersForEnterpriseProjectRequest", string(data)}, " ")
}
