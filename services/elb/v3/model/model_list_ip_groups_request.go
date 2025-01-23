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

	// 参数解释：所属的企业项目ID。 如果enterprise_project_id不传值，默认查询所有企业项目下的资源，鉴权按照细粒度权限鉴权，必须在用户组下分配elb:ipgroups:list权限。 如果enterprise_project_id传值，鉴权按照企业项目权限鉴权，分为传入具体eps_id和all_granted_eps两种场景，前者查询指定eps_id的eps下的资源，后者查询的是所有有list权限的eps下的资源。  支持多值查询，查询条件格式： *enterprise_project_id=xxx&enterprise_project_id=xxx*。  [不支持该字段，请勿使用。](tag:dt,dt_test,hcso_dt)
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
}

func (o ListIpGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListIpGroupsRequest", string(data)}, " ")
}
