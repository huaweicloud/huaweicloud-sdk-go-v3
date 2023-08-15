package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApiCatalogListRequest Request Object
type ListApiCatalogListRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 目录编号
	CatalogId string `json:"catalog_id"`

	// 查询起始坐标, 即跳过前X条数据
	Offset *int32 `json:"offset,omitempty"`

	// 查询条数, 即查询Y条数据
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListApiCatalogListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiCatalogListRequest struct{}"
	}

	return strings.Join([]string{"ListApiCatalogListRequest", string(data)}, " ")
}
