package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateServiceAgencyRequestBody struct {

	// OU列表
	Organizations *[]Organization `json:"organizations,omitempty"`

	// 角色列表
	RoleDescriptions *[]Role `json:"role_descriptions,omitempty"`
}

func (o CreateServiceAgencyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServiceAgencyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateServiceAgencyRequestBody", string(data)}, " ")
}
