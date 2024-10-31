package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetEastWestFirewallResponseBody 查询东西向防火墙接口响应Body体
type GetEastWestFirewallResponseBody struct {

	// 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id。此处仅取type为1的防护对象id，可通过data.records.protect_objects.type（.表示各对象之间层级的区分）获得。
	ObjectId *string `json:"object_id,omitempty"`

	// 项目ID, 可以从调API处获取，也可以从控制台获取。[项目ID获取方式](cfw_02_0015.xml)
	ProjectId *string `json:"project_id,omitempty"`

	// 防护状态：0 已开启防护， 1 未开启防护
	Status *int32 `json:"status,omitempty"`

	// 云防火墙关联子网信息
	FirewallAssociatedSubnets *[]SubnetInfo `json:"firewall_associated_subnets,omitempty"`

	Er *ErInstance `json:"er,omitempty"`

	InspectionVpc *VpcDetail `json:"inspection_vpc,omitempty"`

	// 东西向防护资源信息
	ProtectInfos *[]EwProtectResourceInfo `json:"protect_infos,omitempty"`

	// 防护VPC总数
	Total *int32 `json:"total,omitempty"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示个数，范围为1-1024
	Limit *int32 `json:"limit,omitempty"`

	// 防护模式，值为er
	Mode *string `json:"mode,omitempty"`
}

func (o GetEastWestFirewallResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetEastWestFirewallResponseBody struct{}"
	}

	return strings.Join([]string{"GetEastWestFirewallResponseBody", string(data)}, " ")
}
