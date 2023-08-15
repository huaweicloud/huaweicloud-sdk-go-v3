package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrateCatalogRequest Request Object
type MigrateCatalogRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 目录编号
	CatalogId string `json:"catalog_id"`

	Body *CatalogMoveParaDto `json:"body,omitempty"`
}

func (o MigrateCatalogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateCatalogRequest struct{}"
	}

	return strings.Join([]string{"MigrateCatalogRequest", string(data)}, " ")
}
