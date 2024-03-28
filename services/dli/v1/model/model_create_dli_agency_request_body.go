package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDliAgencyRequestBody struct {

	// 角色目前只支持，obs_adm、dis_adm、ctable_adm、vpc_netadm、smn_adm、te_admin
	Roles []string `json:"roles"`
}

func (o CreateDliAgencyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDliAgencyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDliAgencyRequestBody", string(data)}, " ")
}
