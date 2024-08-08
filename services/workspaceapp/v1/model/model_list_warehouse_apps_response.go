package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWarehouseAppsResponse Response Object
type ListWarehouseAppsResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 应用仓库中的应用列表。
	Items          *[]WarehouseApp `json:"items,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListWarehouseAppsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWarehouseAppsResponse struct{}"
	}

	return strings.Join([]string{"ListWarehouseAppsResponse", string(data)}, " ")
}
