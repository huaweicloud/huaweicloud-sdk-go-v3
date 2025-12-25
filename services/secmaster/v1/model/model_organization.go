package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Organization struct {

	// 角色id
	AccountId string `json:"account_id"`

	// 角色id
	OrganizationId string `json:"organization_id"`
}

func (o Organization) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Organization struct{}"
	}

	return strings.Join([]string{"Organization", string(data)}, " ")
}
