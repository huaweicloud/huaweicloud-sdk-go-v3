package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCatalogDetailRequest Request Object
type ShowCatalogDetailRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 目录编号
	CatalogId string `json:"catalog_id"`
}

func (o ShowCatalogDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCatalogDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowCatalogDetailRequest", string(data)}, " ")
}
