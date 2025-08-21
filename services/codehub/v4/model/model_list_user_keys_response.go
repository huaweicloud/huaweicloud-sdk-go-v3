package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserKeysResponse Response Object
type ListUserKeysResponse struct {

	// **参数解释：** 密钥列表。 **取值范围：** 列表长度不少于0，不超过1000。
	Keys *[]SshKeyDto `json:"keys,omitempty"`

	// **参数解释：** 结果条数。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListUserKeysResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserKeysResponse struct{}"
	}

	return strings.Join([]string{"ListUserKeysResponse", string(data)}, " ")
}
