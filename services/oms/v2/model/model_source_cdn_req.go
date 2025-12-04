package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SourceCdnReq 源端CDN配置。
type SourceCdnReq struct {

	// CDN鉴权密钥，如果CDN需要进行鉴权，则此选项为必选。  无需授权：无需配置此项。 Qiniu：无需配置此项。 Aliyun：根据authentication_type指定的鉴权方式配置此项。 KingsoftCloud：无需配置此项。
	AuthenticationKey *string `json:"authentication_key,omitempty"`

	// 鉴权类型: NONE：公开访问，无安全限制, QINIU_PRIVATE_AUTHENTICATION：七牛自定义URL签名, ALIYUN_OSS_A：阿里云  URL携带签名，简单通用, ALIYUN_OSS_B：阿里云  Header携带签名，用于API调用, ALIYUN_OSS_C：阿里云  STS临时安全令牌，最安全, KSYUN_PRIVATE_AUTHENTICATION：金山云  金山云自定义URL签名, AZURE_SAS_TOKEN：微软Azure  灵活安全的共享访问签名, TENCENT_COS_A:腾讯云  多次有效签名（不推荐）, TENCENT_COS_B:腾讯云  单次有效签名，安全性最高, TENCENT_COS_C:腾讯云  Header携带签名，用于API调用, TENCENT_COS_D:腾讯云  Header携带签名，用于API调用.
	AuthenticationType SourceCdnReqAuthenticationType `json:"authentication_type"`

	//   从指定域名获取对象。
	Domain string `json:"domain"`

	// 协议类型，支持http和https协议。
	Protocol SourceCdnReqProtocol `json:"protocol"`
}

func (o SourceCdnReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceCdnReq struct{}"
	}

	return strings.Join([]string{"SourceCdnReq", string(data)}, " ")
}

type SourceCdnReqAuthenticationType struct {
	value string
}

type SourceCdnReqAuthenticationTypeEnum struct {
	NONE                         SourceCdnReqAuthenticationType
	QINIU_PRIVATE_AUTHENTICATION SourceCdnReqAuthenticationType
	ALIYUN_OSS_A                 SourceCdnReqAuthenticationType
	ALIYUN_OSS_B                 SourceCdnReqAuthenticationType
	ALIYUN_OSS_C                 SourceCdnReqAuthenticationType
	KSYUN_PRIVATE_AUTHENTICATION SourceCdnReqAuthenticationType
	AZURE_SAS_TOKEN              SourceCdnReqAuthenticationType
	TENCENT_COS_A                SourceCdnReqAuthenticationType
	TENCENT_COS_B                SourceCdnReqAuthenticationType
	TENCENT_COS_C                SourceCdnReqAuthenticationType
	TENCENT_COS_D                SourceCdnReqAuthenticationType
}

func GetSourceCdnReqAuthenticationTypeEnum() SourceCdnReqAuthenticationTypeEnum {
	return SourceCdnReqAuthenticationTypeEnum{
		NONE: SourceCdnReqAuthenticationType{
			value: "NONE",
		},
		QINIU_PRIVATE_AUTHENTICATION: SourceCdnReqAuthenticationType{
			value: "QINIU_PRIVATE_AUTHENTICATION",
		},
		ALIYUN_OSS_A: SourceCdnReqAuthenticationType{
			value: "ALIYUN_OSS_A",
		},
		ALIYUN_OSS_B: SourceCdnReqAuthenticationType{
			value: "ALIYUN_OSS_B",
		},
		ALIYUN_OSS_C: SourceCdnReqAuthenticationType{
			value: "ALIYUN_OSS_C",
		},
		KSYUN_PRIVATE_AUTHENTICATION: SourceCdnReqAuthenticationType{
			value: "KSYUN_PRIVATE_AUTHENTICATION",
		},
		AZURE_SAS_TOKEN: SourceCdnReqAuthenticationType{
			value: "AZURE_SAS_TOKEN",
		},
		TENCENT_COS_A: SourceCdnReqAuthenticationType{
			value: "TENCENT_COS_A",
		},
		TENCENT_COS_B: SourceCdnReqAuthenticationType{
			value: "TENCENT_COS_B",
		},
		TENCENT_COS_C: SourceCdnReqAuthenticationType{
			value: "TENCENT_COS_C",
		},
		TENCENT_COS_D: SourceCdnReqAuthenticationType{
			value: "TENCENT_COS_D",
		},
	}
}

func (c SourceCdnReqAuthenticationType) Value() string {
	return c.value
}

func (c SourceCdnReqAuthenticationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceCdnReqAuthenticationType) UnmarshalJSON(b []byte) error {
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

type SourceCdnReqProtocol struct {
	value string
}

type SourceCdnReqProtocolEnum struct {
	HTTP  SourceCdnReqProtocol
	HTTPS SourceCdnReqProtocol
}

func GetSourceCdnReqProtocolEnum() SourceCdnReqProtocolEnum {
	return SourceCdnReqProtocolEnum{
		HTTP: SourceCdnReqProtocol{
			value: "http",
		},
		HTTPS: SourceCdnReqProtocol{
			value: "https",
		},
	}
}

func (c SourceCdnReqProtocol) Value() string {
	return c.value
}

func (c SourceCdnReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceCdnReqProtocol) UnmarshalJSON(b []byte) error {
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
