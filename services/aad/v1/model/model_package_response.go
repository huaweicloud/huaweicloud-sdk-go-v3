package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PackageResponse 包列表响应体
type PackageResponse struct {

	// 防护包id
	PackageId string `json:"package_id"`

	// 防护包名
	PackageName string `json:"package_name"`

	// 资源所属region
	RegionId string `json:"region_id"`

	// 防护类型
	ProtectionType int32 `json:"protection_type"`

	// 防护包类型。cnad_pro：专业版；cnad_ip：标准版；cnad_ep：铂金版；cnad_full_high：全力防高级版；cnad_vic：按需版；cnad_intl_ep：国际站铂金版
	InstanceType PackageResponseInstanceType `json:"instance_type"`

	// 资源id
	ResourceId string `json:"resource_id"`

	// 倒计时相关信息
	CountDownCode *string `json:"count_down_code,omitempty"`

	// 倒计时相关信息
	CountDownInfos *string `json:"count_down_infos,omitempty"`

	// 倒计时相关信息
	CountDownTips *string `json:"count_down_tips,omitempty"`

	// 订单id
	OrderId *string `json:"order_id,omitempty"`

	// 续费用的id
	SubscriptionId *string `json:"subscription_id,omitempty"`

	// ip数
	IpNum *int32 `json:"ip_num,omitempty"`

	// 当前IP数
	IpNumNow *int32 `json:"ip_num_now,omitempty"`

	// 当前防护次数
	ProtectionNumNow *int32 `json:"protection_num_now,omitempty"`

	// 防护次数，9999为无限次
	ProtectionNum *int32 `json:"protection_num,omitempty"`

	// 保底带宽
	BasicBandwidth *int32 `json:"basic_bandwidth,omitempty"`

	// 弹性带宽
	ElasticBandwidth *int32 `json:"elastic_bandwidth,omitempty"`

	// 业务带宽
	ServiceBandwidth *int32 `json:"service_bandwidth,omitempty"`

	// 回源带宽
	CleanBandwidth *int32 `json:"clean_bandwidth,omitempty"`

	// 策略模板数
	PolicyNum int32 `json:"policy_num"`

	// 是否旧防护包（旧防护包不支持升级规格）,默认不传为否
	IsOld *bool `json:"is_old,omitempty"`

	// 专业版铂金版合并之后购买的专业版和铂金版均标识为true
	NewFlag *bool `json:"new_flag,omitempty"`

	// 创建时间
	CreateTime int64 `json:"create_time"`
}

func (o PackageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PackageResponse struct{}"
	}

	return strings.Join([]string{"PackageResponse", string(data)}, " ")
}

type PackageResponseInstanceType struct {
	value string
}

type PackageResponseInstanceTypeEnum struct {
	CNAD_PRO       PackageResponseInstanceType
	CNAD_IP        PackageResponseInstanceType
	CNAD_EP        PackageResponseInstanceType
	CNAD_FULL_HIGH PackageResponseInstanceType
	CNAD_VIC       PackageResponseInstanceType
	CNAD_INTL_EP   PackageResponseInstanceType
}

func GetPackageResponseInstanceTypeEnum() PackageResponseInstanceTypeEnum {
	return PackageResponseInstanceTypeEnum{
		CNAD_PRO: PackageResponseInstanceType{
			value: "cnad_pro",
		},
		CNAD_IP: PackageResponseInstanceType{
			value: "cnad_ip",
		},
		CNAD_EP: PackageResponseInstanceType{
			value: "cnad_ep",
		},
		CNAD_FULL_HIGH: PackageResponseInstanceType{
			value: "cnad_full_high",
		},
		CNAD_VIC: PackageResponseInstanceType{
			value: "cnad_vic",
		},
		CNAD_INTL_EP: PackageResponseInstanceType{
			value: "cnad_intl_ep",
		},
	}
}

func (c PackageResponseInstanceType) Value() string {
	return c.value
}

func (c PackageResponseInstanceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PackageResponseInstanceType) UnmarshalJSON(b []byte) error {
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
