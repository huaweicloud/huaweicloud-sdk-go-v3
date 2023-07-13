package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AssociateDomainV2Response Response Object
type AssociateDomainV2Response struct {

	// 自定义域名
	UrlDomain *string `json:"url_domain,omitempty"`

	// 自定义域名的编号
	Id *string `json:"id,omitempty"`

	// CNAME解析状态 - 1: 未解析 - 2: 解析中 - 3: 解析成功 - 4: 解析失败
	Status *AssociateDomainV2ResponseStatus `json:"status,omitempty"`

	// 支持的最小SSL版本
	MinSslVersion  *string `json:"min_ssl_version,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AssociateDomainV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateDomainV2Response struct{}"
	}

	return strings.Join([]string{"AssociateDomainV2Response", string(data)}, " ")
}

type AssociateDomainV2ResponseStatus struct {
	value int32
}

type AssociateDomainV2ResponseStatusEnum struct {
	E_1 AssociateDomainV2ResponseStatus
	E_2 AssociateDomainV2ResponseStatus
	E_3 AssociateDomainV2ResponseStatus
	E_4 AssociateDomainV2ResponseStatus
}

func GetAssociateDomainV2ResponseStatusEnum() AssociateDomainV2ResponseStatusEnum {
	return AssociateDomainV2ResponseStatusEnum{
		E_1: AssociateDomainV2ResponseStatus{
			value: 1,
		}, E_2: AssociateDomainV2ResponseStatus{
			value: 2,
		}, E_3: AssociateDomainV2ResponseStatus{
			value: 3,
		}, E_4: AssociateDomainV2ResponseStatus{
			value: 4,
		},
	}
}

func (c AssociateDomainV2ResponseStatus) Value() int32 {
	return c.value
}

func (c AssociateDomainV2ResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AssociateDomainV2ResponseStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
