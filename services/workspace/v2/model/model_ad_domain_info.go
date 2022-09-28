package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 域信息。
type AdDomainInfo struct {

	// 域类型。 - LITE_AS：LiteAS。 - LOCAL_AD：本地AD。  说明：域类型为“LOCAL_AD”时，请确保所选VPC网络与“LOCAL_AD”所属网络可连通。
	DomainType AdDomainInfoDomainType `json:"domain_type"`

	// 域管理员帐号。
	DomainAdminAccount string `json:"domain_admin_account"`

	// 域管理员账号密码。
	DomainPassword string `json:"domain_password"`
}

func (o AdDomainInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdDomainInfo struct{}"
	}

	return strings.Join([]string{"AdDomainInfo", string(data)}, " ")
}

type AdDomainInfoDomainType struct {
	value string
}

type AdDomainInfoDomainTypeEnum struct {
	LITE_AS  AdDomainInfoDomainType
	LOCAL_AD AdDomainInfoDomainType
}

func GetAdDomainInfoDomainTypeEnum() AdDomainInfoDomainTypeEnum {
	return AdDomainInfoDomainTypeEnum{
		LITE_AS: AdDomainInfoDomainType{
			value: "LITE_AS",
		},
		LOCAL_AD: AdDomainInfoDomainType{
			value: "LOCAL_AD",
		},
	}
}

func (c AdDomainInfoDomainType) Value() string {
	return c.value
}

func (c AdDomainInfoDomainType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AdDomainInfoDomainType) UnmarshalJSON(b []byte) error {
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
