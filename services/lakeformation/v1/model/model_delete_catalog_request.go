package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCatalogRequest Request Object
type DeleteCatalogRequest struct {

	// 实例Id
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`
}

func (o DeleteCatalogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCatalogRequest struct{}"
	}

	return strings.Join([]string{"DeleteCatalogRequest", string(data)}, " ")
}
