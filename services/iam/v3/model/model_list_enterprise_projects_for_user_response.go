package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnterpriseProjectsForUserResponse Response Object
type ListEnterpriseProjectsForUserResponse struct {

	// 企业项目信息。
	EnterpriseProjects *[]ListEnterpriseProjectsResDetail `json:"enterprise-projects,omitempty"`
	HttpStatusCode     int                                `json:"-"`
}

func (o ListEnterpriseProjectsForUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnterpriseProjectsForUserResponse struct{}"
	}

	return strings.Join([]string{"ListEnterpriseProjectsForUserResponse", string(data)}, " ")
}
