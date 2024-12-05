package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreatePeerLinkRequestBodyPeerLinkPeerSite struct {
	Type *CreatePeerLinkRequestBodyPeerLinkPeerSiteType `json:"type,omitempty"`

	GatewayId string `json:"gateway_id"`

	ProjectId string `json:"project_id"`

	RegionId string `json:"region_id"`
}

func (o CreatePeerLinkRequestBodyPeerLinkPeerSite) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePeerLinkRequestBodyPeerLinkPeerSite struct{}"
	}

	return strings.Join([]string{"CreatePeerLinkRequestBodyPeerLinkPeerSite", string(data)}, " ")
}

type CreatePeerLinkRequestBodyPeerLinkPeerSiteType struct {
	value string
}

type CreatePeerLinkRequestBodyPeerLinkPeerSiteTypeEnum struct {
	ER CreatePeerLinkRequestBodyPeerLinkPeerSiteType
}

func GetCreatePeerLinkRequestBodyPeerLinkPeerSiteTypeEnum() CreatePeerLinkRequestBodyPeerLinkPeerSiteTypeEnum {
	return CreatePeerLinkRequestBodyPeerLinkPeerSiteTypeEnum{
		ER: CreatePeerLinkRequestBodyPeerLinkPeerSiteType{
			value: "ER",
		},
	}
}

func (c CreatePeerLinkRequestBodyPeerLinkPeerSiteType) Value() string {
	return c.value
}

func (c CreatePeerLinkRequestBodyPeerLinkPeerSiteType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePeerLinkRequestBodyPeerLinkPeerSiteType) UnmarshalJSON(b []byte) error {
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
