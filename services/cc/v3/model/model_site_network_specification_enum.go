package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SiteNetworkSpecificationEnum 分支网络的规格类型定义： - site-network.is-support: 支持分支网络。 - site-network.is-support-enterprise-project: 支持分支网络的企业项目。 - site-network.is-support-tag: 支持分支网络的标签能力。 - site-network.is-support-intra-region: 支持同region创建分支网络。 - site-network.support-topologies: 支持分支网络拓扑。 - site-network.support-regions: 支持分支接入的Region列表。 - site-network.support-dscp-regions: 支持分支接入的Region列表 - site-network.support-freeze-regions: 支持分支网络冻结的Region列表 - site-network.support-locations: 支持分支接入点列表。 - site-connection-bandwidth.size-range: 分支连接带宽大小的范围。 - site-connection-bandwidth.charge-mode: 分支连接带宽计费类型。 - site-connection-bandwidth.free-line: 分支连接带宽免费线路。
type SiteNetworkSpecificationEnum struct {
	value string
}

type SiteNetworkSpecificationEnumEnum struct {
	SITE_NETWORK_IS_SUPPORT                    SiteNetworkSpecificationEnum
	SITE_NETWORK_IS_SUPPORT_ENTERPRISE_PROJECT SiteNetworkSpecificationEnum
	SITE_NETWORK_IS_SUPPORT_TAG                SiteNetworkSpecificationEnum
	SITE_NETWORK_IS_SUPPORT_INTRA_REGION       SiteNetworkSpecificationEnum
	SITE_NETWORK_SUPPORT_TOPOLOGIES            SiteNetworkSpecificationEnum
	SITE_NETWORK_SUPPORT_REGIONS               SiteNetworkSpecificationEnum
	SITE_NETWORK_SUPPORT_DSCP_REGIONS          SiteNetworkSpecificationEnum
	SITE_NETWORK_SUPPORT_FREEZE_REGIONS        SiteNetworkSpecificationEnum
	SITE_NETWORK_SUPPORT_LOCATIONS             SiteNetworkSpecificationEnum
	SITE_CONNECTION_BANDWIDTH_SIZE_RANGE       SiteNetworkSpecificationEnum
	SITE_CONNECTION_BANDWIDTH_CHARGE_MODE      SiteNetworkSpecificationEnum
	SITE_CONNECTION_BANDWIDTH_FREE_LINE        SiteNetworkSpecificationEnum
}

func GetSiteNetworkSpecificationEnumEnum() SiteNetworkSpecificationEnumEnum {
	return SiteNetworkSpecificationEnumEnum{
		SITE_NETWORK_IS_SUPPORT: SiteNetworkSpecificationEnum{
			value: "site-network.is-support",
		},
		SITE_NETWORK_IS_SUPPORT_ENTERPRISE_PROJECT: SiteNetworkSpecificationEnum{
			value: "site-network.is-support-enterprise-project",
		},
		SITE_NETWORK_IS_SUPPORT_TAG: SiteNetworkSpecificationEnum{
			value: "site-network.is-support-tag",
		},
		SITE_NETWORK_IS_SUPPORT_INTRA_REGION: SiteNetworkSpecificationEnum{
			value: "site-network.is-support-intra-region",
		},
		SITE_NETWORK_SUPPORT_TOPOLOGIES: SiteNetworkSpecificationEnum{
			value: "site-network.support-topologies",
		},
		SITE_NETWORK_SUPPORT_REGIONS: SiteNetworkSpecificationEnum{
			value: "site-network.support-regions",
		},
		SITE_NETWORK_SUPPORT_DSCP_REGIONS: SiteNetworkSpecificationEnum{
			value: "site-network.support-dscp-regions",
		},
		SITE_NETWORK_SUPPORT_FREEZE_REGIONS: SiteNetworkSpecificationEnum{
			value: "site-network.support-freeze-regions",
		},
		SITE_NETWORK_SUPPORT_LOCATIONS: SiteNetworkSpecificationEnum{
			value: "site-network.support-locations",
		},
		SITE_CONNECTION_BANDWIDTH_SIZE_RANGE: SiteNetworkSpecificationEnum{
			value: "site-connection-bandwidth.size-range",
		},
		SITE_CONNECTION_BANDWIDTH_CHARGE_MODE: SiteNetworkSpecificationEnum{
			value: "site-connection-bandwidth.charge-mode",
		},
		SITE_CONNECTION_BANDWIDTH_FREE_LINE: SiteNetworkSpecificationEnum{
			value: "site-connection-bandwidth.free-line",
		},
	}
}

func (c SiteNetworkSpecificationEnum) Value() string {
	return c.value
}

func (c SiteNetworkSpecificationEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SiteNetworkSpecificationEnum) UnmarshalJSON(b []byte) error {
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
