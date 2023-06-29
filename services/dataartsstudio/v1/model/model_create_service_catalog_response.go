package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateServiceCatalogResponse Response Object
type CreateServiceCatalogResponse struct {

	// 目录编号
	CatalogId *string `json:"catalog_id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 创建者
	CreateUser *string `json:"create_user,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 更新者
	UpdateUser     *string `json:"update_user,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateServiceCatalogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServiceCatalogResponse struct{}"
	}

	return strings.Join([]string{"CreateServiceCatalogResponse", string(data)}, " ")
}
