package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppsResponse Response Object
type ListAppsResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 应用列表
	Apps           *[]AppReturnDto `json:"apps,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListAppsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppsResponse struct{}"
	}

	return strings.Join([]string{"ListAppsResponse", string(data)}, " ")
}
