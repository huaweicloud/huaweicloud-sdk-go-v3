package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifySupabaseApplicationPasswordRequestBody 修改密码请求。application_password和instance_password可单独传，但不能同时为空。
type ModifySupabaseApplicationPasswordRequestBody struct {

	// Supabase应用登录密码。
	ApplicationPassword *string `json:"application_password,omitempty"`

	// 纳管实例数据库密码。
	InstancePassword *string `json:"instance_password,omitempty"`
}

func (o ModifySupabaseApplicationPasswordRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifySupabaseApplicationPasswordRequestBody struct{}"
	}

	return strings.Join([]string{"ModifySupabaseApplicationPasswordRequestBody", string(data)}, " ")
}
