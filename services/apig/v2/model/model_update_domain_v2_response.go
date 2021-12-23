package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateDomainV2Response struct {
	// 自定义域名

	UrlDomain *string `json:"url_domain,omitempty"`
	// 自定义域名的编号

	Id *string `json:"id,omitempty"`
	// CNAME解析状态 - 1: 未解析 - 2: 解析中 - 3: 解析成功 - 4: 解析失败

	Status *UpdateDomainV2ResponseStatus `json:"status,omitempty"`
	// 支持的最小SSL版本

	MinSslVersion  *string `json:"min_ssl_version,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDomainV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainV2Response struct{}"
	}

	return strings.Join([]string{"UpdateDomainV2Response", string(data)}, " ")
}

type UpdateDomainV2ResponseStatus struct {
	value int32
}

type UpdateDomainV2ResponseStatusEnum struct {
	E_1 UpdateDomainV2ResponseStatus
	E_2 UpdateDomainV2ResponseStatus
	E_3 UpdateDomainV2ResponseStatus
	E_4 UpdateDomainV2ResponseStatus
}

func GetUpdateDomainV2ResponseStatusEnum() UpdateDomainV2ResponseStatusEnum {
	return UpdateDomainV2ResponseStatusEnum{
		E_1: UpdateDomainV2ResponseStatus{
			value: 1,
		}, E_2: UpdateDomainV2ResponseStatus{
			value: 2,
		}, E_3: UpdateDomainV2ResponseStatus{
			value: 3,
		}, E_4: UpdateDomainV2ResponseStatus{
			value: 4,
		},
	}
}

func (c UpdateDomainV2ResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDomainV2ResponseStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
