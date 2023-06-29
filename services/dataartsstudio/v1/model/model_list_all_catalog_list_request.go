package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllCatalogListRequest Request Object
type ListAllCatalogListRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 目录编号
	CatalogId string `json:"catalog_id"`

	// 查询起始坐标, 即跳过前X条数据
	Offset *int32 `json:"offset,omitempty"`

	// 查询条数, 即查询Y条数据
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAllCatalogListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllCatalogListRequest struct{}"
	}

	return strings.Join([]string{"ListAllCatalogListRequest", string(data)}, " ")
}
