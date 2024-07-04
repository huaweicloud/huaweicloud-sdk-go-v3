package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UrlDomainRefInfo 域名详情及关联的证书、分组信息
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

	// 是否开启客户端证书校验。只有绑定证书时，该参数才生效。当绑定证书存在trusted_root_ca时，默认开启；当绑定证书不存在trusted_root_ca时，默认关闭。
	VerifiedClientCertificateEnabled *bool `json:"verified_client_certificate_enabled,omitempty"`

	// 访问该域名绑定的http协议入方向端口，-1表示无端口且协议不支持，可使用80默认端口，其他有效端口允许的取值范围为1024~49151，需为实例已开放的HTTP协议的自定义入方向端口。  当创建域名时，该参数未填表示用默认80端口；若填写该参数，则必须同时填写https_port；若要http_port和https_port同时使用默认端口，则两个参数都不填。  当修改域名时，该参数未填表示不修改该端口。
	IngressHttpPort *int32 `json:"ingress_http_port,omitempty"`

	// 访问该域名绑定的http协议入方向端口，-1表示无端口且协议不支持，可使用443默认端口，其他有效端口允许的取值范围为1024~49151，需为实例已开放的HTTPS协议的自定义入方向端口。  当创建域名时，该参数未填表示用默认443端口；若填写该参数，则必须同时填写http_port；若要http_port和https_port同时使用默认端口，则两个参数都不填。  当修改域名时，该参数未填表示不修改该端口。
	IngressHttpsPort *int32 `json:"ingress_https_port,omitempty"`

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
