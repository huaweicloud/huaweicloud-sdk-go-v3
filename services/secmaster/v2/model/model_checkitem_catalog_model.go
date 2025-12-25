package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckitemCatalogModel struct {

	// 检查项uuid
	Uuid string `json:"uuid"`

	// 检查项的名称
	Name string `json:"name"`
}

func (o CheckitemCatalogModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckitemCatalogModel struct{}"
	}

	return strings.Join([]string{"CheckitemCatalogModel", string(data)}, " ")
}
