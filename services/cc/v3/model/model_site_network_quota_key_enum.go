package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SiteNetworkQuotaKeyEnum 分支网络配额类型。 - site_networks_per_account (每个账户的分支网络数) - sites_per_mesh_site_network (网状分支网络的分支数) - spoke_sites_per_star_site_network (星状分支网络的Spoke分支数) - sites_per_hybrid_site_network (混合分支网络的分支数)
type SiteNetworkQuotaKeyEnum struct {
	value string
}

type SiteNetworkQuotaKeyEnumEnum struct {
	SITE_NETWORKS_PER_ACCOUNT         SiteNetworkQuotaKeyEnum
	SITES_PER_MESH_SITE_NETWORK       SiteNetworkQuotaKeyEnum
	SPOKE_SITES_PER_STAR_SITE_NETWORK SiteNetworkQuotaKeyEnum
	SITES_PER_HYBRID_SITE_NETWORK     SiteNetworkQuotaKeyEnum
}

func GetSiteNetworkQuotaKeyEnumEnum() SiteNetworkQuotaKeyEnumEnum {
	return SiteNetworkQuotaKeyEnumEnum{
		SITE_NETWORKS_PER_ACCOUNT: SiteNetworkQuotaKeyEnum{
			value: "site_networks_per_account",
		},
		SITES_PER_MESH_SITE_NETWORK: SiteNetworkQuotaKeyEnum{
			value: "sites_per_mesh_site_network",
		},
		SPOKE_SITES_PER_STAR_SITE_NETWORK: SiteNetworkQuotaKeyEnum{
			value: "spoke_sites_per_star_site_network",
		},
		SITES_PER_HYBRID_SITE_NETWORK: SiteNetworkQuotaKeyEnum{
			value: "sites_per_hybrid_site_network",
		},
	}
}

func (c SiteNetworkQuotaKeyEnum) Value() string {
	return c.value
}

func (c SiteNetworkQuotaKeyEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SiteNetworkQuotaKeyEnum) UnmarshalJSON(b []byte) error {
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
