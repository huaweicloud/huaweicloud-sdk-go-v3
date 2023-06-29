package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBizCatalogDetailRequest Request Object
type ShowBizCatalogDetailRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 实体id
	Id string `json:"id"`
}

func (o ShowBizCatalogDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBizCatalogDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowBizCatalogDetailRequest", string(data)}, " ")
}
