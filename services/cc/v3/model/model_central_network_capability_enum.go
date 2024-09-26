package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CentralNetworkCapabilityEnum 能力类型定义： - central-network.is-support: 支持中心网络 - central-network.is-support-enterprise-project: 支持中心网络的企业项目 - central-network.is-support-tag: 支持中心网络的标签能力 - central-network.is-support-custom-er-table: 支持中心网络的自定义路由表能力 - connection-bandwidth.size-range: 跨地域连接带宽大小的范围 - connection-bandwidth.charge-mode: 跨地域连接带宽计费类型 - connection-bandwidth.free-line: 跨地域连接免费线路 - er-instance.support-regions: 支持ER实例的Region列表 - er-instance.support-ipv6-regions: 支持IPV6的ER实例Region列表 - er-instance.support-dscp-regions:支持带宽包线路等级配置的region列表 - er-instance.support-sts5-regions:支持sts5调用ER服务的region列表 - er-instance.support-sites: 支持ER实例的站点列表 - gdgw-instance.support-dscp-regions:支持GDGW全域互联带宽线路等级配置的region列表 - gdgw-instance.support-freeze-regions:支持GDGW全域互联带宽冻结的region列表 - gdgw-attachment.is-support: 支持GDGW附件 - gdgw-attachment.support-regions: 支持GDGW附件的Region列表 - gdgw-attachment.support-sites: 支持GDGW附件的站点列表 - er-route-table-attachment.is-support: 支持路由表附件 - er-route-table-attachment.support-regions: 支持路由表附件的Region列表 - er-route-table-attachment.support-sites: 支持路由表附件的站点列表 - cloud-alliance.support-regions: 支持云联盟场景的Region列表
type CentralNetworkCapabilityEnum struct {
	value string
}

type CentralNetworkCapabilityEnumEnum struct {
	CENTRAL_NETWORK_IS_SUPPORT                    CentralNetworkCapabilityEnum
	CENTRAL_NETWORK_IS_SUPPORT_ENTERPRISE_PROJECT CentralNetworkCapabilityEnum
	CENTRAL_NETWORK_IS_SUPPORT_TAG                CentralNetworkCapabilityEnum
	CENTRAL_NETWORK_IS_SUPPORT_CUSTOM_ER_TABLE    CentralNetworkCapabilityEnum
	CONNECTION_BANDWIDTH_SIZE_RANGE               CentralNetworkCapabilityEnum
	CONNECTION_BANDWIDTH_CHARGE_MODE              CentralNetworkCapabilityEnum
	CONNECTION_BANDWIDTH_FREE_LINE                CentralNetworkCapabilityEnum
	ER_INSTANCE_SUPPORT_REGIONS                   CentralNetworkCapabilityEnum
	ER_INSTANCE_SUPPORT_IPV6_REGIONS              CentralNetworkCapabilityEnum
	ER_INSTANCE_SUPPORT_DSCP_REGIONS              CentralNetworkCapabilityEnum
	ER_INSTANCE_SUPPORT_STS5_REGIONS              CentralNetworkCapabilityEnum
	ER_INSTANCE_SUPPORT_SITES                     CentralNetworkCapabilityEnum
	GDGW_INSTANCE_SUPPORT_DSCP_REGIONS            CentralNetworkCapabilityEnum
	GDGW_INSTANCE_SUPPORT_FREEZE_REGIONS          CentralNetworkCapabilityEnum
	GDGW_ATTACHMENT_IS_SUPPORT                    CentralNetworkCapabilityEnum
	GDGW_ATTACHMENT_SUPPORT_REGIONS               CentralNetworkCapabilityEnum
	GDGW_ATTACHMENT_SUPPORT_SITES                 CentralNetworkCapabilityEnum
	ER_ROUTE_TABLE_ATTACHMENT_IS_SUPPORT          CentralNetworkCapabilityEnum
	ER_ROUTE_TABLE_ATTACHMENT_SUPPORT_REGIONS     CentralNetworkCapabilityEnum
	ER_ROUTE_TABLE_ATTACHMENT_SUPPORT_SITES       CentralNetworkCapabilityEnum
	CLOUD_ALLIANCE_SUPPORT_REGIONS                CentralNetworkCapabilityEnum
}

