package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SiteNetworkTopologyEnum 拓扑结构。 - p2p(点对点拓扑) - mesh (网状拓扑) - star (星形拓扑) - hybrid (混合拓扑)
type SiteNetworkTopologyEnum struct {
	value string
}

type SiteNetworkTopologyEnumEnum struct {
	P2P    SiteNetworkTopologyEnum
	MESH   SiteNetworkTopologyEnum
	STAR   SiteNetworkTopologyEnum
	HYBRID SiteNetworkTopologyEnum
}

func GetSiteNetworkTopologyEnumEnum() SiteNetworkTopologyEnumEnum {
	return SiteNetworkTopologyEnumEnum{
		P2P: SiteNetworkTopologyEnum{
			value: "p2p",
		},
		MESH: SiteNetworkTopologyEnum{
			value: "mesh",
		},
		STAR: SiteNetworkTopologyEnum{
			value: "star",
		},
		HYBRID: SiteNetworkTopologyEnum{
			value: "hybrid",
		},
	}
}

func (c SiteNetworkTopologyEnum) Value() string {
	return c.value
}

func (c SiteNetworkTopologyEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SiteNetworkTopologyEnum) UnmarshalJSON(b []byte) error {
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
