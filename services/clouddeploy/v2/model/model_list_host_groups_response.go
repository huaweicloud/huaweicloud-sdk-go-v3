package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListHostGroupsResponse struct {

	// 主机组个数
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 主机组详情响应体
	HostGroups     *[]DeploymentGroupDetail `json:"host_groups,omitempty" xml:"host_groups"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListHostGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListHostGroupsResponse", string(data)}, " ")
}
