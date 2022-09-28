package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ModifyWorkspaceAttributesReq struct {
	AdInfo *AdDomainInfo `json:"ad_info,omitempty"`

	AdDomains *AdDomain `json:"ad_domains,omitempty"`

	// 接入模式。 - INTERNET：互联网接入。 - DEDICATED：专线接入。 - BOTH：代表两种接入方式都支持。
	AccessMode *ModifyWorkspaceAttributesReqAccessMode `json:"access_mode,omitempty"`

	// 专线接入网段列表，多个网段信息用分号隔开，列表长度不超过5。
	DedicatedSubnets *string `json:"dedicated_subnets,omitempty"`

	// 子网的网络ID列表。
	SubnetIds *[]string `json:"subnet_ids,omitempty"`

	// 互联网接入端口。
	InternetAccessPort *string `json:"internet_access_port,omitempty"`

	// 企业ID。
	EnterpriseId *string `json:"enterprise_id,omitempty"`
}

func (o ModifyWorkspaceAttributesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyWorkspaceAttributesReq struct{}"
	}

	return strings.Join([]string{"ModifyWorkspaceAttributesReq", string(data)}, " ")
}

type ModifyWorkspaceAttributesReqAccessMode struct {
	value string
}

type ModifyWorkspaceAttributesReqAccessModeEnum struct {
	INTERNET  ModifyWorkspaceAttributesReqAccessMode
	DEDICATED ModifyWorkspaceAttributesReqAccessMode
	BOTH      ModifyWorkspaceAttributesReqAccessMode
}

func GetModifyWorkspaceAttributesReqAccessModeEnum() ModifyWorkspaceAttributesReqAccessModeEnum {
	return ModifyWorkspaceAttributesReqAccessModeEnum{
		INTERNET: ModifyWorkspaceAttributesReqAccessMode{
			value: "INTERNET",
		},
		DEDICATED: ModifyWorkspaceAttributesReqAccessMode{
			value: "DEDICATED",
		},
		BOTH: ModifyWorkspaceAttributesReqAccessMode{
			value: "BOTH",
		},
	}
}

func (c ModifyWorkspaceAttributesReqAccessMode) Value() string {
	return c.value
}

func (c ModifyWorkspaceAttributesReqAccessMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyWorkspaceAttributesReqAccessMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
