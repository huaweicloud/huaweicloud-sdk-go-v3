package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCatalogRequest Request Object
type CreateCatalogRequest struct {

	// 实例Id
	InstanceId string `json:"instance_id"`

	Body *CatalogInput `json:"body,omitempty"`
}

func (o CreateCatalogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCatalogRequest struct{}"
	}

	return strings.Join([]string{"CreateCatalogRequest", string(data)}, " ")
}
