package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListViewResourceResponseBodyData struct {

	// **参数解释：** id。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 视图id。 **取值范围：** 不涉及。
	ViewId *string `json:"view_id,omitempty"`

	// **参数解释：** 对应rms_resource集合中的id值。 **取值范围：** 不涉及。
	RmsResourceId *string `json:"rms_resource_id,omitempty"`

	// **参数解释：** 资源id。。 **取值范围：** 跨账号资源下且视图管理下对应的资源id。
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释：** 租户id。 **取值范围：** 用户登录租户对应的账号ID。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释：** 资源名称。 **取值范围：** 视图下的资源名称。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 企业项目ID。 **取值范围：** 请选择[[企业管理](https://support.huaweicloud.com/usermanual-em/em_eps_qs_0400.html)](tag:hws)[[企业管理](https://support.huaweicloud.com/intl/zh-cn/usermanual-em/em_eps_qs_0400.html)](tag:hws_hk)中存在的项目ID。。
	EpId *string `json:"ep_id,omitempty"`

	// **参数解释：** 企业项目名称。 **取值范围：** 不涉及。
	EpName *string `json:"ep_name,omitempty"`

	// **参数解释：** Openstack中的项目ID。 **取值范围：** 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释：** 云服务名称。 **取值范围：** 不涉及。
	Provider *string `json:"provider,omitempty"`

	// **参数解释：** 资源类型。 **取值范围：** 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释：** 区域id。 **取值范围：** 不涉及。
	RegionId *string `json:"region_id,omitempty"`

	// **参数解释：** 标签键值对。 **取值范围：** 不涉及。
	Tags *[]ListViewResourceResponseBodyTags `json:"tags,omitempty"`

	// **参数解释：** 存储资源的附加字段信息，通常用于展示、筛选等。 **取值范围：** 不涉及。
	Properties *interface{} `json:"properties,omitempty"`

	// **参数解释：** 数据采集属性,描述视图系统采集该资源时所记录的附加信息。 **取值范围：** 不涉及。
	IngestProperties *interface{} `json:"ingest_properties,omitempty"`

	// **参数解释：** uniagent的id值。 **取值范围：** 不涉及。
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释：** uniagent的状态。 **取值范围：** 不涉及。
	AgentState *string `json:"agent_state,omitempty"`

	// **参数解释：** 资源内网ip。 **取值范围：** 不涉及。
	InnerIp *string `json:"inner_ip,omitempty"`

	// **参数解释：** 绑定的资源组信息列表。 **取值范围：** 不涉及。
	AssociatedGroupList *interface{} `json:"associated_group_list,omitempty"`
}

func (o ListViewResourceResponseBodyData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListViewResourceResponseBodyData struct{}"
	}

	return strings.Join([]string{"ListViewResourceResponseBodyData", string(data)}, " ")
}
