package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiCatalogDeleteParaDto struct {
	Ids *[]string `json:"ids,omitempty"`
}

func (o ApiCatalogDeleteParaDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiCatalogDeleteParaDto struct{}"
	}

	return strings.Join([]string{"ApiCatalogDeleteParaDto", string(data)}, " ")
}
