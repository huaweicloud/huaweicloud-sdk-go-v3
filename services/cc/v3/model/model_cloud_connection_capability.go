package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// CloudConnectionCapability 租户能力实例。
type CloudConnectionCapability struct {

	// 实例ID。
	Id string `json:"id"`

	// 实例所属账号ID。
	DomainId string `json:"domain_id"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 实例创建时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 实例更新时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 租户能力类型，分为： 1. v2：V2的API 2. v3：V3的API 3. billing_mode_period_reduce: 实时降配（包周期） 4. billing_mode_demand: 按需计费（每小时统计打点） 5. bwp95: 按需计费-95（传统95计费方式） 6. bwp95Avg: 按需计费-日95（95峰值平均） 7. network-quality: 丢包和时延统计 8. er：是否支持ER 9. domain_bandwidth：租户带宽值 10. ipv6: 是否支持IPV6 11. ipv6_support_regions: IPV6支持的region列表
	ResourceType *CloudConnectionCapabilityResourceType `json:"resource_type,omitempty"`

	Bandwidth *BandwidthCapability `json:"bandwidth,omitempty"`

	// 租户支持的REGION列表。
	SupportRegions *[]string `json:"support_regions,omitempty"`
}

func (o CloudConnectionCapability) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudConnectionCapability struct{}"
	}

	return strings.Join([]string{"CloudConnectionCapability", string(data)}, " ")
}

type CloudConnectionCapabilityResourceType struct {
	value string
}

type CloudConnectionCapabilityResourceTypeEnum struct {
	V2                         CloudConnectionCapabilityResourceType
	V3                         CloudConnectionCapabilityResourceType
	BILLING_MODE_PERIOD_REDUCE CloudConnectionCapabilityResourceType
	BILLING_MODE_DEMAND        CloudConnectionCapabilityResourceType
	BWP95                      CloudConnectionCapabilityResourceType
	BWP95_AVG                  CloudConnectionCapabilityResourceType
	NETWORK_QUALITY            CloudConnectionCapabilityResourceType
	ER                         CloudConnectionCapabilityResourceType
	DOMAIN_BANDWIDTH           CloudConnectionCapabilityResourceType
	IPV6                       CloudConnectionCapabilityResourceType
	IPV6_SUPPORT_REGIONS       CloudConnectionCapabilityResourceType
}

func GetCloudConnectionCapabilityResourceTypeEnum() CloudConnectionCapabilityResourceTypeEnum {
	return CloudConnectionCapabilityResourceTypeEnum{
		V2: CloudConnectionCapabilityResourceType{
			value: "v2",
		},
		V3: CloudConnectionCapabilityResourceType{
			value: "v3",
		},
		BILLING_MODE_PERIOD_REDUCE: CloudConnectionCapabilityResourceType{
			value: "billing_mode_period_reduce",
		},
		BILLING_MODE_DEMAND: CloudConnectionCapabilityResourceType{
			value: "billing_mode_demand",
		},
		BWP95: CloudConnectionCapabilityResourceType{
			value: "bwp95",
		},
		BWP95_AVG: CloudConnectionCapabilityResourceType{
			value: "bwp95Avg",
		},
		NETWORK_QUALITY: CloudConnectionCapabilityResourceType{
			value: "network-quality",
		},
		ER: CloudConnectionCapabilityResourceType{
			value: "er",
		},
		DOMAIN_BANDWIDTH: CloudConnectionCapabilityResourceType{
			value: "domain_bandwidth",
		},
		IPV6: CloudConnectionCapabilityResourceType{
			value: "ipv6",
		},
		IPV6_SUPPORT_REGIONS: CloudConnectionCapabilityResourceType{
			value: "ipv6_support_regions",
		},
	}
}

func (c CloudConnectionCapabilityResourceType) Value() string {
	return c.value
}

func (c CloudConnectionCapabilityResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CloudConnectionCapabilityResourceType) UnmarshalJSON(b []byte) error {
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
