package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryInstanceSimplify struct {

	// **参数说明**：实例类型。 **取值范围**： - standard：标准版实例 - enterprise：企业版实例
	InstanceType *string `json:"instance_type,omitempty"`

	// **参数说明**：实例ID。 **取值范围**：长度不超过36，由小写字母[a-f]、数字、连接符（-）的组成。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数说明**：实例名称 **取值范围**：由中文字符，英文字母、数字及“_”、“-”组成，且长度为[1-64]个字符。
	Name *string `json:"name,omitempty"`

	// **参数说明**：实例的付费方式。 **取值范围**： - prePaid：包年/包月 - postPaid：按需计费
	ChargeMode *string `json:"charge_mode,omitempty"`

	Flavor *Flavor `json:"flavor,omitempty"`

	// **参数说明**：实例状态。 **取值范围**： - CREATING：实例正在创建 - ACTIVE：实例正常 - FROZEN：实例冻结 - MODIFYING：实例正在变更规格 - FAILED：实例创建失败
	Status *string `json:"status,omitempty"`

	// **参数说明**：实例的创建时间。时间格式例如：2023-01-28T06:57:52Z
	CreateTime *string `json:"create_time,omitempty"`

	// **参数说明**：实例的最近一次更新的时间。时间格式例如：2023-01-28T06:57:52Z
	UpdateTime *string `json:"update_time,omitempty"`

	// **参数说明**：企业项目Id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o QueryInstanceSimplify) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryInstanceSimplify struct{}"
	}

	return strings.Join([]string{"QueryInstanceSimplify", string(data)}, " ")
}
