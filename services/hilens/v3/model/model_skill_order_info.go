package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SkillOrderInfo 订单详情
type SkillOrderInfo struct {

	// 技能是否支持永久使用标识。1标识支持，0为不支持
	ExpirationStopFlag *int32 `json:"expiration_stop_flag,omitempty"`

	// 技能套餐包订单ID
	PackageOrderId *string `json:"package_order_id,omitempty"`

	// 技能图标
	Icon *string `json:"icon,omitempty"`

	// 定制技能标识
	CommissionFlag *int32 `json:"commission_flag,omitempty"`

	// 产品收费编码信息
	ProductInfo *[]string `json:"product_info,omitempty"`

	// 套餐包ID
	PackageId *string `json:"package_id,omitempty"`

	// 计费类型，physical_src表示按物理量纲收费，比如包周期 ，src表示一次性收费
	MeasureType *string `json:"measure_type,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 通道数限制
	ChannelLimit *int32 `json:"channel_limit,omitempty"`

	// 步长
	ResourceStepSize *int32 `json:"resource_step_size,omitempty"`

	// 云服务编码
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 开发者ID
	DeveloperId *string `json:"developer_id,omitempty"`

	// 订单数量
	Amount *int32 `json:"amount,omitempty"`

	// 技能类型，文件类型file，镜像类型iamge
	Format *string `json:"format,omitempty"`

	// 资源类别
	ResourceType *string `json:"resource_type,omitempty"`

	// 到期时间
	ExpireTime *int64 `json:"expire_time,omitempty"`

	// 计费单位 qps 表示按qps收费，road表示技能路数instance 表示按实例收费
	MeasureUnit *string `json:"measure_unit,omitempty"`

	// 芯片类别
	SkillChip *string `json:"skill_chip,omitempty"`

	// 技能版本列表
	Versions *[]string `json:"versions,omitempty"`

	// 技能名字
	SkillName *string `json:"skill_name,omitempty"`

	// 技能类别
	SkillType *string `json:"skill_type,omitempty"`

	// 订单使用份数
	UsedAmount *int32 `json:"used_amount,omitempty"`

	// 计费模式
	ChargeModel *int32 `json:"charge_model,omitempty"`

	// 资源编码
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 技能ID
	SkillId *string `json:"skill_id,omitempty"`

	// 技能支持的平台
	SkillPlatform *string `json:"skill_platform,omitempty"`

	// 订单购买限制
	OrderLimit *int32 `json:"order_limit,omitempty"`

	// 订单ID
	OrderId *string `json:"order_id,omitempty"`

	// 订单状态，0表示正常状态，1表示冻结状态，2表示受限状态
	Status *int32 `json:"status,omitempty"`
}

func (o SkillOrderInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SkillOrderInfo struct{}"
	}

	return strings.Join([]string{"SkillOrderInfo", string(data)}, " ")
}
