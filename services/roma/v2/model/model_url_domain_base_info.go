package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type UrlDomainBaseInfo struct {
	// 自定义域名

	UrlDomain string `json:"url_domain"`
	// 自定义域名的编号

	Id string `json:"id"`
	// CNAME解析状态 - 1: 未解析 - 2: 解析中 - 3: 解析成功 - 4: 解析失败

	Status UrlDomainBaseInfoStatus `json:"status"`
	// 支持的最小SSL版本

	MinSslVersion string `json:"min_ssl_version"`
}

func (o UrlDomainBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UrlDomainBaseInfo struct{}"
	}

	return strings.Join([]string{"UrlDomainBaseInfo", string(data)}, " ")
}

type UrlDomainBaseInfoStatus struct {
	value int32
}

type UrlDomainBaseInfoStatusEnum struct {
	E_1 UrlDomainBaseInfoStatus
	E_2 UrlDomainBaseInfoStatus
	E_3 UrlDomainBaseInfoStatus
	E_4 UrlDomainBaseInfoStatus
}

func GetUrlDomainBaseInfoStatusEnum() UrlDomainBaseInfoStatusEnum {
	return UrlDomainBaseInfoStatusEnum{
		E_1: UrlDomainBaseInfoStatus{
			value: 1,
		}, E_2: UrlDomainBaseInfoStatus{
			value: 2,
		}, E_3: UrlDomainBaseInfoStatus{
			value: 3,
		}, E_4: UrlDomainBaseInfoStatus{
			value: 4,
		},
	}
}

func (c UrlDomainBaseInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UrlDomainBaseInfoStatus) UnmarshalJSON(b []byte) error {
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
