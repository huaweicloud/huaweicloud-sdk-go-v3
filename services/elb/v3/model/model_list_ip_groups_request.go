package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListIpGroupsRequest struct {
	// 上一页最后一条记录的ID。  使用说明： - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。

	Marker *string `json:"marker,omitempty"`
	// 每页返回的个数。

	Limit *int32 `json:"limit,omitempty"`
	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。  使用说明： - 必须与limit一起使用。

	PageReverse *bool `json:"page_reverse,omitempty"`
	// IP地址组的ID。

	Id *[]string `json:"id,omitempty"`
	// IP地址组的名称。

	Name *[]string `json:"name,omitempty"`
	// IP地址组的描述信息。

	Description *[]string `json:"description,omitempty"`
	// IP地址，多个用逗号分隔。

	IpList *[]string `json:"ip_list,omitempty"`
}

func (o ListIpGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListIpGroupsRequest", string(data)}, " ")
}
