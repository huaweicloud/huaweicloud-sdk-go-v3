package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SearchApplicationResponse struct {

	// 应用列表
	AppInfoList *[]AppInfo `json:"app_info_list,omitempty" xml:"app_info_list"`

	// 应用总数目
	AppTotalCount *int32 `json:"app_total_count,omitempty" xml:"app_total_count"`

	// 应用名称和应用详情map表
	AppInfoMap     map[string]AppInfo `json:"app_info_map,omitempty" xml:"app_info_map"`
	HttpStatusCode int                `json:"-"`
}

func (o SearchApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchApplicationResponse struct{}"
	}

	return strings.Join([]string{"SearchApplicationResponse", string(data)}, " ")
}
