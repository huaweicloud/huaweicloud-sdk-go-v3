package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GpgSubKeyDto struct {

	// **参数解释：** 子公钥subkey的id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 子公钥的指纹格式。 **取值范围：** 字符串长度不少于1，不超过1000。
	Fingerprint *string `json:"fingerprint,omitempty"`

	// **参数解释：** gpg_key的id。
	GpgKeyId *int32 `json:"gpg_key_id,omitempty"`

	// **参数解释：** 子秘钥的id。 **取值范围：** 字符串长度不少于1，不超过1000。
	Keyid *string `json:"keyid,omitempty"`
}

func (o GpgSubKeyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GpgSubKeyDto struct{}"
	}

	return strings.Join([]string{"GpgSubKeyDto", string(data)}, " ")
}
