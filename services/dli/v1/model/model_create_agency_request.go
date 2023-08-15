package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAgencyRequest
type CreateAgencyRequest struct {

	// 角色目前只支持，obs_adm、dis_adm、ctable_adm、vpc_netadm、smn_adm、te_admin
	Roles []string `json:"roles"`
}

func (o CreateAgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAgencyRequest struct{}"
	}

	return strings.Join([]string{"CreateAgencyRequest", string(data)}, " ")
}
