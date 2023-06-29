package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCatalogRequest Request Object
type UpdateCatalogRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 目录编号
	CatalogId string `json:"catalog_id"`

	Body *ApiCatalogUpdateParaDto `json:"body,omitempty"`
}

func (o UpdateCatalogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCatalogRequest struct{}"
	}

	return strings.Join([]string{"UpdateCatalogRequest", string(data)}, " ")
}
