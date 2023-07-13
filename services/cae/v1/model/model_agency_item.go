package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AgencyItem struct {

	// 委托名称，固定值“cae_trust”，该值不可修改。
	Name *string `json:"name,omitempty"`
}

func (o AgencyItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyItem struct{}"
	}

	return strings.Join([]string{"AgencyItem", string(data)}, " ")
}
