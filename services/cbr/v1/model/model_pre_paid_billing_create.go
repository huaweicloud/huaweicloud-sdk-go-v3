package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PrePaidBillingCreate 创建参数
type PrePaidBillingCreate struct {

	// 云类型，默认为public，支持类型如下。 [public：公有云; hybrid: 混合云](tag:hws,hws_hk,ctc) [public：公有云](tag:dt,ocb,tlf,sbc,g42,tm,hk_g42)
	CloudType *string `json:"cloud_type,omitempty"`

	// [功能描述：存储库规格。取值范围：app_consistent: 应用一致性，crash_consistent: 崩溃一致性。默认取值不涉及。](tag:hws,hws_hk,fcs_vm,ctc,tm,g42,hk_g42) [功能描述：存储库规格。取值范围：crash_consistent: 崩溃一致性。默认取值不涉及。](tag:dt,ocb,tlf,sbc,hcso_dt)
	ConsistentLevel string `json:"consistent_level"`

	// [对象类型，支持\"server\", \"disk\", \"turbo\", \"workspace\", \"vmware\", \"rds\"和\"file\"共七种。server：云服务器，disk：云硬盘，turbo：文件系统，workspace：云桌面，vmware：VMware，rds：关系型数据库，file：文件。默认取值不涉及。](tag:hws,hws_hk) [对象类型，支持\"server\", \"disk\"和\"turbo\"共三种。server：云服务器，disk：云硬盘，turbo：文件系统。默认取值不涉及。](tag:ctc,fcs_vm,ocb,hk_g42,sbc,hws_ocb) [对象类型，支持\"server\"和\"disk\"共两种。server：云服务器，disk：云硬盘。默认取值不涉及。](tag:dt,tlf,tm,cmcc,hcso_dt) [对象类型，支持\"server\", \"disk\", \"turbo\"和\"workspace\"共四种。server：云服务器，disk：云硬盘，turbo：文件系统，workspace：云桌面。默认取值不涉及。](tag:g42)
	ObjectType string `json:"object_type"`

	// 保护类型，默认取值不涉及。取值范围如下： [backup：备份，replication：复制](tag:hws,hws_hk,ocb,hws_ocb) [backup：备份](tag:tlf,tm,cmcc,fcs_vm,g42,dt,hk_g42,sbc,hcso_dt)
	ProtectType string `json:"protect_type"`

	// 资源容量大小，单位GB，取值范围：10-10485760，默认取值不涉及。
	Size int32 `json:"size"`

	// 计费模式，仅支持填写pre_paid：代表包年/包月模式
	ChargingMode string `json:"charging_mode"`

	// 功能说明：订购周期单位。charging_mode参数为pre_paid时period_type参数会生效，并且period_type参数为必选。默认取值不涉及。 取值范围： - month：月 - year：年
	PeriodType PrePaidBillingCreatePeriodType `json:"period_type"`

	// 功能说明：订购周期数，charging_mode为pre_paid时period_num参数会生效，并且period_num参数为为必选。默认取值不涉及。 取值范围：[1-9]
	PeriodNum int32 `json:"period_num"`

	// 功能说明：到期后是否自动续期，默认为false 取值范围： - true：到期后自动续期 - false：到期后不自动续期
	IsAutoRenew *bool `json:"is_auto_renew,omitempty"`

	// 功能说明：是否自动付费，默认为false 取值范围： - true：下单后自动付费 - false：下单后不自动付费
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`

	// 云服务console_url。 订购订单支付完成后，客户可以通过此URL跳转到云服务Console页面查看信息。（仅手动支付时涉及）。默认取值不涉及。
	ConsoleUrl *string `json:"console_url,omitempty"`

	// 功能说明：存储库是否具有多AZ属性，即底层备份是否为多AZ备份，默认为false 取值范围： - true：存储库具有多AZ属性 - false：存储库不具有多AZ属性
	IsMultiAz *bool `json:"is_multi_az,omitempty"`

	// 功能说明：存储库是否具有融合桶属性，即底层备份是否为融合桶备份，默认为false 取值范围： - true：存储库具有融合桶属性 - false：存储库不具有融合桶属性
	IsDoubleAz *bool `json:"is_double_az,omitempty"`

	// 促销信息，包周期时可选参数，取值范围不涉及，默认取值不涉及。
	PromotionInfo *string `json:"promotion_info,omitempty"`

	// 购买模式，包周期时可选参数，取值范围不涉及，默认取值不涉及。
	PurchaseMode *string `json:"purchase_mode,omitempty"`

	// 订单 ID，包周期时可选参数，取值范围不涉及，默认取值不涉及。
	OrderId *string `json:"order_id,omitempty"`
}

func (o PrePaidBillingCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrePaidBillingCreate struct{}"
	}

	return strings.Join([]string{"PrePaidBillingCreate", string(data)}, " ")
}

type PrePaidBillingCreatePeriodType struct {
	value string
}

type PrePaidBillingCreatePeriodTypeEnum struct {
	YEAR  PrePaidBillingCreatePeriodType
	MONTH PrePaidBillingCreatePeriodType
}

func GetPrePaidBillingCreatePeriodTypeEnum() PrePaidBillingCreatePeriodTypeEnum {
	return PrePaidBillingCreatePeriodTypeEnum{
		YEAR: PrePaidBillingCreatePeriodType{
			value: "year",
		},
		MONTH: PrePaidBillingCreatePeriodType{
			value: "month",
		},
	}
}

func (c PrePaidBillingCreatePeriodType) Value() string {
	return c.value
}

func (c PrePaidBillingCreatePeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrePaidBillingCreatePeriodType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
