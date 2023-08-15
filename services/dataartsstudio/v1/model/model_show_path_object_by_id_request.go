package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPathObjectByIdRequest Request Object
type ShowPathObjectByIdRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 目录编号
	CatalogId string `json:"catalog_id"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowPathObjectByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPathObjectByIdRequest struct{}"
	}

	return strings.Join([]string{"ShowPathObjectByIdRequest", string(data)}, " ")
}
