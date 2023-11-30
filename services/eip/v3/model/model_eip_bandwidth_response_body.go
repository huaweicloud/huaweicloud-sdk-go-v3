package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// EipBandwidthResponseBody 带宽对象
type EipBandwidthResponseBody struct {

	// - 功能说明：带宽状态 - 取值范围：normal，freezed
	AdminState *string `json:"admin_state,omitempty"`

	// - 功能说明：入网大小，单位Mbit/s
	IngressSize *int32 `json:"ingress_size,omitempty"`

	// - 功能说明：规则数值，最低阈值可调节
	RuleQuota *int32 `json:"rule_quota,omitempty"`

	// - 功能说明：增强型95带宽保底率，最低保底率为20
	Ratio95peakPlus *int32 `json:"ratio_95peak_plus,omitempty"`

	// - 功能说明：带宽分组使能，表明开启带宽分组限速功能。
	EnableBandwidthRules *bool `json:"enable_bandwidth_rules,omitempty"`

	// - 功能说明：带宽规则对象（该字段仅在上海1局点返回）
	BandwidthRules *[]BandWidthRules `json:"bandwidth_rules,omitempty"`

	// - 功能说明：带宽AZ属性，表征中心和边缘。中心带宽默认为center
	PublicBorderGroup *string `json:"public_border_group,omitempty"`

	// - 功能说明：带宽类型，共享带宽默认为share。 - 取值范围：share，bgp，telcom，sbgp等。  share：共享带宽；  bgp：动态bgp；  telcom ：联通；  sbgp：静态bgp。
	BandwidthType *string `json:"bandwidth_type,omitempty"`

	// - 功能说明：账单信息，
	Billinginfo *string `json:"billinginfo,omitempty"`

	// - 功能说明：带宽唯一标识
	Id *string `json:"id,omitempty"`

	// - 功能说明：带宽名称 - 取值范围：1-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）
	Name *string `json:"name,omitempty"`

	// - 功能说明：带宽对应的弹性公网IP信息 - 约束：WHOLE类型的带宽支持多个弹性公网IP，PER类型的带宽只能对应一个弹性公网IP
	PublicipInfo *[]PublicipInfoResponseBody `json:"publicip_info,omitempty"`

	// - 功能说明：带宽类型，标识是否是共享带宽 - 取值范围：WHOLE，PER。  WHOLE表示共享带宽；PER表示独享带宽
	Type *EipBandwidthResponseBodyType `json:"type,omitempty"`

	// - 功能说明：带宽大小 - 取值范围：默认5Mbit/s~2000Mbit/s（具体范围以各区域配置为准，请参见控制台对应页面显示）。
	Size *int32 `json:"size,omitempty"`

	// - 功能说明：用户所属项目ID
	TenantId *string `json:"tenant_id,omitempty"`

	// - 功能说明：\"公网EIP标签\"
	Tags *[]string `json:"tags,omitempty"`

	// - 功能说明：资源创建时间，采用UTC时间，格式：YYYY-MM-DDTHH:MM:SS
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// - 功能说明：资源更新时间，采用UTC时间，格式：YYYY-MM-DDTHH:MM:SS
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o EipBandwidthResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipBandwidthResponseBody struct{}"
	}

	return strings.Join([]string{"EipBandwidthResponseBody", string(data)}, " ")
}

type EipBandwidthResponseBodyType struct {
	value string
}

type EipBandwidthResponseBodyTypeEnum struct {
	WHOLE EipBandwidthResponseBodyType
	PER   EipBandwidthResponseBodyType
}

func GetEipBandwidthResponseBodyTypeEnum() EipBandwidthResponseBodyTypeEnum {
	return EipBandwidthResponseBodyTypeEnum{
		WHOLE: EipBandwidthResponseBodyType{
			value: "WHOLE",
		},
		PER: EipBandwidthResponseBodyType{
			value: "PER",
		},
	}
}

func (c EipBandwidthResponseBodyType) Value() string {
	return c.value
}

func (c EipBandwidthResponseBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EipBandwidthResponseBodyType) UnmarshalJSON(b []byte) error {
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
