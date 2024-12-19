package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PeerSite 对端的网关的描述信息
type PeerSite struct {

	// 对端网关的ID
	GatewayId *string `json:"gateway_id,omitempty"`

	// 对端网关连接的ID(如：对端为ER时为attachment ID,对端为GDGW时为对端的PeerLink Id)
	LinkId *string `json:"link_id,omitempty"`

	// 对端网关所在的Region
	RegionId *string `json:"region_id,omitempty"`

	// 专线全域接入网关对应的站点位置
	SiteCode *string `json:"site_code,omitempty"`

	// 对等体站点的项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 对等体的类型 - ER 企业路由器 - GDGW 全域接入网关
	Type *PeerSiteType `json:"type,omitempty"`
}

func (o PeerSite) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeerSite struct{}"
	}

	return strings.Join([]string{"PeerSite", string(data)}, " ")
}

type PeerSiteType struct {
	value string
}

type PeerSiteTypeEnum struct {
	ER   PeerSiteType
	GDGW PeerSiteType
}

func GetPeerSiteTypeEnum() PeerSiteTypeEnum {
	return PeerSiteTypeEnum{
		ER: PeerSiteType{
			value: "ER",
		},
		GDGW: PeerSiteType{
			value: "GDGW",
		},
	}
}

func (c PeerSiteType) Value() string {
	return c.value
}

func (c PeerSiteType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PeerSiteType) UnmarshalJSON(b []byte) error {
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
