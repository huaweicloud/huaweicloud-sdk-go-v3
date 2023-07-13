package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServiceCatalogRequest Request Object
type DeleteServiceCatalogRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	Body *ApiCatalogDeleteParaDto `json:"body,omitempty"`
}

func (o DeleteServiceCatalogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServiceCatalogRequest struct{}"
	}

	return strings.Join([]string{"DeleteServiceCatalogRequest", string(data)}, " ")
}
