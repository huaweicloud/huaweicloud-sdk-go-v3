package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListResourceGroupResponse struct {

	// 一个或者多个资源分组信息。
	ResourceGroups *[]ResourceGroupInfo `json:"resource_groups,omitempty" xml:"resource_groups"`

	MetaData       *TotalMetaData `json:"meta_data,omitempty" xml:"meta_data"`
	HttpStatusCode int            `json:"-"`
}

func (o ListResourceGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceGroupResponse struct{}"
	}

	return strings.Join([]string{"ListResourceGroupResponse", string(data)}, " ")
}
