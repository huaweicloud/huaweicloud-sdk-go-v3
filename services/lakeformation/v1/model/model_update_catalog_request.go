package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCatalogRequest Request Object
type UpdateCatalogRequest struct {

	// 实例Id
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	Body *CatalogInput `json:"body,omitempty"`
}

func (o UpdateCatalogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCatalogRequest struct{}"
	}

	return strings.Join([]string{"UpdateCatalogRequest", string(data)}, " ")
}
