package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchApplicationResponse Response Object
type SearchApplicationResponse struct {

	// 组件列表。
	AppInfoList *[]AppInfo `json:"app_info_list,omitempty"`

	// 组件总数目。
	AppTotalCount *int32 `json:"app_total_count,omitempty"`

	// 组件名称和组件详情map表。
	AppInfoMap     map[string]AppInfo `json:"app_info_map,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o SearchApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchApplicationResponse struct{}"
	}

	return strings.Join([]string{"SearchApplicationResponse", string(data)}, " ")
}
