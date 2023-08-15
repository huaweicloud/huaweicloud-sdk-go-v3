package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiCatalogUpdateParaDto struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`
}

func (o ApiCatalogUpdateParaDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiCatalogUpdateParaDto struct{}"
	}

	return strings.Join([]string{"ApiCatalogUpdateParaDto", string(data)}, " ")
}
