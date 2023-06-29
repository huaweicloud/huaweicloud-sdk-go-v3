package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCatalogRequest Request Object
type DeleteCatalogRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *IdsParam `json:"body,omitempty"`
}

func (o DeleteCatalogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCatalogRequest struct{}"
	}

	return strings.Join([]string{"DeleteCatalogRequest", string(data)}, " ")
}
