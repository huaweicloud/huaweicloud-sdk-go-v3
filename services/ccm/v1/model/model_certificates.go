package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Certificates struct {

	// 私有证书ID。
	CertificateId string `json:"certificate_id"`

	// 证书状态：   - **ISSUED** : 已签发；   - **EXPIRED** : 已过期；   - **REVOKED** : 已吊销。
	Status string `json:"status"`

	// 父CA证书ID。
	IssuerId string `json:"issuer_id"`

	// 父CA证书名称。
	IssuerName string `json:"issuer_name"`

	// 密钥算法。
	KeyAlgorithm string `json:"key_algorithm"`

	// 签名算法。
	SignatureAlgorithm string `json:"signature_algorithm"`

	// 冻结标识:   - **0** : 非冻结状态；   - **其它值** : 冻结状态，当前预留。
	FreezeFlag int32 `json:"freeze_flag"`

	// 证书生成方式：  - **GENERATE** : PCA系统生成；  - **IMPORT** : 外部导入；  - **CSR** : 外部提供CSR，内部CA进行签发，即私钥不在PCA进行托管。
	GenMode string `json:"gen_mode"`

	// 序列号。
	SerialNumber string `json:"serial_number"`

	// 证书创建时间，格式为时间戳（毫秒级）。
	CreateTime int64 `json:"create_time"`

	// 证书删除时间，格式为时间戳（毫秒级）。
	DeleteTime int64 `json:"delete_time"`

	// 证书创建时间，格式为时间戳（毫秒级）。
	NotBefore int64 `json:"not_before"`

	// 证书到期时间，格式为时间戳（毫秒级）。
	NotAfter int64 `json:"not_after"`

	DistinguishedName *DistinguishedName `json:"distinguished_name"`

	EncCertInfo *EncCertInfo `json:"enc_cert_info,omitempty"`
}

func (o Certificates) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Certificates struct{}"
	}

	return strings.Join([]string{"Certificates", string(data)}, " ")
}
