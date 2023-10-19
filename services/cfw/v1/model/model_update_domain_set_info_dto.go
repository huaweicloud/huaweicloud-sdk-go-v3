package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDomainSetInfoDto struct {

	// 域名组名称UUID
	Name string `json:"name"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 域名组id
	SetId *string `json:"set_id,omitempty"`

	// 域名组类型，0表示URL过滤，1表示地址解析
	DomainSetType *int32 `json:"domain_set_type,omitempty"`
}

func (o UpdateDomainSetInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainSetInfoDto struct{}"
	}

	return strings.Join([]string{"UpdateDomainSetInfoDto", string(data)}, " ")
}
