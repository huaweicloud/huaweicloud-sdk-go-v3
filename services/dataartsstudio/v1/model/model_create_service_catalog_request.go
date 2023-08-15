package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateServiceCatalogRequest Request Object
type CreateServiceCatalogRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	Body *ApiCatalogCreateParaDto `json:"body,omitempty"`
}

func (o CreateServiceCatalogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServiceCatalogRequest struct{}"
	}

	return strings.Join([]string{"CreateServiceCatalogRequest", string(data)}, " ")
}
