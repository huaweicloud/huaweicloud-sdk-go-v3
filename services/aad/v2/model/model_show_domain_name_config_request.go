package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDomainNameConfigRequest Request Object
type ShowDomainNameConfigRequest struct {

	// 域名id
	DomainId string `json:"domain_id"`

	// 请求类型 domain、waf
	Type ShowDomainNameConfigRequestType `json:"type"`
}

func (o ShowDomainNameConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainNameConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainNameConfigRequest", string(data)}, " ")
}

type ShowDomainNameConfigRequestType struct {
	value string
}

type ShowDomainNameConfigRequestTypeEnum struct {
	DOMAIN ShowDomainNameConfigRequestType
	WAF    ShowDomainNameConfigRequestType
}

func GetShowDomainNameConfigRequestTypeEnum() ShowDomainNameConfigRequestTypeEnum {
	return ShowDomainNameConfigRequestTypeEnum{
		DOMAIN: ShowDomainNameConfigRequestType{
			value: "DOMAIN",
		},
		WAF: ShowDomainNameConfigRequestType{
			value: "WAF",
		},
	}
}

func (c ShowDomainNameConfigRequestType) Value() string {
	return c.value
}

func (c ShowDomainNameConfigRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDomainNameConfigRequestType) UnmarshalJSON(b []byte) error {
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
