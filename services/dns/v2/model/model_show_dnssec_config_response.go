package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDnssecConfigResponse Response Object
type ShowDnssecConfigResponse struct {

	// **参数解释：** 域名。 **取值范围：** 不涉及。
	ZoneName *string `json:"zone_name,omitempty"`

	// **参数解释：** 密钥标签。 **取值范围：** 不涉及。
	KeyTag *int32 `json:"key_tag,omitempty"`

	// **参数解释：** 旗标。 **取值范围：** 不涉及。
	Flag *int32 `json:"flag,omitempty"`

	// **参数解释：** 摘要算法。 **取值范围：** 不涉及。
	DigestAlgorithm *string `json:"digest_algorithm,omitempty"`

	// **参数解释：** 摘要算法类型。 **取值范围：** 不涉及。
	DigestType *int32 `json:"digest_type,omitempty"`

	// **参数解释：** 摘要。 **取值范围：** 不涉及。
	Digest *string `json:"digest,omitempty"`

	// **参数解释：** 签名算法。 **取值范围：** 不涉及。
	Signature *string `json:"signature,omitempty"`

	// **参数解释：** 签名算法类型。 **取值范围：** 不涉及。
	SignatureType *int32 `json:"signature_type,omitempty"`

	// **参数解释：** 公有密钥。 **取值范围：** 不涉及。
	KskPublicKey *string `json:"ksk_public_key,omitempty"`

	// **参数解释：** DS记录。 **取值范围：** 不涉及。
	DsRecord *string `json:"ds_record,omitempty"`

	// **参数解释：** 创建时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。 **取值范围：** 不涉及。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 更新时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。 **取值范围：** 不涉及。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释：** 状态。 **取值范围：** - ENABLE：启用 - DISABLE：关闭
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDnssecConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDnssecConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowDnssecConfigResponse", string(data)}, " ")
}
