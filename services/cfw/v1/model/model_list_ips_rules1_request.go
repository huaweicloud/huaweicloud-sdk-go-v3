package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIpsRules1Request Request Object
type ListIpsRules1Request struct {

	// 攻击对象
	AffectedApplicationLike *int32 `json:"affected_application_like,omitempty"`

	// 创建时间
	CreateTime *int32 `json:"create_time,omitempty"`

	// 防火墙ID，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	// cve
	IpsCveLike *int32 `json:"ips_cve_like,omitempty"`

	// ips组
	IpsGroup *int32 `json:"ips_group,omitempty"`

	// ips规则id
	IpsId *string `json:"ips_id,omitempty"`

	// ips等级
	IpsLevel *int32 `json:"ips_level,omitempty"`

	// ips规则名称
	IpsNameLike *string `json:"ips_name_like,omitempty"`

	// ips规则类型
	IpsRulesTypeLike *int32 `json:"ips_rules_type_like,omitempty"`

	// ips规则状态
	IpsStatus *string `json:"ips_status,omitempty"`

	// 是否查新更新规则
	IsUpdatedIpsRuleQueried *bool `json:"is_updated_ips_rule_queried,omitempty"`

	// 分页查询数据量限制
	Limit int32 `json:"limit"`

	// 防护对象ID，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id。此处仅取type为1的防护对象id，可通过data.records.protect_objects.type（.表示各对象之间层级的区分）获得。
	ObjectId string `json:"object_id"`

	// 查询偏移量
	Offset int32 `json:"offset"`

	// 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取，用户未开启企业项目时为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListIpsRules1Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpsRules1Request struct{}"
	}

	return strings.Join([]string{"ListIpsRules1Request", string(data)}, " ")
}
