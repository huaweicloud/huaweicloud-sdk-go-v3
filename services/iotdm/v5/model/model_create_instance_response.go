package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceResponse Response Object
type CreateInstanceResponse struct {

	// **参数说明**：实例类型。 **取值范围**： - standard：标准版实例 - enterprise：企业版实例
	InstanceType *string `json:"instance_type,omitempty"`

	// **参数说明**：实例ID。 **取值范围**：长度不超过36，由小写字母[a-f]、数字、连接符（-）的组成。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数说明**：实例名称 **取值范围**：由中文字符，英文字母、数字及“_”、“-”组成，且长度为[1-64]个字符。
	Name *string `json:"name,omitempty"`

	Flavor *Flavor `json:"flavor,omitempty"`

	// **参数说明**：实例状态。 **取值范围**： - CREATING：实例正在创建 - ACTIVE：实例正常 - FROZEN：实例冻结 - TRADING: 实例正在进行交易 - MODIFYING：实例正在变更规格 - MODIFY_FAILED: 实例变更失败 - FAILED：实例创建失败
	Status *string `json:"status,omitempty"`

	ChargeInfo *ChargeInfo `json:"charge_info,omitempty"`

	// **参数说明**：设备接入实例的描述信息。 **取值范围**：长度不超过256，只允许中文、字母、数字、以及_，,.。、&-等字符的组合。
	Description *string `json:"description,omitempty"`

	// **参数说明**：企业项目Id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数说明**: 设备接入实例的标签信息。如果实例有标签，则会有该字段，否则该字段为空。
	Tags *[]Tag `json:"tags,omitempty"`

	// **参数说明**：订单号，创建包年包月实例时返回该参数。[查看订单详情请参考[查询订单详情](https://support.huaweicloud.com/api-bpconsole/zh-cn_topic_0075746564.html)。](tag:hws)
	OrderId *string `json:"order_id,omitempty"`

	AdditionalParams *AdditionalParams `json:"additional_params,omitempty"`
	HttpStatusCode   int               `json:"-"`
}

func (o CreateInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceResponse", string(data)}, " ")
}
