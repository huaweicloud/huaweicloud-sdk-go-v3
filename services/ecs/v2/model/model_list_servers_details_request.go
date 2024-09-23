package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServersDetailsRequest Request Object
type ListServersDetailsRequest struct {

	// 查询绑定某个企业项目的弹性云服务器。  若需要查询当前用户所有企业项目绑定的弹性云服务，请传参all_granted_eps。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 云服务器规格ID,已上线的规格请参见《弹性云服务器用户指南》的“实例和应用场景”章节。
	Flavor *string `json:"flavor,omitempty"`

	// IPv4地址过滤结果，匹配规则为模糊匹配。
	Ip *string `json:"ip,omitempty"`

	// 查询返回云服务器当前页面的大小。每页默认值是25，最多返回1000台云服务器的信息。
	Limit *int32 `json:"limit,omitempty"`

	// 云服务器名称，匹配规则为模糊匹配。
	Name *string `json:"name,omitempty"`

	// 查询tag字段中不包含该值的云服务器。
	NotTags *string `json:"not-tags,omitempty"`

	// 页码。 当前页面数，默认为1，取值范围大于等于0。 当取值为0时，系统默认返回第1页，与取值为1时相同。 建议设置该参数大于等于1。
	Offset *int32 `json:"offset,omitempty"`

	// 批量创建弹性云服务器时，指定返回的ID，用于查询本次批量创建的弹性云服务器。
	ReservationId *string `json:"reservation_id,omitempty"`

	// 云服务器状态。  取值范围：  ACTIVE， BUILD，DELETED，ERROR，HARD_REBOOT，MIGRATING，REBOOT，RESIZE，REVERT_RESIZE，SHELVED，SHELVED_OFFLOADED，SHUTOFF，UNKNOWN，VERIFY_RESIZE  只有管理员可以使用“deleted”状态过滤查询已经删除的弹性云服务器。  弹性云服务器状态说明请参考[云服务器状态](https://support.huaweicloud.com/api-ecs/ecs_08_0002.html)
	Status *string `json:"status,omitempty"`

	// 查询tag字段中包含该值的云服务器。
	Tags *string `json:"tags,omitempty"`

	// IPv4地址过滤结果，匹配规则为精确匹配。
	IpEq *string `json:"ip_eq,omitempty"`

	// 云服务器ID，格式为UUID，匹配规则为精确匹配  示例: server_id={id1},{id2}  说明： 在使用server_id作为过滤条件时，不能同时使用其他过滤条件。如果同时指定server_id及其他过滤条件，则以server_id条件为准，其他过滤条件会被忽略 当server_id中含有不存在的云服务器ID时，返回的响应参数中该云服务器ID对应的servers结构体中除了id和fault其它字段均为null 为了避免API的URI过长，建议一次查询的server_id个数不超过100个
	ServerId *string `json:"server_id,omitempty"`

	// 以单页最后一条server的id作为分页标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListServersDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServersDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListServersDetailsRequest", string(data)}, " ")
}
