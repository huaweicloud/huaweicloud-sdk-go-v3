package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCertificateAuthorityResponse Response Object
type ShowCertificateAuthorityResponse struct {

	// CA证书ID。
	CaId *string `json:"ca_id,omitempty"`

	// CA类型:   - **ROOT**: 根CA   - **SUBORDINATE**: 从属CA
	Type *string `json:"type,omitempty"`

	// CA证书状态：   - **PENDING** : 待激活，此状态下，不可用于签发证书；   - **ACTIVED** : 已激活，此状态下，可用于签发证书；   - **DISABLED** : 已禁用，此状态下，不可用于签发证书；   - **DELETED** : 计划删除，此状态下，不可用于签发证书；   - **EXPIRED** : 已过期，此状态下，不可用于签发证书。
	Status *string `json:"status,omitempty"`

	// CA路径长度。 > 注：生成的根CA证书，其路径长度不做限制，但本字段在数据库中统一置为7。从属CA的路径长度在创建时由用户指定，缺省值为0。
	PathLength *int32 `json:"path_length,omitempty"`

	// 父CA证书ID，即签发此证书的CA证书ID。根CA中，此参数为**null**。
	IssuerId *string `json:"issuer_id,omitempty"`

	// 父CA证书名称。根CA中，此参数为**null**。
	IssuerName *string `json:"issuer_name,omitempty"`

	// 密钥算法。
	KeyAlgorithm *string `json:"key_algorithm,omitempty"`

	// 签名哈希算法。
	SignatureAlgorithm *string `json:"signature_algorithm,omitempty"`

	// 冻结标识:   - **0** : 非冻结状态；   - **其它值** : 冻结状态，当前预留。
	FreezeFlag *int32 `json:"freeze_flag,omitempty"`

	// 证书生成方式：  - **GENERATE** : PCA系统生成；  - **IMPORT** : 外部导入；  - **CSR** : 外部提供CSR，内部CA进行签发，即私钥不在PCA进行托管。
	GenMode *string `json:"gen_mode,omitempty"`

	// 证书序列号。
	SerialNumber *string `json:"serial_number,omitempty"`

	// 证书创建时间，格式为时间戳（毫秒级）。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 证书删除时间，格式为时间戳（毫秒级）。
	DeleteTime *int64 `json:"delete_time,omitempty"`

	// 证书创建时间，格式为时间戳（毫秒级）。
	NotBefore *int64 `json:"not_before,omitempty"`

	// 证书到期时间，格式为时间戳（毫秒级）。
	NotAfter *int64 `json:"not_after,omitempty"`

	DistinguishedName *DistinguishedName `json:"distinguished_name,omitempty"`

	CrlConfiguration *ListCrlConfiguration `json:"crl_configuration,omitempty"`
	HttpStatusCode   int                   `json:"-"`
}

func (o ShowCertificateAuthorityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertificateAuthorityResponse struct{}"
	}

	return strings.Join([]string{"ShowCertificateAuthorityResponse", string(data)}, " ")
}
