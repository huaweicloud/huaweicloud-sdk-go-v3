package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgencyMetadata 请求数据。
type AgencyMetadata struct {

	// 委托名称，固定值“cae_trust”，该值不可修改。
	Name string `json:"name"`
}

func (o AgencyMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyMetadata struct{}"
	}

	return strings.Join([]string{"AgencyMetadata", string(data)}, " ")
}
