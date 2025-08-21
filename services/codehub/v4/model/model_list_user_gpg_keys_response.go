package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserGpgKeysResponse Response Object
type ListUserGpgKeysResponse struct {

	// **参数解释：** 记录id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 创建时间。 **取值范围：** 字符串长度不少于1，不超过1000。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 解析到的邮箱列表以及是否生效。
	EmailsWithVerifiedStatus map[string]bool `json:"emails_with_verified_status,omitempty"`

	// **参数解释：** 主密/公钥的指纹格式。 **取值范围：** 字符串长度不少于1，不超过1000。
	Fingerprint *string `json:"fingerprint,omitempty"`

	// **参数解释：** gpg的公钥。 **取值范围：** 字符串长度不少于1，不超过1000。
	Key *string `json:"key,omitempty"`

	// **参数解释：** 描述。 **取值范围：** 字符串长度不少于1，不超过1000。
	Description *string `json:"description,omitempty"`

	// **参数解释：** gpg_key的标题。 **取值范围：** 字符串长度不少于1，不超过1000。
	Title *string `json:"title,omitempty"`

	// **参数解释：** 主密/公钥primary_key的id。 **取值范围：** 字符串长度不少于1，不超过1000。
	PrimaryKeyid *string `json:"primary_keyid,omitempty"`

	// **参数解释：** 是否生效。
	Active *bool `json:"active,omitempty"`

	// **参数解释：** 子钥列表。 **取值范围：** 列表长度不少于0，不超过1000。
	Subkeys        *[]GpgSubKeyDto `json:"subkeys,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListUserGpgKeysResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserGpgKeysResponse struct{}"
	}

	return strings.Join([]string{"ListUserGpgKeysResponse", string(data)}, " ")
}
