package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIegInfoResponse Response Object
type ShowIegInfoResponse struct {

	// 智能企业网关ID
	Id *string `json:"id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 租户账号ID
	DomainId *string `json:"domain_id,omitempty"`

	// 智能企业网关名字
	Name *string `json:"name,omitempty"`

	// 大区ID
	AreaId *string `json:"area_id,omitempty"`

	// 设备类型
	EquipmentType *string `json:"equipment_type,omitempty"`

	// 高可用性
	HighAvailability *string `json:"high_availability,omitempty"`

	// 冻结效果
	FrozenEffect *int32 `json:"frozen_effect,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 绑定的企业连接网络ID
	BindEcn *string `json:"bind_ecn,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 包周期场景下购买的订单号，按需场景下为空
	OrderId *string `json:"order_id,omitempty"`

	// 包周期场景下购买的订单号，按需场景下为空
	ProductId *string `json:"product_id,omitempty"`

	// ieg设备信息
	EquipmentInfos *[]EquipmentItem `json:"equipment_infos,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowIegInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIegInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowIegInfoResponse", string(data)}, " ")
}
