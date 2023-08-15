package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateDomainsResponse Response Object
type CreateDomainsResponse struct {

	// 状态码:   * success - 成功   * failure - 失败
	InfoCode *CreateDomainsResponseInfoCode `json:"info_code,omitempty"`

	// 返回的提示信息
	InfoDescription *string `json:"info_description,omitempty"`

	// 网站域名ID
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

func (c CreateDomainsResponseInfoCode) Value() string {
	return c.value
}

func (c CreateDomainsResponseInfoCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDomainsResponseInfoCode) UnmarshalJSON(b []byte) error {
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
