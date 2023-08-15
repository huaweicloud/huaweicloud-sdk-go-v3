package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AuthorizeDomainsResponse Response Object
type AuthorizeDomainsResponse struct {

	// 状态码:   * success - 成功   * failure - 失败
	InfoCode *AuthorizeDomainsResponseInfoCode `json:"info_code,omitempty"`

	// 返回的提示信息
	InfoDescription *string `json:"info_description,omitempty"`

	// 网站域名认证使用须知
	UsageNotice    *string `json:"usage_notice,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AuthorizeDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeDomainsResponse struct{}"
	}

	return strings.Join([]string{"AuthorizeDomainsResponse", string(data)}, " ")
}

type AuthorizeDomainsResponseInfoCode struct {
	value string
}

type AuthorizeDomainsResponseInfoCodeEnum struct {
	SUCCESS AuthorizeDomainsResponseInfoCode
	FAILURE AuthorizeDomainsResponseInfoCode
}

func GetAuthorizeDomainsResponseInfoCodeEnum() AuthorizeDomainsResponseInfoCodeEnum {
	return AuthorizeDomainsResponseInfoCodeEnum{
		SUCCESS: AuthorizeDomainsResponseInfoCode{
			value: "success",
		},
		FAILURE: AuthorizeDomainsResponseInfoCode{
			value: "failure",
		},
	}
}

func (c AuthorizeDomainsResponseInfoCode) Value() string {
	return c.value
}

func (c AuthorizeDomainsResponseInfoCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthorizeDomainsResponseInfoCode) UnmarshalJSON(b []byte) error {
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
