package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCatalogRequest Request Object
type CreateCatalogRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *BizCatalogVo `json:"body,omitempty"`
}

func (o CreateCatalogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCatalogRequest struct{}"
	}

	return strings.Join([]string{"CreateCatalogRequest", string(data)}, " ")
}
