package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateAdminResponse Response Object
type ValidateAdminResponse struct {

	// 是否是管理员账号
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
