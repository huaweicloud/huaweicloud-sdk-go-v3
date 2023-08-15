package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostGroupsResponse Response Object
type ListHostGroupsResponse struct {

	// 主机集群个数
	Total *int32 `json:"total,omitempty"`

	// 主机集群详情响应体
	HostGroups     *[]DeploymentGroupDetail `json:"host_groups,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListHostGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListHostGroupsResponse", string(data)}, " ")
}
