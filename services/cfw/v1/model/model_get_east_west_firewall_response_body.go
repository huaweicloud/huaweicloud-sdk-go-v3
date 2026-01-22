package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetEastWestFirewallResponseBody 查询东西向防火墙接口响应Body体
type GetEastWestFirewallResponseBody struct {

	// **参数解释**： 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得。type为0时，object_id为互联网边界防护对象ID，type为1时，object_id为VPC边界防护对象ID。此处仅取type为1的防护对象id，可通过data.records.protect_objects.type（.表示各对象之间层级的区分）获得。 **取值范围**： 不涉及
	ObjectId *string `json:"object_id,omitempty"`

	// **参数解释**： 项目ID，可以通过调用API获取，也可以从控制台获取。[项目ID获取方式](cfw_02_0015.xml) **取值范围**： 不涉及
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 防护状态： 不涉及 **取值范围**： 0：已开启防护 1：未开启防护
	Status *int32 `json:"status,omitempty"`

	// **参数解释**： 云防火墙关联子网信息 **取值范围**： 不涉及
	FirewallAssociatedSubnets *[]SubnetInfo `json:"firewall_associated_subnets,omitempty"`

	Er *ErInstance `json:"er,omitempty"`

	InspectionVpc *VpcDetail `json:"inspection_vpc,omitempty"`

	// **参数解释**： 东西向防护资源信息 **取值范围**： 不涉及
	ProtectInfos *[]EwProtectResourceInfo `json:"protect_infos,omitempty"`

	// **参数解释**： 防护VPC总数 **取值范围**： 不涉及
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 偏移量：指定返回记录的开始位置 **取值范围**： 大于或等于0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 每页显示个数 **取值范围**： 1-1024
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 防护模式，值为er **取值范围**： 不涉及
	Mode *string `json:"mode,omitempty"`
}

func (o GetEastWestFirewallResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetEastWestFirewallResponseBody struct{}"
	}

	return strings.Join([]string{"GetEastWestFirewallResponseBody", string(data)}, " ")
}
