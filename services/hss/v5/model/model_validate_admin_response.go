package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateAdminResponse Response Object
type ValidateAdminResponse struct {

	// **参数解释**： 当前账号是否是管理员账号 **取值范围**： - true：是管理员账号 - false：不是管理员账号
	IsAdminAccount *bool `json:"is_admin_account,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ValidateAdminResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateAdminResponse struct{}"
	}

	return strings.Join([]string{"ValidateAdminResponse", string(data)}, " ")
}
