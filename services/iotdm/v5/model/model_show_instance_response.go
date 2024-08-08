package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceResponse Response Object
type ShowInstanceResponse struct {

	// **参数说明**：实例类型。 **取值范围**： - standard：标准版实例 - enterprise：企业版实例
	InstanceType *string `json:"instance_type,omitempty"`

	// **参数说明**：实例ID。 **取值范围**：长度不超过36，由小写字母[a-f]、数字、连接符（-）的组成。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数说明**：实例的付费方式。 **取值范围**： - prePaid：包年/包月 - postPaid：按需计费
	ChargeMode *string `json:"charge_mode,omitempty"`

	// **参数说明**：实例名称 **取值范围**：由中文字符，英文字母、数字及“_”、“-”组成，且长度为[1-64]个字符。
	Name *string `json:"name,omitempty"`

	Flavor *Flavor `json:"flavor,omitempty"`

	// **参数说明**：实例状态。 **取值范围**： - CREATING：实例正在创建 - ACTIVE：实例正常 - FROZEN：实例冻结 - MODIFYING：实例正在变更规格 - FAILED：实例创建失败
	Status *string `json:"status,omitempty"`

	// **参数说明**：设备接入实例的描述信息。 **取值范围**：由中文，字母，数字，句号，逗号，下划线（“_”），中划线（“-”），空格组成，且长度为[1-256]个字符。
	Description *string `json:"description,omitempty"`

	// **参数说明**：设备接入实例的接入信息
	AccessInfos *[]AccessInfo `json:"access_infos,omitempty"`

	// **参数说明**：实例的创建时间。时间格式例如：2023-01-28T06:57:52Z
	CreateTime *string `json:"create_time,omitempty"`

	// **参数说明**：实例的最近一次更新的时间。时间格式例如：2023-01-28T06:57:52Z
	UpdateTime *string `json:"update_time,omitempty"`

	// **参数说明**：企业项目Id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数说明**: 设备接入实例的标签信息。如果实例有标签，则会有该字段，否则该字段为空。
	Tags *[]Tag `json:"tags,omitempty"`

	// **参数说明**：订单号，仅包年包月实例返回该参数。[查看订单详情请参考[[查询订单详情](https://support.huaweicloud.com/api-bpconsole/zh-cn_topic_0075746564.html)。]](tag:hws)
	OrderId *string `json:"order_id,omitempty"`

	OperateWindow *OperateWindow `json:"operate_window,omitempty"`

	AdditionalParams *AdditionalParamsResp `json:"additional_params,omitempty"`
	HttpStatusCode   int                   `json:"-"`
}

func (o ShowInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceResponse", string(data)}, " ")
}
