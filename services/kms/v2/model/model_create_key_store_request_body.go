package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateKeyStoreRequestBody 创建专属密钥库请求体
type CreateKeyStoreRequestBody struct {

	// 专属密钥库别名，取值范围为1到255个字符，满足正则匹配“^[a-zA-Z0-9:/_-]{1,255}$”，且不与已有的专属密钥库别名重名。
	KeystoreAlias string `json:"keystore_alias"`

	// DHSM集群Id，要求集群当前未创建专属密钥库。
	HsmClusterId *string `json:"hsm_cluster_id,omitempty"`

	// DHSM集群的CA证书
	HsmCaCert *string `json:"hsm_ca_cert,omitempty"`

	// 集群ID。当类型为DHSM时，cluster_id为hsm_cluster_id。当类型为CDMS时，为cdms_cluster_id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 专属密钥库集群类型。DHSM表示专属加密服务集群，CDMS表示密码卡集群,DEFAULT表示KMS原有集群
	KeystoreType *CreateKeyStoreRequestBodyKeystoreType `json:"keystore_type,omitempty"`

	// **参数解释：** 仅XKS类型的密钥库支持；可选，表示Xks-Proxy的域名信息 **约束限制：** ^https://[^\\s]+ **取值范围：** 不涉及 **默认取值：** 不涉及
	ProxyUriEndpoint *string `json:"proxy_uri_endpoint,omitempty"`

	// **参数解释：** 仅XKS类型的密钥库支持；必选，Xks-Proxy的路径前缀 **约束限制：** ^/(?:[a-zA-Z0-9/_-]{0,127})$ **取值范围：** 不涉及 **默认取值：** 不涉及
	ProxyUriPath *string `json:"proxy_uri_path,omitempty"`

	// **参数解释：** 仅XKS类型的密钥库支持；必选，访问XKS-Proxy的凭据ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ProxyAccessKeyId *string `json:"proxy_access_key_id,omitempty"`

	// **参数解释：** 仅XKS类型的密钥库支持；必选，访问XKS-Proxy的凭据Key **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ProxySecretAccessKey *string `json:"proxy_secret_access_key,omitempty"`

	// **参数解释：** 仅XKS类型的密钥库支持；必选，Xks-Proxy的终端节点服务名称 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ProxyVpcEndpointServiceName *string `json:"proxy_vpc_endpoint_service_name,omitempty"`

	// **参数解释：** 仅XKS类型的密钥库支持；必选，终端节点服务所属账号 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ProxyVpcEndpointServiceOwner *string `json:"proxy_vpc_endpoint_service_owner,omitempty"`
}

func (o CreateKeyStoreRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKeyStoreRequestBody struct{}"
	}

	return strings.Join([]string{"CreateKeyStoreRequestBody", string(data)}, " ")
}

type CreateKeyStoreRequestBodyKeystoreType struct {
	value string
}

type CreateKeyStoreRequestBodyKeystoreTypeEnum struct {
	DEFAULT CreateKeyStoreRequestBodyKeystoreType
	DHSM    CreateKeyStoreRequestBodyKeystoreType
	CDMS    CreateKeyStoreRequestBodyKeystoreType
	XKS     CreateKeyStoreRequestBodyKeystoreType
}

func GetCreateKeyStoreRequestBodyKeystoreTypeEnum() CreateKeyStoreRequestBodyKeystoreTypeEnum {
	return CreateKeyStoreRequestBodyKeystoreTypeEnum{
		DEFAULT: CreateKeyStoreRequestBodyKeystoreType{
			value: "DEFAULT",
		},
		DHSM: CreateKeyStoreRequestBodyKeystoreType{
			value: "DHSM",
		},
		CDMS: CreateKeyStoreRequestBodyKeystoreType{
			value: "CDMS",
		},
		XKS: CreateKeyStoreRequestBodyKeystoreType{
			value: "XKS",
		},
	}
}

func (c CreateKeyStoreRequestBodyKeystoreType) Value() string {
	return c.value
}

func (c CreateKeyStoreRequestBodyKeystoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateKeyStoreRequestBodyKeystoreType) UnmarshalJSON(b []byte) error {
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
