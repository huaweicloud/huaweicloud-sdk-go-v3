package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomerIpsRequest Request Object
type ListCustomerIpsRequest struct {

	// 动作（0：只记录日志，1：重置/拦截）
	ActionType *int32 `json:"action_type,omitempty"`

	// 操作系统
	AffectedOs *int32 `json:"affected_os,omitempty"`

	// 攻击类型
	AttackType *int32 `json:"attack_type,omitempty"`

	// 防火墙ID，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取
	FwInstanceId string `json:"fw_instance_id"`

	// ips规则名称
	IpsName *string `json:"ips_name,omitempty"`

	// 分页查询数据量限制
	Limit int32 `json:"limit"`

	// 防护对象ID，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id。此处仅取type为1的防护对象id，可通过data.records.protect_objects.type（.表示各对象之间层级的区分）获得。
	ObjectId string `json:"object_id"`

	// 查询偏移量
	Offset int32 `json:"offset"`

	// 协议
	Protocol *int32 `json:"protocol,omitempty"`

	// 严重程度（critical：致命，high：高危，medium:中危，low:低危）
	Severity *int32 `json:"severity,omitempty"`

	// 影响软件
	Software *int32 `json:"software,omitempty"`

	// 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取，用户未开启企业项目时为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListCustomerIpsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomerIpsRequest struct{}"
	}

	return strings.Join([]string{"ListCustomerIpsRequest", string(data)}, " ")
}
