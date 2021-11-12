package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type CreateDomainsResponse struct {
	// 状态码:   * success - 成功   * failure - 失败

	InfoCode *CreateDomainsResponseInfoCode `json:"info_code,omitempty"`
	// 返回的提示信息

	InfoDescription *string `json:"info_description,omitempty"`
	// 域名ID

	DomainId       *string `json:"domain_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDomainsResponse struct{}"
	}

	return strings.Join([]string{"CreateDomainsResponse", string(data)}, " ")
}

type CreateDomainsResponseInfoCode struct {
	value string
}

type CreateDomainsResponseInfoCodeEnum struct {
	SUCCESS CreateDomainsResponseInfoCode
	FAILURE CreateDomainsResponseInfoCode
}

func GetCreateDomainsResponseInfoCodeEnum() CreateDomainsResponseInfoCodeEnum {
	return CreateDomainsResponseInfoCodeEnum{
		SUCCESS: CreateDomainsResponseInfoCode{
			value: "success",
		},
		FAILURE: CreateDomainsResponseInfoCode{
			value: "failure",
		},
	}
}

func (c CreateDomainsResponseInfoCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDomainsResponseInfoCode) UnmarshalJSON(b []byte) error {
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
