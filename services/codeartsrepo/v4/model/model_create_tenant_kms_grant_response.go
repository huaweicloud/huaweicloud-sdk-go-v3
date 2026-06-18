package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTenantKmsGrantResponse Response Object
type CreateTenantKmsGrantResponse struct {

	// **参数解释：** 租户id。 **取值范围：** 字符串长度不少于1，不超过1000。
	TenantId *string `json:"tenant_id,omitempty"`

	// **参数解释：** 是否有委托和授权。
	Assumed        *bool `json:"assumed,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CreateTenantKmsGrantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTenantKmsGrantResponse struct{}"
	}

	return strings.Join([]string{"CreateTenantKmsGrantResponse", string(data)}, " ")
}
