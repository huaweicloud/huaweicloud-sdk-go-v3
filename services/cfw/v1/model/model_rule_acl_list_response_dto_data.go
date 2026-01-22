package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuleAclListResponseDtoData **参数解释**： 查询规则列表返回值数据
type RuleAclListResponseDtoData struct {

	// **参数解释**： 偏移量：指定返回记录的开始位置 **取值范围**： 大于或等于0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 每页显示个数 **取值范围**： 1-1024
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 查询规则列表总条数 **取值范围**： 大于0
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得，type为0时，object_id为互联网边界防护对象ID，type为1时，object_id为VPC边界防护对象ID，type可通过data.records.protect_objects.type（.表示各对象之间层级的区分）获得  **取值范围**：  32位UUID
	ObjectId *string `json:"object_id,omitempty"`

	// **参数解释**： 顶部规则数量 **取值范围**： 不涉及
	UpRulesCount *int32 `json:"up_rules_count,omitempty"`

	// **参数解释**： 查询规则列表记录
	Records *[]RuleAclListResponseDtoDataRecords `json:"records,omitempty"`
}

func (o RuleAclListResponseDtoData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleAclListResponseDtoData struct{}"
	}

	return strings.Join([]string{"RuleAclListResponseDtoData", string(data)}, " ")
}
