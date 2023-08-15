package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiCatalogCreateParaDto struct {

	// 父目录编号
	Pid *string `json:"pid,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`
}

func (o ApiCatalogCreateParaDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiCatalogCreateParaDto struct{}"
	}

	return strings.Join([]string{"ApiCatalogCreateParaDto", string(data)}, " ")
}
