package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIpsRulesRequest Request Object
type ListIpsRulesRequest struct {

	// 受影响对象查询关键字，可包含如下：Others、Sun、Apache、IBM、VMware、WordPress、Adobe、Oracle、Google Chrome等
	AffectedApplicationLike *int32 `json:"affected_application_like,omitempty"`

	// IPS规则创建的年份
	CreateTime *int32 `json:"create_time,omitempty"`

	// 防火墙ID，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	// cve id查询关键字，cve id为cve漏洞库中存储的漏洞id
	IpsCveLike *int32 `json:"ips_cve_like,omitempty"`

	// IPS组，使用IPS规则拦截模式区分，包含，0：观察模式，1：严格模式，2：中等模式，3：宽松模式
	IpsGroup *int32 `json:"ips_group,omitempty"`

	// IPS规则id
	IpsId *string `json:"ips_id,omitempty"`

	// IPS严重等级，包含CRITICAL、HIGH、MEDIUM、LOW
	IpsLevel *int32 `json:"ips_level,omitempty"`

	// IPS规则名称查询关键字
	IpsNameLike *string `json:"ips_name_like,omitempty"`

	// IPS规则类型，包括漏洞扫描、黑客工具、特洛伊木马等
	IpsRulesTypeLike *int32 `json:"ips_rules_type_like,omitempty"`

	// IPS规则状态，包含观察：OBSERVE、拦截：ENABLE、禁用：CLOSE、恢复默认：DEFAULT、全局恢复默认：ALL_DEFAULT
	IpsStatus *string `json:"ips_status,omitempty"`

	// 是否查询虚拟补丁相对基础防御更新规则，是表示true,否表示false
	IsUpdatedIpsRuleQueried *bool `json:"is_updated_ips_rule_queried,omitempty"`

	// 每页显示个数，范围为1-1024
	Limit int32 `json:"limit"`

	// 防护对象ID，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得，type为0时，object_id为互联网边界防护对象ID，type为1时，object_id为VPC边界防护对象ID。此处仅取type为0的防护对象id，可通过data.records.protect_objects.type（.表示各对象之间层级的区分）获得。
	ObjectId string `json:"object_id"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset int32 `json:"offset"`

	// 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取，用户未开启企业项目时为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListIpsRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpsRulesRequest struct{}"
	}

	return strings.Join([]string{"ListIpsRulesRequest", string(data)}, " ")
}
