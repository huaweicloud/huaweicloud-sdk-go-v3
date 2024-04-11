package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CatalogAttributeVoParent 父节点。
type CatalogAttributeVoParent struct {

	// 父节点ID。
	Id *string `json:"id,omitempty"`
}

func (o CatalogAttributeVoParent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CatalogAttributeVoParent struct{}"
	}

	return strings.Join([]string{"CatalogAttributeVoParent", string(data)}, " ")
}
