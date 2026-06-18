package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTenantKmsGrantResponse Response Object
type ShowTenantKmsGrantResponse struct {

	// **参数解释：** 租户id。 **取值范围：** 字符串长度不少于1，不超过1000。
	TenantId *string `json:"tenant_id,omitempty"`

	// **参数解释：** 是否有委托和授权。
	Assumed        *bool `json:"assumed,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowTenantKmsGrantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTenantKmsGrantResponse struct{}"
	}

	return strings.Join([]string{"ShowTenantKmsGrantResponse", string(data)}, " ")
}
