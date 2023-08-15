package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeCatalogRequest Request Object
type ChangeCatalogRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *BizCatalogVo `json:"body,omitempty"`
}

func (o ChangeCatalogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeCatalogRequest struct{}"
	}

	return strings.Join([]string{"ChangeCatalogRequest", string(data)}, " ")
}
