package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCatalogTreeRequest Request Object
type ListCatalogTreeRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`
}

func (o ListCatalogTreeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCatalogTreeRequest struct{}"
	}

	return strings.Join([]string{"ListCatalogTreeRequest", string(data)}, " ")
}
