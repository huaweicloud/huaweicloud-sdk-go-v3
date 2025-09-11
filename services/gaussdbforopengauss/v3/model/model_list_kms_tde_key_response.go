package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListKmsTdeKeyResponse Response Object
type ListKmsTdeKeyResponse struct {

	// **参数解释**: kms秘钥列表。
	KeyDetails *[]ListKmsTdeKeyResponseBodyKeyDetails `json:"key_details,omitempty"`

	// **参数解释**: 授权用户ID，在开启透明加密能力时，必须对当前用户ID进行授权。 授权操作参考创建授权接口 https://support.huaweicloud.com/api-dew/CreateGrant.html。 **取值范围**: 不涉及。
	AuthorizedId *string `json:"authorized_id,omitempty"`

	// **参数解释**: 授权用户名称。 **取值范围**: 不涉及。
	AuthorizedName *string `json:"authorized_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListKmsTdeKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKmsTdeKeyResponse struct{}"
	}

	return strings.Join([]string{"ListKmsTdeKeyResponse", string(data)}, " ")
}
