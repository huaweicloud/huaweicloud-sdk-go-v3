package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnterprisePersonNew struct {

	// 法人姓名。
	LegelName string `json:"legel_name"`

	// 法人身份证号。
	LegelIdNumber string `json:"legel_id_number"`

	// 认证人角色。 legalPerson ：法人代表。
	CertifierRole *string `json:"certifier_role,omitempty"`
}

func (o EnterprisePersonNew) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnterprisePersonNew struct{}"
	}

	return strings.Join([]string{"EnterprisePersonNew", string(data)}, " ")
}
