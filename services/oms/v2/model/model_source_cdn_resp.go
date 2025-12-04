package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SourceCdnResp 源端CDN配置返回值。
type SourceCdnResp struct {

	//   从指定域名获取对象。
	Domain string `json:"domain"`

	// 协议类型，支持http和https协议。
	Protocol SourceCdnRespProtocol `json:"protocol"`

	// 鉴权类型: NONE：公开访问，无安全限制, QINIU_PRIVATE_AUTHENTICATION：七牛自定义URL签名, ALIYUN_OSS_A：阿里云  URL携带签名，简单通用, ALIYUN_OSS_B：阿里云  Header携带签名，用于API调用, ALIYUN_OSS_C：阿里云  STS临时安全令牌，最安全, KSYUN_PRIVATE_AUTHENTICATION：金山云  金山云自定义URL签名, AZURE_SAS_TOKEN：微软Azure  灵活安全的共享访问签名, TENCENT_COS_A:腾讯云  多次有效签名（不推荐）, TENCENT_COS_B:腾讯云  单次有效签名，安全性最高, TENCENT_COS_C:腾讯云  Header携带签名，用于API调用, TENCENT_COS_D:腾讯云  Header携带签名，用于API调用.
	AuthenticationType *SourceCdnRespAuthenticationType `json:"authentication_type,omitempty"`
}

func (o SourceCdnResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceCdnResp struct{}"
	}

	return strings.Join([]string{"SourceCdnResp", string(data)}, " ")
}

type SourceCdnRespProtocol struct {
	value string
}

type SourceCdnRespProtocolEnum struct {
	HTTP  SourceCdnRespProtocol
	HTTPS SourceCdnRespProtocol
}

func GetSourceCdnRespProtocolEnum() SourceCdnRespProtocolEnum {
	return SourceCdnRespProtocolEnum{
		HTTP: SourceCdnRespProtocol{
			value: "http",
		},
		HTTPS: SourceCdnRespProtocol{
			value: "https",
		},
	}
}

func (c SourceCdnRespProtocol) Value() string {
	return c.value
}

func (c SourceCdnRespProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceCdnRespProtocol) UnmarshalJSON(b []byte) error {
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

type SourceCdnRespAuthenticationType struct {
	value string
}

type SourceCdnRespAuthenticationTypeEnum struct {
	NONE                         SourceCdnRespAuthenticationType
	QINIU_PRIVATE_AUTHENTICATION SourceCdnRespAuthenticationType
	ALIYUN_OSS_A                 SourceCdnRespAuthenticationType
	ALIYUN_OSS_B                 SourceCdnRespAuthenticationType
	ALIYUN_OSS_C                 SourceCdnRespAuthenticationType
	KSYUN_PRIVATE_AUTHENTICATION SourceCdnRespAuthenticationType
	AZURE_SAS_TOKEN              SourceCdnRespAuthenticationType
	TENCENT_COS_A                SourceCdnRespAuthenticationType
	TENCENT_COS_B                SourceCdnRespAuthenticationType
	TENCENT_COS_C                SourceCdnRespAuthenticationType
	TENCENT_COS_D                SourceCdnRespAuthenticationType
}

func GetSourceCdnRespAuthenticationTypeEnum() SourceCdnRespAuthenticationTypeEnum {
	return SourceCdnRespAuthenticationTypeEnum{
		NONE: SourceCdnRespAuthenticationType{
			value: "NONE",
		},
		QINIU_PRIVATE_AUTHENTICATION: SourceCdnRespAuthenticationType{
			value: "QINIU_PRIVATE_AUTHENTICATION",
		},
		ALIYUN_OSS_A: SourceCdnRespAuthenticationType{
			value: "ALIYUN_OSS_A",
		},
		ALIYUN_OSS_B: SourceCdnRespAuthenticationType{
			value: "ALIYUN_OSS_B",
		},
		ALIYUN_OSS_C: SourceCdnRespAuthenticationType{
			value: "ALIYUN_OSS_C",
		},
		KSYUN_PRIVATE_AUTHENTICATION: SourceCdnRespAuthenticationType{
			value: "KSYUN_PRIVATE_AUTHENTICATION",
		},
		AZURE_SAS_TOKEN: SourceCdnRespAuthenticationType{
			value: "AZURE_SAS_TOKEN",
		},
		TENCENT_COS_A: SourceCdnRespAuthenticationType{
			value: "TENCENT_COS_A",
		},
		TENCENT_COS_B: SourceCdnRespAuthenticationType{
			value: "TENCENT_COS_B",
		},
		TENCENT_COS_C: SourceCdnRespAuthenticationType{
			value: "TENCENT_COS_C",
		},
		TENCENT_COS_D: SourceCdnRespAuthenticationType{
			value: "TENCENT_COS_D",
		},
	}
}

func (c SourceCdnRespAuthenticationType) Value() string {
	return c.value
}

func (c SourceCdnRespAuthenticationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceCdnRespAuthenticationType) UnmarshalJSON(b []byte) error {
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
