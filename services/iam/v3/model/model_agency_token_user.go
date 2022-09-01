package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type AgencyTokenUser struct {

	// 委托方A账号名/委托名。
	Name string `json:"name" xml:"name"`

	// 委托ID。
	Id string `json:"id" xml:"id"`

	Domain *AgencyTokenUserDomain `json:"domain" xml:"domain"`
}

func (o AgencyTokenUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyTokenUser struct{}"
	}

	return strings.Join([]string{"AgencyTokenUser", string(data)}, " ")
}
