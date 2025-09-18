package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CloudConnectionCapabilityKeyEnum 租户能力类型，分为： - v2（V2的API） - v3（V3的API） - billing_mode_period_reduce（包周期实时降配） - billing_mode_demand（按需计费） - bwp95（按需计费-95） - bwp95Avg（按需计费-日95） - network-quality（丢包和时延统计） - er（是否支持ER） - domain_bandwidth（租户带宽值） - ipv6（是否支持IPV6） - ipv6_support_regions（IPV6支持的区域列表） - enterprise-cloud-connection.is-support（支持企业版云连接） - enterprise-cloud-connection.support-sites（支持企业版云连接的站点列表） - enterprise-cloud-connection-segment.is-support（支持企业版云连接平面） - enterprise-cloud-connection-dc-attachment.is-support（支持企业版云连接的专线网关连接）
type CloudConnectionCapabilityKeyEnum struct {
	value string
}

type CloudConnectionCapabilityKeyEnumEnum struct {
	V2                                                   CloudConnectionCapabilityKeyEnum
	V3                                                   CloudConnectionCapabilityKeyEnum
	BILLING_MODE_PERIOD_REDUCE                           CloudConnectionCapabilityKeyEnum
	BILLING_MODE_DEMAND                                  CloudConnectionCapabilityKeyEnum
	BWP95                                                CloudConnectionCapabilityKeyEnum
	BWP95_AVG                                            CloudConnectionCapabilityKeyEnum
	NETWORK_QUALITY                                      CloudConnectionCapabilityKeyEnum
	ER                                                   CloudConnectionCapabilityKeyEnum
	DOMAIN_BANDWIDTH                                     CloudConnectionCapabilityKeyEnum
	IPV6                                                 CloudConnectionCapabilityKeyEnum
	IPV6_SUPPORT_REGIONS                                 CloudConnectionCapabilityKeyEnum
	ENTERPRISE_CLOUD_CONNECTION_IS_SUPPORT               CloudConnectionCapabilityKeyEnum
	ENTERPRISE_CLOUD_CONNECTION_SUPPORT_SITES            CloudConnectionCapabilityKeyEnum
	ENTERPRISE_CLOUD_CONNECTION_SEGMENT_IS_SUPPORT       CloudConnectionCapabilityKeyEnum
	ENTERPRISE_CLOUD_CONNECTION_DC_ATTACHMENT_IS_SUPPORT CloudConnectionCapabilityKeyEnum
}

func GetCloudConnectionCapabilityKeyEnumEnum() CloudConnectionCapabilityKeyEnumEnum {
	return CloudConnectionCapabilityKeyEnumEnum{
		V2: CloudConnectionCapabilityKeyEnum{
			value: "v2",
		},
		V3: CloudConnectionCapabilityKeyEnum{
			value: "v3",
		},
		BILLING_MODE_PERIOD_REDUCE: CloudConnectionCapabilityKeyEnum{
			value: "billing_mode_period_reduce",
		},
		BILLING_MODE_DEMAND: CloudConnectionCapabilityKeyEnum{
			value: "billing_mode_demand",
		},
		BWP95: CloudConnectionCapabilityKeyEnum{
			value: "bwp95",
		},
		BWP95_AVG: CloudConnectionCapabilityKeyEnum{
			value: "bwp95Avg",
		},
		NETWORK_QUALITY: CloudConnectionCapabilityKeyEnum{
			value: "network-quality",
		},
		ER: CloudConnectionCapabilityKeyEnum{
			value: "er",
		},
		DOMAIN_BANDWIDTH: CloudConnectionCapabilityKeyEnum{
			value: "domain_bandwidth",
		},
		IPV6: CloudConnectionCapabilityKeyEnum{
			value: "ipv6",
		},
		IPV6_SUPPORT_REGIONS: CloudConnectionCapabilityKeyEnum{
			value: "ipv6_support_regions",
		},
		ENTERPRISE_CLOUD_CONNECTION_IS_SUPPORT: CloudConnectionCapabilityKeyEnum{
			value: "enterprise-cloud-connection.is-support",
		},
		ENTERPRISE_CLOUD_CONNECTION_SUPPORT_SITES: CloudConnectionCapabilityKeyEnum{
			value: "enterprise-cloud-connection.support-sites",
		},
		ENTERPRISE_CLOUD_CONNECTION_SEGMENT_IS_SUPPORT: CloudConnectionCapabilityKeyEnum{
			value: "enterprise-cloud-connection-segment.is-support",
		},
		ENTERPRISE_CLOUD_CONNECTION_DC_ATTACHMENT_IS_SUPPORT: CloudConnectionCapabilityKeyEnum{
			value: "enterprise-cloud-connection-dc-attachment.is-support",
		},
	}
}

func (c CloudConnectionCapabilityKeyEnum) Value() string {
	return c.value
}

func (c CloudConnectionCapabilityKeyEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CloudConnectionCapabilityKeyEnum) UnmarshalJSON(b []byte) error {
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
