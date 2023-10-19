package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCatalogRequest Request Object
type CreateCatalogRequest struct {

	// LakeFormation实例ID。创建实例时自动生成。例如：2180518f-42b8-4947-b20b-adfc53981a25。
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
