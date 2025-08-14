package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCatalogApplicationsResponse Response Object
type ListCatalogApplicationsResponse struct {

	// 应用程序目录中的应用程序列表
	Applications *[]ApplicationTemplateDisplayDto `json:"applications,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListCatalogApplicationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCatalogApplicationsResponse struct{}"
	}

	return strings.Join([]string{"ListCatalogApplicationsResponse", string(data)}, " ")
}
