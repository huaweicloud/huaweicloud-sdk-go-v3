package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPathByIdRequest Request Object
type ShowPathByIdRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 目录编号
	CatalogId string `json:"catalog_id"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowPathByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPathByIdRequest struct{}"
	}

	return strings.Join([]string{"ShowPathByIdRequest", string(data)}, " ")
}
