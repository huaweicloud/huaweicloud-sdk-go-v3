package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceTenant struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 账号名称
	DomainName string `json:"domain_name"`

	// 创建的委托类型
	Type string `json:"type"`
}

func (o ResourceTenant) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTenant struct{}"
	}

	return strings.Join([]string{"ResourceTenant", string(data)}, " ")
}
