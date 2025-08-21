package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddSshKeyResponse Response Object
type AddSshKeyResponse struct {

	// **参数解释：** 秘钥id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 秘钥名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	Title *string `json:"title,omitempty"`

	// **参数解释：** 公钥。 **取值范围：** 字符串长度不少于1，不超过1000。
	Key *string `json:"key,omitempty"`

	// **参数解释：** 创建时间。 **取值范围：** 字符串长度不少于1，不超过1000。
	CreatedAt      *string `json:"created_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddSshKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSshKeyResponse struct{}"
	}

	return strings.Join([]string{"AddSshKeyResponse", string(data)}, " ")
}
