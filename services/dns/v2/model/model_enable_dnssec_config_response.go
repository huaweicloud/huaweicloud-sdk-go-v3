package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableDnssecConfigResponse Response Object
type EnableDnssecConfigResponse struct {

	// 域名。
	ZoneName *string `json:"zone_name,omitempty"`

	// 密钥标签。
	KeyTag *int32 `json:"key_tag,omitempty"`

	// 旗标。
	Flag *int32 `json:"flag,omitempty"`

	// 摘要算法。
	DigestAlgorithm *string `json:"digest_algorithm,omitempty"`

	// 摘要算法类型。
	DigestType *int32 `json:"digest_type,omitempty"`

	// 摘要。
	Digest *string `json:"digest,omitempty"`

	// 签名算法。
	Signature *string `json:"signature,omitempty"`

	// 签名算法类型。
	SignatureType *int32 `json:"signature_type,omitempty"`

	// 公有密钥。
	KskPublicKey *string `json:"ksk_public_key,omitempty"`

	// DS记录。
	DsRecord *string `json:"ds_record,omitempty"`

	// 创建时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 状态。  取值范围：  ENABLE：启用 DISABLE：关闭
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o EnableDnssecConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableDnssecConfigResponse struct{}"
	}

	return strings.Join([]string{"EnableDnssecConfigResponse", string(data)}, " ")
}
