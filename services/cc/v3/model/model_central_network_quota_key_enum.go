package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CentralNetworkQuotaKeyEnum 中心网络配额类型。 - central_networks_per_account (每个账户的中心网数) - policy_versions_per_central_network (每个中心网的策略数) - size_of_document_per_central_network_policy_version (中心网络策略文档大小(KB)) - planes_per_central_network (每个中心网的平面数) - er_instances_per_region_per_central_network (每个中心网络每个区域的Er实例数) - connections_per_central_network (每个中心网的连接数) - attachments_per_central_network (每个中心网的附件数) - GDGW_attachments_per_region_per_central_network (每个中心网每个区域的GDGW附件数) - ER_ROUTE_TABLE_attachments_per_region_per_central_network (每个中心网每个区域的ER_ROUTE_TABLE附件数)
type CentralNetworkQuotaKeyEnum struct {
	value string
}

type CentralNetworkQuotaKeyEnumEnum struct {
	CENTRAL_NETWORKS_PER_ACCOUNT                              CentralNetworkQuotaKeyEnum
	POLICY_VERSIONS_PER_CENTRAL_NETWORK                       CentralNetworkQuotaKeyEnum
	SIZE_OF_DOCUMENT_PER_CENTRAL_NETWORK_POLICY_VERSION       CentralNetworkQuotaKeyEnum
	PLANES_PER_CENTRAL_NETWORK                                CentralNetworkQuotaKeyEnum
	ER_INSTANCES_PER_REGION_PER_CENTRAL_NETWORK               CentralNetworkQuotaKeyEnum
	CONNECTIONS_PER_CENTRAL_NETWORK                           CentralNetworkQuotaKeyEnum
	ATTACHMENTS_PER_CENTRAL_NETWORK                           CentralNetworkQuotaKeyEnum
	GDGW_ATTACHMENTS_PER_REGION_PER_CENTRAL_NETWORK           CentralNetworkQuotaKeyEnum
	ER_ROUTE_TABLE_ATTACHMENTS_PER_REGION_PER_CENTRAL_NETWORK CentralNetworkQuotaKeyEnum
}

func GetCentralNetworkQuotaKeyEnumEnum() CentralNetworkQuotaKeyEnumEnum {
	return CentralNetworkQuotaKeyEnumEnum{
		CENTRAL_NETWORKS_PER_ACCOUNT: CentralNetworkQuotaKeyEnum{
			value: "central_networks_per_account",
		},
		POLICY_VERSIONS_PER_CENTRAL_NETWORK: CentralNetworkQuotaKeyEnum{
			value: "policy_versions_per_central_network",
		},
		SIZE_OF_DOCUMENT_PER_CENTRAL_NETWORK_POLICY_VERSION: CentralNetworkQuotaKeyEnum{
			value: "size_of_document_per_central_network_policy_version",
		},
		PLANES_PER_CENTRAL_NETWORK: CentralNetworkQuotaKeyEnum{
			value: "planes_per_central_network",
		},
		ER_INSTANCES_PER_REGION_PER_CENTRAL_NETWORK: CentralNetworkQuotaKeyEnum{
			value: "er_instances_per_region_per_central_network",
		},
		CONNECTIONS_PER_CENTRAL_NETWORK: CentralNetworkQuotaKeyEnum{
			value: "connections_per_central_network",
		},
		ATTACHMENTS_PER_CENTRAL_NETWORK: CentralNetworkQuotaKeyEnum{
			value: "attachments_per_central_network",
		},
		GDGW_ATTACHMENTS_PER_REGION_PER_CENTRAL_NETWORK: CentralNetworkQuotaKeyEnum{
			value: "GDGW_attachments_per_region_per_central_network",
		},
		ER_ROUTE_TABLE_ATTACHMENTS_PER_REGION_PER_CENTRAL_NETWORK: CentralNetworkQuotaKeyEnum{
			value: "ER_ROUTE_TABLE_attachments_per_region_per_central_network",
		},
	}
}

func (c CentralNetworkQuotaKeyEnum) Value() string {
	return c.value
}

func (c CentralNetworkQuotaKeyEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CentralNetworkQuotaKeyEnum) UnmarshalJSON(b []byte) error {
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
