package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListHostsResponse struct {
	// 主机数量

	Total *int32 `json:"total,omitempty"`
	// 主机组名称

	GroupName *string `json:"group_name,omitempty"`
	// 主机列表信息

	Hosts          *[]DeploymentHostDetail `json:"hosts,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListHostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostsResponse struct{}"
	}

	return strings.Join([]string{"ListHostsResponse", string(data)}, " ")
}
