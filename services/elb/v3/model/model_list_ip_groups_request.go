package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIpGroupsRequest Request Object
type ListIpGroupsRequest struct {

	// 参数解释：上一页最后一条记录的ID。  使用说明： - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。
	Marker *string `json:"marker,omitempty"`

	// 参数解释：每页返回的个数。  取值范围：0-2000  默认取值：2000
	Limit *int32 `json:"limit,omitempty"`

	// 参数解释：是否反向查询。  取值： - true：查询上一页。 - false：查询下一页，默认。  使用说明： - 必须与limit一起使用。 - 当page_reverse=true时，若要查询上一页，marker取值为当前页返回值的previous_marker。
	PageReverse *bool `json:"page_reverse,omitempty"`

	// 参数解释：IP地址组的ID。
	Id *[]string `json:"id,omitempty"`

	// 参数解释：IP地址组的名称。
	Name *[]string `json:"name,omitempty"`

	// 参数解释：IP地址组的描述信息。
	Description *[]string `json:"description,omitempty"`

	// 参数解释：IP地址，多个用逗号分隔。
	IpList *[]string `json:"ip_list,omitempty"`

	// 参数解释：企业项目ID。
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
}

func (o ListIpGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListIpGroupsRequest", string(data)}, " ")
}
