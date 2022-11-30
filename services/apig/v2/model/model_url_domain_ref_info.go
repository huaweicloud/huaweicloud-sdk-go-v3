package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 域名详情及关联的证书、分组信息
type UrlDomainRefInfo struct {

	// 自定义域名
	UrlDomain string `json:"url_domain"`

	// 自定义域名的编号
	Id string `json:"id"`

	// CNAME解析状态 - 1: 未解析 - 2: 解析中 - 3: 解析成功 - 4: 解析失败
	Status UrlDomainRefInfoStatus `json:"status"`

	// 支持的最小SSL版本
	MinSslVersion string `json:"min_ssl_version"`

	// 是否开启http到https的重定向，false为关闭，true为开启，默认为false
	IsHttpRedirectToHttps *bool `json:"is_http_redirect_to_https,omitempty"`

	// 证书ID
	SslId *string `json:"ssl_id,omitempty"`

	// 证书名称
	SslName *string `json:"ssl_name,omitempty"`

	// 所属API分组ID
	ApiGroupId string `json:"api_group_id"`

	// 所属API分组名称
	ApiGroupName string `json:"api_group_name"`

	// 所属实例ID
	InstanceId string `json:"instance_id"`
}

func (o UrlDomainRefInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UrlDomainRefInfo struct{}"
	}

	return strings.Join([]string{"UrlDomainRefInfo", string(data)}, " ")
}

type UrlDomainRefInfoStatus struct {
	value int32
}

type UrlDomainRefInfoStatusEnum struct {
	E_1 UrlDomainRefInfoStatus
	E_2 UrlDomainRefInfoStatus
	E_3 UrlDomainRefInfoStatus
	E_4 UrlDomainRefInfoStatus
}

func GetUrlDomainRefInfoStatusEnum() UrlDomainRefInfoStatusEnum {
	return UrlDomainRefInfoStatusEnum{
		E_1: UrlDomainRefInfoStatus{
			value: 1,
		}, E_2: UrlDomainRefInfoStatus{
			value: 2,
		}, E_3: UrlDomainRefInfoStatus{
			value: 3,
		}, E_4: UrlDomainRefInfoStatus{
			value: 4,
		},
	}
}

func (c UrlDomainRefInfoStatus) Value() int32 {
	return c.value
}

func (c UrlDomainRefInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UrlDomainRefInfoStatus) UnmarshalJSON(b []byte) error {
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
