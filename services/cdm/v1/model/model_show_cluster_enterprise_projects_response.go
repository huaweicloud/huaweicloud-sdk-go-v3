package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterEnterpriseProjectsResponse Response Object
type ShowClusterEnterpriseProjectsResponse struct {

	// 企业项目列表
	SysTags        *[]SysTags `json:"sys_tags,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowClusterEnterpriseProjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterEnterpriseProjectsResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterEnterpriseProjectsResponse", string(data)}, " ")
}