func GetCentralNetworkCapabilityEnumEnum() CentralNetworkCapabilityEnumEnum {
	return CentralNetworkCapabilityEnumEnum{
		CENTRAL_NETWORK_IS_SUPPORT: CentralNetworkCapabilityEnum{
			value: "central-network.is-support",
		},
		CENTRAL_NETWORK_IS_SUPPORT_ENTERPRISE_PROJECT: CentralNetworkCapabilityEnum{
			value: "central-network.is-support-enterprise-project",
		},
		CENTRAL_NETWORK_IS_SUPPORT_TAG: CentralNetworkCapabilityEnum{
			value: "central-network.is-support-tag",
		},
		CENTRAL_NETWORK_IS_SUPPORT_CUSTOM_ER_TABLE: CentralNetworkCapabilityEnum{
			value: "central-network.is-support-custom-er-table",
		},
		CONNECTION_BANDWIDTH_SIZE_RANGE: CentralNetworkCapabilityEnum{
			value: "connection-bandwidth.size-range",
		},
		CONNECTION_BANDWIDTH_CHARGE_MODE: CentralNetworkCapabilityEnum{
			value: "connection-bandwidth.charge-mode",
		},
		CONNECTION_BANDWIDTH_FREE_LINE: CentralNetworkCapabilityEnum{
			value: "connection-bandwidth.free-line",
		},
		ER_INSTANCE_SUPPORT_REGIONS: CentralNetworkCapabilityEnum{
			value: "er-instance.support-regions",
		},
		ER_INSTANCE_SUPPORT_IPV6_REGIONS: CentralNetworkCapabilityEnum{
			value: "er-instance.support-ipv6-regions",
		},
		ER_INSTANCE_SUPPORT_DSCP_REGIONS: CentralNetworkCapabilityEnum{
			value: "er-instance.support-dscp-regions",
		},
		ER_INSTANCE_SUPPORT_STS5_REGIONS: CentralNetworkCapabilityEnum{
			value: "er-instance.support-sts5-regions",
		},
		ER_INSTANCE_SUPPORT_SITES: CentralNetworkCapabilityEnum{
			value: "er-instance.support-sites",
		},
		GDGW_INSTANCE_SUPPORT_DSCP_REGIONS: CentralNetworkCapabilityEnum{
			value: "gdgw-instance.support-dscp-regions",
		},
		GDGW_INSTANCE_SUPPORT_FREEZE_REGIONS: CentralNetworkCapabilityEnum{
			value: "gdgw-instance.support-freeze-regions",
		},
		GDGW_ATTACHMENT_IS_SUPPORT: CentralNetworkCapabilityEnum{
			value: "gdgw-attachment.is-support",
		},
		GDGW_ATTACHMENT_SUPPORT_REGIONS: CentralNetworkCapabilityEnum{
			value: "gdgw-attachment.support-regions",
		},
		GDGW_ATTACHMENT_SUPPORT_SITES: CentralNetworkCapabilityEnum{
			value: "gdgw-attachment.support-sites",
		},
		ER_ROUTE_TABLE_ATTACHMENT_IS_SUPPORT: CentralNetworkCapabilityEnum{
			value: "er-route-table-attachment.is-support",
		},
		ER_ROUTE_TABLE_ATTACHMENT_SUPPORT_REGIONS: CentralNetworkCapabilityEnum{
			value: "er-route-table-attachment.support-regions",
		},
		ER_ROUTE_TABLE_ATTACHMENT_SUPPORT_SITES: CentralNetworkCapabilityEnum{
			value: "er-route-table-attachment.support-sites",
		},
		CLOUD_ALLIANCE_SUPPORT_REGIONS: CentralNetworkCapabilityEnum{
			value: "cloud-alliance.support-regions",
		},
	}
}

func (c CentralNetworkCapabilityEnum) Value() string {
	return c.value
}

func (c CentralNetworkCapabilityEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CentralNetworkCapabilityEnum) UnmarshalJSON(b []byte) error {
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
