package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VerifyDomainOwnerResponse Response Object
type VerifyDomainOwnerResponse struct {

	// 验证是否通过，true：通过，false：不通过。
	Result *bool `json:"result,omitempty"`

	// **参数解释：** 域名校验状态，表示当前账号下是否存在该域名 **约束限制：** 不涉及 **取值范围：** - DomainExist: 该域名已存在 - DomainNotExist: 该域名不存在  **默认取值：** 不涉及
	Status *string `json:"status,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o VerifyDomainOwnerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VerifyDomainOwnerResponse struct{}"
	}

	return strings.Join([]string{"VerifyDomainOwnerResponse", string(data)}, " ")
}
