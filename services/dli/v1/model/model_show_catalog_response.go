package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCatalogResponse Response Object
type ShowCatalogResponse struct {

	// 是否成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// DLI侧catalog映射名称。
	Name *string `json:"name,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	Parameters map[string]string `json:"parameters,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// catalog状态。CREATING：catalog创建中；ACTIVE：catalog可使用；FAILED：catalog创建失败。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCatalogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCatalogResponse struct{}"
	}

	return strings.Join([]string{"ShowCatalogResponse", string(data)}, " ")
}
